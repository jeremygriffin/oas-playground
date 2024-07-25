# \MoveTaskOrderAPI

All URIs are relative to */prime/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMoveTaskOrder**](MoveTaskOrderAPI.md#GetMoveTaskOrder) | **Get** /move-task-orders/{moveID} | getMoveTaskOrder



## GetMoveTaskOrder

> MoveTaskOrderV3V3 GetMoveTaskOrder(ctx, moveID).Execute()

getMoveTaskOrder



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
	moveID := "moveID_example" // string | UUID or MoveCode of move task order to use........

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoveTaskOrderAPI.GetMoveTaskOrder(context.Background(), moveID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoveTaskOrderAPI.GetMoveTaskOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMoveTaskOrder`: MoveTaskOrderV3V3
	fmt.Fprintf(os.Stdout, "Response from `MoveTaskOrderAPI.GetMoveTaskOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moveID** | **string** | UUID or MoveCode of move task order to use........ | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoveTaskOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MoveTaskOrderV3V3**](MoveTaskOrderV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

