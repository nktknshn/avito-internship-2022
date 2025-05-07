# \ReserveConfirmAPI

All URIs are relative to *http://localhost:5050*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReserveConfirm**](ReserveConfirmAPI.md#ReserveConfirm) | **Post** /api/v1/balance/reserve/confirm | Reserve confirm



## ReserveConfirm

> GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultEmpty ReserveConfirm(ctx).Payload(payload).Execute()

Reserve confirm



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
	payload := *openapiclient.NewInternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody() // InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody | Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReserveConfirmAPI.ReserveConfirm(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReserveConfirmAPI.ReserveConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReserveConfirm`: GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultEmpty
	fmt.Fprintf(os.Stdout, "Response from `ReserveConfirmAPI.ReserveConfirm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReserveConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody**](InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody.md) | Payload | 

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

