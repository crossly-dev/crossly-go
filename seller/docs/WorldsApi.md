# \WorldsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorldExperience**](WorldsApi.md#CreateWorldExperience) | **Post** /v1/worlds/experiences | Create an experience with a start place (no version yet). Counts against the plan&#39;s experience limit.
[**CreateWorldExperienceAccess**](WorldsApi.md#CreateWorldExperienceAccess) | **Post** /v1/worlds/experiences/{experienceId}/access | Who can play: private, friends, unlisted or public (public/unlisted need a reviewed publish).
[**CreateWorldExperienceAlertAck**](WorldsApi.md#CreateWorldExperienceAlertAck) | **Post** /v1/worlds/experiences/{experienceId}/alerts/{alertId}/ack | Acknowledge an alert.
[**CreateWorldExperienceAnalyticQuery**](WorldsApi.md#CreateWorldExperienceAnalyticQuery) | **Post** /v1/worlds/experiences/{experienceId}/analytics/query | Query one custom event over ≤ 31 days: count, sum, avg or players, by day or hour, optionally grouped by a field. Groups under 50 players are suppressed.
[**CreateWorldExperienceArchive**](WorldsApi.md#CreateWorldExperienceArchive) | **Post** /v1/worlds/experiences/{experienceId}/archive | Archive: closes it, drains its servers, frees its plan slot; data is kept.
[**CreateWorldExperienceBan**](WorldsApi.md#CreateWorldExperienceBan) | **Post** /v1/worlds/experiences/{experienceId}/bans | Ban a player from your experience: { sub, duration (seconds, null &#x3D; permanent), reason, displayReason?, applyToAlts? }. Enforced at join and on every running server.
[**CreateWorldExperienceClose**](WorldsApi.md#CreateWorldExperienceClose) | **Post** /v1/worlds/experiences/{experienceId}/close | Make the experience private again (owner-only joins).
[**CreateWorldExperienceConfigRollback**](WorldsApi.md#CreateWorldExperienceConfigRollback) | **Post** /v1/worlds/experiences/{experienceId}/configs/rollback | Instant rollback: an earlier version becomes the newest one.
[**CreateWorldExperienceEvent**](WorldsApi.md#CreateWorldExperienceEvent) | **Post** /v1/worlds/experiences/{experienceId}/events | Schedule a live event: start and end, config values while live, an optional Messaging topic the platform publishes to at start and end.
[**CreateWorldExperienceEventCancel**](WorldsApi.md#CreateWorldExperienceEventCancel) | **Post** /v1/worlds/experiences/{experienceId}/events/{eventId}/cancel | Cancel a scheduled event, or end a live one now.
[**CreateWorldExperienceExperiment**](WorldsApi.md#CreateWorldExperienceExperiment) | **Post** /v1/worlds/experiences/{experienceId}/experiments | Create a draft experiment: 2–5 variants of config values with traffic weights (+ holdout &#x3D; 100), targeting, a primary metric and guardrails, minimum 7 days.
[**CreateWorldExperienceExperimentShip**](WorldsApi.md#CreateWorldExperienceExperimentShip) | **Post** /v1/worlds/experiences/{experienceId}/experiments/{experimentId}/ship | Ship a variant: its values become the base config for everyone and the experiment ends.
[**CreateWorldExperienceExperimentStart**](WorldsApi.md#CreateWorldExperienceExperimentStart) | **Post** /v1/worlds/experiences/{experienceId}/experiments/{experimentId}/start | Start: running servers begin assigning players by their per-experience id.
[**CreateWorldExperienceExperimentStop**](WorldsApi.md#CreateWorldExperienceExperimentStop) | **Post** /v1/worlds/experiences/{experienceId}/experiments/{experimentId}/stop | Stop: every player sees the base config again.
[**CreateWorldExperienceJoin**](WorldsApi.md#CreateWorldExperienceJoin) | **Post** /v1/worlds/experiences/{experienceId}/join | Join as the token&#39;s owner: a server and a 60-second single-use ticket for wss …/v1/realtime/worlds/{serverId}. 202 while a new server boots.
[**CreateWorldExperienceModerationAppeal**](WorldsApi.md#CreateWorldExperienceModerationAppeal) | **Post** /v1/worlds/experiences/{experienceId}/moderation/appeals | Appeal a moderation decision on your game within 6 months (183 days): { kind: action | game | version | asset, targetId (from GET …/moderation/appeals targets), reason (20-4000 characters), evidence?: up to 5 https links }. The older shape { actionId, statement } still appeals an action. One open appeal per decision; a different reviewer answers and you are notified.
[**CreateWorldExperienceOpen**](WorldsApi.md#CreateWorldExperienceOpen) | **Post** /v1/worlds/experiences/{experienceId}/open | Open the experience to players (needs a published start place).
[**CreateWorldExperiencePlace**](WorldsApi.md#CreateWorldExperiencePlace) | **Post** /v1/worlds/experiences/{experienceId}/places | Add a place (at most 10 per experience in P1).
[**CreateWorldExperiencePlaceCheck**](WorldsApi.md#CreateWorldExperiencePlaceCheck) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/checks | Run the publish checks (script lint, budgets, shop-object listings) on a project. Writes nothing.
[**CreateWorldExperiencePlacePublish**](WorldsApi.md#CreateWorldExperiencePlacePublish) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/publishes | Step 1: run checks and freeze a version (the saved draft, or an uploaded versionId). Blocking problems answer 400 with the list.
[**CreateWorldExperiencePlaceStart**](WorldsApi.md#CreateWorldExperiencePlaceStart) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/start | Make this the start place (where a plain join lands).
[**CreateWorldExperiencePlaceTest**](WorldsApi.md#CreateWorldExperiencePlaceTest) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/test | Cloud playtest: a private test version from the saved draft, joined on a fleet server (owner only, never live). 202 while it boots.
[**CreateWorldExperiencePlaceVersion**](WorldsApi.md#CreateWorldExperiencePlaceVersion) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/versions | Upload a version bundle (crossly-worlds-project/1). Not live until published.
[**CreateWorldExperiencePlaceVersionPublish**](WorldsApi.md#CreateWorldExperiencePlaceVersionPublish) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/versions/{versionId}/publish | Make a version current (publish or roll back). New joins land on it; older servers drain as they empty.
[**CreateWorldExperiencePlaceVersionRollback**](WorldsApi.md#CreateWorldExperiencePlaceVersionRollback) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/versions/{versionId}/rollback | Make an earlier reviewed version current again. Unreviewed versions go through …/publishes.
[**CreateWorldExperiencePlaytester**](WorldsApi.md#CreateWorldExperiencePlaytester) | **Post** /v1/worlds/experiences/{experienceId}/playtesters | Add a playtester by Crossly handle (max 200 per game).
[**CreateWorldExperiencePublishCancel**](WorldsApi.md#CreateWorldExperiencePublishCancel) | **Post** /v1/worlds/experiences/{experienceId}/publishes/{requestId}/cancel | Give up a publish that is not live.
[**CreateWorldExperiencePublishHalt**](WorldsApi.md#CreateWorldExperiencePublishHalt) | **Post** /v1/worlds/experiences/{experienceId}/publishes/{requestId}/halt | Stop a rollout: new joins go back to the current version.
[**CreateWorldExperiencePublishRollout**](WorldsApi.md#CreateWorldExperiencePublishRollout) | **Post** /v1/worlds/experiences/{experienceId}/publishes/{requestId}/rollout | Step 7: serve the approved version to percent (1–100) of new joins; 100 makes it current. Crash-heavy rollouts halt automatically.
[**CreateWorldExperiencePublishSubmit**](WorldsApi.md#CreateWorldExperiencePublishSubmit) | **Post** /v1/worlds/experiences/{experienceId}/publishes/{requestId}/submit | Step 6: moderation. First publishes and script changes go to a person (24 h SLA); layout-only and private publishes are approved automatically.
[**CreateWorldExperienceReservedServer**](WorldsApi.md#CreateWorldExperienceReservedServer) | **Post** /v1/worlds/experiences/{experienceId}/reserved-servers | An access code for a reserved server of a place (teleports, matches, tournaments). The code is shown once.
[**CreateWorldExperienceRestore**](WorldsApi.md#CreateWorldExperienceRestore) | **Post** /v1/worlds/experiences/{experienceId}/restore | Restore an archived game, or a deleted one within 30 days; it comes back private.
[**CreateWorldExperienceServerShutdown**](WorldsApi.md#CreateWorldExperienceServerShutdown) | **Post** /v1/worlds/experiences/{experienceId}/servers/shutdown | Shut down servers: mode soft (default) teleports players to fresh servers of the current version first; hard stops now. scope all | outdated (only servers not on a served version).
[**CreateWorldExperienceTransferRequest**](WorldsApi.md#CreateWorldExperienceTransferRequest) | **Post** /v1/worlds/experiences/{experienceId}/transfer-requests | Ask another Crossly account (handle or email) to take over a 3D game you own. One waiting request per game; it expires after 14 days.
[**CreateWorldMatureVerify**](WorldsApi.md#CreateWorldMatureVerify) | **Post** /v1/worlds/mature/verify | Start the ID check Mature needs (Stripe Identity: photo ID + selfie; the date of birth is read off the ID). Returns a hosted URL and client secret, or alreadyVerified.
[**CreateWorldParentalAccept**](WorldsApi.md#CreateWorldParentalAccept) | **Post** /v1/worlds/parental/accept | Accept a child&#39;s invite as the parent: { code, pin }. The parent must be an ID-verified adult (18+ by the date of birth on the ID). The PIN protects every later change.
[**CreateWorldParentalChildrenPin**](WorldsApi.md#CreateWorldParentalChildrenPin) | **Post** /v1/worlds/parental/children/{childId}/pin | Change the parent PIN: { pin, newPin } (4-8 digits). Five wrong PINs lock changes for 15 minutes.
[**CreateWorldParentalChildrenUnlink**](WorldsApi.md#CreateWorldParentalChildrenUnlink) | **Post** /v1/worlds/parental/children/{childId}/unlink | Remove the parent link (and its limits) with your PIN: { pin }.
[**CreateWorldParentalInvite**](WorldsApi.md#CreateWorldParentalInvite) | **Post** /v1/worlds/parental/invite | An under-18 account asks for a parent: a one-time code (72 hours) the parent enters while signed in to their own account. Replaces any pending code.
[**CreateWorldRatingCompute**](WorldsApi.md#CreateWorldRatingCompute) | **Post** /v1/worlds/rating/compute | Rate a set of questionnaire answers exactly as publishing does: { answers } → rating (null when refused), reasons, content descriptors, interactive notices and refusals.
[**CreateWorldReport**](WorldsApi.md#CreateWorldReport) | **Post** /v1/worlds/reports | Report a player (targetSub), the experience, an object (targetInstanceId) or a chat line (chatLineId). Crossly attaches the server, place version and its own chat log; you send ids only.
[**CreateWorldTransferRequestAccept**](WorldsApi.md#CreateWorldTransferRequestAccept) | **Post** /v1/worlds/transfer-requests/{requestId}/accept | Accept a transfer asked of you: the game becomes yours now. Earnings already credited stay with the previous owner; new sales pay you.
[**CreateWorldTransferRequestCancel**](WorldsApi.md#CreateWorldTransferRequestCancel) | **Post** /v1/worlds/transfer-requests/{requestId}/cancel | Cancel a transfer request you sent while it waits.
[**CreateWorldTransferRequestDecline**](WorldsApi.md#CreateWorldTransferRequestDecline) | **Post** /v1/worlds/transfer-requests/{requestId}/decline | Decline a transfer asked of you; nothing changes.
[**DeleteWorldExperience**](WorldsApi.md#DeleteWorldExperience) | **Delete** /v1/worlds/experiences/{experienceId} | Delete a game (a tombstone, restorable for 30 days). Body { confirmName } &#x3D; the game&#39;s name.
[**DeleteWorldExperienceBan**](WorldsApi.md#DeleteWorldExperienceBan) | **Delete** /v1/worlds/experiences/{experienceId}/bans/{sub} | Lift your experience&#39;s ban on a player. Bans placed by Crossly safety are lifted only by appeal.
[**DeleteWorldExperiencePlaytester**](WorldsApi.md#DeleteWorldExperiencePlaytester) | **Delete** /v1/worlds/experiences/{experienceId}/playtesters/{userId} | Remove a playtester.
[**DeleteWorldParentalInvite**](WorldsApi.md#DeleteWorldParentalInvite) | **Delete** /v1/worlds/parental/invite | Withdraw a pending parent invite.
[**GetWorldChart**](WorldsApi.md#GetWorldChart) | **Get** /v1/worlds/charts/{chart} | A chart: top-playing (players now), trending (this week vs last, retention-weighted), top-rated (Wilson lower bound of thumbs up, 10+ votes), new-rising, friends, most-favorited; ?genre&#x3D; cuts it to one genre. Organic only: nothing paid.
[**GetWorldDiscoverExperience**](WorldsApi.md#GetWorldDiscoverExperience) | **Get** /v1/worlds/discover/experiences/{key} | An experience&#39;s thumbs up/down, favourites, your vote and favourite, and whether you have played enough to vote (10 minutes). key: worlds:&lt;id&gt;, classic:&lt;id&gt;, tcg:&lt;game&gt;, cards:&lt;game&gt;.
[**GetWorldDiscoverRow**](WorldsApi.md#GetWorldDiscoverRow) | **Get** /v1/worlds/discover/rows | The Play hub&#39;s rows for you: Continue playing, Friends are playing, Recommended for you, Featured (editorial, first-party capped), Top Playing, Trending, New &amp; Rising, Top Rated, Made in Carnival Studio, Shop the world, card games, Classic games and genre rows. At most one tile is paid, always labelled \&quot;Sponsored\&quot;, never in the first two rows or a row&#39;s first two places.
[**GetWorldExperience**](WorldsApi.md#GetWorldExperience) | **Get** /v1/worlds/experiences/{experienceId} | One of your experiences with its places, live CCU and server count.
[**GetWorldExperienceAccessLimit**](WorldsApi.md#GetWorldExperienceAccessLimit) | **Get** /v1/worlds/experiences/{experienceId}/access-limits | The game&#39;s crew-only and country limits, and the owner&#39;s crew.
[**GetWorldExperienceAlertSetting**](WorldsApi.md#GetWorldExperienceAlertSetting) | **Get** /v1/worlds/experiences/{experienceId}/alerts/settings | Alert rules (threshold, on/off, inbox, webhook) and quiet hours.
[**GetWorldExperienceAnalytic**](WorldsApi.md#GetWorldExperienceAnalytic) | **Get** /v1/worlds/experiences/{experienceId}/analytics/{report} | An analytics report: overview, engagement, retention, monetization, funnels, economy, performance, custom or costs (fleet CPU-minutes, egress, estimated cost). ?from&#x3D;&amp;to&#x3D; (YYYY-MM-DD), retention ?segment&#x3D;, funnels ?funnel&#x3D;, economy ?currency&#x3D;.
[**GetWorldExperienceBan**](WorldsApi.md#GetWorldExperienceBan) | **Get** /v1/worlds/experiences/{experienceId}/bans/{sub} | The live ban on one player, or 404.
[**GetWorldExperienceChatSetting**](WorldsApi.md#GetWorldExperienceChatSetting) | **Get** /v1/worlds/experiences/{experienceId}/chat-settings | The game&#39;s text chat setting and filter level.
[**GetWorldExperienceConfig**](WorldsApi.md#GetWorldExperienceConfig) | **Get** /v1/worlds/experiences/{experienceId}/configs | The remote config document ({ values, rules }) and its version.
[**GetWorldExperienceExperiment**](WorldsApi.md#GetWorldExperienceExperiment) | **Get** /v1/worlds/experiences/{experienceId}/experiments/{experimentId} | Results per variant: n, mean, lift, fixed and always-valid 95% intervals, the verdict and a sample-size estimate.
[**GetWorldExperienceMaturity**](WorldsApi.md#GetWorldExperienceMaturity) | **Get** /v1/worlds/experiences/{experienceId}/maturity | The content maturity questionnaire: questions, answers, the published rating and any rating waiting for the next publish.
[**GetWorldExperienceModeration**](WorldsApi.md#GetWorldExperienceModeration) | **Get** /v1/worlds/experiences/{experienceId}/moderation | Moderation status of your experience: holds (unlisted / taken down), the latest review, the automated scan, strikes, trusted-lane status, open reports and live bans.
[**GetWorldExperienceModerationAppeal**](WorldsApi.md#GetWorldExperienceModerationAppeal) | **Get** /v1/worlds/experiences/{experienceId}/moderation/appeals | What you can appeal on your game and every appeal of it: data.targets (Crossly&#39;s actions, a moderation hold, rejected versions, blocked assets; each with appealableUntil and its latest appeal) and data.appeals (newest first: kind, status open | upheld | overturned, reason, evidence, the reviewer&#39;s note).
[**GetWorldExperiencePlaceDraft**](WorldsApi.md#GetWorldExperiencePlaceDraft) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/draft | The Studio draft of a place ({ world, place }, revision). Falls back to the current version, or an empty place.
[**GetWorldExperiencePublish**](WorldsApi.md#GetWorldExperiencePublish) | **Get** /v1/worlds/experiences/{experienceId}/publishes/{requestId} | One publish with its review state and rollout.
[**GetWorldExperienceShop**](WorldsApi.md#GetWorldExperienceShop) | **Get** /v1/worlds/experiences/{experienceId}/shop | An open experience&#39;s way back to commerce: its creator&#39;s shop and how many live listings it has.
[**GetWorldExperienceShowcase**](WorldsApi.md#GetWorldExperienceShowcase) | **Get** /v1/worlds/experiences/{experienceId}/showcase | The listings the game&#39;s page showcases, and the owner&#39;s live Crossly listings that can be picked.
[**GetWorldFriendPlaying**](WorldsApi.md#GetWorldFriendPlaying) | **Get** /v1/worlds/friends/playing | Your friends (mutual follows) and where each is playing right now, with the link that joins them.
[**GetWorldGameByGame**](WorldsApi.md#GetWorldGameByGame) | **Get** /v1/worlds/games/{game} | A game&#39;s page, as the /games launcher shows it: the tile (players now, thumbs-up %, friends inside), description, tags, creator (Crossly for its own), rating with its content descriptors, visits, created and updated, release notes, gallery, what it sells (bought through Crossly checkout) and the Crossly World app link. :game is a Worlds experience id or any experience key.
[**GetWorldGameMine**](WorldsApi.md#GetWorldGameMine) | **Get** /v1/worlds/games/mine/{gameId} | A game&#39;s summary: status, timeline, the sub-pages you may open, health and next steps.
[**GetWorldGameMineAssetsUsed**](WorldsApi.md#GetWorldGameMineAssetsUsed) | **Get** /v1/worlds/games/mine/{gameId}/assets-used | Everything a game&#39;s published build ships (Assets used): models, sounds, images, scripts, Toolbox assets and other files, each with the places that ship it. A 3D world reads each place&#39;s current version; a Classic game its latest reviewed package version.
[**GetWorldGameServer**](WorldsApi.md#GetWorldGameServer) | **Get** /v1/worlds/games/{game}/servers | A Crossly Carnival game&#39;s public servers that are up: players, max, whether a stranger still fits, your friends on each (friends&#39; servers first, then the fullest with room); and the private servers you pay for. Join one with POST /v1/worlds/experiences/:id/join { serverId }.
[**GetWorldLinkCardGame**](WorldsApi.md#GetWorldLinkCardGame) | **Get** /v1/worlds/links/card-game/{game} | Where to play a card game: Crossly&#39;s own game, its Commons tables, and creator worlds that tag it.
[**GetWorldLinkListing**](WorldsApi.md#GetWorldLinkListing) | **Get** /v1/worlds/links/listing/{slug} | \&quot;Try it in Crossly Carnival\&quot;: open games whose live version shows this Crossly listing (its shop slug), then other games by the same seller. Links only: a sale is attributed to a game only through in-world touches its server reports.
[**GetWorldLinkSeller**](WorldsApi.md#GetWorldLinkSeller) | **Get** /v1/worlds/links/seller/{username} | \&quot;Games by this seller\&quot;: the seller&#39;s open games, best first.
[**GetWorldMature**](WorldsApi.md#GetWorldMature) | **Get** /v1/worlds/mature | Mature (17+) for this account: eligible or not, the steps still to take in order (birth month, ID check, opt-in; region, parental controls and safety holds are final), and the ID age check status.
[**GetWorldMePref**](WorldsApi.md#GetWorldMePref) | **Get** /v1/worlds/me/prefs | Your Crossly Carnival choices (whether games may see your Crossly wallet address).
[**GetWorldParentalChildren**](WorldsApi.md#GetWorldParentalChildren) | **Get** /v1/worlds/parental/children/{childId} | One linked account with ?days&#x3D; (1-90) of playtime per experience.
[**GetWorldParentalMe**](WorldsApi.md#GetWorldParentalMe) | **Get** /v1/worlds/parental/me | This account&#39;s parent link: none, pending (an invite is out) or active (the parent&#39;s name and the limits they set).
[**GetWorldPlaza**](WorldsApi.md#GetWorldPlaza) | **Get** /v1/worlds/plaza | The Carnival Plaza district: the twelve-portal ring (a sponsored portal can only take the last slot, labelled), the new &amp; rising row, the featured stage&#39;s editorial pick and four chart boards, with players now and friends inside.
[**GetWorldPolicyMe**](WorldsApi.md#GetWorldPolicyMe) | **Get** /v1/worlds/policy/me | What this account may do in experiences, by age and region: chat level, purchases, paid random items, trading, links, ads, contests, rating ceiling — with the reason for each \&quot;no\&quot;.
[**GetWorldRatingQuestionnaire**](WorldsApi.md#GetWorldRatingQuestionnaire) | **Get** /v1/worlds/rating/questionnaire | The content-rating questionnaire (IARC-aligned): every question and choice, the three ratings (Everyone, Teen 13+, Mature 17+), what is never allowed, and the regions where Mature is restricted.
[**GetWorldSearch**](WorldsApi.md#GetWorldSearch) | **Get** /v1/worlds/search | Search experiences by title, creator, tags and genre (typo-tolerant, synonyms such as obby / obstacle). Filters: genre, rating (everyone keeps to Everyone), device, players (one server holds at least N), minThumbs (percent), live, friends. No results suggests a chart.
[**GetWorldTransferRequest**](WorldsApi.md#GetWorldTransferRequest) | **Get** /v1/worlds/transfer-requests/{requestId} | One ownership transfer request (only its two people can read it), with the actions you may take.
[**GetWorldTransparency**](WorldsApi.md#GetWorldTransparency) | **Get** /v1/worlds/transparency/{month} | One month&#39;s transparency report (YYYY-MM): reports by reason and source, decisions and median time to decide, actions by kind and issuer, appeals and overturn rate, publish reviews, second reviews of creator bans.
[**ListWorldAppeals**](WorldsApi.md#ListWorldAppeals) | **Get** /v1/worlds/appeals | Your appeals and their decisions.
[**ListWorldExperienceAlerts**](WorldsApi.md#ListWorldExperienceAlerts) | **Get** /v1/worlds/experiences/{experienceId}/alerts | Alerts that fired, newest first.
[**ListWorldExperienceAudit**](WorldsApi.md#ListWorldExperienceAudit) | **Get** /v1/worlds/experiences/{experienceId}/audit | The live-ops audit log: every config, experiment, event and alert change, who made it, when.
[**ListWorldExperienceBans**](WorldsApi.md#ListWorldExperienceBans) | **Get** /v1/worlds/experiences/{experienceId}/bans | Players banned from your experience (by pairwise sub), oldest first (cursor).
[**ListWorldExperienceConfigVersions**](WorldsApi.md#ListWorldExperienceConfigVersions) | **Get** /v1/worlds/experiences/{experienceId}/configs/versions | Config history, newest first: who changed which keys, when.
[**ListWorldExperienceEvents**](WorldsApi.md#ListWorldExperienceEvents) | **Get** /v1/worlds/experiences/{experienceId}/events | Scheduled and live events (?includeEnded&#x3D;true for history).
[**ListWorldExperienceExperiments**](WorldsApi.md#ListWorldExperienceExperiments) | **Get** /v1/worlds/experiences/{experienceId}/experiments | A/B experiments, newest first, with their live verdict.
[**ListWorldExperienceModerationReports**](WorldsApi.md#ListWorldExperienceModerationReports) | **Get** /v1/worlds/experiences/{experienceId}/moderation/reports | Reports about your experience, newest first (cursor). Reporter identity, account ids and other players&#39; words are removed.
[**ListWorldExperiencePlacePublishes**](WorldsApi.md#ListWorldExperiencePlacePublishes) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/publishes | Publishes of a place, newest first: state, checks, rating, access, review, rollout.
[**ListWorldExperiencePlaceVersions**](WorldsApi.md#ListWorldExperiencePlaceVersions) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/versions | Version history of a place, newest first; &#x60;current&#x60; marks the live one.
[**ListWorldExperiencePlaces**](WorldsApi.md#ListWorldExperiencePlaces) | **Get** /v1/worlds/experiences/{experienceId}/places | The places of an experience, start place first.
[**ListWorldExperiencePlaytesters**](WorldsApi.md#ListWorldExperiencePlaytesters) | **Get** /v1/worlds/experiences/{experienceId}/playtesters | People who may join while the game is private.
[**ListWorldExperienceServers**](WorldsApi.md#ListWorldExperienceServers) | **Get** /v1/worlds/experiences/{experienceId}/servers | Live servers: place, version, region, players, health.
[**ListWorldExperiences**](WorldsApi.md#ListWorldExperiences) | **Get** /v1/worlds/experiences | Your experiences, newest first; cursor-paged (?cursor&#x3D;&amp;limit&#x3D;).
[**ListWorldFavorites**](WorldsApi.md#ListWorldFavorites) | **Get** /v1/worlds/favorites | Your favourite experiences that are still open, newest first.
[**ListWorldGameMine**](WorldsApi.md#ListWorldGameMine) | **Get** /v1/worlds/games/mine | Every game you own or collaborate on, 3D worlds and Classic, filtered (?q&amp;status&amp;engine&amp;genre&amp;template&amp;owner&amp;flag&amp;sort; sort&#x3D;revenue &#x3D; highest earning first). Each row carries revenueCents: your creator-ledger earnings from it in USD cents (null for Classic).
[**ListWorldParentalChildren**](WorldsApi.md#ListWorldParentalChildren) | **Get** /v1/worlds/parental/children | Accounts you are the linked parent of: limits (max rating, chat, monthly spend), spend this month, and the last 7 days of playtime per experience.
[**ListWorldRecent**](WorldsApi.md#ListWorldRecent) | **Get** /v1/worlds/recent | Continue playing: what you played in the last 90 days that is still open, newest first (from what the servers recorded).
[**ListWorldRegions**](WorldsApi.md#ListWorldRegions) | **Get** /v1/worlds/regions | Live server regions with a probe URL each: GET it a few times, send the round trips as &#x60;pings&#x60; with the join to land in the nearest region.
[**ListWorldShopListings**](WorldsApi.md#ListWorldShopListings) | **Get** /v1/worlds/shop-listings | Your live Crossly listings a game can show and sell (shop objects bind to their id).
[**ListWorldTransferRequests**](WorldsApi.md#ListWorldTransferRequests) | **Get** /v1/worlds/transfer-requests | Ownership transfer requests you were asked (incoming) or sent (outgoing), newest first (?direction&amp;status&amp;experienceId).
[**ListWorldTransparency**](WorldsApi.md#ListWorldTransparency) | **Get** /v1/worlds/transparency | Published monthly transparency reports, newest first: reports received and actions taken, by type (counts only).
[**UpdateWorldExperience**](WorldsApi.md#UpdateWorldExperience) | **Patch** /v1/worlds/experiences/{experienceId} | Edit name, description, genre, rating or icon.
[**UpdateWorldExperienceAccessLimit**](WorldsApi.md#UpdateWorldExperienceAccessLimit) | **Put** /v1/worlds/experiences/{experienceId}/access-limits | Limit who can play: { crewOnly } (only members of the owner&#39;s crew), { countryMode: all|allow|block, countries: [&#39;US&#39;, ...] }. The owner and team are never limited.
[**UpdateWorldExperienceAlertSetting**](WorldsApi.md#UpdateWorldExperienceAlertSetting) | **Put** /v1/worlds/experiences/{experienceId}/alerts/settings | Change alert rules and quiet hours.
[**UpdateWorldExperienceChatSetting**](WorldsApi.md#UpdateWorldExperienceChatSetting) | **Put** /v1/worlds/experiences/{experienceId}/chat-settings | Turn text chat off or use the strict filter for everyone (never looser than the platform).
[**UpdateWorldExperienceConfig**](WorldsApi.md#UpdateWorldExperienceConfig) | **Put** /v1/worlds/experiences/{experienceId}/configs | Save the whole config document. Running servers get it within seconds (ConfigService.Changed). baseVersion guards against overwriting a newer save (409).
[**UpdateWorldExperienceConfigValue**](WorldsApi.md#UpdateWorldExperienceConfigValue) | **Patch** /v1/worlds/experiences/{experienceId}/configs/values | Set and/or remove individual config keys.
[**UpdateWorldExperienceEvent**](WorldsApi.md#UpdateWorldExperienceEvent) | **Patch** /v1/worlds/experiences/{experienceId}/events/{eventId} | Reschedule or change a scheduled event; a live event may only move its end.
[**UpdateWorldExperienceExperiment**](WorldsApi.md#UpdateWorldExperienceExperiment) | **Patch** /v1/worlds/experiences/{experienceId}/experiments/{experimentId} | Edit a draft (a running experiment may only rename or reword).
[**UpdateWorldExperienceMaturity**](WorldsApi.md#UpdateWorldExperienceMaturity) | **Put** /v1/worlds/experiences/{experienceId}/maturity | Answer the questionnaire: { answers }. A stricter rating applies now; a looser or first Mature rating waits for the next reviewed publish (takesEffect).
[**UpdateWorldExperiencePlace**](WorldsApi.md#UpdateWorldExperiencePlace) | **Patch** /v1/worlds/experiences/{experienceId}/places/{placeId} | Place settings: name, max players (1-50), fill target, social slots, empty shutdown.
[**UpdateWorldExperiencePlaceDraft**](WorldsApi.md#UpdateWorldExperiencePlaceDraft) | **Put** /v1/worlds/experiences/{experienceId}/places/{placeId}/draft | Save the draft. baseRevision is the revision you loaded (0 for none); a newer save elsewhere answers 409 draft_conflict.
[**UpdateWorldExperiencePlaceVersionNote**](WorldsApi.md#UpdateWorldExperiencePlaceVersionNote) | **Put** /v1/worlds/experiences/{experienceId}/places/{placeId}/versions/{versionId}/notes | Edit a release version&#39;s changelog (the game page&#39;s What&#39;s new): { notes } up to 2000 characters.
[**UpdateWorldExperiencePublish**](WorldsApi.md#UpdateWorldExperiencePublish) | **Patch** /v1/worlds/experiences/{experienceId}/publishes/{requestId} | Steps 2–5: details, rating questionnaire answers, access (private | friends | unlisted | public), money review acknowledgement.
[**UpdateWorldExperienceShowcase**](WorldsApi.md#UpdateWorldExperienceShowcase) | **Put** /v1/worlds/experiences/{experienceId}/showcase | Set the showcased listings: { listingIds } (Crossly shop slugs of the owner&#39;s live listings, max 12, in order).
[**UpdateWorldFavorite**](WorldsApi.md#UpdateWorldFavorite) | **Put** /v1/worlds/favorites | Favourite or un-favourite an experience (idempotent).
[**UpdateWorldMature**](WorldsApi.md#UpdateWorldMature) | **Put** /v1/worlds/mature | Switch \&quot;Show Mature experiences\&quot; on or off: { showMature }. On is refused while a final step blocks (too young, region, parental controls, safety hold); it never skips the ID check.
[**UpdateWorldMePref**](WorldsApi.md#UpdateWorldMePref) | **Put** /v1/worlds/me/prefs | Set your Crossly Carnival choices.
[**UpdateWorldParentalChildren**](WorldsApi.md#UpdateWorldParentalChildren) | **Patch** /v1/worlds/parental/children/{childId} | Change a linked account&#39;s limits with your PIN: { pin, settings: { maxRating: everyone|teen, chat: on|friends|off, monthlySpendLimitCents (0 &#x3D; no purchases, null &#x3D; no limit) } }. Enforced server-side at join, chat, the purchase prompt and checkout.
[**UpdateWorldVote**](WorldsApi.md#UpdateWorldVote) | **Put** /v1/worlds/votes | Thumbs up, down, or null to withdraw: one vote per player per experience, counted only after 10 minutes of play (403 vote_needs_playtime before). A light ranking signal.



## CreateWorldExperience

> CreateWorldExperienceResponse CreateWorldExperience(ctx).Execute()

Create an experience with a start place (no version yet). Counts against the plan's experience limit.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperience(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperience`: CreateWorldExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperience`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceRequest struct via the builder pattern


### Return type

**CreateWorldExperienceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceAccess

> WorldExperienceResponse CreateWorldExperienceAccess(ctx, experienceId).Execute()

Who can play: private, friends, unlisted or public (public/unlisted need a reviewed publish).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceAccess(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceAccess`: WorldExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceAlertAck

> CreateWorldExperienceAlertAckResponse CreateWorldExperienceAlertAck(ctx, experienceId, alertId).Execute()

Acknowledge an alert.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    alertId := "alertId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceAlertAck(context.Background(), experienceId, alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceAlertAck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceAlertAck`: CreateWorldExperienceAlertAckResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceAlertAck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**alertId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceAlertAckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperienceAlertAckResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceAnalyticQuery

> CreateWorldExperienceAnalyticQueryResponse CreateWorldExperienceAnalyticQuery(ctx, experienceId).Execute()

Query one custom event over ≤ 31 days: count, sum, avg or players, by day or hour, optionally grouped by a field. Groups under 50 players are suppressed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceAnalyticQuery(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceAnalyticQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceAnalyticQuery`: CreateWorldExperienceAnalyticQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceAnalyticQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceAnalyticQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceAnalyticQueryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceArchive

> WorldExperienceResponse CreateWorldExperienceArchive(ctx, experienceId).Execute()

Archive: closes it, drains its servers, frees its plan slot; data is kept.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceArchive(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceArchive`: WorldExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceBan

> GetWorldExperienceBanResponse CreateWorldExperienceBan(ctx, experienceId).Execute()

Ban a player from your experience: { sub, duration (seconds, null = permanent), reason, displayReason?, applyToAlts? }. Enforced at join and on every running server.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceBan(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceBan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceBan`: GetWorldExperienceBanResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceBan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceBanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceBanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceClose

> WorldExperienceResponse CreateWorldExperienceClose(ctx, experienceId).Execute()

Make the experience private again (owner-only joins).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceClose(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceClose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceClose`: WorldExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceConfigRollback

> WorldExperienceConfigResponse CreateWorldExperienceConfigRollback(ctx, experienceId).Execute()

Instant rollback: an earlier version becomes the newest one.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceConfigRollback(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceConfigRollback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceConfigRollback`: WorldExperienceConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceConfigRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceConfigRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceConfigResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceEvent

> WorldExperienceEventResponse CreateWorldExperienceEvent(ctx, experienceId).Execute()

Schedule a live event: start and end, config values while live, an optional Messaging topic the platform publishes to at start and end.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceEvent(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceEvent`: WorldExperienceEventResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceEventCancel

> WorldExperienceEventResponse CreateWorldExperienceEventCancel(ctx, experienceId, eventId).Execute()

Cancel a scheduled event, or end a live one now.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    eventId := "eventId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceEventCancel(context.Background(), experienceId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceEventCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceEventCancel`: WorldExperienceEventResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceEventCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceEventCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceExperiment

> WorldExperienceExperimentResponse CreateWorldExperienceExperiment(ctx, experienceId).Execute()

Create a draft experiment: 2–5 variants of config values with traffic weights (+ holdout = 100), targeting, a primary metric and guardrails, minimum 7 days.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceExperiment(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceExperiment`: WorldExperienceExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceExperiment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceExperimentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceExperimentShip

> WorldExperienceExperimentResponse CreateWorldExperienceExperimentShip(ctx, experienceId, experimentId).Execute()

Ship a variant: its values become the base config for everyone and the experiment ends.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    experimentId := "experimentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceExperimentShip(context.Background(), experienceId, experimentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceExperimentShip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceExperimentShip`: WorldExperienceExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceExperimentShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**experimentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceExperimentShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceExperimentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceExperimentStart

> WorldExperienceExperimentResponse CreateWorldExperienceExperimentStart(ctx, experienceId, experimentId).Execute()

Start: running servers begin assigning players by their per-experience id.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    experimentId := "experimentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceExperimentStart(context.Background(), experienceId, experimentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceExperimentStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceExperimentStart`: WorldExperienceExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceExperimentStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**experimentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceExperimentStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceExperimentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceExperimentStop

> WorldExperienceExperimentResponse CreateWorldExperienceExperimentStop(ctx, experienceId, experimentId).Execute()

Stop: every player sees the base config again.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    experimentId := "experimentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceExperimentStop(context.Background(), experienceId, experimentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceExperimentStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceExperimentStop`: WorldExperienceExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceExperimentStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**experimentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceExperimentStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceExperimentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceJoin

> CreateWorldExperienceJoinResponse CreateWorldExperienceJoin(ctx, experienceId).Execute()

Join as the token's owner: a server and a 60-second single-use ticket for wss …/v1/realtime/worlds/{serverId}. 202 while a new server boots.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceJoin(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceJoin`: CreateWorldExperienceJoinResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceJoinResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceModerationAppeal

> CreateWorldExperienceModerationAppealResponse CreateWorldExperienceModerationAppeal(ctx, experienceId).Execute()

Appeal a moderation decision on your game within 6 months (183 days): { kind: action | game | version | asset, targetId (from GET …/moderation/appeals targets), reason (20-4000 characters), evidence?: up to 5 https links }. The older shape { actionId, statement } still appeals an action. One open appeal per decision; a different reviewer answers and you are notified.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceModerationAppeal(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceModerationAppeal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceModerationAppeal`: CreateWorldExperienceModerationAppealResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceModerationAppeal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceModerationAppealRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceModerationAppealResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceOpen

> WorldExperienceResponse CreateWorldExperienceOpen(ctx, experienceId).Execute()

Open the experience to players (needs a published start place).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceOpen(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceOpen`: WorldExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceOpen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlace

> WorldExperiencePlaceResponse CreateWorldExperiencePlace(ctx, experienceId).Execute()

Add a place (at most 10 per experience in P1).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePlace(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePlace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlace`: WorldExperiencePlaceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePlace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperiencePlaceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceCheck

> CreateWorldExperiencePlaceCheckResponse CreateWorldExperiencePlaceCheck(ctx, experienceId, placeId).Execute()

Run the publish checks (script lint, budgets, shop-object listings) on a project. Writes nothing.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePlaceCheck(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePlaceCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceCheck`: CreateWorldExperiencePlaceCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePlaceCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceCheckResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlacePublish

> WorldExperienceResponse2 CreateWorldExperiencePlacePublish(ctx, experienceId, placeId).Execute()

Step 1: run checks and freeze a version (the saved draft, or an uploaded versionId). Blocking problems answer 400 with the list.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePlacePublish(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePlacePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlacePublish`: WorldExperienceResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePlacePublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlacePublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceStart

> WorldExperiencePlaceResponse CreateWorldExperiencePlaceStart(ctx, experienceId, placeId).Execute()

Make this the start place (where a plain join lands).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePlaceStart(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePlaceStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceStart`: WorldExperiencePlaceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePlaceStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperiencePlaceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceTest

> CreateWorldExperiencePlaceTestResponse CreateWorldExperiencePlaceTest(ctx, experienceId, placeId).Execute()

Cloud playtest: a private test version from the saved draft, joined on a fleet server (owner only, never live). 202 while it boots.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePlaceTest(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePlaceTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceTest`: CreateWorldExperiencePlaceTestResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePlaceTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceTestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceVersion

> WorldExperiencePlaceVersionResponse CreateWorldExperiencePlaceVersion(ctx, experienceId, placeId).Execute()

Upload a version bundle (crossly-worlds-project/1). Not live until published.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePlaceVersion(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePlaceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceVersion`: WorldExperiencePlaceVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePlaceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperiencePlaceVersionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceVersionPublish

> WorldExperiencePlaceVersionResponse CreateWorldExperiencePlaceVersionPublish(ctx, experienceId, placeId, versionId).Execute()

Make a version current (publish or roll back). New joins land on it; older servers drain as they empty.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 
    versionId := "versionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePlaceVersionPublish(context.Background(), experienceId, placeId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePlaceVersionPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceVersionPublish`: WorldExperiencePlaceVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePlaceVersionPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceVersionPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**WorldExperiencePlaceVersionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceVersionRollback

> WorldExperiencePlaceVersionResponse CreateWorldExperiencePlaceVersionRollback(ctx, experienceId, placeId, versionId).Execute()

Make an earlier reviewed version current again. Unreviewed versions go through …/publishes.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 
    versionId := "versionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePlaceVersionRollback(context.Background(), experienceId, placeId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePlaceVersionRollback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceVersionRollback`: WorldExperiencePlaceVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePlaceVersionRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceVersionRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**WorldExperiencePlaceVersionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaytester

> CreateWorldExperiencePlaytesterResponse CreateWorldExperiencePlaytester(ctx, experienceId).Execute()

Add a playtester by Crossly handle (max 200 per game).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePlaytester(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePlaytester``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaytester`: CreateWorldExperiencePlaytesterResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePlaytester`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaytesterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperiencePlaytesterResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePublishCancel

> WorldExperienceResponse2 CreateWorldExperiencePublishCancel(ctx, experienceId, requestId).Execute()

Give up a publish that is not live.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePublishCancel(context.Background(), experienceId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePublishCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePublishCancel`: WorldExperienceResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePublishCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePublishCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePublishHalt

> WorldExperienceResponse2 CreateWorldExperiencePublishHalt(ctx, experienceId, requestId).Execute()

Stop a rollout: new joins go back to the current version.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePublishHalt(context.Background(), experienceId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePublishHalt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePublishHalt`: WorldExperienceResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePublishHalt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePublishHaltRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePublishRollout

> WorldExperienceResponse2 CreateWorldExperiencePublishRollout(ctx, experienceId, requestId).Execute()

Step 7: serve the approved version to percent (1–100) of new joins; 100 makes it current. Crash-heavy rollouts halt automatically.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePublishRollout(context.Background(), experienceId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePublishRollout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePublishRollout`: WorldExperienceResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePublishRollout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePublishRolloutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePublishSubmit

> WorldExperienceResponse2 CreateWorldExperiencePublishSubmit(ctx, experienceId, requestId).Execute()

Step 6: moderation. First publishes and script changes go to a person (24 h SLA); layout-only and private publishes are approved automatically.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperiencePublishSubmit(context.Background(), experienceId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperiencePublishSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePublishSubmit`: WorldExperienceResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperiencePublishSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePublishSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceReservedServer

> CreateWorldExperienceReservedServerResponse CreateWorldExperienceReservedServer(ctx, experienceId).Execute()

An access code for a reserved server of a place (teleports, matches, tournaments). The code is shown once.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceReservedServer(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceReservedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceReservedServer`: CreateWorldExperienceReservedServerResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceReservedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceReservedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceReservedServerResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceRestore

> WorldExperienceResponse CreateWorldExperienceRestore(ctx, experienceId).Execute()

Restore an archived game, or a deleted one within 30 days; it comes back private.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceRestore(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceRestore`: WorldExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceServerShutdown

> CreateWorldExperienceServerShutdownResponse CreateWorldExperienceServerShutdown(ctx, experienceId).Execute()

Shut down servers: mode soft (default) teleports players to fresh servers of the current version first; hard stops now. scope all | outdated (only servers not on a served version).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceServerShutdown(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceServerShutdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceServerShutdown`: CreateWorldExperienceServerShutdownResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceServerShutdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceServerShutdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceServerShutdownResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceTransferRequest

> WorldResponse2 CreateWorldExperienceTransferRequest(ctx, experienceId).Execute()

Ask another Crossly account (handle or email) to take over a 3D game you own. One waiting request per game; it expires after 14 days.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldExperienceTransferRequest(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldExperienceTransferRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceTransferRequest`: WorldResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldExperienceTransferRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceTransferRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldMatureVerify

> CreateWorldMatureVerifyResponse CreateWorldMatureVerify(ctx).Execute()

Start the ID check Mature needs (Stripe Identity: photo ID + selfie; the date of birth is read off the ID). Returns a hosted URL and client secret, or alreadyVerified.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldMatureVerify(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldMatureVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldMatureVerify`: CreateWorldMatureVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldMatureVerify`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldMatureVerifyRequest struct via the builder pattern


### Return type

**CreateWorldMatureVerifyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldParentalAccept

> WorldParentalResponse CreateWorldParentalAccept(ctx).Execute()

Accept a child's invite as the parent: { code, pin }. The parent must be an ID-verified adult (18+ by the date of birth on the ID). The PIN protects every later change.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldParentalAccept(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldParentalAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldParentalAccept`: WorldParentalResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldParentalAccept`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldParentalAcceptRequest struct via the builder pattern


### Return type

**WorldParentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldParentalChildrenPin

> CreateWorldParentalChildrenPinResponse CreateWorldParentalChildrenPin(ctx, childId).Execute()

Change the parent PIN: { pin, newPin } (4-8 digits). Five wrong PINs lock changes for 15 minutes.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    childId := "childId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldParentalChildrenPin(context.Background(), childId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldParentalChildrenPin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldParentalChildrenPin`: CreateWorldParentalChildrenPinResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldParentalChildrenPin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldParentalChildrenPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldParentalChildrenPinResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldParentalChildrenUnlink

> CreateWorldParentalChildrenUnlinkResponse CreateWorldParentalChildrenUnlink(ctx, childId).Execute()

Remove the parent link (and its limits) with your PIN: { pin }.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    childId := "childId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldParentalChildrenUnlink(context.Background(), childId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldParentalChildrenUnlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldParentalChildrenUnlink`: CreateWorldParentalChildrenUnlinkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldParentalChildrenUnlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldParentalChildrenUnlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldParentalChildrenUnlinkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldParentalInvite

> CreateWorldParentalInviteResponse CreateWorldParentalInvite(ctx).Execute()

An under-18 account asks for a parent: a one-time code (72 hours) the parent enters while signed in to their own account. Replaces any pending code.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldParentalInvite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldParentalInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldParentalInvite`: CreateWorldParentalInviteResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldParentalInvite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldParentalInviteRequest struct via the builder pattern


### Return type

**CreateWorldParentalInviteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldRatingCompute

> CreateWorldRatingComputeResponse CreateWorldRatingCompute(ctx).Execute()

Rate a set of questionnaire answers exactly as publishing does: { answers } → rating (null when refused), reasons, content descriptors, interactive notices and refusals.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldRatingCompute(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldRatingCompute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldRatingCompute`: CreateWorldRatingComputeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldRatingCompute`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldRatingComputeRequest struct via the builder pattern


### Return type

**CreateWorldRatingComputeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldReport

> CreateWorldReportResponse CreateWorldReport(ctx).Execute()

Report a player (targetSub), the experience, an object (targetInstanceId) or a chat line (chatLineId). Crossly attaches the server, place version and its own chat log; you send ids only.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldReport`: CreateWorldReportResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldReportRequest struct via the builder pattern


### Return type

**CreateWorldReportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldTransferRequestAccept

> WorldResponse2 CreateWorldTransferRequestAccept(ctx, requestId).Execute()

Accept a transfer asked of you: the game becomes yours now. Earnings already credited stay with the previous owner; new sales pay you.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldTransferRequestAccept(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldTransferRequestAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldTransferRequestAccept`: WorldResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldTransferRequestAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldTransferRequestAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldTransferRequestCancel

> WorldResponse2 CreateWorldTransferRequestCancel(ctx, requestId).Execute()

Cancel a transfer request you sent while it waits.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldTransferRequestCancel(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldTransferRequestCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldTransferRequestCancel`: WorldResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldTransferRequestCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldTransferRequestCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldTransferRequestDecline

> WorldResponse2 CreateWorldTransferRequestDecline(ctx, requestId).Execute()

Decline a transfer asked of you; nothing changes.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.CreateWorldTransferRequestDecline(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.CreateWorldTransferRequestDecline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldTransferRequestDecline`: WorldResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.CreateWorldTransferRequestDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldTransferRequestDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperience

> DeleteWorldExperienceResponse DeleteWorldExperience(ctx, experienceId).Execute()

Delete a game (a tombstone, restorable for 30 days). Body { confirmName } = the game's name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.DeleteWorldExperience(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.DeleteWorldExperience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperience`: DeleteWorldExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.DeleteWorldExperience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteWorldExperienceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperienceBan

> DeleteWorldExperienceBanResponse DeleteWorldExperienceBan(ctx, experienceId, sub).Execute()

Lift your experience's ban on a player. Bans placed by Crossly safety are lifted only by appeal.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    sub := "sub_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.DeleteWorldExperienceBan(context.Background(), experienceId, sub).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.DeleteWorldExperienceBan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperienceBan`: DeleteWorldExperienceBanResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.DeleteWorldExperienceBan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**sub** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperienceBanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**DeleteWorldExperienceBanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperiencePlaytester

> DeleteSpatialSceneGrantResponse DeleteWorldExperiencePlaytester(ctx, experienceId, userId).Execute()

Remove a playtester.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.DeleteWorldExperiencePlaytester(context.Background(), experienceId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.DeleteWorldExperiencePlaytester``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperiencePlaytester`: DeleteSpatialSceneGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.DeleteWorldExperiencePlaytester`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperiencePlaytesterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**DeleteSpatialSceneGrantResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldParentalInvite

> DeleteDaoProposalResponse DeleteWorldParentalInvite(ctx).Execute()

Withdraw a pending parent invite.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.DeleteWorldParentalInvite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.DeleteWorldParentalInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldParentalInvite`: DeleteDaoProposalResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.DeleteWorldParentalInvite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldParentalInviteRequest struct via the builder pattern


### Return type

**DeleteDaoProposalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldChart

> GetWorldChartResponse GetWorldChart(ctx, chart).Execute()

A chart: top-playing (players now), trending (this week vs last, retention-weighted), top-rated (Wilson lower bound of thumbs up, 10+ votes), new-rising, friends, most-favorited; ?genre= cuts it to one genre. Organic only: nothing paid.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chart := "chart_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldChart(context.Background(), chart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldChart`: GetWorldChartResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldChart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chart** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldChartResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldDiscoverExperience

> WorldResponse3 GetWorldDiscoverExperience(ctx, key).Execute()

An experience's thumbs up/down, favourites, your vote and favourite, and whether you have played enough to vote (10 minutes). key: worlds:<id>, classic:<id>, tcg:<game>, cards:<game>.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldDiscoverExperience(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldDiscoverExperience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldDiscoverExperience`: WorldResponse3
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldDiscoverExperience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldDiscoverExperienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldResponse3**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldDiscoverRow

> GetWorldDiscoverRowResponse GetWorldDiscoverRow(ctx).Execute()

The Play hub's rows for you: Continue playing, Friends are playing, Recommended for you, Featured (editorial, first-party capped), Top Playing, Trending, New & Rising, Top Rated, Made in Carnival Studio, Shop the world, card games, Classic games and genre rows. At most one tile is paid, always labelled \"Sponsored\", never in the first two rows or a row's first two places.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldDiscoverRow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldDiscoverRow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldDiscoverRow`: GetWorldDiscoverRowResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldDiscoverRow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldDiscoverRowRequest struct via the builder pattern


### Return type

**GetWorldDiscoverRowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperience

> GetWorldExperienceResponse GetWorldExperience(ctx, experienceId).Execute()

One of your experiences with its places, live CCU and server count.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperience(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperience`: GetWorldExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceAccessLimit

> GetWorldExperienceAccessLimitResponse GetWorldExperienceAccessLimit(ctx, experienceId).Execute()

The game's crew-only and country limits, and the owner's crew.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceAccessLimit(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceAccessLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceAccessLimit`: GetWorldExperienceAccessLimitResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceAccessLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceAccessLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceAccessLimitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceAlertSetting

> GetWorldExperienceAlertSettingResponse GetWorldExperienceAlertSetting(ctx, experienceId).Execute()

Alert rules (threshold, on/off, inbox, webhook) and quiet hours.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceAlertSetting(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceAlertSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceAlertSetting`: GetWorldExperienceAlertSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceAlertSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceAlertSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceAlertSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceAnalytic

> GetWorldExperienceAnalyticResponse GetWorldExperienceAnalytic(ctx, experienceId, report).Execute()

An analytics report: overview, engagement, retention, monetization, funnels, economy, performance, custom or costs (fleet CPU-minutes, egress, estimated cost). ?from=&to= (YYYY-MM-DD), retention ?segment=, funnels ?funnel=, economy ?currency=.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    report := "report_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceAnalytic(context.Background(), experienceId, report).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceAnalytic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceAnalytic`: GetWorldExperienceAnalyticResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceAnalytic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**report** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceAnalyticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperienceAnalyticResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceBan

> GetWorldExperienceBanResponse GetWorldExperienceBan(ctx, experienceId, sub).Execute()

The live ban on one player, or 404.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    sub := "sub_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceBan(context.Background(), experienceId, sub).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceBan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceBan`: GetWorldExperienceBanResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceBan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**sub** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceBanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperienceBanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceChatSetting

> GetWorldExperienceChatSettingResponse GetWorldExperienceChatSetting(ctx, experienceId).Execute()

The game's text chat setting and filter level.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceChatSetting(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceChatSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceChatSetting`: GetWorldExperienceChatSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceChatSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceChatSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceChatSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceConfig

> WorldExperienceConfigResponse GetWorldExperienceConfig(ctx, experienceId).Execute()

The remote config document ({ values, rules }) and its version.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceConfig(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceConfig`: WorldExperienceConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceConfigResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceExperiment

> GetWorldExperienceExperimentResponse GetWorldExperienceExperiment(ctx, experienceId, experimentId).Execute()

Results per variant: n, mean, lift, fixed and always-valid 95% intervals, the verdict and a sample-size estimate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    experimentId := "experimentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceExperiment(context.Background(), experienceId, experimentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceExperiment`: GetWorldExperienceExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceExperiment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**experimentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperienceExperimentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceMaturity

> GetWorldExperienceMaturityResponse GetWorldExperienceMaturity(ctx, experienceId).Execute()

The content maturity questionnaire: questions, answers, the published rating and any rating waiting for the next publish.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceMaturity(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceMaturity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceMaturity`: GetWorldExperienceMaturityResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceMaturity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceMaturityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceMaturityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceModeration

> GetWorldExperienceModerationResponse GetWorldExperienceModeration(ctx, experienceId).Execute()

Moderation status of your experience: holds (unlisted / taken down), the latest review, the automated scan, strikes, trusted-lane status, open reports and live bans.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceModeration(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceModeration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceModeration`: GetWorldExperienceModerationResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceModeration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceModerationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceModerationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceModerationAppeal

> GetWorldExperienceModerationAppealResponse GetWorldExperienceModerationAppeal(ctx, experienceId).Execute()

What you can appeal on your game and every appeal of it: data.targets (Crossly's actions, a moderation hold, rejected versions, blocked assets; each with appealableUntil and its latest appeal) and data.appeals (newest first: kind, status open | upheld | overturned, reason, evidence, the reviewer's note).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceModerationAppeal(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceModerationAppeal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceModerationAppeal`: GetWorldExperienceModerationAppealResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceModerationAppeal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceModerationAppealRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceModerationAppealResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperiencePlaceDraft

> GetWorldExperiencePlaceDraftResponse GetWorldExperiencePlaceDraft(ctx, experienceId, placeId).Execute()

The Studio draft of a place ({ world, place }, revision). Falls back to the current version, or an empty place.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperiencePlaceDraft(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperiencePlaceDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperiencePlaceDraft`: GetWorldExperiencePlaceDraftResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperiencePlaceDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperiencePlaceDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperiencePlaceDraftResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperiencePublish

> WorldExperienceResponse2 GetWorldExperiencePublish(ctx, experienceId, requestId).Execute()

One publish with its review state and rollout.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperiencePublish(context.Background(), experienceId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperiencePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperiencePublish`: WorldExperienceResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperiencePublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperiencePublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceShop

> GetWorldExperienceShopResponse GetWorldExperienceShop(ctx, experienceId).Execute()

An open experience's way back to commerce: its creator's shop and how many live listings it has.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceShop(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceShop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceShop`: GetWorldExperienceShopResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceShop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceShopResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceShowcase

> GetWorldExperienceShowcaseResponse GetWorldExperienceShowcase(ctx, experienceId).Execute()

The listings the game's page showcases, and the owner's live Crossly listings that can be picked.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldExperienceShowcase(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldExperienceShowcase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceShowcase`: GetWorldExperienceShowcaseResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldExperienceShowcase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceShowcaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceShowcaseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldFriendPlaying

> GetWorldFriendPlayingResponse GetWorldFriendPlaying(ctx).Execute()

Your friends (mutual follows) and where each is playing right now, with the link that joins them.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldFriendPlaying(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldFriendPlaying``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldFriendPlaying`: GetWorldFriendPlayingResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldFriendPlaying`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldFriendPlayingRequest struct via the builder pattern


### Return type

**GetWorldFriendPlayingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldGameByGame

> GetWorldGameByGameResponse GetWorldGameByGame(ctx, game).Execute()

A game's page, as the /games launcher shows it: the tile (players now, thumbs-up %, friends inside), description, tags, creator (Crossly for its own), rating with its content descriptors, visits, created and updated, release notes, gallery, what it sells (bought through Crossly checkout) and the Crossly World app link. :game is a Worlds experience id or any experience key.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    game := "game_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldGameByGame(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldGameByGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldGameByGame`: GetWorldGameByGameResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldGameByGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldGameByGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldGameByGameResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldGameMine

> GetWorldGameMineResponse GetWorldGameMine(ctx, gameId).Execute()

A game's summary: status, timeline, the sub-pages you may open, health and next steps.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gameId := "gameId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldGameMine(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldGameMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldGameMine`: GetWorldGameMineResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldGameMine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldGameMineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldGameMineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldGameMineAssetsUsed

> GetWorldGameMineAssetsUsedResponse GetWorldGameMineAssetsUsed(ctx, gameId).Execute()

Everything a game's published build ships (Assets used): models, sounds, images, scripts, Toolbox assets and other files, each with the places that ship it. A 3D world reads each place's current version; a Classic game its latest reviewed package version.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gameId := "gameId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldGameMineAssetsUsed(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldGameMineAssetsUsed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldGameMineAssetsUsed`: GetWorldGameMineAssetsUsedResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldGameMineAssetsUsed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldGameMineAssetsUsedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldGameMineAssetsUsedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldGameServer

> GetWorldGameServerResponse GetWorldGameServer(ctx, game).Execute()

A Crossly Carnival game's public servers that are up: players, max, whether a stranger still fits, your friends on each (friends' servers first, then the fullest with room); and the private servers you pay for. Join one with POST /v1/worlds/experiences/:id/join { serverId }.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    game := "game_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldGameServer(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldGameServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldGameServer`: GetWorldGameServerResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldGameServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldGameServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldGameServerResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldLinkCardGame

> WorldLinkResponse GetWorldLinkCardGame(ctx, game).Execute()

Where to play a card game: Crossly's own game, its Commons tables, and creator worlds that tag it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    game := "game_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldLinkCardGame(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldLinkCardGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldLinkCardGame`: WorldLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldLinkCardGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldLinkCardGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldLinkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldLinkListing

> WorldLinkResponse GetWorldLinkListing(ctx, slug).Execute()

\"Try it in Crossly Carnival\": open games whose live version shows this Crossly listing (its shop slug), then other games by the same seller. Links only: a sale is attributed to a game only through in-world touches its server reports.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    slug := "slug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldLinkListing(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldLinkListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldLinkListing`: WorldLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldLinkListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldLinkListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldLinkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldLinkSeller

> WorldLinkResponse GetWorldLinkSeller(ctx, username).Execute()

\"Games by this seller\": the seller's open games, best first.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldLinkSeller(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldLinkSeller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldLinkSeller`: WorldLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldLinkSeller`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldLinkSellerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldLinkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldMature

> GetWorldMatureResponse GetWorldMature(ctx).Execute()

Mature (17+) for this account: eligible or not, the steps still to take in order (birth month, ID check, opt-in; region, parental controls and safety holds are final), and the ID age check status.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldMature(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldMature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldMature`: GetWorldMatureResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldMature`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldMatureRequest struct via the builder pattern


### Return type

**GetWorldMatureResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldMePref

> GetWorldMePrefResponse GetWorldMePref(ctx).Execute()

Your Crossly Carnival choices (whether games may see your Crossly wallet address).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldMePref(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldMePref``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldMePref`: GetWorldMePrefResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldMePref`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldMePrefRequest struct via the builder pattern


### Return type

**GetWorldMePrefResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldParentalChildren

> WorldParentalResponse GetWorldParentalChildren(ctx, childId).Execute()

One linked account with ?days= (1-90) of playtime per experience.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    childId := "childId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldParentalChildren(context.Background(), childId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldParentalChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldParentalChildren`: WorldParentalResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldParentalChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldParentalChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldParentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldParentalMe

> GetWorldParentalMeResponse GetWorldParentalMe(ctx).Execute()

This account's parent link: none, pending (an invite is out) or active (the parent's name and the limits they set).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldParentalMe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldParentalMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldParentalMe`: GetWorldParentalMeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldParentalMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldParentalMeRequest struct via the builder pattern


### Return type

**GetWorldParentalMeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldPlaza

> GetWorldPlazaResponse GetWorldPlaza(ctx).Execute()

The Carnival Plaza district: the twelve-portal ring (a sponsored portal can only take the last slot, labelled), the new & rising row, the featured stage's editorial pick and four chart boards, with players now and friends inside.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldPlaza(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldPlaza``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldPlaza`: GetWorldPlazaResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldPlaza`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldPlazaRequest struct via the builder pattern


### Return type

**GetWorldPlazaResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldPolicyMe

> GetWorldPolicyMeResponse GetWorldPolicyMe(ctx).Execute()

What this account may do in experiences, by age and region: chat level, purchases, paid random items, trading, links, ads, contests, rating ceiling — with the reason for each \"no\".

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldPolicyMe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldPolicyMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldPolicyMe`: GetWorldPolicyMeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldPolicyMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldPolicyMeRequest struct via the builder pattern


### Return type

**GetWorldPolicyMeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldRatingQuestionnaire

> GetWorldRatingQuestionnaireResponse GetWorldRatingQuestionnaire(ctx).Execute()

The content-rating questionnaire (IARC-aligned): every question and choice, the three ratings (Everyone, Teen 13+, Mature 17+), what is never allowed, and the regions where Mature is restricted.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldRatingQuestionnaire(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldRatingQuestionnaire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldRatingQuestionnaire`: GetWorldRatingQuestionnaireResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldRatingQuestionnaire`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldRatingQuestionnaireRequest struct via the builder pattern


### Return type

**GetWorldRatingQuestionnaireResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldSearch

> GetWorldSearchResponse GetWorldSearch(ctx).Execute()

Search experiences by title, creator, tags and genre (typo-tolerant, synonyms such as obby / obstacle). Filters: genre, rating (everyone keeps to Everyone), device, players (one server holds at least N), minThumbs (percent), live, friends. No results suggests a chart.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldSearch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldSearch`: GetWorldSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldSearch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldSearchRequest struct via the builder pattern


### Return type

**GetWorldSearchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldTransferRequest

> WorldResponse2 GetWorldTransferRequest(ctx, requestId).Execute()

One ownership transfer request (only its two people can read it), with the actions you may take.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldTransferRequest(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldTransferRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldTransferRequest`: WorldResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldTransferRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldTransferRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldTransparency

> GetWorldTransparencyResponse GetWorldTransparency(ctx, month).Execute()

One month's transparency report (YYYY-MM): reports by reason and source, decisions and median time to decide, actions by kind and issuer, appeals and overturn rate, publish reviews, second reviews of creator bans.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    month := "month_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.GetWorldTransparency(context.Background(), month).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.GetWorldTransparency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldTransparency`: GetWorldTransparencyResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.GetWorldTransparency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**month** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldTransparencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldTransparencyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldAppeals

> ListWorldAppealsResponse ListWorldAppeals(ctx).Execute()

Your appeals and their decisions.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldAppeals(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldAppeals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldAppeals`: ListWorldAppealsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldAppeals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldAppealsRequest struct via the builder pattern


### Return type

**ListWorldAppealsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceAlerts

> ListWorldExperienceAlertsResponse ListWorldExperienceAlerts(ctx, experienceId).Execute()

Alerts that fired, newest first.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperienceAlerts(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperienceAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceAlerts`: ListWorldExperienceAlertsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperienceAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceAlertsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceAudit

> ListWorldExperienceAuditResponse ListWorldExperienceAudit(ctx, experienceId).Execute()

The live-ops audit log: every config, experiment, event and alert change, who made it, when.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperienceAudit(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperienceAudit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceAudit`: ListWorldExperienceAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperienceAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceAuditResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceBans

> ListWorldExperienceBansResponse ListWorldExperienceBans(ctx, experienceId).Execute()

Players banned from your experience (by pairwise sub), oldest first (cursor).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperienceBans(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperienceBans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceBans`: ListWorldExperienceBansResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperienceBans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceBansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceBansResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceConfigVersions

> ListWorldExperienceConfigVersionsResponse ListWorldExperienceConfigVersions(ctx, experienceId).Execute()

Config history, newest first: who changed which keys, when.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperienceConfigVersions(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperienceConfigVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceConfigVersions`: ListWorldExperienceConfigVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperienceConfigVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceConfigVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceConfigVersionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceEvents

> ListWorldExperienceEventsResponse ListWorldExperienceEvents(ctx, experienceId).Execute()

Scheduled and live events (?includeEnded=true for history).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperienceEvents(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperienceEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceEvents`: ListWorldExperienceEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperienceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceExperiments

> ListWorldExperienceExperimentsResponse ListWorldExperienceExperiments(ctx, experienceId).Execute()

A/B experiments, newest first, with their live verdict.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperienceExperiments(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperienceExperiments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceExperiments`: ListWorldExperienceExperimentsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperienceExperiments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceExperimentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceExperimentsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceModerationReports

> ListWorldExperienceModerationReportsResponse ListWorldExperienceModerationReports(ctx, experienceId).Execute()

Reports about your experience, newest first (cursor). Reporter identity, account ids and other players' words are removed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperienceModerationReports(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperienceModerationReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceModerationReports`: ListWorldExperienceModerationReportsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperienceModerationReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceModerationReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceModerationReportsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperiencePlacePublishes

> ListWorldExperiencePlacePublishesResponse ListWorldExperiencePlacePublishes(ctx, experienceId, placeId).Execute()

Publishes of a place, newest first: state, checks, rating, access, review, rollout.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperiencePlacePublishes(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperiencePlacePublishes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperiencePlacePublishes`: ListWorldExperiencePlacePublishesResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperiencePlacePublishes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperiencePlacePublishesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**ListWorldExperiencePlacePublishesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperiencePlaceVersions

> ListWorldExperiencePlaceVersionsResponse ListWorldExperiencePlaceVersions(ctx, experienceId, placeId).Execute()

Version history of a place, newest first; `current` marks the live one.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperiencePlaceVersions(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperiencePlaceVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperiencePlaceVersions`: ListWorldExperiencePlaceVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperiencePlaceVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperiencePlaceVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**ListWorldExperiencePlaceVersionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperiencePlaces

> ListWorldExperiencePlacesResponse ListWorldExperiencePlaces(ctx, experienceId).Execute()

The places of an experience, start place first.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperiencePlaces(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperiencePlaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperiencePlaces`: ListWorldExperiencePlacesResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperiencePlaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperiencePlacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperiencePlacesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperiencePlaytesters

> ListWorldExperiencePlaytestersResponse ListWorldExperiencePlaytesters(ctx, experienceId).Execute()

People who may join while the game is private.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperiencePlaytesters(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperiencePlaytesters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperiencePlaytesters`: ListWorldExperiencePlaytestersResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperiencePlaytesters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperiencePlaytestersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperiencePlaytestersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceServers

> ListWorldExperienceServersResponse ListWorldExperienceServers(ctx, experienceId).Execute()

Live servers: place, version, region, players, health.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperienceServers(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperienceServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceServers`: ListWorldExperienceServersResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperienceServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceServersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperiences

> ListWorldExperiencesResponse ListWorldExperiences(ctx).Execute()

Your experiences, newest first; cursor-paged (?cursor=&limit=).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldExperiences(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldExperiences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperiences`: ListWorldExperiencesResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldExperiences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperiencesRequest struct via the builder pattern


### Return type

**ListWorldExperiencesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldFavorites

> ListWorldRecentResponse ListWorldFavorites(ctx).Execute()

Your favourite experiences that are still open, newest first.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldFavorites(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldFavorites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldFavorites`: ListWorldRecentResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldFavorites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldFavoritesRequest struct via the builder pattern


### Return type

**ListWorldRecentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldGameMine

> ListWorldGameMineResponse ListWorldGameMine(ctx).Execute()

Every game you own or collaborate on, 3D worlds and Classic, filtered (?q&status&engine&genre&template&owner&flag&sort; sort=revenue = highest earning first). Each row carries revenueCents: your creator-ledger earnings from it in USD cents (null for Classic).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldGameMine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldGameMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldGameMine`: ListWorldGameMineResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldGameMine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldGameMineRequest struct via the builder pattern


### Return type

**ListWorldGameMineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldParentalChildren

> ListWorldParentalChildrenResponse ListWorldParentalChildren(ctx).Execute()

Accounts you are the linked parent of: limits (max rating, chat, monthly spend), spend this month, and the last 7 days of playtime per experience.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldParentalChildren(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldParentalChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldParentalChildren`: ListWorldParentalChildrenResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldParentalChildren`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldParentalChildrenRequest struct via the builder pattern


### Return type

**ListWorldParentalChildrenResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldRecent

> ListWorldRecentResponse ListWorldRecent(ctx).Execute()

Continue playing: what you played in the last 90 days that is still open, newest first (from what the servers recorded).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldRecent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldRecent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldRecent`: ListWorldRecentResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldRecent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldRecentRequest struct via the builder pattern


### Return type

**ListWorldRecentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldRegions

> ListWorldRegionsResponse ListWorldRegions(ctx).Execute()

Live server regions with a probe URL each: GET it a few times, send the round trips as `pings` with the join to land in the nearest region.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldRegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldRegions`: ListWorldRegionsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldRegionsRequest struct via the builder pattern


### Return type

**ListWorldRegionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldShopListings

> ListWorldShopListingsResponse ListWorldShopListings(ctx).Execute()

Your live Crossly listings a game can show and sell (shop objects bind to their id).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldShopListings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldShopListings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldShopListings`: ListWorldShopListingsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldShopListings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldShopListingsRequest struct via the builder pattern


### Return type

**ListWorldShopListingsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldTransferRequests

> ListWorldTransferRequestsResponse ListWorldTransferRequests(ctx).Execute()

Ownership transfer requests you were asked (incoming) or sent (outgoing), newest first (?direction&status&experienceId).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldTransferRequests(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldTransferRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldTransferRequests`: ListWorldTransferRequestsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldTransferRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldTransferRequestsRequest struct via the builder pattern


### Return type

**ListWorldTransferRequestsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldTransparency

> ListWorldTransparencyResponse ListWorldTransparency(ctx).Execute()

Published monthly transparency reports, newest first: reports received and actions taken, by type (counts only).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.ListWorldTransparency(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.ListWorldTransparency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldTransparency`: ListWorldTransparencyResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.ListWorldTransparency`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldTransparencyRequest struct via the builder pattern


### Return type

**ListWorldTransparencyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperience

> WorldExperienceResponse UpdateWorldExperience(ctx, experienceId).Execute()

Edit name, description, genre, rating or icon.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperience(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperience`: WorldExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceAccessLimit

> GetWorldExperienceAccessLimitResponse UpdateWorldExperienceAccessLimit(ctx, experienceId).Execute()

Limit who can play: { crewOnly } (only members of the owner's crew), { countryMode: all|allow|block, countries: ['US', ...] }. The owner and team are never limited.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperienceAccessLimit(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperienceAccessLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceAccessLimit`: GetWorldExperienceAccessLimitResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperienceAccessLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceAccessLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceAccessLimitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceAlertSetting

> GetWorldExperienceAlertSettingResponse UpdateWorldExperienceAlertSetting(ctx, experienceId).Execute()

Change alert rules and quiet hours.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperienceAlertSetting(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperienceAlertSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceAlertSetting`: GetWorldExperienceAlertSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperienceAlertSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceAlertSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceAlertSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceChatSetting

> GetWorldExperienceChatSettingResponse UpdateWorldExperienceChatSetting(ctx, experienceId).Execute()

Turn text chat off or use the strict filter for everyone (never looser than the platform).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperienceChatSetting(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperienceChatSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceChatSetting`: GetWorldExperienceChatSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperienceChatSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceChatSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceChatSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceConfig

> WorldExperienceConfigResponse UpdateWorldExperienceConfig(ctx, experienceId).Execute()

Save the whole config document. Running servers get it within seconds (ConfigService.Changed). baseVersion guards against overwriting a newer save (409).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperienceConfig(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperienceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceConfig`: WorldExperienceConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperienceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceConfigResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceConfigValue

> WorldExperienceConfigResponse UpdateWorldExperienceConfigValue(ctx, experienceId).Execute()

Set and/or remove individual config keys.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperienceConfigValue(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperienceConfigValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceConfigValue`: WorldExperienceConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperienceConfigValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceConfigValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceConfigResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceEvent

> WorldExperienceEventResponse UpdateWorldExperienceEvent(ctx, experienceId, eventId).Execute()

Reschedule or change a scheduled event; a live event may only move its end.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    eventId := "eventId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperienceEvent(context.Background(), experienceId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperienceEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceEvent`: WorldExperienceEventResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperienceEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceExperiment

> WorldExperienceExperimentResponse UpdateWorldExperienceExperiment(ctx, experienceId, experimentId).Execute()

Edit a draft (a running experiment may only rename or reword).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    experimentId := "experimentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperienceExperiment(context.Background(), experienceId, experimentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperienceExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceExperiment`: WorldExperienceExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperienceExperiment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**experimentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceExperimentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceMaturity

> UpdateWorldExperienceMaturityResponse UpdateWorldExperienceMaturity(ctx, experienceId).Execute()

Answer the questionnaire: { answers }. A stricter rating applies now; a looser or first Mature rating waits for the next reviewed publish (takesEffect).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperienceMaturity(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperienceMaturity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceMaturity`: UpdateWorldExperienceMaturityResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperienceMaturity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceMaturityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateWorldExperienceMaturityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperiencePlace

> WorldExperiencePlaceResponse UpdateWorldExperiencePlace(ctx, experienceId, placeId).Execute()

Place settings: name, max players (1-50), fill target, social slots, empty shutdown.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperiencePlace(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperiencePlace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperiencePlace`: WorldExperiencePlaceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperiencePlace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperiencePlaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperiencePlaceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperiencePlaceDraft

> GetWorldExperiencePlaceDraftResponse UpdateWorldExperiencePlaceDraft(ctx, experienceId, placeId).Execute()

Save the draft. baseRevision is the revision you loaded (0 for none); a newer save elsewhere answers 409 draft_conflict.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperiencePlaceDraft(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperiencePlaceDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperiencePlaceDraft`: GetWorldExperiencePlaceDraftResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperiencePlaceDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperiencePlaceDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperiencePlaceDraftResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperiencePlaceVersionNote

> UpdateWorldExperiencePlaceVersionNoteResponse UpdateWorldExperiencePlaceVersionNote(ctx, experienceId, placeId, versionId).Execute()

Edit a release version's changelog (the game page's What's new): { notes } up to 2000 characters.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 
    versionId := "versionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperiencePlaceVersionNote(context.Background(), experienceId, placeId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperiencePlaceVersionNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperiencePlaceVersionNote`: UpdateWorldExperiencePlaceVersionNoteResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperiencePlaceVersionNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperiencePlaceVersionNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**UpdateWorldExperiencePlaceVersionNoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperiencePublish

> WorldExperienceResponse2 UpdateWorldExperiencePublish(ctx, experienceId, requestId).Execute()

Steps 2–5: details, rating questionnaire answers, access (private | friends | unlisted | public), money review acknowledgement.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperiencePublish(context.Background(), experienceId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperiencePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperiencePublish`: WorldExperienceResponse2
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperiencePublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperiencePublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceShowcase

> UpdateWorldExperienceShowcaseResponse UpdateWorldExperienceShowcase(ctx, experienceId).Execute()

Set the showcased listings: { listingIds } (Crossly shop slugs of the owner's live listings, max 12, in order).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldExperienceShowcase(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldExperienceShowcase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceShowcase`: UpdateWorldExperienceShowcaseResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldExperienceShowcase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceShowcaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateWorldExperienceShowcaseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldFavorite

> WorldResponse3 UpdateWorldFavorite(ctx).Execute()

Favourite or un-favourite an experience (idempotent).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldFavorite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldFavorite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldFavorite`: WorldResponse3
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldFavorite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldFavoriteRequest struct via the builder pattern


### Return type

**WorldResponse3**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldMature

> GetWorldMatureResponse UpdateWorldMature(ctx).Execute()

Switch \"Show Mature experiences\" on or off: { showMature }. On is refused while a final step blocks (too young, region, parental controls, safety hold); it never skips the ID check.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldMature(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldMature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldMature`: GetWorldMatureResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldMature`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldMatureRequest struct via the builder pattern


### Return type

**GetWorldMatureResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldMePref

> GetWorldMePrefResponse UpdateWorldMePref(ctx).Execute()

Set your Crossly Carnival choices.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldMePref(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldMePref``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldMePref`: GetWorldMePrefResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldMePref`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldMePrefRequest struct via the builder pattern


### Return type

**GetWorldMePrefResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldParentalChildren

> WorldParentalResponse UpdateWorldParentalChildren(ctx, childId).Execute()

Change a linked account's limits with your PIN: { pin, settings: { maxRating: everyone|teen, chat: on|friends|off, monthlySpendLimitCents (0 = no purchases, null = no limit) } }. Enforced server-side at join, chat, the purchase prompt and checkout.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    childId := "childId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldParentalChildren(context.Background(), childId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldParentalChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldParentalChildren`: WorldParentalResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldParentalChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldParentalChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldParentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldVote

> WorldResponse3 UpdateWorldVote(ctx).Execute()

Thumbs up, down, or null to withdraw: one vote per player per experience, counted only after 10 minutes of play (403 vote_needs_playtime before). A light ranking signal.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsApi.UpdateWorldVote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsApi.UpdateWorldVote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldVote`: WorldResponse3
    fmt.Fprintf(os.Stdout, "Response from `WorldsApi.UpdateWorldVote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldVoteRequest struct via the builder pattern


### Return type

**WorldResponse3**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

