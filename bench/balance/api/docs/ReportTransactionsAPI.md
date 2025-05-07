# \ReportTransactionsAPI

All URIs are relative to *http://localhost:5050*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportTransactions**](ReportTransactionsAPI.md#ReportTransactions) | **Get** /api/v1/report/transactions/{user_id} | Report transactions



## ReportTransactions

> GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersReportTransactionsResponseBody ReportTransactions(ctx, userId).Limit(limit).Sorting(sorting).SortingDirection(sortingDirection).Cursor(cursor).Execute()

Report transactions



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
	limit := int32(56) // int32 | Limit
	sorting := "sorting_example" // string | Sorting
	sortingDirection := "sortingDirection_example" // string | Sorting Direction
	cursor := "cursor_example" // string | Cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportTransactionsAPI.ReportTransactions(context.Background(), userId).Limit(limit).Sorting(sorting).SortingDirection(sortingDirection).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportTransactionsAPI.ReportTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportTransactions`: GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersReportTransactionsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ReportTransactionsAPI.ReportTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit | 
 **sorting** | **string** | Sorting | 
 **sortingDirection** | **string** | Sorting Direction | 
 **cursor** | **string** | Cursor | 

### Return type

[**GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersReportTransactionsResponseBody**](GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersReportTransactionsResponseBody.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

