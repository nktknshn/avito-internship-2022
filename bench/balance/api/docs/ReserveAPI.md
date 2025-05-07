# \ReserveAPI

All URIs are relative to *http://localhost:5050*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Reserve**](ReserveAPI.md#Reserve) | **Post** /api/v1/balance/reserve | Reserve



## Reserve

> GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultEmpty Reserve(ctx).Payload(payload).Execute()

Reserve



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
	payload := *openapiclient.NewInternalBalanceAdaptersHttpHandlersReserveRequestBody() // InternalBalanceAdaptersHttpHandlersReserveRequestBody | Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReserveAPI.Reserve(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReserveAPI.Reserve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Reserve`: GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultEmpty
	fmt.Fprintf(os.Stdout, "Response from `ReserveAPI.Reserve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReserveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**InternalBalanceAdaptersHttpHandlersReserveRequestBody**](InternalBalanceAdaptersHttpHandlersReserveRequestBody.md) | Payload | 

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

