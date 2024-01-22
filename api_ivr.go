
/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type IVRApiService service
/*
IVRApiService Create IVR Call
Use the Create IVR Call endpoint to create an IVR call entry. This endpoint connects your IVR system to Galileo&#x27;s IVR system for PIN setting.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *IVRApiPostCreateivrcallOpts - Optional Parameters:
     * @param "ApiLogin" (optional.String) - 
     * @param "ApiTransKey" (optional.String) - 
     * @param "ProviderId" (optional.Int32) - 
     * @param "TransactionId" (optional.String) - 
     * @param "AccountNo" (optional.String) - 
     * @param "CallType" (optional.String) - 
     * @param "CallParams" (optional.String) - 
     * @param "Phone" (optional.String) - 
     * @param "ResponseContentType" (optional.String) -  Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header.
@return InlineResponseDefault84
*/

type IVRApiPostCreateivrcallOpts struct {
    ApiLogin optional.String
    ApiTransKey optional.String
    ProviderId optional.Int32
    TransactionId optional.String
    AccountNo optional.String
    CallType optional.String
    CallParams optional.String
    Phone optional.String
    ResponseContentType optional.String
}

func (a *IVRApiService) PostCreateivrcall(ctx context.Context, localVarOptionals *IVRApiPostCreateivrcallOpts) (InlineResponseDefault84, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponseDefault84
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/createIvrCall"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ResponseContentType.IsSet() {
		localVarHeaderParams["response-content-type"] = parameterToString(localVarOptionals.ResponseContentType.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ApiLogin.IsSet() {
		localVarFormParams.Add("apiLogin", parameterToString(localVarOptionals.ApiLogin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ApiTransKey.IsSet() {
		localVarFormParams.Add("apiTransKey", parameterToString(localVarOptionals.ApiTransKey.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProviderId.IsSet() {
		localVarFormParams.Add("providerId", parameterToString(localVarOptionals.ProviderId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransactionId.IsSet() {
		localVarFormParams.Add("transactionId", parameterToString(localVarOptionals.TransactionId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountNo.IsSet() {
		localVarFormParams.Add("accountNo", parameterToString(localVarOptionals.AccountNo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CallType.IsSet() {
		localVarFormParams.Add("callType", parameterToString(localVarOptionals.CallType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CallParams.IsSet() {
		localVarFormParams.Add("callParams", parameterToString(localVarOptionals.CallParams.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Phone.IsSet() {
		localVarFormParams.Add("phone", parameterToString(localVarOptionals.Phone.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v InlineResponseDefault84
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
IVRApiService Get IVR Call Identifier
Use the Get IVR Call Identifier endpoint to retrieve the most recent &#x60;actpeg_id&#x60; associated with an inbound phone number. This identifier links a customer&#x27;s phone number to a call record.  The Get IVR Call Identifier endpoint connects your IVR system to Galileo&#x27;s IVR system for PIN setting.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *IVRApiPostGetivrcallidentifierOpts - Optional Parameters:
     * @param "ApiLogin" (optional.String) - 
     * @param "ApiTransKey" (optional.String) - 
     * @param "ProviderId" (optional.Int32) - 
     * @param "TransactionId" (optional.String) - 
     * @param "PhoneNo" (optional.String) - 
     * @param "ResponseContentType" (optional.String) -  Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header.
@return InlineResponseDefault86
*/

type IVRApiPostGetivrcallidentifierOpts struct {
    ApiLogin optional.String
    ApiTransKey optional.String
    ProviderId optional.Int32
    TransactionId optional.String
    PhoneNo optional.String
    ResponseContentType optional.String
}

func (a *IVRApiService) PostGetivrcallidentifier(ctx context.Context, localVarOptionals *IVRApiPostGetivrcallidentifierOpts) (InlineResponseDefault86, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponseDefault86
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getIvrCallIdentifier"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ResponseContentType.IsSet() {
		localVarHeaderParams["response-content-type"] = parameterToString(localVarOptionals.ResponseContentType.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ApiLogin.IsSet() {
		localVarFormParams.Add("apiLogin", parameterToString(localVarOptionals.ApiLogin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ApiTransKey.IsSet() {
		localVarFormParams.Add("apiTransKey", parameterToString(localVarOptionals.ApiTransKey.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProviderId.IsSet() {
		localVarFormParams.Add("providerId", parameterToString(localVarOptionals.ProviderId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransactionId.IsSet() {
		localVarFormParams.Add("transactionId", parameterToString(localVarOptionals.TransactionId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PhoneNo.IsSet() {
		localVarFormParams.Add("phoneNo", parameterToString(localVarOptionals.PhoneNo.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v InlineResponseDefault86
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
IVRApiService Get IVR Call Status
Use the Get IVR Call Status endpoint to retrieve the status of outbound IVR calls. This endpoint connects your IVR system to Galileo&#x27;s IVR system for PIN setting.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *IVRApiPostGetivrcallstatusOpts - Optional Parameters:
     * @param "ApiLogin" (optional.String) - 
     * @param "ApiTransKey" (optional.String) - 
     * @param "ProviderId" (optional.Int32) - 
     * @param "TransactionId" (optional.String) - 
     * @param "AccountNo" (optional.String) - 
     * @param "CallId" (optional.Int32) - 
     * @param "ResponseContentType" (optional.String) -  Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header.
@return InlineResponseDefault83
*/

type IVRApiPostGetivrcallstatusOpts struct {
    ApiLogin optional.String
    ApiTransKey optional.String
    ProviderId optional.Int32
    TransactionId optional.String
    AccountNo optional.String
    CallId optional.Int32
    ResponseContentType optional.String
}

func (a *IVRApiService) PostGetivrcallstatus(ctx context.Context, localVarOptionals *IVRApiPostGetivrcallstatusOpts) (InlineResponseDefault83, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponseDefault83
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getIvrCallStatus"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ResponseContentType.IsSet() {
		localVarHeaderParams["response-content-type"] = parameterToString(localVarOptionals.ResponseContentType.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ApiLogin.IsSet() {
		localVarFormParams.Add("apiLogin", parameterToString(localVarOptionals.ApiLogin.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ApiTransKey.IsSet() {
		localVarFormParams.Add("apiTransKey", parameterToString(localVarOptionals.ApiTransKey.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProviderId.IsSet() {
		localVarFormParams.Add("providerId", parameterToString(localVarOptionals.ProviderId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TransactionId.IsSet() {
		localVarFormParams.Add("transactionId", parameterToString(localVarOptionals.TransactionId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountNo.IsSet() {
		localVarFormParams.Add("accountNo", parameterToString(localVarOptionals.AccountNo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CallId.IsSet() {
		localVarFormParams.Add("callId", parameterToString(localVarOptionals.CallId.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v InlineResponseDefault83
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
