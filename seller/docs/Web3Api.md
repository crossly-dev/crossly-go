# \Web3Api

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWeb3Alert**](Web3Api.md#CreateWeb3Alert) | **Post** /v1/web3/alerts | Set a price alert on a token, or on a whole collection without tokenId: you are notified when one is listed at or below &#x60;price&#x60; (the chain&#39;s native coin). Setting it again moves the price.
[**CreateWeb3Auction**](Web3Api.md#CreateWeb3Auction) | **Post** /v1/web3/auctions | Start an English auction on a token in your linked wallet: &#x60;reserve&#x60; is the opening bid in the chain&#39;s native coin, &#x60;hours&#x60; its length (1–720). Bids must rise 5%; a late bid extends the end by 10 minutes.
[**CreateWeb3AuctionCancel**](Web3Api.md#CreateWeb3AuctionCancel) | **Post** /v1/web3/auctions/{id}/cancel | Call off your auction while it has no bids.
[**CreateWeb3CreatorClaim**](Web3Api.md#CreateWeb3CreatorClaim) | **Post** /v1/web3/creator/claims | Claim an EVM collection as its creator: its owner() must be one of your verified wallets. Returns &#x60;{ chain, contract }&#x60;.
[**CreateWeb3Crosslist**](Web3Api.md#CreateWeb3Crosslist) | **Post** /v1/web3/crosslist | Submit signed Seaport listings for several marketplaces at once (crossly, opensea); results per marketplace.
[**CreateWeb3DropRemind**](Web3Api.md#CreateWeb3DropRemind) | **Post** /v1/web3/drops/{id}/remind | Turn the 15-minute reminder for a drop on, or off if it already is. Returns &#x60;{ remindMe }&#x60;.
[**CreateWeb3Favorite**](Web3Api.md#CreateWeb3Favorite) | **Post** /v1/web3/favorites | Favourite an NFT, or unfavourite it if it already is. Returns &#x60;{ favorited }&#x60;.
[**CreateWeb3Launch**](Web3Api.md#CreateWeb3Launch) | **Post** /v1/web3/launches | Draft a launchpad collection: name, symbol, supply, mint price (wei of the chain’s native currency), schedule, optional holder phase and royalty. Crossly takes 5% of each mint inside the contract.
[**CreateWeb3LaunchAirdrop**](Web3Api.md#CreateWeb3LaunchAirdrop) | **Post** /v1/web3/launches/{id}/airdrop | Plan an airdrop (one free NFT each, up to 500) to &#x60;addresses&#x60; and/or the holders of &#x60;holdersOf&#x60;. Returns the transactions for the owner wallet; no payment moves.
[**CreateWeb3LaunchArt**](Web3Api.md#CreateWeb3LaunchArt) | **Post** /v1/web3/launches/{id}/art | Append per-token art (image URL, name, traits), up to the supply, before the reveal. Returns &#x60;{ count }&#x60;.
[**CreateWeb3LaunchConfirmDeploy**](Web3Api.md#CreateWeb3LaunchConfirmDeploy) | **Post** /v1/web3/launches/{id}/confirm-deploy | Confirm a sent deploy transaction: the contract must be the Crossly launch contract with this draft’s settings; it then goes live.
[**CreateWeb3LaunchFollowCreator**](Web3Api.md#CreateWeb3LaunchFollowCreator) | **Post** /v1/web3/launches/{id}/follow-creator | Follow the creator of a launch: their launches are announced to you, and again when each opens. Returns &#x60;{ following }&#x60;.
[**CreateWeb3LaunchReveal**](Web3Api.md#CreateWeb3LaunchReveal) | **Post** /v1/web3/launches/{id}/reveal | Reveal the art: one random offset maps every token to its asset for good. Returns &#x60;{ revealedAt }&#x60;.
[**CreateWeb3Order**](Web3Api.md#CreateWeb3Order) | **Post** /v1/web3/orders | Place a signed Seaport 1.6 order (listing, offer or collection offer) in the Crossly book.
[**CreateWeb3OrderCancel**](Web3Api.md#CreateWeb3OrderCancel) | **Post** /v1/web3/orders/{chain}/{hash}/cancel | Take one of your orders out of the Crossly book.
[**CreateWeb3OrderFulfillmentBatch**](Web3Api.md#CreateWeb3OrderFulfillmentBatch) | **Post** /v1/web3/orders/fulfillment-batch | Fulfilments for sweeping several listings in one Seaport transaction; unfillable ones are skipped with a reason.
[**CreateWeb3Phygital**](Web3Api.md#CreateWeb3Phygital) | **Post** /v1/web3/phygitals | Make an NFT you hold redeemable for one of your physical inventory items.
[**CreateWeb3PhygitalRedeem**](Web3Api.md#CreateWeb3PhygitalRedeem) | **Post** /v1/web3/phygitals/{id}/redeem | Redeem a phygital you hold (checked on-chain) for its physical item, with a shipping address. Once per phygital.
[**CreateWeb3PhygitalRedemptionShip**](Web3Api.md#CreateWeb3PhygitalRedemptionShip) | **Post** /v1/web3/phygitals/redemptions/{id}/ship | Mark a redemption shipped with its carrier and tracking number; the redeemer is notified.
[**CreateWeb3RaribleDraft**](Web3Api.md#CreateWeb3RaribleDraft) | **Post** /v1/web3/rarible/draft | Draft a Rarible V2 listing for an NFT in your linked wallet: the EIP-712 order to sign and whether Rarible&#39;s transfer proxy is approved.
[**CreateWeb3RaribleListing**](Web3Api.md#CreateWeb3RaribleListing) | **Post** /v1/web3/rarible/listings | Submit a signed Rarible V2 listing: the signature is verified, then the order is posted to Rarible and recorded.
[**CreateWeb3SolanaListTx**](Web3Api.md#CreateWeb3SolanaListTx) | **Post** /v1/web3/solana/{marketplace}/list-tx | Build an unsigned listing transaction on the Crossly book, Tensor or Magic Eden for a Solana NFT in your linked wallet.
[**CreateWeb3SolanaListing**](Web3Api.md#CreateWeb3SolanaListing) | **Post** /v1/web3/solana/{marketplace}/listings | Record a Crossly-book, Tensor or Magic Eden listing your wallet sent (read back from the chain first).
[**CreateWeb3SolanaOffer**](Web3Api.md#CreateWeb3SolanaOffer) | **Post** /v1/web3/solana/offers | Record a sent bid; it is recorded only once its trade state exists on chain.
[**CreateWeb3SolanaOfferAcceptTx**](Web3Api.md#CreateWeb3SolanaOfferAcceptTx) | **Post** /v1/web3/solana/offers/{hash}/accept-tx | Build the sale into a bid from your linked wallet that holds the NFT (sell + execute_sale, one transaction).
[**CreateWeb3SolanaOfferCancelTx**](Web3Api.md#CreateWeb3SolanaOfferCancelTx) | **Post** /v1/web3/solana/offers/{hash}/cancel-tx | Build the cancel of your bid: the SOL your other bids do not need comes back.
[**CreateWeb3SolanaOfferConfirm**](Web3Api.md#CreateWeb3SolanaOfferConfirm) | **Post** /v1/web3/solana/offers/{hash}/confirm | Apply a sent accept or cancel once the chain shows the bid closed.
[**CreateWeb3SolanaOfferTx**](Web3Api.md#CreateWeb3SolanaOfferTx) | **Post** /v1/web3/solana/offers/tx | Build a bid on a Solana NFT on Crossly&#39;s book for your linked wallet to sign; the SOL goes into your Auction House escrow.
[**CreateWeb3SolanaOrderBuyTx**](Web3Api.md#CreateWeb3SolanaOrderBuyTx) | **Post** /v1/web3/solana/orders/{hash}/buy-tx | Build the unsigned transaction that buys a Solana listing (price capped at the listed price).
[**CreateWeb3SolanaOrderConfirm**](Web3Api.md#CreateWeb3SolanaOrderConfirm) | **Post** /v1/web3/solana/orders/{hash}/confirm | Apply a sent Solana buy or delist to the book (the transaction is read back from the chain).
[**CreateWeb3SolanaOrderDelistTx**](Web3Api.md#CreateWeb3SolanaOrderDelistTx) | **Post** /v1/web3/solana/orders/{hash}/delist-tx | Build the unsigned transaction that takes your Solana listing down.
[**CreateWeb3SolanaTensorBidAcceptTx**](Web3Api.md#CreateWeb3SolanaTensorBidAcceptTx) | **Post** /v1/web3/solana/tensor-bids/accept-tx | Build the transaction that sells your NFT into a Tensor bid, never below &#x60;minLamports&#x60;.
[**CreateWeb3SolanaTensorBidCancelTx**](Web3Api.md#CreateWeb3SolanaTensorBidCancelTx) | **Post** /v1/web3/solana/tensor-bids/cancel-tx | Build the transaction that cancels your Tensor bid.
[**CreateWeb3SolanaTensorBidPlaceTx**](Web3Api.md#CreateWeb3SolanaTensorBidPlaceTx) | **Post** /v1/web3/solana/tensor-bids/place-tx | Build a Tensor collection bid (or trait bid with &#x60;trait&#x60;) for your Solana wallet to sign. Price in lamports.
[**CreateWeb3TokenRefresh**](Web3Api.md#CreateWeb3TokenRefresh) | **Post** /v1/web3/tokens/{chain}/{contract}/{tokenId}/refresh | Re-read a token&#39;s metadata from the chain&#39;s provider (name, traits, image — a new image is copied to Crossly&#39;s CDN). Once per token per 10 minutes.
[**CreateWeb3Transaction**](Web3Api.md#CreateWeb3Transaction) | **Post** /v1/web3/transactions | Apply a Seaport transaction&#39;s fills and cancels to the Crossly book (read from the receipt).
[**CreateWeb3WalletChallenge**](Web3Api.md#CreateWeb3WalletChallenge) | **Post** /v1/web3/wallets/challenge | Start linking a wallet: returns the message the wallet must sign.
[**CreateWeb3WalletSync**](Web3Api.md#CreateWeb3WalletSync) | **Post** /v1/web3/wallets/{id}/sync | Import a linked wallet&#39;s NFTs into inventory; archives ones that have left the wallet.
[**CreateWeb3WalletVerify**](Web3Api.md#CreateWeb3WalletVerify) | **Post** /v1/web3/wallets/verify | Finish linking a wallet with its signature over the challenge message.
[**DeleteWeb3Alert**](Web3Api.md#DeleteWeb3Alert) | **Delete** /v1/web3/alerts/{id} | Delete a price alert.
[**DeleteWeb3HolderPerk**](Web3Api.md#DeleteWeb3HolderPerk) | **Delete** /v1/web3/holder-perks/{id} | Remove a holder perk. Returns &#x60;{ ok }&#x60;.
[**DeleteWeb3Launch**](Web3Api.md#DeleteWeb3Launch) | **Delete** /v1/web3/launches/{id} | Delete a launch draft that never went live. Returns &#x60;{ ok }&#x60;.
[**DeleteWeb3LaunchFollowCreator**](Web3Api.md#DeleteWeb3LaunchFollowCreator) | **Delete** /v1/web3/launches/{id}/follow-creator | Stop following the creator of a launch. Returns &#x60;{ following }&#x60;.
[**DeleteWeb3MarketplaceKey**](Web3Api.md#DeleteWeb3MarketplaceKey) | **Delete** /v1/web3/marketplace-keys/{marketplace} | Remove your own key for a marketplace; Crossly&#39;s key is used again.
[**DeleteWeb3Wallet**](Web3Api.md#DeleteWeb3Wallet) | **Delete** /v1/web3/wallets/{id} | Unlink a wallet. Its NFTs stay in inventory but stop syncing.
[**GetWeb3Auction**](Web3Api.md#GetWeb3Auction) | **Get** /v1/web3/auctions/{chain}/{contract}/{tokenId} | A token&#39;s live (or latest) auction: reserve, end, top bid, bid count and the minimum next bid.
[**GetWeb3Browse**](Web3Api.md#GetWeb3Browse) | **Get** /v1/web3/browse | The Web3 directory&#39;s categories and chains, with how many collections each holds.
[**GetWeb3Collection**](Web3Api.md#GetWeb3Collection) | **Get** /v1/web3/collections/{chain}/{contract} | A collection: metadata and Crossly stats (floor, listed, volume, best collection offer).
[**GetWeb3CollectionDepth**](Web3Api.md#GetWeb3CollectionDepth) | **Get** /v1/web3/collections/{chain}/{contract}/depth | A collection&#39;s Crossly order book by price level — asks cheapest first, bids dearest first, with running totals — and the floor of every trait value.
[**GetWeb3CollectionInsight**](Web3Api.md#GetWeb3CollectionInsight) | **Get** /v1/web3/collections/{chain}/{contract}/insights | On-chain insights for a collection: holders, trait counts across the whole collection, recent transfers and a 30-day activity series.
[**GetWeb3CollectionItem**](Web3Api.md#GetWeb3CollectionItem) | **Get** /v1/web3/collections/{chain}/{contract}/items | A collection&#39;s NFTs: filter by traits (repeat trait&#x3D;Type:Value; OR within a type, AND across), search by name or token id, sort by id, rarity or Crossly price. Whole-collection once indexed.
[**GetWeb3CollectionTrait**](Web3Api.md#GetWeb3CollectionTrait) | **Get** /v1/web3/collections/{chain}/{contract}/traits | Every trait type and value in a collection, with how many tokens have each.
[**GetWeb3CollectionTraitToken**](Web3Api.md#GetWeb3CollectionTraitToken) | **Get** /v1/web3/collections/{chain}/{contract}/trait-tokens | Token ids in a collection with one trait — what a trait offer commits to.
[**GetWeb3Collector**](Web3Api.md#GetWeb3Collector) | **Get** /v1/web3/collector | Your collector perks: how many different NFTs you bought through Crossly, your Collector pin tier, relist rewards paid, and each NFT you still hold with the Bought on Crossly mark.
[**GetWeb3Creator**](Web3Api.md#GetWeb3Creator) | **Get** /v1/web3/creator/{chain}/{contract} | A collection’s creator dashboard: sales by marketplace and day, estimated royalties, holders, and for a launch the mint revenue and what is waiting to be withdrawn (wei of the chain’s native currency).
[**GetWeb3CrosslistQuote**](Web3Api.md#GetWeb3CrosslistQuote) | **Get** /v1/web3/crosslist/quote | Each marketplace&#39;s terms for listing a token: Crossly fee, OpenSea required fees and zone, Tensor availability.
[**GetWeb3Featured**](Web3Api.md#GetWeb3Featured) | **Get** /v1/web3/featured | The NFT home&#39;s featured collections: chosen by Crossly, else the hottest.
[**GetWeb3Inventory**](Web3Api.md#GetWeb3Inventory) | **Get** /v1/web3/inventory/{itemId} | The NFT behind an inventory item: its token, wallet, and open listings on every marketplace (null if it is not an NFT).
[**GetWeb3LaunchAllowlistRootTx**](Web3Api.md#GetWeb3LaunchAllowlistRootTx) | **Get** /v1/web3/launches/{id}/allowlist/root-tx | The setAllowlistRoot transaction for the allowlist Crossly holds (EVM, live): &#x60;{ chain, from, to, data }&#x60;.
[**GetWeb3LaunchDeployTx**](Web3Api.md#GetWeb3LaunchDeployTx) | **Get** /v1/web3/launches/{id}/deploy-tx | The contract-creation transaction for the draft: send it from the owner wallet, then confirm it. Returns &#x60;{ chain, from, data }&#x60;.
[**GetWeb3LaunchTool**](Web3Api.md#GetWeb3LaunchTool) | **Get** /v1/web3/launches/{id}/tools | A launch’s creator tools: allowlist size and root (and whether the contract has it yet), uploaded art and reveal, supply minted.
[**GetWeb3OrderFulfillment**](Web3Api.md#GetWeb3OrderFulfillment) | **Get** /v1/web3/orders/{chain}/{hash}/fulfillment | The signed order a taker passes to Seaport, after re-checking it can still fill.
[**GetWeb3Phygital**](Web3Api.md#GetWeb3Phygital) | **Get** /v1/web3/phygitals/{chain}/{contract}/{tokenId} | A token&#39;s phygital: the physical item it can be redeemed for, and whether it has been.
[**GetWeb3Phygital2**](Web3Api.md#GetWeb3Phygital2) | **Get** /v1/web3/phygitals | Phygitals you linked (with each redemption and its shipping address) and ones you redeemed.
[**GetWeb3Portfolio**](Web3Api.md#GetWeb3Portfolio) | **Get** /v1/web3/portfolio | Your NFTs with value (Crossly floor, else recent sales) and cost (what a linked wallet paid), realised sales, and USD totals at each day’s coin price.
[**GetWeb3PortfolioTax**](Web3Api.md#GetWeb3PortfolioTax) | **Get** /v1/web3/portfolio/tax | One calendar year of NFT disposals from your linked wallets as CSV text: acquired, sold, proceeds, cost basis, gain (coin and USD).
[**GetWeb3Profile**](Web3Api.md#GetWeb3Profile) | **Get** /v1/web3/profiles/{address} | A wallet&#39;s public profile: NFTs held, open Crossly listings, and activity.
[**GetWeb3Resolve**](Web3Api.md#GetWeb3Resolve) | **Get** /v1/web3/resolve/{chain}/{slug} | A collection&#39;s chain and contract from its readable URL slug (e.g. eth / boredapeyachtclub).
[**GetWeb3SolanaCollectionTensorBid**](Web3Api.md#GetWeb3SolanaCollectionTensorBid) | **Get** /v1/web3/solana/collections/{collection}/tensor-bids | A Solana collection’s Tensor bids (collection-wide and by trait), best first; price in lamports. Needs your Tensor key.
[**GetWeb3Token**](Web3Api.md#GetWeb3Token) | **Get** /v1/web3/tokens/{chain}/{contract}/{tokenId} | One NFT: metadata, traits, owner, open listings and offers, activity.
[**ListWeb3Activity**](Web3Api.md#ListWeb3Activity) | **Get** /v1/web3/activity | Listings, offers, sales and cancels in the Crossly book, newest first.
[**ListWeb3Alerts**](Web3Api.md#ListWeb3Alerts) | **Get** /v1/web3/alerts | Your NFT price alerts, with the cheapest open Crossly listing each one covers now.
[**ListWeb3Approvals**](Web3Api.md#ListWeb3Approvals) | **Get** /v1/web3/approvals | Your collections the OpenSea conduit is approved to move, per linked wallet.
[**ListWeb3Auctions**](Web3Api.md#ListWeb3Auctions) | **Get** /v1/web3/auctions | Live English auctions on the Crossly book, ending soonest first, each with its top bid and the least the next bid can be.
[**ListWeb3BrowseCollections**](Web3Api.md#ListWeb3BrowseCollections) | **Get** /v1/web3/browse/collections | Browse NFT collections by chain, category and name; sort hot (Crossly sales, listings, favourites), new, volume, floor or name.
[**ListWeb3BrowseTokens**](Web3Api.md#ListWeb3BrowseTokens) | **Get** /v1/web3/browse/tokens | Hot NFTs: what is for sale on Crossly first, then tokens from the hottest collections.
[**ListWeb3Bundles**](Web3Api.md#ListWeb3Bundles) | **Get** /v1/web3/bundles | Open bundle listings (2–30 NFTs for one price), newest first, by chain or by a contract any item is from.
[**ListWeb3CollectionHistory**](Web3Api.md#ListWeb3CollectionHistory) | **Get** /v1/web3/collections/{chain}/{contract}/history | Daily floor, sales, volume, holders and listed count for a collection, oldest first (up to 365 days).
[**ListWeb3Collections**](Web3Api.md#ListWeb3Collections) | **Get** /v1/web3/collections | Collections ranked by 30-day Crossly volume, with floor and listed count.
[**ListWeb3CreatorCollections**](Web3Api.md#ListWeb3CreatorCollections) | **Get** /v1/web3/creator/collections | Collections you are the creator of: your launches, and contracts you claimed by owning them.
[**ListWeb3Drops**](Web3Api.md#ListWeb3Drops) | **Get** /v1/web3/drops | Upcoming and live NFT drops (curated mints on the creator&#39;s own page), soonest first, with whether you asked to be reminded.
[**ListWeb3HolderPerks**](Web3Api.md#ListWeb3HolderPerks) | **Get** /v1/web3/holder-perks | Your holder perks: percent off your Crossly shop for holders of an NFT collection, best first.
[**ListWeb3Launches**](Web3Api.md#ListWeb3Launches) | **Get** /v1/web3/launches | Your launchpad collections (drafts and live), newest first, with minted count and phase.
[**ListWeb3Listings**](Web3Api.md#ListWeb3Listings) | **Get** /v1/web3/listings | Open NFT listings in the Crossly book — the cheapest per token.
[**ListWeb3MarketplaceKeys**](Web3Api.md#ListWeb3MarketplaceKeys) | **Get** /v1/web3/marketplace-keys | Your own marketplace API keys (last four characters only) and whether Crossly&#39;s key covers each marketplace.
[**ListWeb3Nfts**](Web3Api.md#ListWeb3Nfts) | **Get** /v1/web3/nfts | List the NFTs in your inventory, imported from your linked wallets.
[**ListWeb3Orders**](Web3Api.md#ListWeb3Orders) | **Get** /v1/web3/orders | Your orders in the Crossly book (listings and offers), any status.
[**ListWeb3Purchases**](Web3Api.md#ListWeb3Purchases) | **Get** /v1/web3/purchases | NFTs your linked wallets bought — on Crossly or any marketplace the chain feed saw — newest first, with price and whether you still hold each.
[**ListWeb3Rankings**](Web3Api.md#ListWeb3Rankings) | **Get** /v1/web3/rankings | Collections by sales volume over 24h, 7d or 30d, in one coin: the ETH chains together, or Polygon / Solana on their own.
[**ListWeb3Reserved**](Web3Api.md#ListWeb3Reserved) | **Get** /v1/web3/reserved | Private listings reserved for one of your linked wallets — only you can see or buy them.
[**ListWeb3Sales**](Web3Api.md#ListWeb3Sales) | **Get** /v1/web3/sales | The latest NFT sales across the directory, from the chain (Seaport and Blur) and the Crossly book, newest first; page with &#x60;before&#x60;.
[**ListWeb3Search**](Web3Api.md#ListWeb3Search) | **Get** /v1/web3/search | Find collections by contract address (every chain) or by name.
[**ListWeb3SearchTokens**](Web3Api.md#ListWeb3SearchTokens) | **Get** /v1/web3/search/tokens | Find tokens by collection name and number: &#39;pudgy 42&#39;, &#39;Bored Ape #1234&#39;.
[**ListWeb3Wallets**](Web3Api.md#ListWeb3Wallets) | **Get** /v1/web3/wallets | List your verified, linked wallets (EVM and Solana).
[**ListWeb3Watchlist**](Web3Api.md#ListWeb3Watchlist) | **Get** /v1/web3/watchlist | Your favourited NFTs, with the cheapest open Crossly listing and the last sale of each.
[**UpdateWeb3HolderPerk**](Web3Api.md#UpdateWeb3HolderPerk) | **Put** /v1/web3/holder-perks | Set the percent off (100–5000 bps) holders of a collection get across your Crossly shop; one perk per collection.
[**UpdateWeb3LaunchAllowlist**](Web3Api.md#UpdateWeb3LaunchAllowlist) | **Put** /v1/web3/launches/{id}/allowlist | Replace the early-phase allowlist (up to 10,000 wallets). EVM live launches get back &#x60;rootTx&#x60;, the setAllowlistRoot transaction for the owner wallet.
[**UpdateWeb3MarketplaceKey**](Web3Api.md#UpdateWeb3MarketplaceKey) | **Put** /v1/web3/marketplace-keys/{marketplace} | Save your own API key for a marketplace (checked against the marketplace first where it can be); stored encrypted.



## CreateWeb3Alert

> IdResponse CreateWeb3Alert(ctx).Execute()

Set a price alert on a token, or on a whole collection without tokenId: you are notified when one is listed at or below `price` (the chain's native coin). Setting it again moves the price.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3Alert(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3Alert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3Alert`: IdResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3Alert`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3AlertRequest struct via the builder pattern


### Return type

**IdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3Auction

> CreateWeb3AuctionResponse CreateWeb3Auction(ctx).Execute()

Start an English auction on a token in your linked wallet: `reserve` is the opening bid in the chain's native coin, `hours` its length (1–720). Bids must rise 5%; a late bid extends the end by 10 minutes.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3Auction(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3Auction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3Auction`: CreateWeb3AuctionResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3Auction`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3AuctionRequest struct via the builder pattern


### Return type

**CreateWeb3AuctionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3AuctionCancel

> OkResponse CreateWeb3AuctionCancel(ctx, id).Execute()

Call off your auction while it has no bids.

### Example

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
    resp, r, err := apiClient.Web3Api.CreateWeb3AuctionCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3AuctionCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3AuctionCancel`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3AuctionCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3AuctionCancelRequest struct via the builder pattern


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


## CreateWeb3CreatorClaim

> CreateWeb3CreatorClaimResponse CreateWeb3CreatorClaim(ctx).Execute()

Claim an EVM collection as its creator: its owner() must be one of your verified wallets. Returns `{ chain, contract }`.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3CreatorClaim(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3CreatorClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3CreatorClaim`: CreateWeb3CreatorClaimResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3CreatorClaim`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3CreatorClaimRequest struct via the builder pattern


### Return type

**CreateWeb3CreatorClaimResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3Crosslist

> CreateWeb3CrosslistResponse CreateWeb3Crosslist(ctx).Execute()

Submit signed Seaport listings for several marketplaces at once (crossly, opensea); results per marketplace.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3Crosslist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3Crosslist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3Crosslist`: CreateWeb3CrosslistResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3Crosslist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3CrosslistRequest struct via the builder pattern


### Return type

**CreateWeb3CrosslistResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3DropRemind

> CreateWeb3DropRemindResponse CreateWeb3DropRemind(ctx, id).Execute()

Turn the 15-minute reminder for a drop on, or off if it already is. Returns `{ remindMe }`.

### Example

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
    resp, r, err := apiClient.Web3Api.CreateWeb3DropRemind(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3DropRemind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3DropRemind`: CreateWeb3DropRemindResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3DropRemind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3DropRemindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3DropRemindResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3Favorite

> CreateWeb3FavoriteResponse CreateWeb3Favorite(ctx).Execute()

Favourite an NFT, or unfavourite it if it already is. Returns `{ favorited }`.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3Favorite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3Favorite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3Favorite`: CreateWeb3FavoriteResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3Favorite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3FavoriteRequest struct via the builder pattern


### Return type

**CreateWeb3FavoriteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3Launch

> ListWeb3LaunchesItem CreateWeb3Launch(ctx).Execute()

Draft a launchpad collection: name, symbol, supply, mint price (wei of the chain’s native currency), schedule, optional holder phase and royalty. Crossly takes 5% of each mint inside the contract.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3Launch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3Launch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3Launch`: ListWeb3LaunchesItem
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3Launch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3LaunchRequest struct via the builder pattern


### Return type

**ListWeb3LaunchesItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3LaunchAirdrop

> CreateWeb3LaunchAirdropResponse CreateWeb3LaunchAirdrop(ctx, id).Execute()

Plan an airdrop (one free NFT each, up to 500) to `addresses` and/or the holders of `holdersOf`. Returns the transactions for the owner wallet; no payment moves.

### Example

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
    resp, r, err := apiClient.Web3Api.CreateWeb3LaunchAirdrop(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3LaunchAirdrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3LaunchAirdrop`: CreateWeb3LaunchAirdropResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3LaunchAirdrop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3LaunchAirdropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3LaunchAirdropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3LaunchArt

> CreateWeb3LaunchArtResponse CreateWeb3LaunchArt(ctx, id).Execute()

Append per-token art (image URL, name, traits), up to the supply, before the reveal. Returns `{ count }`.

### Example

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
    resp, r, err := apiClient.Web3Api.CreateWeb3LaunchArt(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3LaunchArt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3LaunchArt`: CreateWeb3LaunchArtResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3LaunchArt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3LaunchArtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3LaunchArtResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3LaunchConfirmDeploy

> ListWeb3LaunchesItem CreateWeb3LaunchConfirmDeploy(ctx, id).Execute()

Confirm a sent deploy transaction: the contract must be the Crossly launch contract with this draft’s settings; it then goes live.

### Example

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
    resp, r, err := apiClient.Web3Api.CreateWeb3LaunchConfirmDeploy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3LaunchConfirmDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3LaunchConfirmDeploy`: ListWeb3LaunchesItem
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3LaunchConfirmDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3LaunchConfirmDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWeb3LaunchesItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3LaunchFollowCreator

> CreateWeb3LaunchFollowCreatorResponse CreateWeb3LaunchFollowCreator(ctx, id).Execute()

Follow the creator of a launch: their launches are announced to you, and again when each opens. Returns `{ following }`.

### Example

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
    resp, r, err := apiClient.Web3Api.CreateWeb3LaunchFollowCreator(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3LaunchFollowCreator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3LaunchFollowCreator`: CreateWeb3LaunchFollowCreatorResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3LaunchFollowCreator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3LaunchFollowCreatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3LaunchFollowCreatorResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3LaunchReveal

> CreateWeb3LaunchRevealResponse CreateWeb3LaunchReveal(ctx, id).Execute()

Reveal the art: one random offset maps every token to its asset for good. Returns `{ revealedAt }`.

### Example

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
    resp, r, err := apiClient.Web3Api.CreateWeb3LaunchReveal(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3LaunchReveal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3LaunchReveal`: CreateWeb3LaunchRevealResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3LaunchReveal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3LaunchRevealRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3LaunchRevealResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3Order

> CreateWeb3OrderResponse CreateWeb3Order(ctx).Execute()

Place a signed Seaport 1.6 order (listing, offer or collection offer) in the Crossly book.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3Order(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3Order``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3Order`: CreateWeb3OrderResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3Order`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3OrderRequest struct via the builder pattern


### Return type

**CreateWeb3OrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3OrderCancel

> OkResponse CreateWeb3OrderCancel(ctx, chain, hash).Execute()

Take one of your orders out of the Crossly book.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    hash := "hash_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3OrderCancel(context.Background(), chain, hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3OrderCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3OrderCancel`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3OrderCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3OrderCancelRequest struct via the builder pattern


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


## CreateWeb3OrderFulfillmentBatch

> CreateWeb3OrderFulfillmentBatchResponse CreateWeb3OrderFulfillmentBatch(ctx).Execute()

Fulfilments for sweeping several listings in one Seaport transaction; unfillable ones are skipped with a reason.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3OrderFulfillmentBatch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3OrderFulfillmentBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3OrderFulfillmentBatch`: CreateWeb3OrderFulfillmentBatchResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3OrderFulfillmentBatch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3OrderFulfillmentBatchRequest struct via the builder pattern


### Return type

**CreateWeb3OrderFulfillmentBatchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3Phygital

> CreateWeb3PhygitalResponse CreateWeb3Phygital(ctx).Execute()

Make an NFT you hold redeemable for one of your physical inventory items.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3Phygital(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3Phygital``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3Phygital`: CreateWeb3PhygitalResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3Phygital`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3PhygitalRequest struct via the builder pattern


### Return type

**CreateWeb3PhygitalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3PhygitalRedeem

> CreateWeb3PhygitalRedeemResponse CreateWeb3PhygitalRedeem(ctx, id).Execute()

Redeem a phygital you hold (checked on-chain) for its physical item, with a shipping address. Once per phygital.

### Example

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
    resp, r, err := apiClient.Web3Api.CreateWeb3PhygitalRedeem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3PhygitalRedeem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3PhygitalRedeem`: CreateWeb3PhygitalRedeemResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3PhygitalRedeem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3PhygitalRedeemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3PhygitalRedeemResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3PhygitalRedemptionShip

> OkResponse CreateWeb3PhygitalRedemptionShip(ctx, id).Execute()

Mark a redemption shipped with its carrier and tracking number; the redeemer is notified.

### Example

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
    resp, r, err := apiClient.Web3Api.CreateWeb3PhygitalRedemptionShip(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3PhygitalRedemptionShip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3PhygitalRedemptionShip`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3PhygitalRedemptionShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3PhygitalRedemptionShipRequest struct via the builder pattern


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


## CreateWeb3RaribleDraft

> CreateWeb3RaribleDraftResponse CreateWeb3RaribleDraft(ctx).Execute()

Draft a Rarible V2 listing for an NFT in your linked wallet: the EIP-712 order to sign and whether Rarible's transfer proxy is approved.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3RaribleDraft(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3RaribleDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3RaribleDraft`: CreateWeb3RaribleDraftResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3RaribleDraft`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3RaribleDraftRequest struct via the builder pattern


### Return type

**CreateWeb3RaribleDraftResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3RaribleListing

> OrderHashResponse CreateWeb3RaribleListing(ctx).Execute()

Submit a signed Rarible V2 listing: the signature is verified, then the order is posted to Rarible and recorded.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3RaribleListing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3RaribleListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3RaribleListing`: OrderHashResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3RaribleListing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3RaribleListingRequest struct via the builder pattern


### Return type

**OrderHashResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaListTx

> CreateWeb3SolanaListTxResponse CreateWeb3SolanaListTx(ctx, marketplace).Execute()

Build an unsigned listing transaction on the Crossly book, Tensor or Magic Eden for a Solana NFT in your linked wallet.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    marketplace := "marketplace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaListTx(context.Background(), marketplace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaListTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaListTx`: CreateWeb3SolanaListTxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaListTx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketplace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaListTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3SolanaListTxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaListing

> OrderHashResponse CreateWeb3SolanaListing(ctx, marketplace).Execute()

Record a Crossly-book, Tensor or Magic Eden listing your wallet sent (read back from the chain first).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    marketplace := "marketplace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaListing(context.Background(), marketplace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaListing`: OrderHashResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketplace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OrderHashResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaOffer

> OrderHashResponse CreateWeb3SolanaOffer(ctx).Execute()

Record a sent bid; it is recorded only once its trade state exists on chain.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaOffer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaOffer`: OrderHashResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaOffer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaOfferRequest struct via the builder pattern


### Return type

**OrderHashResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaOfferAcceptTx

> CreateWeb3SolanaListTxResponse CreateWeb3SolanaOfferAcceptTx(ctx, hash).Execute()

Build the sale into a bid from your linked wallet that holds the NFT (sell + execute_sale, one transaction).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hash := "hash_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaOfferAcceptTx(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaOfferAcceptTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaOfferAcceptTx`: CreateWeb3SolanaListTxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaOfferAcceptTx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaOfferAcceptTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3SolanaListTxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaOfferCancelTx

> CreateWeb3SolanaListTxResponse CreateWeb3SolanaOfferCancelTx(ctx, hash).Execute()

Build the cancel of your bid: the SOL your other bids do not need comes back.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hash := "hash_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaOfferCancelTx(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaOfferCancelTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaOfferCancelTx`: CreateWeb3SolanaListTxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaOfferCancelTx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaOfferCancelTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3SolanaListTxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaOfferConfirm

> CreateWeb3SolanaOfferConfirmResponse CreateWeb3SolanaOfferConfirm(ctx, hash).Execute()

Apply a sent accept or cancel once the chain shows the bid closed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hash := "hash_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaOfferConfirm(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaOfferConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaOfferConfirm`: CreateWeb3SolanaOfferConfirmResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaOfferConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaOfferConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3SolanaOfferConfirmResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaOfferTx

> CreateWeb3SolanaOfferTxResponse CreateWeb3SolanaOfferTx(ctx).Execute()

Build a bid on a Solana NFT on Crossly's book for your linked wallet to sign; the SOL goes into your Auction House escrow.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaOfferTx(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaOfferTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaOfferTx`: CreateWeb3SolanaOfferTxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaOfferTx`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaOfferTxRequest struct via the builder pattern


### Return type

**CreateWeb3SolanaOfferTxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaOrderBuyTx

> CreateWeb3SolanaOrderBuyTxResponse CreateWeb3SolanaOrderBuyTx(ctx, hash).Execute()

Build the unsigned transaction that buys a Solana listing (price capped at the listed price).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hash := "hash_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaOrderBuyTx(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaOrderBuyTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaOrderBuyTx`: CreateWeb3SolanaOrderBuyTxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaOrderBuyTx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaOrderBuyTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3SolanaOrderBuyTxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaOrderConfirm

> CreateWeb3SolanaOfferConfirmResponse CreateWeb3SolanaOrderConfirm(ctx, hash).Execute()

Apply a sent Solana buy or delist to the book (the transaction is read back from the chain).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hash := "hash_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaOrderConfirm(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaOrderConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaOrderConfirm`: CreateWeb3SolanaOfferConfirmResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaOrderConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaOrderConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3SolanaOfferConfirmResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaOrderDelistTx

> CreateWeb3SolanaListTxResponse CreateWeb3SolanaOrderDelistTx(ctx, hash).Execute()

Build the unsigned transaction that takes your Solana listing down.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hash := "hash_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaOrderDelistTx(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaOrderDelistTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaOrderDelistTx`: CreateWeb3SolanaListTxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaOrderDelistTx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaOrderDelistTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3SolanaListTxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaTensorBidAcceptTx

> CreateWeb3SolanaTensorBidPlaceTxResponse CreateWeb3SolanaTensorBidAcceptTx(ctx).Execute()

Build the transaction that sells your NFT into a Tensor bid, never below `minLamports`.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaTensorBidAcceptTx(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaTensorBidAcceptTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaTensorBidAcceptTx`: CreateWeb3SolanaTensorBidPlaceTxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaTensorBidAcceptTx`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaTensorBidAcceptTxRequest struct via the builder pattern


### Return type

**CreateWeb3SolanaTensorBidPlaceTxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaTensorBidCancelTx

> CreateWeb3SolanaTensorBidPlaceTxResponse CreateWeb3SolanaTensorBidCancelTx(ctx).Execute()

Build the transaction that cancels your Tensor bid.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaTensorBidCancelTx(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaTensorBidCancelTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaTensorBidCancelTx`: CreateWeb3SolanaTensorBidPlaceTxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaTensorBidCancelTx`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaTensorBidCancelTxRequest struct via the builder pattern


### Return type

**CreateWeb3SolanaTensorBidPlaceTxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3SolanaTensorBidPlaceTx

> CreateWeb3SolanaTensorBidPlaceTxResponse CreateWeb3SolanaTensorBidPlaceTx(ctx).Execute()

Build a Tensor collection bid (or trait bid with `trait`) for your Solana wallet to sign. Price in lamports.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3SolanaTensorBidPlaceTx(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3SolanaTensorBidPlaceTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3SolanaTensorBidPlaceTx`: CreateWeb3SolanaTensorBidPlaceTxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3SolanaTensorBidPlaceTx`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3SolanaTensorBidPlaceTxRequest struct via the builder pattern


### Return type

**CreateWeb3SolanaTensorBidPlaceTxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3TokenRefresh

> CreateWeb3TokenRefreshResponse CreateWeb3TokenRefresh(ctx, chain, contract, tokenId).Execute()

Re-read a token's metadata from the chain's provider (name, traits, image — a new image is copied to Crossly's CDN). Once per token per 10 minutes.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3TokenRefresh(context.Background(), chain, contract, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3TokenRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3TokenRefresh`: CreateWeb3TokenRefreshResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3TokenRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3TokenRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**CreateWeb3TokenRefreshResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3Transaction

> CreateWeb3TransactionResponse CreateWeb3Transaction(ctx).Execute()

Apply a Seaport transaction's fills and cancels to the Crossly book (read from the receipt).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3Transaction(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3Transaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3Transaction`: CreateWeb3TransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3Transaction`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3TransactionRequest struct via the builder pattern


### Return type

**CreateWeb3TransactionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3WalletChallenge

> CreateWeb3WalletChallengeResponse CreateWeb3WalletChallenge(ctx).Execute()

Start linking a wallet: returns the message the wallet must sign.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3WalletChallenge(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3WalletChallenge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3WalletChallenge`: CreateWeb3WalletChallengeResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3WalletChallenge`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3WalletChallengeRequest struct via the builder pattern


### Return type

**CreateWeb3WalletChallengeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3WalletSync

> CreateWeb3WalletSyncResponse CreateWeb3WalletSync(ctx, id).Execute()

Import a linked wallet's NFTs into inventory; archives ones that have left the wallet.

### Example

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
    resp, r, err := apiClient.Web3Api.CreateWeb3WalletSync(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3WalletSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3WalletSync`: CreateWeb3WalletSyncResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3WalletSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3WalletSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3WalletSyncResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWeb3WalletVerify

> CreateWeb3WalletVerifyResponse CreateWeb3WalletVerify(ctx).Execute()

Finish linking a wallet with its signature over the challenge message.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.CreateWeb3WalletVerify(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.CreateWeb3WalletVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWeb3WalletVerify`: CreateWeb3WalletVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.CreateWeb3WalletVerify`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWeb3WalletVerifyRequest struct via the builder pattern


### Return type

**CreateWeb3WalletVerifyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWeb3Alert

> OkResponse DeleteWeb3Alert(ctx, id).Execute()

Delete a price alert.

### Example

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
    resp, r, err := apiClient.Web3Api.DeleteWeb3Alert(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.DeleteWeb3Alert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWeb3Alert`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.DeleteWeb3Alert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWeb3AlertRequest struct via the builder pattern


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


## DeleteWeb3HolderPerk

> OkResponse DeleteWeb3HolderPerk(ctx, id).Execute()

Remove a holder perk. Returns `{ ok }`.

### Example

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
    resp, r, err := apiClient.Web3Api.DeleteWeb3HolderPerk(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.DeleteWeb3HolderPerk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWeb3HolderPerk`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.DeleteWeb3HolderPerk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWeb3HolderPerkRequest struct via the builder pattern


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


## DeleteWeb3Launch

> OkResponse DeleteWeb3Launch(ctx, id).Execute()

Delete a launch draft that never went live. Returns `{ ok }`.

### Example

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
    resp, r, err := apiClient.Web3Api.DeleteWeb3Launch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.DeleteWeb3Launch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWeb3Launch`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.DeleteWeb3Launch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWeb3LaunchRequest struct via the builder pattern


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


## DeleteWeb3LaunchFollowCreator

> CreateWeb3LaunchFollowCreatorResponse DeleteWeb3LaunchFollowCreator(ctx, id).Execute()

Stop following the creator of a launch. Returns `{ following }`.

### Example

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
    resp, r, err := apiClient.Web3Api.DeleteWeb3LaunchFollowCreator(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.DeleteWeb3LaunchFollowCreator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWeb3LaunchFollowCreator`: CreateWeb3LaunchFollowCreatorResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.DeleteWeb3LaunchFollowCreator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWeb3LaunchFollowCreatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWeb3LaunchFollowCreatorResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWeb3MarketplaceKey

> DeleteWeb3MarketplaceKey(ctx, marketplace).Execute()

Remove your own key for a marketplace; Crossly's key is used again.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    marketplace := "marketplace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.DeleteWeb3MarketplaceKey(context.Background(), marketplace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.DeleteWeb3MarketplaceKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketplace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWeb3MarketplaceKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWeb3Wallet

> OkResponse DeleteWeb3Wallet(ctx, id).Execute()

Unlink a wallet. Its NFTs stay in inventory but stop syncing.

### Example

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
    resp, r, err := apiClient.Web3Api.DeleteWeb3Wallet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.DeleteWeb3Wallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWeb3Wallet`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.DeleteWeb3Wallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWeb3WalletRequest struct via the builder pattern


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


## GetWeb3Auction

> GetWeb3AuctionResponse GetWeb3Auction(ctx, chain, contract, tokenId).Execute()

A token's live (or latest) auction: reserve, end, top bid, bid count and the minimum next bid.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Auction(context.Background(), chain, contract, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Auction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Auction`: GetWeb3AuctionResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Auction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3AuctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**GetWeb3AuctionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Browse

> GetWeb3BrowseResponse GetWeb3Browse(ctx).Execute()

The Web3 directory's categories and chains, with how many collections each holds.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Browse(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Browse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Browse`: GetWeb3BrowseResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Browse`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3BrowseRequest struct via the builder pattern


### Return type

**GetWeb3BrowseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Collection

> GetWeb3CollectionResponse GetWeb3Collection(ctx, chain, contract).Execute()

A collection: metadata and Crossly stats (floor, listed, volume, best collection offer).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Collection(context.Background(), chain, contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Collection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Collection`: GetWeb3CollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Collection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3CollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWeb3CollectionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3CollectionDepth

> GetWeb3CollectionDepthResponse GetWeb3CollectionDepth(ctx, chain, contract).Execute()

A collection's Crossly order book by price level — asks cheapest first, bids dearest first, with running totals — and the floor of every trait value.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3CollectionDepth(context.Background(), chain, contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3CollectionDepth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3CollectionDepth`: GetWeb3CollectionDepthResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3CollectionDepth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3CollectionDepthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWeb3CollectionDepthResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3CollectionInsight

> GetWeb3CollectionInsightResponse GetWeb3CollectionInsight(ctx, chain, contract).Execute()

On-chain insights for a collection: holders, trait counts across the whole collection, recent transfers and a 30-day activity series.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3CollectionInsight(context.Background(), chain, contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3CollectionInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3CollectionInsight`: GetWeb3CollectionInsightResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3CollectionInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3CollectionInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWeb3CollectionInsightResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3CollectionItem

> GetWeb3CollectionItemResponse GetWeb3CollectionItem(ctx, chain, contract).Execute()

A collection's NFTs: filter by traits (repeat trait=Type:Value; OR within a type, AND across), search by name or token id, sort by id, rarity or Crossly price. Whole-collection once indexed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3CollectionItem(context.Background(), chain, contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3CollectionItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3CollectionItem`: GetWeb3CollectionItemResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3CollectionItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3CollectionItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWeb3CollectionItemResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3CollectionTrait

> GetWeb3CollectionTraitResponse GetWeb3CollectionTrait(ctx, chain, contract).Execute()

Every trait type and value in a collection, with how many tokens have each.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3CollectionTrait(context.Background(), chain, contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3CollectionTrait``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3CollectionTrait`: GetWeb3CollectionTraitResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3CollectionTrait`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3CollectionTraitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWeb3CollectionTraitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3CollectionTraitToken

> GetWeb3CollectionTraitTokenResponse GetWeb3CollectionTraitToken(ctx, chain, contract).Execute()

Token ids in a collection with one trait — what a trait offer commits to.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3CollectionTraitToken(context.Background(), chain, contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3CollectionTraitToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3CollectionTraitToken`: GetWeb3CollectionTraitTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3CollectionTraitToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3CollectionTraitTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWeb3CollectionTraitTokenResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Collector

> GetWeb3CollectorResponse GetWeb3Collector(ctx).Execute()

Your collector perks: how many different NFTs you bought through Crossly, your Collector pin tier, relist rewards paid, and each NFT you still hold with the Bought on Crossly mark.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Collector(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Collector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Collector`: GetWeb3CollectorResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Collector`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3CollectorRequest struct via the builder pattern


### Return type

**GetWeb3CollectorResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Creator

> GetWeb3CreatorResponse GetWeb3Creator(ctx, chain, contract).Execute()

A collection’s creator dashboard: sales by marketplace and day, estimated royalties, holders, and for a launch the mint revenue and what is waiting to be withdrawn (wei of the chain’s native currency).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Creator(context.Background(), chain, contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Creator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Creator`: GetWeb3CreatorResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Creator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3CreatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWeb3CreatorResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3CrosslistQuote

> GetWeb3CrosslistQuoteResponse GetWeb3CrosslistQuote(ctx).Execute()

Each marketplace's terms for listing a token: Crossly fee, OpenSea required fees and zone, Tensor availability.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3CrosslistQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3CrosslistQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3CrosslistQuote`: GetWeb3CrosslistQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3CrosslistQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3CrosslistQuoteRequest struct via the builder pattern


### Return type

**GetWeb3CrosslistQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Featured

> GetWeb3FeaturedResponse GetWeb3Featured(ctx).Execute()

The NFT home's featured collections: chosen by Crossly, else the hottest.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Featured(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Featured``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Featured`: GetWeb3FeaturedResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Featured`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3FeaturedRequest struct via the builder pattern


### Return type

**GetWeb3FeaturedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Inventory

> GetWeb3InventoryResponse GetWeb3Inventory(ctx, itemId).Execute()

The NFT behind an inventory item: its token, wallet, and open listings on every marketplace (null if it is not an NFT).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemId := "itemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Inventory(context.Background(), itemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Inventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Inventory`: GetWeb3InventoryResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Inventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3InventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWeb3InventoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3LaunchAllowlistRootTx

> GetWeb3LaunchAllowlistRootTxResponse GetWeb3LaunchAllowlistRootTx(ctx, id).Execute()

The setAllowlistRoot transaction for the allowlist Crossly holds (EVM, live): `{ chain, from, to, data }`.

### Example

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
    resp, r, err := apiClient.Web3Api.GetWeb3LaunchAllowlistRootTx(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3LaunchAllowlistRootTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3LaunchAllowlistRootTx`: GetWeb3LaunchAllowlistRootTxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3LaunchAllowlistRootTx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3LaunchAllowlistRootTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWeb3LaunchAllowlistRootTxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3LaunchDeployTx

> GetWeb3LaunchDeployTxResponse GetWeb3LaunchDeployTx(ctx, id).Execute()

The contract-creation transaction for the draft: send it from the owner wallet, then confirm it. Returns `{ chain, from, data }`.

### Example

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
    resp, r, err := apiClient.Web3Api.GetWeb3LaunchDeployTx(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3LaunchDeployTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3LaunchDeployTx`: GetWeb3LaunchDeployTxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3LaunchDeployTx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3LaunchDeployTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWeb3LaunchDeployTxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3LaunchTool

> GetWeb3LaunchToolResponse GetWeb3LaunchTool(ctx, id).Execute()

A launch’s creator tools: allowlist size and root (and whether the contract has it yet), uploaded art and reveal, supply minted.

### Example

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
    resp, r, err := apiClient.Web3Api.GetWeb3LaunchTool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3LaunchTool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3LaunchTool`: GetWeb3LaunchToolResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3LaunchTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3LaunchToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWeb3LaunchToolResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3OrderFulfillment

> GetWeb3OrderFulfillmentResponse GetWeb3OrderFulfillment(ctx, chain, hash).Execute()

The signed order a taker passes to Seaport, after re-checking it can still fill.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    hash := "hash_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3OrderFulfillment(context.Background(), chain, hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3OrderFulfillment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3OrderFulfillment`: GetWeb3OrderFulfillmentResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3OrderFulfillment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3OrderFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWeb3OrderFulfillmentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Phygital

> GetWeb3PhygitalResponse GetWeb3Phygital(ctx, chain, contract, tokenId).Execute()

A token's phygital: the physical item it can be redeemed for, and whether it has been.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Phygital(context.Background(), chain, contract, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Phygital``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Phygital`: GetWeb3PhygitalResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Phygital`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3PhygitalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**GetWeb3PhygitalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Phygital2

> GetWeb3Phygital2Response GetWeb3Phygital2(ctx).Execute()

Phygitals you linked (with each redemption and its shipping address) and ones you redeemed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Phygital2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Phygital2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Phygital2`: GetWeb3Phygital2Response
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Phygital2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3Phygital2Request struct via the builder pattern


### Return type

**GetWeb3Phygital2Response**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Portfolio

> GetWeb3PortfolioResponse GetWeb3Portfolio(ctx).Execute()

Your NFTs with value (Crossly floor, else recent sales) and cost (what a linked wallet paid), realised sales, and USD totals at each day’s coin price.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Portfolio(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Portfolio``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Portfolio`: GetWeb3PortfolioResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Portfolio`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3PortfolioRequest struct via the builder pattern


### Return type

**GetWeb3PortfolioResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3PortfolioTax

> GetWeb3PortfolioTaxResponse GetWeb3PortfolioTax(ctx).Execute()

One calendar year of NFT disposals from your linked wallets as CSV text: acquired, sold, proceeds, cost basis, gain (coin and USD).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3PortfolioTax(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3PortfolioTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3PortfolioTax`: GetWeb3PortfolioTaxResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3PortfolioTax`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3PortfolioTaxRequest struct via the builder pattern


### Return type

**GetWeb3PortfolioTaxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Profile

> GetWeb3ProfileResponse GetWeb3Profile(ctx, address).Execute()

A wallet's public profile: NFTs held, open Crossly listings, and activity.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    address := "address_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Profile(context.Background(), address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Profile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Profile`: GetWeb3ProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Profile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3ProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWeb3ProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Resolve

> GetWeb3ResolveResponse GetWeb3Resolve(ctx, chain, slug).Execute()

A collection's chain and contract from its readable URL slug (e.g. eth / boredapeyachtclub).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    slug := "slug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Resolve(context.Background(), chain, slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Resolve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Resolve`: GetWeb3ResolveResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Resolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3ResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWeb3ResolveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3SolanaCollectionTensorBid

> GetWeb3SolanaCollectionTensorBidResponse GetWeb3SolanaCollectionTensorBid(ctx, collection).Execute()

A Solana collection’s Tensor bids (collection-wide and by trait), best first; price in lamports. Needs your Tensor key.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    collection := "collection_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3SolanaCollectionTensorBid(context.Background(), collection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3SolanaCollectionTensorBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3SolanaCollectionTensorBid`: GetWeb3SolanaCollectionTensorBidResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3SolanaCollectionTensorBid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3SolanaCollectionTensorBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWeb3SolanaCollectionTensorBidResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWeb3Token

> GetWeb3TokenResponse GetWeb3Token(ctx, chain, contract, tokenId).Execute()

One NFT: metadata, traits, owner, open listings and offers, activity.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.GetWeb3Token(context.Background(), chain, contract, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.GetWeb3Token``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWeb3Token`: GetWeb3TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.GetWeb3Token`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3TokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**GetWeb3TokenResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Activity

> ListWeb3ActivityResponse ListWeb3Activity(ctx).Execute()

Listings, offers, sales and cancels in the Crossly book, newest first.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Activity(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Activity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Activity`: ListWeb3ActivityResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Activity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3ActivityRequest struct via the builder pattern


### Return type

**ListWeb3ActivityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Alerts

> ListWeb3AlertsResponse ListWeb3Alerts(ctx).Execute()

Your NFT price alerts, with the cheapest open Crossly listing each one covers now.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Alerts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Alerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Alerts`: ListWeb3AlertsResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Alerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3AlertsRequest struct via the builder pattern


### Return type

**ListWeb3AlertsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Approvals

> ListWeb3ApprovalsResponse ListWeb3Approvals(ctx).Execute()

Your collections the OpenSea conduit is approved to move, per linked wallet.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Approvals(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Approvals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Approvals`: ListWeb3ApprovalsResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Approvals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3ApprovalsRequest struct via the builder pattern


### Return type

**ListWeb3ApprovalsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Auctions

> ListWeb3AuctionsResponse ListWeb3Auctions(ctx).Execute()

Live English auctions on the Crossly book, ending soonest first, each with its top bid and the least the next bid can be.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Auctions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Auctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Auctions`: ListWeb3AuctionsResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Auctions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3AuctionsRequest struct via the builder pattern


### Return type

**ListWeb3AuctionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3BrowseCollections

> ListWeb3BrowseCollectionsResponse ListWeb3BrowseCollections(ctx).Execute()

Browse NFT collections by chain, category and name; sort hot (Crossly sales, listings, favourites), new, volume, floor or name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3BrowseCollections(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3BrowseCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3BrowseCollections`: ListWeb3BrowseCollectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3BrowseCollections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3BrowseCollectionsRequest struct via the builder pattern


### Return type

**ListWeb3BrowseCollectionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3BrowseTokens

> ListWeb3BrowseTokensResponse ListWeb3BrowseTokens(ctx).Execute()

Hot NFTs: what is for sale on Crossly first, then tokens from the hottest collections.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3BrowseTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3BrowseTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3BrowseTokens`: ListWeb3BrowseTokensResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3BrowseTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3BrowseTokensRequest struct via the builder pattern


### Return type

**ListWeb3BrowseTokensResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Bundles

> ListWeb3BundlesResponse ListWeb3Bundles(ctx).Execute()

Open bundle listings (2–30 NFTs for one price), newest first, by chain or by a contract any item is from.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Bundles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Bundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Bundles`: ListWeb3BundlesResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Bundles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3BundlesRequest struct via the builder pattern


### Return type

**ListWeb3BundlesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3CollectionHistory

> ListWeb3CollectionHistoryResponse ListWeb3CollectionHistory(ctx, chain, contract).Execute()

Daily floor, sales, volume, holders and listed count for a collection, oldest first (up to 365 days).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string | 
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3CollectionHistory(context.Background(), chain, contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3CollectionHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3CollectionHistory`: ListWeb3CollectionHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3CollectionHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** |  | 
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3CollectionHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**ListWeb3CollectionHistoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Collections

> ListWeb3CollectionsResponse ListWeb3Collections(ctx).Execute()

Collections ranked by 30-day Crossly volume, with floor and listed count.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Collections(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Collections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Collections`: ListWeb3CollectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Collections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3CollectionsRequest struct via the builder pattern


### Return type

**ListWeb3CollectionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3CreatorCollections

> ListWeb3CreatorCollectionsResponse ListWeb3CreatorCollections(ctx).Execute()

Collections you are the creator of: your launches, and contracts you claimed by owning them.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3CreatorCollections(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3CreatorCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3CreatorCollections`: ListWeb3CreatorCollectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3CreatorCollections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3CreatorCollectionsRequest struct via the builder pattern


### Return type

**ListWeb3CreatorCollectionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Drops

> ListWeb3DropsResponse ListWeb3Drops(ctx).Execute()

Upcoming and live NFT drops (curated mints on the creator's own page), soonest first, with whether you asked to be reminded.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Drops(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Drops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Drops`: ListWeb3DropsResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Drops`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3DropsRequest struct via the builder pattern


### Return type

**ListWeb3DropsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3HolderPerks

> ListWeb3HolderPerksResponse ListWeb3HolderPerks(ctx).Execute()

Your holder perks: percent off your Crossly shop for holders of an NFT collection, best first.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3HolderPerks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3HolderPerks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3HolderPerks`: ListWeb3HolderPerksResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3HolderPerks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3HolderPerksRequest struct via the builder pattern


### Return type

**ListWeb3HolderPerksResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Launches

> ListWeb3LaunchesResponse ListWeb3Launches(ctx).Execute()

Your launchpad collections (drafts and live), newest first, with minted count and phase.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Launches(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Launches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Launches`: ListWeb3LaunchesResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Launches`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3LaunchesRequest struct via the builder pattern


### Return type

**ListWeb3LaunchesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Listings

> ListWeb3ListingsResponse ListWeb3Listings(ctx).Execute()

Open NFT listings in the Crossly book — the cheapest per token.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Listings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Listings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Listings`: ListWeb3ListingsResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Listings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3ListingsRequest struct via the builder pattern


### Return type

**ListWeb3ListingsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3MarketplaceKeys

> ListWeb3MarketplaceKeysResponse ListWeb3MarketplaceKeys(ctx).Execute()

Your own marketplace API keys (last four characters only) and whether Crossly's key covers each marketplace.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3MarketplaceKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3MarketplaceKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3MarketplaceKeys`: ListWeb3MarketplaceKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3MarketplaceKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3MarketplaceKeysRequest struct via the builder pattern


### Return type

**ListWeb3MarketplaceKeysResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Nfts

> ListWeb3NftsResponse ListWeb3Nfts(ctx).Chain(chain).WalletId(walletId).Page(page).Limit(limit).Execute()

List the NFTs in your inventory, imported from your linked wallets.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chain := "chain_example" // string |  (optional)
    walletId := "walletId_example" // string |  (optional)
    page := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Nfts(context.Background()).Chain(chain).WalletId(walletId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Nfts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Nfts`: ListWeb3NftsResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Nfts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3NftsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chain** | **string** |  | 
 **walletId** | **string** |  | 
 **page** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

**ListWeb3NftsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Orders

> ListWeb3OrdersResponse ListWeb3Orders(ctx).Execute()

Your orders in the Crossly book (listings and offers), any status.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Orders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Orders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Orders`: ListWeb3OrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Orders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3OrdersRequest struct via the builder pattern


### Return type

**ListWeb3OrdersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Purchases

> ListWeb3PurchasesResponse ListWeb3Purchases(ctx).Execute()

NFTs your linked wallets bought — on Crossly or any marketplace the chain feed saw — newest first, with price and whether you still hold each.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Purchases(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Purchases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Purchases`: ListWeb3PurchasesResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Purchases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3PurchasesRequest struct via the builder pattern


### Return type

**ListWeb3PurchasesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Rankings

> ListWeb3RankingsResponse ListWeb3Rankings(ctx).Execute()

Collections by sales volume over 24h, 7d or 30d, in one coin: the ETH chains together, or Polygon / Solana on their own.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Rankings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Rankings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Rankings`: ListWeb3RankingsResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Rankings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3RankingsRequest struct via the builder pattern


### Return type

**ListWeb3RankingsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Reserved

> ListWeb3BundlesResponse ListWeb3Reserved(ctx).Execute()

Private listings reserved for one of your linked wallets — only you can see or buy them.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Reserved(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Reserved``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Reserved`: ListWeb3BundlesResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Reserved`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3ReservedRequest struct via the builder pattern


### Return type

**ListWeb3BundlesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Sales

> ListWeb3SalesResponse ListWeb3Sales(ctx).Execute()

The latest NFT sales across the directory, from the chain (Seaport and Blur) and the Crossly book, newest first; page with `before`.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Sales(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Sales``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Sales`: ListWeb3SalesResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Sales`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3SalesRequest struct via the builder pattern


### Return type

**ListWeb3SalesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Search

> ListWeb3SearchResponse ListWeb3Search(ctx).Execute()

Find collections by contract address (every chain) or by name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Search(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Search`: ListWeb3SearchResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Search`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3SearchRequest struct via the builder pattern


### Return type

**ListWeb3SearchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3SearchTokens

> ListWeb3SearchTokensResponse ListWeb3SearchTokens(ctx).Execute()

Find tokens by collection name and number: 'pudgy 42', 'Bored Ape #1234'.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3SearchTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3SearchTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3SearchTokens`: ListWeb3SearchTokensResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3SearchTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3SearchTokensRequest struct via the builder pattern


### Return type

**ListWeb3SearchTokensResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Wallets

> ListWeb3WalletsResponse ListWeb3Wallets(ctx).Execute()

List your verified, linked wallets (EVM and Solana).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Wallets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Wallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Wallets`: ListWeb3WalletsResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Wallets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3WalletsRequest struct via the builder pattern


### Return type

**ListWeb3WalletsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWeb3Watchlist

> ListWeb3WatchlistResponse ListWeb3Watchlist(ctx).Execute()

Your favourited NFTs, with the cheapest open Crossly listing and the last sale of each.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.ListWeb3Watchlist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.ListWeb3Watchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWeb3Watchlist`: ListWeb3WatchlistResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.ListWeb3Watchlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeb3WatchlistRequest struct via the builder pattern


### Return type

**ListWeb3WatchlistResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWeb3HolderPerk

> ListWeb3HolderPerksItem UpdateWeb3HolderPerk(ctx).Execute()

Set the percent off (100–5000 bps) holders of a collection get across your Crossly shop; one perk per collection.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.UpdateWeb3HolderPerk(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.UpdateWeb3HolderPerk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWeb3HolderPerk`: ListWeb3HolderPerksItem
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.UpdateWeb3HolderPerk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWeb3HolderPerkRequest struct via the builder pattern


### Return type

**ListWeb3HolderPerksItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWeb3LaunchAllowlist

> UpdateWeb3LaunchAllowlistResponse UpdateWeb3LaunchAllowlist(ctx, id).Execute()

Replace the early-phase allowlist (up to 10,000 wallets). EVM live launches get back `rootTx`, the setAllowlistRoot transaction for the owner wallet.

### Example

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
    resp, r, err := apiClient.Web3Api.UpdateWeb3LaunchAllowlist(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.UpdateWeb3LaunchAllowlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWeb3LaunchAllowlist`: UpdateWeb3LaunchAllowlistResponse
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.UpdateWeb3LaunchAllowlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWeb3LaunchAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateWeb3LaunchAllowlistResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWeb3MarketplaceKey

> ListWeb3MarketplaceKeysItem UpdateWeb3MarketplaceKey(ctx, marketplace).Execute()

Save your own API key for a marketplace (checked against the marketplace first where it can be); stored encrypted.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    marketplace := "marketplace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Web3Api.UpdateWeb3MarketplaceKey(context.Background(), marketplace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Web3Api.UpdateWeb3MarketplaceKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWeb3MarketplaceKey`: ListWeb3MarketplaceKeysItem
    fmt.Fprintf(os.Stdout, "Response from `Web3Api.UpdateWeb3MarketplaceKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketplace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWeb3MarketplaceKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWeb3MarketplaceKeysItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

