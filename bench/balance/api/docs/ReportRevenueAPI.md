# \ReportRevenueAPI

All URIs are relative to *http://localhost:5050*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportRevenue**](ReportRevenueAPI.md#ReportRevenue) | **Get** /api/v1/report/revenue | Report revenue



## ReportRevenue

> GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersReportRevenueResponseBody ReportRevenue(ctx).Year(year).Month(month).Execute()

Report revenue



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
	year := int32(56) // int32 | Year
	month := int32(56) // int32 | Month

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportRevenueAPI.ReportRevenue(context.Background()).Year(year).Month(month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportRevenueAPI.ReportRevenue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportRevenue`: GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersReportRevenueResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ReportRevenueAPI.ReportRevenue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportRevenueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **int32** | Year | 
 **month** | **int32** | Month | 

### Return type

[**GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersReportRevenueResponseBody**](GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersReportRevenueResponseBody.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

