# GeminiCommerce\Tokentracker\TokenTrackerAPI

All URIs are relative to *https://token-tracker.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokenTrackerAdjustTokenBalance**](TokenTrackerAPI.md#TokenTrackerAdjustTokenBalance) | **Post** /v1/{tenantId}/adjust_token_balance | AdjustTokenBalance
[**TokenTrackerGetTokenBalance**](TokenTrackerAPI.md#TokenTrackerGetTokenBalance) | **Post** /v1/{tenantId}/get_token_balance | GetTokenBalance
[**TokenTrackerStripeWebhook**](TokenTrackerAPI.md#TokenTrackerStripeWebhook) | **Post** /v1/stripe/webhook | StripeWebhook



## TokenTrackerAdjustTokenBalance

> TokentrackerAdjustTokenBalanceResponse TokenTrackerAdjustTokenBalance(ctx, tenantId).Body(body).Execute()

AdjustTokenBalance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-tokentracker"
)

func main() {
	tenantId := "tenantId_example" // string | 
	body := *openapiclient.NewTokenTrackerAdjustTokenBalanceRequest() // TokenTrackerAdjustTokenBalanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenTrackerAPI.TokenTrackerAdjustTokenBalance(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenTrackerAPI.TokenTrackerAdjustTokenBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenTrackerAdjustTokenBalance`: TokentrackerAdjustTokenBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenTrackerAPI.TokenTrackerAdjustTokenBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenTrackerAdjustTokenBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TokenTrackerAdjustTokenBalanceRequest**](TokenTrackerAdjustTokenBalanceRequest.md) |  | 

### Return type

[**TokentrackerAdjustTokenBalanceResponse**](TokentrackerAdjustTokenBalanceResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenTrackerGetTokenBalance

> TokentrackerGetTokenBalanceResponse TokenTrackerGetTokenBalance(ctx, tenantId).Body(body).Execute()

GetTokenBalance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-tokentracker"
)

func main() {
	tenantId := "tenantId_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenTrackerAPI.TokenTrackerGetTokenBalance(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenTrackerAPI.TokenTrackerGetTokenBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenTrackerGetTokenBalance`: TokentrackerGetTokenBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenTrackerAPI.TokenTrackerGetTokenBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenTrackerGetTokenBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TokentrackerGetTokenBalanceResponse**](TokentrackerGetTokenBalanceResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenTrackerStripeWebhook

> map[string]interface{} TokenTrackerStripeWebhook(ctx).Body(body).Execute()

StripeWebhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-tokentracker"
)

func main() {
	body := *openapiclient.NewTokentrackerStripeWebhookRequest() // TokentrackerStripeWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenTrackerAPI.TokenTrackerStripeWebhook(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenTrackerAPI.TokenTrackerStripeWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenTrackerStripeWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TokenTrackerAPI.TokenTrackerStripeWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenTrackerStripeWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TokentrackerStripeWebhookRequest**](TokentrackerStripeWebhookRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

