# \VaultMarketsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVaultLend**](VaultMarketsApi.md#CreateVaultLend) | **Post** /v1/vault/{id}/lend | Turn lending on or off for a unit you own in the vault, optionally with your minimum borrow rate (APR in basis points). Returns &#x60;{ lendEnabled, minAprBps }&#x60;.
[**CreateVaultMarketAccount**](VaultMarketsApi.md#CreateVaultMarketAccount) | **Post** /v1/vault-markets/account | Open your account at the custodian partner (once). Returns &#x60;{ status }&#x60;; the custodian verifies it before it is active.
[**CreateVaultMarketAdminMark**](VaultMarketsApi.md#CreateVaultMarketAdminMark) | **Post** /v1/vault-markets/admin-marks | Upload your fund administrator&#39;s prices (&#x60;administrator&#x60;, &#x60;marks: [{ bookKey, day, cents }]&#x60;). Shown beside Crossly&#39;s marks and labelled as yours; they never change a Crossly figure.
[**CreateVaultMarketAdvance**](VaultMarketsApi.md#CreateVaultMarketAdvance) | **Post** /v1/vault-markets/advances | Take an advance against a vault unit that is listed for sale: paid into your custodian account, netted from the sale. Free for 60 days, then the pool rate. Needs a verified ID.
[**CreateVaultMarketAdvanceRelist**](VaultMarketsApi.md#CreateVaultMarketAdvanceRelist) | **Post** /v1/vault-markets/advances/{id}/relist | Relist an advance&#39;s unit at a new price (never so low its sale could not repay the advance).
[**CreateVaultMarketAgentDraftCancel**](VaultMarketsApi.md#CreateVaultMarketAgentDraftCancel) | **Post** /v1/vault-markets/agent/drafts/{id}/cancel | Throw a drafted order away (only while it is still a draft). Returns { ok, draft }.
[**CreateVaultMarketAgentWatch**](VaultMarketsApi.md#CreateVaultMarketAgentWatch) | **Post** /v1/vault-markets/agent/watches | Watch a book key: tell me once when its best bid, lowest ask, last trade or market price crosses a price (cents). Leave out direction for \&quot;crosses\&quot; (whichever way it has to move). Never trades. At most 50 live, each for 90 days.
[**CreateVaultMarketAlert**](VaultMarketsApi.md#CreateVaultMarketAlert) | **Post** /v1/vault-markets/alerts | Alert me when a key&#39;s market price goes above or below a price (fires once).
[**CreateVaultMarketAlertPrefTest**](VaultMarketsApi.md#CreateVaultMarketAlertPrefTest) | **Post** /v1/vault-markets/alert-prefs/test | Send yourself one test alert of a &#x60;kind&#x60; through your own switches. Returns &#x60;{ sent, kind, plan }&#x60;, where plan says per channel: send, off, quiet (held by quiet hours) or unavailable.
[**CreateVaultMarketAuctionLot**](VaultMarketsApi.md#CreateVaultMarketAuctionLot) | **Post** /v1/vault-markets/auctions/lots | Consign a stored, authenticated unit to the next weekly or premier event, with an optional opening price and reserve. The unit is held until the close. Returns &#x60;{ lotId, eventId, endsAt }&#x60;.
[**CreateVaultMarketAuctionLotBid**](VaultMarketsApi.md#CreateVaultMarketAuctionLotBid) | **Post** /v1/vault-markets/auctions/lots/{id}/bids | Leave or raise your maximum on a lot; the system bids for you up to it. A bid in the last 2 minutes extends the lot by 2 minutes. The winner&#39;s saved card is charged at the close through the order book. ID needed above $10,000.
[**CreateVaultMarketAuctionLotWithdraw**](VaultMarketsApi.md#CreateVaultMarketAuctionLotWithdraw) | **Post** /v1/vault-markets/auctions/lots/{id}/withdraw | Withdraw your lot before any bid; the unit goes back on your shelf. Returns &#x60;{ ok }&#x60;.
[**CreateVaultMarketBlock**](VaultMarketsApi.md#CreateVaultMarketBlock) | **Post** /v1/vault-markets/blocks | Place a block buy for an entity: one order booked to the first sub-account in &#x60;plan&#x60; (it pays from its custodian account), split across every sub-account in the plan at the average price, fees included, once it is done and the custodian confirmed each payment. Each share&#39;s cash moves between the entity&#39;s own custodian accounts; its units move only after that settles. Every sub-account must be one you may trade, and its share must fit its limits.
[**CreateVaultMarketBlockAllocate**](VaultMarketsApi.md#CreateVaultMarketBlockAllocate) | **Post** /v1/vault-markets/blocks/{orderId}/allocate | Allocate a done block now (the keeper does it too). Returns &#x60;{ outcome, block }&#x60;: allocated, waiting (a payment or transfer is not settled yet), nothing_filled or none.
[**CreateVaultMarketBullionFamilyPeg**](VaultMarketsApi.md#CreateVaultMarketBullionFamilyPeg) | **Post** /v1/vault-markets/bullion/family-pegs | Set a family-wide spot-pegged quote: one \&quot;any 1 oz gold coin\&quot; quote resting on each product book of a bullion family at the chosen condition (DS, VNDS or Used). When any leg fills it comes off the others and goes back on only for what is left; two fills at once are settled one at a time against what is left. A bid pays from your custodian account. Replaces your live family quote on the same family, condition and side. ID required.
[**CreateVaultMarketBullionPeg**](VaultMarketsApi.md#CreateVaultMarketBullionPeg) | **Post** /v1/vault-markets/bullion/pegs | Set a spot-pegged maker quote on a bullion book key: \&quot;spot plus or minus N cents\&quot; or \&quot;plus or minus N basis points\&quot;, with an optional floor and cap. It rests as a pro limit order re-priced every spot tick and is pulled while spot is stale. Replaces your live quote on the same key and side. ID required.
[**CreateVaultMarketCollectionWorthWatch**](VaultMarketsApi.md#CreateVaultMarketCollectionWorthWatch) | **Post** /v1/vault-markets/collection-worth/watch | Watch every collection item that sits on a vault book key (up to the watchlist cap). Returns &#x60;{ watched, skipped }&#x60;.
[**CreateVaultMarketCreatorAttribution**](VaultMarketsApi.md#CreateVaultMarketCreatorAttribution) | **Post** /v1/vault-markets/creators/attribution | Use a creator code: the creator earns a capped share of the fee Crossly keeps on your Vault Markets trades for a limited window (paid by Crossly, never by you). Only for people new to Vault Markets, once, and never your own code.
[**CreateVaultMarketCreatorCohostInviteRespond**](VaultMarketsApi.md#CreateVaultMarketCreatorCohostInviteRespond) | **Post** /v1/vault-markets/creators/cohost-invites/{id}/respond | For creators: accept or decline an invite to co-host a Market Open launch show (&#x60;accept&#x60; true or false). Accepting needs your vault room verified and adds you to the show as a co-host; your room and the Commons then say you co-host it, with the creator-programme disclosure when it applies.
[**CreateVaultMarketCreditRepay**](VaultMarketsApi.md#CreateVaultMarketCreditRepay) | **Post** /v1/vault-markets/credit/{id}/repay | Repay from your custodian account, interest first. Omit &#x60;amountCents&#x60; to repay everything; pass &#x60;requestId&#x60; to make a retry safe. Repaid in full, the items are free.
[**CreateVaultMarketDcmLink**](VaultMarketsApi.md#CreateVaultMarketDcmLink) | **Post** /v1/vault-markets/dcm/link | Link a partner exchange account to your account (or to a firm sub-account you manage, &#x60;subaccountId&#x60;) so Crossly can show its positions: &#x60;accountRef&#x60; (the account id there) and &#x60;code&#x60; (the one-time link code it gave you). The partner must confirm the code; the code is never stored. Read-only: Crossly never trades on it.
[**CreateVaultMarketDcmUnlink**](VaultMarketsApi.md#CreateVaultMarketDcmUnlink) | **Post** /v1/vault-markets/dcm/unlink | Unlink the partner exchange account from your account (or the firm sub-account &#x60;subaccountId&#x60;): Crossly stops reading it and its hedges leave portfolio margin.
[**CreateVaultMarketDealerApply**](VaultMarketsApi.md#CreateVaultMarketDealerApply) | **Post** /v1/vault-markets/dealer/apply | Apply to be a dealer (ID check first). Ops review; approved dealers get raised limits and their bids show as dealer bids.
[**CreateVaultMarketDealerBid**](VaultMarketsApi.md#CreateVaultMarketDealerBid) | **Post** /v1/vault-markets/dealer/bids | Place a standing dealer bid: a resting GTC/GTD limit buy, flagged as a dealer bid. Approved dealers only.
[**CreateVaultMarketDealerPool**](VaultMarketsApi.md#CreateVaultMarketDealerPool) | **Post** /v1/vault-markets/dealer/pools | Open a curve pool on a key (approved dealers, ID checked): linear (delta in cents) or exponential (delta in bps), spot, fee, levels per side, units it may sell, and a buy budget from your custodian account. Its ladder goes up at once.
[**CreateVaultMarketDealerPoolClose**](VaultMarketsApi.md#CreateVaultMarketDealerPoolClose) | **Post** /v1/vault-markets/dealer/pools/{id}/close | Close a curve pool: its legs come off the book; what filled stays filled.
[**CreateVaultMarketDealerPoolPreview**](VaultMarketsApi.md#CreateVaultMarketDealerPoolPreview) | **Post** /v1/vault-markets/dealer/pools/preview | The bids and asks a set of curve settings would show, and why each side stops where it does. Places nothing.
[**CreateVaultMarketDeposit**](VaultMarketsApi.md#CreateVaultMarketDeposit) | **Post** /v1/vault-markets/deposits | Open a ship-in deposit for an ungraded catalogue item (&#x60;skuId&#x60;, &#x60;variantId&#x60;, &#x60;condition&#x60; DS / VNDS / Used). It is inbound until the vault receives and checks it, then stored in your name. Returns &#x60;{ vaultItem: { id, status } }&#x60;.
[**CreateVaultMarketDonation**](VaultMarketsApi.md#CreateVaultMarketDonation) | **Post** /v1/vault-markets/donations | Give a unit to Crossly Cares (once the charity can receive; otherwise its gate message). Returns &#x60;{ id, markCents }&#x60;.
[**CreateVaultMarketEntity**](VaultMarketsApi.md#CreateVaultMarketEntity) | **Post** /v1/vault-markets/entities | Open an entity account: you become its admin and the custodian partner starts KYB (its own flow; Crossly keeps only its reference). Needs a verified ID and the custodian (503 names what is missing).
[**CreateVaultMarketEntityEcp**](VaultMarketsApi.md#CreateVaultMarketEntityEcp) | **Post** /v1/vault-markets/entities/{id}/ecp | Admin: record the entity&#39;s attestation that it is an eligible contract participant. Raising its caps is still an ops decision.
[**CreateVaultMarketEntityHalt**](VaultMarketsApi.md#CreateVaultMarketEntityHalt) | **Post** /v1/vault-markets/entities/{id}/halts | Admin or risk: the firm kill switch. Halts the entity (no &#x60;subaccountId&#x60;) or one sub-account: open orders on it are cancelled and no new ones are placed until cleared.
[**CreateVaultMarketEntityHaltClear**](VaultMarketsApi.md#CreateVaultMarketEntityHaltClear) | **Post** /v1/vault-markets/entities/{id}/halts/{haltId}/clear | Admin or risk: lift a firm halt.
[**CreateVaultMarketEntityInvitation**](VaultMarketsApi.md#CreateVaultMarketEntityInvitation) | **Post** /v1/vault-markets/entities/{id}/invitation | Accept or decline an invitation (&#x60;accept&#x60;). Trading, admin and risk roles need a verified ID to accept.
[**CreateVaultMarketEntityKyb**](VaultMarketsApi.md#CreateVaultMarketEntityKyb) | **Post** /v1/vault-markets/entities/{id}/kyb | Refresh the entity&#39;s KYB status from the custodian (and its onboarding link while KYB is open).
[**CreateVaultMarketEntityMember**](VaultMarketsApi.md#CreateVaultMarketEntityMember) | **Post** /v1/vault-markets/entities/{id}/members | Admin: invite a Crossly user by email as admin, trader, risk, ops or viewer (a viewer can never place an order), with optional limits and sub-accounts.
[**CreateVaultMarketEntitySubaccount**](VaultMarketsApi.md#CreateVaultMarketEntitySubaccount) | **Post** /v1/vault-markets/entities/{id}/subaccounts | Admin: open a sub-account once the custodian approved the entity; its limits are held to the entity&#39;s caps.
[**CreateVaultMarketEntitySubaccountClose**](VaultMarketsApi.md#CreateVaultMarketEntitySubaccountClose) | **Post** /v1/vault-markets/entities/{id}/subaccounts/{subId}/close | Admin: close a sub-account that holds nothing and has no open orders.
[**CreateVaultMarketFollowCopyWatchAlertRead**](VaultMarketsApi.md#CreateVaultMarketFollowCopyWatchAlertRead) | **Post** /v1/vault-markets/follows/copy-watch/alerts/read | Mark every copy-watch alert read. Returns &#x60;{ marked }&#x60;.
[**CreateVaultMarketGift**](VaultMarketsApi.md#CreateVaultMarketGift) | **Post** /v1/vault-markets/gifts | Give a unit in your vault to another Crossly user (by @username or email). It stays on its shelf; they have 7 days to accept. ID needed above $1,000 at the mark. Returns &#x60;{ id, expiresAt }&#x60;.
[**CreateVaultMarketGiftAccept**](VaultMarketsApi.md#CreateVaultMarketGiftAccept) | **Post** /v1/vault-markets/gifts/{id}/accept | Accept a gift: the unit becomes yours in place. Returns &#x60;{ status }&#x60;.
[**CreateVaultMarketGiftCancel**](VaultMarketsApi.md#CreateVaultMarketGiftCancel) | **Post** /v1/vault-markets/gifts/{id}/cancel | Cancel a gift you sent before it is accepted. Returns &#x60;{ status }&#x60;.
[**CreateVaultMarketGiftDecline**](VaultMarketsApi.md#CreateVaultMarketGiftDecline) | **Post** /v1/vault-markets/gifts/{id}/decline | Decline a gift: the unit stays with the giver. Returns &#x60;{ status }&#x60;.
[**CreateVaultMarketGradeToVault**](VaultMarketsApi.md#CreateVaultMarketGradeToVault) | **Post** /v1/vault-markets/grade-to-vault | Start Grade to Vault on one item (&#x60;source&#x60; inventory or vault_unit). With &#x60;runId&#x60; the grading slot is booked at once (free; your saved card is charged only if the run locks). Nothing is listed until you confirm a price after the grade.
[**CreateVaultMarketGradeToVaultBook**](VaultMarketsApi.md#CreateVaultMarketGradeToVaultBook) | **Post** /v1/vault-markets/grade-to-vault/{id}/book | Book (or re-book) the grading slot of a flow still at \&quot;picked\&quot;. Answers 503 naming GRADING_HUB_ADDRESS when runs are not open.
[**CreateVaultMarketGradeToVaultCancel**](VaultMarketsApi.md#CreateVaultMarketGradeToVaultCancel) | **Post** /v1/vault-markets/grade-to-vault/{id}/cancel | Stop the flow before the card goes to the grader (still picked, or its slot only reserved: nothing was charged).
[**CreateVaultMarketGradeToVaultKeep**](VaultMarketsApi.md#CreateVaultMarketGradeToVaultKeep) | **Post** /v1/vault-markets/grade-to-vault/{id}/keep | Keep the graded unit in your vault without listing it. The flow ends; list it any time from the vault.
[**CreateVaultMarketGradeToVaultList**](VaultMarketsApi.md#CreateVaultMarketGradeToVaultList) | **Post** /v1/vault-markets/grade-to-vault/{id}/list | Confirm the price for the grade the unit came back with (&#x60;confirm: true&#x60;, &#x60;gradeKey&#x60; must match the unit). Pre-trade risk, the price band and the ID line apply, then one shelf ask is placed on the Crossly order book.
[**CreateVaultMarketGradeToVaultShipped**](VaultMarketsApi.md#CreateVaultMarketGradeToVaultShipped) | **Post** /v1/vault-markets/grade-to-vault/{id}/shipped | You mailed the card to the grading hub (an inventory card, once its run locked).
[**CreateVaultMarketGrading**](VaultMarketsApi.md#CreateVaultMarketGrading) | **Post** /v1/vault-markets/grading | Send a raw vault unit on the grading hub&#39;s shelf to a Grading Run (&#x60;runId&#x60;). It is held until it comes back, graded, to the same vault item. The slot is free; your saved card is charged only if the run locks.
[**CreateVaultMarketGradingPull**](VaultMarketsApi.md#CreateVaultMarketGradingPull) | **Post** /v1/vault-markets/grading/{submissionId}/pull | Pull a vault unit out of its run before the run locks (nothing was charged). It goes back on the shelf.
[**CreateVaultMarketHallLoan**](VaultMarketsApi.md#CreateVaultMarketHallLoan) | **Post** /v1/vault-markets/hall/loans | Lend a unit in your Crossly vault to the Hall of Grails: { vaultItemId, credit: &#39;name&#39; | &#39;anonymous&#39; }. A standing loan of display rights only: the unit goes on a plinth, under a plate naming you (or \&quot;a collector\&quot;), in every weekly exhibition whose category it fits. It never moves and stays yours to list, sell or withdraw. Earns reputation per exhibition shown; no cash. Resumes a loan you opted out of while it still shows.
[**CreateVaultMarketHallLoanCredit**](VaultMarketsApi.md#CreateVaultMarketHallLoanCredit) | **Post** /v1/vault-markets/hall/loans/{id}/credit | Change how the plate on a lent unit&#39;s plinth names you, from now: { credit: &#39;name&#39; } (your display name) or { credit: &#39;anonymous&#39; } (\&quot;a collector\&quot;).
[**CreateVaultMarketHallLoanOptOut**](VaultMarketsApi.md#CreateVaultMarketHallLoanOptOut) | **Post** /v1/vault-markets/hall/loans/{id}/opt-out | Opt a Hall of Grails loan out, at any time. A unit on a plinth this week finishes this exhibition (status ending, showingUntil); otherwise the loan ends now (status ended).
[**CreateVaultMarketIdentityStart**](VaultMarketsApi.md#CreateVaultMarketIdentityStart) | **Post** /v1/vault-markets/identity/start | Start (or resume) the one-time ID check the money features and larger orders need. Returns &#x60;{ url }&#x60; to open, or &#x60;{ alreadyVerified: true }&#x60;.
[**CreateVaultMarketIndex**](VaultMarketsApi.md#CreateVaultMarketIndex) | **Post** /v1/vault-markets/indexes | Make an index of your own: 1 to 50 book keys, weighted by mark or equally. Up to 10 per account.
[**CreateVaultMarketIndexBuy**](VaultMarketsApi.md#CreateVaultMarketIndexBuy) | **Post** /v1/vault-markets/indexes/{slug}/buy | Buy the basket: one limit buy per key in the index for 1-5 whole units each, at most &#x60;maxBpsOverMark&#x60; over each key&#39;s mark, IOC or GTC. Every unit is a real item that lands in your vault. Idempotent on &#x60;clientBasketId&#x60;; returns a result per key.
[**CreateVaultMarketIndexComplaint**](VaultMarketsApi.md#CreateVaultMarketIndexComplaint) | **Post** /v1/vault-markets/index-complaints | File a complaint about an index or the method (&#x60;slug&#x60; optional). Every complaint is answered. Returns &#x60;{ id, status, createdAt }&#x60;.
[**CreateVaultMarketIndexConsultationComment**](VaultMarketsApi.md#CreateVaultMarketIndexConsultationComment) | **Post** /v1/vault-markets/index-consultations/{id}/comments | Comment on an open consultation (5 to 5,000 characters; up to 5 per person). Comments are published without your name. Returns &#x60;{ ok }&#x60;.
[**CreateVaultMarketInstantExit**](VaultMarketsApi.md#CreateVaultMarketInstantExit) | **Post** /v1/vault-markets/instant-exit | Sell a unit now at the quoted price: a sell IOC at &#x60;priceCents&#x60; through the pro engine. Refused, with the new quote, if the best bid dropped below it. Crossly never buys or funds it.
[**CreateVaultMarketJoint**](VaultMarketsApi.md#CreateVaultMarketJoint) | **Post** /v1/vault-markets/joint | Open a joint vault (ID first): you hold full rights and accept the current joint terms. &#x60;invite&#x60; sends invitations (@username or email, each with view, trade or full rights). Returns &#x60;{ id, status, invited }&#x60;.
[**CreateVaultMarketJointClose**](VaultMarketsApi.md#CreateVaultMarketJointClose) | **Post** /v1/vault-markets/joint/{id}/close | Close the joint vault (full rights): open joint orders are cancelled and every unit goes back to its owner. Returns &#x60;{ status, unitsReturned }&#x60;.
[**CreateVaultMarketJointInvite**](VaultMarketsApi.md#CreateVaultMarketJointInvite) | **Post** /v1/vault-markets/joint/{id}/invite | Invite a Crossly user (full rights only). They have 14 days and show ID to join. Returns &#x60;{ userId, rights, expiresAt }&#x60;.
[**CreateVaultMarketJointMemberRemove**](VaultMarketsApi.md#CreateVaultMarketJointMemberRemove) | **Post** /v1/vault-markets/joint/{id}/members/{userId}/remove | Remove a member (full rights), or leave when userId is you. Their units go back to their own account and their joint orders are cancelled. Returns &#x60;{ userId, status, unitsReturned, jointStatus }&#x60;.
[**CreateVaultMarketJointMemberRight**](VaultMarketsApi.md#CreateVaultMarketJointMemberRight) | **Post** /v1/vault-markets/joint/{id}/members/{userId}/rights | Set a member rights to view, trade or full (full rights only; the vault always keeps a full member). Returns &#x60;{ userId, rights }&#x60;.
[**CreateVaultMarketJointRespond**](VaultMarketsApi.md#CreateVaultMarketJointRespond) | **Post** /v1/vault-markets/joint/{id}/respond | Accept (ID first; accepts the current joint terms) or decline an invitation. Returns &#x60;{ status }&#x60;.
[**CreateVaultMarketJointTerm**](VaultMarketsApi.md#CreateVaultMarketJointTerm) | **Post** /v1/vault-markets/joint/{id}/terms | Accept the current joint terms (needed to trade after the terms version changes). Returns &#x60;{ termsVersion }&#x60;.
[**CreateVaultMarketJointTrade**](VaultMarketsApi.md#CreateVaultMarketJointTrade) | **Post** /v1/vault-markets/joint/{id}/trades | Place a limit order for the joint vault (trade or full rights, ID shown). A buy (bookKey) is placed on you, paid from your own Vault Markets account, and what it buys is booked to the joint vault. A sell (vaultItemId of a unit in it) is placed on the unit owner of record, who gets the proceeds; pre-trade checks run against you too. Returns &#x60;{ tradeId, side, bookKey, ownerUserId, order }&#x60;.
[**CreateVaultMarketJointTradeCancel**](VaultMarketsApi.md#CreateVaultMarketJointTradeCancel) | **Post** /v1/vault-markets/joint/{id}/trades/{tradeId}/cancel | Cancel a working joint order (trade or full rights); what filled stays filled. Returns &#x60;{ tradeId, order }&#x60;.
[**CreateVaultMarketJointUnit**](VaultMarketsApi.md#CreateVaultMarketJointUnit) | **Post** /v1/vault-markets/joint/{id}/units | Put units you own into the joint vault (each stored and free). They keep you as owner of record and stay on the shelf. Returns &#x60;{ added, vaultItemIds }&#x60;.
[**CreateVaultMarketJointUnitRemove**](VaultMarketsApi.md#CreateVaultMarketJointUnitRemove) | **Post** /v1/vault-markets/joint/{id}/units/remove | Take units out of the joint vault, back to their owner own account (yours at any rights; anyone else with full rights). Returns &#x60;{ removed, vaultItemIds }&#x60;.
[**CreateVaultMarketLargeTraderReportFile**](VaultMarketsApi.md#CreateVaultMarketLargeTraderReportFile) | **Post** /v1/vault-markets/large-trader-reports/{id}/file | File (or amend, or refile) a large-trader report: why you hold it (purpose), who you act for, other accounts you control, whether you plan to add, and a statement of 20+ characters. Returns the report.
[**CreateVaultMarketLearnQuiz**](VaultMarketsApi.md#CreateVaultMarketLearnQuiz) | **Post** /v1/vault-markets/learn/{concept}/quiz | Answer a concept&#39;s quiz (&#x60;answers&#x60;: the chosen choice index per question). Graded on the server; 5 tries an hour.
[**CreateVaultMarketLiveStreamFeaturedLot**](VaultMarketsApi.md#CreateVaultMarketLiveStreamFeaturedLot) | **Post** /v1/vault-markets/live-streams/{id}/featured-lot | Feature an open lot on your live show (or &#x60;lotId: null&#x60; to take it off). Viewers hear LIVE_PIN_CHANGED. Returns &#x60;{ ok }&#x60;.
[**CreateVaultMarketLoan**](VaultMarketsApi.md#CreateVaultMarketLoan) | **Post** /v1/vault-markets/loans | Borrow against stored vault units: they are pledged (no listing, lending or withdrawal) and the market pool lends &#x60;amountCents&#x60; into your custodian account. Needs a verified ID.
[**CreateVaultMarketLoanRecall**](VaultMarketsApi.md#CreateVaultMarketLoanRecall) | **Post** /v1/vault-markets/loans/{id}/recall | Recall a lent unit: the short has 3 days to return one. Returns &#x60;{ recallDueAt }&#x60;.
[**CreateVaultMarketLoanUnit**](VaultMarketsApi.md#CreateVaultMarketLoanUnit) | **Post** /v1/vault-markets/loans/{id}/units | Pledge more stored units to a loan (answers a loan call, or makes room).
[**CreateVaultMarketLoanUnitRelease**](VaultMarketsApi.md#CreateVaultMarketLoanUnitRelease) | **Post** /v1/vault-markets/loans/{id}/units/{vaultItemId}/release | Free one pledged unit, if what is owed still fits inside the other units&#39; limit.
[**CreateVaultMarketLoanUnitSell**](VaultMarketsApi.md#CreateVaultMarketLoanUnitSell) | **Post** /v1/vault-markets/loans/{id}/units/{vaultItemId}/sell | Sell a pledged unit to repay its loan: it is listed at &#x60;priceCents&#x60; (at least the sale floor) and stays pledged until it sells; the proceeds repay the loan first and the rest is yours. Calling it again re-prices it.
[**CreateVaultMarketLocate**](VaultMarketsApi.md#CreateVaultMarketLocate) | **Post** /v1/vault-markets/locates | Hold up to &#x60;qty&#x60; lendable units of &#x60;bookKey&#x60; for 15 minutes so a short can open against them (&#x60;termDays&#x60; 30/90 for a term loan). Needs lending (VAULT_LENDING_TERMS_VERSION) and the ID check. Returns the locate with how many units it holds.
[**CreateVaultMarketLong**](VaultMarketsApi.md#CreateVaultMarketLong) | **Post** /v1/vault-markets/longs | Buy &#x60;units&#x60; units (default 1; at most 25/10/3/1 by liquidity tier A/B/C/D) into your vault with leverage (up to 2x) from the market staking pool, paying up to &#x60;maxCents&#x60; each. A partial fill opens with the units bought and the rest is cancelled. Returns &#x60;{ positionId, matched, units }&#x60;.
[**CreateVaultMarketMidpoint**](VaultMarketsApi.md#CreateVaultMarketMidpoint) | **Post** /v1/vault-markets/midpoint | A hidden block order that trades only at the firm-book midpoint, with another midpoint order of at least your minimum size.
[**CreateVaultMarketMidpointCancel**](VaultMarketsApi.md#CreateVaultMarketMidpointCancel) | **Post** /v1/vault-markets/midpoint/{id}/cancel | Cancel a midpoint order.
[**CreateVaultMarketMmDmm**](VaultMarketsApi.md#CreateVaultMarketMmDmm) | **Post** /v1/vault-markets/mm/dmm | Apply to be the designated market maker on a tier B or C key. Ops approves; one per key.
[**CreateVaultMarketMmLeave**](VaultMarketsApi.md#CreateVaultMarketMmLeave) | **Post** /v1/vault-markets/mm/{id}/leave | Leave a market-maker programme.
[**CreateVaultMarketMmLp**](VaultMarketsApi.md#CreateVaultMarketMmLp) | **Post** /v1/vault-markets/mm/lp | Join a key as a Registered Liquidity Provider (ID needed). Meeting coverage earns the LP rates.
[**CreateVaultMarketMmMassCancel**](VaultMarketsApi.md#CreateVaultMarketMmMassCancel) | **Post** /v1/vault-markets/mm/mass-cancel | Cancel every open order (one key, or only quotes) and every resting leg in one statement: nothing fills between the first cancel and the last.
[**CreateVaultMarketMmMmpReset**](VaultMarketsApi.md#CreateVaultMarketMmMmpReset) | **Post** /v1/vault-markets/mm/mmp/reset | End a market-maker protection freeze early.
[**CreateVaultMarketOpeningWaitlist**](VaultMarketsApi.md#CreateVaultMarketOpeningWaitlist) | **Post** /v1/vault-markets/opening/{slug}/waitlist | Join a key&#39;s waitlist, optionally with a friend&#39;s invite code (&#x60;ref&#x60;). Inviting moves you up the queue; nobody is paid. Returns your place and your own code.
[**CreateVaultMarketP2pLoanCall**](VaultMarketsApi.md#CreateVaultMarketP2pLoanCall) | **Post** /v1/vault-markets/p2p/loans/{id}/call | Lender: call the loan (once it is a day old). A refinancing auction runs for about 30 hours, the rate a new lender gets climbing to the ceiling; repaid or refinanced, you are paid in full; neither by the end, the unit moves to your vault.
[**CreateVaultMarketP2pLoanCancelCall**](VaultMarketsApi.md#CreateVaultMarketP2pLoanCancelCall) | **Post** /v1/vault-markets/p2p/loans/{id}/cancel-call | Lender: take back your call before any other lender has taken the loan over. The refinancing auction stops and the loan carries on as before on its terms.
[**CreateVaultMarketP2pLoanRefiAuction**](VaultMarketsApi.md#CreateVaultMarketP2pLoanRefiAuction) | **Post** /v1/vault-markets/p2p/loans/{id}/refi-auction | Borrower: look for a cheaper lender. A refinancing auction runs for &#x60;hours&#x60; (1-72, default 24); the rate a new lender gets climbs from &#x60;startAprBps&#x60; (default half the loan rate) to just under the loan rate now. The first lender to take it pays your lender off and the loan carries on at that lower rate. Nobody by the end: nothing changes.
[**CreateVaultMarketP2pLoanRefiAuctionCancel**](VaultMarketsApi.md#CreateVaultMarketP2pLoanRefiAuctionCancel) | **Post** /v1/vault-markets/p2p/loans/{id}/refi-auction/cancel | Borrower: stop looking for a cheaper lender before anyone has taken the loan over. The loan carries on unchanged.
[**CreateVaultMarketP2pLoanRefinance**](VaultMarketsApi.md#CreateVaultMarketP2pLoanRefinance) | **Post** /v1/vault-markets/p2p/loans/{id}/refinance | Take over a loan in a refinancing auction (a lender call or a borrower search for a cheaper lender): you pay the old lender everything owed now (refused above &#x60;maxPayoffCents&#x60;), which becomes your loan to the same borrower at the auction rate now. Needs a verified ID, an active account and the current lending terms. Pass &#x60;requestId&#x60; to make a retry safe.
[**CreateVaultMarketP2pLoanRepay**](VaultMarketsApi.md#CreateVaultMarketP2pLoanRepay) | **Post** /v1/vault-markets/p2p/loans/{id}/repay | Borrower: repay from your custodian account to the lender. Without &#x60;amountCents&#x60;, in full (principal plus interest to the second; the unit is free once the custodian has it). With &#x60;amountCents&#x60; below what is owed, a part payment: interest first, then principal, at least the interest owed now and at least $1; the loan stays open on the principal left. Pass &#x60;requestId&#x60; to make a retry safe.
[**CreateVaultMarketP2pOffer**](VaultMarketsApi.md#CreateVaultMarketP2pOffer) | **Post** /v1/vault-markets/p2p/offers | Offer to lend: &#x60;principalCents&#x60; per loan at &#x60;aprBps&#x60; a year against an item (&#x60;bookKey&#x60;) or one unit (&#x60;vaultItemId&#x60;), for up to &#x60;maxLoans&#x60; loans, lasting &#x60;expiresInDays&#x60; (default 90). Nothing is held: the custodian pays from your own account when a borrower takes it. Needs a verified ID, an active account and the current lending terms.
[**CreateVaultMarketP2pOfferCancel**](VaultMarketsApi.md#CreateVaultMarketP2pOfferCancel) | **Post** /v1/vault-markets/p2p/offers/{id}/cancel | Withdraw one of your open offers. Loans already taken from it carry on.
[**CreateVaultMarketP2pOfferTake**](VaultMarketsApi.md#CreateVaultMarketP2pOfferTake) | **Post** /v1/vault-markets/p2p/offers/{id}/take | Borrow: take an offer against one of your stored vault units (&#x60;vaultItemId&#x60;). The unit is pledged (no listing, lending or withdrawal) and the principal is paid into your custodian account. Interest accrues by the second; repay any time. Pass &#x60;requestId&#x60; to make a retry safe.
[**CreateVaultMarketPaperMissionMarkCheck**](VaultMarketsApi.md#CreateVaultMarketPaperMissionMarkCheck) | **Post** /v1/vault-markets/paper/missions/mark-check | Answer the \&quot;Read the mark\&quot; mission: the quality letter (A, B or C) you read on a market page, checked against the stored mark. One answer per market, 10 a day. Returns &#x60;{ bookKey, answer, correct, quality, mission, next }&#x60;.
[**CreateVaultMarketPaperOrder**](VaultMarketsApi.md#CreateVaultMarketPaperOrder) | **Post** /v1/vault-markets/paper/orders | Place a paper order (play money only). It takes visible real liquidity now at book prices with the taker fee simulated; a GTC limit rests the rest with a queue-position estimate, filled later as the real tape reaches it. Sells only paper units you hold (no paper shorts). Never touches the real book. Returns &#x60;{ order, fills }&#x60;.
[**CreateVaultMarketPaperOrderCancel**](VaultMarketsApi.md#CreateVaultMarketPaperOrderCancel) | **Post** /v1/vault-markets/paper/orders/{id}/cancel | Cancel a working paper order: the play cash or paper units it held come back. Returns &#x60;{ order }&#x60;.
[**CreateVaultMarketPaperReset**](VaultMarketsApi.md#CreateVaultMarketPaperReset) | **Post** /v1/vault-markets/paper/reset | Reset your paper account to $10,000 of play money: working paper orders are cancelled, paper positions cleared, and your leaderboard week starts again. Once a day.
[**CreateVaultMarketPartnerApply**](VaultMarketsApi.md#CreateVaultMarketPartnerApply) | **Post** /v1/vault-markets/partners/apply | Apply as a Vault Drop-off Partner (a card shop that takes walk-ins for the vault). Needs the one-time ID check; ops approve.
[**CreateVaultMarketPartnerBatch**](VaultMarketsApi.md#CreateVaultMarketPartnerBatch) | **Post** /v1/vault-markets/partners/batches | An approved partner sends walk-ins in under its code: each item (catalogue variant, condition, grade, photos) becomes an inbound vault unit; the box gets one prepaid label. The partner earns a capped share of Crossly&#39;s fee on each unit&#39;s first sale.
[**CreateVaultMarketPoolStake**](VaultMarketsApi.md#CreateVaultMarketPoolStake) | **Post** /v1/vault-markets/pool/stakes | Stake CBX you hold at the custodian into the market pool (base units, as a string). Returns &#x60;{ stakeId, shares }&#x60;.
[**CreateVaultMarketPoolStakeUnstake**](VaultMarketsApi.md#CreateVaultMarketPoolStakeUnstake) | **Post** /v1/vault-markets/pool/stakes/{id}/unstake | Unstake: paid out after the 7-day cooldown. Returns &#x60;{ cooldownUntil }&#x60;.
[**CreateVaultMarketPortfolioMarginEnroll**](VaultMarketsApi.md#CreateVaultMarketPortfolioMarginEnroll) | **Post** /v1/vault-markets/portfolio-margin/enroll | Admin: put a firm sub-account under portfolio margin. Needs VAULT_LEVERAGE_VENUE, your ID, an ECP-attested entity that passed KYB, and the minimum equity. Its positions are then margined on the whole account, recalculated every 5 minutes; calls and auctions are unchanged. Audited.
[**CreateVaultMarketPortfolioMarginLeave**](VaultMarketsApi.md#CreateVaultMarketPortfolioMarginLeave) | **Post** /v1/vault-markets/portfolio-margin/leave | Admin: take a firm sub-account off portfolio margin. Each position is then held to its own tier floors again and may be called at once. Audited.
[**CreateVaultMarketPortfolioMarginWhatIf**](VaultMarketsApi.md#CreateVaultMarketPortfolioMarginWhatIf) | **Post** /v1/vault-markets/portfolio-margin/what-if | The portfolio margin requirement of a firm sub-account before and after a trade (&#x60;bookKey&#x60;, &#x60;side&#x60; buy or sell, &#x60;units&#x60;), at the Crossly mark with fees left out. Changes nothing.
[**CreateVaultMarketPositionCancel**](VaultMarketsApi.md#CreateVaultMarketPositionCancel) | **Post** /v1/vault-markets/positions/{id}/cancel | Cancel the working order: an unfilled opening order cancels the position; an unfilled cover or close leaves it open. Returns &#x60;{ ok }&#x60;.
[**CreateVaultMarketPositionClose**](VaultMarketsApi.md#CreateVaultMarketPositionClose) | **Post** /v1/vault-markets/positions/{id}/close | Close a leveraged long: all its units, or &#x60;units&#x60; of them, each sold for at least &#x60;minCents&#x60;; each sale repays the pool first.
[**CreateVaultMarketPositionCollateral**](VaultMarketsApi.md#CreateVaultMarketPositionCollateral) | **Post** /v1/vault-markets/positions/{id}/collateral | Answer a margin call: a short posts more collateral (USD, or CBX at half value); a leveraged long pays down its pool loan. Returns the position.
[**CreateVaultMarketPositionCover**](VaultMarketsApi.md#CreateVaultMarketPositionCover) | **Post** /v1/vault-markets/positions/{id}/cover | Cover a short: all its units, or &#x60;units&#x60; of them, each buying any unit of the key into the vault for up to &#x60;maxCents&#x60;; each goes back to its lender.
[**CreateVaultMarketPriceWidget**](VaultMarketsApi.md#CreateVaultMarketPriceWidget) | **Post** /v1/vault-markets/price-widgets | Make a price widget for one vault market (&#x60;kind&#x60; mark + &#x60;bookKey&#x60;) or one public index (&#x60;kind&#x60; index + &#x60;indexSlug&#x60;): range, layout (compact or chart), theme (auto, light, dark), optionally your own affiliate code for the link back and the sites allowed to show it (empty &#x3D; any). Returns the widget with its snippets.
[**CreateVaultMarketPriceWidgetPreview**](VaultMarketsApi.md#CreateVaultMarketPriceWidgetPreview) | **Post** /v1/vault-markets/price-widgets/preview | What a widget would show before you save it: &#x60;view&#x60; (value, move, quality, n, as-of, link back) and &#x60;html&#x60;, the exact page the iframe serves. Nothing is stored.
[**CreateVaultMarketProBracket**](VaultMarketsApi.md#CreateVaultMarketProBracket) | **Post** /v1/vault-markets/pro/brackets | A buy order with a take-profit and/or stop-loss attached; the exits arm as an OCO pair for what the entry fills.
[**CreateVaultMarketProLinkCancel**](VaultMarketsApi.md#CreateVaultMarketProLinkCancel) | **Post** /v1/vault-markets/pro/links/{id}/cancel | Cancel a linked order: its open child orders come off; a bracket&#39;s entry stays.
[**CreateVaultMarketProOco**](VaultMarketsApi.md#CreateVaultMarketProOco) | **Post** /v1/vault-markets/pro/oco | Two orders on one key, one cancels the other: when one fills or its stop triggers, the other is cancelled (a partial fill shrinks it).
[**CreateVaultMarketProOrder**](VaultMarketsApi.md#CreateVaultMarketProOrder) | **Post** /v1/vault-markets/pro/orders | Place an order on a vault book: limit, market (protected, IOC), stop or stop-limit; GTC, GTD, IOC or FOK; post-only; an iceberg &#x60;displayQuantity&#x60;; a TWAP or VWAP algo over a window. Idempotent on &#x60;clientOrderId&#x60;. &#x60;subaccountId&#x60; trades for a firm sub-account you may trade (your role, its KYB, its limits, no firm halt; buys pay from its custodian account and land in it).
[**CreateVaultMarketProOrderCancel**](VaultMarketsApi.md#CreateVaultMarketProOrderCancel) | **Post** /v1/vault-markets/pro/orders/{id}/cancel | Cancel an order; what filled stays filled.
[**CreateVaultMarketProOrderCancelAll**](VaultMarketsApi.md#CreateVaultMarketProOrderCancelAll) | **Post** /v1/vault-markets/pro/orders/cancel-all | Cancel every open order, or every one on a book key. Returns &#x60;{ cancelled }&#x60;.
[**CreateVaultMarketProOrderReplace**](VaultMarketsApi.md#CreateVaultMarketProOrderReplace) | **Post** /v1/vault-markets/pro/orders/{id}/replace | Amend an order&#39;s price, stop or quantity (cancel-replace: it re-queues at its new price).
[**CreateVaultMarketProQuote**](VaultMarketsApi.md#CreateVaultMarketProQuote) | **Post** /v1/vault-markets/pro/quotes | Market makers: set a two-sided post-only quote on a book key (replaces the last). With cancel-on-disconnect (default) heartbeat every 30 s or it is pulled.
[**CreateVaultMarketProQuoteHeartbeat**](VaultMarketsApi.md#CreateVaultMarketProQuoteHeartbeat) | **Post** /v1/vault-markets/pro/quotes/heartbeat | Keep your quotes alive (cancel-on-disconnect). Returns &#x60;{ quotes, at }&#x60;.
[**CreateVaultMarketProTrailingStop**](VaultMarketsApi.md#CreateVaultMarketProTrailingStop) | **Post** /v1/vault-markets/pro/trailing-stops | A trailing stop: its stop follows the best price since it was placed, &#x60;trailBps&#x60; or &#x60;trailCents&#x60; behind (under the high for a sell, over the low for a buy).
[**CreateVaultMarketProfileFollow**](VaultMarketsApi.md#CreateVaultMarketProfileFollow) | **Post** /v1/vault-markets/profiles/{handle}/follow | Follow someone with a public vault: their new vault buys appear in your feed. No copy trading.
[**CreateVaultMarketRecurring**](VaultMarketsApi.md#CreateVaultMarketRecurring) | **Post** /v1/vault-markets/recurring | Start a recurring buy: one whole unit of a key every week, two weeks or month, at no more than &#x60;maxPriceCents&#x60; and/or the mark plus &#x60;capBpsOverMark&#x60;; card (with an optional backup) or custodian; optional auto-lend.
[**CreateVaultMarketRfq**](VaultMarketsApi.md#CreateVaultMarketRfq) | **Post** /v1/vault-markets/rfq | Ask for a quote on a block: dealers, the key&#39;s DMM and LPs, and holders are told, anonymously. Answers are firm for 60 s. Your limit is never shown.
[**CreateVaultMarketRfqAccept**](VaultMarketsApi.md#CreateVaultMarketRfqAccept) | **Post** /v1/vault-markets/rfq/{id}/accept | Take one answer: it fills as a block through the normal fill path (fees by role, self-trade prevention, halts, LULD).
[**CreateVaultMarketRfqCancel**](VaultMarketsApi.md#CreateVaultMarketRfqCancel) | **Post** /v1/vault-markets/rfq/{id}/cancel | Cancel your open quote request.
[**CreateVaultMarketRfqRespond**](VaultMarketsApi.md#CreateVaultMarketRfqRespond) | **Post** /v1/vault-markets/rfq/{id}/respond | Answer a quote request for the full quantity with a firm price, good for 60 s. Selling needs the units free in your vault; buying pays from your Vault Markets custodian account.
[**CreateVaultMarketRfqResponseWithdraw**](VaultMarketsApi.md#CreateVaultMarketRfqResponseWithdraw) | **Post** /v1/vault-markets/rfq/responses/{id}/withdraw | Withdraw your live answer to a quote request.
[**CreateVaultMarketSetGoal**](VaultMarketsApi.md#CreateVaultMarketSetGoal) | **Post** /v1/vault-markets/set-goals | Start a goal on a catalogue set (category, brandSlug, modelSlug from the set search), optionally with a target: one condition, and any grade, ungraded, or one exact grade. One live goal per set.
[**CreateVaultMarketSetGoalBasket**](VaultMarketsApi.md#CreateVaultMarketSetGoalBasket) | **Post** /v1/vault-markets/set-goals/{id}/basket | Place the \&quot;complete my set\&quot; basket bid: one GTC limit buy per missing piece, under a hard budget, paid by one of your saved cards or your custodian account. Every leg passes pre-trade risk and the ID line like any order. Idempotent on clientBasketId. One working basket per goal; its legs are cancelled once the set is complete.
[**CreateVaultMarketSetGoalBasketCancel**](VaultMarketsApi.md#CreateVaultMarketSetGoalBasketCancel) | **Post** /v1/vault-markets/set-goals/{id}/basket/cancel | Cancel the goal working basket bid: every working leg comes off the book; what already filled stays in your vault.
[**CreateVaultMarketShareCard**](VaultMarketsApi.md#CreateVaultMarketShareCard) | **Post** /v1/vault-markets/share-cards | Make a verified share card of a position, a vault unit, a public index or your portfolio (&#x60;kind&#x60;, &#x60;refId&#x60;). The % on it is the price move since your real fill; amounts stay hidden unless &#x60;showAmounts&#x60;.
[**CreateVaultMarketShort**](VaultMarketsApi.md#CreateVaultMarketShort) | **Post** /v1/vault-markets/shorts | Short a book key: borrow &#x60;units&#x60; lent units (default 1; at most 25/10/3/1 by liquidity tier A/B/C/D) and sell each at &#x60;limitCents&#x60; or better, posting collateral for all of them (50% of the mark each; CBX counts at half, up to 30%). All units are borrowed or none unless &#x60;allowPartial&#x60;. A partial sale opens with the units sold. Returns &#x60;{ positionId, matched, units }&#x60;.
[**CreateVaultMarketSupplyBatchLabel**](VaultMarketsApi.md#CreateVaultMarketSupplyBatchLabel) | **Post** /v1/vault-markets/supply/batches/{id}/label | Buy a box&#39;s prepaid label again after it failed (for example once a ship-from address is saved). A box that has a label keeps it.
[**CreateVaultMarketSupplyInvite**](VaultMarketsApi.md#CreateVaultMarketSupplyInvite) | **Post** /v1/vault-markets/supply/invites/{id} | Accept or decline a Market Open invite (&#x60;accept&#x60;).
[**CreateVaultMarketSupplySend**](VaultMarketsApi.md#CreateVaultMarketSupplySend) | **Post** /v1/vault-markets/supply/send | Send inventory items matched to a catalogue variant to the vault as one box: each becomes an inbound vault unit, and the box gets one prepaid inbound label (a PDF to print). Optional: auto-list each unit in the vault book at a share of the mark once stored, and lend when idle. Items that cannot go say why. Crosslisted listings stay live.
[**CreateVaultMarketTerm**](VaultMarketsApi.md#CreateVaultMarketTerm) | **Post** /v1/vault-markets/terms | Accept the current lending terms (needed to lend or short). Returns &#x60;{ termsVersion }&#x60;.
[**CreateVaultMarketTodClaimObject**](VaultMarketsApi.md#CreateVaultMarketTodClaimObject) | **Post** /v1/vault-markets/tod/claims/{id}/object | Stop a transfer-on-death claim made on your own account. Returns &#x60;{ id, status, estateUnlocked }&#x60;.
[**CreateVaultMarketTodMyClaim**](VaultMarketsApi.md#CreateVaultMarketTodMyClaim) | **Post** /v1/vault-markets/tod/my-claims | Start a claim as a named beneficiary (ID first), naming the account holder by @username or email. It stays a draft until you add documents and submit. Returns &#x60;{ id, status, existing }&#x60;.
[**CreateVaultMarketTodMyClaimDocument**](VaultMarketsApi.md#CreateVaultMarketTodMyClaimDocument) | **Post** /v1/vault-markets/tod/my-claims/{id}/documents | Add one document to a draft claim (or one ops asked more of): kind, filename and the file as base64 (PDF, JPEG or PNG, up to 8 MB). Kept privately for ops review. Returns &#x60;{ id, kind, label, filename, bytes }&#x60;.
[**CreateVaultMarketTodMyClaimSubmit**](VaultMarketsApi.md#CreateVaultMarketTodMyClaimSubmit) | **Post** /v1/vault-markets/tod/my-claims/{id}/submit | Send a claim for ops review (it needs a death certificate). The account holder is told and can stop it. Returns &#x60;{ id, status }&#x60;.
[**CreateVaultMarketTodMyClaimWithdraw**](VaultMarketsApi.md#CreateVaultMarketTodMyClaimWithdraw) | **Post** /v1/vault-markets/tod/my-claims/{id}/withdraw | Withdraw a claim before it is released. Returns &#x60;{ id, status }&#x60;.
[**CreateVaultMarketTodTerm**](VaultMarketsApi.md#CreateVaultMarketTodTerm) | **Post** /v1/vault-markets/tod/terms | Accept the current transfer-on-death terms. Returns &#x60;{ termsVersion }&#x60;.
[**CreateVaultMarketUnitEvidenceScan**](VaultMarketsApi.md#CreateVaultMarketUnitEvidenceScan) | **Post** /v1/vault-markets/units/{id}/evidence/scans | Add a finished scan of your own (&#x60;captureId&#x60;, of the same catalogue item) to your vault unit. Shown as added by the owner, not as intake evidence.
[**CreateVaultMarketUnitSellNow**](VaultMarketsApi.md#CreateVaultMarketUnitSellNow) | **Post** /v1/vault-markets/units/{id}/sell-now | Sell one unit of this item now into the best bid (Instant Exit: an immediate-or-cancel pro order), never below &#x60;minCents&#x60;. Returns &#x60;{ status: filled | missed, filledCents, quote, order }&#x60;.
[**CreateVaultMarketValuationStatement**](VaultMarketsApi.md#CreateVaultMarketValuationStatement) | **Post** /v1/vault-markets/valuation-statements | Issue an insurance or appraisal statement for a day: frozen with the sha256 of its record, sealed in the Crossly audit chain, with a public verify link. preparedFor and policyRef are printed on it and never shown on verify. Moves no money.
[**CreateVaultMarketWatchlist**](VaultMarketsApi.md#CreateVaultMarketWatchlist) | **Post** /v1/vault-markets/watchlist | Watch a book key. Returns &#x60;{ bookKey, watching }&#x60;.
[**CreateVaultMarketWidgetToken**](VaultMarketsApi.md#CreateVaultMarketWidgetToken) | **Post** /v1/vault-markets/widget/tokens | Mint a widget token for one phone (&#x60;deviceName&#x60;, &#x60;installId&#x60;); minting again with the same pair replaces that phone token. The token reads the widget summary and nothing else, and is returned once. At most five phones.
[**CreateVaultReceiptTokenConfirm**](VaultMarketsApi.md#CreateVaultReceiptTokenConfirm) | **Post** /v1/vault/receipt-tokens/{id}/confirm | After your wallet sent the export: read the chain and mark the token live once its one token exists. Pass the &#x60;signature&#x60; your wallet returned to record it (kept only if it landed).
[**CreateVaultReceiptTokenExport**](VaultMarketsApi.md#CreateVaultReceiptTokenExport) | **Post** /v1/vault/receipt-tokens/export | Send one unit (&#x60;vaultItemId&#x60;) warehouse receipt to one of your verified Solana wallets (&#x60;wallet&#x60;) as a token, accepting the current terms (&#x60;acceptTerms&#x60;). Returns a base64 transaction already signed by the new mint and the receipts authority: your wallet pays the rent and must sign and send it within about two minutes, then call confirm. Nothing is minted until it does. Needs a verified ID; the unit must be stored, unpledged, not lent and not for sale.
[**CreateVaultReceiptTokenRedeem**](VaultMarketsApi.md#CreateVaultReceiptTokenRedeem) | **Post** /v1/vault/receipt-tokens/{id}/redeem | Redeem a receipt token held in one of your verified Solana wallets (&#x60;wallet&#x60;): returns an unsigned base64 burn your wallet signs and sends (it also closes the token account and returns its rent to you). Once it lands, call redeem/confirm and the unit is yours in the vault. Refused while the token is frozen.
[**CreateVaultReceiptTokenRedeemConfirm**](VaultMarketsApi.md#CreateVaultReceiptTokenRedeemConfirm) | **Post** /v1/vault/receipt-tokens/{id}/redeem/confirm | After your wallet sent the redeem: read the chain; once the token is burned the unit moves into your vault (if it was not yours already). Pass the &#x60;signature&#x60; to record it.
[**CreateVaultReceiptTokenWalletAllow**](VaultMarketsApi.md#CreateVaultReceiptTokenWalletAllow) | **Post** /v1/vault/receipt-tokens/wallets/allow | Put one of your verified Solana wallets (&#x60;wallet&#x60;) on the receipt-token allow-list so it can hold and receive receipt tokens. Returns a base64 transaction already signed by the receipts authority: the wallet pays the small rent (about 0.001 SOL) and must sign and send it within about two minutes. Needs a verified ID and a wallet that passes the sanctions screen.
[**DeleteVaultMarketAgentWatch**](VaultMarketsApi.md#DeleteVaultMarketAgentWatch) | **Delete** /v1/vault-markets/agent/watches/{id} | Cancel a live price watch. Returns the watch.
[**DeleteVaultMarketAlert**](VaultMarketsApi.md#DeleteVaultMarketAlert) | **Delete** /v1/vault-markets/alerts/{id} | Delete a price alert. Returns &#x60;{ ok }&#x60;.
[**DeleteVaultMarketBullionFamilyPeg**](VaultMarketsApi.md#DeleteVaultMarketBullionFamilyPeg) | **Delete** /v1/vault-markets/bullion/family-pegs/{id} | Stop a family-wide quote: every leg order comes off the book; what filled stays filled.
[**DeleteVaultMarketBullionPeg**](VaultMarketsApi.md#DeleteVaultMarketBullionPeg) | **Delete** /v1/vault-markets/bullion/pegs/{id} | Stop a spot-pegged quote: its resting order comes off the book; what filled stays filled.
[**DeleteVaultMarketDealerBand**](VaultMarketsApi.md#DeleteVaultMarketDealerBand) | **Delete** /v1/vault-markets/dealer/bands | Stop a band on a key; its standing bid comes off.
[**DeleteVaultMarketEntityMember**](VaultMarketsApi.md#DeleteVaultMarketEntityMember) | **Delete** /v1/vault-markets/entities/{id}/members/{memberId} | Admin: remove a member (or leave, yourself). An entity always keeps one admin.
[**DeleteVaultMarketFollowCopyWatch**](VaultMarketsApi.md#DeleteVaultMarketFollowCopyWatch) | **Delete** /v1/vault-markets/follows/copy-watch/{handle} | Stop copy-watching a collector.
[**DeleteVaultMarketIndex**](VaultMarketsApi.md#DeleteVaultMarketIndex) | **Delete** /v1/vault-markets/indexes/{slug} | Delete an index of your own. Returns &#x60;{ ok }&#x60;.
[**DeleteVaultMarketLocate**](VaultMarketsApi.md#DeleteVaultMarketLocate) | **Delete** /v1/vault-markets/locates/{id} | Release a live locate before it lapses.
[**DeleteVaultMarketOpeningWaitlist**](VaultMarketsApi.md#DeleteVaultMarketOpeningWaitlist) | **Delete** /v1/vault-markets/opening/{slug}/waitlist | Leave a key&#39;s waitlist.
[**DeleteVaultMarketPriceWidget**](VaultMarketsApi.md#DeleteVaultMarketPriceWidget) | **Delete** /v1/vault-markets/price-widgets/{id} | Delete a widget. Pages that still paste it show that it is no longer available. Returns &#x60;{ id, archived }&#x60;.
[**DeleteVaultMarketProQuote**](VaultMarketsApi.md#DeleteVaultMarketProQuote) | **Delete** /v1/vault-markets/pro/quotes | Pull your quote on a book key, or every quote. Returns &#x60;{ pulled }&#x60;.
[**DeleteVaultMarketProfileFollow**](VaultMarketsApi.md#DeleteVaultMarketProfileFollow) | **Delete** /v1/vault-markets/profiles/{handle}/follow | Stop following someone.
[**DeleteVaultMarketRecurring**](VaultMarketsApi.md#DeleteVaultMarketRecurring) | **Delete** /v1/vault-markets/recurring/{id} | Stop a recurring buy for good; its history stays.
[**DeleteVaultMarketSetGoal**](VaultMarketsApi.md#DeleteVaultMarketSetGoal) | **Delete** /v1/vault-markets/set-goals/{id} | Remove a set goal. A working basket bid is cancelled first; what filled stays in your vault.
[**DeleteVaultMarketShareCard**](VaultMarketsApi.md#DeleteVaultMarketShareCard) | **Delete** /v1/vault-markets/share-cards/{id} | Take a share card down: its link and image stop working. Returns &#x60;{ ok }&#x60;.
[**DeleteVaultMarketSupplyRoundUp**](VaultMarketsApi.md#DeleteVaultMarketSupplyRoundUp) | **Delete** /v1/vault-markets/supply/round-up | Stop the round-up (its history stays; an armed buy is cancelled).
[**DeleteVaultMarketUnitEvidence**](VaultMarketsApi.md#DeleteVaultMarketUnitEvidence) | **Delete** /v1/vault-markets/units/{id}/evidence/{evidenceId} | Take back a scan you added to your vault unit (intake evidence stays).
[**DeleteVaultMarketWatchlist**](VaultMarketsApi.md#DeleteVaultMarketWatchlist) | **Delete** /v1/vault-markets/watchlist | Stop watching a book key (&#x60;?bookKey&#x3D;&#x60;). Returns &#x60;{ bookKey, watching }&#x60;.
[**DeleteVaultMarketWidgetToken**](VaultMarketsApi.md#DeleteVaultMarketWidgetToken) | **Delete** /v1/vault-markets/widget/tokens | Turn the widget off on every phone: revokes all your widget tokens. Returns &#x60;{ revoked }&#x60;, how many.
[**DeleteVaultMarketWidgetTokenById**](VaultMarketsApi.md#DeleteVaultMarketWidgetTokenById) | **Delete** /v1/vault-markets/widget/tokens/{id} | Turn the widget off on one phone: revokes that widget token at once. Only a widget token can be revoked here.
[**GetVaultAttestation**](VaultMarketsApi.md#GetVaultAttestation) | **Get** /v1/vault/attestations | Vault attestations: independent physical counts, the insurance certificate (with a warning when insured value in the vault exceeds its aggregate or per-occurrence limit), SOC 1 Type II and SOC 2 reports. Each carries its signed document.
[**GetVaultMarketAccount**](VaultMarketsApi.md#GetVaultMarketAccount) | **Get** /v1/vault-markets/account | Which Vault Markets registrations are live (and the env each waits on), your custodian account, limits and balances.
[**GetVaultMarketAdl**](VaultMarketsApi.md#GetVaultMarketAdl) | **Get** /v1/vault-markets/adl | The insurance fund (booked and free to draw), the most one event may draw before auto-deleveraging (ADL) starts, the loss waterfall in order, and the latest ADL events (no user or position ids). Money in cents.
[**GetVaultMarketAdlPosition**](VaultMarketsApi.md#GetVaultMarketAdlPosition) | **Get** /v1/vault-markets/adl/positions/{id} | One of your positions with its ADL lights (0 once it has ended).
[**GetVaultMarketAdvanceQuote**](VaultMarketsApi.md#GetVaultMarketAdvanceQuote) | **Get** /v1/vault-markets/advances/quote | What an advance on a vault unit for sale could be: up to half of the lower of its mark and asking price, free for 60 days.
[**GetVaultMarketAgentDraft**](VaultMarketsApi.md#GetVaultMarketAgentDraft) | **Get** /v1/vault-markets/agent/drafts/{id} | One drafted order: item, side, units, limit (cents), the book when drafted, status, and the orderBody that places it.
[**GetVaultMarketAlertPref**](VaultMarketsApi.md#GetVaultMarketAlertPref) | **Get** /v1/vault-markets/alert-prefs | Your vault alert settings: every kind (price, fill, margin_call, adl, p2p_loan, auction, market_open, receipt_token, set_goal, follow, account), every channel (push, telegram, terminal, email, webhook) with whether it is set up, the full on/off grid, the locked switches (a critical kind always pushes) and quiet hours.
[**GetVaultMarketAuctionLot**](VaultMarketsApi.md#GetVaultMarketAuctionLot) | **Get** /v1/vault-markets/auctions/lots/{id} | One lot: the price showing, the next minimum bid, whether a reserve exists and is met, the end (extended by late bids), the bid history and your standing.
[**GetVaultMarketAuctionMine**](VaultMarketsApi.md#GetVaultMarketAuctionMine) | **Get** /v1/vault-markets/auctions/mine | Lots you consigned (with your reserve) and lots you bid on (with your maximum and whether you lead).
[**GetVaultMarketBlock**](VaultMarketsApi.md#GetVaultMarketBlock) | **Get** /v1/vault-markets/blocks/{orderId} | One block order: its order, plan lines (planned and allocated units, average price, cash share, status).
[**GetVaultMarketBullionFor**](VaultMarketsApi.md#GetVaultMarketBullionFor) | **Get** /v1/vault-markets/bullion/for | The bullion key families one book key belongs to, with spot for its metal as a labelled reference and its melt value at that spot. Empty families when the item is not keyed bullion.
[**GetVaultMarketBullionKey**](VaultMarketsApi.md#GetVaultMarketBullionKey) | **Get** /v1/vault-markets/bullion/keys/{slug} | One bullion key family: its catalogue products, every live book key with best bid, best ask, vault units and its mark (from the book only), live pegged quote counts, and spot as a labelled reference.
[**GetVaultMarketBullionPegPreview**](VaultMarketsApi.md#GetVaultMarketBullionPegPreview) | **Get** /v1/vault-markets/bullion/pegs/preview | The price a spot-pegged quote would rest at now (spot plus a premium in cents or basis points, clamped to floor and cap). No price when spot is stale or missing: such a quote waits off the book.
[**GetVaultMarketCall**](VaultMarketsApi.md#GetVaultMarketCall) | **Get** /v1/vault-markets/calls/{bookKey} | Recent call-auction uncrosses on a book key: kind (open, reopen, periodic, daily, buy-in, sell-off), clearing price in cents and units.
[**GetVaultMarketCapacity**](VaultMarketsApi.md#GetVaultMarketCapacity) | **Get** /v1/vault-markets/capacity | Capacity from the book: &#x60;bookKey&#x60; (+ &#x60;units&#x60;) gives ADV, depth within ±5% of the mark, vault stock, lendable units and days to liquidate at 20% of ADV; &#x60;index&#x60; + &#x60;notional&#x60; (cents) gives how much of the basket the asks fill now and the days the rest needs.
[**GetVaultMarketChannel**](VaultMarketsApi.md#GetVaultMarketChannel) | **Get** /v1/vault-markets/channels | Your Vault Markets channel switches: Telegram alerts and digest, the daily cap on vault nudges by push (default 5), whether a Telegram chat is linked, and how many nudges went out today.
[**GetVaultMarketCompare**](VaultMarketsApi.md#GetVaultMarketCompare) | **Get** /v1/vault-markets/compare | Up to 5 book keys and/or index slugs (&#x60;items&#x60;, comma-separated) as daily series on the same days, each rebased to 100.
[**GetVaultMarketConfirmation**](VaultMarketsApi.md#GetVaultMarketConfirmation) | **Get** /v1/vault-markets/confirmations/{tradeId} | One trade confirmation.
[**GetVaultMarketConfirmationPdf**](VaultMarketsApi.md#GetVaultMarketConfirmationPdf) | **Get** /v1/vault-markets/confirmations/{tradeId}/pdf | One trade confirmation as a PDF.
[**GetVaultMarketCorrelation**](VaultMarketsApi.md#GetVaultMarketCorrelation) | **Get** /v1/vault-markets/correlations | The correlation matrix of the sector indexes&#39; weekly returns over &#x60;weeks&#x60; (default 52), pairwise-complete, with n per pair; a pair under 26 shared weeks has no number.
[**GetVaultMarketCreatorAttribution**](VaultMarketsApi.md#GetVaultMarketCreatorAttribution) | **Get** /v1/vault-markets/creators/attribution | The creator who brought you to Vault Markets, when one did: when it started and ends, whether it still earns, and the disclosure. Your prices and fees are the same either way.
[**GetVaultMarketCreatorDashboard**](VaultMarketsApi.md#GetVaultMarketCreatorDashboard) | **Get** /v1/vault-markets/creators/dashboard | For creators: whether you joined the programme, each of your codes with traders attributed and still earning (counts only), shares accrued, paid and voided (cents) and the latest shares.
[**GetVaultMarketCreatorProgramme**](VaultMarketsApi.md#GetVaultMarketCreatorProgramme) | **Get** /v1/vault-markets/creators/programme | The Vault Markets creator programme: the share of the fee Crossly keeps that a creator earns, the per-fill and lifetime caps (cents), the attribution window in days, the 2% net floor it sits under, and the disclosure words.
[**GetVaultMarketCreatorRoom**](VaultMarketsApi.md#GetVaultMarketCreatorRoom) | **Get** /v1/vault-markets/creators/room | For creators: your vault shelf room (applied, verified, rejected or paused) and whether your public vault profile is on, which the room needs.
[**GetVaultMarketCredit**](VaultMarketsApi.md#GetVaultMarketCredit) | **Get** /v1/vault-markets/credit/{id} | One loan or advance with its pledged units and its money log (draws, interest, repayments, sale proceeds).
[**GetVaultMarketDcm**](VaultMarketsApi.md#GetVaultMarketDcm) | **Get** /v1/vault-markets/dcm | The partner exchange (DCM) that lists and clears futures and options on Crossly indexes (needs VAULT_INDEX_DCM + VAULT_INDEX_DCM_API_URL + VAULT_INDEX_DCM_API_KEY), and its contracts as last read: perps, futures and options with the index each settles on, multiplier, tick, expiry and strike. Read-only: Crossly is the licensor and data provider, never the venue.
[**GetVaultMarketDcmContract**](VaultMarketsApi.md#GetVaultMarketDcmContract) | **Get** /v1/vault-markets/dcm/contracts/{symbol} | One partner contract: its spec, its funding prints (newest first; the partner computes and settles funding), annualised funding over the last 21 prints, and the Crossly index level it settles on.
[**GetVaultMarketDcmContractBook**](VaultMarketsApi.md#GetVaultMarketDcmContractBook) | **Get** /v1/vault-markets/dcm/contracts/{symbol}/book | The partner exchange live order book for a contract (cached 5 seconds), with mark, index, last, open interest, 24 h volume, spread in bps and the basis of its mark over the Crossly index in bps. Prices are index points.
[**GetVaultMarketDcmOptionsInput**](VaultMarketsApi.md#GetVaultMarketDcmOptionsInput) | **Get** /v1/vault-markets/dcm/options-inputs/{slug} | Options inputs for an index (indexes only, never a single item): Crossly realised volatility (raw and Geltner-desmoothed, from weekly returns), the forward (the partner future mark, else the Crossly close), Black-76 greeks on a strike x expiry grid and on every option the partner lists on the index. Model inputs, not quotes.
[**GetVaultMarketDcmPosition**](VaultMarketsApi.md#GetVaultMarketDcmPosition) | **Get** /v1/vault-markets/dcm/positions | The partner exchange account linked to your account (or to the firm sub-account &#x60;subaccountId&#x60;) and its positions, read-only (cached 15 seconds); index futures on category indexes count as portfolio-margin hedges for that account: contracts, side, entry and mark price, unrealised P&amp;L and margin in cents, liquidation price and notional in cents. Not linked: &#x60;linked: false&#x60;.
[**GetVaultMarketDealer**](VaultMarketsApi.md#GetVaultMarketDealer) | **Get** /v1/vault-markets/dealer | Your dealer status (applied, approved, revoked) and your Instant Exit bands.
[**GetVaultMarketDealerPool**](VaultMarketsApi.md#GetVaultMarketDealerPool) | **Get** /v1/vault-markets/dealer/pools/{id} | One of your curve pools, with its ladder and live legs.
[**GetVaultMarketDigestPreview**](VaultMarketsApi.md#GetVaultMarketDigestPreview) | **Get** /v1/vault-markets/digest/preview | This week&#39;s vault digest as it would go out: units at the display marks (unmarked units counted, not valued), the like-for-like change from a week ago, offers on your units, watched keys that got a first mark, and one Learn card.
[**GetVaultMarketEntity**](VaultMarketsApi.md#GetVaultMarketEntity) | **Get** /v1/vault-markets/entities | The entities (funds, LLCs, trusts, SPVs) you belong to, and invitations waiting for you.
[**GetVaultMarketEntityById**](VaultMarketsApi.md#GetVaultMarketEntityById) | **Get** /v1/vault-markets/entities/{id} | One entity: members and roles, sub-accounts with their limits, firm halts, caps.
[**GetVaultMarketEstimate**](VaultMarketsApi.md#GetVaultMarketEstimate) | **Get** /v1/vault-markets/estimates/{bookKey} | The Crossly Estimate for a book key — a model, not a price. Method (own trades, index-adjusted last trade, grade ratio), cents with 68% and 95% bands, High/Medium/Low from the walk-forward backtest&#39;s FSD, and n. Hidden (&#x60;available: false&#x60;) under 50 backtested trades or when FSD &gt; 0.30. Never used for margin, loans or settlement.
[**GetVaultMarketExecutionQuality**](VaultMarketsApi.md#GetVaultMarketExecutionQuality) | **Get** /v1/vault-markets/execution-quality | The quarterly execution-quality report (Rule 605 style): fill rates, speed, effective vs quoted spread and shortfall by order type and size, and the busiest keys. &#x60;quarter&#x60; like 2026-Q4; the current quarter is provisional.
[**GetVaultMarketFeeQuote**](VaultMarketsApi.md#GetVaultMarketFeeQuote) | **Get** /v1/vault-markets/fees/quote | Your fee for an order at a price: as maker (the order rests) and as taker (it fills at once), with the programme that sets your rate and any Crossly One trim. A buyer&#39;s fee is added to the price; a seller&#39;s comes out of the proceeds.
[**GetVaultMarketFeeSchedule**](VaultMarketsApi.md#GetVaultMarketFeeSchedule) | **Get** /v1/vault-markets/fees/schedule | The fee schedule by role: maker and taker rates for retail, 30-day volume, Registered LPs and DMMs, the CBX discount, the referral share and the 2% net floor Crossly always keeps.
[**GetVaultMarketFollowSharing**](VaultMarketsApi.md#GetVaultMarketFollowSharing) | **Get** /v1/vault-markets/follows/sharing | Your sharing switches: the public vault profile, new buys, units sent in, sales, Hall loans, bids and asks for copy-watchers, and the leaderboards. All off until you turn them on.
[**GetVaultMarketGemRate**](VaultMarketsApi.md#GetVaultMarketGemRate) | **Get** /v1/vault-markets/gem-rate/{bookKey} | The Crossly gem rate: of cards graded through Crossly, the share that came back at each grader&#39;s gem grade, for this item and its category, with a Wilson 95% band. 10 returns per grader or no number. Not a grader population report.
[**GetVaultMarketGradeToVault**](VaultMarketsApi.md#GetVaultMarketGradeToVault) | **Get** /v1/vault-markets/grade-to-vault/{id} | One Grade to Vault flow, brought up to date with its submission and unit.
[**GetVaultMarketGradeToVaultCandidate**](VaultMarketsApi.md#GetVaultMarketGradeToVaultCandidate) | **Get** /v1/vault-markets/grade-to-vault/candidates | What you can start Grade to Vault on: raw inventory matched to a catalogue item, and raw vault units on the grading hub shelf; plus whether Grading Runs are open.
[**GetVaultMarketGradingSignal**](VaultMarketsApi.md#GetVaultMarketGradingSignal) | **Get** /v1/vault-markets/grading-signal | For a raw book key: the marks of the same item&#39;s graded keys, the spread over raw, grading cost per grader, the margin tier grading lifts it to, your raw vault units it pays to grade, and where to act.
[**GetVaultMarketHallLoan**](VaultMarketsApi.md#GetVaultMarketHallLoan) | **Get** /v1/vault-markets/hall/loans | Your Hall of Grails desk: this week&#39;s exhibition, the rules (maxLive loans, shownPerHolder per exhibition, repPoints per exhibition shown), reputation (exhibitions finished), your loans (status active, ending or ended; onShowNow, showingUntil, exhibitions, the next exhibition it fits, the Lent to the Hall badge) and the vault units you could lend (canLend, or the blocker in words, and the next exhibition each fits). Display only.
[**GetVaultMarketHealth**](VaultMarketsApi.md#GetVaultMarketHealth) | **Get** /v1/vault-markets/health | Portfolio health of your vault holdings at Crossly marks: concentration by single item, category, grade and grader (shares, HHI on 0 to 10000, its band and effective number of holdings), value by liquidity tier, the stress grid (each category and everything down 10, 20 and 30 percent), value by mark quality, plain observations, and the held items a what-if can run on. Information only, not advice.
[**GetVaultMarketHealthWhatIf**](VaultMarketsApi.md#GetVaultMarketHealthWhatIf) | **Get** /v1/vault-markets/health/what-if | What if one held item moved by moveBps (-10000 to 10000 basis points of its mark): the item, your vault total, the item and its category share of the vault, and the single-item HHI, before and after. A scenario on the mark, not a price you could sell at.
[**GetVaultMarketHomeUnit**](VaultMarketsApi.md#GetVaultMarketHomeUnit) | **Get** /v1/vault-markets/home/units | \&quot;Show my vault\&quot;: the vault units you hold, the ones standing in your Crossly World home first. Each: vaultItemId, title, grade, imageUrl, bookKey (its market), inHome, visitors (visitors to your home see it). Also your home (slug, privacy, path), max units one home shows, shown, visitorsSee and a note on who sees what. Your home is made on first read. No prices.
[**GetVaultMarketIndex**](VaultMarketsApi.md#GetVaultMarketIndex) | **Get** /v1/vault-markets/indexes/{slug} | One index with its daily closes (oldest first) and its level now.
[**GetVaultMarketIndexCalendar**](VaultMarketsApi.md#GetVaultMarketIndexCalendar) | **Get** /v1/vault-markets/index-calendar | The next four quarterly reviews: reference day (2nd Friday of Mar/Jun/Sep/Dec), announcement day, last close under the old membership and the effective day; June and December also review the Blue-chip 50.
[**GetVaultMarketIndexConsultation**](VaultMarketsApi.md#GetVaultMarketIndexConsultation) | **Get** /v1/vault-markets/index-consultations/{id} | One consultation with its proposal, versions, dates, outcome and the public comments.
[**GetVaultMarketIndexFeed**](VaultMarketsApi.md#GetVaultMarketIndexFeed) | **Get** /v1/vault-markets/index-feed/{slug} | The licensed index feed for one built-in index (needs VAULT_INDEX_DCM + VAULT_INDEX_FEED_SECRET): the level now (intraday from A/B marks, or the last close below 50% constituent coverage), the last close, the frozen 7-day TWAP settlement value and every constituent with its mark source, quality and trade count. Signed: &#x60;x-crossly-signature: t&#x3D;&lt;ms&gt;,v1&#x3D;&lt;hex HMAC-SHA256(secret, t + \&quot;.\&quot; + body)&gt;&#x60;. Every read is logged for the licence. Crossly is the licensor and data feed only.
[**GetVaultMarketIndexFeedFunding**](VaultMarketsApi.md#GetVaultMarketIndexFeedFunding) | **Get** /v1/vault-markets/index-feed/{slug}/funding | Reference funding for a perpetual-style future on the index: your premium &#x60;premiumBps&#x60; (TWAP of (impact mid - index)/index) through F &#x3D; P + clamp(I - P, ±0.05%) per 8 h, |F| capped at 0.5%/day, with the index level it is relative to. The partner DCM computes and settles funding; this is the reference math.
[**GetVaultMarketIndexFeedSettlement**](VaultMarketsApi.md#GetVaultMarketIndexFeedSettlement) | **Get** /v1/vault-markets/index-feed/{slug}/settlement | The frozen 7-day TWAP settlement series for a licensed index (written once a day, never restated), with the closes behind each value. Signed like the feed.
[**GetVaultMarketIndexLevel**](VaultMarketsApi.md#GetVaultMarketIndexLevel) | **Get** /v1/vault-markets/indexes/{slug}/levels | An index&#39;s daily closes (up to 5 years), oldest first.
[**GetVaultMarketIndexLicenceUsage**](VaultMarketsApi.md#GetVaultMarketIndexLicenceUsage) | **Get** /v1/vault-markets/index-licence/usage | Your licensed index feed reads over &#x60;days&#x60;, per index and endpoint (licence-usage report).
[**GetVaultMarketIndexMethodology**](VaultMarketsApi.md#GetVaultMarketIndexMethodology) | **Get** /v1/vault-markets/index-methodology | The index method in machine-readable form: version (semver), inputs, weightings and caps, the family and its screens, the engines and their publication gate, revisions and the settlement series, the calendar, governance and limitations.
[**GetVaultMarketInstantExit**](VaultMarketsApi.md#GetVaultMarketInstantExit) | **Get** /v1/vault-markets/instant-exit | The Instant Exit quote for one vault unit you hold: the best resting bid on its book key (anyone&#39;s, dealers&#39; included), its share of the mark, the bid depth, and whether it can be taken now.
[**GetVaultMarketJoint**](VaultMarketsApi.md#GetVaultMarketJoint) | **Get** /v1/vault-markets/joint | Your joint vaults and the invitations waiting for you, with the gate (open, missing env), the joint terms in words and what each right allows. A JointVaultsHome object, not a collection.
[**GetVaultMarketJointById**](VaultMarketsApi.md#GetVaultMarketJointById) | **Get** /v1/vault-markets/joint/{id} | One joint vault (active members only): members and rights, the units in it at the mark with each owner of record, totals, recent joint trades, and the activity from the audit trail.
[**GetVaultMarketKeyStat**](VaultMarketsApi.md#GetVaultMarketKeyStat) | **Get** /v1/vault-markets/key-stats/{bookKey} | Everything derived about one book key in one read: the Crossly Estimate, the liquidity score, risk statistics and the gem rate. A part that cannot be read is &#x60;{ error }&#x60;.
[**GetVaultMarketLearn**](VaultMarketsApi.md#GetVaultMarketLearn) | **Get** /v1/vault-markets/learn | The Learn hub: each concept with whether you passed its quiz. Passing earns cosmetic pins only.
[**GetVaultMarketLearnByConcept**](VaultMarketsApi.md#GetVaultMarketLearnByConcept) | **Get** /v1/vault-markets/learn/{concept} | One Learn concept: its sections and quiz questions (no answers).
[**GetVaultMarketLendingElection**](VaultMarketsApi.md#GetVaultMarketLendingElection) | **Get** /v1/vault-markets/lending/election | Your lender election: &#x60;returnInKind&#x60; (shorts post 175% collateral, a 14-day buy-in, never cash-settled) and &#x60;allowTerm&#x60; (your units may go out on 30/90-day fixed-rate term loans, not recallable before term).
[**GetVaultMarketLendingProgramme**](VaultMarketsApi.md#GetVaultMarketLendingProgramme) | **Get** /v1/vault-markets/lending/programme | Your lender programme report: lendable units, units out, utilisation, fee revenue (accrued / paid), term loans and return-in-kind loans in force, and each loan&#39;s rate against the key&#39;s rate today.
[**GetVaultMarketLendingStatement**](VaultMarketsApi.md#GetVaultMarketLendingStatement) | **Get** /v1/vault-markets/lending/statement | Your daily collateral statement as a lender (today live, or &#x60;day&#x60;): per open loan the mark and its quality, fees owed to you, the short&#39;s collateral share at the custodian and its coverage against 105% of mark + fees. A unit without a mark usable for risk shows no coverage number.
[**GetVaultMarketLiquidity**](VaultMarketsApi.md#GetVaultMarketLiquidity) | **Get** /v1/vault-markets/liquidity/{bookKey} | The liquidity score (0-100) of a book key and every component behind it: trades in 30/90 days, distinct counterparties, median days between trades, time-weighted quoted spread, two-sided share and depth within 5% (hourly book snapshots), Amihud and Roll. Each component is a percentile within the category. No score under the data threshold.
[**GetVaultMarketLoanQuote**](VaultMarketsApi.md#GetVaultMarketLoanQuote) | **Get** /v1/vault-markets/loans/quote | How much you can borrow against vault units (&#x60;vaultItemIds&#x60;, comma-separated; none &#x3D; every stored unit you hold): each unit&#39;s tier, mark and limit, the pool&#39;s rate and yours.
[**GetVaultMarketLoanUnitSaleFloor**](VaultMarketsApi.md#GetVaultMarketLoanUnitSaleFloor) | **Get** /v1/vault-markets/loans/{id}/units/{vaultItemId}/sale-floor | The lowest price a pledged unit can be sold at to repay its loan: after the sale (less the vault fee and storage) what is still owed must fit inside the other units&#39; limit.
[**GetVaultMarketLocate**](VaultMarketsApi.md#GetVaultMarketLocate) | **Get** /v1/vault-markets/locates | A locate quote: how many units of &#x60;key&#x60; (a book key) you could borrow now (up to &#x60;qty&#x60;), the indicative rate (APR bps; &#x60;termDays&#x60; 30 or 90 for a fixed-rate term loan), and the collateral they need (150%, or 175% for units whose lenders elected return in kind only). Holds nothing; POST to hold.
[**GetVaultMarketMark**](VaultMarketsApi.md#GetVaultMarketMark) | **Get** /v1/vault-markets/marks/{bookKey} | One book key mark with everything behind it: the price in cents, its quality (A trade-anchored, B firm two-sided, C indicative, none), its confidence in bps and the band in cents, when it was computed, whether risk may act on it (judged for margin), the related key it was implied from (impliedFrom: a grade, the model median of the other sizes, or an index; null when none), every stored input, and the learned relations it may be implied from.
[**GetVaultMarketMarket**](VaultMarketsApi.md#GetVaultMarketMarket) | **Get** /v1/vault-markets/market | One book key (variantId:condition:gradeKey): its mark and history, lending depth and borrow rate, open interest and live auctions.
[**GetVaultMarketMarketModel**](VaultMarketsApi.md#GetVaultMarketMarketModel) | **Get** /v1/vault-markets/market-model/{bookKey} | How a book key trades: its liquidity tier (A continuous, B call auction every 15 minutes, C daily call, D listings only), the call collecting now with its indicative price and volume and nominal close, the next call, the LULD band and any limit state, pause or halt.
[**GetVaultMarketMm**](VaultMarketsApi.md#GetVaultMarketMm) | **Get** /v1/vault-markets/mm | Your market-maker programmes, the last 30 days of obligations, rebates earned (paid only from fees already taken), and the rules.
[**GetVaultMarketMmKey**](VaultMarketsApi.md#GetVaultMarketMmKey) | **Get** /v1/vault-markets/mm/key | A key&#39;s designated market maker (by the name it chose) and how many LPs quote it.
[**GetVaultMarketMmMmp**](VaultMarketsApi.md#GetVaultMarketMmMmp) | **Get** /v1/vault-markets/mm/mmp | Your market-maker protection: the fill trigger, the freeze, and whether you are frozen now.
[**GetVaultMarketMode**](VaultMarketsApi.md#GetVaultMarketMode) | **Get** /v1/vault-markets/mode | Your Vault Markets mode: simple (buy, sell and what it is worth) or pro (the full terminal), whether you chose it or it was set for you (new people start in simple; people already trading start in pro), and when.
[**GetVaultMarketNav**](VaultMarketsApi.md#GetVaultMarketNav) | **Get** /v1/vault-markets/nav | Mark NAV, bid NAV (the resting bids) and liquidation NAV (the bids at 20% of ADV for 10 days), holding by holding, each with its ASC 820 level, mark quality and source, and your administrator&#39;s price when uploaded.
[**GetVaultMarketNavFile**](VaultMarketsApi.md#GetVaultMarketNavFile) | **Get** /v1/vault-markets/nav/file/{format} | The NAV as a csv, pdf or json file.
[**GetVaultMarketOpening**](VaultMarketsApi.md#GetVaultMarketOpening) | **Get** /v1/vault-markets/opening/{slug} | One opening key.
[**GetVaultMarketOracle**](VaultMarketsApi.md#GetVaultMarketOracle) | **Get** /v1/vault-markets/oracle | Crossly as an oracle publisher (needs CROSSLY_ORACLE_PUBLISHER_KEY + CROSSLY_ORACLE_NETWORK): the network (pyth or switchboard), the ed25519 key every value is signed with, each published index feed with its latest attempt, and 24 h counts of published, skipped and failed. Only built-in indexes priced from A/B marks are published; Crossly never reads its own feed back.
[**GetVaultMarketOracleBySlug**](VaultMarketsApi.md#GetVaultMarketOracleBySlug) | **Get** /v1/vault-markets/oracle/{slug} | The latest published oracle value for an index: level, fixed-point price and confidence at the feed exponent, confidence in bps, staleness in seconds, coverage, the exact signed message and its ed25519 signature with the signer key, and whether it is fresh.
[**GetVaultMarketOrderHandling**](VaultMarketsApi.md#GetVaultMarketOrderHandling) | **Get** /v1/vault-markets/order-handling | How orders are handled (the 606/607-style disclosure): one book, self-trade prevention, protected market orders, liquidations by auction, halts, dealers and fees.
[**GetVaultMarketP2pDashboard**](VaultMarketsApi.md#GetVaultMarketP2pDashboard) | **Get** /v1/vault-markets/p2p/dashboard | Peer-to-peer vault loans at a glance: what you owe as a borrower and are owed as a lender (principal + interest to the second), daily interest, called loans, your open offers, the gates and the rules (auction hours, rate ceiling, size bounds, call grace).
[**GetVaultMarketP2pLoan**](VaultMarketsApi.md#GetVaultMarketP2pLoan) | **Get** /v1/vault-markets/p2p/loans/{id} | One P2P loan with its money log (fund, repay, refinance). A loan in a running refinancing auction can be read by any lender, without the log.
[**GetVaultMarketPaper**](VaultMarketsApi.md#GetVaultMarketPaper) | **Get** /v1/vault-markets/paper | Your paper trading account (play money only, never real or convertible; opened at $10,000 the first time): play cash, buying power, paper positions at the real marks, P&amp;L, working paper orders, recent paper fills and this week on the leaderboard.
[**GetVaultMarketPaperQuote**](VaultMarketsApi.md#GetVaultMarketPaperQuote) | **Get** /v1/vault-markets/paper/quote | What a paper order would fill against the real book now, without placing it: units fillable inside the limit (a market order uses the 5% protection band around the mark), average price, simulated fee, and the queue ahead of any rest.
[**GetVaultMarketPartnerMe**](VaultMarketsApi.md#GetVaultMarketPartnerMe) | **Get** /v1/vault-markets/partners/me | Your Vault Drop-off Partnership (null when you never applied): status, partner code, fee share and cap, totals and batches.
[**GetVaultMarketPool**](VaultMarketsApi.md#GetVaultMarketPool) | **Get** /v1/vault-markets/pool | The CBX market staking pool: size, value, what it lends, utilisation, rate, insurance fund.
[**GetVaultMarketPopReport**](VaultMarketsApi.md#GetVaultMarketPopReport) | **Get** /v1/vault-markets/pop-reports | Grading population report for one market (&#x60;bookKey&#x60;): focus (on a graded market: its grade, pop &#x3D; copies graded at exactly that grade, higher &#x3D; copies graded better by the same grader, total, asOf and source), graders (each grader ladder for the card, best grade first, from one report: the newest, the population feed over a hand entry on the same date), provider.connected, and status (ok, not_synced, no_report or no_source) with explain, which says in plain words what is missing. source is provider (the population feed) or ops (entered by Crossly from the grader published report). Counts only, never a price.
[**GetVaultMarketPortfolio**](VaultMarketsApi.md#GetVaultMarketPortfolio) | **Get** /v1/vault-markets/portfolio | Your vault holdings at the market price, with cost and P&amp;L, and units you have out on loan.
[**GetVaultMarketPortfolioHistory**](VaultMarketsApi.md#GetVaultMarketPortfolioHistory) | **Get** /v1/vault-markets/portfolio/history | Your Vault Markets value by day at the marks (vault units, units on loan, positions) and today&#39;s return against yesterday&#39;s close.
[**GetVaultMarketPortfolioMargin**](VaultMarketsApi.md#GetVaultMarketPortfolioMargin) | **Get** /v1/vault-markets/portfolio-margin | Portfolio margin for a firm sub-account (&#x60;subaccountId&#x60;): eligibility with every reason (ECP, KYB, minimum equity $500k or $100k with evidenced intraday monitoring), whether it is enrolled, the latest requirement, equity, excess and state (healthy, call or buy_in), the breakdown by category and item, the partner index futures counted, and the policy. Before enrolling, a live preview.
[**GetVaultMarketPosition**](VaultMarketsApi.md#GetVaultMarketPosition) | **Get** /v1/vault-markets/positions/{id} | One position with its loans, auctions and custodian instruction log.
[**GetVaultMarketPositionExit**](VaultMarketsApi.md#GetVaultMarketPositionExit) | **Get** /v1/vault-markets/positions/{id}/exits | A position&#39;s take-profit / stop-loss, if set.
[**GetVaultMarketPositionRisk**](VaultMarketsApi.md#GetVaultMarketPositionRisk) | **Get** /v1/vault-markets/positions/{id}/risk | One position&#39;s call and liquidation marks.
[**GetVaultMarketPress**](VaultMarketsApi.md#GetVaultMarketPress) | **Get** /v1/vault-markets/press | The press kit: each built-in index with level, quality, n, CSV link and a \&quot;cite as\&quot; line, plus the methodology links.
[**GetVaultMarketPriceWidget**](VaultMarketsApi.md#GetVaultMarketPriceWidget) | **Get** /v1/vault-markets/price-widgets/{id} | One of your price widgets, with its iframe and script-tag snippets and URLs.
[**GetVaultMarketProDepth**](VaultMarketsApi.md#GetVaultMarketProDepth) | **Get** /v1/vault-markets/pro/depth | The L2 depth ladder: price levels and units per side, best first; asks show how many units are authenticated vault stock.
[**GetVaultMarketProOrder**](VaultMarketsApi.md#GetVaultMarketProOrder) | **Get** /v1/vault-markets/pro/orders/{id} | One order with its full event log.
[**GetVaultMarketProSnapshot**](VaultMarketsApi.md#GetVaultMarketProSnapshot) | **Get** /v1/vault-markets/pro/snapshot | A book key&#39;s depth (20 levels a side), last 50 trades, 24 h stats, mark and any halt, in one call.
[**GetVaultMarketProStat**](VaultMarketsApi.md#GetVaultMarketProStat) | **Get** /v1/vault-markets/pro/stats | 24 h open, high, low, last, volume and change on a book key.
[**GetVaultMarketProStream**](VaultMarketsApi.md#GetVaultMarketProStream) | **Get** /v1/vault-markets/pro/stream | A long-lived text/event-stream for a book key: &#x60;snapshot&#x60; first, then &#x60;depth&#x60;, &#x60;trades&#x60;, &#x60;mark&#x60; and &#x60;halt&#x60; events as they change.
[**GetVaultMarketProfile**](VaultMarketsApi.md#GetVaultMarketProfile) | **Get** /v1/vault-markets/profile | Your public vault profile switch: on or off, whether followers see new buys, your profile URL and follower count.
[**GetVaultMarketProfileByHandle**](VaultMarketsApi.md#GetVaultMarketProfileByHandle) | **Get** /v1/vault-markets/profiles/{handle} | Someone&#39;s public vault: each market&#39;s weight (bps of the marked value), followers, whether you follow, their recent buys (no prices).
[**GetVaultMarketProofOfReserve**](VaultMarketsApi.md#GetVaultMarketProofOfReserve) | **Get** /v1/vault-markets/proof-of-reserves | Proof of reserves: the latest daily Merkle root over every unit in custody (leaf count, stored / listed / grading / withdrawing, pledged, on loan, open loans, receipts, insured value), its Solana memo anchor when VAULT_POR_ANCHOR_KEY is set, the last &#x60;days&#x60; roots, and the exact hashing method.
[**GetVaultMarketProofOfReserveByDay**](VaultMarketsApi.md#GetVaultMarketProofOfReserveByDay) | **Get** /v1/vault-markets/proof-of-reserves/{day} | One day&#39;s proof-of-reserves snapshot (root, counts, anchor).
[**GetVaultMarketProofOfReserveMe**](VaultMarketsApi.md#GetVaultMarketProofOfReserveMe) | **Get** /v1/vault-markets/proof-of-reserves/me | Your inclusion proofs for a day (latest by default): each of your units&#39; leaf fields, salt, owner commitment, leaf hash and sibling path. Verify with your OWN user id: commitment &#x3D; SHA-256(JSON[\&quot;crossly-por-owner-v1\&quot;, userId, salt]); then the leaf; then climb the path to the published root.
[**GetVaultMarketRate**](VaultMarketsApi.md#GetVaultMarketRate) | **Get** /v1/vault-markets/rates | Daily borrow APR (lent units) and pool APR on a book key, with the pool as a whole, oldest first.
[**GetVaultMarketReport**](VaultMarketsApi.md#GetVaultMarketReport) | **Get** /v1/vault-markets/report | The quarterly Crossly Vault Report (&#x60;quarter&#x60; like 2026-q3; default the latest): index levels, top movers, most-lent markets, grade premiums. Book data only.
[**GetVaultMarketRfq**](VaultMarketsApi.md#GetVaultMarketRfq) | **Get** /v1/vault-markets/rfq/{id} | One quote request: the requester sees every answer (anonymous); a responder sees the request and their own answer.
[**GetVaultMarketRiskReport**](VaultMarketsApi.md#GetVaultMarketRiskReport) | **Get** /v1/vault-markets/risk-report | Historical VaR 95/99 1d/10d and Bangia LVaR (model figures: labelled, with n and a band, null below the data threshold), stress scenarios (category −30%, CBX −50%, books thin to 25%, index drawdowns), concentration and counterparty exposure.
[**GetVaultMarketRiskStat**](VaultMarketsApi.md#GetVaultMarketRiskStat) | **Get** /v1/vault-markets/risk-stats/{bookKey} | Risk statistics for a book key (weekly returns from A-quality marks only) or an index slug (its closes): annualised volatility raw and Geltner-desmoothed (labelled), max drawdown and beta to the Composite. Needs 26 weekly returns, else no numbers.
[**GetVaultMarketSetGoal**](VaultMarketsApi.md#GetVaultMarketSetGoal) | **Get** /v1/vault-markets/set-goals/{id} | One set goal: every piece in the set order, held (with the vault unit) or missing (with the cheapest marked book key that meets the target and its mark), the conditions and grades seen, and the working basket bid with its legs.
[**GetVaultMarketSetGoalBasketPlan**](VaultMarketsApi.md#GetVaultMarketSetGoalBasketPlan) | **Get** /v1/vault-markets/set-goals/{id}/basket-plan | What a basket bid would place now, without placing it: one leg per missing piece (cheapest first, the mark plus maxBpsOverMark as the limit), the pieces skipped and why, and the most it can cost (limits plus the worst-case buyer fee), never more than budgetCents.
[**GetVaultMarketShareCard**](VaultMarketsApi.md#GetVaultMarketShareCard) | **Get** /v1/vault-markets/share-cards/{id} | One share card as the public sees it, with its image and link paths.
[**GetVaultMarketStatement**](VaultMarketsApi.md#GetVaultMarketStatement) | **Get** /v1/vault-markets/statements | The last 24 months with whether each statement is issued (frozen, with its sha256).
[**GetVaultMarketStatementByMonth**](VaultMarketsApi.md#GetVaultMarketStatementByMonth) | **Get** /v1/vault-markets/statements/{month} | A monthly statement (&#x60;month&#x60; like 2026-10): units held at month end, loans, pledges, positions, fills, fees, storage, and cash quoted from the custodian and labelled as its record. Issued (frozen) once the month ends.
[**GetVaultMarketStatementFile**](VaultMarketsApi.md#GetVaultMarketStatementFile) | **Get** /v1/vault-markets/statements/{month}/file/{format} | A monthly statement as a csv, pdf or json file.
[**GetVaultMarketSupply**](VaultMarketsApi.md#GetVaultMarketSupply) | **Get** /v1/vault-markets/supply | A book key&#39;s units in the vault and authentications passed, by day (oldest first, today live).
[**GetVaultMarketSupplyBatch**](VaultMarketsApi.md#GetVaultMarketSupplyBatch) | **Get** /v1/vault-markets/supply/batches/{id} | A box sent to the vault: its prepaid label and tracking, or why it has no label.
[**GetVaultMarketSupplyRoundUp**](VaultMarketsApi.md#GetVaultMarketSupplyRoundUp) | **Get** /v1/vault-markets/supply/round-up | Your seller round-up: the share of each payout that sets a vault buy&#39;s budget, the budget so far and what it is doing. &#x60;roundUp&#x60; is null when none is set.
[**GetVaultMarketTax**](VaultMarketsApi.md#GetVaultMarketTax) | **Get** /v1/vault-markets/tax | Your tax year on the vault markets: units bought and sold with cost, proceeds, gain and term; lending fees, staking interest (estimate), market-maker rewards; closed positions. Links to the CSV and PDF.
[**GetVaultMarketTaxCsv**](VaultMarketsApi.md#GetVaultMarketTaxCsv) | **Get** /v1/vault-markets/tax.csv | The tax year as a CSV: one row per sale, purchase, gift, income line and closed position.
[**GetVaultMarketTaxLossSuggestion**](VaultMarketsApi.md#GetVaultMarketTaxLossSuggestion) | **Get** /v1/vault-markets/tax/loss-suggestions | Tax-loss suggestions: the units you hold whose mark is below your cost, each with its lot (acquired when and how, cost in cents), short or long term if sold today and the day a short lot turns long, the estimated loss at the mark before selling fees, mark quality, what stops a sale, and a link to the ordinary sell ticket. With totals by term, the gains realised so far this year, and a not-tax-advice disclosure. Never sells.
[**GetVaultMarketTaxPdf**](VaultMarketsApi.md#GetVaultMarketTaxPdf) | **Get** /v1/vault-markets/tax.pdf | The tax year as a PDF summary.
[**GetVaultMarketTca**](VaultMarketsApi.md#GetVaultMarketTca) | **Get** /v1/vault-markets/tca | Your pro orders in a window (&#x60;from&#x60;, &#x60;to&#x60;; default 90 days) with arrival mark and mid, average fill, implementation shortfall, slippage, fill rate and time to fill.
[**GetVaultMarketTcaByOrderId**](VaultMarketsApi.md#GetVaultMarketTcaByOrderId) | **Get** /v1/vault-markets/tca/{orderId} | One order&#39;s TCA, including slippage against the VWAP of eligible trades while it worked.
[**GetVaultMarketTod**](VaultMarketsApi.md#GetVaultMarketTod) | **Get** /v1/vault-markets/tod | Your transfer-on-death setup: the gate and terms, whether you accepted them, your beneficiaries and shares, units with their designations and who the plan gives them to today, and open claims on your account (which you can stop). A TodOverview object, not a collection.
[**GetVaultMarketTodMyClaim**](VaultMarketsApi.md#GetVaultMarketTodMyClaim) | **Get** /v1/vault-markets/tod/my-claims | Claims you filed as a beneficiary: status in words, documents sent (names only), the waiting period, and units moved or held. A TodMyClaims object, not a collection.
[**GetVaultMarketUnitAr**](VaultMarketsApi.md#GetVaultMarketUnitAr) | **Get** /v1/vault-markets/units/{id}/ar | A vault unit as an AR model, for \&quot;See it in your room\&quot; (the web&#39;s /ar/vault-&lt;id&gt; page hands it to Scene Viewer, WebXR or Quick Look): title, model (a path under /models, without .glb, at real size), scanUrl (the unit&#39;s own captured mesh, when it has one), imageUrl, art, tint, scale, bookKey, backPath and mine. Your own units always; somebody else&#39;s only while they show it to visitors in a home you may enter (else 404).
[**GetVaultMarketUnitEvidence**](VaultMarketsApi.md#GetVaultMarketUnitEvidence) | **Get** /v1/vault-markets/units/{id}/evidence | Evidence on a vault unit:its 3D scans (model, texture, listing photos) and its grade report (grader, grade, cert, verified against the grader record, hub intake photos). No owner identity.
[**GetVaultMarketUnitQuote**](VaultMarketsApi.md#GetVaultMarketUnitQuote) | **Get** /v1/vault-markets/units/{id}/quote | The market around a vault unit: its mark and source, the best bid that could take it, the Instant Exit price (the best bid) and a pack pull&#39;s own sell-back credit.
[**GetVaultMarketUnitStory**](VaultMarketsApi.md#GetVaultMarketUnitStory) | **Get** /v1/vault-markets/units/{id}/story | One vault unit&#39;s story, oldest first: intake (with scan and grade evidence when it has some), grades, owners, sales at their prices, loans, pledges, Hall of Grails exhibitions, receipt-token exports and redemptions, and each warehouse-receipt hash-chain step (re-verified on read: chain.ok). People are named only while their public vault profile is on; others are &#39;A collector&#39;.
[**GetVaultMarketValuationStatement**](VaultMarketsApi.md#GetVaultMarketValuationStatement) | **Get** /v1/vault-markets/valuation-statements/{id} | One statement you issued: every unit, the totals by mark quality, the methodology, its sha256, its audit-chain seal and the verify link.
[**GetVaultMarketValuationStatementPdf**](VaultMarketsApi.md#GetVaultMarketValuationStatementPdf) | **Get** /v1/vault-markets/valuation-statements/{id}/pdf | One statement you issued as a PDF.
[**GetVaultMarketValuationStatementPreview**](VaultMarketsApi.md#GetVaultMarketValuationStatementPreview) | **Get** /v1/vault-markets/valuation-statements/preview | An insurance or appraisal statement for a day (purpose, asOf YYYY-MM-DD, default today), built but not issued: every unit you held that day at its mark, with quality, last trade, grade and cert, and the total. No id, no hash.
[**GetVaultMarketValuationStatementVerify**](VaultMarketsApi.md#GetVaultMarketValuationStatementVerify) | **Get** /v1/vault-markets/valuation-statements/{id}/verify | Check any issued statement by its id, as an insurer would: purpose, date, totals by mark quality and the sha256; intact (the stored record still hashes to it), seal.ok (its audit-chain entry still carries it) and, with ?hash&#x3D;, whether yours matches. Never the holder or the units.
[**GetVaultMarketValue**](VaultMarketsApi.md#GetVaultMarketValue) | **Get** /v1/vault-markets/value | A catalogue item&#39;s vault book keys (&#x60;?skuId&#x3D;&#x60;, optional &#x60;variantId&#x60;): each key&#39;s display mark with quality and n, units in the vault, the Instant Exit price a holder would get now, and whether it can be shipped in (ungraded keys).
[**GetVaultMarketVaultSuggestionItem**](VaultMarketsApi.md#GetVaultMarketVaultSuggestionItem) | **Get** /v1/vault-markets/vault-suggestions/items/{id} | One inventory item: its vault suggestion, or in plain words why there is none (&#x60;status&#x60;: suggested, not_found, no_stock, not_matched, in_vault or no_market).
[**GetVaultMarketWidgetIndex**](VaultMarketsApi.md#GetVaultMarketWidgetIndex) | **Get** /v1/vault-markets/widgets/index/{slug} | What the &lt;crossly-index&gt; widget shows: level, closes, quality, n, as-of, or &#x60;insufficient&#x60;.
[**GetVaultMarketWidgetMark**](VaultMarketsApi.md#GetVaultMarketWidgetMark) | **Get** /v1/vault-markets/widgets/mark | What the &lt;crossly-mark&gt; widget shows for &#x60;bookKey&#x60;: mark, quality, trades over 30 days, as-of, sparkline, or &#x60;insufficient&#x60;.
[**GetVaultMarketWidgetSummary**](VaultMarketsApi.md#GetVaultMarketWidgetSummary) | **Get** /v1/vault-markets/widget/summary | The home-screen widget summary: portfolio value at the mark with its move since the previous marks (cents and bps) and the share resting on reliable A/B marks, the three biggest movers among what you hold (the market movers when you hold nothing; &#x60;movers.source&#x60; says which), and one watchlist key (&#x60;?key&#x3D;&#x60; a book key on your watchlist; else your newest watch) with a 14-day sparkline. Every price is from the Crossly order book and carries its mark quality (A, B, C or none). Cached for five minutes.
[**GetVaultMarketWorldFloor**](VaultMarketsApi.md#GetVaultMarketWorldFloor) | **Get** /v1/vault-markets/world/floor | The Trading Floor in Crossly World: trades (the newest trades on every vault book in the last day, at most 40, newest first: bookKey, name, priceCents, priceText, side -- buy when a buyer took an ask, sell when a seller hit a bid, cross for an auction -- and ageS), tickers (the big board: the most-traded markets, each with its last trade price, the change from the last trade before today in basis points, trades today and the display mark with its quality), bell (Market Open: ringing for ten minutes after a new key opens, with the key and its market link, and what opens next) and trades24h. Every price is a trade on the Crossly order book; it never shows who traded. Nothing can be traded in the room.
[**GetVaultMarketWorldLounge**](VaultMarketsApi.md#GetVaultMarketWorldLounge) | **Get** /v1/vault-markets/world/lounge | The Vault Holders Lounge: the rule to join (held value at reliable marks, or a unit count), whether you may join now and what you hold toward it. Anyone may look in; only holders are seen there. No perks, nothing to trade.
[**GetVaultMarketWorldPlaceVaultGate**](VaultMarketsApi.md#GetVaultMarketWorldPlaceVaultGate) | **Get** /v1/vault-markets/world/places/{place}/vault-gate | The vault-holder door on a Crossly World room or Space, by its presence slug (a room share slug, or space-&lt;slug&gt;): the rule (minValueCents held at reliable A/B marks and/or minUnits units, mode any or all, optionally one product, bullion family or category), the rule in words, and whether you may join now with what you hold toward it. gate is null when the place has no vault door. Anyone may look in; only holders join.
[**GetVaultMarketWorth**](VaultMarketsApi.md#GetVaultMarketWorth) | **Get** /v1/vault-markets/worth | What one book key (variantId:condition:gradeKey, &#x60;bookKey&#x60;) is worth: the mark in cents with its quality (A from trades, B from firm quotes, C indicative, none) and band, what Instant Exit would pay for one unit now from the best resting bid (yours left out), and your units: how many you hold, how many are free to sell, their value at the mark and what selling every free one now would fetch from today’s bids. Plain-words &#x60;headline&#x60; and &#x60;lines&#x60;. Read only.
[**GetVaultMarketWorthVault**](VaultMarketsApi.md#GetVaultMarketWorthVault) | **Get** /v1/vault-markets/worth/vault | What your whole vault is worth now: value at the marks split into reliable and indicative marks, cost and P&amp;L, today’s change, and an estimate of what the units free to sell would fetch from the bids resting on the book now. Read only.
[**GetVaultReceipt**](VaultMarketsApi.md#GetVaultReceipt) | **Get** /v1/vault/receipts/{receiptNo} | One of your receipts with its full hash chain (issued, transferred, pledged, released, surrendered), re-verified on read: &#x60;chain.ok&#x60; is false if any event was altered.
[**GetVaultReceiptToken**](VaultMarketsApi.md#GetVaultReceiptToken) | **Get** /v1/vault/receipt-tokens/{id} | One receipt token (yours, or on a unit you own now) with its chain log: built, minted, frozen, thawed, new holder seen, redeem built, burned, with explorer links.
[**GetVaultReceiptTokenByMint**](VaultMarketsApi.md#GetVaultReceiptTokenByMint) | **Get** /v1/vault/receipt-tokens/by-mint/{mint} | What a Solana mint is: the warehouse receipt number and item it stands for, and whether the token is live, frozen (unit pledged or lent) or burned, and whether the unit is still on a vault shelf. Names no person and no wallet.
[**ListVaultMarketAccounts**](VaultMarketsApi.md#ListVaultMarketAccounts) | **Get** /v1/vault-markets/accounts | Every account you can act on: your own, and each sub-account your entity memberships reach, with your role.
[**ListVaultMarketAdlEvents**](VaultMarketsApi.md#ListVaultMarketAdlEvents) | **Get** /v1/vault-markets/adl/events | The ADL events your positions were part of, newest first: role bankrupt (yours went under water) or counterparty (yours was first in line), the price both closed at and the transfer status.
[**ListVaultMarketAdlPositions**](VaultMarketsApi.md#ListVaultMarketAdlPositions) | **Get** /v1/vault-markets/adl/positions | Your live positions, each with its place in its item ADL queue: lights 5 (first in line) to 1 (last), 0 when not in the queue (only profitable positions are), ranked by P&amp;L% x effective leverage.
[**ListVaultMarketAgentDrafts**](VaultMarketsApi.md#ListVaultMarketAgentDrafts) | **Get** /v1/vault-markets/agent/drafts | Limit orders the shopping agent drafted for you, newest first. A draft places nothing: send its orderBody to POST /api/v1/vault-markets/pro/orders to place it.
[**ListVaultMarketAgentWatches**](VaultMarketsApi.md#ListVaultMarketAgentWatches) | **Get** /v1/vault-markets/agent/watches | Your price watches on book keys (best bid, lowest ask, last trade or market price, in cents), live ones first. &#x60;?live&#x3D;true&#x60; for live only.
[**ListVaultMarketAlertLog**](VaultMarketsApi.md#ListVaultMarketAlertLog) | **Get** /v1/vault-markets/alert-log | What the alerts hub did with your recent vault alerts (&#x60;?limit&#x3D;&#x60; 1 to 200, default 50): kind, title, the channels it went to, what was held (channel:reason) and the email digest state.
[**ListVaultMarketAlerts**](VaultMarketsApi.md#ListVaultMarketAlerts) | **Get** /v1/vault-markets/alerts | Your price alerts on book keys.
[**ListVaultMarketAuctions**](VaultMarketsApi.md#ListVaultMarketAuctions) | **Get** /v1/vault-markets/auctions | Open vault auction events (weekly, ending Sunday; premier, units marked $10,000+) and their lots, soonest ending first.
[**ListVaultMarketBlocks**](VaultMarketsApi.md#ListVaultMarketBlocks) | **Get** /v1/vault-markets/blocks | An entity&#39;s block orders, newest first, each with its plan and how it was allocated.
[**ListVaultMarketBullionFamilyPegs**](VaultMarketsApi.md#ListVaultMarketBullionFamilyPegs) | **Get** /v1/vault-markets/bullion/family-pegs | Your family-wide spot-pegged quotes (live, or all with status&#x3D;all): one quote resting across a bullion family books at one condition, with units left, the peg, its status and every leg (the pegged quote on each book).
[**ListVaultMarketBullionKeys**](VaultMarketsApi.md#ListVaultMarketBullionKeys) | **Get** /v1/vault-markets/bullion/keys | Bullion key families (metal, form, size, mint or any mint): members, vault units, best resting bid and ask across their books, and spot as a labelled reference with what one unit of metal is worth at it. Spot is never a Crossly price.
[**ListVaultMarketBullionPegs**](VaultMarketsApi.md#ListVaultMarketBullionPegs) | **Get** /v1/vault-markets/bullion/pegs | Your spot-pegged quotes (live, or all with status&#x3D;all): side, units, the peg, the order it rests as, the last price and spot it was priced from, and why it is pulled when it is.
[**ListVaultMarketCollectionWorth**](VaultMarketsApi.md#ListVaultMarketCollectionWorth) | **Get** /v1/vault-markets/collection-worth | Your collection valued at the display marks: rows on a vault book key with a mark are summed; rows with no mark or no key are counted, never valued.
[**ListVaultMarketConfirmations**](VaultMarketsApi.md#ListVaultMarketConfirmations) | **Get** /v1/vault-markets/confirmations | Trade confirmations for your fills in a window (default 31 days): price, sale fee, maker or taker, order, vault unit and warehouse receipt.
[**ListVaultMarketCreatorCohostInvites**](VaultMarketsApi.md#ListVaultMarketCreatorCohostInvites) | **Get** /v1/vault-markets/creators/cohost-invites | For creators: your invites from Crossly to co-host a Market Open launch show (invited, accepted, declined or withdrawn), each with the key, its show and the guidance for co-hosts.
[**ListVaultMarketCredit**](VaultMarketsApi.md#ListVaultMarketCredit) | **Get** /v1/vault-markets/credit | Your vault loans and advances (live by default): what is owed, interest accrued, loan-to-value, the liquidation line and the collateral value it is called at.
[**ListVaultMarketDealerBands**](VaultMarketsApi.md#ListVaultMarketDealerBands) | **Get** /v1/vault-markets/dealer/bands | Your Instant Exit bands and whether each is being kept now.
[**ListVaultMarketDealerBids**](VaultMarketsApi.md#ListVaultMarketDealerBids) | **Get** /v1/vault-markets/dealer/bids | Your standing dealer bids: open (default) or all.
[**ListVaultMarketDealerPoolFills**](VaultMarketsApi.md#ListVaultMarketDealerPoolFills) | **Get** /v1/vault-markets/dealer/pools/{id}/fills | A curve pool&#39;s fills, newest first: side, price and spot after.
[**ListVaultMarketDealerPools**](VaultMarketsApi.md#ListVaultMarketDealerPools) | **Get** /v1/vault-markets/dealer/pools | Your curve pools: open (default) or all, each with its curve, spot, units, buy budget, running totals, the ladder it shows and its live legs.
[**ListVaultMarketDigestHistory**](VaultMarketsApi.md#ListVaultMarketDigestHistory) | **Get** /v1/vault-markets/digest/history | Past weekly digests: week, status (sent / skipped / failed), the channels used and a summary.
[**ListVaultMarketEntityAudit**](VaultMarketsApi.md#ListVaultMarketEntityAudit) | **Get** /v1/vault-markets/entities/{id}/audit | The entity&#39;s audit trail, newest first: hash-chained rows with actor, account, API key, IP, session and source.
[**ListVaultMarketEntityConversions**](VaultMarketsApi.md#ListVaultMarketEntityConversions) | **Get** /v1/vault-markets/entities/{id}/conversions | The custodian&#39;s CBX-to-USD reward conversions for an entity that elected USD, newest first, with their status.
[**ListVaultMarketEstimatesAccuracy**](VaultMarketsApi.md#ListVaultMarketEstimatesAccuracy) | **Get** /v1/vault-markets/estimates-accuracy | The Crossly Estimate&#39;s track record: walk-forward backtest n, FSD and bias per method, category and liquidity bucket, and whether estimates in that cell are shown.
[**ListVaultMarketFeed**](VaultMarketsApi.md#ListVaultMarketFeed) | **Get** /v1/vault-markets/feed | New vault buys by the people you follow who share them, newest first (no prices). Page with &#x60;before&#x60; (the last item&#39;s &#x60;at&#x60;).
[**ListVaultMarketFollowCopyWatch**](VaultMarketsApi.md#ListVaultMarketFollowCopyWatch) | **Get** /v1/vault-markets/follows/copy-watch | The collectors you copy-watch and whether each still shares bids and asks.
[**ListVaultMarketFollowCopyWatchAlerts**](VaultMarketsApi.md#ListVaultMarketFollowCopyWatchAlerts) | **Get** /v1/vault-markets/follows/copy-watch/alerts | Your copy-watch alerts, newest first: who, bid or ask, quantity, limit price, market, and whether it still rests. Page with &#x60;before&#x60;.
[**ListVaultMarketFollowFeed**](VaultMarketsApi.md#ListVaultMarketFollowFeed) | **Get** /v1/vault-markets/follows/feed | Public vault activity of the collectors you follow (buys, units sent in, sales, Hall loans), newest first, never a price. Filter with &#x60;kinds&#x60;; page with &#x60;before&#x60;.
[**ListVaultMarketFollowFollowing**](VaultMarketsApi.md#ListVaultMarketFollowFollowing) | **Get** /v1/vault-markets/follows/following | The collectors you follow who have a public vault: what each shares and whether you copy-watch them.
[**ListVaultMarketFollowLeaderboards**](VaultMarketsApi.md#ListVaultMarketFollowLeaderboards) | **Get** /v1/vault-markets/follows/leaderboards | An opt-in return leaderboard (&#x60;period&#x60; 30d, 90d or 365d; &#x60;category&#x60; all or a category slug), ranked by risk-adjusted return from the book A/B marks. Percent and risk only: no dollars, no volume.
[**ListVaultMarketFollowProfileActivity**](VaultMarketsApi.md#ListVaultMarketFollowProfileActivity) | **Get** /v1/vault-markets/follows/profiles/{handle}/activity | One collector public vault activity, newest first, never a price. 404 when their vault is not public.
[**ListVaultMarketGifts**](VaultMarketsApi.md#ListVaultMarketGifts) | **Get** /v1/vault-markets/gifts | Gifts you sent and received, newest first, with their status.
[**ListVaultMarketGradeToVault**](VaultMarketsApi.md#ListVaultMarketGradeToVault) | **Get** /v1/vault-markets/grade-to-vault | Your Grade to Vault flows, newest first, each brought up to date: its step, the grading submission, the unit, the grade, the market for it, and the ask once listed. &#x60;live&#x3D;1&#x60; for those still going.
[**ListVaultMarketGrading**](VaultMarketsApi.md#ListVaultMarketGrading) | **Get** /v1/vault-markets/grading | Your vault units sent to grading: where each card is in its run, and the grade once it is back.
[**ListVaultMarketHallLenders**](VaultMarketsApi.md#ListVaultMarketHallLenders) | **Get** /v1/vault-markets/hall/lenders | The vault units lent to this week&#39;s Hall of Grails exhibition and on show now, in plinth order: the item (skuId, name, imageUrl, grade), the lender as the plate names them (a display name, or \&quot;a collector\&quot;), the plate lines and how many exhibitions the unit has stood in. meta.exhibition is the week (key, title, endsAt). Names no ids and no amounts.
[**ListVaultMarketIndexCommitteeMinutes**](VaultMarketsApi.md#ListVaultMarketIndexCommitteeMinutes) | **Get** /v1/vault-markets/index-committee/minutes | Published Index Committee minutes, newest first: attendees (names and roles), quorum, agenda, minutes and decisions.
[**ListVaultMarketIndexComplaints**](VaultMarketsApi.md#ListVaultMarketIndexComplaints) | **Get** /v1/vault-markets/index-complaints | Your complaints and their answers, newest first.
[**ListVaultMarketIndexConstituents**](VaultMarketsApi.md#ListVaultMarketIndexConstituents) | **Get** /v1/vault-markets/indexes/{slug}/constituents | The book keys in an index with their weight, move since the last close and contribution to the index&#39;s move (basis points).
[**ListVaultMarketIndexConsultations**](VaultMarketsApi.md#ListVaultMarketIndexConsultations) | **Get** /v1/vault-markets/index-consultations | Consultations on index method changes (open for at least 30 days), newest first. Filter by &#x60;status&#x60;.
[**ListVaultMarketIndexEvents**](VaultMarketsApi.md#ListVaultMarketIndexEvents) | **Get** /v1/vault-markets/index-events | Every index event, newest first: add, delete, rebalance, method, correction, halt; each with when it was announced, when it takes effect, the detail and the reason. Filter by &#x60;kind&#x60;.
[**ListVaultMarketIndexEventsBySlug**](VaultMarketsApi.md#ListVaultMarketIndexEventsBySlug) | **Get** /v1/vault-markets/indexes/{slug}/events | One index&#39;s events, newest first.
[**ListVaultMarketIndexJudgments**](VaultMarketsApi.md#ListVaultMarketIndexJudgments) | **Get** /v1/vault-markets/index-judgments | The expert-judgment log: every Index Committee action (halt, resume, exclude a key, void a trade, restate, retire) with its written reason.
[**ListVaultMarketIndexLevelSeries**](VaultMarketsApi.md#ListVaultMarketIndexLevelSeries) | **Get** /v1/vault-markets/indexes/{slug}/level-series | An index&#39;s published levels. &#x60;engine&#x3D;chain&#x60; (default): daily closes. &#x60;engine&#x3D;tpd|repeat_sales&#x60;: weekly (else monthly) model estimates with a 95% band and n; labelled, never a price. &#x60;series&#x3D;settlement&#x60; returns each value as first published (never restated). Below the data gate &#x60;data&#x60; is empty and &#x60;notice&#x60; explains.
[**ListVaultMarketIndexLicenceMine**](VaultMarketsApi.md#ListVaultMarketIndexLicenceMine) | **Get** /v1/vault-markets/index-licence/mine | Your index feed licences (the indexes each covers, or all), as ops recorded them. No licence: the feed answers 403 not_a_licensee.
[**ListVaultMarketIndexMembers**](VaultMarketsApi.md#ListVaultMarketIndexMembers) | **Get** /v1/vault-markets/indexes/{slug}/members | The members of an index on a day (&#x60;asOf&#x60;, default today) with their weights, shares, moves and contributions; for a past day, the inputs stored with that day&#39;s close and each member&#39;s effective dates.
[**ListVaultMarketIndexMovers**](VaultMarketsApi.md#ListVaultMarketIndexMovers) | **Get** /v1/vault-markets/indexes/movers | The built-in indexes that moved most since their last daily close.
[**ListVaultMarketIndexes**](VaultMarketsApi.md#ListVaultMarketIndexes) | **Get** /v1/vault-markets/indexes | Indexes over book keys: per category, per top brand, the Crossly 50 and your own. Each with its level now (base 1000 at inception) and change since the last daily close.
[**ListVaultMarketInstantExitHistory**](VaultMarketsApi.md#ListVaultMarketInstantExitHistory) | **Get** /v1/vault-markets/instant-exit/history | Your Instant Exits, newest first: quoted price, mark, and whether it filled.
[**ListVaultMarketInstantExitHoldings**](VaultMarketsApi.md#ListVaultMarketInstantExitHoldings) | **Get** /v1/vault-markets/instant-exit/holdings | Instant Exit quotes for everything you hold in the vault, one per book key, with units held and which are free to sell.
[**ListVaultMarketJointTrades**](VaultMarketsApi.md#ListVaultMarketJointTrades) | **Get** /v1/vault-markets/joint/{id}/trades | The joint vault trades, newest first: side, book key, limit, who started it, whose account the order sits on, and the order standing on the book.
[**ListVaultMarketLargeTraderReports**](VaultMarketsApi.md#ListVaultMarketLargeTraderReports) | **Get** /v1/vault-markets/large-trader-reports | Your large-trader reports (with any linked accounts, as one party): due first, then filed, then cleared in the last 90 days. Each has &#x60;action&#x60;: file | amend | refile | none.
[**ListVaultMarketLendingStatements**](VaultMarketsApi.md#ListVaultMarketLendingStatements) | **Get** /v1/vault-markets/lending/statements | The days you have stored collateral statements for, newest first, with their summary.
[**ListVaultMarketLoans**](VaultMarketsApi.md#ListVaultMarketLoans) | **Get** /v1/vault-markets/loans | Units you lent: rate, fees earned, recall due.
[**ListVaultMarketLocateMine**](VaultMarketsApi.md#ListVaultMarketLocateMine) | **Get** /v1/vault-markets/locates/mine | Your locates, newest first (held / used / released / expired).
[**ListVaultMarketMidpoint**](VaultMarketsApi.md#ListVaultMarketMidpoint) | **Get** /v1/vault-markets/midpoint | Your midpoint block orders.
[**ListVaultMarketMovers**](VaultMarketsApi.md#ListVaultMarketMovers) | **Get** /v1/vault-markets/movers | Vault items with the biggest 24 h moves in their market price (from the order book), with units in the vault.
[**ListVaultMarketNavHistory**](VaultMarketsApi.md#ListVaultMarketNavHistory) | **Get** /v1/vault-markets/nav/history | Daily NAV snapshots (mark, bid, liquidation, Level 2 / Level 3).
[**ListVaultMarketOffers**](VaultMarketsApi.md#ListVaultMarketOffers) | **Get** /v1/vault-markets/offers | The offers inbox: for each book key you hold, the best bid from someone else, units at it, the top bid levels, the mark, and your alert.
[**ListVaultMarketOpening**](VaultMarketsApi.md#ListVaultMarketOpening) | **Get** /v1/vault-markets/opening | Keys opening soon on the vault book (and just opened): units on the shelf, dealers quoting, the waitlist, the mark when the book has one, and your place.
[**ListVaultMarketOracleHistory**](VaultMarketsApi.md#ListVaultMarketOracleHistory) | **Get** /v1/vault-markets/oracle/{slug}/history | Recent publish attempts for an index over 30 days, newest first: published, skipped (with the reason: too few items priced, stale marks, confidence too wide...) or failed (with the error). Filter with &#x60;status&#x60;, cap with &#x60;limit&#x60; (default 100).
[**ListVaultMarketP2pAuctions**](VaultMarketsApi.md#ListVaultMarketP2pAuctions) | **Get** /v1/vault-markets/p2p/auctions | Running refinancing auctions, soonest to end first: lenders calls (&#x60;auction&#x60;) and borrowers looking for a cheaper lender (&#x60;refiAuction&#x60;). The payoff now, the rate a new lender gets now, when it ends, and the item.
[**ListVaultMarketP2pLoans**](VaultMarketsApi.md#ListVaultMarketP2pLoans) | **Get** /v1/vault-markets/p2p/loans | Your P2P loans as borrower and/or lender (&#x60;role&#x60;), live by default (&#x60;status&#x60;): what is owed now, daily interest, the refinancing auction when called, and the item mark for information only (it never triggers anything).
[**ListVaultMarketP2pOfferMine**](VaultMarketsApi.md#ListVaultMarketP2pOfferMine) | **Get** /v1/vault-markets/p2p/offers/mine | Your P2P loan offers, newest first, with how many loans each has funded.
[**ListVaultMarketP2pOffers**](VaultMarketsApi.md#ListVaultMarketP2pOffers) | **Get** /v1/vault-markets/p2p/offers | The open P2P loan offers book, lowest rate first: cents offered per loan, APR (bps), loans left, and the item (book key) or unit each lends against. &#x60;bookKey&#x60; narrows to one item. Lender identities are never shown.
[**ListVaultMarketP2pUnits**](VaultMarketsApi.md#ListVaultMarketP2pUnits) | **Get** /v1/vault-markets/p2p/units | Your vault units for P2P borrowing: whether each can be pledged now (and why not), the best offers covering it, and its loan when it has one.
[**ListVaultMarketPaperFills**](VaultMarketsApi.md#ListVaultMarketPaperFills) | **Get** /v1/vault-markets/paper/fills | Your paper fills since your last reset, newest first: price, simulated fee, taker or maker, and whether a visible real order (book) or a real trade (tape) filled it. Play money only.
[**ListVaultMarketPaperLeaderboard**](VaultMarketsApi.md#ListVaultMarketPaperLeaderboard) | **Get** /v1/vault-markets/paper/leaderboard | The weekly paper trading leaderboard (opt-in, display names only): best return on play equity this week (or a past &#x60;week&#x60;, its Monday as YYYY-MM-DD), with your own line. Nothing is won or paid.
[**ListVaultMarketPaperMissions**](VaultMarketsApi.md#ListVaultMarketPaperMissions) | **Get** /v1/vault-markets/paper/missions | Your paper-trading missions (play money only): each mission with what to do, your count toward its target, whether it is done and when, plus a suggested market for the \&quot;Read the mark\&quot; mission. Progress is observed by the server from your paper orders and fills, price watches, set goals and graded mark checks. Missions earn world quest progress and cosmetic pins, never money.
[**ListVaultMarketPaperOrders**](VaultMarketsApi.md#ListVaultMarketPaperOrders) | **Get** /v1/vault-markets/paper/orders | Your paper orders since your last reset, newest first: working, done (filled, cancelled, expired) or all. Play money only.
[**ListVaultMarketPartnerAccruals**](VaultMarketsApi.md#ListVaultMarketPartnerAccruals) | **Get** /v1/vault-markets/partners/accruals | Your partner fee shares, newest first: accrued on a unit&#39;s first sale, paid in Bucks once the sale is final, void if it was refunded.
[**ListVaultMarketPoolStakes**](VaultMarketsApi.md#ListVaultMarketPoolStakes) | **Get** /v1/vault-markets/pool/stakes | Your stakes in the market pool and what each redeems for now.
[**ListVaultMarketPortfolioMarginHistory**](VaultMarketsApi.md#ListVaultMarketPortfolioMarginHistory) | **Get** /v1/vault-markets/portfolio-margin/history | The 5-minute portfolio margin recalculations for a firm sub-account over the last &#x60;hours&#x60; (default 24, at most 720), newest first: equity, requirement, excess, scenario and floor figures, gross and state.
[**ListVaultMarketPositionUnits**](VaultMarketsApi.md#ListVaultMarketPositionUnits) | **Get** /v1/vault-markets/positions/{id}/units | The vault units of one of your positions that you hold (a leveraged long&#39;s collateral), each with its story link.
[**ListVaultMarketPositions**](VaultMarketsApi.md#ListVaultMarketPositions) | **Get** /v1/vault-markets/positions | Your positions, live (default) or closed.
[**ListVaultMarketPositionsRisk**](VaultMarketsApi.md#ListVaultMarketPositionsRisk) | **Get** /v1/vault-markets/positions-risk | For each live position: its tier, health, margin state, the mark at which it is called and the mark at which it is liquidated, and the distance to it.
[**ListVaultMarketPriceWidgets**](VaultMarketsApi.md#ListVaultMarketPriceWidgets) | **Get** /v1/vault-markets/price-widgets | Your price widgets, newest first, each with its iframe and script-tag snippets. &#x60;meta.codes&#x60; lists your affiliate codes a widget can link through; &#x60;meta.max&#x60; is the cap on live widgets.
[**ListVaultMarketProCandles**](VaultMarketsApi.md#ListVaultMarketProCandles) | **Get** /v1/vault-markets/pro/candles | OHLCV candles (1m, 5m, 15m, 1h, 4h, 1d) from the trades on a book key.
[**ListVaultMarketProEvents**](VaultMarketsApi.md#ListVaultMarketProEvents) | **Get** /v1/vault-markets/pro/events | Your order event log after a sequence number (new, leg, fill, replaced, cancelled, expired, done): a replayable drop copy.
[**ListVaultMarketProLinks**](VaultMarketsApi.md#ListVaultMarketProLinks) | **Get** /v1/vault-markets/pro/links | Your linked orders (brackets, OCO pairs, trailing stops, position exits): live (default) or done.
[**ListVaultMarketProOrders**](VaultMarketsApi.md#ListVaultMarketProOrders) | **Get** /v1/vault-markets/pro/orders | Your orders: open (default) or done, optionally on one book key, or only those booked to one sub-account (&#x60;subaccountId&#x60;).
[**ListVaultMarketProQuotes**](VaultMarketsApi.md#ListVaultMarketProQuotes) | **Get** /v1/vault-markets/pro/quotes | Your live quotes.
[**ListVaultMarketProRewards**](VaultMarketsApi.md#ListVaultMarketProRewards) | **Get** /v1/vault-markets/pro/rewards | Your daily market-maker liquidity rewards (paid in CBX).
[**ListVaultMarketProTape**](VaultMarketsApi.md#ListVaultMarketProTape) | **Get** /v1/vault-markets/pro/tape | The latest trades on a book key, newest first, with which side made and which took.
[**ListVaultMarketProofOfReserveLeaves**](VaultMarketsApi.md#ListVaultMarketProofOfReserveLeaves) | **Get** /v1/vault-markets/proof-of-reserves/{day}/leaves | One day&#39;s leaf hashes in tree order, a page at a time: enough to recount the units and rebuild the root yourself. Leaves are salted, so they name no owner.
[**ListVaultMarketRecurring**](VaultMarketsApi.md#ListVaultMarketRecurring) | **Get** /v1/vault-markets/recurring | Your recurring buys: key, schedule, price cap, payment, auto-lend, next run.
[**ListVaultMarketRecurringRuns**](VaultMarketsApi.md#ListVaultMarketRecurringRuns) | **Get** /v1/vault-markets/recurring/{id}/runs | One recurring buy&#39;s runs: filled, skipped (with the reason, e.g. the price was over the cap) or failed.
[**ListVaultMarketRecurringRuns2**](VaultMarketsApi.md#ListVaultMarketRecurringRuns2) | **Get** /v1/vault-markets/recurring-runs | Every run of all your recurring buys, newest first.
[**ListVaultMarketRfq**](VaultMarketsApi.md#ListVaultMarketRfq) | **Get** /v1/vault-markets/rfq | Your quote requests, newest first.
[**ListVaultMarketRfqInbox**](VaultMarketsApi.md#ListVaultMarketRfqInbox) | **Get** /v1/vault-markets/rfq/inbox | Open quote requests you were sent (as a dealer, DMM, LP or holder), with your live answer.
[**ListVaultMarketRfqOpen**](VaultMarketsApi.md#ListVaultMarketRfqOpen) | **Get** /v1/vault-markets/rfq/open | Open quote requests on one key (all-to-all: anyone with ID may answer). Anonymous.
[**ListVaultMarketSetGoalBaskets**](VaultMarketsApi.md#ListVaultMarketSetGoalBaskets) | **Get** /v1/vault-markets/set-goals/{id}/baskets | A set goal basket bids, newest first, each with its budget, what it committed and spent, and every leg (piece, book key, limit, pro order id, status).
[**ListVaultMarketSetGoalSets**](VaultMarketsApi.md#ListVaultMarketSetGoalSets) | **Get** /v1/vault-markets/set-goals/sets | Sets to start a goal on. With q (2+ characters): a word-start search of every catalogue set. Without: the sets you hold pieces of in the vault, then the collections you are building. Each with your live goal on it, if any.
[**ListVaultMarketSetGoals**](VaultMarketsApi.md#ListVaultMarketSetGoals) | **Get** /v1/vault-markets/set-goals | Your set goals: the set, the target a piece must meet, pieces held of the total (from your vault), and the working basket bid.
[**ListVaultMarketShareCards**](VaultMarketsApi.md#ListVaultMarketShareCards) | **Get** /v1/vault-markets/share-cards | Your share cards (not taken down), each as the public sees it.
[**ListVaultMarketSupplyDeposits**](VaultMarketsApi.md#ListVaultMarketSupplyDeposits) | **Get** /v1/vault-markets/supply/deposits | Your units sent to the vault, newest first, each with its state and its key&#39;s Instant Exit line (\&quot;Instant Exit now: $X (N% of mark)\&quot; or \&quot;No market yet\&quot;). Filter with &#x60;inventoryItemIds&#x60; (comma-separated).
[**ListVaultMarketSupplyInvites**](VaultMarketsApi.md#ListVaultMarketSupplyInvites) | **Get** /v1/vault-markets/supply/invites | Market Open invites to you as a dealer: keys Crossly is opening and asks you to quote.
[**ListVaultMarketValuationStatements**](VaultMarketsApi.md#ListVaultMarketValuationStatements) | **Get** /v1/vault-markets/valuation-statements | The insurance and appraisal statements you issued, newest first: purpose, date, units, total at the marks, hash and verify link.
[**ListVaultMarketVaultSuggestions**](VaultMarketsApi.md#ListVaultMarketVaultSuggestions) | **Get** /v1/vault-markets/vault-suggestions | What to send to the vault: your inventory items whose book key has a liquid vault market (tier A, B or C), best first. Each has a score (0 to 100) and its reasons, the market (tier, mark and its quality, best bid, best vault ask, spread), Instant Exit now, your list price, days listed without selling, and the vault fee. &#x60;meta.summary&#x60; counts what was left out and why. Inform only; send with POST /api/v1/vault-markets/supply/send.
[**ListVaultMarketWatchlist**](VaultMarketsApi.md#ListVaultMarketWatchlist) | **Get** /v1/vault-markets/watchlist | Your watched book keys with mark, 24 h and 7 d change, and the last 30 daily marks (sparkline).
[**ListVaultMarketWidgetTokens**](VaultMarketsApi.md#ListVaultMarketWidgetTokens) | **Get** /v1/vault-markets/widget/tokens | Your live home-screen widget tokens, one per phone: id, name (the phone), public prefix, created and last used. The secret is never returned again.
[**ListVaultMarketWorldCohosts**](VaultMarketsApi.md#ListVaultMarketWorldCohosts) | **Get** /v1/vault-markets/world/cohosts | Market Open launch shows a verified creator co-hosts that are on now or opening soon: the key, its phase, the show, the creator (named only while their vault profile is public), their room link and, when the creator is in the creator programme, the disclosure that they may earn a share of the fee Crossly keeps. Never a prompt to trade.
[**ListVaultMarketWorldCreatorRooms**](VaultMarketsApi.md#ListVaultMarketWorldCreatorRooms) | **Get** /v1/vault-markets/world/creator-rooms | Verified creators whose public vault shelf is a world room: title, the creator, the room link and the disclosure. Rooms show each market share of the shelf, never amounts.
[**ListVaultMarketWorldMarketOpen**](VaultMarketsApi.md#ListVaultMarketWorldMarketOpen) | **Get** /v1/vault-markets/world/market-open | Market Open events in the Crossly World Commons: keys opening soon and those that opened in the last few hours, live first then soonest, each with its board line, units on the shelf, dealers quoting and the launch show when there is one. Never a prompt to trade.
[**ListVaultReceiptTokenWallets**](VaultMarketsApi.md#ListVaultReceiptTokenWallets) | **Get** /v1/vault/receipt-tokens/wallets | Your verified Solana wallets and whether each is on the receipt-token allow-list (receipt tokens move only between allow-listed wallets). meta.state is the gate.
[**ListVaultReceiptTokens**](VaultMarketsApi.md#ListVaultReceiptTokens) | **Get** /v1/vault/receipt-tokens | Your vault units with their warehouse receipt and receipt-token state (not exported, waiting for a signature, in a wallet, frozen, redeemed), whether each can be exported now and why not. meta.state is the gate (open, missing envs, program id, authority, terms version, cluster); meta.history is your last 50 exports.
[**ListVaultReceipts**](VaultMarketsApi.md#ListVaultReceipts) | **Get** /v1/vault/receipts | Your warehouse receipts (documents of title for your vault units), live first. meta.receipts says whether receipts are registered (VAULT_WAREHOUSE_ISSUER) and their terms version.
[**ListVaultUnitList**](VaultMarketsApi.md#ListVaultUnitList) | **Get** /v1/vault/unit-list | The published unit list (bar-list equivalent): every unit on a vault shelf with its warehouse receipt number (when receipts are registered), book key, shelf location code, insured value, status, pledged / on-loan flags and a per-receipt hashed holder (never a name or id). Filter with &#x60;bookKey&#x60;; pages with &#x60;page&#x60; / &#x60;limit&#x60;.
[**UpdateVaultMarketAlertPref**](VaultMarketsApi.md#UpdateVaultMarketAlertPref) | **Put** /v1/vault-markets/alert-prefs | Change vault alert settings: &#x60;matrix&#x60; ({ kind: { channel: true|false } }, merged onto what is set), &#x60;quietHours&#x60; ({ start: \&quot;HH:MM\&quot;, end: \&quot;HH:MM\&quot;, tz } or null to clear) and &#x60;telegram&#x60; (the Telegram master switch). Locked switches are ignored. Returns the new settings.
[**UpdateVaultMarketBlockAllocation**](VaultMarketsApi.md#UpdateVaultMarketBlockAllocation) | **Put** /v1/vault-markets/blocks/{orderId}/allocation | Post-trade allocation: set or replace the plan for a buy you placed for one of the entity&#39;s sub-accounts (that sub-account leads). Refused once any share has moved.
[**UpdateVaultMarketChannel**](VaultMarketsApi.md#UpdateVaultMarketChannel) | **Put** /v1/vault-markets/channels | Change the switches: &#x60;telegramAlerts&#x60;, &#x60;telegramDigest&#x60;, &#x60;pushCapPerDay&#x60; (0 to 20). Margin calls and recalls are never capped. Returns the new view.
[**UpdateVaultMarketDealerBand**](VaultMarketsApi.md#UpdateVaultMarketDealerBand) | **Put** /v1/vault-markets/dealer/bands | Commit to an Instant Exit band on a key: a standing bid for &#x60;units&#x60; kept at &#x60;bandBps&#x60; of the mark, never above &#x60;maxPriceCents&#x60;, re-quoted as the mark moves.
[**UpdateVaultMarketDealerPool**](VaultMarketsApi.md#UpdateVaultMarketDealerPool) | **Patch** /v1/vault-markets/dealer/pools/{id} | Change a curve pool: spot, step, fee, levels, units, buy budget, floor, ceiling, or pause / resume. Re-quoted at once.
[**UpdateVaultMarketEntityMember**](VaultMarketsApi.md#UpdateVaultMarketEntityMember) | **Patch** /v1/vault-markets/entities/{id}/members/{memberId} | Admin: change a member&#39;s role, limits or sub-accounts. The last admin cannot be demoted.
[**UpdateVaultMarketEntitySubaccountLimit**](VaultMarketsApi.md#UpdateVaultMarketEntitySubaccountLimit) | **Put** /v1/vault-markets/entities/{id}/subaccounts/{subId}/limits | Admin: set a sub-account&#39;s limits (per-order notional, units per key, gross, net, leverage), held to the entity&#39;s caps. &#x60;limits: null&#x60; clears them.
[**UpdateVaultMarketEntityUsdElection**](VaultMarketsApi.md#UpdateVaultMarketEntityUsdElection) | **Put** /v1/vault-markets/entities/{id}/usd-election | Admin: elect USD for the entity (&#x60;usdOnly&#x60;). Its sub-accounts then post collateral in USD only and pay every fee in USD, and any CBX reward paid to one of them is converted to USD by the custodian in that account. Audited; it can be turned off again.
[**UpdateVaultMarketFollowCopyWatch**](VaultMarketsApi.md#UpdateVaultMarketFollowCopyWatch) | **Put** /v1/vault-markets/follows/copy-watch/{handle} | Copy-watch a collector who shares bids and asks: you get an alert when they rest a new one. It never places an order for you.
[**UpdateVaultMarketFollowSharing**](VaultMarketsApi.md#UpdateVaultMarketFollowSharing) | **Put** /v1/vault-markets/follows/sharing | Turn sharing switches on or off (&#x60;isPublic&#x60;, &#x60;showBuys&#x60;, &#x60;deposits&#x60;, &#x60;sales&#x60;, &#x60;hallLoans&#x60;, &#x60;orders&#x60;, &#x60;leaderboard&#x60;). A switch shows only what happens after it goes on; all but the profile itself need the profile public.
[**UpdateVaultMarketHomeUnit**](VaultMarketsApi.md#UpdateVaultMarketHomeUnit) | **Put** /v1/vault-markets/home/units | Show exactly these vault units in your Crossly World home: { units: [{ vaultItemId, visitors }] }. Units left out come out of their cases; units already standing stay where they are; new ones go into a free display-case or shelf slot, else a glass case of their own along a wall. visitors: true lets people who may enter your home see that unit (default false: only you). Returns placed, unplaced (no free case: the home is full), refused (with why) and the desk. Display only: nothing moves.
[**UpdateVaultMarketIndex**](VaultMarketsApi.md#UpdateVaultMarketIndex) | **Patch** /v1/vault-markets/indexes/{slug} | Rename, re-weight, share or change the keys of an index of your own; its level carries on.
[**UpdateVaultMarketLendingElection**](VaultMarketsApi.md#UpdateVaultMarketLendingElection) | **Put** /v1/vault-markets/lending/election | Change your lender election (applies to loans opened after the change). Needs lending.
[**UpdateVaultMarketMmMmp**](VaultMarketsApi.md#UpdateVaultMarketMmMmp) | **Put** /v1/vault-markets/mm/mmp | Set market-maker protection: N units or $X filled on your quotes within T seconds pulls every quote and freezes quoting for a set time.
[**UpdateVaultMarketMode**](VaultMarketsApi.md#UpdateVaultMarketMode) | **Put** /v1/vault-markets/mode | Choose your Vault Markets mode: &#x60;mode&#x60; simple or pro. Remembered for you on the web and in the app. A display preference only: it places nothing and moves no money.
[**UpdateVaultMarketOfferAlert**](VaultMarketsApi.md#UpdateVaultMarketOfferAlert) | **Put** /v1/vault-markets/offers/alert | Set (or clear with null) the bid price that should always notify you on a key you hold, and mute or unmute offer notices on it.
[**UpdateVaultMarketPaperLeaderboard**](VaultMarketsApi.md#UpdateVaultMarketPaperLeaderboard) | **Put** /v1/vault-markets/paper/leaderboard | Join or leave the weekly paper leaderboard. Joining needs a display name (2 to 24 letters, numbers, spaces, dots, dashes or underscores), never your real name. Returns &#x60;{ optIn, displayName }&#x60;.
[**UpdateVaultMarketPositionExit**](VaultMarketsApi.md#UpdateVaultMarketPositionExit) | **Put** /v1/vault-markets/positions/{id}/exits | Set a position&#39;s take-profit / stop-loss: a long is closed, a short covered, when the last trade reaches one. Both null clears them.
[**UpdateVaultMarketPriceWidget**](VaultMarketsApi.md#UpdateVaultMarketPriceWidget) | **Patch** /v1/vault-markets/price-widgets/{id} | Change a widget: name, range, layout, theme, affiliate code (null for none) or allowed sites. The market or index it shows never changes: make a new widget for another.
[**UpdateVaultMarketProfile**](VaultMarketsApi.md#UpdateVaultMarketProfile) | **Put** /v1/vault-markets/profile | Turn your public vault profile on or off (&#x60;isPublic&#x60;) and whether followers see your new buys (&#x60;showBuys&#x60;). It shows weights by market, never dollars.
[**UpdateVaultMarketRecurring**](VaultMarketsApi.md#UpdateVaultMarketRecurring) | **Patch** /v1/vault-markets/recurring/{id} | Change a recurring buy&#39;s schedule, caps, cards or auto-lend, or pause / resume it.
[**UpdateVaultMarketSetGoal**](VaultMarketsApi.md#UpdateVaultMarketSetGoal) | **Patch** /v1/vault-markets/set-goals/{id} | Change what a piece must be to count (condition, grade mode, grade). Refused while a basket bid works: cancel it first.
[**UpdateVaultMarketShareCard**](VaultMarketsApi.md#UpdateVaultMarketShareCard) | **Patch** /v1/vault-markets/share-cards/{id} | Show or hide the dollar amounts on one of your share cards (&#x60;showAmounts&#x60;).
[**UpdateVaultMarketSupplyDeposit**](VaultMarketsApi.md#UpdateVaultMarketSupplyDeposit) | **Patch** /v1/vault-markets/supply/deposits/{id} | Change a deposit&#39;s toggles: &#x60;autoListBps&#x60; (list at that share of the mark once stored; null &#x3D; do not) and &#x60;lendWhenIdle&#x60;.
[**UpdateVaultMarketSupplyRoundUp**](VaultMarketsApi.md#UpdateVaultMarketSupplyRoundUp) | **Put** /v1/vault-markets/supply/round-up | Set the round-up: &#x60;shareBps&#x60; (1%-50%) of each payout released to you adds to a budget; when it covers the cheapest vault unit on &#x60;bookKey&#x60;, one unit is bought with your saved card, never above the budget.
[**UpdateVaultMarketTodBeneficiary**](VaultMarketsApi.md#UpdateVaultMarketTodBeneficiary) | **Put** /v1/vault-markets/tod/beneficiaries | Replace your beneficiary list (keep an id to edit one): full name, contact (email or @username), relationship, and share in basis points; shares add up to 10000. Returns &#x60;{ beneficiaries, removed }&#x60;.
[**UpdateVaultMarketTodDesignation**](VaultMarketsApi.md#UpdateVaultMarketTodDesignation) | **Put** /v1/vault-markets/tod/designations | Leave one of your units to one beneficiary (beneficiaryId null clears it; the unit then follows the shares). Returns &#x60;{ vaultItemId, beneficiaryId }&#x60;.
[**UpdateVaultMarketWorldRoomVaultGate**](VaultMarketsApi.md#UpdateVaultMarketWorldRoomVaultGate) | **Put** /v1/vault-markets/world/rooms/{id}/vault-gate | Set the vault-holder door on your own room (its collection rooms follow): visitors must hold minValueCents (cents, at reliable A/B marks) and/or minUnits units in the Crossly vault, mode any (either) or all (both), optionally only one product (scope kind product, ref the catalogue product id), bullion family (bullion, the family slug) or category (category, the slug). Send null to open it again. Returns &#x60;{ gate }&#x60;. Moves no money.
[**UpdateVaultMarketWorldSpaceVaultGate**](VaultMarketsApi.md#UpdateVaultMarketWorldSpaceVaultGate) | **Put** /v1/vault-markets/world/spaces/{slug}/vault-gate | Set the vault-holder door on a Space you host: the same rule as a room (minValueCents at reliable marks and/or minUnits, mode, optional scope). Send null to open it again. Returns &#x60;{ gate }&#x60;. Moves no money.



## CreateVaultLend

> CreateVaultLendResponse CreateVaultLend(ctx, id).Execute()

Turn lending on or off for a unit you own in the vault, optionally with your minimum borrow rate (APR in basis points). Returns `{ lendEnabled, minAprBps }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultLend(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultLend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultLend`: CreateVaultLendResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultLend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultLendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultLendResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketAccount

> StatusResponse CreateVaultMarketAccount(ctx).Execute()

Open your account at the custodian partner (once). Returns `{ status }`; the custodian verifies it before it is active.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketAccount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketAccount`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketAccountRequest struct via the builder pattern


### Return type

**StatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketAdminMark

> CreateVaultMarketAdminMarkResponse CreateVaultMarketAdminMark(ctx).Execute()

Upload your fund administrator's prices (`administrator`, `marks: [{ bookKey, day, cents }]`). Shown beside Crossly's marks and labelled as yours; they never change a Crossly figure.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketAdminMark(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketAdminMark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketAdminMark`: CreateVaultMarketAdminMarkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketAdminMark`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketAdminMarkRequest struct via the builder pattern


### Return type

**CreateVaultMarketAdminMarkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketAdvance

> GetVaultMarketCreditResponse CreateVaultMarketAdvance(ctx).Execute()

Take an advance against a vault unit that is listed for sale: paid into your custodian account, netted from the sale. Free for 60 days, then the pool rate. Needs a verified ID.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketAdvance(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketAdvance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketAdvance`: GetVaultMarketCreditResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketAdvance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketAdvanceRequest struct via the builder pattern


### Return type

**GetVaultMarketCreditResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketAdvanceRelist

> CreateVaultMarketAdvanceRelistResponse CreateVaultMarketAdvanceRelist(ctx, id).Execute()

Relist an advance's unit at a new price (never so low its sale could not repay the advance).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketAdvanceRelist(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketAdvanceRelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketAdvanceRelist`: CreateVaultMarketAdvanceRelistResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketAdvanceRelist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketAdvanceRelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketAdvanceRelistResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketAgentDraftCancel

> CreateVaultMarketAgentDraftCancelResponse CreateVaultMarketAgentDraftCancel(ctx, id).Execute()

Throw a drafted order away (only while it is still a draft). Returns { ok, draft }.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketAgentDraftCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketAgentDraftCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketAgentDraftCancel`: CreateVaultMarketAgentDraftCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketAgentDraftCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketAgentDraftCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketAgentDraftCancelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketAgentWatch

> CreateVaultMarketAgentWatchResponse CreateVaultMarketAgentWatch(ctx).Execute()

Watch a book key: tell me once when its best bid, lowest ask, last trade or market price crosses a price (cents). Leave out direction for \"crosses\" (whichever way it has to move). Never trades. At most 50 live, each for 90 days.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketAgentWatch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketAgentWatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketAgentWatch`: CreateVaultMarketAgentWatchResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketAgentWatch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketAgentWatchRequest struct via the builder pattern


### Return type

**CreateVaultMarketAgentWatchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketAlert

> ListVaultMarketAlertsItem CreateVaultMarketAlert(ctx).Execute()

Alert me when a key's market price goes above or below a price (fires once).

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketAlert(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketAlert`: ListVaultMarketAlertsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketAlert`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketAlertRequest struct via the builder pattern


### Return type

**ListVaultMarketAlertsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketAlertPrefTest

> CreateVaultMarketAlertPrefTestResponse CreateVaultMarketAlertPrefTest(ctx).Execute()

Send yourself one test alert of a `kind` through your own switches. Returns `{ sent, kind, plan }`, where plan says per channel: send, off, quiet (held by quiet hours) or unavailable.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketAlertPrefTest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketAlertPrefTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketAlertPrefTest`: CreateVaultMarketAlertPrefTestResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketAlertPrefTest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketAlertPrefTestRequest struct via the builder pattern


### Return type

**CreateVaultMarketAlertPrefTestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketAuctionLot

> CreateVaultMarketAuctionLotResponse CreateVaultMarketAuctionLot(ctx).Execute()

Consign a stored, authenticated unit to the next weekly or premier event, with an optional opening price and reserve. The unit is held until the close. Returns `{ lotId, eventId, endsAt }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketAuctionLot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketAuctionLot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketAuctionLot`: CreateVaultMarketAuctionLotResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketAuctionLot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketAuctionLotRequest struct via the builder pattern


### Return type

**CreateVaultMarketAuctionLotResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketAuctionLotBid

> CreateVaultMarketAuctionLotBidResponse CreateVaultMarketAuctionLotBid(ctx, id).Execute()

Leave or raise your maximum on a lot; the system bids for you up to it. A bid in the last 2 minutes extends the lot by 2 minutes. The winner's saved card is charged at the close through the order book. ID needed above $10,000.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketAuctionLotBid(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketAuctionLotBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketAuctionLotBid`: CreateVaultMarketAuctionLotBidResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketAuctionLotBid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketAuctionLotBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketAuctionLotBidResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketAuctionLotWithdraw

> OkResponse CreateVaultMarketAuctionLotWithdraw(ctx, id).Execute()

Withdraw your lot before any bid; the unit goes back on your shelf. Returns `{ ok }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketAuctionLotWithdraw(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketAuctionLotWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketAuctionLotWithdraw`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketAuctionLotWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketAuctionLotWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketBlock

> ListVaultMarketBlocksItem CreateVaultMarketBlock(ctx).Execute()

Place a block buy for an entity: one order booked to the first sub-account in `plan` (it pays from its custodian account), split across every sub-account in the plan at the average price, fees included, once it is done and the custodian confirmed each payment. Each share's cash moves between the entity's own custodian accounts; its units move only after that settles. Every sub-account must be one you may trade, and its share must fit its limits.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketBlock(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketBlock`: ListVaultMarketBlocksItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketBlock`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketBlockRequest struct via the builder pattern


### Return type

**ListVaultMarketBlocksItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketBlockAllocate

> CreateVaultMarketBlockAllocateResponse CreateVaultMarketBlockAllocate(ctx, orderId).Execute()

Allocate a done block now (the keeper does it too). Returns `{ outcome, block }`: allocated, waiting (a payment or transfer is not settled yet), nothing_filled or none.

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
    orderId := "orderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketBlockAllocate(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketBlockAllocate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketBlockAllocate`: CreateVaultMarketBlockAllocateResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketBlockAllocate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketBlockAllocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketBlockAllocateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketBullionFamilyPeg

> ListVaultMarketBullionFamilyPegsItem CreateVaultMarketBullionFamilyPeg(ctx).Execute()

Set a family-wide spot-pegged quote: one \"any 1 oz gold coin\" quote resting on each product book of a bullion family at the chosen condition (DS, VNDS or Used). When any leg fills it comes off the others and goes back on only for what is left; two fills at once are settled one at a time against what is left. A bid pays from your custodian account. Replaces your live family quote on the same family, condition and side. ID required.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketBullionFamilyPeg(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketBullionFamilyPeg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketBullionFamilyPeg`: ListVaultMarketBullionFamilyPegsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketBullionFamilyPeg`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketBullionFamilyPegRequest struct via the builder pattern


### Return type

**ListVaultMarketBullionFamilyPegsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketBullionPeg

> ListVaultMarketBullionPegsItem CreateVaultMarketBullionPeg(ctx).Execute()

Set a spot-pegged maker quote on a bullion book key: \"spot plus or minus N cents\" or \"plus or minus N basis points\", with an optional floor and cap. It rests as a pro limit order re-priced every spot tick and is pulled while spot is stale. Replaces your live quote on the same key and side. ID required.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketBullionPeg(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketBullionPeg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketBullionPeg`: ListVaultMarketBullionPegsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketBullionPeg`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketBullionPegRequest struct via the builder pattern


### Return type

**ListVaultMarketBullionPegsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketCollectionWorthWatch

> CreateVaultMarketCollectionWorthWatchResponse CreateVaultMarketCollectionWorthWatch(ctx).Execute()

Watch every collection item that sits on a vault book key (up to the watchlist cap). Returns `{ watched, skipped }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketCollectionWorthWatch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketCollectionWorthWatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketCollectionWorthWatch`: CreateVaultMarketCollectionWorthWatchResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketCollectionWorthWatch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketCollectionWorthWatchRequest struct via the builder pattern


### Return type

**CreateVaultMarketCollectionWorthWatchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketCreatorAttribution

> CreateVaultMarketCreatorAttributionResponse CreateVaultMarketCreatorAttribution(ctx).Execute()

Use a creator code: the creator earns a capped share of the fee Crossly keeps on your Vault Markets trades for a limited window (paid by Crossly, never by you). Only for people new to Vault Markets, once, and never your own code.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketCreatorAttribution(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketCreatorAttribution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketCreatorAttribution`: CreateVaultMarketCreatorAttributionResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketCreatorAttribution`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketCreatorAttributionRequest struct via the builder pattern


### Return type

**CreateVaultMarketCreatorAttributionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketCreatorCohostInviteRespond

> CreateVaultMarketCreatorCohostInviteRespondResponse CreateVaultMarketCreatorCohostInviteRespond(ctx, id).Execute()

For creators: accept or decline an invite to co-host a Market Open launch show (`accept` true or false). Accepting needs your vault room verified and adds you to the show as a co-host; your room and the Commons then say you co-host it, with the creator-programme disclosure when it applies.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketCreatorCohostInviteRespond(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketCreatorCohostInviteRespond``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketCreatorCohostInviteRespond`: CreateVaultMarketCreatorCohostInviteRespondResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketCreatorCohostInviteRespond`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketCreatorCohostInviteRespondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketCreatorCohostInviteRespondResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketCreditRepay

> CreateVaultMarketCreditRepayResponse CreateVaultMarketCreditRepay(ctx, id).Execute()

Repay from your custodian account, interest first. Omit `amountCents` to repay everything; pass `requestId` to make a retry safe. Repaid in full, the items are free.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketCreditRepay(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketCreditRepay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketCreditRepay`: CreateVaultMarketCreditRepayResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketCreditRepay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketCreditRepayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketCreditRepayResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketDcmLink

> CreateVaultMarketDcmLinkResponse CreateVaultMarketDcmLink(ctx).Execute()

Link a partner exchange account to your account (or to a firm sub-account you manage, `subaccountId`) so Crossly can show its positions: `accountRef` (the account id there) and `code` (the one-time link code it gave you). The partner must confirm the code; the code is never stored. Read-only: Crossly never trades on it.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketDcmLink(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketDcmLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketDcmLink`: CreateVaultMarketDcmLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketDcmLink`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketDcmLinkRequest struct via the builder pattern


### Return type

**CreateVaultMarketDcmLinkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketDcmUnlink

> CreateVaultMarketDcmUnlinkResponse CreateVaultMarketDcmUnlink(ctx).Execute()

Unlink the partner exchange account from your account (or the firm sub-account `subaccountId`): Crossly stops reading it and its hedges leave portfolio margin.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketDcmUnlink(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketDcmUnlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketDcmUnlink`: CreateVaultMarketDcmUnlinkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketDcmUnlink`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketDcmUnlinkRequest struct via the builder pattern


### Return type

**CreateVaultMarketDcmUnlinkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketDealerApply

> CreateVaultMarketDealerApplyResponse CreateVaultMarketDealerApply(ctx).Execute()

Apply to be a dealer (ID check first). Ops review; approved dealers get raised limits and their bids show as dealer bids.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketDealerApply(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketDealerApply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketDealerApply`: CreateVaultMarketDealerApplyResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketDealerApply`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketDealerApplyRequest struct via the builder pattern


### Return type

**CreateVaultMarketDealerApplyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketDealerBid

> CreateVaultMarketProOrderResponse CreateVaultMarketDealerBid(ctx).Execute()

Place a standing dealer bid: a resting GTC/GTD limit buy, flagged as a dealer bid. Approved dealers only.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketDealerBid(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketDealerBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketDealerBid`: CreateVaultMarketProOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketDealerBid`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketDealerBidRequest struct via the builder pattern


### Return type

**CreateVaultMarketProOrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketDealerPool

> ListVaultMarketDealerPoolsItem CreateVaultMarketDealerPool(ctx).Execute()

Open a curve pool on a key (approved dealers, ID checked): linear (delta in cents) or exponential (delta in bps), spot, fee, levels per side, units it may sell, and a buy budget from your custodian account. Its ladder goes up at once.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketDealerPool(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketDealerPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketDealerPool`: ListVaultMarketDealerPoolsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketDealerPool`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketDealerPoolRequest struct via the builder pattern


### Return type

**ListVaultMarketDealerPoolsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketDealerPoolClose

> ListVaultMarketDealerPoolsItem CreateVaultMarketDealerPoolClose(ctx, id).Execute()

Close a curve pool: its legs come off the book; what filled stays filled.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketDealerPoolClose(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketDealerPoolClose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketDealerPoolClose`: ListVaultMarketDealerPoolsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketDealerPoolClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketDealerPoolCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketDealerPoolsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketDealerPoolPreview

> CreateVaultMarketDealerPoolPreviewResponse CreateVaultMarketDealerPoolPreview(ctx).Execute()

The bids and asks a set of curve settings would show, and why each side stops where it does. Places nothing.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketDealerPoolPreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketDealerPoolPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketDealerPoolPreview`: CreateVaultMarketDealerPoolPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketDealerPoolPreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketDealerPoolPreviewRequest struct via the builder pattern


### Return type

**CreateVaultMarketDealerPoolPreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketDeposit

> CreateVaultMarketDepositResponse CreateVaultMarketDeposit(ctx).Execute()

Open a ship-in deposit for an ungraded catalogue item (`skuId`, `variantId`, `condition` DS / VNDS / Used). It is inbound until the vault receives and checks it, then stored in your name. Returns `{ vaultItem: { id, status } }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketDeposit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketDeposit`: CreateVaultMarketDepositResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketDeposit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketDepositRequest struct via the builder pattern


### Return type

**CreateVaultMarketDepositResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketDonation

> CreateVaultMarketDonationResponse CreateVaultMarketDonation(ctx).Execute()

Give a unit to Crossly Cares (once the charity can receive; otherwise its gate message). Returns `{ id, markCents }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketDonation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketDonation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketDonation`: CreateVaultMarketDonationResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketDonation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketDonationRequest struct via the builder pattern


### Return type

**CreateVaultMarketDonationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketEntity

> CreateVaultMarketEntityResponse CreateVaultMarketEntity(ctx).Execute()

Open an entity account: you become its admin and the custodian partner starts KYB (its own flow; Crossly keeps only its reference). Needs a verified ID and the custodian (503 names what is missing).

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketEntity(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketEntity`: CreateVaultMarketEntityResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketEntity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketEntityRequest struct via the builder pattern


### Return type

**CreateVaultMarketEntityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketEntityEcp

> GetVaultMarketEntityByIdResponse CreateVaultMarketEntityEcp(ctx, id).Execute()

Admin: record the entity's attestation that it is an eligible contract participant. Raising its caps is still an ops decision.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketEntityEcp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketEntityEcp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketEntityEcp`: GetVaultMarketEntityByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketEntityEcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketEntityEcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketEntityByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketEntityHalt

> CreateVaultMarketEntityHaltResponse CreateVaultMarketEntityHalt(ctx, id).Execute()

Admin or risk: the firm kill switch. Halts the entity (no `subaccountId`) or one sub-account: open orders on it are cancelled and no new ones are placed until cleared.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketEntityHalt(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketEntityHalt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketEntityHalt`: CreateVaultMarketEntityHaltResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketEntityHalt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketEntityHaltRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketEntityHaltResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketEntityHaltClear

> GetVaultMarketEntityByIdResponse CreateVaultMarketEntityHaltClear(ctx, id, haltId).Execute()

Admin or risk: lift a firm halt.

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
    id := "id_example" // string | 
    haltId := "haltId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketEntityHaltClear(context.Background(), id, haltId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketEntityHaltClear``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketEntityHaltClear`: GetVaultMarketEntityByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketEntityHaltClear`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**haltId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketEntityHaltClearRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetVaultMarketEntityByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketEntityInvitation

> GetVaultMarketEntityResponse CreateVaultMarketEntityInvitation(ctx, id).Execute()

Accept or decline an invitation (`accept`). Trading, admin and risk roles need a verified ID to accept.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketEntityInvitation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketEntityInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketEntityInvitation`: GetVaultMarketEntityResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketEntityInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketEntityInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketEntityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketEntityKyb

> CreateVaultMarketEntityKybResponse CreateVaultMarketEntityKyb(ctx, id).Execute()

Refresh the entity's KYB status from the custodian (and its onboarding link while KYB is open).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketEntityKyb(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketEntityKyb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketEntityKyb`: CreateVaultMarketEntityKybResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketEntityKyb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketEntityKybRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketEntityKybResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketEntityMember

> GetVaultMarketEntityByIdResponse CreateVaultMarketEntityMember(ctx, id).Execute()

Admin: invite a Crossly user by email as admin, trader, risk, ops or viewer (a viewer can never place an order), with optional limits and sub-accounts.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketEntityMember(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketEntityMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketEntityMember`: GetVaultMarketEntityByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketEntityMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketEntityMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketEntityByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketEntitySubaccount

> GetVaultMarketEntityByIdResponse CreateVaultMarketEntitySubaccount(ctx, id).Execute()

Admin: open a sub-account once the custodian approved the entity; its limits are held to the entity's caps.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketEntitySubaccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketEntitySubaccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketEntitySubaccount`: GetVaultMarketEntityByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketEntitySubaccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketEntitySubaccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketEntityByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketEntitySubaccountClose

> GetVaultMarketEntityByIdResponse CreateVaultMarketEntitySubaccountClose(ctx, id, subId).Execute()

Admin: close a sub-account that holds nothing and has no open orders.

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
    id := "id_example" // string | 
    subId := "subId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketEntitySubaccountClose(context.Background(), id, subId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketEntitySubaccountClose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketEntitySubaccountClose`: GetVaultMarketEntityByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketEntitySubaccountClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**subId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketEntitySubaccountCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetVaultMarketEntityByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketFollowCopyWatchAlertRead

> CreateVaultMarketFollowCopyWatchAlertReadResponse CreateVaultMarketFollowCopyWatchAlertRead(ctx).Execute()

Mark every copy-watch alert read. Returns `{ marked }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketFollowCopyWatchAlertRead(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketFollowCopyWatchAlertRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketFollowCopyWatchAlertRead`: CreateVaultMarketFollowCopyWatchAlertReadResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketFollowCopyWatchAlertRead`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketFollowCopyWatchAlertReadRequest struct via the builder pattern


### Return type

**CreateVaultMarketFollowCopyWatchAlertReadResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGift

> CreateVaultMarketGiftResponse CreateVaultMarketGift(ctx).Execute()

Give a unit in your vault to another Crossly user (by @username or email). It stays on its shelf; they have 7 days to accept. ID needed above $1,000 at the mark. Returns `{ id, expiresAt }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGift(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGift``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGift`: CreateVaultMarketGiftResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGift`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGiftRequest struct via the builder pattern


### Return type

**CreateVaultMarketGiftResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGiftAccept

> StatusResponse CreateVaultMarketGiftAccept(ctx, id).Execute()

Accept a gift: the unit becomes yours in place. Returns `{ status }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGiftAccept(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGiftAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGiftAccept`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGiftAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGiftAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGiftCancel

> StatusResponse CreateVaultMarketGiftCancel(ctx, id).Execute()

Cancel a gift you sent before it is accepted. Returns `{ status }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGiftCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGiftCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGiftCancel`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGiftCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGiftCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGiftDecline

> StatusResponse CreateVaultMarketGiftDecline(ctx, id).Execute()

Decline a gift: the unit stays with the giver. Returns `{ status }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGiftDecline(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGiftDecline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGiftDecline`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGiftDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGiftDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGradeToVault

> ListVaultMarketGradeToVaultItem CreateVaultMarketGradeToVault(ctx).Execute()

Start Grade to Vault on one item (`source` inventory or vault_unit). With `runId` the grading slot is booked at once (free; your saved card is charged only if the run locks). Nothing is listed until you confirm a price after the grade.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGradeToVault(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGradeToVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGradeToVault`: ListVaultMarketGradeToVaultItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGradeToVault`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradeToVaultRequest struct via the builder pattern


### Return type

**ListVaultMarketGradeToVaultItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGradeToVaultBook

> ListVaultMarketGradeToVaultItem CreateVaultMarketGradeToVaultBook(ctx, id).Execute()

Book (or re-book) the grading slot of a flow still at \"picked\". Answers 503 naming GRADING_HUB_ADDRESS when runs are not open.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGradeToVaultBook(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGradeToVaultBook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGradeToVaultBook`: ListVaultMarketGradeToVaultItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGradeToVaultBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradeToVaultBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketGradeToVaultItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGradeToVaultCancel

> ListVaultMarketGradeToVaultItem CreateVaultMarketGradeToVaultCancel(ctx, id).Execute()

Stop the flow before the card goes to the grader (still picked, or its slot only reserved: nothing was charged).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGradeToVaultCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGradeToVaultCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGradeToVaultCancel`: ListVaultMarketGradeToVaultItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGradeToVaultCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradeToVaultCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketGradeToVaultItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGradeToVaultKeep

> ListVaultMarketGradeToVaultItem CreateVaultMarketGradeToVaultKeep(ctx, id).Execute()

Keep the graded unit in your vault without listing it. The flow ends; list it any time from the vault.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGradeToVaultKeep(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGradeToVaultKeep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGradeToVaultKeep`: ListVaultMarketGradeToVaultItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGradeToVaultKeep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradeToVaultKeepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketGradeToVaultItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGradeToVaultList

> ListVaultMarketGradeToVaultItem CreateVaultMarketGradeToVaultList(ctx, id).Execute()

Confirm the price for the grade the unit came back with (`confirm: true`, `gradeKey` must match the unit). Pre-trade risk, the price band and the ID line apply, then one shelf ask is placed on the Crossly order book.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGradeToVaultList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGradeToVaultList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGradeToVaultList`: ListVaultMarketGradeToVaultItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGradeToVaultList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradeToVaultListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketGradeToVaultItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGradeToVaultShipped

> ListVaultMarketGradeToVaultItem CreateVaultMarketGradeToVaultShipped(ctx, id).Execute()

You mailed the card to the grading hub (an inventory card, once its run locked).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGradeToVaultShipped(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGradeToVaultShipped``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGradeToVaultShipped`: ListVaultMarketGradeToVaultItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGradeToVaultShipped`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradeToVaultShippedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketGradeToVaultItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGrading

> CreateVaultMarketGradingResponse CreateVaultMarketGrading(ctx).Execute()

Send a raw vault unit on the grading hub's shelf to a Grading Run (`runId`). It is held until it comes back, graded, to the same vault item. The slot is free; your saved card is charged only if the run locks.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGrading(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGrading``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGrading`: CreateVaultMarketGradingResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGrading`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradingRequest struct via the builder pattern


### Return type

**CreateVaultMarketGradingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGradingPull

> StatusResponse CreateVaultMarketGradingPull(ctx, submissionId).Execute()

Pull a vault unit out of its run before the run locks (nothing was charged). It goes back on the shelf.

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
    submissionId := "submissionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketGradingPull(context.Background(), submissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketGradingPull``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGradingPull`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketGradingPull`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradingPullRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketHallLoan

> CreateVaultMarketHallLoanResponse CreateVaultMarketHallLoan(ctx).Execute()

Lend a unit in your Crossly vault to the Hall of Grails: { vaultItemId, credit: 'name' | 'anonymous' }. A standing loan of display rights only: the unit goes on a plinth, under a plate naming you (or \"a collector\"), in every weekly exhibition whose category it fits. It never moves and stays yours to list, sell or withdraw. Earns reputation per exhibition shown; no cash. Resumes a loan you opted out of while it still shows.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketHallLoan(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketHallLoan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketHallLoan`: CreateVaultMarketHallLoanResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketHallLoan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketHallLoanRequest struct via the builder pattern


### Return type

**CreateVaultMarketHallLoanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketHallLoanCredit

> CreateVaultMarketHallLoanCreditResponse CreateVaultMarketHallLoanCredit(ctx, id).Execute()

Change how the plate on a lent unit's plinth names you, from now: { credit: 'name' } (your display name) or { credit: 'anonymous' } (\"a collector\").

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketHallLoanCredit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketHallLoanCredit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketHallLoanCredit`: CreateVaultMarketHallLoanCreditResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketHallLoanCredit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketHallLoanCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketHallLoanCreditResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketHallLoanOptOut

> CreateVaultMarketHallLoanOptOutResponse CreateVaultMarketHallLoanOptOut(ctx, id).Execute()

Opt a Hall of Grails loan out, at any time. A unit on a plinth this week finishes this exhibition (status ending, showingUntil); otherwise the loan ends now (status ended).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketHallLoanOptOut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketHallLoanOptOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketHallLoanOptOut`: CreateVaultMarketHallLoanOptOutResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketHallLoanOptOut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketHallLoanOptOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketHallLoanOptOutResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketIdentityStart

> CreateVaultMarketIdentityStartResponse CreateVaultMarketIdentityStart(ctx).Execute()

Start (or resume) the one-time ID check the money features and larger orders need. Returns `{ url }` to open, or `{ alreadyVerified: true }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketIdentityStart(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketIdentityStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketIdentityStart`: CreateVaultMarketIdentityStartResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketIdentityStart`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketIdentityStartRequest struct via the builder pattern


### Return type

**CreateVaultMarketIdentityStartResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketIndex

> GetVaultMarketIndexResponse CreateVaultMarketIndex(ctx).Execute()

Make an index of your own: 1 to 50 book keys, weighted by mark or equally. Up to 10 per account.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketIndex(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketIndex`: GetVaultMarketIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketIndex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketIndexRequest struct via the builder pattern


### Return type

**GetVaultMarketIndexResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketIndexBuy

> CreateVaultMarketIndexBuyResponse CreateVaultMarketIndexBuy(ctx, slug).Execute()

Buy the basket: one limit buy per key in the index for 1-5 whole units each, at most `maxBpsOverMark` over each key's mark, IOC or GTC. Every unit is a real item that lands in your vault. Idempotent on `clientBasketId`; returns a result per key.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketIndexBuy(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketIndexBuy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketIndexBuy`: CreateVaultMarketIndexBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketIndexBuy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketIndexBuyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketIndexBuyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketIndexComplaint

> CreateVaultMarketIndexComplaintResponse CreateVaultMarketIndexComplaint(ctx).Execute()

File a complaint about an index or the method (`slug` optional). Every complaint is answered. Returns `{ id, status, createdAt }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketIndexComplaint(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketIndexComplaint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketIndexComplaint`: CreateVaultMarketIndexComplaintResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketIndexComplaint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketIndexComplaintRequest struct via the builder pattern


### Return type

**CreateVaultMarketIndexComplaintResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketIndexConsultationComment

> OkResponse CreateVaultMarketIndexConsultationComment(ctx, id).Execute()

Comment on an open consultation (5 to 5,000 characters; up to 5 per person). Comments are published without your name. Returns `{ ok }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketIndexConsultationComment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketIndexConsultationComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketIndexConsultationComment`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketIndexConsultationComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketIndexConsultationCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketInstantExit

> CreateVaultMarketInstantExitResponse CreateVaultMarketInstantExit(ctx).Execute()

Sell a unit now at the quoted price: a sell IOC at `priceCents` through the pro engine. Refused, with the new quote, if the best bid dropped below it. Crossly never buys or funds it.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketInstantExit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketInstantExit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketInstantExit`: CreateVaultMarketInstantExitResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketInstantExit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketInstantExitRequest struct via the builder pattern


### Return type

**CreateVaultMarketInstantExitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketJoint

> CreateVaultMarketJointResponse CreateVaultMarketJoint(ctx).Execute()

Open a joint vault (ID first): you hold full rights and accept the current joint terms. `invite` sends invitations (@username or email, each with view, trade or full rights). Returns `{ id, status, invited }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketJoint(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketJoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketJoint`: CreateVaultMarketJointResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketJoint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketJointRequest struct via the builder pattern


### Return type

**CreateVaultMarketJointResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketJointClose

> CreateVaultMarketJointCloseResponse CreateVaultMarketJointClose(ctx, id).Execute()

Close the joint vault (full rights): open joint orders are cancelled and every unit goes back to its owner. Returns `{ status, unitsReturned }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketJointClose(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketJointClose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketJointClose`: CreateVaultMarketJointCloseResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketJointClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketJointCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketJointCloseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketJointInvite

> CreateVaultMarketJointInviteResponse CreateVaultMarketJointInvite(ctx, id).Execute()

Invite a Crossly user (full rights only). They have 14 days and show ID to join. Returns `{ userId, rights, expiresAt }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketJointInvite(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketJointInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketJointInvite`: CreateVaultMarketJointInviteResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketJointInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketJointInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketJointInviteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketJointMemberRemove

> CreateVaultMarketJointMemberRemoveResponse CreateVaultMarketJointMemberRemove(ctx, id, userId).Execute()

Remove a member (full rights), or leave when userId is you. Their units go back to their own account and their joint orders are cancelled. Returns `{ userId, status, unitsReturned, jointStatus }`.

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
    id := "id_example" // string | 
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketJointMemberRemove(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketJointMemberRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketJointMemberRemove`: CreateVaultMarketJointMemberRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketJointMemberRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketJointMemberRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateVaultMarketJointMemberRemoveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketJointMemberRight

> CreateVaultMarketJointMemberRightResponse CreateVaultMarketJointMemberRight(ctx, id, userId).Execute()

Set a member rights to view, trade or full (full rights only; the vault always keeps a full member). Returns `{ userId, rights }`.

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
    id := "id_example" // string | 
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketJointMemberRight(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketJointMemberRight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketJointMemberRight`: CreateVaultMarketJointMemberRightResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketJointMemberRight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketJointMemberRightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateVaultMarketJointMemberRightResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketJointRespond

> CreateVaultMarketJointRespondResponse CreateVaultMarketJointRespond(ctx, id).Execute()

Accept (ID first; accepts the current joint terms) or decline an invitation. Returns `{ status }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketJointRespond(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketJointRespond``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketJointRespond`: CreateVaultMarketJointRespondResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketJointRespond`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketJointRespondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketJointRespondResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketJointTerm

> TermsVersionResponse CreateVaultMarketJointTerm(ctx, id).Execute()

Accept the current joint terms (needed to trade after the terms version changes). Returns `{ termsVersion }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketJointTerm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketJointTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketJointTerm`: TermsVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketJointTerm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketJointTermRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**TermsVersionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketJointTrade

> CreateVaultMarketJointTradeResponse CreateVaultMarketJointTrade(ctx, id).Execute()

Place a limit order for the joint vault (trade or full rights, ID shown). A buy (bookKey) is placed on you, paid from your own Vault Markets account, and what it buys is booked to the joint vault. A sell (vaultItemId of a unit in it) is placed on the unit owner of record, who gets the proceeds; pre-trade checks run against you too. Returns `{ tradeId, side, bookKey, ownerUserId, order }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketJointTrade(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketJointTrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketJointTrade`: CreateVaultMarketJointTradeResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketJointTrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketJointTradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketJointTradeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketJointTradeCancel

> CreateVaultMarketJointTradeCancelResponse CreateVaultMarketJointTradeCancel(ctx, id, tradeId).Execute()

Cancel a working joint order (trade or full rights); what filled stays filled. Returns `{ tradeId, order }`.

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
    id := "id_example" // string | 
    tradeId := "tradeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketJointTradeCancel(context.Background(), id, tradeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketJointTradeCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketJointTradeCancel`: CreateVaultMarketJointTradeCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketJointTradeCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**tradeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketJointTradeCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateVaultMarketJointTradeCancelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketJointUnit

> CreateVaultMarketJointUnitResponse CreateVaultMarketJointUnit(ctx, id).Execute()

Put units you own into the joint vault (each stored and free). They keep you as owner of record and stay on the shelf. Returns `{ added, vaultItemIds }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketJointUnit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketJointUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketJointUnit`: CreateVaultMarketJointUnitResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketJointUnit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketJointUnitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketJointUnitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketJointUnitRemove

> CreateVaultMarketJointUnitRemoveResponse CreateVaultMarketJointUnitRemove(ctx, id).Execute()

Take units out of the joint vault, back to their owner own account (yours at any rights; anyone else with full rights). Returns `{ removed, vaultItemIds }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketJointUnitRemove(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketJointUnitRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketJointUnitRemove`: CreateVaultMarketJointUnitRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketJointUnitRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketJointUnitRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketJointUnitRemoveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketLargeTraderReportFile

> ListVaultMarketLargeTraderReportsItem CreateVaultMarketLargeTraderReportFile(ctx, id).Execute()

File (or amend, or refile) a large-trader report: why you hold it (purpose), who you act for, other accounts you control, whether you plan to add, and a statement of 20+ characters. Returns the report.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketLargeTraderReportFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketLargeTraderReportFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketLargeTraderReportFile`: ListVaultMarketLargeTraderReportsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketLargeTraderReportFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketLargeTraderReportFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketLargeTraderReportsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketLearnQuiz

> CreateVaultMarketLearnQuizResponse CreateVaultMarketLearnQuiz(ctx, concept).Execute()

Answer a concept's quiz (`answers`: the chosen choice index per question). Graded on the server; 5 tries an hour.

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
    concept := "concept_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketLearnQuiz(context.Background(), concept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketLearnQuiz``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketLearnQuiz`: CreateVaultMarketLearnQuizResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketLearnQuiz`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**concept** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketLearnQuizRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketLearnQuizResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketLiveStreamFeaturedLot

> OkResponse CreateVaultMarketLiveStreamFeaturedLot(ctx, id).Execute()

Feature an open lot on your live show (or `lotId: null` to take it off). Viewers hear LIVE_PIN_CHANGED. Returns `{ ok }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketLiveStreamFeaturedLot(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketLiveStreamFeaturedLot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketLiveStreamFeaturedLot`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketLiveStreamFeaturedLot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketLiveStreamFeaturedLotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketLoan

> GetVaultMarketCreditResponse CreateVaultMarketLoan(ctx).Execute()

Borrow against stored vault units: they are pledged (no listing, lending or withdrawal) and the market pool lends `amountCents` into your custodian account. Needs a verified ID.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketLoan(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketLoan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketLoan`: GetVaultMarketCreditResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketLoan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketLoanRequest struct via the builder pattern


### Return type

**GetVaultMarketCreditResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketLoanRecall

> CreateVaultMarketLoanRecallResponse CreateVaultMarketLoanRecall(ctx, id).Execute()

Recall a lent unit: the short has 3 days to return one. Returns `{ recallDueAt }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketLoanRecall(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketLoanRecall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketLoanRecall`: CreateVaultMarketLoanRecallResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketLoanRecall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketLoanRecallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketLoanRecallResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketLoanUnit

> GetVaultMarketCreditResponse CreateVaultMarketLoanUnit(ctx, id).Execute()

Pledge more stored units to a loan (answers a loan call, or makes room).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketLoanUnit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketLoanUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketLoanUnit`: GetVaultMarketCreditResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketLoanUnit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketLoanUnitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketCreditResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketLoanUnitRelease

> GetVaultMarketCreditResponse CreateVaultMarketLoanUnitRelease(ctx, id, vaultItemId).Execute()

Free one pledged unit, if what is owed still fits inside the other units' limit.

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
    id := "id_example" // string | 
    vaultItemId := "vaultItemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketLoanUnitRelease(context.Background(), id, vaultItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketLoanUnitRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketLoanUnitRelease`: GetVaultMarketCreditResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketLoanUnitRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**vaultItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketLoanUnitReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetVaultMarketCreditResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketLoanUnitSell

> CreateVaultMarketLoanUnitSellResponse CreateVaultMarketLoanUnitSell(ctx, id, vaultItemId).Execute()

Sell a pledged unit to repay its loan: it is listed at `priceCents` (at least the sale floor) and stays pledged until it sells; the proceeds repay the loan first and the rest is yours. Calling it again re-prices it.

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
    id := "id_example" // string | 
    vaultItemId := "vaultItemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketLoanUnitSell(context.Background(), id, vaultItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketLoanUnitSell``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketLoanUnitSell`: CreateVaultMarketLoanUnitSellResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketLoanUnitSell`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**vaultItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketLoanUnitSellRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateVaultMarketLoanUnitSellResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketLocate

> ListVaultMarketLocateMineItem CreateVaultMarketLocate(ctx).Execute()

Hold up to `qty` lendable units of `bookKey` for 15 minutes so a short can open against them (`termDays` 30/90 for a term loan). Needs lending (VAULT_LENDING_TERMS_VERSION) and the ID check. Returns the locate with how many units it holds.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketLocate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketLocate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketLocate`: ListVaultMarketLocateMineItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketLocate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketLocateRequest struct via the builder pattern


### Return type

**ListVaultMarketLocateMineItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketLong

> CreateVaultMarketLongResponse CreateVaultMarketLong(ctx).Execute()

Buy `units` units (default 1; at most 25/10/3/1 by liquidity tier A/B/C/D) into your vault with leverage (up to 2x) from the market staking pool, paying up to `maxCents` each. A partial fill opens with the units bought and the rest is cancelled. Returns `{ positionId, matched, units }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketLong(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketLong``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketLong`: CreateVaultMarketLongResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketLong`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketLongRequest struct via the builder pattern


### Return type

**CreateVaultMarketLongResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketMidpoint

> ListVaultMarketMidpointItem CreateVaultMarketMidpoint(ctx).Execute()

A hidden block order that trades only at the firm-book midpoint, with another midpoint order of at least your minimum size.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketMidpoint(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketMidpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketMidpoint`: ListVaultMarketMidpointItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketMidpoint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketMidpointRequest struct via the builder pattern


### Return type

**ListVaultMarketMidpointItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketMidpointCancel

> ListVaultMarketMidpointItem CreateVaultMarketMidpointCancel(ctx, id).Execute()

Cancel a midpoint order.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketMidpointCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketMidpointCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketMidpointCancel`: ListVaultMarketMidpointItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketMidpointCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketMidpointCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketMidpointItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketMmDmm

> CreateVaultMarketMmLpResponse CreateVaultMarketMmDmm(ctx).Execute()

Apply to be the designated market maker on a tier B or C key. Ops approves; one per key.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketMmDmm(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketMmDmm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketMmDmm`: CreateVaultMarketMmLpResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketMmDmm`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketMmDmmRequest struct via the builder pattern


### Return type

**CreateVaultMarketMmLpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketMmLeave

> CreateVaultMarketMmLpResponse CreateVaultMarketMmLeave(ctx, id).Execute()

Leave a market-maker programme.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketMmLeave(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketMmLeave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketMmLeave`: CreateVaultMarketMmLpResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketMmLeave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketMmLeaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketMmLpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketMmLp

> CreateVaultMarketMmLpResponse CreateVaultMarketMmLp(ctx).Execute()

Join a key as a Registered Liquidity Provider (ID needed). Meeting coverage earns the LP rates.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketMmLp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketMmLp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketMmLp`: CreateVaultMarketMmLpResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketMmLp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketMmLpRequest struct via the builder pattern


### Return type

**CreateVaultMarketMmLpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketMmMassCancel

> CreateVaultMarketMmMassCancelResponse CreateVaultMarketMmMassCancel(ctx).Execute()

Cancel every open order (one key, or only quotes) and every resting leg in one statement: nothing fills between the first cancel and the last.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketMmMassCancel(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketMmMassCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketMmMassCancel`: CreateVaultMarketMmMassCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketMmMassCancel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketMmMassCancelRequest struct via the builder pattern


### Return type

**CreateVaultMarketMmMassCancelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketMmMmpReset

> GetVaultMarketMmMmpResponse CreateVaultMarketMmMmpReset(ctx).Execute()

End a market-maker protection freeze early.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketMmMmpReset(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketMmMmpReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketMmMmpReset`: GetVaultMarketMmMmpResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketMmMmpReset`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketMmMmpResetRequest struct via the builder pattern


### Return type

**GetVaultMarketMmMmpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketOpeningWaitlist

> CreateVaultMarketOpeningWaitlistResponse CreateVaultMarketOpeningWaitlist(ctx, slug).Execute()

Join a key's waitlist, optionally with a friend's invite code (`ref`). Inviting moves you up the queue; nobody is paid. Returns your place and your own code.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketOpeningWaitlist(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketOpeningWaitlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketOpeningWaitlist`: CreateVaultMarketOpeningWaitlistResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketOpeningWaitlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketOpeningWaitlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketOpeningWaitlistResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketP2pLoanCall

> GetVaultMarketP2pLoanResponse CreateVaultMarketP2pLoanCall(ctx, id).Execute()

Lender: call the loan (once it is a day old). A refinancing auction runs for about 30 hours, the rate a new lender gets climbing to the ceiling; repaid or refinanced, you are paid in full; neither by the end, the unit moves to your vault.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketP2pLoanCall(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketP2pLoanCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketP2pLoanCall`: GetVaultMarketP2pLoanResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketP2pLoanCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketP2pLoanCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketP2pLoanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketP2pLoanCancelCall

> GetVaultMarketP2pLoanResponse CreateVaultMarketP2pLoanCancelCall(ctx, id).Execute()

Lender: take back your call before any other lender has taken the loan over. The refinancing auction stops and the loan carries on as before on its terms.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketP2pLoanCancelCall(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketP2pLoanCancelCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketP2pLoanCancelCall`: GetVaultMarketP2pLoanResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketP2pLoanCancelCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketP2pLoanCancelCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketP2pLoanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketP2pLoanRefiAuction

> GetVaultMarketP2pLoanResponse CreateVaultMarketP2pLoanRefiAuction(ctx, id).Execute()

Borrower: look for a cheaper lender. A refinancing auction runs for `hours` (1-72, default 24); the rate a new lender gets climbs from `startAprBps` (default half the loan rate) to just under the loan rate now. The first lender to take it pays your lender off and the loan carries on at that lower rate. Nobody by the end: nothing changes.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketP2pLoanRefiAuction(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketP2pLoanRefiAuction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketP2pLoanRefiAuction`: GetVaultMarketP2pLoanResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketP2pLoanRefiAuction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketP2pLoanRefiAuctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketP2pLoanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketP2pLoanRefiAuctionCancel

> GetVaultMarketP2pLoanResponse CreateVaultMarketP2pLoanRefiAuctionCancel(ctx, id).Execute()

Borrower: stop looking for a cheaper lender before anyone has taken the loan over. The loan carries on unchanged.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketP2pLoanRefiAuctionCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketP2pLoanRefiAuctionCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketP2pLoanRefiAuctionCancel`: GetVaultMarketP2pLoanResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketP2pLoanRefiAuctionCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketP2pLoanRefiAuctionCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketP2pLoanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketP2pLoanRefinance

> GetVaultMarketP2pLoanResponse CreateVaultMarketP2pLoanRefinance(ctx, id).Execute()

Take over a loan in a refinancing auction (a lender call or a borrower search for a cheaper lender): you pay the old lender everything owed now (refused above `maxPayoffCents`), which becomes your loan to the same borrower at the auction rate now. Needs a verified ID, an active account and the current lending terms. Pass `requestId` to make a retry safe.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketP2pLoanRefinance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketP2pLoanRefinance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketP2pLoanRefinance`: GetVaultMarketP2pLoanResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketP2pLoanRefinance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketP2pLoanRefinanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketP2pLoanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketP2pLoanRepay

> GetVaultMarketP2pLoanResponse CreateVaultMarketP2pLoanRepay(ctx, id).Execute()

Borrower: repay from your custodian account to the lender. Without `amountCents`, in full (principal plus interest to the second; the unit is free once the custodian has it). With `amountCents` below what is owed, a part payment: interest first, then principal, at least the interest owed now and at least $1; the loan stays open on the principal left. Pass `requestId` to make a retry safe.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketP2pLoanRepay(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketP2pLoanRepay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketP2pLoanRepay`: GetVaultMarketP2pLoanResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketP2pLoanRepay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketP2pLoanRepayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketP2pLoanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketP2pOffer

> ListVaultMarketP2pOffersItem CreateVaultMarketP2pOffer(ctx).Execute()

Offer to lend: `principalCents` per loan at `aprBps` a year against an item (`bookKey`) or one unit (`vaultItemId`), for up to `maxLoans` loans, lasting `expiresInDays` (default 90). Nothing is held: the custodian pays from your own account when a borrower takes it. Needs a verified ID, an active account and the current lending terms.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketP2pOffer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketP2pOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketP2pOffer`: ListVaultMarketP2pOffersItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketP2pOffer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketP2pOfferRequest struct via the builder pattern


### Return type

**ListVaultMarketP2pOffersItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketP2pOfferCancel

> ListVaultMarketP2pOffersItem CreateVaultMarketP2pOfferCancel(ctx, id).Execute()

Withdraw one of your open offers. Loans already taken from it carry on.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketP2pOfferCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketP2pOfferCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketP2pOfferCancel`: ListVaultMarketP2pOffersItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketP2pOfferCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketP2pOfferCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketP2pOffersItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketP2pOfferTake

> GetVaultMarketP2pLoanResponse CreateVaultMarketP2pOfferTake(ctx, id).Execute()

Borrow: take an offer against one of your stored vault units (`vaultItemId`). The unit is pledged (no listing, lending or withdrawal) and the principal is paid into your custodian account. Interest accrues by the second; repay any time. Pass `requestId` to make a retry safe.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketP2pOfferTake(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketP2pOfferTake``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketP2pOfferTake`: GetVaultMarketP2pLoanResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketP2pOfferTake`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketP2pOfferTakeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketP2pLoanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPaperMissionMarkCheck

> CreateVaultMarketPaperMissionMarkCheckResponse CreateVaultMarketPaperMissionMarkCheck(ctx).Execute()

Answer the \"Read the mark\" mission: the quality letter (A, B or C) you read on a market page, checked against the stored mark. One answer per market, 10 a day. Returns `{ bookKey, answer, correct, quality, mission, next }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPaperMissionMarkCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPaperMissionMarkCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPaperMissionMarkCheck`: CreateVaultMarketPaperMissionMarkCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPaperMissionMarkCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPaperMissionMarkCheckRequest struct via the builder pattern


### Return type

**CreateVaultMarketPaperMissionMarkCheckResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPaperOrder

> CreateVaultMarketPaperOrderResponse CreateVaultMarketPaperOrder(ctx).Execute()

Place a paper order (play money only). It takes visible real liquidity now at book prices with the taker fee simulated; a GTC limit rests the rest with a queue-position estimate, filled later as the real tape reaches it. Sells only paper units you hold (no paper shorts). Never touches the real book. Returns `{ order, fills }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPaperOrder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPaperOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPaperOrder`: CreateVaultMarketPaperOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPaperOrder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPaperOrderRequest struct via the builder pattern


### Return type

**CreateVaultMarketPaperOrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPaperOrderCancel

> CreateVaultMarketPaperOrderCancelResponse CreateVaultMarketPaperOrderCancel(ctx, id).Execute()

Cancel a working paper order: the play cash or paper units it held come back. Returns `{ order }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPaperOrderCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPaperOrderCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPaperOrderCancel`: CreateVaultMarketPaperOrderCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPaperOrderCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPaperOrderCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketPaperOrderCancelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPaperReset

> CreateVaultMarketPaperResetResponse CreateVaultMarketPaperReset(ctx).Execute()

Reset your paper account to $10,000 of play money: working paper orders are cancelled, paper positions cleared, and your leaderboard week starts again. Once a day.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPaperReset(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPaperReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPaperReset`: CreateVaultMarketPaperResetResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPaperReset`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPaperResetRequest struct via the builder pattern


### Return type

**CreateVaultMarketPaperResetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPartnerApply

> CreateVaultMarketPartnerApplyResponse CreateVaultMarketPartnerApply(ctx).Execute()

Apply as a Vault Drop-off Partner (a card shop that takes walk-ins for the vault). Needs the one-time ID check; ops approve.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPartnerApply(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPartnerApply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPartnerApply`: CreateVaultMarketPartnerApplyResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPartnerApply`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPartnerApplyRequest struct via the builder pattern


### Return type

**CreateVaultMarketPartnerApplyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPartnerBatch

> CreateVaultMarketPartnerBatchResponse CreateVaultMarketPartnerBatch(ctx).Execute()

An approved partner sends walk-ins in under its code: each item (catalogue variant, condition, grade, photos) becomes an inbound vault unit; the box gets one prepaid label. The partner earns a capped share of Crossly's fee on each unit's first sale.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPartnerBatch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPartnerBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPartnerBatch`: CreateVaultMarketPartnerBatchResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPartnerBatch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPartnerBatchRequest struct via the builder pattern


### Return type

**CreateVaultMarketPartnerBatchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPoolStake

> CreateVaultMarketPoolStakeResponse CreateVaultMarketPoolStake(ctx).Execute()

Stake CBX you hold at the custodian into the market pool (base units, as a string). Returns `{ stakeId, shares }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPoolStake(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPoolStake``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPoolStake`: CreateVaultMarketPoolStakeResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPoolStake`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPoolStakeRequest struct via the builder pattern


### Return type

**CreateVaultMarketPoolStakeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPoolStakeUnstake

> CreateVaultMarketPoolStakeUnstakeResponse CreateVaultMarketPoolStakeUnstake(ctx, id).Execute()

Unstake: paid out after the 7-day cooldown. Returns `{ cooldownUntil }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPoolStakeUnstake(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPoolStakeUnstake``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPoolStakeUnstake`: CreateVaultMarketPoolStakeUnstakeResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPoolStakeUnstake`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPoolStakeUnstakeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketPoolStakeUnstakeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPortfolioMarginEnroll

> CreateVaultMarketPortfolioMarginEnrollResponse CreateVaultMarketPortfolioMarginEnroll(ctx).Execute()

Admin: put a firm sub-account under portfolio margin. Needs VAULT_LEVERAGE_VENUE, your ID, an ECP-attested entity that passed KYB, and the minimum equity. Its positions are then margined on the whole account, recalculated every 5 minutes; calls and auctions are unchanged. Audited.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPortfolioMarginEnroll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPortfolioMarginEnroll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPortfolioMarginEnroll`: CreateVaultMarketPortfolioMarginEnrollResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPortfolioMarginEnroll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPortfolioMarginEnrollRequest struct via the builder pattern


### Return type

**CreateVaultMarketPortfolioMarginEnrollResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPortfolioMarginLeave

> CreateVaultMarketPortfolioMarginLeaveResponse CreateVaultMarketPortfolioMarginLeave(ctx).Execute()

Admin: take a firm sub-account off portfolio margin. Each position is then held to its own tier floors again and may be called at once. Audited.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPortfolioMarginLeave(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPortfolioMarginLeave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPortfolioMarginLeave`: CreateVaultMarketPortfolioMarginLeaveResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPortfolioMarginLeave`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPortfolioMarginLeaveRequest struct via the builder pattern


### Return type

**CreateVaultMarketPortfolioMarginLeaveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPortfolioMarginWhatIf

> CreateVaultMarketPortfolioMarginWhatIfResponse CreateVaultMarketPortfolioMarginWhatIf(ctx).Execute()

The portfolio margin requirement of a firm sub-account before and after a trade (`bookKey`, `side` buy or sell, `units`), at the Crossly mark with fees left out. Changes nothing.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPortfolioMarginWhatIf(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPortfolioMarginWhatIf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPortfolioMarginWhatIf`: CreateVaultMarketPortfolioMarginWhatIfResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPortfolioMarginWhatIf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPortfolioMarginWhatIfRequest struct via the builder pattern


### Return type

**CreateVaultMarketPortfolioMarginWhatIfResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPositionCancel

> OkResponse CreateVaultMarketPositionCancel(ctx, id).Execute()

Cancel the working order: an unfilled opening order cancels the position; an unfilled cover or close leaves it open. Returns `{ ok }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPositionCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPositionCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPositionCancel`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPositionCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPositionCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPositionClose

> CreateVaultMarketPositionCloseResponse CreateVaultMarketPositionClose(ctx, id).Execute()

Close a leveraged long: all its units, or `units` of them, each sold for at least `minCents`; each sale repays the pool first.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPositionClose(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPositionClose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPositionClose`: CreateVaultMarketPositionCloseResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPositionClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPositionCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketPositionCloseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPositionCollateral

> CreateVaultMarketPositionCollateralResponse CreateVaultMarketPositionCollateral(ctx, id).Execute()

Answer a margin call: a short posts more collateral (USD, or CBX at half value); a leveraged long pays down its pool loan. Returns the position.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPositionCollateral(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPositionCollateral``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPositionCollateral`: CreateVaultMarketPositionCollateralResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPositionCollateral`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPositionCollateralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketPositionCollateralResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPositionCover

> CreateVaultMarketPositionCloseResponse CreateVaultMarketPositionCover(ctx, id).Execute()

Cover a short: all its units, or `units` of them, each buying any unit of the key into the vault for up to `maxCents`; each goes back to its lender.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPositionCover(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPositionCover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPositionCover`: CreateVaultMarketPositionCloseResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPositionCover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPositionCoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketPositionCloseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPriceWidget

> ListVaultMarketPriceWidgetsItem CreateVaultMarketPriceWidget(ctx).Execute()

Make a price widget for one vault market (`kind` mark + `bookKey`) or one public index (`kind` index + `indexSlug`): range, layout (compact or chart), theme (auto, light, dark), optionally your own affiliate code for the link back and the sites allowed to show it (empty = any). Returns the widget with its snippets.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPriceWidget(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPriceWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPriceWidget`: ListVaultMarketPriceWidgetsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPriceWidget`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPriceWidgetRequest struct via the builder pattern


### Return type

**ListVaultMarketPriceWidgetsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketPriceWidgetPreview

> CreateVaultMarketPriceWidgetPreviewResponse CreateVaultMarketPriceWidgetPreview(ctx).Execute()

What a widget would show before you save it: `view` (value, move, quality, n, as-of, link back) and `html`, the exact page the iframe serves. Nothing is stored.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketPriceWidgetPreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketPriceWidgetPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketPriceWidgetPreview`: CreateVaultMarketPriceWidgetPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketPriceWidgetPreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketPriceWidgetPreviewRequest struct via the builder pattern


### Return type

**CreateVaultMarketPriceWidgetPreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketProBracket

> CreateVaultMarketProBracketResponse CreateVaultMarketProBracket(ctx).Execute()

A buy order with a take-profit and/or stop-loss attached; the exits arm as an OCO pair for what the entry fills.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketProBracket(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketProBracket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketProBracket`: CreateVaultMarketProBracketResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketProBracket`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketProBracketRequest struct via the builder pattern


### Return type

**CreateVaultMarketProBracketResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketProLinkCancel

> ListVaultMarketProLinksItem CreateVaultMarketProLinkCancel(ctx, id).Execute()

Cancel a linked order: its open child orders come off; a bracket's entry stays.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketProLinkCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketProLinkCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketProLinkCancel`: ListVaultMarketProLinksItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketProLinkCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketProLinkCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketProLinksItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketProOco

> CreateVaultMarketProOcoResponse CreateVaultMarketProOco(ctx).Execute()

Two orders on one key, one cancels the other: when one fills or its stop triggers, the other is cancelled (a partial fill shrinks it).

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketProOco(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketProOco``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketProOco`: CreateVaultMarketProOcoResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketProOco`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketProOcoRequest struct via the builder pattern


### Return type

**CreateVaultMarketProOcoResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketProOrder

> CreateVaultMarketProOrderResponse CreateVaultMarketProOrder(ctx).Execute()

Place an order on a vault book: limit, market (protected, IOC), stop or stop-limit; GTC, GTD, IOC or FOK; post-only; an iceberg `displayQuantity`; a TWAP or VWAP algo over a window. Idempotent on `clientOrderId`. `subaccountId` trades for a firm sub-account you may trade (your role, its KYB, its limits, no firm halt; buys pay from its custodian account and land in it).

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketProOrder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketProOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketProOrder`: CreateVaultMarketProOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketProOrder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketProOrderRequest struct via the builder pattern


### Return type

**CreateVaultMarketProOrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketProOrderCancel

> CreateVaultMarketProOrderResponse CreateVaultMarketProOrderCancel(ctx, id).Execute()

Cancel an order; what filled stays filled.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketProOrderCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketProOrderCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketProOrderCancel`: CreateVaultMarketProOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketProOrderCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketProOrderCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketProOrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketProOrderCancelAll

> CreateVaultMarketProOrderCancelAllResponse CreateVaultMarketProOrderCancelAll(ctx).Execute()

Cancel every open order, or every one on a book key. Returns `{ cancelled }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketProOrderCancelAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketProOrderCancelAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketProOrderCancelAll`: CreateVaultMarketProOrderCancelAllResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketProOrderCancelAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketProOrderCancelAllRequest struct via the builder pattern


### Return type

**CreateVaultMarketProOrderCancelAllResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketProOrderReplace

> CreateVaultMarketProOrderResponse CreateVaultMarketProOrderReplace(ctx, id).Execute()

Amend an order's price, stop or quantity (cancel-replace: it re-queues at its new price).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketProOrderReplace(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketProOrderReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketProOrderReplace`: CreateVaultMarketProOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketProOrderReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketProOrderReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketProOrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketProQuote

> CreateVaultMarketProQuoteResponse CreateVaultMarketProQuote(ctx).Execute()

Market makers: set a two-sided post-only quote on a book key (replaces the last). With cancel-on-disconnect (default) heartbeat every 30 s or it is pulled.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketProQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketProQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketProQuote`: CreateVaultMarketProQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketProQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketProQuoteRequest struct via the builder pattern


### Return type

**CreateVaultMarketProQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketProQuoteHeartbeat

> CreateVaultMarketProQuoteHeartbeatResponse CreateVaultMarketProQuoteHeartbeat(ctx).Execute()

Keep your quotes alive (cancel-on-disconnect). Returns `{ quotes, at }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketProQuoteHeartbeat(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketProQuoteHeartbeat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketProQuoteHeartbeat`: CreateVaultMarketProQuoteHeartbeatResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketProQuoteHeartbeat`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketProQuoteHeartbeatRequest struct via the builder pattern


### Return type

**CreateVaultMarketProQuoteHeartbeatResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketProTrailingStop

> CreateVaultMarketProTrailingStopResponse CreateVaultMarketProTrailingStop(ctx).Execute()

A trailing stop: its stop follows the best price since it was placed, `trailBps` or `trailCents` behind (under the high for a sell, over the low for a buy).

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketProTrailingStop(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketProTrailingStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketProTrailingStop`: CreateVaultMarketProTrailingStopResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketProTrailingStop`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketProTrailingStopRequest struct via the builder pattern


### Return type

**CreateVaultMarketProTrailingStopResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketProfileFollow

> CreateVaultMarketProfileFollowResponse CreateVaultMarketProfileFollow(ctx, handle).Execute()

Follow someone with a public vault: their new vault buys appear in your feed. No copy trading.

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketProfileFollow(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketProfileFollow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketProfileFollow`: CreateVaultMarketProfileFollowResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketProfileFollow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketProfileFollowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketProfileFollowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketRecurring

> ListVaultMarketRecurringItem CreateVaultMarketRecurring(ctx).Execute()

Start a recurring buy: one whole unit of a key every week, two weeks or month, at no more than `maxPriceCents` and/or the mark plus `capBpsOverMark`; card (with an optional backup) or custodian; optional auto-lend.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketRecurring(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketRecurring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketRecurring`: ListVaultMarketRecurringItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketRecurring`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketRecurringRequest struct via the builder pattern


### Return type

**ListVaultMarketRecurringItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketRfq

> CreateVaultMarketRfqResponse CreateVaultMarketRfq(ctx).Execute()

Ask for a quote on a block: dealers, the key's DMM and LPs, and holders are told, anonymously. Answers are firm for 60 s. Your limit is never shown.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketRfq(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketRfq``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketRfq`: CreateVaultMarketRfqResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketRfq`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketRfqRequest struct via the builder pattern


### Return type

**CreateVaultMarketRfqResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketRfqAccept

> CreateVaultMarketRfqAcceptResponse CreateVaultMarketRfqAccept(ctx, id).Execute()

Take one answer: it fills as a block through the normal fill path (fees by role, self-trade prevention, halts, LULD).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketRfqAccept(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketRfqAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketRfqAccept`: CreateVaultMarketRfqAcceptResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketRfqAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketRfqAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketRfqAcceptResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketRfqCancel

> ListVaultMarketRfqItem CreateVaultMarketRfqCancel(ctx, id).Execute()

Cancel your open quote request.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketRfqCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketRfqCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketRfqCancel`: ListVaultMarketRfqItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketRfqCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketRfqCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketRfqItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketRfqRespond

> CreateVaultMarketRfqRespondResponse CreateVaultMarketRfqRespond(ctx, id).Execute()

Answer a quote request for the full quantity with a firm price, good for 60 s. Selling needs the units free in your vault; buying pays from your Vault Markets custodian account.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketRfqRespond(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketRfqRespond``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketRfqRespond`: CreateVaultMarketRfqRespondResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketRfqRespond`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketRfqRespondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketRfqRespondResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketRfqResponseWithdraw

> CreateVaultMarketRfqRespondResponse CreateVaultMarketRfqResponseWithdraw(ctx, id).Execute()

Withdraw your live answer to a quote request.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketRfqResponseWithdraw(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketRfqResponseWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketRfqResponseWithdraw`: CreateVaultMarketRfqRespondResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketRfqResponseWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketRfqResponseWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketRfqRespondResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketSetGoal

> GetVaultMarketSetGoalResponse CreateVaultMarketSetGoal(ctx).Execute()

Start a goal on a catalogue set (category, brandSlug, modelSlug from the set search), optionally with a target: one condition, and any grade, ungraded, or one exact grade. One live goal per set.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketSetGoal(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketSetGoal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketSetGoal`: GetVaultMarketSetGoalResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketSetGoal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketSetGoalRequest struct via the builder pattern


### Return type

**GetVaultMarketSetGoalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketSetGoalBasket

> ListVaultMarketSetGoalBasketsItem CreateVaultMarketSetGoalBasket(ctx, id).Execute()

Place the \"complete my set\" basket bid: one GTC limit buy per missing piece, under a hard budget, paid by one of your saved cards or your custodian account. Every leg passes pre-trade risk and the ID line like any order. Idempotent on clientBasketId. One working basket per goal; its legs are cancelled once the set is complete.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketSetGoalBasket(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketSetGoalBasket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketSetGoalBasket`: ListVaultMarketSetGoalBasketsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketSetGoalBasket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketSetGoalBasketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketSetGoalBasketsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketSetGoalBasketCancel

> ListVaultMarketSetGoalBasketsItem CreateVaultMarketSetGoalBasketCancel(ctx, id).Execute()

Cancel the goal working basket bid: every working leg comes off the book; what already filled stays in your vault.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketSetGoalBasketCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketSetGoalBasketCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketSetGoalBasketCancel`: ListVaultMarketSetGoalBasketsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketSetGoalBasketCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketSetGoalBasketCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketSetGoalBasketsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketShareCard

> ListVaultMarketShareCardsItem CreateVaultMarketShareCard(ctx).Execute()

Make a verified share card of a position, a vault unit, a public index or your portfolio (`kind`, `refId`). The % on it is the price move since your real fill; amounts stay hidden unless `showAmounts`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketShareCard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketShareCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketShareCard`: ListVaultMarketShareCardsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketShareCard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketShareCardRequest struct via the builder pattern


### Return type

**ListVaultMarketShareCardsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketShort

> CreateVaultMarketLongResponse CreateVaultMarketShort(ctx).Execute()

Short a book key: borrow `units` lent units (default 1; at most 25/10/3/1 by liquidity tier A/B/C/D) and sell each at `limitCents` or better, posting collateral for all of them (50% of the mark each; CBX counts at half, up to 30%). All units are borrowed or none unless `allowPartial`. A partial sale opens with the units sold. Returns `{ positionId, matched, units }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketShort(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketShort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketShort`: CreateVaultMarketLongResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketShort`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketShortRequest struct via the builder pattern


### Return type

**CreateVaultMarketLongResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketSupplyBatchLabel

> GetVaultMarketSupplyBatchResponse CreateVaultMarketSupplyBatchLabel(ctx, id).Execute()

Buy a box's prepaid label again after it failed (for example once a ship-from address is saved). A box that has a label keeps it.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketSupplyBatchLabel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketSupplyBatchLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketSupplyBatchLabel`: GetVaultMarketSupplyBatchResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketSupplyBatchLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketSupplyBatchLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketSupplyBatchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketSupplyInvite

> IdStatusResponse CreateVaultMarketSupplyInvite(ctx, id).Execute()

Accept or decline a Market Open invite (`accept`).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketSupplyInvite(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketSupplyInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketSupplyInvite`: IdStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketSupplyInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketSupplyInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**IdStatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketSupplySend

> CreateVaultMarketSupplySendResponse CreateVaultMarketSupplySend(ctx).Execute()

Send inventory items matched to a catalogue variant to the vault as one box: each becomes an inbound vault unit, and the box gets one prepaid inbound label (a PDF to print). Optional: auto-list each unit in the vault book at a share of the mark once stored, and lend when idle. Items that cannot go say why. Crosslisted listings stay live.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketSupplySend(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketSupplySend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketSupplySend`: CreateVaultMarketSupplySendResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketSupplySend`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketSupplySendRequest struct via the builder pattern


### Return type

**CreateVaultMarketSupplySendResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketTerm

> TermsVersionResponse CreateVaultMarketTerm(ctx).Execute()

Accept the current lending terms (needed to lend or short). Returns `{ termsVersion }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketTerm(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketTerm`: TermsVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketTerm`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketTermRequest struct via the builder pattern


### Return type

**TermsVersionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketTodClaimObject

> CreateVaultMarketTodClaimObjectResponse CreateVaultMarketTodClaimObject(ctx, id).Execute()

Stop a transfer-on-death claim made on your own account. Returns `{ id, status, estateUnlocked }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketTodClaimObject(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketTodClaimObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketTodClaimObject`: CreateVaultMarketTodClaimObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketTodClaimObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketTodClaimObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketTodClaimObjectResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketTodMyClaim

> CreateVaultMarketTodMyClaimResponse CreateVaultMarketTodMyClaim(ctx).Execute()

Start a claim as a named beneficiary (ID first), naming the account holder by @username or email. It stays a draft until you add documents and submit. Returns `{ id, status, existing }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketTodMyClaim(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketTodMyClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketTodMyClaim`: CreateVaultMarketTodMyClaimResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketTodMyClaim`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketTodMyClaimRequest struct via the builder pattern


### Return type

**CreateVaultMarketTodMyClaimResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketTodMyClaimDocument

> CreateVaultMarketTodMyClaimDocumentResponse CreateVaultMarketTodMyClaimDocument(ctx, id).Execute()

Add one document to a draft claim (or one ops asked more of): kind, filename and the file as base64 (PDF, JPEG or PNG, up to 8 MB). Kept privately for ops review. Returns `{ id, kind, label, filename, bytes }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketTodMyClaimDocument(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketTodMyClaimDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketTodMyClaimDocument`: CreateVaultMarketTodMyClaimDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketTodMyClaimDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketTodMyClaimDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketTodMyClaimDocumentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketTodMyClaimSubmit

> IdStatusResponse CreateVaultMarketTodMyClaimSubmit(ctx, id).Execute()

Send a claim for ops review (it needs a death certificate). The account holder is told and can stop it. Returns `{ id, status }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketTodMyClaimSubmit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketTodMyClaimSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketTodMyClaimSubmit`: IdStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketTodMyClaimSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketTodMyClaimSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**IdStatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketTodMyClaimWithdraw

> IdStatusResponse CreateVaultMarketTodMyClaimWithdraw(ctx, id).Execute()

Withdraw a claim before it is released. Returns `{ id, status }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketTodMyClaimWithdraw(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketTodMyClaimWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketTodMyClaimWithdraw`: IdStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketTodMyClaimWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketTodMyClaimWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**IdStatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketTodTerm

> TermsVersionResponse CreateVaultMarketTodTerm(ctx).Execute()

Accept the current transfer-on-death terms. Returns `{ termsVersion }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketTodTerm(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketTodTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketTodTerm`: TermsVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketTodTerm`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketTodTermRequest struct via the builder pattern


### Return type

**TermsVersionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketUnitEvidenceScan

> GetVaultMarketUnitEvidenceResponse CreateVaultMarketUnitEvidenceScan(ctx, id).Execute()

Add a finished scan of your own (`captureId`, of the same catalogue item) to your vault unit. Shown as added by the owner, not as intake evidence.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketUnitEvidenceScan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketUnitEvidenceScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketUnitEvidenceScan`: GetVaultMarketUnitEvidenceResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketUnitEvidenceScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketUnitEvidenceScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketUnitEvidenceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketUnitSellNow

> CreateVaultMarketInstantExitResponse CreateVaultMarketUnitSellNow(ctx, id).Execute()

Sell one unit of this item now into the best bid (Instant Exit: an immediate-or-cancel pro order), never below `minCents`. Returns `{ status: filled | missed, filledCents, quote, order }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketUnitSellNow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketUnitSellNow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketUnitSellNow`: CreateVaultMarketInstantExitResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketUnitSellNow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketUnitSellNowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketInstantExitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketValuationStatement

> CreateVaultMarketValuationStatementResponse CreateVaultMarketValuationStatement(ctx).Execute()

Issue an insurance or appraisal statement for a day: frozen with the sha256 of its record, sealed in the Crossly audit chain, with a public verify link. preparedFor and policyRef are printed on it and never shown on verify. Moves no money.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketValuationStatement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketValuationStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketValuationStatement`: CreateVaultMarketValuationStatementResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketValuationStatement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketValuationStatementRequest struct via the builder pattern


### Return type

**CreateVaultMarketValuationStatementResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketWatchlist

> CreateVaultMarketWatchlistResponse CreateVaultMarketWatchlist(ctx).Execute()

Watch a book key. Returns `{ bookKey, watching }`.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketWatchlist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketWatchlist`: CreateVaultMarketWatchlistResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketWatchlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketWatchlistRequest struct via the builder pattern


### Return type

**CreateVaultMarketWatchlistResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketWidgetToken

> CreateVaultMarketWidgetTokenResponse CreateVaultMarketWidgetToken(ctx).Execute()

Mint a widget token for one phone (`deviceName`, `installId`); minting again with the same pair replaces that phone token. The token reads the widget summary and nothing else, and is returned once. At most five phones.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultMarketWidgetToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultMarketWidgetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketWidgetToken`: CreateVaultMarketWidgetTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultMarketWidgetToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketWidgetTokenRequest struct via the builder pattern


### Return type

**CreateVaultMarketWidgetTokenResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultReceiptTokenConfirm

> CreateVaultReceiptTokenConfirmResponse CreateVaultReceiptTokenConfirm(ctx, id).Execute()

After your wallet sent the export: read the chain and mark the token live once its one token exists. Pass the `signature` your wallet returned to record it (kept only if it landed).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultReceiptTokenConfirm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultReceiptTokenConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultReceiptTokenConfirm`: CreateVaultReceiptTokenConfirmResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultReceiptTokenConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultReceiptTokenConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultReceiptTokenConfirmResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultReceiptTokenExport

> CreateVaultReceiptTokenExportResponse CreateVaultReceiptTokenExport(ctx).Execute()

Send one unit (`vaultItemId`) warehouse receipt to one of your verified Solana wallets (`wallet`) as a token, accepting the current terms (`acceptTerms`). Returns a base64 transaction already signed by the new mint and the receipts authority: your wallet pays the rent and must sign and send it within about two minutes, then call confirm. Nothing is minted until it does. Needs a verified ID; the unit must be stored, unpledged, not lent and not for sale.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultReceiptTokenExport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultReceiptTokenExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultReceiptTokenExport`: CreateVaultReceiptTokenExportResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultReceiptTokenExport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultReceiptTokenExportRequest struct via the builder pattern


### Return type

**CreateVaultReceiptTokenExportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultReceiptTokenRedeem

> CreateVaultReceiptTokenRedeemResponse CreateVaultReceiptTokenRedeem(ctx, id).Execute()

Redeem a receipt token held in one of your verified Solana wallets (`wallet`): returns an unsigned base64 burn your wallet signs and sends (it also closes the token account and returns its rent to you). Once it lands, call redeem/confirm and the unit is yours in the vault. Refused while the token is frozen.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultReceiptTokenRedeem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultReceiptTokenRedeem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultReceiptTokenRedeem`: CreateVaultReceiptTokenRedeemResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultReceiptTokenRedeem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultReceiptTokenRedeemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultReceiptTokenRedeemResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultReceiptTokenRedeemConfirm

> CreateVaultReceiptTokenRedeemConfirmResponse CreateVaultReceiptTokenRedeemConfirm(ctx, id).Execute()

After your wallet sent the redeem: read the chain; once the token is burned the unit moves into your vault (if it was not yours already). Pass the `signature` to record it.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultReceiptTokenRedeemConfirm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultReceiptTokenRedeemConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultReceiptTokenRedeemConfirm`: CreateVaultReceiptTokenRedeemConfirmResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultReceiptTokenRedeemConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultReceiptTokenRedeemConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultReceiptTokenRedeemConfirmResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultReceiptTokenWalletAllow

> CreateVaultReceiptTokenWalletAllowResponse CreateVaultReceiptTokenWalletAllow(ctx).Execute()

Put one of your verified Solana wallets (`wallet`) on the receipt-token allow-list so it can hold and receive receipt tokens. Returns a base64 transaction already signed by the receipts authority: the wallet pays the small rent (about 0.001 SOL) and must sign and send it within about two minutes. Needs a verified ID and a wallet that passes the sanctions screen.

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
    resp, r, err := apiClient.VaultMarketsApi.CreateVaultReceiptTokenWalletAllow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.CreateVaultReceiptTokenWalletAllow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultReceiptTokenWalletAllow`: CreateVaultReceiptTokenWalletAllowResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.CreateVaultReceiptTokenWalletAllow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultReceiptTokenWalletAllowRequest struct via the builder pattern


### Return type

**CreateVaultReceiptTokenWalletAllowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketAgentWatch

> ListVaultMarketAgentWatchesItem DeleteVaultMarketAgentWatch(ctx, id).Execute()

Cancel a live price watch. Returns the watch.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketAgentWatch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketAgentWatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketAgentWatch`: ListVaultMarketAgentWatchesItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketAgentWatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketAgentWatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketAgentWatchesItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketAlert

> OkResponse DeleteVaultMarketAlert(ctx, id).Execute()

Delete a price alert. Returns `{ ok }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketAlert(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketAlert`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketBullionFamilyPeg

> ListVaultMarketBullionFamilyPegsItem DeleteVaultMarketBullionFamilyPeg(ctx, id).Execute()

Stop a family-wide quote: every leg order comes off the book; what filled stays filled.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketBullionFamilyPeg(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketBullionFamilyPeg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketBullionFamilyPeg`: ListVaultMarketBullionFamilyPegsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketBullionFamilyPeg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketBullionFamilyPegRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketBullionFamilyPegsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketBullionPeg

> ListVaultMarketBullionPegsItem DeleteVaultMarketBullionPeg(ctx, id).Execute()

Stop a spot-pegged quote: its resting order comes off the book; what filled stays filled.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketBullionPeg(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketBullionPeg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketBullionPeg`: ListVaultMarketBullionPegsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketBullionPeg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketBullionPegRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketBullionPegsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketDealerBand

> ListVaultMarketDealerBandsItem DeleteVaultMarketDealerBand(ctx).Execute()

Stop a band on a key; its standing bid comes off.

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
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketDealerBand(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketDealerBand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketDealerBand`: ListVaultMarketDealerBandsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketDealerBand`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketDealerBandRequest struct via the builder pattern


### Return type

**ListVaultMarketDealerBandsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketEntityMember

> OkResponse DeleteVaultMarketEntityMember(ctx, id, memberId).Execute()

Admin: remove a member (or leave, yourself). An entity always keeps one admin.

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
    id := "id_example" // string | 
    memberId := "memberId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketEntityMember(context.Background(), id, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketEntityMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketEntityMember`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketEntityMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketEntityMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketFollowCopyWatch

> DeleteVaultMarketFollowCopyWatchResponse DeleteVaultMarketFollowCopyWatch(ctx, handle).Execute()

Stop copy-watching a collector.

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketFollowCopyWatch(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketFollowCopyWatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketFollowCopyWatch`: DeleteVaultMarketFollowCopyWatchResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketFollowCopyWatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketFollowCopyWatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteVaultMarketFollowCopyWatchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketIndex

> OkResponse DeleteVaultMarketIndex(ctx, slug).Execute()

Delete an index of your own. Returns `{ ok }`.

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
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketIndex(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketIndex`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketLocate

> ListVaultMarketLocateMineItem DeleteVaultMarketLocate(ctx, id).Execute()

Release a live locate before it lapses.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketLocate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketLocate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketLocate`: ListVaultMarketLocateMineItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketLocate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketLocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketLocateMineItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketOpeningWaitlist

> DeleteVaultMarketOpeningWaitlistResponse DeleteVaultMarketOpeningWaitlist(ctx, slug).Execute()

Leave a key's waitlist.

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
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketOpeningWaitlist(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketOpeningWaitlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketOpeningWaitlist`: DeleteVaultMarketOpeningWaitlistResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketOpeningWaitlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketOpeningWaitlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteVaultMarketOpeningWaitlistResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketPriceWidget

> DeleteVaultMarketPriceWidgetResponse DeleteVaultMarketPriceWidget(ctx, id).Execute()

Delete a widget. Pages that still paste it show that it is no longer available. Returns `{ id, archived }`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketPriceWidget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketPriceWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketPriceWidget`: DeleteVaultMarketPriceWidgetResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketPriceWidget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketPriceWidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteVaultMarketPriceWidgetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketProQuote

> DeleteVaultMarketProQuoteResponse DeleteVaultMarketProQuote(ctx).Execute()

Pull your quote on a book key, or every quote. Returns `{ pulled }`.

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
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketProQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketProQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketProQuote`: DeleteVaultMarketProQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketProQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketProQuoteRequest struct via the builder pattern


### Return type

**DeleteVaultMarketProQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketProfileFollow

> CreateVaultMarketProfileFollowResponse DeleteVaultMarketProfileFollow(ctx, handle).Execute()

Stop following someone.

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketProfileFollow(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketProfileFollow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketProfileFollow`: CreateVaultMarketProfileFollowResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketProfileFollow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketProfileFollowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketProfileFollowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketRecurring

> ListVaultMarketRecurringItem DeleteVaultMarketRecurring(ctx, id).Execute()

Stop a recurring buy for good; its history stays.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketRecurring(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketRecurring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketRecurring`: ListVaultMarketRecurringItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketRecurring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketRecurringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketRecurringItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketSetGoal

> DeleteVaultMarketSetGoalResponse DeleteVaultMarketSetGoal(ctx, id).Execute()

Remove a set goal. A working basket bid is cancelled first; what filled stays in your vault.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketSetGoal(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketSetGoal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketSetGoal`: DeleteVaultMarketSetGoalResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketSetGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketSetGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteVaultMarketSetGoalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketShareCard

> OkResponse DeleteVaultMarketShareCard(ctx, id).Execute()

Take a share card down: its link and image stop working. Returns `{ ok }`.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketShareCard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketShareCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketShareCard`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketShareCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketShareCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketSupplyRoundUp

> GetVaultMarketSupplyRoundUpResponse DeleteVaultMarketSupplyRoundUp(ctx).Execute()

Stop the round-up (its history stays; an armed buy is cancelled).

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
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketSupplyRoundUp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketSupplyRoundUp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketSupplyRoundUp`: GetVaultMarketSupplyRoundUpResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketSupplyRoundUp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketSupplyRoundUpRequest struct via the builder pattern


### Return type

**GetVaultMarketSupplyRoundUpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketUnitEvidence

> GetVaultMarketUnitEvidenceResponse DeleteVaultMarketUnitEvidence(ctx, id, evidenceId).Execute()

Take back a scan you added to your vault unit (intake evidence stays).

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
    id := "id_example" // string | 
    evidenceId := "evidenceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketUnitEvidence(context.Background(), id, evidenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketUnitEvidence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketUnitEvidence`: GetVaultMarketUnitEvidenceResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketUnitEvidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**evidenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketUnitEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetVaultMarketUnitEvidenceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketWatchlist

> CreateVaultMarketWatchlistResponse DeleteVaultMarketWatchlist(ctx).Execute()

Stop watching a book key (`?bookKey=`). Returns `{ bookKey, watching }`.

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
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketWatchlist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketWatchlist`: CreateVaultMarketWatchlistResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketWatchlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketWatchlistRequest struct via the builder pattern


### Return type

**CreateVaultMarketWatchlistResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketWidgetToken

> DeleteVaultMarketWidgetTokenResponse DeleteVaultMarketWidgetToken(ctx).Execute()

Turn the widget off on every phone: revokes all your widget tokens. Returns `{ revoked }`, how many.

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
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketWidgetToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketWidgetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketWidgetToken`: DeleteVaultMarketWidgetTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketWidgetToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketWidgetTokenRequest struct via the builder pattern


### Return type

**DeleteVaultMarketWidgetTokenResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteVaultMarketWidgetTokenById

> DeleteVaultMarketWidgetTokenByIdResponse DeleteVaultMarketWidgetTokenById(ctx, id).Execute()

Turn the widget off on one phone: revokes that widget token at once. Only a widget token can be revoked here.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.DeleteVaultMarketWidgetTokenById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.DeleteVaultMarketWidgetTokenById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVaultMarketWidgetTokenById`: DeleteVaultMarketWidgetTokenByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.DeleteVaultMarketWidgetTokenById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultMarketWidgetTokenByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteVaultMarketWidgetTokenByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultAttestation

> GetVaultAttestationResponse GetVaultAttestation(ctx).Execute()

Vault attestations: independent physical counts, the insurance certificate (with a warning when insured value in the vault exceeds its aggregate or per-occurrence limit), SOC 1 Type II and SOC 2 reports. Each carries its signed document.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultAttestation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultAttestation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultAttestation`: GetVaultAttestationResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultAttestation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultAttestationRequest struct via the builder pattern


### Return type

**GetVaultAttestationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketAccount

> GetVaultMarketAccountResponse GetVaultMarketAccount(ctx).Execute()

Which Vault Markets registrations are live (and the env each waits on), your custodian account, limits and balances.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketAccount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketAccount`: GetVaultMarketAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketAccountRequest struct via the builder pattern


### Return type

**GetVaultMarketAccountResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketAdl

> GetVaultMarketAdlResponse GetVaultMarketAdl(ctx).Execute()

The insurance fund (booked and free to draw), the most one event may draw before auto-deleveraging (ADL) starts, the loss waterfall in order, and the latest ADL events (no user or position ids). Money in cents.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketAdl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketAdl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketAdl`: GetVaultMarketAdlResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketAdl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketAdlRequest struct via the builder pattern


### Return type

**GetVaultMarketAdlResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketAdlPosition

> ListVaultMarketAdlPositionsItem GetVaultMarketAdlPosition(ctx, id).Execute()

One of your positions with its ADL lights (0 once it has ended).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketAdlPosition(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketAdlPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketAdlPosition`: ListVaultMarketAdlPositionsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketAdlPosition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketAdlPositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketAdlPositionsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketAdvanceQuote

> GetVaultMarketAdvanceQuoteResponse GetVaultMarketAdvanceQuote(ctx).Execute()

What an advance on a vault unit for sale could be: up to half of the lower of its mark and asking price, free for 60 days.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketAdvanceQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketAdvanceQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketAdvanceQuote`: GetVaultMarketAdvanceQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketAdvanceQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketAdvanceQuoteRequest struct via the builder pattern


### Return type

**GetVaultMarketAdvanceQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketAgentDraft

> ListVaultMarketAgentDraftsItem GetVaultMarketAgentDraft(ctx, id).Execute()

One drafted order: item, side, units, limit (cents), the book when drafted, status, and the orderBody that places it.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketAgentDraft(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketAgentDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketAgentDraft`: ListVaultMarketAgentDraftsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketAgentDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketAgentDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketAgentDraftsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketAlertPref

> GetVaultMarketAlertPrefResponse GetVaultMarketAlertPref(ctx).Execute()

Your vault alert settings: every kind (price, fill, margin_call, adl, p2p_loan, auction, market_open, receipt_token, set_goal, follow, account), every channel (push, telegram, terminal, email, webhook) with whether it is set up, the full on/off grid, the locked switches (a critical kind always pushes) and quiet hours.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketAlertPref(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketAlertPref``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketAlertPref`: GetVaultMarketAlertPrefResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketAlertPref`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketAlertPrefRequest struct via the builder pattern


### Return type

**GetVaultMarketAlertPrefResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketAuctionLot

> GetVaultMarketAuctionLotResponse GetVaultMarketAuctionLot(ctx, id).Execute()

One lot: the price showing, the next minimum bid, whether a reserve exists and is met, the end (extended by late bids), the bid history and your standing.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketAuctionLot(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketAuctionLot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketAuctionLot`: GetVaultMarketAuctionLotResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketAuctionLot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketAuctionLotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketAuctionLotResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketAuctionMine

> GetVaultMarketAuctionMineResponse GetVaultMarketAuctionMine(ctx).Execute()

Lots you consigned (with your reserve) and lots you bid on (with your maximum and whether you lead).

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketAuctionMine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketAuctionMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketAuctionMine`: GetVaultMarketAuctionMineResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketAuctionMine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketAuctionMineRequest struct via the builder pattern


### Return type

**GetVaultMarketAuctionMineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketBlock

> ListVaultMarketBlocksItem GetVaultMarketBlock(ctx, orderId).Execute()

One block order: its order, plan lines (planned and allocated units, average price, cash share, status).

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
    orderId := "orderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketBlock(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketBlock`: ListVaultMarketBlocksItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketBlocksItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketBullionFor

> GetVaultMarketBullionForResponse GetVaultMarketBullionFor(ctx).Execute()

The bullion key families one book key belongs to, with spot for its metal as a labelled reference and its melt value at that spot. Empty families when the item is not keyed bullion.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketBullionFor(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketBullionFor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketBullionFor`: GetVaultMarketBullionForResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketBullionFor`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketBullionForRequest struct via the builder pattern


### Return type

**GetVaultMarketBullionForResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketBullionKey

> GetVaultMarketBullionKeyResponse GetVaultMarketBullionKey(ctx, slug).Execute()

One bullion key family: its catalogue products, every live book key with best bid, best ask, vault units and its mark (from the book only), live pegged quote counts, and spot as a labelled reference.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketBullionKey(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketBullionKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketBullionKey`: GetVaultMarketBullionKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketBullionKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketBullionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketBullionKeyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketBullionPegPreview

> GetVaultMarketBullionPegPreviewResponse GetVaultMarketBullionPegPreview(ctx).Execute()

The price a spot-pegged quote would rest at now (spot plus a premium in cents or basis points, clamped to floor and cap). No price when spot is stale or missing: such a quote waits off the book.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketBullionPegPreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketBullionPegPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketBullionPegPreview`: GetVaultMarketBullionPegPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketBullionPegPreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketBullionPegPreviewRequest struct via the builder pattern


### Return type

**GetVaultMarketBullionPegPreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketCall

> GetVaultMarketCallResponse GetVaultMarketCall(ctx, bookKey).Execute()

Recent call-auction uncrosses on a book key: kind (open, reopen, periodic, daily, buy-in, sell-off), clearing price in cents and units.

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
    bookKey := "bookKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketCall(context.Background(), bookKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketCall`: GetVaultMarketCallResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketCallResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketCapacity

> GetVaultMarketCapacityResponse GetVaultMarketCapacity(ctx).Execute()

Capacity from the book: `bookKey` (+ `units`) gives ADV, depth within ±5% of the mark, vault stock, lendable units and days to liquidate at 20% of ADV; `index` + `notional` (cents) gives how much of the basket the asks fill now and the days the rest needs.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketCapacity(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketCapacity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketCapacity`: GetVaultMarketCapacityResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketCapacity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketCapacityRequest struct via the builder pattern


### Return type

**GetVaultMarketCapacityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketChannel

> GetVaultMarketChannelResponse GetVaultMarketChannel(ctx).Execute()

Your Vault Markets channel switches: Telegram alerts and digest, the daily cap on vault nudges by push (default 5), whether a Telegram chat is linked, and how many nudges went out today.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketChannel(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketChannel`: GetVaultMarketChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketChannel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketChannelRequest struct via the builder pattern


### Return type

**GetVaultMarketChannelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketCompare

> GetVaultMarketCompareResponse GetVaultMarketCompare(ctx).Execute()

Up to 5 book keys and/or index slugs (`items`, comma-separated) as daily series on the same days, each rebased to 100.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketCompare(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketCompare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketCompare`: GetVaultMarketCompareResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketCompare`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketCompareRequest struct via the builder pattern


### Return type

**GetVaultMarketCompareResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketConfirmation

> GetVaultMarketConfirmationResponse GetVaultMarketConfirmation(ctx, tradeId).Execute()

One trade confirmation.

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
    tradeId := "tradeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketConfirmation(context.Background(), tradeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketConfirmation`: GetVaultMarketConfirmationResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketConfirmationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketConfirmationPdf

> *os.File GetVaultMarketConfirmationPdf(ctx, tradeId).Execute()

One trade confirmation as a PDF.

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
    tradeId := "tradeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketConfirmationPdf(context.Background(), tradeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketConfirmationPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketConfirmationPdf`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketConfirmationPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketConfirmationPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

***os.File**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketCorrelation

> GetVaultMarketCorrelationResponse GetVaultMarketCorrelation(ctx).Execute()

The correlation matrix of the sector indexes' weekly returns over `weeks` (default 52), pairwise-complete, with n per pair; a pair under 26 shared weeks has no number.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketCorrelation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketCorrelation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketCorrelation`: GetVaultMarketCorrelationResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketCorrelation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketCorrelationRequest struct via the builder pattern


### Return type

**GetVaultMarketCorrelationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketCreatorAttribution

> GetVaultMarketCreatorAttributionResponse GetVaultMarketCreatorAttribution(ctx).Execute()

The creator who brought you to Vault Markets, when one did: when it started and ends, whether it still earns, and the disclosure. Your prices and fees are the same either way.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketCreatorAttribution(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketCreatorAttribution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketCreatorAttribution`: GetVaultMarketCreatorAttributionResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketCreatorAttribution`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketCreatorAttributionRequest struct via the builder pattern


### Return type

**GetVaultMarketCreatorAttributionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketCreatorDashboard

> GetVaultMarketCreatorDashboardResponse GetVaultMarketCreatorDashboard(ctx).Execute()

For creators: whether you joined the programme, each of your codes with traders attributed and still earning (counts only), shares accrued, paid and voided (cents) and the latest shares.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketCreatorDashboard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketCreatorDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketCreatorDashboard`: GetVaultMarketCreatorDashboardResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketCreatorDashboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketCreatorDashboardRequest struct via the builder pattern


### Return type

**GetVaultMarketCreatorDashboardResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketCreatorProgramme

> GetVaultMarketCreatorProgrammeResponse GetVaultMarketCreatorProgramme(ctx).Execute()

The Vault Markets creator programme: the share of the fee Crossly keeps that a creator earns, the per-fill and lifetime caps (cents), the attribution window in days, the 2% net floor it sits under, and the disclosure words.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketCreatorProgramme(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketCreatorProgramme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketCreatorProgramme`: GetVaultMarketCreatorProgrammeResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketCreatorProgramme`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketCreatorProgrammeRequest struct via the builder pattern


### Return type

**GetVaultMarketCreatorProgrammeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketCreatorRoom

> GetVaultMarketCreatorRoomResponse GetVaultMarketCreatorRoom(ctx).Execute()

For creators: your vault shelf room (applied, verified, rejected or paused) and whether your public vault profile is on, which the room needs.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketCreatorRoom(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketCreatorRoom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketCreatorRoom`: GetVaultMarketCreatorRoomResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketCreatorRoom`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketCreatorRoomRequest struct via the builder pattern


### Return type

**GetVaultMarketCreatorRoomResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketCredit

> GetVaultMarketCreditResponse GetVaultMarketCredit(ctx, id).Execute()

One loan or advance with its pledged units and its money log (draws, interest, repayments, sale proceeds).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketCredit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketCredit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketCredit`: GetVaultMarketCreditResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketCredit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketCreditResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketDcm

> GetVaultMarketDcmResponse GetVaultMarketDcm(ctx).Execute()

The partner exchange (DCM) that lists and clears futures and options on Crossly indexes (needs VAULT_INDEX_DCM + VAULT_INDEX_DCM_API_URL + VAULT_INDEX_DCM_API_KEY), and its contracts as last read: perps, futures and options with the index each settles on, multiplier, tick, expiry and strike. Read-only: Crossly is the licensor and data provider, never the venue.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketDcm(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketDcm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketDcm`: GetVaultMarketDcmResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketDcm`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketDcmRequest struct via the builder pattern


### Return type

**GetVaultMarketDcmResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketDcmContract

> GetVaultMarketDcmContractResponse GetVaultMarketDcmContract(ctx, symbol).Execute()

One partner contract: its spec, its funding prints (newest first; the partner computes and settles funding), annualised funding over the last 21 prints, and the Crossly index level it settles on.

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
    symbol := "symbol_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketDcmContract(context.Background(), symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketDcmContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketDcmContract`: GetVaultMarketDcmContractResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketDcmContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketDcmContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketDcmContractResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketDcmContractBook

> GetVaultMarketDcmContractBookResponse GetVaultMarketDcmContractBook(ctx, symbol).Execute()

The partner exchange live order book for a contract (cached 5 seconds), with mark, index, last, open interest, 24 h volume, spread in bps and the basis of its mark over the Crossly index in bps. Prices are index points.

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
    symbol := "symbol_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketDcmContractBook(context.Background(), symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketDcmContractBook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketDcmContractBook`: GetVaultMarketDcmContractBookResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketDcmContractBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketDcmContractBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketDcmContractBookResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketDcmOptionsInput

> GetVaultMarketDcmOptionsInputResponse GetVaultMarketDcmOptionsInput(ctx, slug).Execute()

Options inputs for an index (indexes only, never a single item): Crossly realised volatility (raw and Geltner-desmoothed, from weekly returns), the forward (the partner future mark, else the Crossly close), Black-76 greeks on a strike x expiry grid and on every option the partner lists on the index. Model inputs, not quotes.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketDcmOptionsInput(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketDcmOptionsInput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketDcmOptionsInput`: GetVaultMarketDcmOptionsInputResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketDcmOptionsInput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketDcmOptionsInputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketDcmOptionsInputResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketDcmPosition

> GetVaultMarketDcmPositionResponse GetVaultMarketDcmPosition(ctx).Execute()

The partner exchange account linked to your account (or to the firm sub-account `subaccountId`) and its positions, read-only (cached 15 seconds); index futures on category indexes count as portfolio-margin hedges for that account: contracts, side, entry and mark price, unrealised P&L and margin in cents, liquidation price and notional in cents. Not linked: `linked: false`.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketDcmPosition(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketDcmPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketDcmPosition`: GetVaultMarketDcmPositionResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketDcmPosition`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketDcmPositionRequest struct via the builder pattern


### Return type

**GetVaultMarketDcmPositionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketDealer

> GetVaultMarketDealerResponse GetVaultMarketDealer(ctx).Execute()

Your dealer status (applied, approved, revoked) and your Instant Exit bands.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketDealer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketDealer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketDealer`: GetVaultMarketDealerResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketDealer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketDealerRequest struct via the builder pattern


### Return type

**GetVaultMarketDealerResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketDealerPool

> ListVaultMarketDealerPoolsItem GetVaultMarketDealerPool(ctx, id).Execute()

One of your curve pools, with its ladder and live legs.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketDealerPool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketDealerPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketDealerPool`: ListVaultMarketDealerPoolsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketDealerPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketDealerPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketDealerPoolsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketDigestPreview

> GetVaultMarketDigestPreviewResponse GetVaultMarketDigestPreview(ctx).Execute()

This week's vault digest as it would go out: units at the display marks (unmarked units counted, not valued), the like-for-like change from a week ago, offers on your units, watched keys that got a first mark, and one Learn card.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketDigestPreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketDigestPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketDigestPreview`: GetVaultMarketDigestPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketDigestPreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketDigestPreviewRequest struct via the builder pattern


### Return type

**GetVaultMarketDigestPreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketEntity

> GetVaultMarketEntityResponse GetVaultMarketEntity(ctx).Execute()

The entities (funds, LLCs, trusts, SPVs) you belong to, and invitations waiting for you.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketEntity(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketEntity`: GetVaultMarketEntityResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketEntity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketEntityRequest struct via the builder pattern


### Return type

**GetVaultMarketEntityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketEntityById

> GetVaultMarketEntityByIdResponse GetVaultMarketEntityById(ctx, id).Execute()

One entity: members and roles, sub-accounts with their limits, firm halts, caps.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketEntityById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketEntityById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketEntityById`: GetVaultMarketEntityByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketEntityById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketEntityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketEntityByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketEstimate

> GetVaultMarketEstimateResponse GetVaultMarketEstimate(ctx, bookKey).Execute()

The Crossly Estimate for a book key — a model, not a price. Method (own trades, index-adjusted last trade, grade ratio), cents with 68% and 95% bands, High/Medium/Low from the walk-forward backtest's FSD, and n. Hidden (`available: false`) under 50 backtested trades or when FSD > 0.30. Never used for margin, loans or settlement.

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
    bookKey := "bookKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketEstimate(context.Background(), bookKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketEstimate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketEstimate`: GetVaultMarketEstimateResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketEstimate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketEstimateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketEstimateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketExecutionQuality

> GetVaultMarketExecutionQualityResponse GetVaultMarketExecutionQuality(ctx).Execute()

The quarterly execution-quality report (Rule 605 style): fill rates, speed, effective vs quoted spread and shortfall by order type and size, and the busiest keys. `quarter` like 2026-Q4; the current quarter is provisional.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketExecutionQuality(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketExecutionQuality``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketExecutionQuality`: GetVaultMarketExecutionQualityResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketExecutionQuality`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketExecutionQualityRequest struct via the builder pattern


### Return type

**GetVaultMarketExecutionQualityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketFeeQuote

> GetVaultMarketFeeQuoteResponse GetVaultMarketFeeQuote(ctx).Execute()

Your fee for an order at a price: as maker (the order rests) and as taker (it fills at once), with the programme that sets your rate and any Crossly One trim. A buyer's fee is added to the price; a seller's comes out of the proceeds.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketFeeQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketFeeQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketFeeQuote`: GetVaultMarketFeeQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketFeeQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketFeeQuoteRequest struct via the builder pattern


### Return type

**GetVaultMarketFeeQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketFeeSchedule

> GetVaultMarketFeeScheduleResponse GetVaultMarketFeeSchedule(ctx).Execute()

The fee schedule by role: maker and taker rates for retail, 30-day volume, Registered LPs and DMMs, the CBX discount, the referral share and the 2% net floor Crossly always keeps.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketFeeSchedule(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketFeeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketFeeSchedule`: GetVaultMarketFeeScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketFeeSchedule`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketFeeScheduleRequest struct via the builder pattern


### Return type

**GetVaultMarketFeeScheduleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketFollowSharing

> GetVaultMarketFollowSharingResponse GetVaultMarketFollowSharing(ctx).Execute()

Your sharing switches: the public vault profile, new buys, units sent in, sales, Hall loans, bids and asks for copy-watchers, and the leaderboards. All off until you turn them on.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketFollowSharing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketFollowSharing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketFollowSharing`: GetVaultMarketFollowSharingResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketFollowSharing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketFollowSharingRequest struct via the builder pattern


### Return type

**GetVaultMarketFollowSharingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketGemRate

> GetVaultMarketGemRateResponse GetVaultMarketGemRate(ctx, bookKey).Execute()

The Crossly gem rate: of cards graded through Crossly, the share that came back at each grader's gem grade, for this item and its category, with a Wilson 95% band. 10 returns per grader or no number. Not a grader population report.

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
    bookKey := "bookKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketGemRate(context.Background(), bookKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketGemRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketGemRate`: GetVaultMarketGemRateResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketGemRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketGemRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketGemRateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketGradeToVault

> ListVaultMarketGradeToVaultItem GetVaultMarketGradeToVault(ctx, id).Execute()

One Grade to Vault flow, brought up to date with its submission and unit.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketGradeToVault(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketGradeToVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketGradeToVault`: ListVaultMarketGradeToVaultItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketGradeToVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketGradeToVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketGradeToVaultItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketGradeToVaultCandidate

> GetVaultMarketGradeToVaultCandidateResponse GetVaultMarketGradeToVaultCandidate(ctx).Execute()

What you can start Grade to Vault on: raw inventory matched to a catalogue item, and raw vault units on the grading hub shelf; plus whether Grading Runs are open.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketGradeToVaultCandidate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketGradeToVaultCandidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketGradeToVaultCandidate`: GetVaultMarketGradeToVaultCandidateResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketGradeToVaultCandidate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketGradeToVaultCandidateRequest struct via the builder pattern


### Return type

**GetVaultMarketGradeToVaultCandidateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketGradingSignal

> GetVaultMarketGradingSignalResponse GetVaultMarketGradingSignal(ctx).Execute()

For a raw book key: the marks of the same item's graded keys, the spread over raw, grading cost per grader, the margin tier grading lifts it to, your raw vault units it pays to grade, and where to act.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketGradingSignal(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketGradingSignal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketGradingSignal`: GetVaultMarketGradingSignalResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketGradingSignal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketGradingSignalRequest struct via the builder pattern


### Return type

**GetVaultMarketGradingSignalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketHallLoan

> GetVaultMarketHallLoanResponse GetVaultMarketHallLoan(ctx).Execute()

Your Hall of Grails desk: this week's exhibition, the rules (maxLive loans, shownPerHolder per exhibition, repPoints per exhibition shown), reputation (exhibitions finished), your loans (status active, ending or ended; onShowNow, showingUntil, exhibitions, the next exhibition it fits, the Lent to the Hall badge) and the vault units you could lend (canLend, or the blocker in words, and the next exhibition each fits). Display only.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketHallLoan(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketHallLoan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketHallLoan`: GetVaultMarketHallLoanResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketHallLoan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketHallLoanRequest struct via the builder pattern


### Return type

**GetVaultMarketHallLoanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketHealth

> GetVaultMarketHealthResponse GetVaultMarketHealth(ctx).Execute()

Portfolio health of your vault holdings at Crossly marks: concentration by single item, category, grade and grader (shares, HHI on 0 to 10000, its band and effective number of holdings), value by liquidity tier, the stress grid (each category and everything down 10, 20 and 30 percent), value by mark quality, plain observations, and the held items a what-if can run on. Information only, not advice.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketHealth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketHealth`: GetVaultMarketHealthResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketHealthRequest struct via the builder pattern


### Return type

**GetVaultMarketHealthResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketHealthWhatIf

> GetVaultMarketHealthWhatIfResponse GetVaultMarketHealthWhatIf(ctx).BookKey(bookKey).MoveBps(moveBps).Execute()

What if one held item moved by moveBps (-10000 to 10000 basis points of its mark): the item, your vault total, the item and its category share of the vault, and the single-item HHI, before and after. A scenario on the mark, not a price you could sell at.

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
    bookKey := "bookKey_example" // string | 
    moveBps := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketHealthWhatIf(context.Background()).BookKey(bookKey).MoveBps(moveBps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketHealthWhatIf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketHealthWhatIf`: GetVaultMarketHealthWhatIfResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketHealthWhatIf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketHealthWhatIfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookKey** | **string** |  | 
 **moveBps** | **int32** |  | 

### Return type

**GetVaultMarketHealthWhatIfResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketHomeUnit

> GetVaultMarketHomeUnitResponse GetVaultMarketHomeUnit(ctx).Execute()

\"Show my vault\": the vault units you hold, the ones standing in your Crossly World home first. Each: vaultItemId, title, grade, imageUrl, bookKey (its market), inHome, visitors (visitors to your home see it). Also your home (slug, privacy, path), max units one home shows, shown, visitorsSee and a note on who sees what. Your home is made on first read. No prices.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketHomeUnit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketHomeUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketHomeUnit`: GetVaultMarketHomeUnitResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketHomeUnit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketHomeUnitRequest struct via the builder pattern


### Return type

**GetVaultMarketHomeUnitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketIndex

> GetVaultMarketIndexResponse GetVaultMarketIndex(ctx, slug).Execute()

One index with its daily closes (oldest first) and its level now.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketIndex(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketIndex`: GetVaultMarketIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketIndexResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketIndexCalendar

> GetVaultMarketIndexCalendarResponse GetVaultMarketIndexCalendar(ctx).Execute()

The next four quarterly reviews: reference day (2nd Friday of Mar/Jun/Sep/Dec), announcement day, last close under the old membership and the effective day; June and December also review the Blue-chip 50.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketIndexCalendar(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketIndexCalendar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketIndexCalendar`: GetVaultMarketIndexCalendarResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketIndexCalendar`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketIndexCalendarRequest struct via the builder pattern


### Return type

**GetVaultMarketIndexCalendarResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketIndexConsultation

> ListVaultMarketIndexConsultationsItem GetVaultMarketIndexConsultation(ctx, id).Execute()

One consultation with its proposal, versions, dates, outcome and the public comments.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketIndexConsultation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketIndexConsultation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketIndexConsultation`: ListVaultMarketIndexConsultationsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketIndexConsultation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketIndexConsultationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketIndexConsultationsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketIndexFeed

> GetVaultMarketIndexFeedResponse GetVaultMarketIndexFeed(ctx, slug).Execute()

The licensed index feed for one built-in index (needs VAULT_INDEX_DCM + VAULT_INDEX_FEED_SECRET): the level now (intraday from A/B marks, or the last close below 50% constituent coverage), the last close, the frozen 7-day TWAP settlement value and every constituent with its mark source, quality and trade count. Signed: `x-crossly-signature: t=<ms>,v1=<hex HMAC-SHA256(secret, t + \".\" + body)>`. Every read is logged for the licence. Crossly is the licensor and data feed only.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketIndexFeed(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketIndexFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketIndexFeed`: GetVaultMarketIndexFeedResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketIndexFeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketIndexFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketIndexFeedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketIndexFeedFunding

> GetVaultMarketIndexFeedFundingResponse GetVaultMarketIndexFeedFunding(ctx, slug).Execute()

Reference funding for a perpetual-style future on the index: your premium `premiumBps` (TWAP of (impact mid - index)/index) through F = P + clamp(I - P, ±0.05%) per 8 h, |F| capped at 0.5%/day, with the index level it is relative to. The partner DCM computes and settles funding; this is the reference math.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketIndexFeedFunding(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketIndexFeedFunding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketIndexFeedFunding`: GetVaultMarketIndexFeedFundingResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketIndexFeedFunding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketIndexFeedFundingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketIndexFeedFundingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketIndexFeedSettlement

> GetVaultMarketIndexFeedSettlementResponse GetVaultMarketIndexFeedSettlement(ctx, slug).Execute()

The frozen 7-day TWAP settlement series for a licensed index (written once a day, never restated), with the closes behind each value. Signed like the feed.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketIndexFeedSettlement(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketIndexFeedSettlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketIndexFeedSettlement`: GetVaultMarketIndexFeedSettlementResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketIndexFeedSettlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketIndexFeedSettlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketIndexFeedSettlementResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketIndexLevel

> GetVaultMarketIndexLevelResponse GetVaultMarketIndexLevel(ctx, slug).Execute()

An index's daily closes (up to 5 years), oldest first.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketIndexLevel(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketIndexLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketIndexLevel`: GetVaultMarketIndexLevelResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketIndexLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketIndexLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketIndexLevelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketIndexLicenceUsage

> GetVaultMarketIndexLicenceUsageResponse GetVaultMarketIndexLicenceUsage(ctx).Execute()

Your licensed index feed reads over `days`, per index and endpoint (licence-usage report).

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketIndexLicenceUsage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketIndexLicenceUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketIndexLicenceUsage`: GetVaultMarketIndexLicenceUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketIndexLicenceUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketIndexLicenceUsageRequest struct via the builder pattern


### Return type

**GetVaultMarketIndexLicenceUsageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketIndexMethodology

> GetVaultMarketIndexMethodologyResponse GetVaultMarketIndexMethodology(ctx).Execute()

The index method in machine-readable form: version (semver), inputs, weightings and caps, the family and its screens, the engines and their publication gate, revisions and the settlement series, the calendar, governance and limitations.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketIndexMethodology(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketIndexMethodology``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketIndexMethodology`: GetVaultMarketIndexMethodologyResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketIndexMethodology`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketIndexMethodologyRequest struct via the builder pattern


### Return type

**GetVaultMarketIndexMethodologyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketInstantExit

> GetVaultMarketInstantExitResponse GetVaultMarketInstantExit(ctx).Execute()

The Instant Exit quote for one vault unit you hold: the best resting bid on its book key (anyone's, dealers' included), its share of the mark, the bid depth, and whether it can be taken now.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketInstantExit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketInstantExit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketInstantExit`: GetVaultMarketInstantExitResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketInstantExit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketInstantExitRequest struct via the builder pattern


### Return type

**GetVaultMarketInstantExitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketJoint

> GetVaultMarketJointResponse GetVaultMarketJoint(ctx).Execute()

Your joint vaults and the invitations waiting for you, with the gate (open, missing env), the joint terms in words and what each right allows. A JointVaultsHome object, not a collection.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketJoint(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketJoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketJoint`: GetVaultMarketJointResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketJoint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketJointRequest struct via the builder pattern


### Return type

**GetVaultMarketJointResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketJointById

> GetVaultMarketJointByIdResponse GetVaultMarketJointById(ctx, id).Execute()

One joint vault (active members only): members and rights, the units in it at the mark with each owner of record, totals, recent joint trades, and the activity from the audit trail.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketJointById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketJointById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketJointById`: GetVaultMarketJointByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketJointById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketJointByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketJointByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketKeyStat

> GetVaultMarketKeyStatResponse GetVaultMarketKeyStat(ctx, bookKey).Execute()

Everything derived about one book key in one read: the Crossly Estimate, the liquidity score, risk statistics and the gem rate. A part that cannot be read is `{ error }`.

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
    bookKey := "bookKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketKeyStat(context.Background(), bookKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketKeyStat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketKeyStat`: GetVaultMarketKeyStatResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketKeyStat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketKeyStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketKeyStatResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketLearn

> GetVaultMarketLearnResponse GetVaultMarketLearn(ctx).Execute()

The Learn hub: each concept with whether you passed its quiz. Passing earns cosmetic pins only.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketLearn(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketLearn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketLearn`: GetVaultMarketLearnResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketLearn`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketLearnRequest struct via the builder pattern


### Return type

**GetVaultMarketLearnResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketLearnByConcept

> GetVaultMarketLearnByConceptResponse GetVaultMarketLearnByConcept(ctx, concept).Execute()

One Learn concept: its sections and quiz questions (no answers).

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
    concept := "concept_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketLearnByConcept(context.Background(), concept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketLearnByConcept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketLearnByConcept`: GetVaultMarketLearnByConceptResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketLearnByConcept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**concept** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketLearnByConceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketLearnByConceptResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketLendingElection

> GetVaultMarketLendingElectionResponse GetVaultMarketLendingElection(ctx).Execute()

Your lender election: `returnInKind` (shorts post 175% collateral, a 14-day buy-in, never cash-settled) and `allowTerm` (your units may go out on 30/90-day fixed-rate term loans, not recallable before term).

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketLendingElection(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketLendingElection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketLendingElection`: GetVaultMarketLendingElectionResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketLendingElection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketLendingElectionRequest struct via the builder pattern


### Return type

**GetVaultMarketLendingElectionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketLendingProgramme

> GetVaultMarketLendingProgrammeResponse GetVaultMarketLendingProgramme(ctx).Execute()

Your lender programme report: lendable units, units out, utilisation, fee revenue (accrued / paid), term loans and return-in-kind loans in force, and each loan's rate against the key's rate today.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketLendingProgramme(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketLendingProgramme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketLendingProgramme`: GetVaultMarketLendingProgrammeResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketLendingProgramme`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketLendingProgrammeRequest struct via the builder pattern


### Return type

**GetVaultMarketLendingProgrammeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketLendingStatement

> GetVaultMarketLendingStatementResponse GetVaultMarketLendingStatement(ctx).Execute()

Your daily collateral statement as a lender (today live, or `day`): per open loan the mark and its quality, fees owed to you, the short's collateral share at the custodian and its coverage against 105% of mark + fees. A unit without a mark usable for risk shows no coverage number.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketLendingStatement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketLendingStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketLendingStatement`: GetVaultMarketLendingStatementResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketLendingStatement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketLendingStatementRequest struct via the builder pattern


### Return type

**GetVaultMarketLendingStatementResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketLiquidity

> GetVaultMarketLiquidityResponse GetVaultMarketLiquidity(ctx, bookKey).Execute()

The liquidity score (0-100) of a book key and every component behind it: trades in 30/90 days, distinct counterparties, median days between trades, time-weighted quoted spread, two-sided share and depth within 5% (hourly book snapshots), Amihud and Roll. Each component is a percentile within the category. No score under the data threshold.

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
    bookKey := "bookKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketLiquidity(context.Background(), bookKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketLiquidity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketLiquidity`: GetVaultMarketLiquidityResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketLiquidity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketLiquidityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketLiquidityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketLoanQuote

> GetVaultMarketLoanQuoteResponse GetVaultMarketLoanQuote(ctx).Execute()

How much you can borrow against vault units (`vaultItemIds`, comma-separated; none = every stored unit you hold): each unit's tier, mark and limit, the pool's rate and yours.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketLoanQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketLoanQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketLoanQuote`: GetVaultMarketLoanQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketLoanQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketLoanQuoteRequest struct via the builder pattern


### Return type

**GetVaultMarketLoanQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketLoanUnitSaleFloor

> GetVaultMarketLoanUnitSaleFloorResponse GetVaultMarketLoanUnitSaleFloor(ctx, id, vaultItemId).Execute()

The lowest price a pledged unit can be sold at to repay its loan: after the sale (less the vault fee and storage) what is still owed must fit inside the other units' limit.

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
    id := "id_example" // string | 
    vaultItemId := "vaultItemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketLoanUnitSaleFloor(context.Background(), id, vaultItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketLoanUnitSaleFloor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketLoanUnitSaleFloor`: GetVaultMarketLoanUnitSaleFloorResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketLoanUnitSaleFloor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**vaultItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketLoanUnitSaleFloorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetVaultMarketLoanUnitSaleFloorResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketLocate

> GetVaultMarketLocateResponse GetVaultMarketLocate(ctx).Execute()

A locate quote: how many units of `key` (a book key) you could borrow now (up to `qty`), the indicative rate (APR bps; `termDays` 30 or 90 for a fixed-rate term loan), and the collateral they need (150%, or 175% for units whose lenders elected return in kind only). Holds nothing; POST to hold.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketLocate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketLocate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketLocate`: GetVaultMarketLocateResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketLocate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketLocateRequest struct via the builder pattern


### Return type

**GetVaultMarketLocateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketMark

> GetVaultMarketMarkResponse GetVaultMarketMark(ctx, bookKey).Execute()

One book key mark with everything behind it: the price in cents, its quality (A trade-anchored, B firm two-sided, C indicative, none), its confidence in bps and the band in cents, when it was computed, whether risk may act on it (judged for margin), the related key it was implied from (impliedFrom: a grade, the model median of the other sizes, or an index; null when none), every stored input, and the learned relations it may be implied from.

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
    bookKey := "bookKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketMark(context.Background(), bookKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketMark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketMark`: GetVaultMarketMarkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketMark`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketMarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketMarkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketMarket

> GetVaultMarketMarketResponse GetVaultMarketMarket(ctx).Execute()

One book key (variantId:condition:gradeKey): its mark and history, lending depth and borrow rate, open interest and live auctions.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketMarket(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketMarket`: GetVaultMarketMarketResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketMarket`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketMarketRequest struct via the builder pattern


### Return type

**GetVaultMarketMarketResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketMarketModel

> GetVaultMarketMarketModelResponse GetVaultMarketMarketModel(ctx, bookKey).Execute()

How a book key trades: its liquidity tier (A continuous, B call auction every 15 minutes, C daily call, D listings only), the call collecting now with its indicative price and volume and nominal close, the next call, the LULD band and any limit state, pause or halt.

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
    bookKey := "bookKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketMarketModel(context.Background(), bookKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketMarketModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketMarketModel`: GetVaultMarketMarketModelResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketMarketModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketMarketModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketMarketModelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketMm

> GetVaultMarketMmResponse GetVaultMarketMm(ctx).Execute()

Your market-maker programmes, the last 30 days of obligations, rebates earned (paid only from fees already taken), and the rules.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketMm(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketMm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketMm`: GetVaultMarketMmResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketMm`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketMmRequest struct via the builder pattern


### Return type

**GetVaultMarketMmResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketMmKey

> GetVaultMarketMmKeyResponse GetVaultMarketMmKey(ctx).Execute()

A key's designated market maker (by the name it chose) and how many LPs quote it.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketMmKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketMmKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketMmKey`: GetVaultMarketMmKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketMmKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketMmKeyRequest struct via the builder pattern


### Return type

**GetVaultMarketMmKeyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketMmMmp

> GetVaultMarketMmMmpResponse GetVaultMarketMmMmp(ctx).Execute()

Your market-maker protection: the fill trigger, the freeze, and whether you are frozen now.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketMmMmp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketMmMmp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketMmMmp`: GetVaultMarketMmMmpResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketMmMmp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketMmMmpRequest struct via the builder pattern


### Return type

**GetVaultMarketMmMmpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketMode

> GetVaultMarketModeResponse GetVaultMarketMode(ctx).Execute()

Your Vault Markets mode: simple (buy, sell and what it is worth) or pro (the full terminal), whether you chose it or it was set for you (new people start in simple; people already trading start in pro), and when.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketMode(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketMode`: GetVaultMarketModeResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketMode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketModeRequest struct via the builder pattern


### Return type

**GetVaultMarketModeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketNav

> GetVaultMarketNavResponse GetVaultMarketNav(ctx).Execute()

Mark NAV, bid NAV (the resting bids) and liquidation NAV (the bids at 20% of ADV for 10 days), holding by holding, each with its ASC 820 level, mark quality and source, and your administrator's price when uploaded.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketNav(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketNav``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketNav`: GetVaultMarketNavResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketNav`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketNavRequest struct via the builder pattern


### Return type

**GetVaultMarketNavResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketNavFile

> string GetVaultMarketNavFile(ctx, format).Execute()

The NAV as a csv, pdf or json file.

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
    format := "format_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketNavFile(context.Background(), format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketNavFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketNavFile`: string
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketNavFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketNavFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketOpening

> ListVaultMarketOpeningItem GetVaultMarketOpening(ctx, slug).Execute()

One opening key.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketOpening(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketOpening``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketOpening`: ListVaultMarketOpeningItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketOpening`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketOpeningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketOpeningItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketOracle

> GetVaultMarketOracleResponse GetVaultMarketOracle(ctx).Execute()

Crossly as an oracle publisher (needs CROSSLY_ORACLE_PUBLISHER_KEY + CROSSLY_ORACLE_NETWORK): the network (pyth or switchboard), the ed25519 key every value is signed with, each published index feed with its latest attempt, and 24 h counts of published, skipped and failed. Only built-in indexes priced from A/B marks are published; Crossly never reads its own feed back.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketOracle(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketOracle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketOracle`: GetVaultMarketOracleResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketOracle`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketOracleRequest struct via the builder pattern


### Return type

**GetVaultMarketOracleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketOracleBySlug

> GetVaultMarketOracleBySlugResponse GetVaultMarketOracleBySlug(ctx, slug).Execute()

The latest published oracle value for an index: level, fixed-point price and confidence at the feed exponent, confidence in bps, staleness in seconds, coverage, the exact signed message and its ed25519 signature with the signer key, and whether it is fresh.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketOracleBySlug(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketOracleBySlug``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketOracleBySlug`: GetVaultMarketOracleBySlugResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketOracleBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketOracleBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketOracleBySlugResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketOrderHandling

> GetVaultMarketOrderHandlingResponse GetVaultMarketOrderHandling(ctx).Execute()

How orders are handled (the 606/607-style disclosure): one book, self-trade prevention, protected market orders, liquidations by auction, halts, dealers and fees.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketOrderHandling(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketOrderHandling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketOrderHandling`: GetVaultMarketOrderHandlingResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketOrderHandling`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketOrderHandlingRequest struct via the builder pattern


### Return type

**GetVaultMarketOrderHandlingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketP2pDashboard

> GetVaultMarketP2pDashboardResponse GetVaultMarketP2pDashboard(ctx).Execute()

Peer-to-peer vault loans at a glance: what you owe as a borrower and are owed as a lender (principal + interest to the second), daily interest, called loans, your open offers, the gates and the rules (auction hours, rate ceiling, size bounds, call grace).

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketP2pDashboard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketP2pDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketP2pDashboard`: GetVaultMarketP2pDashboardResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketP2pDashboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketP2pDashboardRequest struct via the builder pattern


### Return type

**GetVaultMarketP2pDashboardResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketP2pLoan

> GetVaultMarketP2pLoanResponse GetVaultMarketP2pLoan(ctx, id).Execute()

One P2P loan with its money log (fund, repay, refinance). A loan in a running refinancing auction can be read by any lender, without the log.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketP2pLoan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketP2pLoan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketP2pLoan`: GetVaultMarketP2pLoanResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketP2pLoan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketP2pLoanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketP2pLoanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPaper

> GetVaultMarketPaperResponse GetVaultMarketPaper(ctx).Execute()

Your paper trading account (play money only, never real or convertible; opened at $10,000 the first time): play cash, buying power, paper positions at the real marks, P&L, working paper orders, recent paper fills and this week on the leaderboard.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPaper(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPaper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPaper`: GetVaultMarketPaperResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPaper`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPaperRequest struct via the builder pattern


### Return type

**GetVaultMarketPaperResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPaperQuote

> GetVaultMarketPaperQuoteResponse GetVaultMarketPaperQuote(ctx).Execute()

What a paper order would fill against the real book now, without placing it: units fillable inside the limit (a market order uses the 5% protection band around the mark), average price, simulated fee, and the queue ahead of any rest.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPaperQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPaperQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPaperQuote`: GetVaultMarketPaperQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPaperQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPaperQuoteRequest struct via the builder pattern


### Return type

**GetVaultMarketPaperQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPartnerMe

> GetVaultMarketPartnerMeResponse GetVaultMarketPartnerMe(ctx).Execute()

Your Vault Drop-off Partnership (null when you never applied): status, partner code, fee share and cap, totals and batches.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPartnerMe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPartnerMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPartnerMe`: GetVaultMarketPartnerMeResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPartnerMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPartnerMeRequest struct via the builder pattern


### Return type

**GetVaultMarketPartnerMeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPool

> GetVaultMarketPoolResponse GetVaultMarketPool(ctx).Execute()

The CBX market staking pool: size, value, what it lends, utilisation, rate, insurance fund.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPool(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPool`: GetVaultMarketPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPool`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPoolRequest struct via the builder pattern


### Return type

**GetVaultMarketPoolResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPopReport

> GetVaultMarketPopReportResponse GetVaultMarketPopReport(ctx).BookKey(bookKey).Execute()

Grading population report for one market (`bookKey`): focus (on a graded market: its grade, pop = copies graded at exactly that grade, higher = copies graded better by the same grader, total, asOf and source), graders (each grader ladder for the card, best grade first, from one report: the newest, the population feed over a hand entry on the same date), provider.connected, and status (ok, not_synced, no_report or no_source) with explain, which says in plain words what is missing. source is provider (the population feed) or ops (entered by Crossly from the grader published report). Counts only, never a price.

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
    bookKey := "bookKey_example" // string | variantId:condition:gradeKey, grade \"-\" when ungraded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPopReport(context.Background()).BookKey(bookKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPopReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPopReport`: GetVaultMarketPopReportResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPopReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPopReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookKey** | **string** | variantId:condition:gradeKey, grade \&quot;-\&quot; when ungraded. | 

### Return type

**GetVaultMarketPopReportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPortfolio

> GetVaultMarketPortfolioResponse GetVaultMarketPortfolio(ctx).Execute()

Your vault holdings at the market price, with cost and P&L, and units you have out on loan.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPortfolio(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPortfolio``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPortfolio`: GetVaultMarketPortfolioResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPortfolio`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPortfolioRequest struct via the builder pattern


### Return type

**GetVaultMarketPortfolioResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPortfolioHistory

> GetVaultMarketPortfolioHistoryResponse GetVaultMarketPortfolioHistory(ctx).Execute()

Your Vault Markets value by day at the marks (vault units, units on loan, positions) and today's return against yesterday's close.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPortfolioHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPortfolioHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPortfolioHistory`: GetVaultMarketPortfolioHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPortfolioHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPortfolioHistoryRequest struct via the builder pattern


### Return type

**GetVaultMarketPortfolioHistoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPortfolioMargin

> GetVaultMarketPortfolioMarginResponse GetVaultMarketPortfolioMargin(ctx).Execute()

Portfolio margin for a firm sub-account (`subaccountId`): eligibility with every reason (ECP, KYB, minimum equity $500k or $100k with evidenced intraday monitoring), whether it is enrolled, the latest requirement, equity, excess and state (healthy, call or buy_in), the breakdown by category and item, the partner index futures counted, and the policy. Before enrolling, a live preview.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPortfolioMargin(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPortfolioMargin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPortfolioMargin`: GetVaultMarketPortfolioMarginResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPortfolioMargin`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPortfolioMarginRequest struct via the builder pattern


### Return type

**GetVaultMarketPortfolioMarginResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPosition

> GetVaultMarketPositionResponse GetVaultMarketPosition(ctx, id).Execute()

One position with its loans, auctions and custodian instruction log.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPosition(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPosition`: GetVaultMarketPositionResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPosition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketPositionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPositionExit

> GetVaultMarketPositionExitResponse GetVaultMarketPositionExit(ctx, id).Execute()

A position's take-profit / stop-loss, if set.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPositionExit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPositionExit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPositionExit`: GetVaultMarketPositionExitResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPositionExit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPositionExitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketPositionExitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPositionRisk

> ListVaultMarketPositionsRiskItem GetVaultMarketPositionRisk(ctx, id).Execute()

One position's call and liquidation marks.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPositionRisk(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPositionRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPositionRisk`: ListVaultMarketPositionsRiskItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPositionRisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPositionRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketPositionsRiskItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPress

> GetVaultMarketPressResponse GetVaultMarketPress(ctx).Execute()

The press kit: each built-in index with level, quality, n, CSV link and a \"cite as\" line, plus the methodology links.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPress`: GetVaultMarketPressResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPressRequest struct via the builder pattern


### Return type

**GetVaultMarketPressResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketPriceWidget

> ListVaultMarketPriceWidgetsItem GetVaultMarketPriceWidget(ctx, id).Execute()

One of your price widgets, with its iframe and script-tag snippets and URLs.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketPriceWidget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketPriceWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketPriceWidget`: ListVaultMarketPriceWidgetsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketPriceWidget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketPriceWidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketPriceWidgetsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketProDepth

> GetVaultMarketProDepthResponse GetVaultMarketProDepth(ctx).Execute()

The L2 depth ladder: price levels and units per side, best first; asks show how many units are authenticated vault stock.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketProDepth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketProDepth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketProDepth`: GetVaultMarketProDepthResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketProDepth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketProDepthRequest struct via the builder pattern


### Return type

**GetVaultMarketProDepthResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketProOrder

> GetVaultMarketProOrderResponse GetVaultMarketProOrder(ctx, id).Execute()

One order with its full event log.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketProOrder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketProOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketProOrder`: GetVaultMarketProOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketProOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketProOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketProOrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketProSnapshot

> GetVaultMarketProSnapshotResponse GetVaultMarketProSnapshot(ctx).Execute()

A book key's depth (20 levels a side), last 50 trades, 24 h stats, mark and any halt, in one call.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketProSnapshot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketProSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketProSnapshot`: GetVaultMarketProSnapshotResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketProSnapshot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketProSnapshotRequest struct via the builder pattern


### Return type

**GetVaultMarketProSnapshotResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketProStat

> GetVaultMarketProStatResponse GetVaultMarketProStat(ctx).Execute()

24 h open, high, low, last, volume and change on a book key.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketProStat(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketProStat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketProStat`: GetVaultMarketProStatResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketProStat`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketProStatRequest struct via the builder pattern


### Return type

**GetVaultMarketProStatResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketProStream

> string GetVaultMarketProStream(ctx).Execute()

A long-lived text/event-stream for a book key: `snapshot` first, then `depth`, `trades`, `mark` and `halt` events as they change.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketProStream(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketProStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketProStream`: string
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketProStream`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketProStreamRequest struct via the builder pattern


### Return type

**string**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketProfile

> GetVaultMarketProfileResponse GetVaultMarketProfile(ctx).Execute()

Your public vault profile switch: on or off, whether followers see new buys, your profile URL and follower count.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketProfile`: GetVaultMarketProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketProfileRequest struct via the builder pattern


### Return type

**GetVaultMarketProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketProfileByHandle

> GetVaultMarketProfileByHandleResponse GetVaultMarketProfileByHandle(ctx, handle).Execute()

Someone's public vault: each market's weight (bps of the marked value), followers, whether you follow, their recent buys (no prices).

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketProfileByHandle(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketProfileByHandle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketProfileByHandle`: GetVaultMarketProfileByHandleResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketProfileByHandle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketProfileByHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketProfileByHandleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketProofOfReserve

> GetVaultMarketProofOfReserveResponse GetVaultMarketProofOfReserve(ctx).Execute()

Proof of reserves: the latest daily Merkle root over every unit in custody (leaf count, stored / listed / grading / withdrawing, pledged, on loan, open loans, receipts, insured value), its Solana memo anchor when VAULT_POR_ANCHOR_KEY is set, the last `days` roots, and the exact hashing method.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketProofOfReserve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketProofOfReserve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketProofOfReserve`: GetVaultMarketProofOfReserveResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketProofOfReserve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketProofOfReserveRequest struct via the builder pattern


### Return type

**GetVaultMarketProofOfReserveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketProofOfReserveByDay

> GetVaultMarketProofOfReserveByDayResponse GetVaultMarketProofOfReserveByDay(ctx, day).Execute()

One day's proof-of-reserves snapshot (root, counts, anchor).

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
    day := "day_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketProofOfReserveByDay(context.Background(), day).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketProofOfReserveByDay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketProofOfReserveByDay`: GetVaultMarketProofOfReserveByDayResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketProofOfReserveByDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**day** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketProofOfReserveByDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketProofOfReserveByDayResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketProofOfReserveMe

> GetVaultMarketProofOfReserveMeResponse GetVaultMarketProofOfReserveMe(ctx).Execute()

Your inclusion proofs for a day (latest by default): each of your units' leaf fields, salt, owner commitment, leaf hash and sibling path. Verify with your OWN user id: commitment = SHA-256(JSON[\"crossly-por-owner-v1\", userId, salt]); then the leaf; then climb the path to the published root.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketProofOfReserveMe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketProofOfReserveMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketProofOfReserveMe`: GetVaultMarketProofOfReserveMeResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketProofOfReserveMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketProofOfReserveMeRequest struct via the builder pattern


### Return type

**GetVaultMarketProofOfReserveMeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketRate

> GetVaultMarketRateResponse GetVaultMarketRate(ctx).Execute()

Daily borrow APR (lent units) and pool APR on a book key, with the pool as a whole, oldest first.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketRate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketRate`: GetVaultMarketRateResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketRate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketRateRequest struct via the builder pattern


### Return type

**GetVaultMarketRateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketReport

> GetVaultMarketReportResponse GetVaultMarketReport(ctx).Execute()

The quarterly Crossly Vault Report (`quarter` like 2026-q3; default the latest): index levels, top movers, most-lent markets, grade premiums. Book data only.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketReport`: GetVaultMarketReportResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketReportRequest struct via the builder pattern


### Return type

**GetVaultMarketReportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketRfq

> GetVaultMarketRfqResponse GetVaultMarketRfq(ctx, id).Execute()

One quote request: the requester sees every answer (anonymous); a responder sees the request and their own answer.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketRfq(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketRfq``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketRfq`: GetVaultMarketRfqResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketRfq`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketRfqRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketRfqResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketRiskReport

> GetVaultMarketRiskReportResponse GetVaultMarketRiskReport(ctx).Execute()

Historical VaR 95/99 1d/10d and Bangia LVaR (model figures: labelled, with n and a band, null below the data threshold), stress scenarios (category −30%, CBX −50%, books thin to 25%, index drawdowns), concentration and counterparty exposure.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketRiskReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketRiskReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketRiskReport`: GetVaultMarketRiskReportResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketRiskReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketRiskReportRequest struct via the builder pattern


### Return type

**GetVaultMarketRiskReportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketRiskStat

> GetVaultMarketRiskStatResponse GetVaultMarketRiskStat(ctx, bookKey).Execute()

Risk statistics for a book key (weekly returns from A-quality marks only) or an index slug (its closes): annualised volatility raw and Geltner-desmoothed (labelled), max drawdown and beta to the Composite. Needs 26 weekly returns, else no numbers.

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
    bookKey := "bookKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketRiskStat(context.Background(), bookKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketRiskStat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketRiskStat`: GetVaultMarketRiskStatResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketRiskStat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketRiskStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketRiskStatResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketSetGoal

> GetVaultMarketSetGoalResponse GetVaultMarketSetGoal(ctx, id).Execute()

One set goal: every piece in the set order, held (with the vault unit) or missing (with the cheapest marked book key that meets the target and its mark), the conditions and grades seen, and the working basket bid with its legs.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketSetGoal(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketSetGoal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketSetGoal`: GetVaultMarketSetGoalResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketSetGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketSetGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketSetGoalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketSetGoalBasketPlan

> GetVaultMarketSetGoalBasketPlanResponse GetVaultMarketSetGoalBasketPlan(ctx, id).Execute()

What a basket bid would place now, without placing it: one leg per missing piece (cheapest first, the mark plus maxBpsOverMark as the limit), the pieces skipped and why, and the most it can cost (limits plus the worst-case buyer fee), never more than budgetCents.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketSetGoalBasketPlan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketSetGoalBasketPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketSetGoalBasketPlan`: GetVaultMarketSetGoalBasketPlanResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketSetGoalBasketPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketSetGoalBasketPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketSetGoalBasketPlanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketShareCard

> ListVaultMarketShareCardsItem GetVaultMarketShareCard(ctx, id).Execute()

One share card as the public sees it, with its image and link paths.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketShareCard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketShareCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketShareCard`: ListVaultMarketShareCardsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketShareCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketShareCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketShareCardsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketStatement

> GetVaultMarketStatementResponse GetVaultMarketStatement(ctx).Execute()

The last 24 months with whether each statement is issued (frozen, with its sha256).

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketStatement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketStatement`: GetVaultMarketStatementResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketStatement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketStatementRequest struct via the builder pattern


### Return type

**GetVaultMarketStatementResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketStatementByMonth

> GetVaultMarketStatementByMonthResponse GetVaultMarketStatementByMonth(ctx, month).Execute()

A monthly statement (`month` like 2026-10): units held at month end, loans, pledges, positions, fills, fees, storage, and cash quoted from the custodian and labelled as its record. Issued (frozen) once the month ends.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketStatementByMonth(context.Background(), month).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketStatementByMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketStatementByMonth`: GetVaultMarketStatementByMonthResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketStatementByMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**month** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketStatementByMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketStatementByMonthResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketStatementFile

> string GetVaultMarketStatementFile(ctx, month, format).Execute()

A monthly statement as a csv, pdf or json file.

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
    format := "format_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketStatementFile(context.Background(), month, format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketStatementFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketStatementFile`: string
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketStatementFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**month** | **string** |  | 
**format** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketStatementFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketSupply

> GetVaultMarketSupplyResponse GetVaultMarketSupply(ctx).Execute()

A book key's units in the vault and authentications passed, by day (oldest first, today live).

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketSupply(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketSupply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketSupply`: GetVaultMarketSupplyResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketSupply`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketSupplyRequest struct via the builder pattern


### Return type

**GetVaultMarketSupplyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketSupplyBatch

> GetVaultMarketSupplyBatchResponse GetVaultMarketSupplyBatch(ctx, id).Execute()

A box sent to the vault: its prepaid label and tracking, or why it has no label.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketSupplyBatch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketSupplyBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketSupplyBatch`: GetVaultMarketSupplyBatchResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketSupplyBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketSupplyBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketSupplyBatchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketSupplyRoundUp

> GetVaultMarketSupplyRoundUpResponse GetVaultMarketSupplyRoundUp(ctx).Execute()

Your seller round-up: the share of each payout that sets a vault buy's budget, the budget so far and what it is doing. `roundUp` is null when none is set.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketSupplyRoundUp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketSupplyRoundUp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketSupplyRoundUp`: GetVaultMarketSupplyRoundUpResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketSupplyRoundUp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketSupplyRoundUpRequest struct via the builder pattern


### Return type

**GetVaultMarketSupplyRoundUpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketTax

> GetVaultMarketTaxResponse GetVaultMarketTax(ctx).Execute()

Your tax year on the vault markets: units bought and sold with cost, proceeds, gain and term; lending fees, staking interest (estimate), market-maker rewards; closed positions. Links to the CSV and PDF.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketTax(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketTax`: GetVaultMarketTaxResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketTax`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketTaxRequest struct via the builder pattern


### Return type

**GetVaultMarketTaxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketTaxCsv

> string GetVaultMarketTaxCsv(ctx).Execute()

The tax year as a CSV: one row per sale, purchase, gift, income line and closed position.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketTaxCsv(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketTaxCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketTaxCsv`: string
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketTaxCsv`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketTaxCsvRequest struct via the builder pattern


### Return type

**string**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketTaxLossSuggestion

> GetVaultMarketTaxLossSuggestionResponse GetVaultMarketTaxLossSuggestion(ctx).Execute()

Tax-loss suggestions: the units you hold whose mark is below your cost, each with its lot (acquired when and how, cost in cents), short or long term if sold today and the day a short lot turns long, the estimated loss at the mark before selling fees, mark quality, what stops a sale, and a link to the ordinary sell ticket. With totals by term, the gains realised so far this year, and a not-tax-advice disclosure. Never sells.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketTaxLossSuggestion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketTaxLossSuggestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketTaxLossSuggestion`: GetVaultMarketTaxLossSuggestionResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketTaxLossSuggestion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketTaxLossSuggestionRequest struct via the builder pattern


### Return type

**GetVaultMarketTaxLossSuggestionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketTaxPdf

> *os.File GetVaultMarketTaxPdf(ctx).Execute()

The tax year as a PDF summary.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketTaxPdf(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketTaxPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketTaxPdf`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketTaxPdf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketTaxPdfRequest struct via the builder pattern


### Return type

***os.File**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketTca

> GetVaultMarketTcaResponse GetVaultMarketTca(ctx).Execute()

Your pro orders in a window (`from`, `to`; default 90 days) with arrival mark and mid, average fill, implementation shortfall, slippage, fill rate and time to fill.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketTca(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketTca``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketTca`: GetVaultMarketTcaResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketTca`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketTcaRequest struct via the builder pattern


### Return type

**GetVaultMarketTcaResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketTcaByOrderId

> GetVaultMarketTcaByOrderIdResponse GetVaultMarketTcaByOrderId(ctx, orderId).Execute()

One order's TCA, including slippage against the VWAP of eligible trades while it worked.

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
    orderId := "orderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketTcaByOrderId(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketTcaByOrderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketTcaByOrderId`: GetVaultMarketTcaByOrderIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketTcaByOrderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketTcaByOrderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketTcaByOrderIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketTod

> GetVaultMarketTodResponse GetVaultMarketTod(ctx).Execute()

Your transfer-on-death setup: the gate and terms, whether you accepted them, your beneficiaries and shares, units with their designations and who the plan gives them to today, and open claims on your account (which you can stop). A TodOverview object, not a collection.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketTod(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketTod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketTod`: GetVaultMarketTodResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketTod`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketTodRequest struct via the builder pattern


### Return type

**GetVaultMarketTodResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketTodMyClaim

> GetVaultMarketTodMyClaimResponse GetVaultMarketTodMyClaim(ctx).Execute()

Claims you filed as a beneficiary: status in words, documents sent (names only), the waiting period, and units moved or held. A TodMyClaims object, not a collection.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketTodMyClaim(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketTodMyClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketTodMyClaim`: GetVaultMarketTodMyClaimResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketTodMyClaim`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketTodMyClaimRequest struct via the builder pattern


### Return type

**GetVaultMarketTodMyClaimResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketUnitAr

> GetVaultMarketUnitArResponse GetVaultMarketUnitAr(ctx, id).Execute()

A vault unit as an AR model, for \"See it in your room\" (the web's /ar/vault-<id> page hands it to Scene Viewer, WebXR or Quick Look): title, model (a path under /models, without .glb, at real size), scanUrl (the unit's own captured mesh, when it has one), imageUrl, art, tint, scale, bookKey, backPath and mine. Your own units always; somebody else's only while they show it to visitors in a home you may enter (else 404).

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketUnitAr(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketUnitAr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketUnitAr`: GetVaultMarketUnitArResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketUnitAr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketUnitArRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketUnitArResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketUnitEvidence

> GetVaultMarketUnitEvidenceResponse GetVaultMarketUnitEvidence(ctx, id).Execute()

Evidence on a vault unit:its 3D scans (model, texture, listing photos) and its grade report (grader, grade, cert, verified against the grader record, hub intake photos). No owner identity.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketUnitEvidence(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketUnitEvidence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketUnitEvidence`: GetVaultMarketUnitEvidenceResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketUnitEvidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketUnitEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketUnitEvidenceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketUnitQuote

> GetVaultMarketUnitQuoteResponse GetVaultMarketUnitQuote(ctx, id).Execute()

The market around a vault unit: its mark and source, the best bid that could take it, the Instant Exit price (the best bid) and a pack pull's own sell-back credit.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketUnitQuote(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketUnitQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketUnitQuote`: GetVaultMarketUnitQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketUnitQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketUnitQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketUnitQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketUnitStory

> GetVaultMarketUnitStoryResponse GetVaultMarketUnitStory(ctx, id).Execute()

One vault unit's story, oldest first: intake (with scan and grade evidence when it has some), grades, owners, sales at their prices, loans, pledges, Hall of Grails exhibitions, receipt-token exports and redemptions, and each warehouse-receipt hash-chain step (re-verified on read: chain.ok). People are named only while their public vault profile is on; others are 'A collector'.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketUnitStory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketUnitStory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketUnitStory`: GetVaultMarketUnitStoryResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketUnitStory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketUnitStoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketUnitStoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketValuationStatement

> GetVaultMarketValuationStatementResponse GetVaultMarketValuationStatement(ctx, id).Execute()

One statement you issued: every unit, the totals by mark quality, the methodology, its sha256, its audit-chain seal and the verify link.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketValuationStatement(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketValuationStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketValuationStatement`: GetVaultMarketValuationStatementResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketValuationStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketValuationStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketValuationStatementResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketValuationStatementPdf

> *os.File GetVaultMarketValuationStatementPdf(ctx, id).Execute()

One statement you issued as a PDF.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketValuationStatementPdf(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketValuationStatementPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketValuationStatementPdf`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketValuationStatementPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketValuationStatementPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

***os.File**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketValuationStatementPreview

> GetVaultMarketValuationStatementPreviewResponse GetVaultMarketValuationStatementPreview(ctx).Execute()

An insurance or appraisal statement for a day (purpose, asOf YYYY-MM-DD, default today), built but not issued: every unit you held that day at its mark, with quality, last trade, grade and cert, and the total. No id, no hash.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketValuationStatementPreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketValuationStatementPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketValuationStatementPreview`: GetVaultMarketValuationStatementPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketValuationStatementPreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketValuationStatementPreviewRequest struct via the builder pattern


### Return type

**GetVaultMarketValuationStatementPreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketValuationStatementVerify

> GetVaultMarketValuationStatementVerifyResponse GetVaultMarketValuationStatementVerify(ctx, id).Execute()

Check any issued statement by its id, as an insurer would: purpose, date, totals by mark quality and the sha256; intact (the stored record still hashes to it), seal.ok (its audit-chain entry still carries it) and, with ?hash=, whether yours matches. Never the holder or the units.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketValuationStatementVerify(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketValuationStatementVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketValuationStatementVerify`: GetVaultMarketValuationStatementVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketValuationStatementVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketValuationStatementVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketValuationStatementVerifyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketValue

> GetVaultMarketValueResponse GetVaultMarketValue(ctx).Execute()

A catalogue item's vault book keys (`?skuId=`, optional `variantId`): each key's display mark with quality and n, units in the vault, the Instant Exit price a holder would get now, and whether it can be shipped in (ungraded keys).

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketValue(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketValue`: GetVaultMarketValueResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketValue`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketValueRequest struct via the builder pattern


### Return type

**GetVaultMarketValueResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketVaultSuggestionItem

> GetVaultMarketVaultSuggestionItemResponse GetVaultMarketVaultSuggestionItem(ctx, id).Execute()

One inventory item: its vault suggestion, or in plain words why there is none (`status`: suggested, not_found, no_stock, not_matched, in_vault or no_market).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketVaultSuggestionItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketVaultSuggestionItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketVaultSuggestionItem`: GetVaultMarketVaultSuggestionItemResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketVaultSuggestionItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketVaultSuggestionItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketVaultSuggestionItemResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketWidgetIndex

> GetVaultMarketWidgetIndexResponse GetVaultMarketWidgetIndex(ctx, slug).Execute()

What the <crossly-index> widget shows: level, closes, quality, n, as-of, or `insufficient`.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketWidgetIndex(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketWidgetIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketWidgetIndex`: GetVaultMarketWidgetIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketWidgetIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketWidgetIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketWidgetIndexResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketWidgetMark

> GetVaultMarketWidgetMarkResponse GetVaultMarketWidgetMark(ctx).Execute()

What the <crossly-mark> widget shows for `bookKey`: mark, quality, trades over 30 days, as-of, sparkline, or `insufficient`.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketWidgetMark(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketWidgetMark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketWidgetMark`: GetVaultMarketWidgetMarkResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketWidgetMark`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketWidgetMarkRequest struct via the builder pattern


### Return type

**GetVaultMarketWidgetMarkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketWidgetSummary

> GetVaultMarketWidgetSummaryResponse GetVaultMarketWidgetSummary(ctx).Execute()

The home-screen widget summary: portfolio value at the mark with its move since the previous marks (cents and bps) and the share resting on reliable A/B marks, the three biggest movers among what you hold (the market movers when you hold nothing; `movers.source` says which), and one watchlist key (`?key=` a book key on your watchlist; else your newest watch) with a 14-day sparkline. Every price is from the Crossly order book and carries its mark quality (A, B, C or none). Cached for five minutes.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketWidgetSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketWidgetSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketWidgetSummary`: GetVaultMarketWidgetSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketWidgetSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketWidgetSummaryRequest struct via the builder pattern


### Return type

**GetVaultMarketWidgetSummaryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketWorldFloor

> GetVaultMarketWorldFloorResponse GetVaultMarketWorldFloor(ctx).Execute()

The Trading Floor in Crossly World: trades (the newest trades on every vault book in the last day, at most 40, newest first: bookKey, name, priceCents, priceText, side -- buy when a buyer took an ask, sell when a seller hit a bid, cross for an auction -- and ageS), tickers (the big board: the most-traded markets, each with its last trade price, the change from the last trade before today in basis points, trades today and the display mark with its quality), bell (Market Open: ringing for ten minutes after a new key opens, with the key and its market link, and what opens next) and trades24h. Every price is a trade on the Crossly order book; it never shows who traded. Nothing can be traded in the room.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketWorldFloor(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketWorldFloor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketWorldFloor`: GetVaultMarketWorldFloorResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketWorldFloor`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketWorldFloorRequest struct via the builder pattern


### Return type

**GetVaultMarketWorldFloorResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketWorldLounge

> GetVaultMarketWorldLoungeResponse GetVaultMarketWorldLounge(ctx).Execute()

The Vault Holders Lounge: the rule to join (held value at reliable marks, or a unit count), whether you may join now and what you hold toward it. Anyone may look in; only holders are seen there. No perks, nothing to trade.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketWorldLounge(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketWorldLounge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketWorldLounge`: GetVaultMarketWorldLoungeResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketWorldLounge`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketWorldLoungeRequest struct via the builder pattern


### Return type

**GetVaultMarketWorldLoungeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketWorldPlaceVaultGate

> GetVaultMarketWorldPlaceVaultGateResponse GetVaultMarketWorldPlaceVaultGate(ctx, place).Execute()

The vault-holder door on a Crossly World room or Space, by its presence slug (a room share slug, or space-<slug>): the rule (minValueCents held at reliable A/B marks and/or minUnits units, mode any or all, optionally one product, bullion family or category), the rule in words, and whether you may join now with what you hold toward it. gate is null when the place has no vault door. Anyone may look in; only holders join.

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
    place := "place_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketWorldPlaceVaultGate(context.Background(), place).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketWorldPlaceVaultGate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketWorldPlaceVaultGate`: GetVaultMarketWorldPlaceVaultGateResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketWorldPlaceVaultGate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**place** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketWorldPlaceVaultGateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketWorldPlaceVaultGateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketWorth

> GetVaultMarketWorthResponse GetVaultMarketWorth(ctx).Execute()

What one book key (variantId:condition:gradeKey, `bookKey`) is worth: the mark in cents with its quality (A from trades, B from firm quotes, C indicative, none) and band, what Instant Exit would pay for one unit now from the best resting bid (yours left out), and your units: how many you hold, how many are free to sell, their value at the mark and what selling every free one now would fetch from today’s bids. Plain-words `headline` and `lines`. Read only.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketWorth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketWorth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketWorth`: GetVaultMarketWorthResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketWorth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketWorthRequest struct via the builder pattern


### Return type

**GetVaultMarketWorthResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketWorthVault

> GetVaultMarketWorthVaultResponse GetVaultMarketWorthVault(ctx).Execute()

What your whole vault is worth now: value at the marks split into reliable and indicative marks, cost and P&L, today’s change, and an estimate of what the units free to sell would fetch from the bids resting on the book now. Read only.

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
    resp, r, err := apiClient.VaultMarketsApi.GetVaultMarketWorthVault(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultMarketWorthVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketWorthVault`: GetVaultMarketWorthVaultResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultMarketWorthVault`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketWorthVaultRequest struct via the builder pattern


### Return type

**GetVaultMarketWorthVaultResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultReceipt

> GetVaultReceiptResponse GetVaultReceipt(ctx, receiptNo).Execute()

One of your receipts with its full hash chain (issued, transferred, pledged, released, surrendered), re-verified on read: `chain.ok` is false if any event was altered.

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
    receiptNo := "receiptNo_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultReceipt(context.Background(), receiptNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultReceipt`: GetVaultReceiptResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptNo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultReceiptResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultReceiptToken

> GetVaultReceiptTokenResponse GetVaultReceiptToken(ctx, id).Execute()

One receipt token (yours, or on a unit you own now) with its chain log: built, minted, frozen, thawed, new holder seen, redeem built, burned, with explorer links.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultReceiptToken(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultReceiptToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultReceiptToken`: GetVaultReceiptTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultReceiptToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultReceiptTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultReceiptTokenResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultReceiptTokenByMint

> GetVaultReceiptTokenByMintResponse GetVaultReceiptTokenByMint(ctx, mint).Execute()

What a Solana mint is: the warehouse receipt number and item it stands for, and whether the token is live, frozen (unit pledged or lent) or burned, and whether the unit is still on a vault shelf. Names no person and no wallet.

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
    mint := "mint_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.GetVaultReceiptTokenByMint(context.Background(), mint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.GetVaultReceiptTokenByMint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultReceiptTokenByMint`: GetVaultReceiptTokenByMintResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.GetVaultReceiptTokenByMint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mint** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultReceiptTokenByMintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultReceiptTokenByMintResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketAccounts

> ListVaultMarketAccountsResponse ListVaultMarketAccounts(ctx).Execute()

Every account you can act on: your own, and each sub-account your entity memberships reach, with your role.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketAccounts`: ListVaultMarketAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketAccountsRequest struct via the builder pattern


### Return type

**ListVaultMarketAccountsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketAdlEvents

> ListVaultMarketAdlEventsResponse ListVaultMarketAdlEvents(ctx).Execute()

The ADL events your positions were part of, newest first: role bankrupt (yours went under water) or counterparty (yours was first in line), the price both closed at and the transfer status.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketAdlEvents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketAdlEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketAdlEvents`: ListVaultMarketAdlEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketAdlEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketAdlEventsRequest struct via the builder pattern


### Return type

**ListVaultMarketAdlEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketAdlPositions

> ListVaultMarketAdlPositionsResponse ListVaultMarketAdlPositions(ctx).Execute()

Your live positions, each with its place in its item ADL queue: lights 5 (first in line) to 1 (last), 0 when not in the queue (only profitable positions are), ranked by P&L% x effective leverage.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketAdlPositions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketAdlPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketAdlPositions`: ListVaultMarketAdlPositionsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketAdlPositions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketAdlPositionsRequest struct via the builder pattern


### Return type

**ListVaultMarketAdlPositionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketAgentDrafts

> ListVaultMarketAgentDraftsResponse ListVaultMarketAgentDrafts(ctx).Execute()

Limit orders the shopping agent drafted for you, newest first. A draft places nothing: send its orderBody to POST /api/v1/vault-markets/pro/orders to place it.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketAgentDrafts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketAgentDrafts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketAgentDrafts`: ListVaultMarketAgentDraftsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketAgentDrafts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketAgentDraftsRequest struct via the builder pattern


### Return type

**ListVaultMarketAgentDraftsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketAgentWatches

> ListVaultMarketAgentWatchesResponse ListVaultMarketAgentWatches(ctx).Execute()

Your price watches on book keys (best bid, lowest ask, last trade or market price, in cents), live ones first. `?live=true` for live only.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketAgentWatches(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketAgentWatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketAgentWatches`: ListVaultMarketAgentWatchesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketAgentWatches`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketAgentWatchesRequest struct via the builder pattern


### Return type

**ListVaultMarketAgentWatchesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketAlertLog

> ListVaultMarketAlertLogResponse ListVaultMarketAlertLog(ctx).Execute()

What the alerts hub did with your recent vault alerts (`?limit=` 1 to 200, default 50): kind, title, the channels it went to, what was held (channel:reason) and the email digest state.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketAlertLog(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketAlertLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketAlertLog`: ListVaultMarketAlertLogResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketAlertLog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketAlertLogRequest struct via the builder pattern


### Return type

**ListVaultMarketAlertLogResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketAlerts

> ListVaultMarketAlertsResponse ListVaultMarketAlerts(ctx).Execute()

Your price alerts on book keys.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketAlerts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketAlerts`: ListVaultMarketAlertsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketAlerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketAlertsRequest struct via the builder pattern


### Return type

**ListVaultMarketAlertsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketAuctions

> ListVaultMarketAuctionsResponse ListVaultMarketAuctions(ctx).Execute()

Open vault auction events (weekly, ending Sunday; premier, units marked $10,000+) and their lots, soonest ending first.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketAuctions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketAuctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketAuctions`: ListVaultMarketAuctionsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketAuctions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketAuctionsRequest struct via the builder pattern


### Return type

**ListVaultMarketAuctionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketBlocks

> ListVaultMarketBlocksResponse ListVaultMarketBlocks(ctx).Execute()

An entity's block orders, newest first, each with its plan and how it was allocated.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketBlocks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketBlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketBlocks`: ListVaultMarketBlocksResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketBlocks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketBlocksRequest struct via the builder pattern


### Return type

**ListVaultMarketBlocksResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketBullionFamilyPegs

> ListVaultMarketBullionFamilyPegsResponse ListVaultMarketBullionFamilyPegs(ctx).Execute()

Your family-wide spot-pegged quotes (live, or all with status=all): one quote resting across a bullion family books at one condition, with units left, the peg, its status and every leg (the pegged quote on each book).

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketBullionFamilyPegs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketBullionFamilyPegs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketBullionFamilyPegs`: ListVaultMarketBullionFamilyPegsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketBullionFamilyPegs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketBullionFamilyPegsRequest struct via the builder pattern


### Return type

**ListVaultMarketBullionFamilyPegsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketBullionKeys

> ListVaultMarketBullionKeysResponse ListVaultMarketBullionKeys(ctx).Execute()

Bullion key families (metal, form, size, mint or any mint): members, vault units, best resting bid and ask across their books, and spot as a labelled reference with what one unit of metal is worth at it. Spot is never a Crossly price.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketBullionKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketBullionKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketBullionKeys`: ListVaultMarketBullionKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketBullionKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketBullionKeysRequest struct via the builder pattern


### Return type

**ListVaultMarketBullionKeysResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketBullionPegs

> ListVaultMarketBullionPegsResponse ListVaultMarketBullionPegs(ctx).Execute()

Your spot-pegged quotes (live, or all with status=all): side, units, the peg, the order it rests as, the last price and spot it was priced from, and why it is pulled when it is.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketBullionPegs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketBullionPegs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketBullionPegs`: ListVaultMarketBullionPegsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketBullionPegs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketBullionPegsRequest struct via the builder pattern


### Return type

**ListVaultMarketBullionPegsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketCollectionWorth

> ListVaultMarketCollectionWorthResponse ListVaultMarketCollectionWorth(ctx).Execute()

Your collection valued at the display marks: rows on a vault book key with a mark are summed; rows with no mark or no key are counted, never valued.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketCollectionWorth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketCollectionWorth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketCollectionWorth`: ListVaultMarketCollectionWorthResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketCollectionWorth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketCollectionWorthRequest struct via the builder pattern


### Return type

**ListVaultMarketCollectionWorthResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketConfirmations

> ListVaultMarketConfirmationsResponse ListVaultMarketConfirmations(ctx).Execute()

Trade confirmations for your fills in a window (default 31 days): price, sale fee, maker or taker, order, vault unit and warehouse receipt.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketConfirmations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketConfirmations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketConfirmations`: ListVaultMarketConfirmationsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketConfirmations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketConfirmationsRequest struct via the builder pattern


### Return type

**ListVaultMarketConfirmationsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketCreatorCohostInvites

> ListVaultMarketCreatorCohostInvitesResponse ListVaultMarketCreatorCohostInvites(ctx).Execute()

For creators: your invites from Crossly to co-host a Market Open launch show (invited, accepted, declined or withdrawn), each with the key, its show and the guidance for co-hosts.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketCreatorCohostInvites(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketCreatorCohostInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketCreatorCohostInvites`: ListVaultMarketCreatorCohostInvitesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketCreatorCohostInvites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketCreatorCohostInvitesRequest struct via the builder pattern


### Return type

**ListVaultMarketCreatorCohostInvitesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketCredit

> ListVaultMarketCreditResponse ListVaultMarketCredit(ctx).Execute()

Your vault loans and advances (live by default): what is owed, interest accrued, loan-to-value, the liquidation line and the collateral value it is called at.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketCredit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketCredit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketCredit`: ListVaultMarketCreditResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketCredit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketCreditRequest struct via the builder pattern


### Return type

**ListVaultMarketCreditResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketDealerBands

> ListVaultMarketDealerBandsResponse ListVaultMarketDealerBands(ctx).Execute()

Your Instant Exit bands and whether each is being kept now.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketDealerBands(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketDealerBands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketDealerBands`: ListVaultMarketDealerBandsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketDealerBands`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketDealerBandsRequest struct via the builder pattern


### Return type

**ListVaultMarketDealerBandsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketDealerBids

> ListVaultMarketDealerBidsResponse ListVaultMarketDealerBids(ctx).Execute()

Your standing dealer bids: open (default) or all.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketDealerBids(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketDealerBids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketDealerBids`: ListVaultMarketDealerBidsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketDealerBids`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketDealerBidsRequest struct via the builder pattern


### Return type

**ListVaultMarketDealerBidsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketDealerPoolFills

> ListVaultMarketDealerPoolFillsResponse ListVaultMarketDealerPoolFills(ctx, id).Execute()

A curve pool's fills, newest first: side, price and spot after.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketDealerPoolFills(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketDealerPoolFills``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketDealerPoolFills`: ListVaultMarketDealerPoolFillsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketDealerPoolFills`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketDealerPoolFillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketDealerPoolFillsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketDealerPools

> ListVaultMarketDealerPoolsResponse ListVaultMarketDealerPools(ctx).Execute()

Your curve pools: open (default) or all, each with its curve, spot, units, buy budget, running totals, the ladder it shows and its live legs.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketDealerPools(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketDealerPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketDealerPools`: ListVaultMarketDealerPoolsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketDealerPools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketDealerPoolsRequest struct via the builder pattern


### Return type

**ListVaultMarketDealerPoolsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketDigestHistory

> ListVaultMarketDigestHistoryResponse ListVaultMarketDigestHistory(ctx).Execute()

Past weekly digests: week, status (sent / skipped / failed), the channels used and a summary.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketDigestHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketDigestHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketDigestHistory`: ListVaultMarketDigestHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketDigestHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketDigestHistoryRequest struct via the builder pattern


### Return type

**ListVaultMarketDigestHistoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketEntityAudit

> ListVaultMarketEntityAuditResponse ListVaultMarketEntityAudit(ctx, id).Execute()

The entity's audit trail, newest first: hash-chained rows with actor, account, API key, IP, session and source.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketEntityAudit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketEntityAudit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketEntityAudit`: ListVaultMarketEntityAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketEntityAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketEntityAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketEntityAuditResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketEntityConversions

> ListVaultMarketEntityConversionsResponse ListVaultMarketEntityConversions(ctx, id).Execute()

The custodian's CBX-to-USD reward conversions for an entity that elected USD, newest first, with their status.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketEntityConversions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketEntityConversions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketEntityConversions`: ListVaultMarketEntityConversionsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketEntityConversions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketEntityConversionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketEntityConversionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketEstimatesAccuracy

> ListVaultMarketEstimatesAccuracyResponse ListVaultMarketEstimatesAccuracy(ctx).Execute()

The Crossly Estimate's track record: walk-forward backtest n, FSD and bias per method, category and liquidity bucket, and whether estimates in that cell are shown.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketEstimatesAccuracy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketEstimatesAccuracy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketEstimatesAccuracy`: ListVaultMarketEstimatesAccuracyResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketEstimatesAccuracy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketEstimatesAccuracyRequest struct via the builder pattern


### Return type

**ListVaultMarketEstimatesAccuracyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketFeed

> ListVaultMarketFeedResponse ListVaultMarketFeed(ctx).Execute()

New vault buys by the people you follow who share them, newest first (no prices). Page with `before` (the last item's `at`).

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketFeed(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketFeed`: ListVaultMarketFeedResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketFeed`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketFeedRequest struct via the builder pattern


### Return type

**ListVaultMarketFeedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketFollowCopyWatch

> ListVaultMarketFollowCopyWatchResponse ListVaultMarketFollowCopyWatch(ctx).Execute()

The collectors you copy-watch and whether each still shares bids and asks.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketFollowCopyWatch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketFollowCopyWatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketFollowCopyWatch`: ListVaultMarketFollowCopyWatchResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketFollowCopyWatch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketFollowCopyWatchRequest struct via the builder pattern


### Return type

**ListVaultMarketFollowCopyWatchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketFollowCopyWatchAlerts

> ListVaultMarketFollowCopyWatchAlertsResponse ListVaultMarketFollowCopyWatchAlerts(ctx).Execute()

Your copy-watch alerts, newest first: who, bid or ask, quantity, limit price, market, and whether it still rests. Page with `before`.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketFollowCopyWatchAlerts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketFollowCopyWatchAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketFollowCopyWatchAlerts`: ListVaultMarketFollowCopyWatchAlertsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketFollowCopyWatchAlerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketFollowCopyWatchAlertsRequest struct via the builder pattern


### Return type

**ListVaultMarketFollowCopyWatchAlertsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketFollowFeed

> ListVaultMarketFollowFeedResponse ListVaultMarketFollowFeed(ctx).Execute()

Public vault activity of the collectors you follow (buys, units sent in, sales, Hall loans), newest first, never a price. Filter with `kinds`; page with `before`.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketFollowFeed(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketFollowFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketFollowFeed`: ListVaultMarketFollowFeedResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketFollowFeed`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketFollowFeedRequest struct via the builder pattern


### Return type

**ListVaultMarketFollowFeedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketFollowFollowing

> ListVaultMarketFollowFollowingResponse ListVaultMarketFollowFollowing(ctx).Execute()

The collectors you follow who have a public vault: what each shares and whether you copy-watch them.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketFollowFollowing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketFollowFollowing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketFollowFollowing`: ListVaultMarketFollowFollowingResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketFollowFollowing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketFollowFollowingRequest struct via the builder pattern


### Return type

**ListVaultMarketFollowFollowingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketFollowLeaderboards

> ListVaultMarketFollowLeaderboardsResponse ListVaultMarketFollowLeaderboards(ctx).Execute()

An opt-in return leaderboard (`period` 30d, 90d or 365d; `category` all or a category slug), ranked by risk-adjusted return from the book A/B marks. Percent and risk only: no dollars, no volume.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketFollowLeaderboards(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketFollowLeaderboards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketFollowLeaderboards`: ListVaultMarketFollowLeaderboardsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketFollowLeaderboards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketFollowLeaderboardsRequest struct via the builder pattern


### Return type

**ListVaultMarketFollowLeaderboardsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketFollowProfileActivity

> ListVaultMarketFollowFeedResponse ListVaultMarketFollowProfileActivity(ctx, handle).Execute()

One collector public vault activity, newest first, never a price. 404 when their vault is not public.

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketFollowProfileActivity(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketFollowProfileActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketFollowProfileActivity`: ListVaultMarketFollowFeedResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketFollowProfileActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketFollowProfileActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketFollowFeedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketGifts

> ListVaultMarketGiftsResponse ListVaultMarketGifts(ctx).Execute()

Gifts you sent and received, newest first, with their status.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketGifts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketGifts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketGifts`: ListVaultMarketGiftsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketGifts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketGiftsRequest struct via the builder pattern


### Return type

**ListVaultMarketGiftsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketGradeToVault

> ListVaultMarketGradeToVaultResponse ListVaultMarketGradeToVault(ctx).Execute()

Your Grade to Vault flows, newest first, each brought up to date: its step, the grading submission, the unit, the grade, the market for it, and the ask once listed. `live=1` for those still going.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketGradeToVault(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketGradeToVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketGradeToVault`: ListVaultMarketGradeToVaultResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketGradeToVault`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketGradeToVaultRequest struct via the builder pattern


### Return type

**ListVaultMarketGradeToVaultResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketGrading

> ListVaultMarketGradingResponse ListVaultMarketGrading(ctx).Execute()

Your vault units sent to grading: where each card is in its run, and the grade once it is back.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketGrading(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketGrading``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketGrading`: ListVaultMarketGradingResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketGrading`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketGradingRequest struct via the builder pattern


### Return type

**ListVaultMarketGradingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketHallLenders

> ListVaultMarketHallLendersResponse ListVaultMarketHallLenders(ctx).Execute()

The vault units lent to this week's Hall of Grails exhibition and on show now, in plinth order: the item (skuId, name, imageUrl, grade), the lender as the plate names them (a display name, or \"a collector\"), the plate lines and how many exhibitions the unit has stood in. meta.exhibition is the week (key, title, endsAt). Names no ids and no amounts.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketHallLenders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketHallLenders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketHallLenders`: ListVaultMarketHallLendersResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketHallLenders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketHallLendersRequest struct via the builder pattern


### Return type

**ListVaultMarketHallLendersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexCommitteeMinutes

> ListVaultMarketIndexCommitteeMinutesResponse ListVaultMarketIndexCommitteeMinutes(ctx).Execute()

Published Index Committee minutes, newest first: attendees (names and roles), quorum, agenda, minutes and decisions.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexCommitteeMinutes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexCommitteeMinutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexCommitteeMinutes`: ListVaultMarketIndexCommitteeMinutesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexCommitteeMinutes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexCommitteeMinutesRequest struct via the builder pattern


### Return type

**ListVaultMarketIndexCommitteeMinutesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexComplaints

> ListVaultMarketIndexComplaintsResponse ListVaultMarketIndexComplaints(ctx).Execute()

Your complaints and their answers, newest first.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexComplaints(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexComplaints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexComplaints`: ListVaultMarketIndexComplaintsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexComplaints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexComplaintsRequest struct via the builder pattern


### Return type

**ListVaultMarketIndexComplaintsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexConstituents

> ListVaultMarketIndexMembersResponse ListVaultMarketIndexConstituents(ctx, slug).Execute()

The book keys in an index with their weight, move since the last close and contribution to the index's move (basis points).

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexConstituents(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexConstituents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexConstituents`: ListVaultMarketIndexMembersResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexConstituents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexConstituentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketIndexMembersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexConsultations

> ListVaultMarketIndexConsultationsResponse ListVaultMarketIndexConsultations(ctx).Execute()

Consultations on index method changes (open for at least 30 days), newest first. Filter by `status`.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexConsultations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexConsultations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexConsultations`: ListVaultMarketIndexConsultationsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexConsultations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexConsultationsRequest struct via the builder pattern


### Return type

**ListVaultMarketIndexConsultationsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexEvents

> ListVaultMarketIndexEventsResponse ListVaultMarketIndexEvents(ctx).Execute()

Every index event, newest first: add, delete, rebalance, method, correction, halt; each with when it was announced, when it takes effect, the detail and the reason. Filter by `kind`.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexEvents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexEvents`: ListVaultMarketIndexEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexEventsRequest struct via the builder pattern


### Return type

**ListVaultMarketIndexEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexEventsBySlug

> ListVaultMarketIndexEventsResponse ListVaultMarketIndexEventsBySlug(ctx, slug).Execute()

One index's events, newest first.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexEventsBySlug(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexEventsBySlug``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexEventsBySlug`: ListVaultMarketIndexEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexEventsBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexEventsBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketIndexEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexJudgments

> ListVaultMarketIndexJudgmentsResponse ListVaultMarketIndexJudgments(ctx).Execute()

The expert-judgment log: every Index Committee action (halt, resume, exclude a key, void a trade, restate, retire) with its written reason.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexJudgments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexJudgments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexJudgments`: ListVaultMarketIndexJudgmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexJudgments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexJudgmentsRequest struct via the builder pattern


### Return type

**ListVaultMarketIndexJudgmentsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexLevelSeries

> ListVaultMarketIndexLevelSeriesResponse ListVaultMarketIndexLevelSeries(ctx, slug).Execute()

An index's published levels. `engine=chain` (default): daily closes. `engine=tpd|repeat_sales`: weekly (else monthly) model estimates with a 95% band and n; labelled, never a price. `series=settlement` returns each value as first published (never restated). Below the data gate `data` is empty and `notice` explains.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexLevelSeries(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexLevelSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexLevelSeries`: ListVaultMarketIndexLevelSeriesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexLevelSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexLevelSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketIndexLevelSeriesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexLicenceMine

> ListVaultMarketIndexLicenceMineResponse ListVaultMarketIndexLicenceMine(ctx).Execute()

Your index feed licences (the indexes each covers, or all), as ops recorded them. No licence: the feed answers 403 not_a_licensee.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexLicenceMine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexLicenceMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexLicenceMine`: ListVaultMarketIndexLicenceMineResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexLicenceMine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexLicenceMineRequest struct via the builder pattern


### Return type

**ListVaultMarketIndexLicenceMineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexMembers

> ListVaultMarketIndexMembersResponse ListVaultMarketIndexMembers(ctx, slug).Execute()

The members of an index on a day (`asOf`, default today) with their weights, shares, moves and contributions; for a past day, the inputs stored with that day's close and each member's effective dates.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexMembers(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexMembers`: ListVaultMarketIndexMembersResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketIndexMembersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexMovers

> ListVaultMarketIndexesResponse ListVaultMarketIndexMovers(ctx).Execute()

The built-in indexes that moved most since their last daily close.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexMovers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexMovers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexMovers`: ListVaultMarketIndexesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexMovers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexMoversRequest struct via the builder pattern


### Return type

**ListVaultMarketIndexesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketIndexes

> ListVaultMarketIndexesResponse ListVaultMarketIndexes(ctx).Execute()

Indexes over book keys: per category, per top brand, the Crossly 50 and your own. Each with its level now (base 1000 at inception) and change since the last daily close.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketIndexes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketIndexes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketIndexes`: ListVaultMarketIndexesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketIndexes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketIndexesRequest struct via the builder pattern


### Return type

**ListVaultMarketIndexesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketInstantExitHistory

> ListVaultMarketInstantExitHistoryResponse ListVaultMarketInstantExitHistory(ctx).Execute()

Your Instant Exits, newest first: quoted price, mark, and whether it filled.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketInstantExitHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketInstantExitHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketInstantExitHistory`: ListVaultMarketInstantExitHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketInstantExitHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketInstantExitHistoryRequest struct via the builder pattern


### Return type

**ListVaultMarketInstantExitHistoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketInstantExitHoldings

> ListVaultMarketInstantExitHoldingsResponse ListVaultMarketInstantExitHoldings(ctx).Execute()

Instant Exit quotes for everything you hold in the vault, one per book key, with units held and which are free to sell.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketInstantExitHoldings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketInstantExitHoldings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketInstantExitHoldings`: ListVaultMarketInstantExitHoldingsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketInstantExitHoldings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketInstantExitHoldingsRequest struct via the builder pattern


### Return type

**ListVaultMarketInstantExitHoldingsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketJointTrades

> ListVaultMarketJointTradesResponse ListVaultMarketJointTrades(ctx, id).Execute()

The joint vault trades, newest first: side, book key, limit, who started it, whose account the order sits on, and the order standing on the book.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketJointTrades(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketJointTrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketJointTrades`: ListVaultMarketJointTradesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketJointTrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketJointTradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketJointTradesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketLargeTraderReports

> ListVaultMarketLargeTraderReportsResponse ListVaultMarketLargeTraderReports(ctx).Execute()

Your large-trader reports (with any linked accounts, as one party): due first, then filed, then cleared in the last 90 days. Each has `action`: file | amend | refile | none.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketLargeTraderReports(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketLargeTraderReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketLargeTraderReports`: ListVaultMarketLargeTraderReportsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketLargeTraderReports`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketLargeTraderReportsRequest struct via the builder pattern


### Return type

**ListVaultMarketLargeTraderReportsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketLendingStatements

> ListVaultMarketLendingStatementsResponse ListVaultMarketLendingStatements(ctx).Execute()

The days you have stored collateral statements for, newest first, with their summary.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketLendingStatements(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketLendingStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketLendingStatements`: ListVaultMarketLendingStatementsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketLendingStatements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketLendingStatementsRequest struct via the builder pattern


### Return type

**ListVaultMarketLendingStatementsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketLoans

> ListVaultMarketLoansResponse ListVaultMarketLoans(ctx).Execute()

Units you lent: rate, fees earned, recall due.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketLoans(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketLoans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketLoans`: ListVaultMarketLoansResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketLoans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketLoansRequest struct via the builder pattern


### Return type

**ListVaultMarketLoansResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketLocateMine

> ListVaultMarketLocateMineResponse ListVaultMarketLocateMine(ctx).Execute()

Your locates, newest first (held / used / released / expired).

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketLocateMine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketLocateMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketLocateMine`: ListVaultMarketLocateMineResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketLocateMine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketLocateMineRequest struct via the builder pattern


### Return type

**ListVaultMarketLocateMineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketMidpoint

> ListVaultMarketMidpointResponse ListVaultMarketMidpoint(ctx).Execute()

Your midpoint block orders.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketMidpoint(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketMidpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketMidpoint`: ListVaultMarketMidpointResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketMidpoint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketMidpointRequest struct via the builder pattern


### Return type

**ListVaultMarketMidpointResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketMovers

> ListVaultMarketMoversResponse ListVaultMarketMovers(ctx).Execute()

Vault items with the biggest 24 h moves in their market price (from the order book), with units in the vault.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketMovers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketMovers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketMovers`: ListVaultMarketMoversResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketMovers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketMoversRequest struct via the builder pattern


### Return type

**ListVaultMarketMoversResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketNavHistory

> ListVaultMarketNavHistoryResponse ListVaultMarketNavHistory(ctx).Execute()

Daily NAV snapshots (mark, bid, liquidation, Level 2 / Level 3).

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketNavHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketNavHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketNavHistory`: ListVaultMarketNavHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketNavHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketNavHistoryRequest struct via the builder pattern


### Return type

**ListVaultMarketNavHistoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketOffers

> ListVaultMarketOffersResponse ListVaultMarketOffers(ctx).Execute()

The offers inbox: for each book key you hold, the best bid from someone else, units at it, the top bid levels, the mark, and your alert.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketOffers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketOffers`: ListVaultMarketOffersResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketOffers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketOffersRequest struct via the builder pattern


### Return type

**ListVaultMarketOffersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketOpening

> ListVaultMarketOpeningResponse ListVaultMarketOpening(ctx).Execute()

Keys opening soon on the vault book (and just opened): units on the shelf, dealers quoting, the waitlist, the mark when the book has one, and your place.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketOpening(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketOpening``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketOpening`: ListVaultMarketOpeningResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketOpening`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketOpeningRequest struct via the builder pattern


### Return type

**ListVaultMarketOpeningResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketOracleHistory

> ListVaultMarketOracleHistoryResponse ListVaultMarketOracleHistory(ctx, slug).Execute()

Recent publish attempts for an index over 30 days, newest first: published, skipped (with the reason: too few items priced, stale marks, confidence too wide...) or failed (with the error). Filter with `status`, cap with `limit` (default 100).

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketOracleHistory(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketOracleHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketOracleHistory`: ListVaultMarketOracleHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketOracleHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketOracleHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketOracleHistoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketP2pAuctions

> ListVaultMarketP2pLoansResponse ListVaultMarketP2pAuctions(ctx).Execute()

Running refinancing auctions, soonest to end first: lenders calls (`auction`) and borrowers looking for a cheaper lender (`refiAuction`). The payoff now, the rate a new lender gets now, when it ends, and the item.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketP2pAuctions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketP2pAuctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketP2pAuctions`: ListVaultMarketP2pLoansResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketP2pAuctions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketP2pAuctionsRequest struct via the builder pattern


### Return type

**ListVaultMarketP2pLoansResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketP2pLoans

> ListVaultMarketP2pLoansResponse ListVaultMarketP2pLoans(ctx).Execute()

Your P2P loans as borrower and/or lender (`role`), live by default (`status`): what is owed now, daily interest, the refinancing auction when called, and the item mark for information only (it never triggers anything).

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketP2pLoans(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketP2pLoans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketP2pLoans`: ListVaultMarketP2pLoansResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketP2pLoans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketP2pLoansRequest struct via the builder pattern


### Return type

**ListVaultMarketP2pLoansResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketP2pOfferMine

> ListVaultMarketP2pOffersResponse ListVaultMarketP2pOfferMine(ctx).Execute()

Your P2P loan offers, newest first, with how many loans each has funded.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketP2pOfferMine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketP2pOfferMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketP2pOfferMine`: ListVaultMarketP2pOffersResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketP2pOfferMine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketP2pOfferMineRequest struct via the builder pattern


### Return type

**ListVaultMarketP2pOffersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketP2pOffers

> ListVaultMarketP2pOffersResponse ListVaultMarketP2pOffers(ctx).Execute()

The open P2P loan offers book, lowest rate first: cents offered per loan, APR (bps), loans left, and the item (book key) or unit each lends against. `bookKey` narrows to one item. Lender identities are never shown.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketP2pOffers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketP2pOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketP2pOffers`: ListVaultMarketP2pOffersResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketP2pOffers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketP2pOffersRequest struct via the builder pattern


### Return type

**ListVaultMarketP2pOffersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketP2pUnits

> ListVaultMarketP2pUnitsResponse ListVaultMarketP2pUnits(ctx).Execute()

Your vault units for P2P borrowing: whether each can be pledged now (and why not), the best offers covering it, and its loan when it has one.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketP2pUnits(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketP2pUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketP2pUnits`: ListVaultMarketP2pUnitsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketP2pUnits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketP2pUnitsRequest struct via the builder pattern


### Return type

**ListVaultMarketP2pUnitsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketPaperFills

> ListVaultMarketPaperFillsResponse ListVaultMarketPaperFills(ctx).Execute()

Your paper fills since your last reset, newest first: price, simulated fee, taker or maker, and whether a visible real order (book) or a real trade (tape) filled it. Play money only.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketPaperFills(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketPaperFills``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketPaperFills`: ListVaultMarketPaperFillsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketPaperFills`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketPaperFillsRequest struct via the builder pattern


### Return type

**ListVaultMarketPaperFillsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketPaperLeaderboard

> ListVaultMarketPaperLeaderboardResponse ListVaultMarketPaperLeaderboard(ctx).Execute()

The weekly paper trading leaderboard (opt-in, display names only): best return on play equity this week (or a past `week`, its Monday as YYYY-MM-DD), with your own line. Nothing is won or paid.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketPaperLeaderboard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketPaperLeaderboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketPaperLeaderboard`: ListVaultMarketPaperLeaderboardResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketPaperLeaderboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketPaperLeaderboardRequest struct via the builder pattern


### Return type

**ListVaultMarketPaperLeaderboardResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketPaperMissions

> ListVaultMarketPaperMissionsResponse ListVaultMarketPaperMissions(ctx).Execute()

Your paper-trading missions (play money only): each mission with what to do, your count toward its target, whether it is done and when, plus a suggested market for the \"Read the mark\" mission. Progress is observed by the server from your paper orders and fills, price watches, set goals and graded mark checks. Missions earn world quest progress and cosmetic pins, never money.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketPaperMissions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketPaperMissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketPaperMissions`: ListVaultMarketPaperMissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketPaperMissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketPaperMissionsRequest struct via the builder pattern


### Return type

**ListVaultMarketPaperMissionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketPaperOrders

> ListVaultMarketPaperOrdersResponse ListVaultMarketPaperOrders(ctx).Execute()

Your paper orders since your last reset, newest first: working, done (filled, cancelled, expired) or all. Play money only.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketPaperOrders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketPaperOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketPaperOrders`: ListVaultMarketPaperOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketPaperOrders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketPaperOrdersRequest struct via the builder pattern


### Return type

**ListVaultMarketPaperOrdersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketPartnerAccruals

> ListVaultMarketPartnerAccrualsResponse ListVaultMarketPartnerAccruals(ctx).Execute()

Your partner fee shares, newest first: accrued on a unit's first sale, paid in Bucks once the sale is final, void if it was refunded.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketPartnerAccruals(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketPartnerAccruals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketPartnerAccruals`: ListVaultMarketPartnerAccrualsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketPartnerAccruals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketPartnerAccrualsRequest struct via the builder pattern


### Return type

**ListVaultMarketPartnerAccrualsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketPoolStakes

> ListVaultMarketPoolStakesResponse ListVaultMarketPoolStakes(ctx).Execute()

Your stakes in the market pool and what each redeems for now.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketPoolStakes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketPoolStakes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketPoolStakes`: ListVaultMarketPoolStakesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketPoolStakes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketPoolStakesRequest struct via the builder pattern


### Return type

**ListVaultMarketPoolStakesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketPortfolioMarginHistory

> ListVaultMarketPortfolioMarginHistoryResponse ListVaultMarketPortfolioMarginHistory(ctx).Execute()

The 5-minute portfolio margin recalculations for a firm sub-account over the last `hours` (default 24, at most 720), newest first: equity, requirement, excess, scenario and floor figures, gross and state.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketPortfolioMarginHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketPortfolioMarginHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketPortfolioMarginHistory`: ListVaultMarketPortfolioMarginHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketPortfolioMarginHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketPortfolioMarginHistoryRequest struct via the builder pattern


### Return type

**ListVaultMarketPortfolioMarginHistoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketPositionUnits

> ListVaultMarketPositionUnitsResponse ListVaultMarketPositionUnits(ctx, id).Execute()

The vault units of one of your positions that you hold (a leveraged long's collateral), each with its story link.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketPositionUnits(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketPositionUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketPositionUnits`: ListVaultMarketPositionUnitsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketPositionUnits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketPositionUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketPositionUnitsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketPositions

> ListVaultMarketPositionsResponse ListVaultMarketPositions(ctx).Execute()

Your positions, live (default) or closed.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketPositions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketPositions`: ListVaultMarketPositionsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketPositions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketPositionsRequest struct via the builder pattern


### Return type

**ListVaultMarketPositionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketPositionsRisk

> ListVaultMarketPositionsRiskResponse ListVaultMarketPositionsRisk(ctx).Execute()

For each live position: its tier, health, margin state, the mark at which it is called and the mark at which it is liquidated, and the distance to it.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketPositionsRisk(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketPositionsRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketPositionsRisk`: ListVaultMarketPositionsRiskResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketPositionsRisk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketPositionsRiskRequest struct via the builder pattern


### Return type

**ListVaultMarketPositionsRiskResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketPriceWidgets

> ListVaultMarketPriceWidgetsResponse ListVaultMarketPriceWidgets(ctx).Execute()

Your price widgets, newest first, each with its iframe and script-tag snippets. `meta.codes` lists your affiliate codes a widget can link through; `meta.max` is the cap on live widgets.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketPriceWidgets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketPriceWidgets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketPriceWidgets`: ListVaultMarketPriceWidgetsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketPriceWidgets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketPriceWidgetsRequest struct via the builder pattern


### Return type

**ListVaultMarketPriceWidgetsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketProCandles

> ListVaultMarketProCandlesResponse ListVaultMarketProCandles(ctx).Execute()

OHLCV candles (1m, 5m, 15m, 1h, 4h, 1d) from the trades on a book key.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketProCandles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketProCandles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketProCandles`: ListVaultMarketProCandlesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketProCandles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketProCandlesRequest struct via the builder pattern


### Return type

**ListVaultMarketProCandlesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketProEvents

> ListVaultMarketProEventsResponse ListVaultMarketProEvents(ctx).Execute()

Your order event log after a sequence number (new, leg, fill, replaced, cancelled, expired, done): a replayable drop copy.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketProEvents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketProEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketProEvents`: ListVaultMarketProEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketProEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketProEventsRequest struct via the builder pattern


### Return type

**ListVaultMarketProEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketProLinks

> ListVaultMarketProLinksResponse ListVaultMarketProLinks(ctx).Execute()

Your linked orders (brackets, OCO pairs, trailing stops, position exits): live (default) or done.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketProLinks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketProLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketProLinks`: ListVaultMarketProLinksResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketProLinks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketProLinksRequest struct via the builder pattern


### Return type

**ListVaultMarketProLinksResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketProOrders

> ListVaultMarketProOrdersResponse ListVaultMarketProOrders(ctx).Execute()

Your orders: open (default) or done, optionally on one book key, or only those booked to one sub-account (`subaccountId`).

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketProOrders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketProOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketProOrders`: ListVaultMarketProOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketProOrders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketProOrdersRequest struct via the builder pattern


### Return type

**ListVaultMarketProOrdersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketProQuotes

> ListVaultMarketProQuotesResponse ListVaultMarketProQuotes(ctx).Execute()

Your live quotes.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketProQuotes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketProQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketProQuotes`: ListVaultMarketProQuotesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketProQuotes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketProQuotesRequest struct via the builder pattern


### Return type

**ListVaultMarketProQuotesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketProRewards

> ListVaultMarketProRewardsResponse ListVaultMarketProRewards(ctx).Execute()

Your daily market-maker liquidity rewards (paid in CBX).

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketProRewards(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketProRewards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketProRewards`: ListVaultMarketProRewardsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketProRewards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketProRewardsRequest struct via the builder pattern


### Return type

**ListVaultMarketProRewardsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketProTape

> ListVaultMarketProTapeResponse ListVaultMarketProTape(ctx).Execute()

The latest trades on a book key, newest first, with which side made and which took.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketProTape(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketProTape``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketProTape`: ListVaultMarketProTapeResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketProTape`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketProTapeRequest struct via the builder pattern


### Return type

**ListVaultMarketProTapeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketProofOfReserveLeaves

> ListVaultMarketProofOfReserveLeavesResponse ListVaultMarketProofOfReserveLeaves(ctx, day).Execute()

One day's leaf hashes in tree order, a page at a time: enough to recount the units and rebuild the root yourself. Leaves are salted, so they name no owner.

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
    day := "day_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketProofOfReserveLeaves(context.Background(), day).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketProofOfReserveLeaves``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketProofOfReserveLeaves`: ListVaultMarketProofOfReserveLeavesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketProofOfReserveLeaves`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**day** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketProofOfReserveLeavesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketProofOfReserveLeavesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketRecurring

> ListVaultMarketRecurringResponse ListVaultMarketRecurring(ctx).Execute()

Your recurring buys: key, schedule, price cap, payment, auto-lend, next run.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketRecurring(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketRecurring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketRecurring`: ListVaultMarketRecurringResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketRecurring`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketRecurringRequest struct via the builder pattern


### Return type

**ListVaultMarketRecurringResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketRecurringRuns

> ListVaultMarketRecurringRunsResponse ListVaultMarketRecurringRuns(ctx, id).Execute()

One recurring buy's runs: filled, skipped (with the reason, e.g. the price was over the cap) or failed.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketRecurringRuns(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketRecurringRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketRecurringRuns`: ListVaultMarketRecurringRunsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketRecurringRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketRecurringRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketRecurringRunsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketRecurringRuns2

> ListVaultMarketRecurringRunsResponse ListVaultMarketRecurringRuns2(ctx).Execute()

Every run of all your recurring buys, newest first.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketRecurringRuns2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketRecurringRuns2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketRecurringRuns2`: ListVaultMarketRecurringRunsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketRecurringRuns2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketRecurringRuns2Request struct via the builder pattern


### Return type

**ListVaultMarketRecurringRunsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketRfq

> ListVaultMarketRfqResponse ListVaultMarketRfq(ctx).Execute()

Your quote requests, newest first.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketRfq(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketRfq``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketRfq`: ListVaultMarketRfqResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketRfq`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketRfqRequest struct via the builder pattern


### Return type

**ListVaultMarketRfqResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketRfqInbox

> ListVaultMarketRfqInboxResponse ListVaultMarketRfqInbox(ctx).Execute()

Open quote requests you were sent (as a dealer, DMM, LP or holder), with your live answer.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketRfqInbox(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketRfqInbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketRfqInbox`: ListVaultMarketRfqInboxResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketRfqInbox`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketRfqInboxRequest struct via the builder pattern


### Return type

**ListVaultMarketRfqInboxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketRfqOpen

> ListVaultMarketRfqResponse ListVaultMarketRfqOpen(ctx).Execute()

Open quote requests on one key (all-to-all: anyone with ID may answer). Anonymous.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketRfqOpen(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketRfqOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketRfqOpen`: ListVaultMarketRfqResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketRfqOpen`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketRfqOpenRequest struct via the builder pattern


### Return type

**ListVaultMarketRfqResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketSetGoalBaskets

> ListVaultMarketSetGoalBasketsResponse ListVaultMarketSetGoalBaskets(ctx, id).Execute()

A set goal basket bids, newest first, each with its budget, what it committed and spent, and every leg (piece, book key, limit, pro order id, status).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketSetGoalBaskets(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketSetGoalBaskets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketSetGoalBaskets`: ListVaultMarketSetGoalBasketsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketSetGoalBaskets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketSetGoalBasketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketSetGoalBasketsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketSetGoalSets

> ListVaultMarketSetGoalSetsResponse ListVaultMarketSetGoalSets(ctx).Execute()

Sets to start a goal on. With q (2+ characters): a word-start search of every catalogue set. Without: the sets you hold pieces of in the vault, then the collections you are building. Each with your live goal on it, if any.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketSetGoalSets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketSetGoalSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketSetGoalSets`: ListVaultMarketSetGoalSetsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketSetGoalSets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketSetGoalSetsRequest struct via the builder pattern


### Return type

**ListVaultMarketSetGoalSetsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketSetGoals

> ListVaultMarketSetGoalsResponse ListVaultMarketSetGoals(ctx).Execute()

Your set goals: the set, the target a piece must meet, pieces held of the total (from your vault), and the working basket bid.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketSetGoals(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketSetGoals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketSetGoals`: ListVaultMarketSetGoalsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketSetGoals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketSetGoalsRequest struct via the builder pattern


### Return type

**ListVaultMarketSetGoalsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketShareCards

> ListVaultMarketShareCardsResponse ListVaultMarketShareCards(ctx).Execute()

Your share cards (not taken down), each as the public sees it.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketShareCards(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketShareCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketShareCards`: ListVaultMarketShareCardsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketShareCards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketShareCardsRequest struct via the builder pattern


### Return type

**ListVaultMarketShareCardsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketSupplyDeposits

> ListVaultMarketSupplyDepositsResponse ListVaultMarketSupplyDeposits(ctx).Execute()

Your units sent to the vault, newest first, each with its state and its key's Instant Exit line (\"Instant Exit now: $X (N% of mark)\" or \"No market yet\"). Filter with `inventoryItemIds` (comma-separated).

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketSupplyDeposits(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketSupplyDeposits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketSupplyDeposits`: ListVaultMarketSupplyDepositsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketSupplyDeposits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketSupplyDepositsRequest struct via the builder pattern


### Return type

**ListVaultMarketSupplyDepositsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketSupplyInvites

> ListVaultMarketSupplyInvitesResponse ListVaultMarketSupplyInvites(ctx).Execute()

Market Open invites to you as a dealer: keys Crossly is opening and asks you to quote.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketSupplyInvites(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketSupplyInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketSupplyInvites`: ListVaultMarketSupplyInvitesResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketSupplyInvites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketSupplyInvitesRequest struct via the builder pattern


### Return type

**ListVaultMarketSupplyInvitesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketValuationStatements

> ListVaultMarketValuationStatementsResponse ListVaultMarketValuationStatements(ctx).Execute()

The insurance and appraisal statements you issued, newest first: purpose, date, units, total at the marks, hash and verify link.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketValuationStatements(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketValuationStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketValuationStatements`: ListVaultMarketValuationStatementsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketValuationStatements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketValuationStatementsRequest struct via the builder pattern


### Return type

**ListVaultMarketValuationStatementsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketVaultSuggestions

> ListVaultMarketVaultSuggestionsResponse ListVaultMarketVaultSuggestions(ctx).Execute()

What to send to the vault: your inventory items whose book key has a liquid vault market (tier A, B or C), best first. Each has a score (0 to 100) and its reasons, the market (tier, mark and its quality, best bid, best vault ask, spread), Instant Exit now, your list price, days listed without selling, and the vault fee. `meta.summary` counts what was left out and why. Inform only; send with POST /api/v1/vault-markets/supply/send.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketVaultSuggestions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketVaultSuggestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketVaultSuggestions`: ListVaultMarketVaultSuggestionsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketVaultSuggestions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketVaultSuggestionsRequest struct via the builder pattern


### Return type

**ListVaultMarketVaultSuggestionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketWatchlist

> ListVaultMarketWatchlistResponse ListVaultMarketWatchlist(ctx).Execute()

Your watched book keys with mark, 24 h and 7 d change, and the last 30 daily marks (sparkline).

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketWatchlist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketWatchlist`: ListVaultMarketWatchlistResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketWatchlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketWatchlistRequest struct via the builder pattern


### Return type

**ListVaultMarketWatchlistResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketWidgetTokens

> ListVaultMarketWidgetTokensResponse ListVaultMarketWidgetTokens(ctx).Execute()

Your live home-screen widget tokens, one per phone: id, name (the phone), public prefix, created and last used. The secret is never returned again.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketWidgetTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketWidgetTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketWidgetTokens`: ListVaultMarketWidgetTokensResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketWidgetTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketWidgetTokensRequest struct via the builder pattern


### Return type

**ListVaultMarketWidgetTokensResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketWorldCohosts

> ListVaultMarketWorldCohostsResponse ListVaultMarketWorldCohosts(ctx).Execute()

Market Open launch shows a verified creator co-hosts that are on now or opening soon: the key, its phase, the show, the creator (named only while their vault profile is public), their room link and, when the creator is in the creator programme, the disclosure that they may earn a share of the fee Crossly keeps. Never a prompt to trade.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketWorldCohosts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketWorldCohosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketWorldCohosts`: ListVaultMarketWorldCohostsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketWorldCohosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketWorldCohostsRequest struct via the builder pattern


### Return type

**ListVaultMarketWorldCohostsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketWorldCreatorRooms

> ListVaultMarketWorldCreatorRoomsResponse ListVaultMarketWorldCreatorRooms(ctx).Execute()

Verified creators whose public vault shelf is a world room: title, the creator, the room link and the disclosure. Rooms show each market share of the shelf, never amounts.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketWorldCreatorRooms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketWorldCreatorRooms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketWorldCreatorRooms`: ListVaultMarketWorldCreatorRoomsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketWorldCreatorRooms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketWorldCreatorRoomsRequest struct via the builder pattern


### Return type

**ListVaultMarketWorldCreatorRoomsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketWorldMarketOpen

> ListVaultMarketWorldMarketOpenResponse ListVaultMarketWorldMarketOpen(ctx).Execute()

Market Open events in the Crossly World Commons: keys opening soon and those that opened in the last few hours, live first then soonest, each with its board line, units on the shelf, dealers quoting and the launch show when there is one. Never a prompt to trade.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultMarketWorldMarketOpen(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultMarketWorldMarketOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketWorldMarketOpen`: ListVaultMarketWorldMarketOpenResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultMarketWorldMarketOpen`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketWorldMarketOpenRequest struct via the builder pattern


### Return type

**ListVaultMarketWorldMarketOpenResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultReceiptTokenWallets

> ListVaultReceiptTokenWalletsResponse ListVaultReceiptTokenWallets(ctx).Execute()

Your verified Solana wallets and whether each is on the receipt-token allow-list (receipt tokens move only between allow-listed wallets). meta.state is the gate.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultReceiptTokenWallets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultReceiptTokenWallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultReceiptTokenWallets`: ListVaultReceiptTokenWalletsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultReceiptTokenWallets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultReceiptTokenWalletsRequest struct via the builder pattern


### Return type

**ListVaultReceiptTokenWalletsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultReceiptTokens

> ListVaultReceiptTokensResponse ListVaultReceiptTokens(ctx).Execute()

Your vault units with their warehouse receipt and receipt-token state (not exported, waiting for a signature, in a wallet, frozen, redeemed), whether each can be exported now and why not. meta.state is the gate (open, missing envs, program id, authority, terms version, cluster); meta.history is your last 50 exports.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultReceiptTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultReceiptTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultReceiptTokens`: ListVaultReceiptTokensResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultReceiptTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultReceiptTokensRequest struct via the builder pattern


### Return type

**ListVaultReceiptTokensResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultReceipts

> ListVaultReceiptsResponse ListVaultReceipts(ctx).Execute()

Your warehouse receipts (documents of title for your vault units), live first. meta.receipts says whether receipts are registered (VAULT_WAREHOUSE_ISSUER) and their terms version.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultReceipts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultReceipts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultReceipts`: ListVaultReceiptsResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultReceipts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultReceiptsRequest struct via the builder pattern


### Return type

**ListVaultReceiptsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultUnitList

> ListVaultUnitListResponse ListVaultUnitList(ctx).Execute()

The published unit list (bar-list equivalent): every unit on a vault shelf with its warehouse receipt number (when receipts are registered), book key, shelf location code, insured value, status, pledged / on-loan flags and a per-receipt hashed holder (never a name or id). Filter with `bookKey`; pages with `page` / `limit`.

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
    resp, r, err := apiClient.VaultMarketsApi.ListVaultUnitList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.ListVaultUnitList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultUnitList`: ListVaultUnitListResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.ListVaultUnitList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultUnitListRequest struct via the builder pattern


### Return type

**ListVaultUnitListResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketAlertPref

> GetVaultMarketAlertPrefResponse UpdateVaultMarketAlertPref(ctx).Execute()

Change vault alert settings: `matrix` ({ kind: { channel: true|false } }, merged onto what is set), `quietHours` ({ start: \"HH:MM\", end: \"HH:MM\", tz } or null to clear) and `telegram` (the Telegram master switch). Locked switches are ignored. Returns the new settings.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketAlertPref(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketAlertPref``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketAlertPref`: GetVaultMarketAlertPrefResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketAlertPref`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketAlertPrefRequest struct via the builder pattern


### Return type

**GetVaultMarketAlertPrefResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketBlockAllocation

> ListVaultMarketBlocksItem UpdateVaultMarketBlockAllocation(ctx, orderId).Execute()

Post-trade allocation: set or replace the plan for a buy you placed for one of the entity's sub-accounts (that sub-account leads). Refused once any share has moved.

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
    orderId := "orderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketBlockAllocation(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketBlockAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketBlockAllocation`: ListVaultMarketBlocksItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketBlockAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketBlockAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketBlocksItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketChannel

> GetVaultMarketChannelResponse UpdateVaultMarketChannel(ctx).Execute()

Change the switches: `telegramAlerts`, `telegramDigest`, `pushCapPerDay` (0 to 20). Margin calls and recalls are never capped. Returns the new view.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketChannel(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketChannel`: GetVaultMarketChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketChannel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketChannelRequest struct via the builder pattern


### Return type

**GetVaultMarketChannelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketDealerBand

> ListVaultMarketDealerBandsItem UpdateVaultMarketDealerBand(ctx).Execute()

Commit to an Instant Exit band on a key: a standing bid for `units` kept at `bandBps` of the mark, never above `maxPriceCents`, re-quoted as the mark moves.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketDealerBand(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketDealerBand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketDealerBand`: ListVaultMarketDealerBandsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketDealerBand`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketDealerBandRequest struct via the builder pattern


### Return type

**ListVaultMarketDealerBandsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketDealerPool

> ListVaultMarketDealerPoolsItem UpdateVaultMarketDealerPool(ctx, id).Execute()

Change a curve pool: spot, step, fee, levels, units, buy budget, floor, ceiling, or pause / resume. Re-quoted at once.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketDealerPool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketDealerPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketDealerPool`: ListVaultMarketDealerPoolsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketDealerPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketDealerPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketDealerPoolsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketEntityMember

> GetVaultMarketEntityByIdResponse UpdateVaultMarketEntityMember(ctx, id, memberId).Execute()

Admin: change a member's role, limits or sub-accounts. The last admin cannot be demoted.

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
    id := "id_example" // string | 
    memberId := "memberId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketEntityMember(context.Background(), id, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketEntityMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketEntityMember`: GetVaultMarketEntityByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketEntityMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketEntityMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetVaultMarketEntityByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketEntitySubaccountLimit

> GetVaultMarketEntityByIdResponse UpdateVaultMarketEntitySubaccountLimit(ctx, id, subId).Execute()

Admin: set a sub-account's limits (per-order notional, units per key, gross, net, leverage), held to the entity's caps. `limits: null` clears them.

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
    id := "id_example" // string | 
    subId := "subId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketEntitySubaccountLimit(context.Background(), id, subId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketEntitySubaccountLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketEntitySubaccountLimit`: GetVaultMarketEntityByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketEntitySubaccountLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**subId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketEntitySubaccountLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetVaultMarketEntityByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketEntityUsdElection

> GetVaultMarketEntityByIdResponse UpdateVaultMarketEntityUsdElection(ctx, id).Execute()

Admin: elect USD for the entity (`usdOnly`). Its sub-accounts then post collateral in USD only and pay every fee in USD, and any CBX reward paid to one of them is converted to USD by the custodian in that account. Audited; it can be turned off again.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketEntityUsdElection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketEntityUsdElection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketEntityUsdElection`: GetVaultMarketEntityByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketEntityUsdElection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketEntityUsdElectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketEntityByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketFollowCopyWatch

> DeleteVaultMarketFollowCopyWatchResponse UpdateVaultMarketFollowCopyWatch(ctx, handle).Execute()

Copy-watch a collector who shares bids and asks: you get an alert when they rest a new one. It never places an order for you.

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketFollowCopyWatch(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketFollowCopyWatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketFollowCopyWatch`: DeleteVaultMarketFollowCopyWatchResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketFollowCopyWatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketFollowCopyWatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteVaultMarketFollowCopyWatchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketFollowSharing

> GetVaultMarketFollowSharingResponse UpdateVaultMarketFollowSharing(ctx).Execute()

Turn sharing switches on or off (`isPublic`, `showBuys`, `deposits`, `sales`, `hallLoans`, `orders`, `leaderboard`). A switch shows only what happens after it goes on; all but the profile itself need the profile public.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketFollowSharing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketFollowSharing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketFollowSharing`: GetVaultMarketFollowSharingResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketFollowSharing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketFollowSharingRequest struct via the builder pattern


### Return type

**GetVaultMarketFollowSharingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketHomeUnit

> UpdateVaultMarketHomeUnitResponse UpdateVaultMarketHomeUnit(ctx).Execute()

Show exactly these vault units in your Crossly World home: { units: [{ vaultItemId, visitors }] }. Units left out come out of their cases; units already standing stay where they are; new ones go into a free display-case or shelf slot, else a glass case of their own along a wall. visitors: true lets people who may enter your home see that unit (default false: only you). Returns placed, unplaced (no free case: the home is full), refused (with why) and the desk. Display only: nothing moves.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketHomeUnit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketHomeUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketHomeUnit`: UpdateVaultMarketHomeUnitResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketHomeUnit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketHomeUnitRequest struct via the builder pattern


### Return type

**UpdateVaultMarketHomeUnitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketIndex

> GetVaultMarketIndexResponse UpdateVaultMarketIndex(ctx, slug).Execute()

Rename, re-weight, share or change the keys of an index of your own; its level carries on.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketIndex(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketIndex`: GetVaultMarketIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketIndexResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketLendingElection

> GetVaultMarketLendingElectionResponse UpdateVaultMarketLendingElection(ctx).Execute()

Change your lender election (applies to loans opened after the change). Needs lending.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketLendingElection(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketLendingElection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketLendingElection`: GetVaultMarketLendingElectionResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketLendingElection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketLendingElectionRequest struct via the builder pattern


### Return type

**GetVaultMarketLendingElectionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketMmMmp

> GetVaultMarketMmMmpResponse UpdateVaultMarketMmMmp(ctx).Execute()

Set market-maker protection: N units or $X filled on your quotes within T seconds pulls every quote and freezes quoting for a set time.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketMmMmp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketMmMmp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketMmMmp`: GetVaultMarketMmMmpResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketMmMmp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketMmMmpRequest struct via the builder pattern


### Return type

**GetVaultMarketMmMmpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketMode

> GetVaultMarketModeResponse UpdateVaultMarketMode(ctx).Execute()

Choose your Vault Markets mode: `mode` simple or pro. Remembered for you on the web and in the app. A display preference only: it places nothing and moves no money.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketMode(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketMode`: GetVaultMarketModeResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketMode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketModeRequest struct via the builder pattern


### Return type

**GetVaultMarketModeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketOfferAlert

> UpdateVaultMarketOfferAlertResponse UpdateVaultMarketOfferAlert(ctx).Execute()

Set (or clear with null) the bid price that should always notify you on a key you hold, and mute or unmute offer notices on it.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketOfferAlert(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketOfferAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketOfferAlert`: UpdateVaultMarketOfferAlertResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketOfferAlert`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketOfferAlertRequest struct via the builder pattern


### Return type

**UpdateVaultMarketOfferAlertResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketPaperLeaderboard

> UpdateVaultMarketPaperLeaderboardResponse UpdateVaultMarketPaperLeaderboard(ctx).Execute()

Join or leave the weekly paper leaderboard. Joining needs a display name (2 to 24 letters, numbers, spaces, dots, dashes or underscores), never your real name. Returns `{ optIn, displayName }`.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketPaperLeaderboard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketPaperLeaderboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketPaperLeaderboard`: UpdateVaultMarketPaperLeaderboardResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketPaperLeaderboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketPaperLeaderboardRequest struct via the builder pattern


### Return type

**UpdateVaultMarketPaperLeaderboardResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketPositionExit

> UpdateVaultMarketPositionExitResponse UpdateVaultMarketPositionExit(ctx, id).Execute()

Set a position's take-profit / stop-loss: a long is closed, a short covered, when the last trade reaches one. Both null clears them.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketPositionExit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketPositionExit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketPositionExit`: UpdateVaultMarketPositionExitResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketPositionExit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketPositionExitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateVaultMarketPositionExitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketPriceWidget

> ListVaultMarketPriceWidgetsItem UpdateVaultMarketPriceWidget(ctx, id).Execute()

Change a widget: name, range, layout, theme, affiliate code (null for none) or allowed sites. The market or index it shows never changes: make a new widget for another.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketPriceWidget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketPriceWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketPriceWidget`: ListVaultMarketPriceWidgetsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketPriceWidget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketPriceWidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketPriceWidgetsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketProfile

> GetVaultMarketProfileResponse UpdateVaultMarketProfile(ctx).Execute()

Turn your public vault profile on or off (`isPublic`) and whether followers see your new buys (`showBuys`). It shows weights by market, never dollars.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketProfile`: GetVaultMarketProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketProfileRequest struct via the builder pattern


### Return type

**GetVaultMarketProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketRecurring

> ListVaultMarketRecurringItem UpdateVaultMarketRecurring(ctx, id).Execute()

Change a recurring buy's schedule, caps, cards or auto-lend, or pause / resume it.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketRecurring(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketRecurring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketRecurring`: ListVaultMarketRecurringItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketRecurring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketRecurringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketRecurringItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketSetGoal

> GetVaultMarketSetGoalResponse UpdateVaultMarketSetGoal(ctx, id).Execute()

Change what a piece must be to count (condition, grade mode, grade). Refused while a basket bid works: cancel it first.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketSetGoal(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketSetGoal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketSetGoal`: GetVaultMarketSetGoalResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketSetGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketSetGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketSetGoalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketShareCard

> ListVaultMarketShareCardsItem UpdateVaultMarketShareCard(ctx, id).Execute()

Show or hide the dollar amounts on one of your share cards (`showAmounts`).

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketShareCard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketShareCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketShareCard`: ListVaultMarketShareCardsItem
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketShareCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketShareCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListVaultMarketShareCardsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketSupplyDeposit

> UpdateVaultMarketSupplyDepositResponse UpdateVaultMarketSupplyDeposit(ctx, id).Execute()

Change a deposit's toggles: `autoListBps` (list at that share of the mark once stored; null = do not) and `lendWhenIdle`.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketSupplyDeposit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketSupplyDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketSupplyDeposit`: UpdateVaultMarketSupplyDepositResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketSupplyDeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketSupplyDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateVaultMarketSupplyDepositResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketSupplyRoundUp

> GetVaultMarketSupplyRoundUpResponse UpdateVaultMarketSupplyRoundUp(ctx).Execute()

Set the round-up: `shareBps` (1%-50%) of each payout released to you adds to a budget; when it covers the cheapest vault unit on `bookKey`, one unit is bought with your saved card, never above the budget.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketSupplyRoundUp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketSupplyRoundUp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketSupplyRoundUp`: GetVaultMarketSupplyRoundUpResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketSupplyRoundUp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketSupplyRoundUpRequest struct via the builder pattern


### Return type

**GetVaultMarketSupplyRoundUpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketTodBeneficiary

> UpdateVaultMarketTodBeneficiaryResponse UpdateVaultMarketTodBeneficiary(ctx).Execute()

Replace your beneficiary list (keep an id to edit one): full name, contact (email or @username), relationship, and share in basis points; shares add up to 10000. Returns `{ beneficiaries, removed }`.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketTodBeneficiary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketTodBeneficiary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketTodBeneficiary`: UpdateVaultMarketTodBeneficiaryResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketTodBeneficiary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketTodBeneficiaryRequest struct via the builder pattern


### Return type

**UpdateVaultMarketTodBeneficiaryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketTodDesignation

> UpdateVaultMarketTodDesignationResponse UpdateVaultMarketTodDesignation(ctx).Execute()

Leave one of your units to one beneficiary (beneficiaryId null clears it; the unit then follows the shares). Returns `{ vaultItemId, beneficiaryId }`.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketTodDesignation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketTodDesignation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketTodDesignation`: UpdateVaultMarketTodDesignationResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketTodDesignation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketTodDesignationRequest struct via the builder pattern


### Return type

**UpdateVaultMarketTodDesignationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketWorldRoomVaultGate

> UpdateVaultMarketWorldRoomVaultGateResponse UpdateVaultMarketWorldRoomVaultGate(ctx, id).Execute()

Set the vault-holder door on your own room (its collection rooms follow): visitors must hold minValueCents (cents, at reliable A/B marks) and/or minUnits units in the Crossly vault, mode any (either) or all (both), optionally only one product (scope kind product, ref the catalogue product id), bullion family (bullion, the family slug) or category (category, the slug). Send null to open it again. Returns `{ gate }`. Moves no money.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketWorldRoomVaultGate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketWorldRoomVaultGate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketWorldRoomVaultGate`: UpdateVaultMarketWorldRoomVaultGateResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketWorldRoomVaultGate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketWorldRoomVaultGateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateVaultMarketWorldRoomVaultGateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVaultMarketWorldSpaceVaultGate

> UpdateVaultMarketWorldRoomVaultGateResponse UpdateVaultMarketWorldSpaceVaultGate(ctx, slug).Execute()

Set the vault-holder door on a Space you host: the same rule as a room (minValueCents at reliable marks and/or minUnits, mode, optional scope). Send null to open it again. Returns `{ gate }`. Moves no money.

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
    resp, r, err := apiClient.VaultMarketsApi.UpdateVaultMarketWorldSpaceVaultGate(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultMarketsApi.UpdateVaultMarketWorldSpaceVaultGate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultMarketWorldSpaceVaultGate`: UpdateVaultMarketWorldRoomVaultGateResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultMarketsApi.UpdateVaultMarketWorldSpaceVaultGate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultMarketWorldSpaceVaultGateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateVaultMarketWorldRoomVaultGateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

