# \BalanceAPI

All URIs are relative to *http://localhost:5050*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalance**](BalanceAPI.md#GetBalance) | **Get** /api/v1/balance/{user_id} | Get balance



## GetBalance

> GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersGetBalanceResponseBody GetBalance(ctx, userId).Execute()

Get balance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nktknshn/avito-internship-2022-bench"
)

func main() {
	userId := int32(56) // int32 | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BalanceAPI.GetBalance(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalanceAPI.GetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalance`: GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersGetBalanceResponseBody
	fmt.Fprintf(os.Stdout, "Response from `BalanceAPI.GetBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersGetBalanceResponseBody**](GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersGetBalanceResponseBody.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

