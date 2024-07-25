# \MoveTaskOrderAPI

All URIs are relative to */prime/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExcessWeightRecord**](MoveTaskOrderAPI.md#CreateExcessWeightRecord) | **Post** /move-task-orders/{moveTaskOrderID}/excess-weight-record | createExcessWeightRecord
[**DownloadMoveOrder**](MoveTaskOrderAPI.md#DownloadMoveOrder) | **Get** /moves/{locator}/documents | Downloads move order as a PDF
[**GetMoveTaskOrder**](MoveTaskOrderAPI.md#GetMoveTaskOrder) | **Get** /move-task-orders/{moveID} | getMoveTaskOrder
[**ListMoves**](MoveTaskOrderAPI.md#ListMoves) | **Get** /moves | listMoves
[**UpdateMTOPostCounselingInformation**](MoveTaskOrderAPI.md#UpdateMTOPostCounselingInformation) | **Patch** /move-task-orders/{moveTaskOrderID}/post-counseling-info | updateMTOPostCounselingInformation



## CreateExcessWeightRecord

> ExcessWeightRecord CreateExcessWeightRecord(ctx, moveTaskOrderID).File(file).Execute()

createExcessWeightRecord



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
	moveTaskOrderID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the move being updated.
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoveTaskOrderAPI.CreateExcessWeightRecord(context.Background(), moveTaskOrderID).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoveTaskOrderAPI.CreateExcessWeightRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExcessWeightRecord`: ExcessWeightRecord
	fmt.Fprintf(os.Stdout, "Response from `MoveTaskOrderAPI.CreateExcessWeightRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moveTaskOrderID** | **string** | UUID of the move being updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExcessWeightRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The file to upload. | 

### Return type

[**ExcessWeightRecord**](ExcessWeightRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadMoveOrder

> *os.File DownloadMoveOrder(ctx, locator).Type_(type_).Execute()

Downloads move order as a PDF



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
	locator := "locator_example" // string | the locator code for move order to be downloaded
	type_ := "type__example" // string | upload type (optional) (default to "ALL")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoveTaskOrderAPI.DownloadMoveOrder(context.Background(), locator).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoveTaskOrderAPI.DownloadMoveOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadMoveOrder`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `MoveTaskOrderAPI.DownloadMoveOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | the locator code for move order to be downloaded | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadMoveOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | upload type | [default to &quot;ALL&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMoveTaskOrder

> MoveTaskOrder GetMoveTaskOrder(ctx, moveID).Execute()

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
	moveID := "moveID_example" // string | UUID or MoveCode of move task order to use.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoveTaskOrderAPI.GetMoveTaskOrder(context.Background(), moveID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoveTaskOrderAPI.GetMoveTaskOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMoveTaskOrder`: MoveTaskOrder
	fmt.Fprintf(os.Stdout, "Response from `MoveTaskOrderAPI.GetMoveTaskOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moveID** | **string** | UUID or MoveCode of move task order to use. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoveTaskOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MoveTaskOrder**](MoveTaskOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMoves

> []ListMove ListMoves(ctx).Since(since).Execute()

listMoves



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	since := time.Now() // time.Time | Only return moves updated since this time. Formatted like \"2021-07-23T18:30:47.116Z\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoveTaskOrderAPI.ListMoves(context.Background()).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoveTaskOrderAPI.ListMoves``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMoves`: []ListMove
	fmt.Fprintf(os.Stdout, "Response from `MoveTaskOrderAPI.ListMoves`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMovesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **time.Time** | Only return moves updated since this time. Formatted like \&quot;2021-07-23T18:30:47.116Z\&quot; | 

### Return type

[**[]ListMove**](ListMove.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMTOPostCounselingInformation

> MoveTaskOrder UpdateMTOPostCounselingInformation(ctx, moveTaskOrderID).IfMatch(ifMatch).Execute()

updateMTOPostCounselingInformation



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
	moveTaskOrderID := "moveTaskOrderID_example" // string | ID of move task order to use.
	ifMatch := "ifMatch_example" // string | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoveTaskOrderAPI.UpdateMTOPostCounselingInformation(context.Background(), moveTaskOrderID).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoveTaskOrderAPI.UpdateMTOPostCounselingInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMTOPostCounselingInformation`: MoveTaskOrder
	fmt.Fprintf(os.Stdout, "Response from `MoveTaskOrderAPI.UpdateMTOPostCounselingInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moveTaskOrderID** | **string** | ID of move task order to use. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMTOPostCounselingInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  | 

### Return type

[**MoveTaskOrder**](MoveTaskOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

