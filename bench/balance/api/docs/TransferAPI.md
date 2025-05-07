# \TransferAPI

All URIs are relative to *http://localhost:5050*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Transfer**](TransferAPI.md#Transfer) | **Post** /api/v1/balance/transfer | Transfer



## Transfer

> GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultEmpty Transfer(ctx).Payload(payload).Execute()

Transfer



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
	payload := *openapiclient.NewInternalBalanceAdaptersHttpHandlersTransferRequestBody() // InternalBalanceAdaptersHttpHandlersTransferRequestBody | Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.Transfer(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.Transfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Transfer`: GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultEmpty
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.Transfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**InternalBalanceAdaptersHttpHandlersTransferRequestBody**](InternalBalanceAdaptersHttpHandlersTransferRequestBody.md) | Payload | 

### Return type

[**GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultEmpty**](GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultEmpty.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

