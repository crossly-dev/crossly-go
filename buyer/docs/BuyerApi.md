# \BuyerApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBuyerActivity**](BuyerApi.md#CreateBuyerActivity) | **Post** /v1/buyer/activity | Report an item your user is looking at, and get our answer.
[**CreateBuyerAuctionBid**](BuyerApi.md#CreateBuyerAuctionBid) | **Post** /v1/buyer/auctions/{slug}/bids | Place a bid, optionally with a proxy maximum. A WINNING BID IS CHARGED AUTOMATICALLY.
[**CreateBuyerAuctionPay**](BuyerApi.md#CreateBuyerAuctionPay) | **Post** /v1/buyer/auctions/{slug}/pay | Pay for an auction you won whose automatic charge did not go through.
[**CreateBuyerAuctionRetract**](BuyerApi.md#CreateBuyerAuctionRetract) | **Post** /v1/buyer/auctions/{slug}/retract | Retract your bids under eBay’s rules — a typo, a changed listing, or an unreachable seller.
[**CreateBuyerBoard**](BuyerApi.md#CreateBuyerBoard) | **Post** /v1/buyer/boards | Create an outfit board.
[**CreateBuyerBoardItem**](BuyerApi.md#CreateBuyerBoardItem) | **Post** /v1/buyer/boards/{id}/items | Add a listing (by slug) to a board, at the next free spot.
[**CreateBuyerBreakSpotHold**](BuyerApi.md#CreateBuyerBreakSpotHold) | **Post** /v1/buyer/breaks/{id}/spots/{spotId}/hold | Hold a spot in a pick-your-spot break, then buy it through checkout before the hold lapses.
[**CreateBuyerBuybackRequest**](BuyerApi.md#CreateBuyerBuybackRequest) | **Post** /v1/buyer/buybacks/{id}/request | Ask Crossly to buy an item back under its buyback guarantee.
[**CreateBuyerCartItem**](BuyerApi.md#CreateBuyerCartItem) | **Post** /v1/buyer/cart/items | Add a listing to your cart.
[**CreateBuyerCartItemSaveForLater**](BuyerApi.md#CreateBuyerCartItemSaveForLater) | **Post** /v1/buyer/cart/items/{id}/save-for-later | Move a cart line to saved-for-later.
[**CreateBuyerCartOptimizer**](BuyerApi.md#CreateBuyerCartOptimizer) | **Post** /v1/buyer/cart-optimizer | Cart optimizer: send a shopping list (one item per line, \&quot;2 Charizard Base Set\&quot;) and get the cheapest delivered basket across sellers — one pick per line, items + each seller’s shipping with free-shipping thresholds applied. fewestPackages prefers fewer sellers. Nothing is added to the cart.
[**CreateBuyerCartQuote**](BuyerApi.md#CreateBuyerCartQuote) | **Post** /v1/buyer/cart/quote | Price the cart, delivered — item, shipping, tax, total.
[**CreateBuyerCartSavedMoveToCart**](BuyerApi.md#CreateBuyerCartSavedMoveToCart) | **Post** /v1/buyer/cart/saved/{id}/move-to-cart | Move a saved item back into your cart, at today’s price.
[**CreateBuyerCaseCounter**](BuyerApi.md#CreateBuyerCaseCounter) | **Post** /v1/buyer/cases/{id}/counter | Propose your own refund figure on a case.
[**CreateBuyerCaseEscalate**](BuyerApi.md#CreateBuyerCaseEscalate) | **Post** /v1/buyer/cases/{id}/escalate | Ask Crossly to step in on a case.
[**CreateBuyerCaseMessage**](BuyerApi.md#CreateBuyerCaseMessage) | **Post** /v1/buyer/cases/{id}/messages | Post a message on a case thread.
[**CreateBuyerCaseRespond**](BuyerApi.md#CreateBuyerCaseRespond) | **Post** /v1/buyer/cases/{id}/respond | Accept or reject the seller’s partial-refund offer on a case.
[**CreateBuyerCaseWithdraw**](BuyerApi.md#CreateBuyerCaseWithdraw) | **Post** /v1/buyer/cases/{id}/withdraw | Withdraw a case you opened.
[**CreateBuyerCatalogListingRental**](BuyerApi.md#CreateBuyerCatalogListingRental) | **Post** /v1/buyer/catalog/listings/{slug}/rental | Book a rental. Returns a Stripe PaymentIntent client secret for rent + shipping both ways + deposit; the dates are held 30 minutes for payment. Crossly buys both labels; the deposit comes back after the item does.
[**CreateBuyerCatalogListingRentalQuote**](BuyerApi.md#CreateBuyerCatalogListingRentalQuote) | **Post** /v1/buyer/catalog/listings/{slug}/rental/quote | Price a rental: rent (the cheaper of by-the-day and by-the-week), both shipping legs once a ship-to is given, and the deposit. &#x60;problem&#x60; says why the dates can’t be booked.
[**CreateBuyerCatalogListingVideoRequest**](BuyerApi.md#CreateBuyerCatalogListingVideoRequest) | **Post** /v1/buyer/catalog/listings/{slug}/video-requests | Ask the seller for a video of an item (one open request per item). The seller is notified; their answer appears on the listing for every buyer.
[**CreateBuyerConversationInvite**](BuyerApi.md#CreateBuyerConversationInvite) | **Post** /v1/buyer/conversations/{id}/invite | Invite a friend into a chat with a seller: a single-use link (7 days) that makes whoever opens it a guest who can read and post. Only the thread’s buyer can invite; at most 5 guests.
[**CreateBuyerConversationJoin**](BuyerApi.md#CreateBuyerConversationJoin) | **Post** /v1/buyer/conversations/join/{token} | Join a chat you were invited into, as a guest.
[**CreateBuyerFeedback**](BuyerApi.md#CreateBuyerFeedback) | **Post** /v1/buyer/feedback | Leave feedback on up to 25 orders at once, with optional detailed seller ratings.
[**CreateBuyerFitRef**](BuyerApi.md#CreateBuyerFitRef) | **Post** /v1/buyer/fit/refs | Save a reference garment (flat measurements, inches) to compare listings against.
[**CreateBuyerFollowingSellerSeen**](BuyerApi.md#CreateBuyerFollowingSellerSeen) | **Post** /v1/buyer/following/sellers/seen | Clear the \&quot;new listings\&quot; count on a followed seller, or on all of them.
[**CreateBuyerGiftCard**](BuyerApi.md#CreateBuyerGiftCard) | **Post** /v1/buyer/gift-cards | Start a gift card purchase.
[**CreateBuyerGiftCardRedeem**](BuyerApi.md#CreateBuyerGiftCardRedeem) | **Post** /v1/buyer/gift-cards/redeem | Redeem a gift card code into your own balance.
[**CreateBuyerGroupBuyJoin**](BuyerApi.md#CreateBuyerGroupBuyJoin) | **Post** /v1/buyer/group-buys/{id}/join | Join a group buy: a hold at the starting price, captured at the price the group reaches. Checked against this key’s spending controls; send an Idempotency-Key.
[**CreateBuyerGroupBuyLeave**](BuyerApi.md#CreateBuyerGroupBuyLeave) | **Post** /v1/buyer/group-buys/{id}/leave | Leave a group buy and release your hold.
[**CreateBuyerHunt**](BuyerApi.md#CreateBuyerHunt) | **Post** /v1/buyer/hunts | Start a buyer-agent hunt — over a search (filters), or auto-buy on one listing (listingSlug: buy it once its price, sale applied, is within maxPriceCents). &#x60;offer&#x60; and &#x60;buy&#x60; modes also need buyer:checkout:write, and the budget is checked against this key’s spending controls.
[**CreateBuyerInstantOfferAccept**](BuyerApi.md#CreateBuyerInstantOfferAccept) | **Post** /v1/buyer/instant-offers/{id}/accept | Accept Crossly’s instant offer on a collection item. Ops then send a prepaid label; the store credit is paid once the item arrives and checks out.
[**CreateBuyerLayawayCancel**](BuyerApi.md#CreateBuyerLayawayCancel) | **Post** /v1/buyer/layaways/{id}/cancel | Cancel a layaway: everything paid comes back except the layaway fee (10% of the price, at most $15).
[**CreateBuyerLayawayPay**](BuyerApi.md#CreateBuyerLayawayPay) | **Post** /v1/buyer/layaways/{id}/pay | Pay some (at least $1) or all of what is left now, on the saved card. Paid in full, it ships.
[**CreateBuyerListingAudioSummary**](BuyerApi.md#CreateBuyerListingAudioSummary) | **Post** /v1/buyer/listings/{slug}/audio-summary | Record the spoken summary of a listing (cached per listing version). 503 naming TTS_API_KEY when no speech provider is configured.
[**CreateBuyerListingSplit**](BuyerApi.md#CreateBuyerListingSplit) | **Post** /v1/buyer/listings/{slug}/split | Split a purchase: one pay link per person; nobody is charged until all have paid.
[**CreateBuyerListingSubscribe**](BuyerApi.md#CreateBuyerListingSubscribe) | **Post** /v1/buyer/listings/{slug}/subscribe | Subscribe &amp; save: a standing order charged to your saved card every interval. Checked against this key’s spending controls; send an Idempotency-Key.
[**CreateBuyerListingTryOn**](BuyerApi.md#CreateBuyerListingTryOn) | **Post** /v1/buyer/listings/{slug}/try-on | Render an AI try-on of this listing on the buyer’s own photo (base64 &#x60;photo&#x60;, or &#x60;reuseLastPhoto&#x60;). &#x60;confirmSelfAdult&#x60; must be true: the photo is of the buyer and they are 18+. Wearables only; capped per day; 503 naming TRYON_API_KEY when no image model is configured.
[**CreateBuyerMarketBid**](BuyerApi.md#CreateBuyerMarketBid) | **Post** /v1/buyer/market/bids | Place a fixed-price market bid. SPENDS MONEY: it fills against the lowest ask at or under your price and charges your saved card (plus the fast-ship premium when fastShip is set and it fills from vault stock). deliverTo vault keeps it in your Crossly vault; gradeWith sends a raw card to a Grading run first. Checked against this key’s spending controls; send an Idempotency-Key.
[**CreateBuyerMarketWatch**](BuyerApi.md#CreateBuyerMarketWatch) | **Post** /v1/buyer/market/watches | Watch a market variant: alert at or below a price, or on any drop.
[**CreateBuyerOffer**](BuyerApi.md#CreateBuyerOffer) | **Post** /v1/buyer/offers | Offer a price on a listing.
[**CreateBuyerOfferRespond**](BuyerApi.md#CreateBuyerOfferRespond) | **Post** /v1/buyer/offers/{id}/respond | Answer a seller’s counteroffer — accept, decline, or counter back.
[**CreateBuyerOrderCancel**](BuyerApi.md#CreateBuyerOrderCancel) | **Post** /v1/buyer/orders/{id}/cancel | Cancel an order before it ships — refunds it in full.
[**CreateBuyerOrderInr**](BuyerApi.md#CreateBuyerOrderInr) | **Post** /v1/buyer/orders/{id}/inr | Open an item-not-received case.
[**CreateBuyerOrderReceived**](BuyerApi.md#CreateBuyerOrderReceived) | **Post** /v1/buyer/orders/{id}/received | Confirm you received an order.
[**CreateBuyerOrderTip**](BuyerApi.md#CreateBuyerOrderTip) | **Post** /v1/buyer/orders/{id}/tip | Tip the seller on one of your shipped orders, charged to your saved card. One tip per order; $1–$500; checked against this key’s spending controls.
[**CreateBuyerPackFairnessRotate**](BuyerApi.md#CreateBuyerPackFairnessRotate) | **Post** /v1/buyer/packs/fairness/rotate | Reveal your current server seed and start a new one (optionally with your own client seed). Refused while an opening is awaiting payment.
[**CreateBuyerPackOpen**](BuyerApi.md#CreateBuyerPackOpen) | **Post** /v1/buyer/packs/{slug}/open | Open a pack. SPENDS MONEY: charges the pack price to your saved card, then pulls one real item by the published odds (provably fair: see /buyer/packs/fairness) into your Crossly vault. confirmAdult must be true. Checked against this key’s spending controls; send an Idempotency-Key.
[**CreateBuyerPackOpeningSellBack**](BuyerApi.md#CreateBuyerPackOpeningSellBack) | **Post** /v1/buyer/pack-openings/{id}/sell-back | Sell a pull straight back for Crossly credit at the pack’s buyback rate, within its window, while it is still in your vault.
[**CreateBuyerPassportClaim**](BuyerApi.md#CreateBuyerPassportClaim) | **Post** /v1/buyer/passports/{code}/claim | Claim a passport as the item’s current owner (the buyer of its latest sale) — shown as “in its owner’s hands” until it sells again.
[**CreateBuyerProtectionPlanCancel**](BuyerApi.md#CreateBuyerProtectionPlanCancel) | **Post** /v1/buyer/protection-plans/{id}/cancel | Cancel a protection plan. The insurer decides what is refundable (pro-rata after the free-look window); that amount goes back to the card.
[**CreateBuyerPurchaseDownload**](BuyerApi.md#CreateBuyerPurchaseDownload) | **Post** /v1/buyer/purchases/{id}/download | A ten-minute link to a digital purchase’s file; each call uses one of its downloads.
[**CreateBuyerPurchaseFit**](BuyerApi.md#CreateBuyerPurchaseFit) | **Post** /v1/buyer/purchases/{id}/fit | Say how a delivered purchase fit: true_to_size, small or large. It teaches your size memory and, pooled with other buyers, how the brand runs.
[**CreateBuyerPurchaseReturn**](BuyerApi.md#CreateBuyerPurchaseReturn) | **Post** /v1/buyer/purchases/{id}/return | Return a purchase under the seller’s policy, or send back a try-at-home item. With reason doesnt_fit, &#x60;fit&#x60; (small | large) says which way, and it feeds your size memory.
[**CreateBuyerRegistry**](BuyerApi.md#CreateBuyerRegistry) | **Post** /v1/buyer/registries | Start a gift registry (wedding, baby, birthday, housewarming, graduation, other). Items are added like any wishlist item; set shipAddressId to one of your saved addresses so guests can buy — they never see it.
[**CreateBuyerRegistryGiftThanked**](BuyerApi.md#CreateBuyerRegistryGiftThanked) | **Post** /v1/buyer/registries/{id}/gifts/{giftId}/thanked | Tick a gift as thanked (or untick it).
[**CreateBuyerRentalArrivalPhoto**](BuyerApi.md#CreateBuyerRentalArrivalPhoto) | **Post** /v1/buyer/rentals/{id}/arrival-photos | Photograph a rental on arrival (URLs you uploaded): your side of the condition record.
[**CreateBuyerRentalCancel**](BuyerApi.md#CreateBuyerRentalCancel) | **Post** /v1/buyer/rentals/{id}/cancel | Cancel a rental that has not shipped: everything is refunded.
[**CreateBuyerRentalClaimResponse**](BuyerApi.md#CreateBuyerRentalClaimResponse) | **Post** /v1/buyer/rentals/{id}/claim-response | Answer a deposit claim: accept (the claim is paid, the rest refunded) or dispute (Crossly decides).
[**CreateBuyerRentalIdentityStart**](BuyerApi.md#CreateBuyerRentalIdentityStart) | **Post** /v1/buyer/rentals/identity/start | Start the ID check renting needs (Stripe Identity: photo ID + selfie). Open &#x60;url&#x60; (hosted) or use &#x60;clientSecret&#x60; with Stripe.js.
[**CreateBuyerReturnCancel**](BuyerApi.md#CreateBuyerReturnCancel) | **Post** /v1/buyer/returns/{id}/cancel | Cancel a return before it ships.
[**CreateBuyerReturnTracking**](BuyerApi.md#CreateBuyerReturnTracking) | **Post** /v1/buyer/returns/{id}/tracking | Add tracking to a return you shipped on your own label.
[**CreateBuyerSellerClubJoin**](BuyerApi.md#CreateBuyerSellerClubJoin) | **Post** /v1/buyer/sellers/{sellerUserId}/club/join | Join a seller’s loyalty club (free). Your tier’s discount then applies to their items — the best of it and any other rate you have with them, never stacked.
[**CreateBuyerSellerReport**](BuyerApi.md#CreateBuyerSellerReport) | **Post** /v1/buyer/sellers/{username}/report | Report a seller to Crossly’s trust team.
[**CreateBuyerSplitCancel**](BuyerApi.md#CreateBuyerSplitCancel) | **Post** /v1/buyer/splits/{id}/cancel | Cancel a split you organised; every held share is released.
[**CreateBuyerTogether**](BuyerApi.md#CreateBuyerTogether) | **Post** /v1/buyer/together | Start a shop-together room.
[**CreateBuyerTogetherItem**](BuyerApi.md#CreateBuyerTogetherItem) | **Post** /v1/buyer/together/{code}/items | Add a listing to a shop-together board.
[**CreateBuyerTogetherItemVote**](BuyerApi.md#CreateBuyerTogetherItemVote) | **Post** /v1/buyer/together/{code}/items/{id}/vote | Vote an item on a room’s board up (1), down (-1) or clear (0).
[**CreateBuyerTogetherJoin**](BuyerApi.md#CreateBuyerTogetherJoin) | **Post** /v1/buyer/together/{code}/join | Join a shop-together room by its code.
[**CreateBuyerTogetherLeave**](BuyerApi.md#CreateBuyerTogetherLeave) | **Post** /v1/buyer/together/{code}/leave | Leave a shop-together room.
[**CreateBuyerTogetherMessage**](BuyerApi.md#CreateBuyerTogetherMessage) | **Post** /v1/buyer/together/{code}/messages | Say something in a shop-together room.
[**CreateBuyerTotalRequest**](BuyerApi.md#CreateBuyerTotalRequest) | **Post** /v1/buyer/total-requests | Ask a seller for one combined shipping total on everything of theirs in your cart.
[**CreateBuyerVaultList**](BuyerApi.md#CreateBuyerVaultList) | **Post** /v1/buyer/vault/{id}/list | Sell a vault item on the market at your price. It is already authenticated, so when it sells nothing ships in: the buyer keeps it vaulted or it ships straight to them.
[**CreateBuyerVaultUnlist**](BuyerApi.md#CreateBuyerVaultUnlist) | **Post** /v1/buyer/vault/{id}/unlist | Take a vault item off the market; it stays stored.
[**CreateBuyerVaultWithdraw**](BuyerApi.md#CreateBuyerVaultWithdraw) | **Post** /v1/buyer/vault/{id}/withdraw | Ship a vault item home, standard or expedited. The shipping, any storage it owes and the early-withdrawal fee (within 90 days of arriving) are charged to your saved card now, all checked against this key’s spending controls; send an Idempotency-Key.
[**CreateBuyerWishlist**](BuyerApi.md#CreateBuyerWishlist) | **Post** /v1/buyer/wishlists | Create a wishlist.
[**CreateBuyerWishlistInvite**](BuyerApi.md#CreateBuyerWishlistInvite) | **Post** /v1/buyer/wishlists/{id}/invite | Share a wishlist with a co-editor: a single-use link (14 days). Members add and remove items; only the owner renames, deletes or removes members. At most 20.
[**CreateBuyerWishlistItem**](BuyerApi.md#CreateBuyerWishlistItem) | **Post** /v1/buyer/wishlists/{id}/items | Add a listing to a wishlist.
[**CreateBuyerWishlistJoin**](BuyerApi.md#CreateBuyerWishlistJoin) | **Post** /v1/buyer/wishlists/join/{token} | Join a wishlist you were invited to edit.
[**CreateBuyerWishlistShare**](BuyerApi.md#CreateBuyerWishlistShare) | **Post** /v1/buyer/wishlists/{id}/share | Share a wishlist — makes it public and returns its link.
[**DeleteBuyerBoard**](BuyerApi.md#DeleteBuyerBoard) | **Delete** /v1/buyer/boards/{id} | Delete a board.
[**DeleteBuyerBoardItem**](BuyerApi.md#DeleteBuyerBoardItem) | **Delete** /v1/buyer/boards/{id}/items/{platformListingId} | Take a listing off a board.
[**DeleteBuyerBreakSpotHold**](BuyerApi.md#DeleteBuyerBreakSpotHold) | **Delete** /v1/buyer/breaks/{id}/spots/{spotId}/hold | Release a break spot you hold.
[**DeleteBuyerCartItem**](BuyerApi.md#DeleteBuyerCartItem) | **Delete** /v1/buyer/cart/items/{id} | Remove a line from your cart.
[**DeleteBuyerCartSaved**](BuyerApi.md#DeleteBuyerCartSaved) | **Delete** /v1/buyer/cart/saved/{id} | Remove an item from saved-for-later. Returns the remaining list.
[**DeleteBuyerConversationGuest**](BuyerApi.md#DeleteBuyerConversationGuest) | **Delete** /v1/buyer/conversations/{id}/guests/{userId} | Remove a guest (the thread’s buyer), or leave (a guest, with your own user id).
[**DeleteBuyerFitRef**](BuyerApi.md#DeleteBuyerFitRef) | **Delete** /v1/buyer/fit/refs/{id} | Delete a reference garment.
[**DeleteBuyerMarketBid**](BuyerApi.md#DeleteBuyerMarketBid) | **Delete** /v1/buyer/market/bids/{id} | Cancel a resting market bid.
[**DeleteBuyerMarketWatch**](BuyerApi.md#DeleteBuyerMarketWatch) | **Delete** /v1/buyer/market/watches/{id} | Stop watching a market variant.
[**DeleteBuyerOffer**](BuyerApi.md#DeleteBuyerOffer) | **Delete** /v1/buyer/offers/{id} | Withdraw an offer you made that the seller has not answered.
[**DeleteBuyerRecentlyViewed**](BuyerApi.md#DeleteBuyerRecentlyViewed) | **Delete** /v1/buyer/recently-viewed/{slug} | Remove one listing from your recently viewed.
[**DeleteBuyerRecentlyViewed2**](BuyerApi.md#DeleteBuyerRecentlyViewed2) | **Delete** /v1/buyer/recently-viewed | Clear your recently viewed list.
[**DeleteBuyerRegistry**](BuyerApi.md#DeleteBuyerRegistry) | **Delete** /v1/buyer/registries/{id} | Delete a registry (and the wishlist behind it).
[**DeleteBuyerSellerClubJoin**](BuyerApi.md#DeleteBuyerSellerClubJoin) | **Delete** /v1/buyer/sellers/{sellerUserId}/club/join | Leave a seller’s loyalty club.
[**DeleteBuyerTogetherItem**](BuyerApi.md#DeleteBuyerTogetherItem) | **Delete** /v1/buyer/together/{code}/items/{id} | Take an item you added off a room’s board.
[**DeleteBuyerTotalRequest**](BuyerApi.md#DeleteBuyerTotalRequest) | **Delete** /v1/buyer/total-requests/{id} | Cancel a total request you made.
[**DeleteBuyerTryOn**](BuyerApi.md#DeleteBuyerTryOn) | **Delete** /v1/buyer/try-ons/{id} | Delete a try-on and its images — or pass &#x60;all&#x60; to delete every photo and result.
[**DeleteBuyerWishlist**](BuyerApi.md#DeleteBuyerWishlist) | **Delete** /v1/buyer/wishlists/{id} | Delete a wishlist and everything on it.
[**DeleteBuyerWishlistItem**](BuyerApi.md#DeleteBuyerWishlistItem) | **Delete** /v1/buyer/wishlists/{id}/items/{slug} | Take a listing off a wishlist.
[**DeleteBuyerWishlistMember**](BuyerApi.md#DeleteBuyerWishlistMember) | **Delete** /v1/buyer/wishlists/{id}/members/{userId} | Remove a co-editor (the owner), or leave a list shared with you (your own user id).
[**GetBuyerAuctionMe**](BuyerApi.md#GetBuyerAuctionMe) | **Get** /v1/buyer/auctions/{slug}/me | Where you stand in one auction — your max, leading, won, must you pay.
[**GetBuyerBadge**](BuyerApi.md#GetBuyerBadge) | **Get** /v1/buyer/badges | The counts behind the account badge — unread messages, offers waiting, orders to review.
[**GetBuyerBoard**](BuyerApi.md#GetBuyerBoard) | **Get** /v1/buyer/boards/{id} | One of your boards with its layout: each item’s x, y and width as fractions of the board, and stacking order.
[**GetBuyerBoardShared**](BuyerApi.md#GetBuyerBoardShared) | **Get** /v1/buyer/boards/shared/{token} | A shared outfit board, by its link token.
[**GetBuyerBreak**](BuyerApi.md#GetBuyerBreak) | **Get** /v1/buyer/breaks/{id} | A card break’s board: spots, holders, and the seed after the draw.
[**GetBuyerCase**](BuyerApi.md#GetBuyerCase) | **Get** /v1/buyer/cases/{id} | One case — its state, the message thread, and the return parcel if there is one.
[**GetBuyerCatalogListingMySize**](BuyerApi.md#GetBuyerCatalogListingMySize) | **Get** /v1/buyer/catalog/listings/{slug}/my-size | This item against your size: the size you keep in this kind of item (and in this brand), whether this one is your size, smaller or larger, a warning if this brand and size ran small or large on you before, and how other buyers say the brand runs. null for an unsized item.
[**GetBuyerCatalogListingPrice**](BuyerApi.md#GetBuyerCatalogListingPrice) | **Get** /v1/buyer/catalog/listings/{slug}/price | The price today. When the seller has the same item live on another marketplace for less, this IS that price (&#x60;priceMatch&#x60; says where from), and &#x60;lowestAnywhere&#x60; lists the marketplaces it is guaranteed not to be cheaper on. &#x60;priceCheck&#x60; rates it against real sales of the same catalog item in the same condition over the last 180 days (great / good / fair / high, with the median and middle-half range in cents); null until there are 3+ sales from 2+ sellers.
[**GetBuyerCatalogListingRental**](BuyerApi.md#GetBuyerCatalogListingRental) | **Get** /v1/buyer/catalog/listings/{slug}/rental | Whether a listing is for rent: day and week rates, the refundable deposit, length limits, the turnaround buffer, and the dates already booked. null when it is not for rent.
[**GetBuyerCheckoutUsdc**](BuyerApi.md#GetBuyerCheckoutUsdc) | **Get** /v1/buyer/checkout/usdc | Whether Crossly takes USDC (Stripe stablecoin payments: Solana, Ethereum, Base, Polygon), and its per-order limit. Paying needs the buyer’s wallet, so it happens in the browser checkout, never unattended.
[**GetBuyerCollectionOffer**](BuyerApi.md#GetBuyerCollectionOffer) | **Get** /v1/buyer/collection/{id}/offer | What an item in your collection would fetch now: the top open bid on the market for the same variant, condition and grade, and Crossly’s instant offer in store credit when that is switched on (a quote valid for a limited time).
[**GetBuyerCollectionSet**](BuyerApi.md#GetBuyerCollectionSet) | **Get** /v1/buyer/collection/sets/{category}/{brand}/{model} | One set: every piece (owned or missing, with its lowest ask on the market), what completing it would cost now, and the leaderboard of collectors who opted in, with your rank.
[**GetBuyerCreator**](BuyerApi.md#GetBuyerCreator) | **Get** /v1/buyer/creators/{handle} | A creator storefront: profile, affiliate code, and shelves of live picks.
[**GetBuyerFitVocabulary**](BuyerApi.md#GetBuyerFitVocabulary) | **Get** /v1/buyer/fit/vocabulary | Garment kinds and the measurements each takes.
[**GetBuyerGiftFinderOption**](BuyerApi.md#GetBuyerGiftFinderOption) | **Get** /v1/buyer/gift-finder/options | The gift finder’s questions: recipients, interests, budgets.
[**GetBuyerImpact**](BuyerApi.md#GetBuyerImpact) | **Get** /v1/buyer/impact | The sourced footprint your secondhand purchases avoided.
[**GetBuyerLayaway**](BuyerApi.md#GetBuyerLayaway) | **Get** /v1/buyer/layaways/{id} | One layaway and its schedule.
[**GetBuyerListingAudioSummary**](BuyerApi.md#GetBuyerListingAudioSummary) | **Get** /v1/buyer/listings/{slug}/audio-summary | A spoken summary of a listing (accessibility): the text, built only from the listing’s own fields and the seller’s description, and the recording of it when one exists.
[**GetBuyerListingBoughtTogether**](BuyerApi.md#GetBuyerListingBoughtTogether) | **Get** /v1/buyer/listings/{slug}/bought-together | Frequently bought together / complete the look.
[**GetBuyerListingBreak**](BuyerApi.md#GetBuyerListingBreak) | **Get** /v1/buyer/listings/{slug}/break | The card break this listing sells spots in, if any; the seed is revealed after the draw.
[**GetBuyerListingCert**](BuyerApi.md#GetBuyerListingCert) | **Get** /v1/buyer/listings/{slug}/cert | A graded listing’s cert, checked with the grader: the seller’s claim (grader, grade, cert number), status (verified, mismatch, not_found, or unchecked where the grader has no lookup), population and how many graded higher where the grader publishes them, and the grader’s own cert page.
[**GetBuyerListingFit**](BuyerApi.md#GetBuyerListingFit) | **Get** /v1/buyer/listings/{slug}/fit | A listing’s flat measurements, compared with your saved reference garments.
[**GetBuyerListingGroupBuy**](BuyerApi.md#GetBuyerListingGroupBuy) | **Get** /v1/buyer/listings/{slug}/group-buy | The group buy open on this listing, if any.
[**GetBuyerListingImpact**](BuyerApi.md#GetBuyerListingImpact) | **Get** /v1/buyer/listings/{slug}/impact | The sourced footprint of making a new equivalent; null when no published figure applies.
[**GetBuyerListingResaleValue**](BuyerApi.md#GetBuyerListingResaleValue) | **Get** /v1/buyer/listings/{slug}/resale-value | What the same brand and category resold for on Crossly, and the buyback offer when enabled.
[**GetBuyerListingSubscribeOffer**](BuyerApi.md#GetBuyerListingSubscribeOffer) | **Get** /v1/buyer/listings/{slug}/subscribe-offer | The subscribe &amp; save discount and intervals on a listing; null when the seller has none.
[**GetBuyerListingTryOn**](BuyerApi.md#GetBuyerListingTryOn) | **Get** /v1/buyer/listings/{slug}/try-on | Whether an AI try-on is offered for this listing (a wearable with a photo) and whether the image model is configured.
[**GetBuyerLiveStreamAuction**](BuyerApi.md#GetBuyerLiveStreamAuction) | **Get** /v1/buyer/live-streams/{id}/auction | The auction running on a live stream right now, if any. Bid on it by its &#x60;slug&#x60; at POST /v1/buyer/auctions/:slug/bids.
[**GetBuyerMarketPurchaseOption**](BuyerApi.md#GetBuyerMarketPurchaseOption) | **Get** /v1/buyer/market/purchase-options | Before you bid or buy: whether the lowest ask is vault stock (keep it vaulted, or pay fastShipCents to rush it), and the grading runs a raw card can go straight to (pass gradeWith on the bid).
[**GetBuyerOrderEligibility**](BuyerApi.md#GetBuyerOrderEligibility) | **Get** /v1/buyer/orders/{id}/eligibility | What you can do with one order right now — cancel, confirm receipt, report not received, leave feedback.
[**GetBuyerOrderExport**](BuyerApi.md#GetBuyerOrderExport) | **Get** /v1/buyer/orders/export | Your order history as CSV text (up to 5,000 orders).
[**GetBuyerOrderPriceProtection**](BuyerApi.md#GetBuyerOrderPriceProtection) | **Get** /v1/buyer/orders/{id}/price-protection | Price protection on one order: if the item gets cheaper before &#x60;until&#x60;, the difference is credited to your Bucks automatically. &#x60;creditedCents&#x60; is what it has paid so far.
[**GetBuyerPack**](BuyerApi.md#GetBuyerPack) | **Get** /v1/buyer/packs/{slug} | One pack: tiers and odds, items left, expected value, the chase items still in it.
[**GetBuyerPackFairness**](BuyerApi.md#GetBuyerPackFairness) | **Get** /v1/buyer/packs/fairness | Your provably-fair seeds: the active server-seed commitment (sha256), your client seed and nonce, and every retired seed in full so past openings can be re-run.
[**GetBuyerParty**](BuyerApi.md#GetBuyerParty) | **Get** /v1/buyer/parties/{slug} | One shopping party.
[**GetBuyerPassport**](BuyerApi.md#GetBuyerPassport) | **Get** /v1/buyer/passports/{code} | What a product passport tag opens: the item, whether it was authenticated, how many owners on Crossly and since when — never who or for how much.
[**GetBuyerPreference**](BuyerApi.md#GetBuyerPreference) | **Get** /v1/buyer/preferences | The shopping profile derived from that activity.
[**GetBuyerPriceGuide**](BuyerApi.md#GetBuyerPriceGuide) | **Get** /v1/buyer/price-guide | What matching items actually sold for on Crossly: median, quartiles, by month, and recent sales. Needs q, brand or category.
[**GetBuyerProfile**](BuyerApi.md#GetBuyerProfile) | **Get** /v1/buyer/profile | Your Crossly shopping profile — name, email, saved address, Bucks balance.
[**GetBuyerPurchaseDelivery**](BuyerApi.md#GetBuyerPurchaseDelivery) | **Get** /v1/buyer/purchases/{id}/delivery | Where a purchase is going when not to your door: the pickup point (with its release code) or its ship-together consolidation.
[**GetBuyerPurchaseLiveMoment**](BuyerApi.md#GetBuyerPurchaseLiveMoment) | **Get** /v1/buyer/purchases/{id}/live-moment | The live show a purchase was bought in, with the offset (seconds) into its replay where it sold. null when it was not bought in a show.
[**GetBuyerPurchaseReturn**](BuyerApi.md#GetBuyerPurchaseReturn) | **Get** /v1/buyer/purchases/{id}/return | What a return would refund for each reason, by when, and any return already open.
[**GetBuyerPurchaseStory**](BuyerApi.md#GetBuyerPurchaseStory) | **Get** /v1/buyer/purchases/{id}/story | Your owner story on a purchase, passed on to the next owner if you resell it.
[**GetBuyerRegistry**](BuyerApi.md#GetBuyerRegistry) | **Get** /v1/buyer/registries/{id} | One of your registries: items with how many are wanted and bought.
[**GetBuyerRegistryByToken**](BuyerApi.md#GetBuyerRegistryByToken) | **Get** /v1/buyer/registry/{token} | A registry as a guest sees it: items, how many are still wanted, and whether it can ship. Never the address.
[**GetBuyerRentalEligibility**](BuyerApi.md#GetBuyerRentalEligibility) | **Get** /v1/buyer/rentals/eligibility | Can you rent right now: ID-verified (a photo ID and a matching selfie, once), the verified name rentals ship to, and your standing (one rental at a time until your first comes back; none after one that never did).
[**GetBuyerSellerClub**](BuyerApi.md#GetBuyerSellerClub) | **Get** /v1/buyer/sellers/{sellerUserId}/club | A seller’s loyalty club: its tiers (spend threshold, discount, perks) and your standing. null when the seller has none.
[**GetBuyerSellerTrust**](BuyerApi.md#GetBuyerSellerTrust) | **Get** /v1/buyer/sellers/{username}/trust | Seller at a glance: member since, completed sales, independent rating, typical handling time, dispute rate, trust tier, badges — and the plain sentences the profile shows.
[**GetBuyerTogether**](BuyerApi.md#GetBuyerTogether) | **Get** /v1/buyer/together/{code} | A shop-together room: members, the board with votes, and chat.
[**GetBuyerVaultMarketHistory**](BuyerApi.md#GetBuyerVaultMarketHistory) | **Get** /v1/buyer/vault-markets/history | Your Vault Markets value by day at the marks (oldest first, today live) over &#x60;days&#x60; (default 90), and today’s change against yesterday’s close.
[**GetBuyerVaultMarketIndex**](BuyerApi.md#GetBuyerVaultMarketIndex) | **Get** /v1/buyer/vault-markets/indexes/{slug} | One index with its published daily closes (oldest first, &#x60;days&#x60; default 365) and its level now.
[**GetBuyerVaultMarketInstantExit**](BuyerApi.md#GetBuyerVaultMarketInstantExit) | **Get** /v1/buyer/vault-markets/instant-exit/{id} | The Instant Exit quote for one vault unit you hold, and whether it can be sold now.
[**GetBuyerVaultMarketWorth**](BuyerApi.md#GetBuyerVaultMarketWorth) | **Get** /v1/buyer/vault-markets/worth | What your vault is worth now: value at the market prices (marks come only from trades) split into reliable and indicative marks, cost and P&amp;L, today’s change, and an estimate of what the units free to sell would fetch from the bids resting on the book now. Read only.
[**ListBuyerActivity**](BuyerApi.md#ListBuyerActivity) | **Get** /v1/buyer/activity | What this buyer has compared lately.
[**ListBuyerAuctionBids**](BuyerApi.md#ListBuyerAuctionBids) | **Get** /v1/buyer/auctions/{slug}/bids | An auction’s bid history, highest first, bidders anonymised.
[**ListBuyerAuctions**](BuyerApi.md#ListBuyerAuctions) | **Get** /v1/buyer/auctions | Auctions you have bid in — bidding, won or lost (last 100).
[**ListBuyerBoards**](BuyerApi.md#ListBuyerBoards) | **Get** /v1/buyer/boards | Your outfit boards: collages of listings, each with a share link.
[**ListBuyerBreaks**](BuyerApi.md#ListBuyerBreaks) | **Get** /v1/buyer/breaks | Card-break spots you hold or bought, with what each pulled.
[**ListBuyerBuybacks**](BuyerApi.md#ListBuyerBuybacks) | **Get** /v1/buyer/buybacks | Your buyback guarantees and requests.
[**ListBuyerCart**](BuyerApi.md#ListBuyerCart) | **Get** /v1/buyer/cart | What is in your Crossly cart.
[**ListBuyerCartSaved**](BuyerApi.md#ListBuyerCartSaved) | **Get** /v1/buyer/cart/saved | Items you saved for later out of your cart (last 200).
[**ListBuyerCases**](BuyerApi.md#ListBuyerCases) | **Get** /v1/buyer/cases | Resolution center — every case you opened (returns, not received, problems).
[**ListBuyerCashback**](BuyerApi.md#ListBuyerCashback) | **Get** /v1/buyer/cashback | Your Scout cashback — pending, confirmed, paid.
[**ListBuyerCatalogListingVideos**](BuyerApi.md#ListBuyerCatalogListingVideos) | **Get** /v1/buyer/catalog/listings/{slug}/videos | Videos the seller recorded for this item on request, newest first, with the question each one answers.
[**ListBuyerClubs**](BuyerApi.md#ListBuyerClubs) | **Get** /v1/buyer/clubs | Sellers’ loyalty clubs you belong to: your spend with each, your tier and its discount, and the next tier.
[**ListBuyerCollectionSets**](BuyerApi.md#ListBuyerCollectionSets) | **Get** /v1/buyer/collection/sets | Sets you are collecting — a catalogue brand + model you own at least two pieces of — with owned / total and percent complete, most complete first; and whether you show on set leaderboards.
[**ListBuyerConsolidations**](BuyerApi.md#ListBuyerConsolidations) | **Get** /v1/buyer/consolidations | Your ship-together orders: parcels at the hub, and the one box out.
[**ListBuyerConversationGuests**](BuyerApi.md#ListBuyerConversationGuests) | **Get** /v1/buyer/conversations/{id}/guests | Guests in a chat.
[**ListBuyerFeedback**](BuyerApi.md#ListBuyerFeedback) | **Get** /v1/buyer/feedback | Feedback you have left sellers, newest first — with the ids you revise by.
[**ListBuyerFeedbackPending**](BuyerApi.md#ListBuyerFeedbackPending) | **Get** /v1/buyer/feedback/pending | Orders from the last 60 days that are on their way or arrived and still need your feedback.
[**ListBuyerFitRefs**](BuyerApi.md#ListBuyerFitRefs) | **Get** /v1/buyer/fit/refs | Your saved reference garments for fit comparison.
[**ListBuyerFollowingSellers**](BuyerApi.md#ListBuyerFollowingSellers) | **Get** /v1/buyer/following/sellers | Sellers you follow, with how many new listings each has since you last looked.
[**ListBuyerGiftCardCredit**](BuyerApi.md#ListBuyerGiftCardCredit) | **Get** /v1/buyer/gift-cards/credit | Gift card credit on your account, and when it expires.
[**ListBuyerGiftCards**](BuyerApi.md#ListBuyerGiftCards) | **Get** /v1/buyer/gift-cards | Gift cards you bought, newest first.
[**ListBuyerGiftFinder**](BuyerApi.md#ListBuyerGiftFinder) | **Get** /v1/buyer/gift-finder | Live listings for a recipient, interest and budget.
[**ListBuyerGiftSent**](BuyerApi.md#ListBuyerGiftSent) | **Get** /v1/buyer/gifts/sent | Items you sent as gifts: whether the teaser went out, whether it arrived and was opened, and any thank-you. Send one with &#x60;gift&#x60; on checkout.
[**ListBuyerHotspots**](BuyerApi.md#ListBuyerHotspots) | **Get** /v1/buyer/hotspots | Shop-the-look tags on a story or lookbook photo: each tagged listing and where it sits on the image (x, y as fractions).
[**ListBuyerHunts**](BuyerApi.md#ListBuyerHunts) | **Get** /v1/buyer/hunts | Your buyer-agent hunts and what they did.
[**ListBuyerInstantOffers**](BuyerApi.md#ListBuyerInstantOffers) | **Get** /v1/buyer/instant-offers | Your instant offers: live quotes, accepted ones awaiting a label, approved ones (with the prepaid label) and paid ones.
[**ListBuyerLayaways**](BuyerApi.md#ListBuyerLayaways) | **Get** /v1/buyer/layaways | Your layaways: price, paid so far, what is left, and the installment schedule. Put an item on layaway at checkout with layawayItemIds.
[**ListBuyerListingQuickQuestions**](BuyerApi.md#ListBuyerListingQuickQuestions) | **Get** /v1/buyer/listings/{slug}/quick-questions | One-tap questions this listing supports; send one as &#x60;quick&#x60; on a message.
[**ListBuyerListingSellerReviews**](BuyerApi.md#ListBuyerListingSellerReviews) | **Get** /v1/buyer/listings/{slug}/seller-reviews | The seller’s public reviews: summary, text, photos, replies.
[**ListBuyerLiveShipments**](BuyerApi.md#ListBuyerLiveShipments) | **Get** /v1/buyer/live-shipments | Your one-box-per-show shipments: every order won in a show ships together, and postage is charged once (the rest refunded when it ships).
[**ListBuyerLiveStreamLineup**](BuyerApi.md#ListBuyerLiveStreamLineup) | **Get** /v1/buyer/live-streams/{id}/lineup | What a live show will sell, in order: each item, its price, and whether it is still available.
[**ListBuyerMarketBids**](BuyerApi.md#ListBuyerMarketBids) | **Get** /v1/buyer/market/bids | Your resting market bids.
[**ListBuyerMarketWatches**](BuyerApi.md#ListBuyerMarketWatches) | **Get** /v1/buyer/market/watches | Market variants you watch, with the price you are waiting for.
[**ListBuyerOffers**](BuyerApi.md#ListBuyerOffers) | **Get** /v1/buyer/offers | Offers you have made, newest first (last 100).
[**ListBuyerOrders**](BuyerApi.md#ListBuyerOrders) | **Get** /v1/buyer/orders | What you have bought on Crossly, newest first.
[**ListBuyerPackOpenings**](BuyerApi.md#ListBuyerPackOpenings) | **Get** /v1/buyer/pack-openings | Your pulls: what you got, its value, sell-back status, and the fairness proof for each (the server seed once rotated).
[**ListBuyerPacks**](BuyerApi.md#ListBuyerPacks) | **Get** /v1/buyer/packs | Packs you can open: price, every tier with its published odds, what is left in it and its value range, and the expected value of one opening.
[**ListBuyerParties**](BuyerApi.md#ListBuyerParties) | **Get** /v1/buyer/parties | Themed shopping parties, live and upcoming.
[**ListBuyerPartyListings**](BuyerApi.md#ListBuyerPartyListings) | **Get** /v1/buyer/parties/{slug}/listings | Listings shared to a party, featured first.
[**ListBuyerPickupPoints**](BuyerApi.md#ListBuyerPickupPoints) | **Get** /v1/buyer/pickup-points | Carrier pickup points (UPS Access Point / FedEx Hold at Location) near an address; pass one as &#x60;pickupPoint&#x60; to POST /v1/buyer/checkout.
[**ListBuyerProtectionPlans**](BuyerApi.md#ListBuyerProtectionPlans) | **Get** /v1/buyer/protection-plans | Your protection plans (bought at checkout with &#x60;protection&#x60;): status, term, coverage, contract id, and any refund.
[**ListBuyerRegistries**](BuyerApi.md#ListBuyerRegistries) | **Get** /v1/buyer/registries | Your gift registries.
[**ListBuyerRegistryGifts**](BuyerApi.md#ListBuyerRegistryGifts) | **Get** /v1/buyer/registries/{id}/gifts | Your thank-you list: who bought what off the registry, and whether you have thanked them.
[**ListBuyerRegistrySearch**](BuyerApi.md#ListBuyerRegistrySearch) | **Get** /v1/buyer/registries/search | Find a registry by its name or the registrant’s name. Current and recent registries only.
[**ListBuyerRentals**](BuyerApi.md#ListBuyerRentals) | **Get** /v1/buyer/rentals | Your rentals: dates, status (reserved, out, returned, claimed, closed, cancelled), the return label, and what happened to the deposit.
[**ListBuyerSizeMemoryPrompts**](BuyerApi.md#ListBuyerSizeMemoryPrompts) | **Get** /v1/buyer/size-memory/prompts | Delivered orders with a size that you haven’t said how they fit (the last 60 days).
[**ListBuyerSizeMemorySizes**](BuyerApi.md#ListBuyerSizeMemorySizes) | **Get** /v1/buyer/size-memory/sizes | The sizes learned from what you bought and kept, by kind of item (shoes, tops, bottoms, dresses, outerwear) and department.
[**ListBuyerSplits**](BuyerApi.md#ListBuyerSplits) | **Get** /v1/buyer/splits | Purchases you split, with each share’s status.
[**ListBuyerSubscriptions**](BuyerApi.md#ListBuyerSubscriptions) | **Get** /v1/buyer/subscriptions | Your subscribe &amp; save standing orders.
[**ListBuyerTogether**](BuyerApi.md#ListBuyerTogether) | **Get** /v1/buyer/together | Shop-together rooms you are in.
[**ListBuyerTotalRequests**](BuyerApi.md#ListBuyerTotalRequests) | **Get** /v1/buyer/total-requests | Your combined-shipping total requests, newest per seller.
[**ListBuyerTryOns**](BuyerApi.md#ListBuyerTryOns) | **Get** /v1/buyer/try-ons | Your AI try-ons: the rendered images (labelled AI-generated wherever shown).
[**ListBuyerVault**](BuyerApi.md#ListBuyerVault) | **Get** /v1/buyer/vault | What you own in the Crossly vault: stored (authenticated, on a shelf), listed for sale, or on its way home — with the price to ship one home.
[**ListBuyerVaultMarketHoldings**](BuyerApi.md#ListBuyerVaultMarketHoldings) | **Get** /v1/buyer/vault-markets/holdings | Each unit you hold in the vault at its market price, with the mark’s quality (A or B reliable, C indicative, none), cost, P&amp;L and where it stands (stored, listed, lent, pledged, at the grader); &#x60;meta&#x60; carries the totals and the units you have out on loan.
[**ListBuyerVaultMarketIndexes**](BuyerApi.md#ListBuyerVaultMarketIndexes) | **Get** /v1/buyer/vault-markets/indexes | Crossly indexes over vault book keys (category, brand, curated, composite, segment, and your own), each with its level now and change since the last daily close. Index levels are valued only from the book’s marks.
[**ListBuyerVaultMarketInstantExit**](BuyerApi.md#ListBuyerVaultMarketInstantExit) | **Get** /v1/buyer/vault-markets/instant-exit | Instant Exit quotes for what you hold, one per book key: the best resting bid (yours left out), its share of the mark, the bid levels, units held and which are free to sell, and why a key cannot be sold instantly. A quote, not an order.
[**ListBuyerVaultMarketWatchlist**](BuyerApi.md#ListBuyerVaultMarketWatchlist) | **Get** /v1/buyer/vault-markets/watchlist | The book keys you watch, with mark and its quality, 24 h and 7 d change, and the last 30 daily marks.
[**ListBuyerVideoRequests**](BuyerApi.md#ListBuyerVideoRequests) | **Get** /v1/buyer/video-requests | Your video requests and their answers.
[**ListBuyerWishlistItems**](BuyerApi.md#ListBuyerWishlistItems) | **Get** /v1/buyer/wishlists/{id}/items | What is on one wishlist.
[**ListBuyerWishlistMembers**](BuyerApi.md#ListBuyerWishlistMembers) | **Get** /v1/buyer/wishlists/{id}/members | Co-editors of a shared wishlist.
[**ListBuyerWishlists**](BuyerApi.md#ListBuyerWishlists) | **Get** /v1/buyer/wishlists | Your wishlists.
[**UpdateBuyerBoard**](BuyerApi.md#UpdateBuyerBoard) | **Patch** /v1/buyer/boards/{id} | Rename a board, change its description, or make it public / private.
[**UpdateBuyerBoardLayout**](BuyerApi.md#UpdateBuyerBoardLayout) | **Put** /v1/buyer/boards/{id}/layout | Replace a board’s layout: items by platform listing id or slug, with x, y, w (0–1) and z. At most 40.
[**UpdateBuyerCollectionSetLeaderboard**](BuyerApi.md#UpdateBuyerCollectionSetLeaderboard) | **Put** /v1/buyer/collection/sets/leaderboard | Show or hide yourself (by your public handle) on set leaderboards.
[**UpdateBuyerFeedback**](BuyerApi.md#UpdateBuyerFeedback) | **Patch** /v1/buyer/feedback/{id} | Revise feedback you left, within 60 days.
[**UpdateBuyerHunt**](BuyerApi.md#UpdateBuyerHunt) | **Patch** /v1/buyer/hunts/{id} | Pause, resume or stop a hunt.
[**UpdateBuyerPurchaseStory**](BuyerApi.md#UpdateBuyerPurchaseStory) | **Put** /v1/buyer/purchases/{id}/story | Write your owner story (280 characters); it travels with the item if you resell it.
[**UpdateBuyerRegistry**](BuyerApi.md#UpdateBuyerRegistry) | **Patch** /v1/buyer/registries/{id} | Edit a registry: title, kind, date, message, ship-to address.
[**UpdateBuyerRegistryItem**](BuyerApi.md#UpdateBuyerRegistryItem) | **Put** /v1/buyer/registries/{id}/items/{platformListingId} | How many of an item you want (1–50).
[**UpdateBuyerSubscription**](BuyerApi.md#UpdateBuyerSubscription) | **Patch** /v1/buyer/subscriptions/{id} | Pause, resume, skip the next delivery, change the interval, or cancel.
[**UpdateBuyerWishlist**](BuyerApi.md#UpdateBuyerWishlist) | **Patch** /v1/buyer/wishlists/{id} | Rename a wishlist, or make it public / private.



## CreateBuyerActivity

> CreateBuyerActivityResponse CreateBuyerActivity(ctx).CreateBuyerActivityRequest(createBuyerActivityRequest).Execute()

Report an item your user is looking at, and get our answer.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBuyerActivityRequest := *openapiclient.NewCreateBuyerActivityRequest("Namespace_example", "Value_example") // CreateBuyerActivityRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerActivity(context.Background()).CreateBuyerActivityRequest(createBuyerActivityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerActivity`: CreateBuyerActivityResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuyerActivityRequest** | **CreateBuyerActivityRequest** |  | 

### Return type

**CreateBuyerActivityResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerAuctionBid

> CreateBuyerAuctionBidResponse CreateBuyerAuctionBid(ctx, slug).CreateBuyerAuctionBidRequest(createBuyerAuctionBidRequest).Execute()

Place a bid, optionally with a proxy maximum. A WINNING BID IS CHARGED AUTOMATICALLY.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    slug := "slug_example" // string | The auction listing slug.
    createBuyerAuctionBidRequest := *openapiclient.NewCreateBuyerAuctionBidRequest(int32(123)) // CreateBuyerAuctionBidRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerAuctionBid(context.Background(), slug).CreateBuyerAuctionBidRequest(createBuyerAuctionBidRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerAuctionBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerAuctionBid`: CreateBuyerAuctionBidResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerAuctionBid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | The auction listing slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerAuctionBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBuyerAuctionBidRequest** | **CreateBuyerAuctionBidRequest** |  | 

### Return type

**CreateBuyerAuctionBidResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerAuctionPay

> CreateBuyerAuctionPayResponse CreateBuyerAuctionPay(ctx, slug).Execute()

Pay for an auction you won whose automatic charge did not go through.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    slug := "slug_example" // string | The auction listing slug.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerAuctionPay(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerAuctionPay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerAuctionPay`: CreateBuyerAuctionPayResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerAuctionPay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | The auction listing slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerAuctionPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerAuctionPayResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerAuctionRetract

> CreateBuyerAuctionRetractResponse CreateBuyerAuctionRetract(ctx, slug).CreateBuyerAuctionRetractRequest(createBuyerAuctionRetractRequest).Execute()

Retract your bids under eBay’s rules — a typo, a changed listing, or an unreachable seller.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    slug := "slug_example" // string | The auction listing slug.
    createBuyerAuctionRetractRequest := *openapiclient.NewCreateBuyerAuctionRetractRequest("Reason_example") // CreateBuyerAuctionRetractRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerAuctionRetract(context.Background(), slug).CreateBuyerAuctionRetractRequest(createBuyerAuctionRetractRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerAuctionRetract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerAuctionRetract`: CreateBuyerAuctionRetractResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerAuctionRetract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | The auction listing slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerAuctionRetractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBuyerAuctionRetractRequest** | **CreateBuyerAuctionRetractRequest** |  | 

### Return type

**CreateBuyerAuctionRetractResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerBoard

> CreateBuyerBoardResponse CreateBuyerBoard(ctx).Execute()

Create an outfit board.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerBoard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerBoard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerBoard`: CreateBuyerBoardResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerBoard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerBoardRequest struct via the builder pattern


### Return type

**CreateBuyerBoardResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerBoardItem

> BuyerResponse CreateBuyerBoardItem(ctx, id).Execute()

Add a listing (by slug) to a board, at the next free spot.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerBoardItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerBoardItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerBoardItem`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerBoardItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerBoardItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerBreakSpotHold

> CreateBuyerBreakSpotHoldResponse CreateBuyerBreakSpotHold(ctx, id, spotId).Execute()

Hold a spot in a pick-your-spot break, then buy it through checkout before the hold lapses.

### Example

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
    spotId := "spotId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerBreakSpotHold(context.Background(), id, spotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerBreakSpotHold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerBreakSpotHold`: CreateBuyerBreakSpotHoldResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerBreakSpotHold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**spotId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerBreakSpotHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateBuyerBreakSpotHoldResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerBuybackRequest

> BuyerResponse CreateBuyerBuybackRequest(ctx, id).Execute()

Ask Crossly to buy an item back under its buyback guarantee.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerBuybackRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerBuybackRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerBuybackRequest`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerBuybackRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerBuybackRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCartItem

> CreateBuyerCartItemResponse CreateBuyerCartItem(ctx).CreateBuyerCartItemRequest(createBuyerCartItemRequest).Execute()

Add a listing to your cart.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBuyerCartItemRequest := *openapiclient.NewCreateBuyerCartItemRequest("ListingSlug_example") // CreateBuyerCartItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerCartItem(context.Background()).CreateBuyerCartItemRequest(createBuyerCartItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCartItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCartItem`: CreateBuyerCartItemResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCartItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCartItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuyerCartItemRequest** | **CreateBuyerCartItemRequest** |  | 

### Return type

**CreateBuyerCartItemResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCartItemSaveForLater

> BuyerCartResponse CreateBuyerCartItemSaveForLater(ctx, id).Execute()

Move a cart line to saved-for-later.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerCartItemSaveForLater(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCartItemSaveForLater``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCartItemSaveForLater`: BuyerCartResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCartItemSaveForLater`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCartItemSaveForLaterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerCartResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCartOptimizer

> CreateBuyerCartOptimizerResponse CreateBuyerCartOptimizer(ctx).Execute()

Cart optimizer: send a shopping list (one item per line, \"2 Charizard Base Set\") and get the cheapest delivered basket across sellers — one pick per line, items + each seller’s shipping with free-shipping thresholds applied. fewestPackages prefers fewer sellers. Nothing is added to the cart.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerCartOptimizer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCartOptimizer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCartOptimizer`: CreateBuyerCartOptimizerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCartOptimizer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCartOptimizerRequest struct via the builder pattern


### Return type

**CreateBuyerCartOptimizerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCartQuote

> CreateBuyerCartQuoteResponse CreateBuyerCartQuote(ctx).Execute()

Price the cart, delivered — item, shipping, tax, total.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerCartQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCartQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCartQuote`: CreateBuyerCartQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCartQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCartQuoteRequest struct via the builder pattern


### Return type

**CreateBuyerCartQuoteResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCartSavedMoveToCart

> BuyerCartResponse CreateBuyerCartSavedMoveToCart(ctx, id).Execute()

Move a saved item back into your cart, at today’s price.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerCartSavedMoveToCart(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCartSavedMoveToCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCartSavedMoveToCart`: BuyerCartResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCartSavedMoveToCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCartSavedMoveToCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerCartResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCaseCounter

> BuyerResponse CreateBuyerCaseCounter(ctx, id).CreateBuyerCaseCounterRequest(createBuyerCaseCounterRequest).Execute()

Propose your own refund figure on a case.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.
    createBuyerCaseCounterRequest := *openapiclient.NewCreateBuyerCaseCounterRequest(int32(123)) // CreateBuyerCaseCounterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerCaseCounter(context.Background(), id).CreateBuyerCaseCounterRequest(createBuyerCaseCounterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCaseCounter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCaseCounter`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCaseCounter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCaseCounterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBuyerCaseCounterRequest** | **CreateBuyerCaseCounterRequest** |  | 

### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCaseEscalate

> BuyerResponse CreateBuyerCaseEscalate(ctx, id).Execute()

Ask Crossly to step in on a case.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerCaseEscalate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCaseEscalate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCaseEscalate`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCaseEscalate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCaseEscalateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCaseMessage

> BuyerResponse CreateBuyerCaseMessage(ctx, id).CreateBuyerCaseMessageRequest(createBuyerCaseMessageRequest).Execute()

Post a message on a case thread.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.
    createBuyerCaseMessageRequest := *openapiclient.NewCreateBuyerCaseMessageRequest("Body_example") // CreateBuyerCaseMessageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerCaseMessage(context.Background(), id).CreateBuyerCaseMessageRequest(createBuyerCaseMessageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCaseMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCaseMessage`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCaseMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCaseMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBuyerCaseMessageRequest** | **CreateBuyerCaseMessageRequest** |  | 

### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCaseRespond

> BuyerResponse CreateBuyerCaseRespond(ctx, id).CreateBuyerCaseRespondRequest(createBuyerCaseRespondRequest).Execute()

Accept or reject the seller’s partial-refund offer on a case.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.
    createBuyerCaseRespondRequest := *openapiclient.NewCreateBuyerCaseRespondRequest("Decision_example") // CreateBuyerCaseRespondRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerCaseRespond(context.Background(), id).CreateBuyerCaseRespondRequest(createBuyerCaseRespondRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCaseRespond``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCaseRespond`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCaseRespond`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCaseRespondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBuyerCaseRespondRequest** | **CreateBuyerCaseRespondRequest** |  | 

### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCaseWithdraw

> BuyerResponse CreateBuyerCaseWithdraw(ctx, id).Execute()

Withdraw a case you opened.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerCaseWithdraw(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCaseWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCaseWithdraw`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCaseWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCaseWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCatalogListingRental

> CreateBuyerCatalogListingRentalResponse CreateBuyerCatalogListingRental(ctx, slug).Execute()

Book a rental. Returns a Stripe PaymentIntent client secret for rent + shipping both ways + deposit; the dates are held 30 minutes for payment. Crossly buys both labels; the deposit comes back after the item does.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerCatalogListingRental(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCatalogListingRental``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCatalogListingRental`: CreateBuyerCatalogListingRentalResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCatalogListingRental`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCatalogListingRentalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerCatalogListingRentalResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCatalogListingRentalQuote

> CreateBuyerCatalogListingRentalQuoteResponse CreateBuyerCatalogListingRentalQuote(ctx, slug).Execute()

Price a rental: rent (the cheaper of by-the-day and by-the-week), both shipping legs once a ship-to is given, and the deposit. `problem` says why the dates can’t be booked.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerCatalogListingRentalQuote(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCatalogListingRentalQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCatalogListingRentalQuote`: CreateBuyerCatalogListingRentalQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCatalogListingRentalQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCatalogListingRentalQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerCatalogListingRentalQuoteResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerCatalogListingVideoRequest

> CreateBuyerCatalogListingVideoRequestResponse CreateBuyerCatalogListingVideoRequest(ctx, slug).Execute()

Ask the seller for a video of an item (one open request per item). The seller is notified; their answer appears on the listing for every buyer.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerCatalogListingVideoRequest(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerCatalogListingVideoRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCatalogListingVideoRequest`: CreateBuyerCatalogListingVideoRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerCatalogListingVideoRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCatalogListingVideoRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerCatalogListingVideoRequestResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerConversationInvite

> CreateBuyerWishlistInviteResponse CreateBuyerConversationInvite(ctx, id).Execute()

Invite a friend into a chat with a seller: a single-use link (7 days) that makes whoever opens it a guest who can read and post. Only the thread’s buyer can invite; at most 5 guests.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerConversationInvite(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerConversationInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerConversationInvite`: CreateBuyerWishlistInviteResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerConversationInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerConversationInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerWishlistInviteResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerConversationJoin

> CreateBuyerConversationJoinResponse CreateBuyerConversationJoin(ctx, token).Execute()

Join a chat you were invited into, as a guest.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerConversationJoin(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerConversationJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerConversationJoin`: CreateBuyerConversationJoinResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerConversationJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerConversationJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerConversationJoinResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerFeedback

> CreateBuyerFeedbackResponse CreateBuyerFeedback(ctx).CreateBuyerFeedbackRequest(createBuyerFeedbackRequest).Execute()

Leave feedback on up to 25 orders at once, with optional detailed seller ratings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBuyerFeedbackRequest := *openapiclient.NewCreateBuyerFeedbackRequest([]openapiclient.CreateBuyerFeedbackRequestReviews{*openapiclient.NewCreateBuyerFeedbackRequestReviews("PurchaseId_example", int32(123))}) // CreateBuyerFeedbackRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerFeedback(context.Background()).CreateBuyerFeedbackRequest(createBuyerFeedbackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerFeedback`: CreateBuyerFeedbackResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuyerFeedbackRequest** | **CreateBuyerFeedbackRequest** |  | 

### Return type

**CreateBuyerFeedbackResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerFitRef

> BuyerResponse2 CreateBuyerFitRef(ctx).Execute()

Save a reference garment (flat measurements, inches) to compare listings against.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerFitRef(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerFitRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerFitRef`: BuyerResponse2
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerFitRef`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerFitRefRequest struct via the builder pattern


### Return type

**BuyerResponse2**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerFollowingSellerSeen

> CreateBuyerFollowingSellerSeenResponse CreateBuyerFollowingSellerSeen(ctx).CreateBuyerFollowingSellerSeenRequest(createBuyerFollowingSellerSeenRequest).Execute()

Clear the \"new listings\" count on a followed seller, or on all of them.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBuyerFollowingSellerSeenRequest := *openapiclient.NewCreateBuyerFollowingSellerSeenRequest() // CreateBuyerFollowingSellerSeenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerFollowingSellerSeen(context.Background()).CreateBuyerFollowingSellerSeenRequest(createBuyerFollowingSellerSeenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerFollowingSellerSeen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerFollowingSellerSeen`: CreateBuyerFollowingSellerSeenResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerFollowingSellerSeen`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerFollowingSellerSeenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuyerFollowingSellerSeenRequest** | **CreateBuyerFollowingSellerSeenRequest** |  | 

### Return type

**CreateBuyerFollowingSellerSeenResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerGiftCard

> CreateBuyerGiftCardResponse CreateBuyerGiftCard(ctx).CreateBuyerGiftCardRequest(createBuyerGiftCardRequest).Execute()

Start a gift card purchase.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBuyerGiftCardRequest := *openapiclient.NewCreateBuyerGiftCardRequest(int32(123), "RecipientEmail_example") // CreateBuyerGiftCardRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerGiftCard(context.Background()).CreateBuyerGiftCardRequest(createBuyerGiftCardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerGiftCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerGiftCard`: CreateBuyerGiftCardResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerGiftCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerGiftCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuyerGiftCardRequest** | **CreateBuyerGiftCardRequest** |  | 

### Return type

**CreateBuyerGiftCardResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerGiftCardRedeem

> CreateBuyerGiftCardRedeemResponse CreateBuyerGiftCardRedeem(ctx).CreateBuyerGiftCardRedeemRequest(createBuyerGiftCardRedeemRequest).Execute()

Redeem a gift card code into your own balance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBuyerGiftCardRedeemRequest := *openapiclient.NewCreateBuyerGiftCardRedeemRequest("Code_example") // CreateBuyerGiftCardRedeemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerGiftCardRedeem(context.Background()).CreateBuyerGiftCardRedeemRequest(createBuyerGiftCardRedeemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerGiftCardRedeem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerGiftCardRedeem`: CreateBuyerGiftCardRedeemResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerGiftCardRedeem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerGiftCardRedeemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuyerGiftCardRedeemRequest** | **CreateBuyerGiftCardRedeemRequest** |  | 

### Return type

**CreateBuyerGiftCardRedeemResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerGroupBuyJoin

> CreateBuyerGroupBuyJoinResponse CreateBuyerGroupBuyJoin(ctx, id).Execute()

Join a group buy: a hold at the starting price, captured at the price the group reaches. Checked against this key’s spending controls; send an Idempotency-Key.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerGroupBuyJoin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerGroupBuyJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerGroupBuyJoin`: CreateBuyerGroupBuyJoinResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerGroupBuyJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerGroupBuyJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerGroupBuyJoinResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerGroupBuyLeave

> BuyerResponse CreateBuyerGroupBuyLeave(ctx, id).Execute()

Leave a group buy and release your hold.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerGroupBuyLeave(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerGroupBuyLeave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerGroupBuyLeave`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerGroupBuyLeave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerGroupBuyLeaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerHunt

> BuyerResponse2 CreateBuyerHunt(ctx).Execute()

Start a buyer-agent hunt — over a search (filters), or auto-buy on one listing (listingSlug: buy it once its price, sale applied, is within maxPriceCents). `offer` and `buy` modes also need buyer:checkout:write, and the budget is checked against this key’s spending controls.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerHunt(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerHunt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerHunt`: BuyerResponse2
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerHunt`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerHuntRequest struct via the builder pattern


### Return type

**BuyerResponse2**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerInstantOfferAccept

> BuyerResponse CreateBuyerInstantOfferAccept(ctx, id).Execute()

Accept Crossly’s instant offer on a collection item. Ops then send a prepaid label; the store credit is paid once the item arrives and checks out.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerInstantOfferAccept(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerInstantOfferAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerInstantOfferAccept`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerInstantOfferAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerInstantOfferAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerLayawayCancel

> BuyerLayawayResponse CreateBuyerLayawayCancel(ctx, id).Execute()

Cancel a layaway: everything paid comes back except the layaway fee (10% of the price, at most $15).

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerLayawayCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerLayawayCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerLayawayCancel`: BuyerLayawayResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerLayawayCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerLayawayCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerLayawayResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerLayawayPay

> BuyerLayawayResponse CreateBuyerLayawayPay(ctx, id).Execute()

Pay some (at least $1) or all of what is left now, on the saved card. Paid in full, it ships.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerLayawayPay(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerLayawayPay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerLayawayPay`: BuyerLayawayResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerLayawayPay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerLayawayPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerLayawayResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerListingAudioSummary

> GetBuyerListingAudioSummaryResponse CreateBuyerListingAudioSummary(ctx, slug).Execute()

Record the spoken summary of a listing (cached per listing version). 503 naming TTS_API_KEY when no speech provider is configured.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerListingAudioSummary(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerListingAudioSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerListingAudioSummary`: GetBuyerListingAudioSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerListingAudioSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerListingAudioSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerListingAudioSummaryResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerListingSplit

> CreateBuyerListingSplitResponse CreateBuyerListingSplit(ctx, slug).Execute()

Split a purchase: one pay link per person; nobody is charged until all have paid.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerListingSplit(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerListingSplit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerListingSplit`: CreateBuyerListingSplitResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerListingSplit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerListingSplitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerListingSplitResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerListingSubscribe

> BuyerResponse2 CreateBuyerListingSubscribe(ctx, slug).Execute()

Subscribe & save: a standing order charged to your saved card every interval. Checked against this key’s spending controls; send an Idempotency-Key.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerListingSubscribe(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerListingSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerListingSubscribe`: BuyerResponse2
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerListingSubscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerListingSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse2**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerListingTryOn

> CreateBuyerListingTryOnResponse CreateBuyerListingTryOn(ctx, slug).Execute()

Render an AI try-on of this listing on the buyer’s own photo (base64 `photo`, or `reuseLastPhoto`). `confirmSelfAdult` must be true: the photo is of the buyer and they are 18+. Wearables only; capped per day; 503 naming TRYON_API_KEY when no image model is configured.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerListingTryOn(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerListingTryOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerListingTryOn`: CreateBuyerListingTryOnResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerListingTryOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerListingTryOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerListingTryOnResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerMarketBid

> CreateBuyerMarketBidResponse CreateBuyerMarketBid(ctx).Execute()

Place a fixed-price market bid. SPENDS MONEY: it fills against the lowest ask at or under your price and charges your saved card (plus the fast-ship premium when fastShip is set and it fills from vault stock). deliverTo vault keeps it in your Crossly vault; gradeWith sends a raw card to a Grading run first. Checked against this key’s spending controls; send an Idempotency-Key.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerMarketBid(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerMarketBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerMarketBid`: CreateBuyerMarketBidResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerMarketBid`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerMarketBidRequest struct via the builder pattern


### Return type

**CreateBuyerMarketBidResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerMarketWatch

> BuyerResponse CreateBuyerMarketWatch(ctx).Execute()

Watch a market variant: alert at or below a price, or on any drop.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerMarketWatch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerMarketWatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerMarketWatch`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerMarketWatch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerMarketWatchRequest struct via the builder pattern


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerOffer

> CreateBuyerOfferResponse CreateBuyerOffer(ctx).CreateBuyerOfferRequest(createBuyerOfferRequest).Execute()

Offer a price on a listing.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBuyerOfferRequest := *openapiclient.NewCreateBuyerOfferRequest("ListingSlug_example", int32(123)) // CreateBuyerOfferRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerOffer(context.Background()).CreateBuyerOfferRequest(createBuyerOfferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerOffer`: CreateBuyerOfferResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerOffer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuyerOfferRequest** | **CreateBuyerOfferRequest** |  | 

### Return type

**CreateBuyerOfferResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerOfferRespond

> CreateBuyerOfferRespondResponse CreateBuyerOfferRespond(ctx, id).CreateBuyerOfferRespondRequest(createBuyerOfferRespondRequest).Execute()

Answer a seller’s counteroffer — accept, decline, or counter back.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The offer id.
    createBuyerOfferRespondRequest := *openapiclient.NewCreateBuyerOfferRespondRequest("Action_example") // CreateBuyerOfferRespondRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerOfferRespond(context.Background(), id).CreateBuyerOfferRespondRequest(createBuyerOfferRespondRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerOfferRespond``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerOfferRespond`: CreateBuyerOfferRespondResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerOfferRespond`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The offer id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerOfferRespondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBuyerOfferRespondRequest** | **CreateBuyerOfferRespondRequest** |  | 

### Return type

**CreateBuyerOfferRespondResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerOrderCancel

> CreateBuyerOrderCancelResponse CreateBuyerOrderCancel(ctx, id).Execute()

Cancel an order before it ships — refunds it in full.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerOrderCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerOrderCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerOrderCancel`: CreateBuyerOrderCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerOrderCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerOrderCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerOrderCancelResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerOrderInr

> CreateBuyerOrderInrResponse CreateBuyerOrderInr(ctx, id).CreateBuyerOrderInrRequest(createBuyerOrderInrRequest).Execute()

Open an item-not-received case.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.
    createBuyerOrderInrRequest := *openapiclient.NewCreateBuyerOrderInrRequest("Body_example") // CreateBuyerOrderInrRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerOrderInr(context.Background(), id).CreateBuyerOrderInrRequest(createBuyerOrderInrRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerOrderInr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerOrderInr`: CreateBuyerOrderInrResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerOrderInr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerOrderInrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBuyerOrderInrRequest** | **CreateBuyerOrderInrRequest** |  | 

### Return type

**CreateBuyerOrderInrResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerOrderReceived

> CreateBuyerOrderReceivedResponse CreateBuyerOrderReceived(ctx, id).Execute()

Confirm you received an order.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerOrderReceived(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerOrderReceived``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerOrderReceived`: CreateBuyerOrderReceivedResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerOrderReceived`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerOrderReceivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerOrderReceivedResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerOrderTip

> CreateBuyerOrderTipResponse CreateBuyerOrderTip(ctx, id).Execute()

Tip the seller on one of your shipped orders, charged to your saved card. One tip per order; $1–$500; checked against this key’s spending controls.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerOrderTip(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerOrderTip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerOrderTip`: CreateBuyerOrderTipResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerOrderTip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerOrderTipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerOrderTipResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerPackFairnessRotate

> GetBuyerPackFairnessResponse CreateBuyerPackFairnessRotate(ctx).Execute()

Reveal your current server seed and start a new one (optionally with your own client seed). Refused while an opening is awaiting payment.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerPackFairnessRotate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerPackFairnessRotate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerPackFairnessRotate`: GetBuyerPackFairnessResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerPackFairnessRotate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerPackFairnessRotateRequest struct via the builder pattern


### Return type

**GetBuyerPackFairnessResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerPackOpen

> CreateBuyerPackOpenResponse CreateBuyerPackOpen(ctx, slug).Execute()

Open a pack. SPENDS MONEY: charges the pack price to your saved card, then pulls one real item by the published odds (provably fair: see /buyer/packs/fairness) into your Crossly vault. confirmAdult must be true. Checked against this key’s spending controls; send an Idempotency-Key.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerPackOpen(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerPackOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerPackOpen`: CreateBuyerPackOpenResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerPackOpen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerPackOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerPackOpenResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerPackOpeningSellBack

> CreateBuyerPackOpeningSellBackResponse CreateBuyerPackOpeningSellBack(ctx, id).Execute()

Sell a pull straight back for Crossly credit at the pack’s buyback rate, within its window, while it is still in your vault.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerPackOpeningSellBack(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerPackOpeningSellBack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerPackOpeningSellBack`: CreateBuyerPackOpeningSellBackResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerPackOpeningSellBack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerPackOpeningSellBackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerPackOpeningSellBackResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerPassportClaim

> GetBuyerPassportResponse CreateBuyerPassportClaim(ctx, code).Execute()

Claim a passport as the item’s current owner (the buyer of its latest sale) — shown as “in its owner’s hands” until it sells again.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "code_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerPassportClaim(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerPassportClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerPassportClaim`: GetBuyerPassportResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerPassportClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerPassportClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerPassportResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerProtectionPlanCancel

> CreateBuyerProtectionPlanCancelResponse CreateBuyerProtectionPlanCancel(ctx, id).Execute()

Cancel a protection plan. The insurer decides what is refundable (pro-rata after the free-look window); that amount goes back to the card.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerProtectionPlanCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerProtectionPlanCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerProtectionPlanCancel`: CreateBuyerProtectionPlanCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerProtectionPlanCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerProtectionPlanCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerProtectionPlanCancelResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerPurchaseDownload

> CreateBuyerPurchaseDownloadResponse CreateBuyerPurchaseDownload(ctx, id).Execute()

A ten-minute link to a digital purchase’s file; each call uses one of its downloads.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerPurchaseDownload(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerPurchaseDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerPurchaseDownload`: CreateBuyerPurchaseDownloadResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerPurchaseDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerPurchaseDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerPurchaseDownloadResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerPurchaseFit

> BuyerResponse CreateBuyerPurchaseFit(ctx, id).Execute()

Say how a delivered purchase fit: true_to_size, small or large. It teaches your size memory and, pooled with other buyers, how the brand runs.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerPurchaseFit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerPurchaseFit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerPurchaseFit`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerPurchaseFit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerPurchaseFitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerPurchaseReturn

> BuyerResponse2 CreateBuyerPurchaseReturn(ctx, id).Execute()

Return a purchase under the seller’s policy, or send back a try-at-home item. With reason doesnt_fit, `fit` (small | large) says which way, and it feeds your size memory.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerPurchaseReturn(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerPurchaseReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerPurchaseReturn`: BuyerResponse2
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerPurchaseReturn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerPurchaseReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse2**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerRegistry

> CreateBuyerRegistryResponse CreateBuyerRegistry(ctx).Execute()

Start a gift registry (wedding, baby, birthday, housewarming, graduation, other). Items are added like any wishlist item; set shipAddressId to one of your saved addresses so guests can buy — they never see it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerRegistry(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerRegistry`: CreateBuyerRegistryResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerRegistry`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerRegistryRequest struct via the builder pattern


### Return type

**CreateBuyerRegistryResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerRegistryGiftThanked

> CreateBuyerRegistryGiftThankedResponse CreateBuyerRegistryGiftThanked(ctx, id, giftId).Execute()

Tick a gift as thanked (or untick it).

### Example

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
    giftId := "giftId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerRegistryGiftThanked(context.Background(), id, giftId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerRegistryGiftThanked``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerRegistryGiftThanked`: CreateBuyerRegistryGiftThankedResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerRegistryGiftThanked`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**giftId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerRegistryGiftThankedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateBuyerRegistryGiftThankedResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerRentalArrivalPhoto

> CreateBuyerRentalArrivalPhotoResponse CreateBuyerRentalArrivalPhoto(ctx, id).Execute()

Photograph a rental on arrival (URLs you uploaded): your side of the condition record.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerRentalArrivalPhoto(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerRentalArrivalPhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerRentalArrivalPhoto`: CreateBuyerRentalArrivalPhotoResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerRentalArrivalPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerRentalArrivalPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerRentalArrivalPhotoResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerRentalCancel

> BuyerResponse CreateBuyerRentalCancel(ctx, id).Execute()

Cancel a rental that has not shipped: everything is refunded.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerRentalCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerRentalCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerRentalCancel`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerRentalCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerRentalCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerRentalClaimResponse

> BuyerResponse CreateBuyerRentalClaimResponse(ctx, id).Execute()

Answer a deposit claim: accept (the claim is paid, the rest refunded) or dispute (Crossly decides).

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerRentalClaimResponse(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerRentalClaimResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerRentalClaimResponse`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerRentalClaimResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerRentalClaimResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerRentalIdentityStart

> CreateBuyerRentalIdentityStartResponse CreateBuyerRentalIdentityStart(ctx).Execute()

Start the ID check renting needs (Stripe Identity: photo ID + selfie). Open `url` (hosted) or use `clientSecret` with Stripe.js.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerRentalIdentityStart(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerRentalIdentityStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerRentalIdentityStart`: CreateBuyerRentalIdentityStartResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerRentalIdentityStart`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerRentalIdentityStartRequest struct via the builder pattern


### Return type

**CreateBuyerRentalIdentityStartResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerReturnCancel

> BuyerResponse CreateBuyerReturnCancel(ctx, id).Execute()

Cancel a return before it ships.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerReturnCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerReturnCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerReturnCancel`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerReturnCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerReturnCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerReturnTracking

> BuyerResponse CreateBuyerReturnTracking(ctx, id).Execute()

Add tracking to a return you shipped on your own label.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerReturnTracking(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerReturnTracking``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerReturnTracking`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerReturnTracking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerReturnTrackingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerSellerClubJoin

> CreateBuyerSellerClubJoinResponse CreateBuyerSellerClubJoin(ctx, sellerUserId).Execute()

Join a seller’s loyalty club (free). Your tier’s discount then applies to their items — the best of it and any other rate you have with them, never stacked.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sellerUserId := "sellerUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerSellerClubJoin(context.Background(), sellerUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerSellerClubJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerSellerClubJoin`: CreateBuyerSellerClubJoinResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerSellerClubJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerSellerClubJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerSellerClubJoinResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerSellerReport

> CreateBuyerSellerReportResponse CreateBuyerSellerReport(ctx, username).CreateBuyerSellerReportRequest(createBuyerSellerReportRequest).Execute()

Report a seller to Crossly’s trust team.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | The seller’s shop handle.
    createBuyerSellerReportRequest := *openapiclient.NewCreateBuyerSellerReportRequest("Reason_example") // CreateBuyerSellerReportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerSellerReport(context.Background(), username).CreateBuyerSellerReportRequest(createBuyerSellerReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerSellerReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerSellerReport`: CreateBuyerSellerReportResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerSellerReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The seller’s shop handle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerSellerReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBuyerSellerReportRequest** | **CreateBuyerSellerReportRequest** |  | 

### Return type

**CreateBuyerSellerReportResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerSplitCancel

> BuyerResponse CreateBuyerSplitCancel(ctx, id).Execute()

Cancel a split you organised; every held share is released.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerSplitCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerSplitCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerSplitCancel`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerSplitCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerSplitCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerTogether

> CreateBuyerTogetherResponse CreateBuyerTogether(ctx).Execute()

Start a shop-together room.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerTogether(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerTogether``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerTogether`: CreateBuyerTogetherResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerTogether`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerTogetherRequest struct via the builder pattern


### Return type

**CreateBuyerTogetherResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerTogetherItem

> BuyerResponse CreateBuyerTogetherItem(ctx, code).Execute()

Add a listing to a shop-together board.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "code_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerTogetherItem(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerTogetherItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerTogetherItem`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerTogetherItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerTogetherItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerTogetherItemVote

> BuyerResponse CreateBuyerTogetherItemVote(ctx, code, id).Execute()

Vote an item on a room’s board up (1), down (-1) or clear (0).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "code_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerTogetherItemVote(context.Background(), code, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerTogetherItemVote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerTogetherItemVote`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerTogetherItemVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerTogetherItemVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerTogetherJoin

> CreateBuyerTogetherResponse CreateBuyerTogetherJoin(ctx, code).Execute()

Join a shop-together room by its code.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "code_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerTogetherJoin(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerTogetherJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerTogetherJoin`: CreateBuyerTogetherResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerTogetherJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerTogetherJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerTogetherResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerTogetherLeave

> BuyerResponse CreateBuyerTogetherLeave(ctx, code).Execute()

Leave a shop-together room.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "code_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerTogetherLeave(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerTogetherLeave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerTogetherLeave`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerTogetherLeave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerTogetherLeaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerTogetherMessage

> BuyerResponse CreateBuyerTogetherMessage(ctx, code).Execute()

Say something in a shop-together room.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "code_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerTogetherMessage(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerTogetherMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerTogetherMessage`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerTogetherMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerTogetherMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerTotalRequest

> CreateBuyerTotalRequestResponse CreateBuyerTotalRequest(ctx).CreateBuyerTotalRequestRequest(createBuyerTotalRequestRequest).Execute()

Ask a seller for one combined shipping total on everything of theirs in your cart.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBuyerTotalRequestRequest := *openapiclient.NewCreateBuyerTotalRequestRequest("SellerUserId_example") // CreateBuyerTotalRequestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerTotalRequest(context.Background()).CreateBuyerTotalRequestRequest(createBuyerTotalRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerTotalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerTotalRequest`: CreateBuyerTotalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerTotalRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerTotalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuyerTotalRequestRequest** | **CreateBuyerTotalRequestRequest** |  | 

### Return type

**CreateBuyerTotalRequestResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerVaultList

> CreateBuyerVaultListResponse CreateBuyerVaultList(ctx, id).Execute()

Sell a vault item on the market at your price. It is already authenticated, so when it sells nothing ships in: the buyer keeps it vaulted or it ships straight to them.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerVaultList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerVaultList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerVaultList`: CreateBuyerVaultListResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerVaultList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerVaultListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerVaultListResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerVaultUnlist

> BuyerResponse CreateBuyerVaultUnlist(ctx, id).Execute()

Take a vault item off the market; it stays stored.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerVaultUnlist(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerVaultUnlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerVaultUnlist`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerVaultUnlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerVaultUnlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerVaultWithdraw

> CreateBuyerVaultWithdrawResponse CreateBuyerVaultWithdraw(ctx, id).Execute()

Ship a vault item home, standard or expedited. The shipping, any storage it owes and the early-withdrawal fee (within 90 days of arriving) are charged to your saved card now, all checked against this key’s spending controls; send an Idempotency-Key.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerVaultWithdraw(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerVaultWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerVaultWithdraw`: CreateBuyerVaultWithdrawResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerVaultWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerVaultWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerVaultWithdrawResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerWishlist

> CreateBuyerWishlistResponse CreateBuyerWishlist(ctx).CreateBuyerWishlistRequest(createBuyerWishlistRequest).Execute()

Create a wishlist.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBuyerWishlistRequest := *openapiclient.NewCreateBuyerWishlistRequest("Name_example") // CreateBuyerWishlistRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerWishlist(context.Background()).CreateBuyerWishlistRequest(createBuyerWishlistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerWishlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerWishlist`: CreateBuyerWishlistResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerWishlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerWishlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuyerWishlistRequest** | **CreateBuyerWishlistRequest** |  | 

### Return type

**CreateBuyerWishlistResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerWishlistInvite

> CreateBuyerWishlistInviteResponse CreateBuyerWishlistInvite(ctx, id).Execute()

Share a wishlist with a co-editor: a single-use link (14 days). Members add and remove items; only the owner renames, deletes or removes members. At most 20.

### Example

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
    resp, r, err := apiClient.BuyerApi.CreateBuyerWishlistInvite(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerWishlistInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerWishlistInvite`: CreateBuyerWishlistInviteResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerWishlistInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerWishlistInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerWishlistInviteResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerWishlistItem

> BuyerResponse2 CreateBuyerWishlistItem(ctx, id).CreateBuyerWishlistItemRequest(createBuyerWishlistItemRequest).Execute()

Add a listing to a wishlist.

### Example

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
    createBuyerWishlistItemRequest := *openapiclient.NewCreateBuyerWishlistItemRequest("ListingSlug_example") // CreateBuyerWishlistItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerWishlistItem(context.Background(), id).CreateBuyerWishlistItemRequest(createBuyerWishlistItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerWishlistItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerWishlistItem`: BuyerResponse2
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerWishlistItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerWishlistItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBuyerWishlistItemRequest** | **CreateBuyerWishlistItemRequest** |  | 

### Return type

**BuyerResponse2**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerWishlistJoin

> CreateBuyerWishlistJoinResponse CreateBuyerWishlistJoin(ctx, token).Execute()

Join a wishlist you were invited to edit.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerWishlistJoin(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerWishlistJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerWishlistJoin`: CreateBuyerWishlistJoinResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerWishlistJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerWishlistJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerWishlistJoinResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerWishlistShare

> CreateBuyerWishlistShareResponse CreateBuyerWishlistShare(ctx, id).CreateBuyerWishlistShareRequest(createBuyerWishlistShareRequest).Execute()

Share a wishlist — makes it public and returns its link.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.
    createBuyerWishlistShareRequest := *openapiclient.NewCreateBuyerWishlistShareRequest() // CreateBuyerWishlistShareRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateBuyerWishlistShare(context.Background(), id).CreateBuyerWishlistShareRequest(createBuyerWishlistShareRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateBuyerWishlistShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerWishlistShare`: CreateBuyerWishlistShareResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateBuyerWishlistShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerWishlistShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBuyerWishlistShareRequest** | **CreateBuyerWishlistShareRequest** |  | 

### Return type

**CreateBuyerWishlistShareResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerBoard

> BuyerResponse DeleteBuyerBoard(ctx, id).Execute()

Delete a board.

### Example

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
    resp, r, err := apiClient.BuyerApi.DeleteBuyerBoard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerBoard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerBoard`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerBoard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerBoardItem

> BuyerResponse DeleteBuyerBoardItem(ctx, id, platformListingId).Execute()

Take a listing off a board.

### Example

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
    platformListingId := "platformListingId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.DeleteBuyerBoardItem(context.Background(), id, platformListingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerBoardItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerBoardItem`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerBoardItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**platformListingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerBoardItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerBreakSpotHold

> BuyerResponse DeleteBuyerBreakSpotHold(ctx, id, spotId).Execute()

Release a break spot you hold.

### Example

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
    spotId := "spotId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.DeleteBuyerBreakSpotHold(context.Background(), id, spotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerBreakSpotHold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerBreakSpotHold`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerBreakSpotHold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**spotId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerBreakSpotHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerCartItem

> DeleteBuyerCartItem(ctx, id).Execute()

Remove a line from your cart.

### Example

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
    resp, r, err := apiClient.BuyerApi.DeleteBuyerCartItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerCartItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerCartItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerCartSaved

> BuyerCartResponse DeleteBuyerCartSaved(ctx, id).Execute()

Remove an item from saved-for-later. Returns the remaining list.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.DeleteBuyerCartSaved(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerCartSaved``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerCartSaved`: BuyerCartResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerCartSaved`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerCartSavedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerCartResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerConversationGuest

> BuyerResponse DeleteBuyerConversationGuest(ctx, id, userId).Execute()

Remove a guest (the thread’s buyer), or leave (a guest, with your own user id).

### Example

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
    resp, r, err := apiClient.BuyerApi.DeleteBuyerConversationGuest(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerConversationGuest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerConversationGuest`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerConversationGuest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerConversationGuestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerFitRef

> BuyerResponse DeleteBuyerFitRef(ctx, id).Execute()

Delete a reference garment.

### Example

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
    resp, r, err := apiClient.BuyerApi.DeleteBuyerFitRef(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerFitRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerFitRef`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerFitRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerFitRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerMarketBid

> BuyerResponse DeleteBuyerMarketBid(ctx, id).Execute()

Cancel a resting market bid.

### Example

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
    resp, r, err := apiClient.BuyerApi.DeleteBuyerMarketBid(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerMarketBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerMarketBid`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerMarketBid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerMarketBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerMarketWatch

> BuyerResponse DeleteBuyerMarketWatch(ctx, id).Execute()

Stop watching a market variant.

### Example

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
    resp, r, err := apiClient.BuyerApi.DeleteBuyerMarketWatch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerMarketWatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerMarketWatch`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerMarketWatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerMarketWatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerOffer

> BuyerResponse DeleteBuyerOffer(ctx, id).Execute()

Withdraw an offer you made that the seller has not answered.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The offer id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.DeleteBuyerOffer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerOffer`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The offer id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerRecentlyViewed

> BuyerResponse DeleteBuyerRecentlyViewed(ctx, slug).Execute()

Remove one listing from your recently viewed.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    slug := "slug_example" // string | The listing slug.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.DeleteBuyerRecentlyViewed(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerRecentlyViewed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerRecentlyViewed`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerRecentlyViewed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | The listing slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerRecentlyViewedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerRecentlyViewed2

> BuyerResponse DeleteBuyerRecentlyViewed2(ctx).Execute()

Clear your recently viewed list.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.DeleteBuyerRecentlyViewed2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerRecentlyViewed2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerRecentlyViewed2`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerRecentlyViewed2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerRecentlyViewed2Request struct via the builder pattern


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerRegistry

> BuyerResponse DeleteBuyerRegistry(ctx, id).Execute()

Delete a registry (and the wishlist behind it).

### Example

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
    resp, r, err := apiClient.BuyerApi.DeleteBuyerRegistry(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerRegistry`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerSellerClubJoin

> BuyerResponse DeleteBuyerSellerClubJoin(ctx, sellerUserId).Execute()

Leave a seller’s loyalty club.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sellerUserId := "sellerUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.DeleteBuyerSellerClubJoin(context.Background(), sellerUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerSellerClubJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerSellerClubJoin`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerSellerClubJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerSellerClubJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerTogetherItem

> BuyerResponse DeleteBuyerTogetherItem(ctx, code, id).Execute()

Take an item you added off a room’s board.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "code_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.DeleteBuyerTogetherItem(context.Background(), code, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerTogetherItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerTogetherItem`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerTogetherItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerTogetherItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerTotalRequest

> BuyerResponse DeleteBuyerTotalRequest(ctx, id).Execute()

Cancel a total request you made.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.DeleteBuyerTotalRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerTotalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerTotalRequest`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerTotalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerTotalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerTryOn

> DeleteBuyerTryOnResponse DeleteBuyerTryOn(ctx, id).Execute()

Delete a try-on and its images — or pass `all` to delete every photo and result.

### Example

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
    resp, r, err := apiClient.BuyerApi.DeleteBuyerTryOn(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerTryOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerTryOn`: DeleteBuyerTryOnResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerTryOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerTryOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteBuyerTryOnResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerWishlist

> BuyerResponse DeleteBuyerWishlist(ctx, id).Execute()

Delete a wishlist and everything on it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.DeleteBuyerWishlist(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerWishlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerWishlist`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerWishlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerWishlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerWishlistItem

> DeleteBuyerWishlistItemResponse DeleteBuyerWishlistItem(ctx, id, slug).Execute()

Take a listing off a wishlist.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The wishlist id.
    slug := "slug_example" // string | The listing slug.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.DeleteBuyerWishlistItem(context.Background(), id, slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerWishlistItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerWishlistItem`: DeleteBuyerWishlistItemResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerWishlistItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The wishlist id. | 
**slug** | **string** | The listing slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerWishlistItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**DeleteBuyerWishlistItemResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerWishlistMember

> BuyerResponse DeleteBuyerWishlistMember(ctx, id, userId).Execute()

Remove a co-editor (the owner), or leave a list shared with you (your own user id).

### Example

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
    resp, r, err := apiClient.BuyerApi.DeleteBuyerWishlistMember(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.DeleteBuyerWishlistMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuyerWishlistMember`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.DeleteBuyerWishlistMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerWishlistMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerAuctionMe

> GetBuyerAuctionMeResponse GetBuyerAuctionMe(ctx, slug).Execute()

Where you stand in one auction — your max, leading, won, must you pay.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    slug := "slug_example" // string | The auction listing slug.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerAuctionMe(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerAuctionMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerAuctionMe`: GetBuyerAuctionMeResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerAuctionMe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | The auction listing slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerAuctionMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerAuctionMeResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerBadge

> GetBuyerBadgeResponse GetBuyerBadge(ctx).Execute()

The counts behind the account badge — unread messages, offers waiting, orders to review.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerBadge(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerBadge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerBadge`: GetBuyerBadgeResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerBadge`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerBadgeRequest struct via the builder pattern


### Return type

**GetBuyerBadgeResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerBoard

> BuyerBoardResponse GetBuyerBoard(ctx, id).Execute()

One of your boards with its layout: each item’s x, y and width as fractions of the board, and stacking order.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerBoard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerBoard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerBoard`: BuyerBoardResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerBoard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerBoardResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerBoardShared

> BuyerBoardResponse GetBuyerBoardShared(ctx, token).Execute()

A shared outfit board, by its link token.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerBoardShared(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerBoardShared``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerBoardShared`: BuyerBoardResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerBoardShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerBoardSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerBoardResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerBreak

> GetBuyerBreakResponse GetBuyerBreak(ctx, id).Execute()

A card break’s board: spots, holders, and the seed after the draw.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerBreak(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerBreak``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerBreak`: GetBuyerBreakResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerBreak`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerBreakRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerBreakResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCase

> GetBuyerCaseResponse GetBuyerCase(ctx, id).Execute()

One case — its state, the message thread, and the return parcel if there is one.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerCase(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerCase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCase`: GetBuyerCaseResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerCaseResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCatalogListingMySize

> GetBuyerCatalogListingMySizeResponse GetBuyerCatalogListingMySize(ctx, slug).Execute()

This item against your size: the size you keep in this kind of item (and in this brand), whether this one is your size, smaller or larger, a warning if this brand and size ran small or large on you before, and how other buyers say the brand runs. null for an unsized item.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerCatalogListingMySize(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerCatalogListingMySize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCatalogListingMySize`: GetBuyerCatalogListingMySizeResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerCatalogListingMySize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCatalogListingMySizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerCatalogListingMySizeResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCatalogListingPrice

> GetBuyerCatalogListingPriceResponse GetBuyerCatalogListingPrice(ctx, slug).Execute()

The price today. When the seller has the same item live on another marketplace for less, this IS that price (`priceMatch` says where from), and `lowestAnywhere` lists the marketplaces it is guaranteed not to be cheaper on. `priceCheck` rates it against real sales of the same catalog item in the same condition over the last 180 days (great / good / fair / high, with the median and middle-half range in cents); null until there are 3+ sales from 2+ sellers.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerCatalogListingPrice(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerCatalogListingPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCatalogListingPrice`: GetBuyerCatalogListingPriceResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerCatalogListingPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCatalogListingPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerCatalogListingPriceResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCatalogListingRental

> GetBuyerCatalogListingRentalResponse GetBuyerCatalogListingRental(ctx, slug).Execute()

Whether a listing is for rent: day and week rates, the refundable deposit, length limits, the turnaround buffer, and the dates already booked. null when it is not for rent.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerCatalogListingRental(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerCatalogListingRental``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCatalogListingRental`: GetBuyerCatalogListingRentalResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerCatalogListingRental`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCatalogListingRentalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerCatalogListingRentalResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCheckoutUsdc

> GetBuyerCheckoutUsdcResponse GetBuyerCheckoutUsdc(ctx).Execute()

Whether Crossly takes USDC (Stripe stablecoin payments: Solana, Ethereum, Base, Polygon), and its per-order limit. Paying needs the buyer’s wallet, so it happens in the browser checkout, never unattended.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerCheckoutUsdc(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerCheckoutUsdc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCheckoutUsdc`: GetBuyerCheckoutUsdcResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerCheckoutUsdc`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCheckoutUsdcRequest struct via the builder pattern


### Return type

**GetBuyerCheckoutUsdcResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCollectionOffer

> GetBuyerCollectionOfferResponse GetBuyerCollectionOffer(ctx, id).Execute()

What an item in your collection would fetch now: the top open bid on the market for the same variant, condition and grade, and Crossly’s instant offer in store credit when that is switched on (a quote valid for a limited time).

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerCollectionOffer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerCollectionOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCollectionOffer`: GetBuyerCollectionOfferResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerCollectionOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCollectionOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerCollectionOfferResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCollectionSet

> GetBuyerCollectionSetResponse GetBuyerCollectionSet(ctx, category, brand, model).Execute()

One set: every piece (owned or missing, with its lowest ask on the market), what completing it would cost now, and the leaderboard of collectors who opted in, with your rank.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    category := "category_example" // string | 
    brand := "brand_example" // string | 
    model := "model_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerCollectionSet(context.Background(), category, brand, model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerCollectionSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCollectionSet`: GetBuyerCollectionSetResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerCollectionSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** |  | 
**brand** | **string** |  | 
**model** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCollectionSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**GetBuyerCollectionSetResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCreator

> GetBuyerCreatorResponse GetBuyerCreator(ctx, handle).Execute()

A creator storefront: profile, affiliate code, and shelves of live picks.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerCreator(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerCreator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCreator`: GetBuyerCreatorResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerCreator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCreatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerCreatorResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerFitVocabulary

> GetBuyerFitVocabularyResponse GetBuyerFitVocabulary(ctx).Execute()

Garment kinds and the measurements each takes.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerFitVocabulary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerFitVocabulary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerFitVocabulary`: GetBuyerFitVocabularyResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerFitVocabulary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerFitVocabularyRequest struct via the builder pattern


### Return type

**GetBuyerFitVocabularyResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerGiftFinderOption

> GetBuyerGiftFinderOptionResponse GetBuyerGiftFinderOption(ctx).Execute()

The gift finder’s questions: recipients, interests, budgets.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerGiftFinderOption(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerGiftFinderOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerGiftFinderOption`: GetBuyerGiftFinderOptionResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerGiftFinderOption`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerGiftFinderOptionRequest struct via the builder pattern


### Return type

**GetBuyerGiftFinderOptionResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerImpact

> GetBuyerImpactResponse GetBuyerImpact(ctx).Execute()

The sourced footprint your secondhand purchases avoided.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerImpact(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerImpact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerImpact`: GetBuyerImpactResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerImpact`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerImpactRequest struct via the builder pattern


### Return type

**GetBuyerImpactResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerLayaway

> BuyerLayawayResponse GetBuyerLayaway(ctx, id).Execute()

One layaway and its schedule.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerLayaway(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerLayaway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerLayaway`: BuyerLayawayResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerLayaway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerLayawayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerLayawayResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerListingAudioSummary

> GetBuyerListingAudioSummaryResponse GetBuyerListingAudioSummary(ctx, slug).Execute()

A spoken summary of a listing (accessibility): the text, built only from the listing’s own fields and the seller’s description, and the recording of it when one exists.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerListingAudioSummary(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerListingAudioSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerListingAudioSummary`: GetBuyerListingAudioSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerListingAudioSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerListingAudioSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerListingAudioSummaryResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerListingBoughtTogether

> GetBuyerListingBoughtTogetherResponse GetBuyerListingBoughtTogether(ctx, slug).Execute()

Frequently bought together / complete the look.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerListingBoughtTogether(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerListingBoughtTogether``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerListingBoughtTogether`: GetBuyerListingBoughtTogetherResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerListingBoughtTogether`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerListingBoughtTogetherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerListingBoughtTogetherResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerListingBreak

> GetBuyerListingBreakResponse GetBuyerListingBreak(ctx, slug).Execute()

The card break this listing sells spots in, if any; the seed is revealed after the draw.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerListingBreak(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerListingBreak``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerListingBreak`: GetBuyerListingBreakResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerListingBreak`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerListingBreakRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerListingBreakResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerListingCert

> GetBuyerListingCertResponse GetBuyerListingCert(ctx, slug).Execute()

A graded listing’s cert, checked with the grader: the seller’s claim (grader, grade, cert number), status (verified, mismatch, not_found, or unchecked where the grader has no lookup), population and how many graded higher where the grader publishes them, and the grader’s own cert page.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerListingCert(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerListingCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerListingCert`: GetBuyerListingCertResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerListingCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerListingCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerListingCertResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerListingFit

> GetBuyerListingFitResponse GetBuyerListingFit(ctx, slug).Execute()

A listing’s flat measurements, compared with your saved reference garments.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerListingFit(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerListingFit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerListingFit`: GetBuyerListingFitResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerListingFit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerListingFitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerListingFitResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerListingGroupBuy

> GetBuyerListingGroupBuyResponse GetBuyerListingGroupBuy(ctx, slug).Execute()

The group buy open on this listing, if any.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerListingGroupBuy(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerListingGroupBuy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerListingGroupBuy`: GetBuyerListingGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerListingGroupBuy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerListingGroupBuyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerListingGroupBuyResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerListingImpact

> GetBuyerListingImpactResponse GetBuyerListingImpact(ctx, slug).Execute()

The sourced footprint of making a new equivalent; null when no published figure applies.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerListingImpact(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerListingImpact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerListingImpact`: GetBuyerListingImpactResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerListingImpact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerListingImpactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerListingImpactResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerListingResaleValue

> GetBuyerListingResaleValueResponse GetBuyerListingResaleValue(ctx, slug).Execute()

What the same brand and category resold for on Crossly, and the buyback offer when enabled.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerListingResaleValue(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerListingResaleValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerListingResaleValue`: GetBuyerListingResaleValueResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerListingResaleValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerListingResaleValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerListingResaleValueResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerListingSubscribeOffer

> GetBuyerListingSubscribeOfferResponse GetBuyerListingSubscribeOffer(ctx, slug).Execute()

The subscribe & save discount and intervals on a listing; null when the seller has none.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerListingSubscribeOffer(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerListingSubscribeOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerListingSubscribeOffer`: GetBuyerListingSubscribeOfferResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerListingSubscribeOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerListingSubscribeOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerListingSubscribeOfferResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerListingTryOn

> GetBuyerListingTryOnResponse GetBuyerListingTryOn(ctx, slug).Execute()

Whether an AI try-on is offered for this listing (a wearable with a photo) and whether the image model is configured.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerListingTryOn(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerListingTryOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerListingTryOn`: GetBuyerListingTryOnResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerListingTryOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerListingTryOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerListingTryOnResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerLiveStreamAuction

> GetBuyerLiveStreamAuctionResponse GetBuyerLiveStreamAuction(ctx, id).Execute()

The auction running on a live stream right now, if any. Bid on it by its `slug` at POST /v1/buyer/auctions/:slug/bids.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerLiveStreamAuction(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerLiveStreamAuction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerLiveStreamAuction`: GetBuyerLiveStreamAuctionResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerLiveStreamAuction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerLiveStreamAuctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerLiveStreamAuctionResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerMarketPurchaseOption

> GetBuyerMarketPurchaseOptionResponse GetBuyerMarketPurchaseOption(ctx).Execute()

Before you bid or buy: whether the lowest ask is vault stock (keep it vaulted, or pay fastShipCents to rush it), and the grading runs a raw card can go straight to (pass gradeWith on the bid).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerMarketPurchaseOption(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerMarketPurchaseOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerMarketPurchaseOption`: GetBuyerMarketPurchaseOptionResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerMarketPurchaseOption`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerMarketPurchaseOptionRequest struct via the builder pattern


### Return type

**GetBuyerMarketPurchaseOptionResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerOrderEligibility

> GetBuyerOrderEligibilityResponse GetBuyerOrderEligibility(ctx, id).Execute()

What you can do with one order right now — cancel, confirm receipt, report not received, leave feedback.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerOrderEligibility(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerOrderEligibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerOrderEligibility`: GetBuyerOrderEligibilityResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerOrderEligibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerOrderEligibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerOrderEligibilityResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerOrderExport

> GetBuyerOrderExportResponse GetBuyerOrderExport(ctx).From(from).To(to).Execute()

Your order history as CSV text (up to 5,000 orders).



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    from := "from_example" // string | First day to include, YYYY-MM-DD. (optional)
    to := "to_example" // string | Last day to include, YYYY-MM-DD (inclusive). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerOrderExport(context.Background()).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerOrderExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerOrderExport`: GetBuyerOrderExportResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerOrderExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerOrderExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | First day to include, YYYY-MM-DD. | 
 **to** | **string** | Last day to include, YYYY-MM-DD (inclusive). | 

### Return type

**GetBuyerOrderExportResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerOrderPriceProtection

> GetBuyerOrderPriceProtectionResponse GetBuyerOrderPriceProtection(ctx, id).Execute()

Price protection on one order: if the item gets cheaper before `until`, the difference is credited to your Bucks automatically. `creditedCents` is what it has paid so far.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerOrderPriceProtection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerOrderPriceProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerOrderPriceProtection`: GetBuyerOrderPriceProtectionResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerOrderPriceProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerOrderPriceProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerOrderPriceProtectionResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerPack

> GetBuyerPackResponse GetBuyerPack(ctx, slug).Execute()

One pack: tiers and odds, items left, expected value, the chase items still in it.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerPack(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerPack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerPack`: GetBuyerPackResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerPack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerPackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerPackResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerPackFairness

> GetBuyerPackFairnessResponse GetBuyerPackFairness(ctx).Execute()

Your provably-fair seeds: the active server-seed commitment (sha256), your client seed and nonce, and every retired seed in full so past openings can be re-run.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerPackFairness(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerPackFairness``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerPackFairness`: GetBuyerPackFairnessResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerPackFairness`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerPackFairnessRequest struct via the builder pattern


### Return type

**GetBuyerPackFairnessResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerParty

> GetBuyerPartyResponse GetBuyerParty(ctx, slug).Execute()

One shopping party.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerParty(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerParty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerParty`: GetBuyerPartyResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerPartyResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerPassport

> GetBuyerPassportResponse GetBuyerPassport(ctx, code).Execute()

What a product passport tag opens: the item, whether it was authenticated, how many owners on Crossly and since when — never who or for how much.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "code_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerPassport(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerPassport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerPassport`: GetBuyerPassportResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerPassport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerPassportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerPassportResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerPreference

> GetBuyerPreferenceResponse GetBuyerPreference(ctx).Execute()

The shopping profile derived from that activity.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerPreference(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerPreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerPreference`: GetBuyerPreferenceResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerPreference`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerPreferenceRequest struct via the builder pattern


### Return type

**GetBuyerPreferenceResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerPriceGuide

> GetBuyerPriceGuideResponse GetBuyerPriceGuide(ctx).Execute()

What matching items actually sold for on Crossly: median, quartiles, by month, and recent sales. Needs q, brand or category.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerPriceGuide(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerPriceGuide``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerPriceGuide`: GetBuyerPriceGuideResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerPriceGuide`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerPriceGuideRequest struct via the builder pattern


### Return type

**GetBuyerPriceGuideResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerProfile

> GetBuyerProfileResponse GetBuyerProfile(ctx).Execute()

Your Crossly shopping profile — name, email, saved address, Bucks balance.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerProfile`: GetBuyerProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerProfileRequest struct via the builder pattern


### Return type

**GetBuyerProfileResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerPurchaseDelivery

> GetBuyerPurchaseDeliveryResponse GetBuyerPurchaseDelivery(ctx, id).Execute()

Where a purchase is going when not to your door: the pickup point (with its release code) or its ship-together consolidation.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerPurchaseDelivery(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerPurchaseDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerPurchaseDelivery`: GetBuyerPurchaseDeliveryResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerPurchaseDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerPurchaseDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerPurchaseDeliveryResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerPurchaseLiveMoment

> GetBuyerPurchaseLiveMomentResponse GetBuyerPurchaseLiveMoment(ctx, id).Execute()

The live show a purchase was bought in, with the offset (seconds) into its replay where it sold. null when it was not bought in a show.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerPurchaseLiveMoment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerPurchaseLiveMoment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerPurchaseLiveMoment`: GetBuyerPurchaseLiveMomentResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerPurchaseLiveMoment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerPurchaseLiveMomentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerPurchaseLiveMomentResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerPurchaseReturn

> GetBuyerPurchaseReturnResponse GetBuyerPurchaseReturn(ctx, id).Execute()

What a return would refund for each reason, by when, and any return already open.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerPurchaseReturn(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerPurchaseReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerPurchaseReturn`: GetBuyerPurchaseReturnResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerPurchaseReturn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerPurchaseReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerPurchaseReturnResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerPurchaseStory

> GetBuyerPurchaseStoryResponse GetBuyerPurchaseStory(ctx, id).Execute()

Your owner story on a purchase, passed on to the next owner if you resell it.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerPurchaseStory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerPurchaseStory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerPurchaseStory`: GetBuyerPurchaseStoryResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerPurchaseStory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerPurchaseStoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerPurchaseStoryResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerRegistry

> BuyerRegistryResponse GetBuyerRegistry(ctx, id).Execute()

One of your registries: items with how many are wanted and bought.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerRegistry(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerRegistry`: BuyerRegistryResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerRegistryResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerRegistryByToken

> BuyerRegistryResponse GetBuyerRegistryByToken(ctx, token).Execute()

A registry as a guest sees it: items, how many are still wanted, and whether it can ship. Never the address.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerRegistryByToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerRegistryByToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerRegistryByToken`: BuyerRegistryResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerRegistryByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerRegistryByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerRegistryResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerRentalEligibility

> GetBuyerRentalEligibilityResponse GetBuyerRentalEligibility(ctx).Execute()

Can you rent right now: ID-verified (a photo ID and a matching selfie, once), the verified name rentals ship to, and your standing (one rental at a time until your first comes back; none after one that never did).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerRentalEligibility(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerRentalEligibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerRentalEligibility`: GetBuyerRentalEligibilityResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerRentalEligibility`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerRentalEligibilityRequest struct via the builder pattern


### Return type

**GetBuyerRentalEligibilityResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerSellerClub

> GetBuyerSellerClubResponse GetBuyerSellerClub(ctx, sellerUserId).Execute()

A seller’s loyalty club: its tiers (spend threshold, discount, perks) and your standing. null when the seller has none.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sellerUserId := "sellerUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerSellerClub(context.Background(), sellerUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerSellerClub``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerSellerClub`: GetBuyerSellerClubResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerSellerClub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerSellerClubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerSellerClubResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerSellerTrust

> GetBuyerSellerTrustResponse GetBuyerSellerTrust(ctx, username).Execute()

Seller at a glance: member since, completed sales, independent rating, typical handling time, dispute rate, trust tier, badges — and the plain sentences the profile shows.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerSellerTrust(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerSellerTrust``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerSellerTrust`: GetBuyerSellerTrustResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerSellerTrust`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerSellerTrustRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerSellerTrustResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerTogether

> GetBuyerTogetherResponse GetBuyerTogether(ctx, code).Execute()

A shop-together room: members, the board with votes, and chat.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "code_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerTogether(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerTogether``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerTogether`: GetBuyerTogetherResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerTogether`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerTogetherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerTogetherResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerVaultMarketHistory

> GetBuyerVaultMarketHistoryResponse GetBuyerVaultMarketHistory(ctx).Execute()

Your Vault Markets value by day at the marks (oldest first, today live) over `days` (default 90), and today’s change against yesterday’s close.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerVaultMarketHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerVaultMarketHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerVaultMarketHistory`: GetBuyerVaultMarketHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerVaultMarketHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerVaultMarketHistoryRequest struct via the builder pattern


### Return type

**GetBuyerVaultMarketHistoryResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerVaultMarketIndex

> GetBuyerVaultMarketIndexResponse GetBuyerVaultMarketIndex(ctx, slug).Execute()

One index with its published daily closes (oldest first, `days` default 365) and its level now.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerVaultMarketIndex(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerVaultMarketIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerVaultMarketIndex`: GetBuyerVaultMarketIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerVaultMarketIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerVaultMarketIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerVaultMarketIndexResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerVaultMarketInstantExit

> GetBuyerVaultMarketInstantExitResponse GetBuyerVaultMarketInstantExit(ctx, id).Execute()

The Instant Exit quote for one vault unit you hold, and whether it can be sold now.

### Example

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
    resp, r, err := apiClient.BuyerApi.GetBuyerVaultMarketInstantExit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerVaultMarketInstantExit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerVaultMarketInstantExit`: GetBuyerVaultMarketInstantExitResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerVaultMarketInstantExit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerVaultMarketInstantExitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerVaultMarketInstantExitResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerVaultMarketWorth

> GetBuyerVaultMarketWorthResponse GetBuyerVaultMarketWorth(ctx).Execute()

What your vault is worth now: value at the market prices (marks come only from trades) split into reliable and indicative marks, cost and P&L, today’s change, and an estimate of what the units free to sell would fetch from the bids resting on the book now. Read only.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.GetBuyerVaultMarketWorth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.GetBuyerVaultMarketWorth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerVaultMarketWorth`: GetBuyerVaultMarketWorthResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.GetBuyerVaultMarketWorth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerVaultMarketWorthRequest struct via the builder pattern


### Return type

**GetBuyerVaultMarketWorthResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerActivity

> ListBuyerActivityResponse ListBuyerActivity(ctx).Execute()

What this buyer has compared lately.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerActivity(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerActivity`: ListBuyerActivityResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerActivity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerActivityRequest struct via the builder pattern


### Return type

**ListBuyerActivityResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerAuctionBids

> ListBuyerAuctionBidsResponse ListBuyerAuctionBids(ctx, slug).Execute()

An auction’s bid history, highest first, bidders anonymised.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    slug := "slug_example" // string | The auction listing slug.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerAuctionBids(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerAuctionBids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerAuctionBids`: ListBuyerAuctionBidsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerAuctionBids`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | The auction listing slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerAuctionBidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerAuctionBidsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerAuctions

> ListBuyerAuctionsResponse ListBuyerAuctions(ctx).Tab(tab).Execute()

Auctions you have bid in — bidding, won or lost (last 100).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tab := "tab_example" // string | bidding = still running; won = you are the winner; lost = ended, someone else won. (optional) (default to "bidding")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerAuctions(context.Background()).Tab(tab).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerAuctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerAuctions`: ListBuyerAuctionsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerAuctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerAuctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tab** | **string** | bidding &#x3D; still running; won &#x3D; you are the winner; lost &#x3D; ended, someone else won. | [default to &quot;bidding&quot;]

### Return type

**ListBuyerAuctionsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerBoards

> ListBuyerBoardsResponse ListBuyerBoards(ctx).Execute()

Your outfit boards: collages of listings, each with a share link.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerBoards(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerBoards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerBoards`: ListBuyerBoardsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerBoards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerBoardsRequest struct via the builder pattern


### Return type

**ListBuyerBoardsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerBreaks

> BuyerResponse3 ListBuyerBreaks(ctx).Execute()

Card-break spots you hold or bought, with what each pulled.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerBreaks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerBreaks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerBreaks`: BuyerResponse3
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerBreaks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerBreaksRequest struct via the builder pattern


### Return type

**BuyerResponse3**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerBuybacks

> BuyerResponse3 ListBuyerBuybacks(ctx).Execute()

Your buyback guarantees and requests.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerBuybacks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerBuybacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerBuybacks`: BuyerResponse3
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerBuybacks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerBuybacksRequest struct via the builder pattern


### Return type

**BuyerResponse3**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerCart

> ListBuyerCartResponse ListBuyerCart(ctx).Execute()

What is in your Crossly cart.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerCart(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerCart`: ListBuyerCartResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerCart`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerCartRequest struct via the builder pattern


### Return type

**ListBuyerCartResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerCartSaved

> BuyerCartResponse ListBuyerCartSaved(ctx).Execute()

Items you saved for later out of your cart (last 200).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerCartSaved(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerCartSaved``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerCartSaved`: BuyerCartResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerCartSaved`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerCartSavedRequest struct via the builder pattern


### Return type

**BuyerCartResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerCases

> ListBuyerCasesResponse ListBuyerCases(ctx).Status(status).Execute()

Resolution center — every case you opened (returns, not received, problems).



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    status := "status_example" // string | open = still being worked; closed = resolved or withdrawn; all = both. (optional) (default to "open")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerCases(context.Background()).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerCases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerCases`: ListBuyerCasesResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerCases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerCasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | open &#x3D; still being worked; closed &#x3D; resolved or withdrawn; all &#x3D; both. | [default to &quot;open&quot;]

### Return type

**ListBuyerCasesResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerCashback

> ListBuyerCashbackResponse ListBuyerCashback(ctx).Status(status).Page(page).Limit(limit).Execute()

Your Scout cashback — pending, confirmed, paid.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    status := "status_example" // string |  (optional)
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerCashback(context.Background()).Status(status).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerCashback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerCashback`: ListBuyerCashbackResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerCashback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerCashbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 25]

### Return type

**ListBuyerCashbackResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerCatalogListingVideos

> ListBuyerCatalogListingVideosResponse ListBuyerCatalogListingVideos(ctx, slug).Execute()

Videos the seller recorded for this item on request, newest first, with the question each one answers.

### Example

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
    resp, r, err := apiClient.BuyerApi.ListBuyerCatalogListingVideos(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerCatalogListingVideos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerCatalogListingVideos`: ListBuyerCatalogListingVideosResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerCatalogListingVideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerCatalogListingVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerCatalogListingVideosResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerClubs

> ListBuyerClubsResponse ListBuyerClubs(ctx).Execute()

Sellers’ loyalty clubs you belong to: your spend with each, your tier and its discount, and the next tier.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerClubs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerClubs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerClubs`: ListBuyerClubsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerClubs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerClubsRequest struct via the builder pattern


### Return type

**ListBuyerClubsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerCollectionSets

> BuyerResponse3 ListBuyerCollectionSets(ctx).Execute()

Sets you are collecting — a catalogue brand + model you own at least two pieces of — with owned / total and percent complete, most complete first; and whether you show on set leaderboards.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerCollectionSets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerCollectionSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerCollectionSets`: BuyerResponse3
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerCollectionSets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerCollectionSetsRequest struct via the builder pattern


### Return type

**BuyerResponse3**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerConsolidations

> BuyerResponse3 ListBuyerConsolidations(ctx).Execute()

Your ship-together orders: parcels at the hub, and the one box out.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerConsolidations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerConsolidations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerConsolidations`: BuyerResponse3
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerConsolidations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerConsolidationsRequest struct via the builder pattern


### Return type

**BuyerResponse3**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerConversationGuests

> ListBuyerWishlistMembersResponse ListBuyerConversationGuests(ctx, id).Execute()

Guests in a chat.

### Example

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
    resp, r, err := apiClient.BuyerApi.ListBuyerConversationGuests(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerConversationGuests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerConversationGuests`: ListBuyerWishlistMembersResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerConversationGuests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerConversationGuestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerWishlistMembersResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerFeedback

> ListBuyerFeedbackResponse ListBuyerFeedback(ctx).Execute()

Feedback you have left sellers, newest first — with the ids you revise by.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerFeedback(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerFeedback`: ListBuyerFeedbackResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerFeedback`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerFeedbackRequest struct via the builder pattern


### Return type

**ListBuyerFeedbackResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerFeedbackPending

> ListBuyerFeedbackPendingResponse ListBuyerFeedbackPending(ctx).Execute()

Orders from the last 60 days that are on their way or arrived and still need your feedback.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerFeedbackPending(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerFeedbackPending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerFeedbackPending`: ListBuyerFeedbackPendingResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerFeedbackPending`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerFeedbackPendingRequest struct via the builder pattern


### Return type

**ListBuyerFeedbackPendingResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerFitRefs

> BuyerResponse3 ListBuyerFitRefs(ctx).Execute()

Your saved reference garments for fit comparison.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerFitRefs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerFitRefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerFitRefs`: BuyerResponse3
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerFitRefs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerFitRefsRequest struct via the builder pattern


### Return type

**BuyerResponse3**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerFollowingSellers

> ListBuyerFollowingSellersResponse ListBuyerFollowingSellers(ctx).Execute()

Sellers you follow, with how many new listings each has since you last looked.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerFollowingSellers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerFollowingSellers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerFollowingSellers`: ListBuyerFollowingSellersResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerFollowingSellers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerFollowingSellersRequest struct via the builder pattern


### Return type

**ListBuyerFollowingSellersResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerGiftCardCredit

> ListBuyerGiftCardCreditResponse ListBuyerGiftCardCredit(ctx).Execute()

Gift card credit on your account, and when it expires.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerGiftCardCredit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerGiftCardCredit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerGiftCardCredit`: ListBuyerGiftCardCreditResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerGiftCardCredit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerGiftCardCreditRequest struct via the builder pattern


### Return type

**ListBuyerGiftCardCreditResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerGiftCards

> ListBuyerGiftCardsResponse ListBuyerGiftCards(ctx).Execute()

Gift cards you bought, newest first.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerGiftCards(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerGiftCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerGiftCards`: ListBuyerGiftCardsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerGiftCards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerGiftCardsRequest struct via the builder pattern


### Return type

**ListBuyerGiftCardsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerGiftFinder

> ListBuyerGiftFinderResponse ListBuyerGiftFinder(ctx).Execute()

Live listings for a recipient, interest and budget.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerGiftFinder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerGiftFinder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerGiftFinder`: ListBuyerGiftFinderResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerGiftFinder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerGiftFinderRequest struct via the builder pattern


### Return type

**ListBuyerGiftFinderResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerGiftSent

> ListBuyerGiftSentResponse ListBuyerGiftSent(ctx).Execute()

Items you sent as gifts: whether the teaser went out, whether it arrived and was opened, and any thank-you. Send one with `gift` on checkout.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerGiftSent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerGiftSent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerGiftSent`: ListBuyerGiftSentResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerGiftSent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerGiftSentRequest struct via the builder pattern


### Return type

**ListBuyerGiftSentResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerHotspots

> ListBuyerHotspotsResponse ListBuyerHotspots(ctx).Execute()

Shop-the-look tags on a story or lookbook photo: each tagged listing and where it sits on the image (x, y as fractions).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerHotspots(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerHotspots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerHotspots`: ListBuyerHotspotsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerHotspots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerHotspotsRequest struct via the builder pattern


### Return type

**ListBuyerHotspotsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerHunts

> ListBuyerHuntsResponse ListBuyerHunts(ctx).Execute()

Your buyer-agent hunts and what they did.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerHunts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerHunts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerHunts`: ListBuyerHuntsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerHunts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerHuntsRequest struct via the builder pattern


### Return type

**ListBuyerHuntsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerInstantOffers

> BuyerResponse3 ListBuyerInstantOffers(ctx).Execute()

Your instant offers: live quotes, accepted ones awaiting a label, approved ones (with the prepaid label) and paid ones.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerInstantOffers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerInstantOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerInstantOffers`: BuyerResponse3
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerInstantOffers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerInstantOffersRequest struct via the builder pattern


### Return type

**BuyerResponse3**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerLayaways

> ListBuyerLayawaysResponse ListBuyerLayaways(ctx).Execute()

Your layaways: price, paid so far, what is left, and the installment schedule. Put an item on layaway at checkout with layawayItemIds.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerLayaways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerLayaways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerLayaways`: ListBuyerLayawaysResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerLayaways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerLayawaysRequest struct via the builder pattern


### Return type

**ListBuyerLayawaysResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerListingQuickQuestions

> ListBuyerListingQuickQuestionsResponse ListBuyerListingQuickQuestions(ctx, slug).Execute()

One-tap questions this listing supports; send one as `quick` on a message.

### Example

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
    resp, r, err := apiClient.BuyerApi.ListBuyerListingQuickQuestions(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerListingQuickQuestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerListingQuickQuestions`: ListBuyerListingQuickQuestionsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerListingQuickQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerListingQuickQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerListingQuickQuestionsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerListingSellerReviews

> ListBuyerListingSellerReviewsResponse ListBuyerListingSellerReviews(ctx, slug).Execute()

The seller’s public reviews: summary, text, photos, replies.

### Example

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
    resp, r, err := apiClient.BuyerApi.ListBuyerListingSellerReviews(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerListingSellerReviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerListingSellerReviews`: ListBuyerListingSellerReviewsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerListingSellerReviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerListingSellerReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerListingSellerReviewsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerLiveShipments

> ListBuyerLiveShipmentsResponse ListBuyerLiveShipments(ctx).Execute()

Your one-box-per-show shipments: every order won in a show ships together, and postage is charged once (the rest refunded when it ships).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerLiveShipments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerLiveShipments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerLiveShipments`: ListBuyerLiveShipmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerLiveShipments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerLiveShipmentsRequest struct via the builder pattern


### Return type

**ListBuyerLiveShipmentsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerLiveStreamLineup

> ListBuyerLiveStreamLineupResponse ListBuyerLiveStreamLineup(ctx, id).Execute()

What a live show will sell, in order: each item, its price, and whether it is still available.

### Example

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
    resp, r, err := apiClient.BuyerApi.ListBuyerLiveStreamLineup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerLiveStreamLineup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerLiveStreamLineup`: ListBuyerLiveStreamLineupResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerLiveStreamLineup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerLiveStreamLineupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerLiveStreamLineupResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerMarketBids

> BuyerResponse3 ListBuyerMarketBids(ctx).Execute()

Your resting market bids.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerMarketBids(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerMarketBids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerMarketBids`: BuyerResponse3
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerMarketBids`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerMarketBidsRequest struct via the builder pattern


### Return type

**BuyerResponse3**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerMarketWatches

> BuyerResponse3 ListBuyerMarketWatches(ctx).Execute()

Market variants you watch, with the price you are waiting for.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerMarketWatches(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerMarketWatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerMarketWatches`: BuyerResponse3
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerMarketWatches`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerMarketWatchesRequest struct via the builder pattern


### Return type

**BuyerResponse3**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerOffers

> ListBuyerOffersResponse ListBuyerOffers(ctx).AwaitingBuyer(awaitingBuyer).Execute()

Offers you have made, newest first (last 100).



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awaitingBuyer := "awaitingBuyer_example" // string | true returns only seller counters waiting on you. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerOffers(context.Background()).AwaitingBuyer(awaitingBuyer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerOffers`: ListBuyerOffersResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerOffers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awaitingBuyer** | **string** | true returns only seller counters waiting on you. | 

### Return type

**ListBuyerOffersResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerOrders

> ListBuyerOrdersResponse ListBuyerOrders(ctx).Page(page).Limit(limit).Execute()

What you have bought on Crossly, newest first.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerOrders(context.Background()).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerOrders`: ListBuyerOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 25]

### Return type

**ListBuyerOrdersResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerPackOpenings

> ListBuyerPackOpeningsResponse ListBuyerPackOpenings(ctx).Execute()

Your pulls: what you got, its value, sell-back status, and the fairness proof for each (the server seed once rotated).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerPackOpenings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerPackOpenings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerPackOpenings`: ListBuyerPackOpeningsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerPackOpenings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerPackOpeningsRequest struct via the builder pattern


### Return type

**ListBuyerPackOpeningsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerPacks

> ListBuyerPacksResponse ListBuyerPacks(ctx).Execute()

Packs you can open: price, every tier with its published odds, what is left in it and its value range, and the expected value of one opening.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerPacks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerPacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerPacks`: ListBuyerPacksResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerPacks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerPacksRequest struct via the builder pattern


### Return type

**ListBuyerPacksResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerParties

> ListBuyerPartiesResponse ListBuyerParties(ctx).Execute()

Themed shopping parties, live and upcoming.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerParties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerParties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerParties`: ListBuyerPartiesResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerParties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerPartiesRequest struct via the builder pattern


### Return type

**ListBuyerPartiesResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerPartyListings

> BuyerResponse3 ListBuyerPartyListings(ctx, slug).Execute()

Listings shared to a party, featured first.

### Example

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
    resp, r, err := apiClient.BuyerApi.ListBuyerPartyListings(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerPartyListings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerPartyListings`: BuyerResponse3
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerPartyListings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerPartyListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse3**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerPickupPoints

> BuyerResponse3 ListBuyerPickupPoints(ctx).Execute()

Carrier pickup points (UPS Access Point / FedEx Hold at Location) near an address; pass one as `pickupPoint` to POST /v1/buyer/checkout.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerPickupPoints(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerPickupPoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerPickupPoints`: BuyerResponse3
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerPickupPoints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerPickupPointsRequest struct via the builder pattern


### Return type

**BuyerResponse3**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerProtectionPlans

> ListBuyerProtectionPlansResponse ListBuyerProtectionPlans(ctx).Execute()

Your protection plans (bought at checkout with `protection`): status, term, coverage, contract id, and any refund.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerProtectionPlans(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerProtectionPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerProtectionPlans`: ListBuyerProtectionPlansResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerProtectionPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerProtectionPlansRequest struct via the builder pattern


### Return type

**ListBuyerProtectionPlansResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerRegistries

> ListBuyerRegistriesResponse ListBuyerRegistries(ctx).Execute()

Your gift registries.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerRegistries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerRegistries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerRegistries`: ListBuyerRegistriesResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerRegistries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerRegistriesRequest struct via the builder pattern


### Return type

**ListBuyerRegistriesResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerRegistryGifts

> ListBuyerRegistryGiftsResponse ListBuyerRegistryGifts(ctx, id).Execute()

Your thank-you list: who bought what off the registry, and whether you have thanked them.

### Example

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
    resp, r, err := apiClient.BuyerApi.ListBuyerRegistryGifts(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerRegistryGifts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerRegistryGifts`: ListBuyerRegistryGiftsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerRegistryGifts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerRegistryGiftsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerRegistryGiftsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerRegistrySearch

> ListBuyerRegistrySearchResponse ListBuyerRegistrySearch(ctx).Execute()

Find a registry by its name or the registrant’s name. Current and recent registries only.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerRegistrySearch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerRegistrySearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerRegistrySearch`: ListBuyerRegistrySearchResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerRegistrySearch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerRegistrySearchRequest struct via the builder pattern


### Return type

**ListBuyerRegistrySearchResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerRentals

> ListBuyerRentalsResponse ListBuyerRentals(ctx).Execute()

Your rentals: dates, status (reserved, out, returned, claimed, closed, cancelled), the return label, and what happened to the deposit.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerRentals(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerRentals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerRentals`: ListBuyerRentalsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerRentals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerRentalsRequest struct via the builder pattern


### Return type

**ListBuyerRentalsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerSizeMemoryPrompts

> ListBuyerSizeMemoryPromptsResponse ListBuyerSizeMemoryPrompts(ctx).Execute()

Delivered orders with a size that you haven’t said how they fit (the last 60 days).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerSizeMemoryPrompts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerSizeMemoryPrompts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerSizeMemoryPrompts`: ListBuyerSizeMemoryPromptsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerSizeMemoryPrompts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerSizeMemoryPromptsRequest struct via the builder pattern


### Return type

**ListBuyerSizeMemoryPromptsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerSizeMemorySizes

> ListBuyerSizeMemorySizesResponse ListBuyerSizeMemorySizes(ctx).Execute()

The sizes learned from what you bought and kept, by kind of item (shoes, tops, bottoms, dresses, outerwear) and department.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerSizeMemorySizes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerSizeMemorySizes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerSizeMemorySizes`: ListBuyerSizeMemorySizesResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerSizeMemorySizes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerSizeMemorySizesRequest struct via the builder pattern


### Return type

**ListBuyerSizeMemorySizesResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerSplits

> BuyerResponse3 ListBuyerSplits(ctx).Execute()

Purchases you split, with each share’s status.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerSplits(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerSplits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerSplits`: BuyerResponse3
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerSplits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerSplitsRequest struct via the builder pattern


### Return type

**BuyerResponse3**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerSubscriptions

> ListBuyerSubscriptionsResponse ListBuyerSubscriptions(ctx).Execute()

Your subscribe & save standing orders.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerSubscriptions`: ListBuyerSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerSubscriptionsRequest struct via the builder pattern


### Return type

**ListBuyerSubscriptionsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerTogether

> ListBuyerTogetherResponse ListBuyerTogether(ctx).Execute()

Shop-together rooms you are in.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerTogether(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerTogether``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerTogether`: ListBuyerTogetherResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerTogether`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerTogetherRequest struct via the builder pattern


### Return type

**ListBuyerTogetherResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerTotalRequests

> ListBuyerTotalRequestsResponse ListBuyerTotalRequests(ctx).Execute()

Your combined-shipping total requests, newest per seller.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerTotalRequests(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerTotalRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerTotalRequests`: ListBuyerTotalRequestsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerTotalRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerTotalRequestsRequest struct via the builder pattern


### Return type

**ListBuyerTotalRequestsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerTryOns

> ListBuyerTryOnsResponse ListBuyerTryOns(ctx).Execute()

Your AI try-ons: the rendered images (labelled AI-generated wherever shown).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerTryOns(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerTryOns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerTryOns`: ListBuyerTryOnsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerTryOns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerTryOnsRequest struct via the builder pattern


### Return type

**ListBuyerTryOnsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerVault

> ListBuyerVaultResponse ListBuyerVault(ctx).Execute()

What you own in the Crossly vault: stored (authenticated, on a shelf), listed for sale, or on its way home — with the price to ship one home.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerVault(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerVault`: ListBuyerVaultResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerVault`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerVaultRequest struct via the builder pattern


### Return type

**ListBuyerVaultResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerVaultMarketHoldings

> ListBuyerVaultMarketHoldingsResponse ListBuyerVaultMarketHoldings(ctx).Execute()

Each unit you hold in the vault at its market price, with the mark’s quality (A or B reliable, C indicative, none), cost, P&L and where it stands (stored, listed, lent, pledged, at the grader); `meta` carries the totals and the units you have out on loan.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerVaultMarketHoldings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerVaultMarketHoldings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerVaultMarketHoldings`: ListBuyerVaultMarketHoldingsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerVaultMarketHoldings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerVaultMarketHoldingsRequest struct via the builder pattern


### Return type

**ListBuyerVaultMarketHoldingsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerVaultMarketIndexes

> ListBuyerVaultMarketIndexesResponse ListBuyerVaultMarketIndexes(ctx).Execute()

Crossly indexes over vault book keys (category, brand, curated, composite, segment, and your own), each with its level now and change since the last daily close. Index levels are valued only from the book’s marks.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerVaultMarketIndexes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerVaultMarketIndexes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerVaultMarketIndexes`: ListBuyerVaultMarketIndexesResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerVaultMarketIndexes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerVaultMarketIndexesRequest struct via the builder pattern


### Return type

**ListBuyerVaultMarketIndexesResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerVaultMarketInstantExit

> ListBuyerVaultMarketInstantExitResponse ListBuyerVaultMarketInstantExit(ctx).Execute()

Instant Exit quotes for what you hold, one per book key: the best resting bid (yours left out), its share of the mark, the bid levels, units held and which are free to sell, and why a key cannot be sold instantly. A quote, not an order.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerVaultMarketInstantExit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerVaultMarketInstantExit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerVaultMarketInstantExit`: ListBuyerVaultMarketInstantExitResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerVaultMarketInstantExit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerVaultMarketInstantExitRequest struct via the builder pattern


### Return type

**ListBuyerVaultMarketInstantExitResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerVaultMarketWatchlist

> ListBuyerVaultMarketWatchlistResponse ListBuyerVaultMarketWatchlist(ctx).Execute()

The book keys you watch, with mark and its quality, 24 h and 7 d change, and the last 30 daily marks.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerVaultMarketWatchlist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerVaultMarketWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerVaultMarketWatchlist`: ListBuyerVaultMarketWatchlistResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerVaultMarketWatchlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerVaultMarketWatchlistRequest struct via the builder pattern


### Return type

**ListBuyerVaultMarketWatchlistResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerVideoRequests

> ListBuyerVideoRequestsResponse ListBuyerVideoRequests(ctx).Execute()

Your video requests and their answers.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerVideoRequests(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerVideoRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerVideoRequests`: ListBuyerVideoRequestsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerVideoRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerVideoRequestsRequest struct via the builder pattern


### Return type

**ListBuyerVideoRequestsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerWishlistItems

> ListBuyerWishlistItemsResponse ListBuyerWishlistItems(ctx, id).Execute()

What is on one wishlist.

### Example

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
    resp, r, err := apiClient.BuyerApi.ListBuyerWishlistItems(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerWishlistItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerWishlistItems`: ListBuyerWishlistItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerWishlistItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerWishlistItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerWishlistItemsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerWishlistMembers

> ListBuyerWishlistMembersResponse ListBuyerWishlistMembers(ctx, id).Execute()

Co-editors of a shared wishlist.

### Example

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
    resp, r, err := apiClient.BuyerApi.ListBuyerWishlistMembers(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerWishlistMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerWishlistMembers`: ListBuyerWishlistMembersResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerWishlistMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerWishlistMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerWishlistMembersResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerWishlists

> ListBuyerWishlistsResponse ListBuyerWishlists(ctx).Execute()

Your wishlists.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.ListBuyerWishlists(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.ListBuyerWishlists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerWishlists`: ListBuyerWishlistsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.ListBuyerWishlists`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerWishlistsRequest struct via the builder pattern


### Return type

**ListBuyerWishlistsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBuyerBoard

> BuyerResponse UpdateBuyerBoard(ctx, id).Execute()

Rename a board, change its description, or make it public / private.

### Example

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
    resp, r, err := apiClient.BuyerApi.UpdateBuyerBoard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.UpdateBuyerBoard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerBoard`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.UpdateBuyerBoard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBuyerBoardLayout

> BuyerBoardResponse UpdateBuyerBoardLayout(ctx, id).Execute()

Replace a board’s layout: items by platform listing id or slug, with x, y, w (0–1) and z. At most 40.

### Example

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
    resp, r, err := apiClient.BuyerApi.UpdateBuyerBoardLayout(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.UpdateBuyerBoardLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerBoardLayout`: BuyerBoardResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.UpdateBuyerBoardLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerBoardLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerBoardResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBuyerCollectionSetLeaderboard

> UpdateBuyerCollectionSetLeaderboardResponse UpdateBuyerCollectionSetLeaderboard(ctx).Execute()

Show or hide yourself (by your public handle) on set leaderboards.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.UpdateBuyerCollectionSetLeaderboard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.UpdateBuyerCollectionSetLeaderboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerCollectionSetLeaderboard`: UpdateBuyerCollectionSetLeaderboardResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.UpdateBuyerCollectionSetLeaderboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerCollectionSetLeaderboardRequest struct via the builder pattern


### Return type

**UpdateBuyerCollectionSetLeaderboardResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBuyerFeedback

> BuyerResponse UpdateBuyerFeedback(ctx, id).UpdateBuyerFeedbackRequest(updateBuyerFeedbackRequest).Execute()

Revise feedback you left, within 60 days.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.
    updateBuyerFeedbackRequest := *openapiclient.NewUpdateBuyerFeedbackRequest(int32(123)) // UpdateBuyerFeedbackRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.UpdateBuyerFeedback(context.Background(), id).UpdateBuyerFeedbackRequest(updateBuyerFeedbackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.UpdateBuyerFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerFeedback`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.UpdateBuyerFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBuyerFeedbackRequest** | **UpdateBuyerFeedbackRequest** |  | 

### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBuyerHunt

> BuyerResponse UpdateBuyerHunt(ctx, id).Execute()

Pause, resume or stop a hunt.

### Example

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
    resp, r, err := apiClient.BuyerApi.UpdateBuyerHunt(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.UpdateBuyerHunt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerHunt`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.UpdateBuyerHunt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerHuntRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBuyerPurchaseStory

> BuyerResponse UpdateBuyerPurchaseStory(ctx, id).Execute()

Write your owner story (280 characters); it travels with the item if you resell it.

### Example

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
    resp, r, err := apiClient.BuyerApi.UpdateBuyerPurchaseStory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.UpdateBuyerPurchaseStory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerPurchaseStory`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.UpdateBuyerPurchaseStory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerPurchaseStoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBuyerRegistry

> BuyerRegistryResponse UpdateBuyerRegistry(ctx, id).Execute()

Edit a registry: title, kind, date, message, ship-to address.

### Example

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
    resp, r, err := apiClient.BuyerApi.UpdateBuyerRegistry(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.UpdateBuyerRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerRegistry`: BuyerRegistryResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.UpdateBuyerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerRegistryResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBuyerRegistryItem

> UpdateBuyerRegistryItemResponse UpdateBuyerRegistryItem(ctx, id, platformListingId).Execute()

How many of an item you want (1–50).

### Example

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
    platformListingId := "platformListingId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.UpdateBuyerRegistryItem(context.Background(), id, platformListingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.UpdateBuyerRegistryItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerRegistryItem`: UpdateBuyerRegistryItemResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.UpdateBuyerRegistryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**platformListingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerRegistryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**UpdateBuyerRegistryItemResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBuyerSubscription

> BuyerResponse UpdateBuyerSubscription(ctx, id).Execute()

Pause, resume, skip the next delivery, change the interval, or cancel.

### Example

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
    resp, r, err := apiClient.BuyerApi.UpdateBuyerSubscription(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.UpdateBuyerSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerSubscription`: BuyerResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.UpdateBuyerSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**BuyerResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBuyerWishlist

> UpdateBuyerWishlistResponse UpdateBuyerWishlist(ctx, id).UpdateBuyerWishlistRequest(updateBuyerWishlistRequest).Execute()

Rename a wishlist, or make it public / private.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id.
    updateBuyerWishlistRequest := *openapiclient.NewUpdateBuyerWishlistRequest() // UpdateBuyerWishlistRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.UpdateBuyerWishlist(context.Background(), id).UpdateBuyerWishlistRequest(updateBuyerWishlistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.UpdateBuyerWishlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerWishlist`: UpdateBuyerWishlistResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.UpdateBuyerWishlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerWishlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBuyerWishlistRequest** | **UpdateBuyerWishlistRequest** |  | 

### Return type

**UpdateBuyerWishlistResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

