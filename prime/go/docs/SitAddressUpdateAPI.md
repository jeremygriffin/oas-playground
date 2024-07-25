# \SitAddressUpdateAPI

All URIs are relative to */prime/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSITAddressUpdateRequest**](SitAddressUpdateAPI.md#CreateSITAddressUpdateRequest) | **Post** /sit-address-updates | createSITAddressUpdateRequest



## CreateSITAddressUpdateRequest

> SitAddressUpdate CreateSITAddressUpdateRequest(ctx).Body(body).Execute()

createSITAddressUpdateRequest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := *openapiclient.NewCreateSITAddressUpdateRequest("Customer reached out to me this week & let me know they want to move closer to family.") // CreateSITAddressUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitAddressUpdateAPI.CreateSITAddressUpdateRequest(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitAddressUpdateAPI.CreateSITAddressUpdateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSITAddressUpdateRequest`: SitAddressUpdate
	fmt.Fprintf(os.Stdout, "Response from `SitAddressUpdateAPI.CreateSITAddressUpdateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSITAddressUpdateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateSITAddressUpdateRequest**](CreateSITAddressUpdateRequest.md) |  | 

### Return type

[**SitAddressUpdate**](SitAddressUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

