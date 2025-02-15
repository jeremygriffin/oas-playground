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


// PaymentRequestAPIService PaymentRequestAPI service
type PaymentRequestAPIService service

type ApiCreatePaymentRequestRequest struct {
	ctx context.Context
	ApiService *PaymentRequestAPIService
	body *CreatePaymentRequest
}

func (r ApiCreatePaymentRequestRequest) Body(body CreatePaymentRequest) ApiCreatePaymentRequestRequest {
	r.body = &body
	return r
}

func (r ApiCreatePaymentRequestRequest) Execute() (*PaymentRequest, *http.Response, error) {
	return r.ApiService.CreatePaymentRequestExecute(r)
}

/*
CreatePaymentRequest createPaymentRequest

Creates a new instance of a paymentRequest and is assigned the status `PENDING`.
A move task order can have multiple payment requests, and
a final payment request can be marked using boolean `isFinal`.

If a `PENDING` payment request is recalculated,
a new payment request is created and the original request is
marked with the status `DEPRECATED`.

**NOTE**: In order to create a payment request for most service items, the shipment *must*
be updated with the `PrimeActualWeight` value via [updateMTOShipment](#operation/updateMTOShipment).

**FSC - Fuel Surcharge** service items require `ActualPickupDate` to be updated on the shipment.

A service item can be on several payment requests in the case of partial payment requests and payments.

In the request, if no params are necessary, then just the `serviceItem` `id` is required. For example:
```json
{
  "isFinal": false,
  "moveTaskOrderID": "uuid",
  "serviceItems": [
    {
      "id": "uuid",
    },
    {
      "id": "uuid",
      "params": [
        {
          "key": "Service Item Parameter Name",
          "value": "Service Item Parameter Value"
        }
      ]
    }
  ],
  "pointOfContact": "string"
}
```

SIT Service Items & Accepted Payment Request Parameters:
---
If `WeightBilled` is not provided then the full shipment weight (`PrimeActualWeight`) will be considered in the calculation.

**NOTE**: Diversions have a unique calcuation for payment requests without a `WeightBilled` parameter.

If you created a payment request for a diversion and `WeightBilled` is not provided, then the following will be used in the calculation:
- The lowest shipment weight (`PrimeActualWeight`) found in the diverted shipment chain.
- The lowest reweigh weight found in the diverted shipment chain.

The diverted shipment chain is created by referencing the `diversion` boolean, `divertedFromShipmentId` UUID, and matching destination to pickup addresses.
If the chain cannot be established it will fall back to the `PrimeActualWeight` of the current shipment. This is utilized because diverted shipments are all one single shipment, but going to different locations.
The lowest weight found is the true shipment weight, and thus we search the chain of shipments for the lowest weight found.

**DOFSIT - Domestic origin 1st day SIT**
```json
  "params": [
    {
      "key": "WeightBilled",
      "value": "integer"
    }
  ]
```

**DOASIT - Domestic origin add'l SIT** *(SITPaymentRequestStart & SITPaymentRequestEnd are **REQUIRED**)*
*To create a paymentRequest for this service item, the `SITPaymentRequestStart` and `SITPaymentRequestEnd` dates must not overlap previously requested SIT dates.*
```json
  "params": [
    {
      "key": "WeightBilled",
      "value": "integer"
    },
    {
      "key": "SITPaymentRequestStart",
      "value": "date"
    },
    {
      "key": "SITPaymentRequestEnd",
      "value": "date"
    }
  ]
```

**DOPSIT - Domestic origin SIT pickup**
```json
  "params": [
    {
      "key": "WeightBilled",
      "value": "integer"
    }
  ]
```

**DOSHUT - Domestic origin shuttle service**
```json
  "params": [
    {
      "key": "WeightBilled",
      "value": "integer"
    }
  ]
```

**DDFSIT - Domestic destination 1st day SIT**
```json
  "params": [
    {
      "key": "WeightBilled",
      "value": "integer"
    }
  ]
```

**DDASIT - Domestic destination add'l SIT** *(SITPaymentRequestStart & SITPaymentRequestEnd are **REQUIRED**)*
*To create a paymentRequest for this service item, the `SITPaymentRequestStart` and `SITPaymentRequestEnd` dates must not overlap previously requested SIT dates.*
```json
  "params": [
    {
      "key": "WeightBilled",
      "value": "integer"
    },
    {
      "key": "SITPaymentRequestStart",
      "value": "date"
    },
    {
      "key": "SITPaymentRequestEnd",
      "value": "date"
    }
  ]
```

**DDDSIT - Domestic destination SIT delivery**
*To create a paymentRequest for this service item, it must first have a final address set via [updateMTOServiceItem](#operation/updateMTOServiceItem).*
```json
  "params": [
    {
      "key": "WeightBilled",
      "value": "integer"
    }
  ]
```

**DDSHUT - Domestic destination shuttle service**
```json
  "params": [
    {
      "key": "WeightBilled",
      "value": "integer"
    }
  ]
```
---


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePaymentRequestRequest
*/
func (a *PaymentRequestAPIService) CreatePaymentRequest(ctx context.Context) ApiCreatePaymentRequestRequest {
	return ApiCreatePaymentRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentRequest
func (a *PaymentRequestAPIService) CreatePaymentRequestExecute(r ApiCreatePaymentRequestRequest) (*PaymentRequest, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentRequest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentRequestAPIService.CreatePaymentRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-requests"

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

type ApiCreateUploadRequest struct {
	ctx context.Context
	ApiService *PaymentRequestAPIService
	paymentRequestID string
	file *os.File
	isWeightTicket *bool
}

// The file to upload.
func (r ApiCreateUploadRequest) File(file *os.File) ApiCreateUploadRequest {
	r.file = file
	return r
}

// Indicates whether the file is a weight ticket.
func (r ApiCreateUploadRequest) IsWeightTicket(isWeightTicket bool) ApiCreateUploadRequest {
	r.isWeightTicket = &isWeightTicket
	return r
}

func (r ApiCreateUploadRequest) Execute() (*UploadWithOmissions, *http.Response, error) {
	return r.ApiService.CreateUploadExecute(r)
}

/*
CreateUpload createUpload

### Functionality
This endpoint **uploads** a Proof of Service document for a PaymentRequest.

The PaymentRequest should already exist.

Optional key of **isWeightTicket** indicates if the document is a weight ticket or not.
This will be used for partial and full deliveries and makes it easier for the Task Invoicing Officers to locate and review service item documents.
If left empty, it will assume it is NOT a weight ticket.

The formdata in the body of the POST request that is sent should look like this if it IS a weight ticket being attached to an existing payment request:
  ```json
  {
    "file": "filePath",
    "isWeightTicket": true
  }
  ```
  If the proof of service doc is NOT a weight ticket, it will look like this - or you can leave it empty:
  ```json
  {
    "file": "filePath",
    "isWeightTicket": false
  }
  ```
  ```json
  {
    "file": "filePath",
  }
  ```

PaymentRequests are created with the [createPaymentRequest](#operation/createPaymentRequest) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentRequestID UUID of payment request to use.
 @return ApiCreateUploadRequest
*/
func (a *PaymentRequestAPIService) CreateUpload(ctx context.Context, paymentRequestID string) ApiCreateUploadRequest {
	return ApiCreateUploadRequest{
		ApiService: a,
		ctx: ctx,
		paymentRequestID: paymentRequestID,
	}
}

// Execute executes the request
//  @return UploadWithOmissions
func (a *PaymentRequestAPIService) CreateUploadExecute(r ApiCreateUploadRequest) (*UploadWithOmissions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UploadWithOmissions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentRequestAPIService.CreateUpload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-requests/{paymentRequestID}/uploads"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentRequestID"+"}", url.PathEscape(parameterValueToString(r.paymentRequestID, "paymentRequestID")), -1)

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
	if r.isWeightTicket != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "isWeightTicket", r.isWeightTicket, "")
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
