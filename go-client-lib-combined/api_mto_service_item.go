/*
MilMove Prime API

The Prime API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v1/`. 

API version: 0.0.1
Contact: milmove-developers@caci.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// MtoServiceItemAPIService MtoServiceItemAPI service
type MtoServiceItemAPIService service

type ApiCreateMTOServiceItemRequest struct {
	ctx context.Context
	ApiService *MtoServiceItemAPIService
	body *MTOServiceItem
}

func (r ApiCreateMTOServiceItemRequest) Body(body MTOServiceItem) ApiCreateMTOServiceItemRequest {
	r.body = &body
	return r
}

func (r ApiCreateMTOServiceItemRequest) Execute() ([]MTOServiceItem, *http.Response, error) {
	return r.ApiService.CreateMTOServiceItemExecute(r)
}

/*
CreateMTOServiceItem createMTOServiceItem

Creates one or more MTOServiceItems. Not all service items may be created, please see details below.

This endpoint supports different body definitions. In the modelType field below, select the modelType corresponding
 to the service item you wish to create and the documentation will update with the new definition.

Upon creation these items are associated with a Move Task Order and an MTO Shipment.
The request must include UUIDs for the MTO and MTO Shipment connected to this service item. Some service item types require
additional service items to be autogenerated when added - all created service items, autogenerated included,
will be returned in the response.

To update a service item, please use [updateMTOServiceItem](#operation/updateMTOServiceItem) endpoint.

---

**`MTOServiceItemOriginSIT`**

MTOServiceItemOriginSIT is a subtype of MTOServiceItem.

This model type describes a domestic origin SIT service item. Items can be created using this
model type with the following codes:

**DOFSIT**

**1st day origin SIT service item**. When a DOFSIT is requested, the API will auto-create the following group of service items:
  * DOFSIT - Domestic origin 1st day SIT
  * DOASIT - Domestic origin Additional day SIT
  * DOPSIT - Domestic origin SIT pickup
  * DOSFSC - Domestic origin SIT fuel surcharge

**DOASIT**

**Addt'l days origin SIT service item**. This represents an additional day of storage for the same item.
Additional DOASIT service items can be created and added to an existing shipment that **includes a DOFSIT service item**.

---

**`MTOServiceItemDestSIT`**

MTOServiceItemDestSIT is a subtype of MTOServiceItem.

This model type describes a domestic destination SIT service item. Items can be created using this
model type with the following codes:

**DDFSIT**

**1st day destination SIT service item**.

These additional fields are optional for creating a DDFSIT:
  * `firstAvailableDeliveryDate1`
    * string <date>
    * First available date that Prime can deliver SIT service item.
    * firstAvailableDeliveryDate1, dateOfContact1, and timeMilitary1 are required together
  * `dateOfContact1`
    * string <date>
    * Date of attempted contact by the prime corresponding to `timeMilitary1`
    * dateOfContact1, timeMilitary1, and firstAvailableDeliveryDate1 are required together
  * `timeMilitary1`
    * string\d{4}Z
    * Time of attempted contact corresponding to `dateOfContact1`, in military format.
    * timeMilitary1, dateOfContact1, and firstAvailableDeliveryDate1 are required together
  * `firstAvailableDeliveryDate2`
    * string <date>
    * Second available date that Prime can deliver SIT service item.
    * firstAvailableDeliveryDate2, dateOfContact2, and timeMilitary2 are required together
  * `dateOfContact2`
    * string <date>
    * Date of attempted contact delivery by the prime corresponding to `timeMilitary2`
    * dateOfContact2, timeMilitary2, and firstAvailableDeliveryDate2 are required together
  * `timeMilitary2`
    * string\d{4}Z
    * Time of attempted contact corresponding to `dateOfContact2`, in military format.
    * timeMilitary2, dateOfContact2, and firstAvailableDeliveryDate2 are required together

When a DDFSIT is requested, the API will auto-create the following group of service items:
  * DDFSIT - Domestic destination 1st day SIT
  * DDASIT - Domestic destination Additional day SIT
  * DDDSIT - Domestic destination SIT delivery
  * DDSFSC - Domestic destination SIT fuel surcharge

**DDASIT**

**Addt'l days destination SIT service item**. This represents an additional day of storage for the same item.
Additional DDASIT service items can be created and added to an existing shipment that **includes a DDFSIT service item**.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateMTOServiceItemRequest
*/
func (a *MtoServiceItemAPIService) CreateMTOServiceItem(ctx context.Context) ApiCreateMTOServiceItemRequest {
	return ApiCreateMTOServiceItemRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MTOServiceItem
func (a *MtoServiceItemAPIService) CreateMTOServiceItemExecute(r ApiCreateMTOServiceItemRequest) ([]MTOServiceItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MTOServiceItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MtoServiceItemAPIService.CreateMTOServiceItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mto-service-items"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateServiceRequestDocumentUploadRequest struct {
	ctx context.Context
	ApiService *MtoServiceItemAPIService
	mtoServiceItemID string
	file *os.File
}

// The file to upload.
func (r ApiCreateServiceRequestDocumentUploadRequest) File(file *os.File) ApiCreateServiceRequestDocumentUploadRequest {
	r.file = file
	return r
}

func (r ApiCreateServiceRequestDocumentUploadRequest) Execute() (*UploadWithOmissions, *http.Response, error) {
	return r.ApiService.CreateServiceRequestDocumentUploadExecute(r)
}

/*
CreateServiceRequestDocumentUpload createServiceRequestDocumentUpload

### Functionality

This endpoint **uploads** a Service Request document for a
ServiceItem.

The ServiceItem should already exist.

ServiceItems are created with the
[createMTOServiceItem](#operation/createMTOServiceItem)
endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mtoServiceItemID UUID of the service item to use.
 @return ApiCreateServiceRequestDocumentUploadRequest
*/
func (a *MtoServiceItemAPIService) CreateServiceRequestDocumentUpload(ctx context.Context, mtoServiceItemID string) ApiCreateServiceRequestDocumentUploadRequest {
	return ApiCreateServiceRequestDocumentUploadRequest{
		ApiService: a,
		ctx: ctx,
		mtoServiceItemID: mtoServiceItemID,
	}
}

// Execute executes the request
//  @return UploadWithOmissions
func (a *MtoServiceItemAPIService) CreateServiceRequestDocumentUploadExecute(r ApiCreateServiceRequestDocumentUploadRequest) (*UploadWithOmissions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UploadWithOmissions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MtoServiceItemAPIService.CreateServiceRequestDocumentUpload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mto-service-items/{mtoServiceItemID}/uploads"
	localVarPath = strings.Replace(localVarPath, "{"+"mtoServiceItemID"+"}", url.PathEscape(parameterValueToString(r.mtoServiceItemID, "mtoServiceItemID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateMTOServiceItemRequest struct {
	ctx context.Context
	ApiService *MtoServiceItemAPIService
	mtoServiceItemID string
	ifMatch *string
	body *UpdateMTOServiceItem
}

// Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error. 
func (r ApiUpdateMTOServiceItemRequest) IfMatch(ifMatch string) ApiUpdateMTOServiceItemRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiUpdateMTOServiceItemRequest) Body(body UpdateMTOServiceItem) ApiUpdateMTOServiceItemRequest {
	r.body = &body
	return r
}

func (r ApiUpdateMTOServiceItemRequest) Execute() (*MTOServiceItem, *http.Response, error) {
	return r.ApiService.UpdateMTOServiceItemExecute(r)
}

/*
UpdateMTOServiceItem updateMTOServiceItem

Updates MTOServiceItems after creation. Not all service items or fields may be updated, please see details below.

This endpoint supports different body definitions. In the modelType field below, select the modelType corresponding
 to the service item you wish to update and the documentation will update with the new definition.

* Addresses: To update a destination service item's SIT destination final address, update the shipment destination address.
For approved shipments, please use [updateShipmentDestinationAddress](#mtoShipment/updateShipmentDestinationAddress).
For shipments not yet approved, please use [updateMTOShipmentAddress](#mtoShipment/updateMTOShipmentAddress).

* SIT Service Items: Take note that when updating `sitCustomerContacted`, `sitDepartureDate`, or `sitRequestedDelivery`, we want
those to be updated on `DOASIT` (for origin SIT) and `DDASIT` (for destination SIT). If updating those values in other service
items, the office users will not have as much attention to those values.

To create a service item, please use [createMTOServiceItem](#mtoServiceItem/createMTOServiceItem)) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mtoServiceItemID UUID of service item to update.
 @return ApiUpdateMTOServiceItemRequest
*/
func (a *MtoServiceItemAPIService) UpdateMTOServiceItem(ctx context.Context, mtoServiceItemID string) ApiUpdateMTOServiceItemRequest {
	return ApiUpdateMTOServiceItemRequest{
		ApiService: a,
		ctx: ctx,
		mtoServiceItemID: mtoServiceItemID,
	}
}

// Execute executes the request
//  @return MTOServiceItem
func (a *MtoServiceItemAPIService) UpdateMTOServiceItemExecute(r ApiUpdateMTOServiceItemRequest) (*MTOServiceItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MTOServiceItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MtoServiceItemAPIService.UpdateMTOServiceItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mto-service-items/{mtoServiceItemID}"
	localVarPath = strings.Replace(localVarPath, "{"+"mtoServiceItemID"+"}", url.PathEscape(parameterValueToString(r.mtoServiceItemID, "mtoServiceItemID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ifMatch == nil {
		return localVarReturnValue, nil, reportError("ifMatch is required and must be specified")
	}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "")
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 412 {
			var v ClientError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
