# \BuyerApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpatialAuctionRoomBid**](BuyerApi.md#CreateSpatialAuctionRoomBid) | **Post** /v1/spatial/auction/rooms/{sceneId}/bid | Bid on the lot currently up in a room.



## CreateSpatialAuctionRoomBid

> CreateSpatialAuctionRoomBidResponse CreateSpatialAuctionRoomBid(ctx, sceneId).Execute()

Bid on the lot currently up in a room.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sceneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerApi.CreateSpatialAuctionRoomBid(context.Background(), sceneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerApi.CreateSpatialAuctionRoomBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpatialAuctionRoomBid`: CreateSpatialAuctionRoomBidResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerApi.CreateSpatialAuctionRoomBid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sceneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpatialAuctionRoomBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateSpatialAuctionRoomBidResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

