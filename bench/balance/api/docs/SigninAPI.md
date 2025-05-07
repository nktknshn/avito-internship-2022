# \SigninAPI

All URIs are relative to *http://localhost:5050*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SignIn**](SigninAPI.md#SignIn) | **Post** /api/v1/signin | Sign in



## SignIn

> GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersSigninResponseBody SignIn(ctx).Payload(payload).Execute()

Sign in



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
	payload := *openapiclient.NewInternalBalanceAdaptersHttpHandlersSigninRequestBody() // InternalBalanceAdaptersHttpHandlersSigninRequestBody | Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SigninAPI.SignIn(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SigninAPI.SignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignIn`: GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersSigninResponseBody
	fmt.Fprintf(os.Stdout, "Response from `SigninAPI.SignIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**InternalBalanceAdaptersHttpHandlersSigninRequestBody**](InternalBalanceAdaptersHttpHandlersSigninRequestBody.md) | Payload | 

### Return type

[**GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersSigninResponseBody**](GithubComNktknshnAvitoInternship2022InternalBalanceAdaptersHttpHandlersHandlersBuilderResultInternalBalanceAdaptersHttpHandlersSigninResponseBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

