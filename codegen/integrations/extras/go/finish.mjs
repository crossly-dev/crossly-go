#!/usr/bin/env node
/**
 * Finish one generated Go client after assemble.mjs: the part generation does not do.
 *
 *   node codegen/integrations/extras/go/finish.mjs <seller|buyer>      (run by codegen/ci/check-go.sh)
 *
 *   1. module path: the 5.4.0 template writes `module github.com/GIT_USER_ID/GIT_REPO_ID`. Each client is its own
 *      module in a subdirectory, so it must be github.com/crossly-dev/crossly-go/<dir> (tags <dir>/vX.Y.Z), and
 *      the README/doc imports follow. A no-op once generate-sdks.mjs passes gitUserId/gitRepoId for go.
 *   2. go 1.23 in go.mod: the ergonomics below use generics and iter.Seq2.
 *   3. vendor: client ergonomics into the package root (package clause rewritten for buyer), the webhook verifier
 *      plus its net/http middleware into <dir>/webhooks, and (seller) the Gin and Echo adapters as their own
 *      modules under webhooks/, so the core module never depends on a web framework.
 *   4. version: SDKVersion from codegen/state.json; the generated default User-Agent becomes crossly-go/<version>.
 *
 * Idempotent: ci.yml runs it again over the committed tree and must find nothing to change.
 */
import { execFileSync } from 'node:child_process';
import { copyFileSync, existsSync, mkdirSync, readdirSync, readFileSync, statSync, writeFileSync } from 'node:fs';
import { basename, dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const HERE = dirname(fileURLToPath(import.meta.url));
const CODEGEN = resolve(HERE, '..', '..', '..');
const dir = resolve(process.argv[2] ?? '');
const api = basename(dir);
if (!['seller', 'buyer'].includes(api) || !existsSync(join(dir, 'go.mod'))) {
  console.error(`finish.mjs: ${dir} is not a generated Go client (seller/ or buyer/ with a go.mod)`);
  process.exit(2);
}
const MODULE = `github.com/crossly-dev/crossly-go/${api}`;
const version = JSON.parse(readFileSync(join(CODEGEN, 'state.json'), 'utf8')).version;
const go = (cwd, ...args) => execFileSync('go', args, { cwd, stdio: ['ignore', 'inherit', 'inherit'] });

const write = (file, text) => {
  mkdirSync(dirname(file), { recursive: true });
  if (!existsSync(file) || readFileSync(file, 'utf8') !== text) writeFileSync(file, text);
};

// 1 + 2. module path and go version.
go(dir, 'mod', 'edit', `-module=${MODULE}`, '-go=1.23');
const walk = (d) =>
  readdirSync(d).flatMap((n) => {
    const p = join(d, n);
    return statSync(p).isDirectory() ? walk(p) : [p];
  });
for (const f of walk(dir).filter((p) => /\.(md|sh)$/.test(p))) {
  const t = readFileSync(f, 'utf8');
  const fixed = t
    .replaceAll('github.com/GIT_USER_ID/GIT_REPO_ID', MODULE)
    .replaceAll('GIT_USER_ID', 'crossly-dev')
    .replaceAll('GIT_REPO_ID', `crossly-go/${api}`);
  if (fixed !== t) writeFileSync(f, fixed);
}

// 1b. An operation with two tags is written into BOTH services' files, each declaring the same
// `Api<Op>Request` type (with a different ApiService field), which does not compile. Keep the first file's
// name and prefix the others with their service (`InventoryApiGetAnalyticInventoryRequest`). A no-op once the
// generator keeps one tag per operation for Go.
const apiFiles = readdirSync(dir).filter((n) => /^api_.*\.go$/.test(n)).sort();
const owner = new Map();
let renamed = 0;
for (const f of apiFiles) {
  const p = join(dir, f);
  let t = readFileSync(p, 'utf8');
  const service = t.match(/^type (\w+)ApiService service$/m)?.[1];
  for (const [, name] of [...t.matchAll(/^type (Api\w+Request) struct/gm)]) {
    if (!owner.has(name)) {
      owner.set(name, f);
      continue;
    }
    if (!service) {
      console.error(`finish.mjs: ${name} is declared in ${owner.get(name)} and ${f}, which has no service type`);
      process.exit(1);
    }
    t = t.replace(new RegExp(`\\b${name}\\b`, 'g'), `${service}${name}`);
    renamed++;
  }
  if (t !== readFileSync(p, 'utf8')) writeFileSync(p, t);
}
if (renamed) console.log(`renamed ${renamed} request type(s) duplicated across multi-tag operations`);

// 1c. A query parameter declared `anyOf: [string, array]` (e.g. listListings' `tag`) is typed AnyOfstringarray,
// a type the 5.4.0 Go generator never writes. Declare each missing one: a string where a string is allowed (the
// API takes a comma-separated list for those), interface{} otherwise. A no-op once the generator handles it.
const goFiles = readdirSync(dir).filter((n) => n.endsWith('.go') && n !== 'crossly_anyof.go');
const used = new Set();
const declared = new Set();
for (const f of goFiles) {
  const t = readFileSync(join(dir, f), 'utf8');
  for (const m of t.matchAll(/\b((?:AnyOf|OneOf)[A-Za-z0-9]+)\b/g)) used.add(m[1]);
  for (const m of t.matchAll(/^type ((?:AnyOf|OneOf)[A-Za-z0-9]+)\b/gm)) declared.add(m[1]);
}
const missing = [...used].filter((n) => !declared.has(n) && !/^(?:AnyOf|OneOf)$/.test(n)).sort();
const pkgName = readFileSync(join(dir, 'client.go'), 'utf8').match(/^package (\w+)/m)[1];
if (missing.length) {
  write(
    join(dir, 'crossly_anyof.go'),
    `// Written by codegen/integrations/extras/go/finish.mjs: parameter types the generator references but never\n` +
      `// declares (an anyOf in the spec).\n\npackage ${pkgName}\n\n` +
      missing
        .map((n) =>
          /string/.test(n)
            ? `// ${n} is a string; where the API accepts a list, pass it comma-separated.\ntype ${n} = string\n`
            : `// ${n} is any JSON value.\ntype ${n} = interface{}\n`,
        )
        .join('\n'),
  );
  console.log(`declared ${missing.length} undeclared anyOf/oneOf type(s): ${missing.join(', ')}`);
}

// 3. package root: client ergonomics.
const pkg = readFileSync(join(dir, 'client.go'), 'utf8').match(/^package (\w+)/m)[1];
for (const f of readdirSync(join(HERE, 'client'))) {
  if (f.endsWith('_test.go')) continue;
  const src = readFileSync(join(HERE, 'client', f), 'utf8').replace(/^package crossly$/m, `package ${pkg}`);
  write(join(dir, f), src);
}
write(
  join(dir, 'crossly_version.go'),
  `// Written by codegen/integrations/extras/go/finish.mjs from codegen/state.json.\n\npackage ${pkg}\n\n` +
    `// SDKVersion is this module's release (its tag is ${api}/v<SDKVersion>).\nconst SDKVersion = "${version}"\n`,
);
const cfgFile = join(dir, 'configuration.go');
const cfg = readFileSync(cfgFile, 'utf8');
const cfgFixed = cfg.replace(/UserAgent:(\s+)"OpenAPI-Generator\/[^"]*"/, 'UserAgent:$1UserAgent');
if (cfgFixed === cfg && !/UserAgent:\s+UserAgent,/.test(cfg)) {
  console.error('finish.mjs: the default UserAgent in configuration.go was not found (template changed?)');
  process.exit(1);
}
if (cfgFixed !== cfg) writeFileSync(cfgFile, cfgFixed);

// 3. <dir>/webhooks: the conformance-tested verifier, byte-identical, plus the net/http middleware.
const verifier = join(CODEGEN, 'integrations', 'webhooks', 'go', 'webhooks.go');
mkdirSync(join(dir, 'webhooks'), { recursive: true });
copyFileSync(verifier, join(dir, 'webhooks', 'webhooks.go'));
if (!readFileSync(verifier).equals(readFileSync(join(dir, 'webhooks', 'webhooks.go')))) process.exit(1);
copyFileSync(join(HERE, 'webhooks', 'middleware.go'), join(dir, 'webhooks', 'middleware.go'));

go(dir, 'mod', 'tidy');

// 3. seller: framework adapters, each its own module (tag seller/webhooks/<name>/vX.Y.Z).
if (api === 'seller') {
  for (const name of ['ginx', 'echox']) {
    const to = join(dir, 'webhooks', name);
    write(join(to, `${name}.go`), readFileSync(join(HERE, name, `${name}.go`), 'utf8'));
    const mod = readFileSync(join(HERE, name, 'go.mod.tmpl'), 'utf8').replaceAll('__VERSION__', version);
    // Keep go.mod as tidy left it when only the go.sum-relevant lines are unchanged; rewrite on a version change.
    const cur = existsSync(join(to, 'go.mod')) ? readFileSync(join(to, 'go.mod'), 'utf8') : '';
    if (!cur.includes(`github.com/crossly-dev/crossly-go/seller v${version}`)) writeFileSync(join(to, 'go.mod'), mod);
    go(to, 'mod', 'tidy');
  }
}

console.log(`finished ${api}: module ${MODULE}, version ${version}`);
