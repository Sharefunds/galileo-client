
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

type FederalBenefitEnrollmentApiService service
/*
FederalBenefitEnrollmentApiService Create Federal Benefit Enrollment
Use the Create Federal Benefit Enrollment endpoint to initiate the process for the specified customer to receive U.S. federal benefit funds via ACH deposit. All enrollments that you initiate with this endpoint are queued into a batch file that is sent regularly (usually daily) to the federal government for processing.  [block:callout]  { \&quot;type\&quot;: \&quot;info\&quot;, \&quot;title\&quot;: \&quot;Note\&quot;, \&quot;body\&quot;: \&quot;This endpoint automatically creates a secondary account for benefits deposits that shares its balance with the primary account.\&quot; } [/block]  
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *FederalBenefitEnrollmentApiPostCreatefbenrollmentOpts - Optional Parameters:
     * @param "ApiLogin" (optional.String) - 
     * @param "ApiTransKey" (optional.String) - 
     * @param "ProviderId" (optional.Int32) - 
     * @param "TransactionId" (optional.String) - 
     * @param "AccountNo" (optional.String) - 
     * @param "Location" (optional.String) - 
     * @param "LocationType" (optional.Int32) - 
     * @param "RecipientType" (optional.Int32) - 
     * @param "FirstName" (optional.String) - 
     * @param "LastName" (optional.String) - 
     * @param "Ssn" (optional.String) - 
     * @param "AgencyType" (optional.String) - 
     * @param "Alerts" (optional.Int32) - 
     * @param "AgentId" (optional.Int32) - 
     * @param "ResponseContentType" (optional.String) -  Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header.
@return InlineResponseDefault43
*/

type FederalBenefitEnrollmentApiPostCreatefbenrollmentOpts struct {
    ApiLogin optional.String
    ApiTransKey optional.String
    ProviderId optional.Int32
    TransactionId optional.String
    AccountNo optional.String
    Location optional.String
    LocationType optional.Int32
    RecipientType optional.Int32
    FirstName optional.String
    LastName optional.String
    Ssn optional.String
    AgencyType optional.String
    Alerts optional.Int32
    AgentId optional.Int32
    ResponseContentType optional.String
}

func (a *FederalBenefitEnrollmentApiService) PostCreatefbenrollment(ctx context.Context, localVarOptionals *FederalBenefitEnrollmentApiPostCreatefbenrollmentOpts) (InlineResponseDefault43, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponseDefault43
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/createFbEnrollment"

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
	if localVarOptionals != nil && localVarOptionals.Location.IsSet() {
		localVarFormParams.Add("location", parameterToString(localVarOptionals.Location.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LocationType.IsSet() {
		localVarFormParams.Add("locationType", parameterToString(localVarOptionals.LocationType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RecipientType.IsSet() {
		localVarFormParams.Add("recipientType", parameterToString(localVarOptionals.RecipientType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FirstName.IsSet() {
		localVarFormParams.Add("firstName", parameterToString(localVarOptionals.FirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LastName.IsSet() {
		localVarFormParams.Add("lastName", parameterToString(localVarOptionals.LastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ssn.IsSet() {
		localVarFormParams.Add("ssn", parameterToString(localVarOptionals.Ssn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AgencyType.IsSet() {
		localVarFormParams.Add("agencyType", parameterToString(localVarOptionals.AgencyType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Alerts.IsSet() {
		localVarFormParams.Add("alerts", parameterToString(localVarOptionals.Alerts.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AgentId.IsSet() {
		localVarFormParams.Add("agentId", parameterToString(localVarOptionals.AgentId.Value(), ""))
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
			var v InlineResponseDefault43
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
FederalBenefitEnrollmentApiService Get Federal Benefit Enrollments
Use the Get Federal Benefit Enrollments endpoint to retrieve information on the federal benefit enrollments of the specified account.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *FederalBenefitEnrollmentApiPostGetfbenrollmentsOpts - Optional Parameters:
     * @param "ApiLogin" (optional.String) - 
     * @param "ApiTransKey" (optional.String) - 
     * @param "ProviderId" (optional.Int32) - 
     * @param "TransactionId" (optional.String) - 
     * @param "AccountNo" (optional.String) - 
     * @param "FbeDda" (optional.String) - 
     * @param "AgencyType" (optional.String) - 
     * @param "Status" (optional.String) - 
     * @param "BeneficiarySsn" (optional.String) - 
     * @param "BeneficiaryFirstName" (optional.String) - 
     * @param "BeneficiaryLastName" (optional.String) - 
     * @param "EnrolledFrom" (optional.String) - 
     * @param "EnrolledTo" (optional.String) - 
     * @param "SubmittedFrom" (optional.String) - 
     * @param "SubmittedTo" (optional.String) - 
     * @param "NotedFrom" (optional.String) - 
     * @param "NotedTo" (optional.String) - 
     * @param "ResponseContentType" (optional.String) -  Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header.
@return InlineResponseDefault44
*/

type FederalBenefitEnrollmentApiPostGetfbenrollmentsOpts struct {
    ApiLogin optional.String
    ApiTransKey optional.String
    ProviderId optional.Int32
    TransactionId optional.String
    AccountNo optional.String
    FbeDda optional.String
    AgencyType optional.String
    Status optional.String
    BeneficiarySsn optional.String
    BeneficiaryFirstName optional.String
    BeneficiaryLastName optional.String
    EnrolledFrom optional.String
    EnrolledTo optional.String
    SubmittedFrom optional.String
    SubmittedTo optional.String
    NotedFrom optional.String
    NotedTo optional.String
    ResponseContentType optional.String
}

func (a *FederalBenefitEnrollmentApiService) PostGetfbenrollments(ctx context.Context, localVarOptionals *FederalBenefitEnrollmentApiPostGetfbenrollmentsOpts) (InlineResponseDefault44, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponseDefault44
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getFbEnrollments"

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
	if localVarOptionals != nil && localVarOptionals.FbeDda.IsSet() {
		localVarFormParams.Add("fbeDda", parameterToString(localVarOptionals.FbeDda.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AgencyType.IsSet() {
		localVarFormParams.Add("agencyType", parameterToString(localVarOptionals.AgencyType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Status.IsSet() {
		localVarFormParams.Add("status", parameterToString(localVarOptionals.Status.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BeneficiarySsn.IsSet() {
		localVarFormParams.Add("beneficiarySsn", parameterToString(localVarOptionals.BeneficiarySsn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BeneficiaryFirstName.IsSet() {
		localVarFormParams.Add("beneficiaryFirstName", parameterToString(localVarOptionals.BeneficiaryFirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BeneficiaryLastName.IsSet() {
		localVarFormParams.Add("beneficiaryLastName", parameterToString(localVarOptionals.BeneficiaryLastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnrolledFrom.IsSet() {
		localVarFormParams.Add("enrolledFrom", parameterToString(localVarOptionals.EnrolledFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnrolledTo.IsSet() {
		localVarFormParams.Add("enrolledTo", parameterToString(localVarOptionals.EnrolledTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SubmittedFrom.IsSet() {
		localVarFormParams.Add("submittedFrom", parameterToString(localVarOptionals.SubmittedFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SubmittedTo.IsSet() {
		localVarFormParams.Add("submittedTo", parameterToString(localVarOptionals.SubmittedTo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NotedFrom.IsSet() {
		localVarFormParams.Add("notedFrom", parameterToString(localVarOptionals.NotedFrom.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NotedTo.IsSet() {
		localVarFormParams.Add("notedTo", parameterToString(localVarOptionals.NotedTo.Value(), ""))
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
			var v InlineResponseDefault44
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
FederalBenefitEnrollmentApiService Resubmit Federal Benefit Enrollment
Use the Resubmit Federal Benefit Enrollment endpoint to submit an amended federal-benefit enrollment record and send it in for processing. You can modify the name and SSN parameters as part of resubmittal.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *FederalBenefitEnrollmentApiPostResubmitfbenrollmentOpts - Optional Parameters:
     * @param "ApiLogin" (optional.String) - 
     * @param "ApiTransKey" (optional.String) - 
     * @param "ProviderId" (optional.Int32) - 
     * @param "TransactionId" (optional.String) - 
     * @param "FbeStatusId" (optional.Int32) - 
     * @param "RecipientType" (optional.Int32) - 
     * @param "FirstName" (optional.String) - 
     * @param "LastName" (optional.String) - 
     * @param "Ssn" (optional.String) - 
     * @param "AgencyType" (optional.String) - 
     * @param "AgentId" (optional.Int32) - 
     * @param "ResponseContentType" (optional.String) -  Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header.
@return InlineResponseDefault16
*/

type FederalBenefitEnrollmentApiPostResubmitfbenrollmentOpts struct {
    ApiLogin optional.String
    ApiTransKey optional.String
    ProviderId optional.Int32
    TransactionId optional.String
    FbeStatusId optional.Int32
    RecipientType optional.Int32
    FirstName optional.String
    LastName optional.String
    Ssn optional.String
    AgencyType optional.String
    AgentId optional.Int32
    ResponseContentType optional.String
}

func (a *FederalBenefitEnrollmentApiService) PostResubmitfbenrollment(ctx context.Context, localVarOptionals *FederalBenefitEnrollmentApiPostResubmitfbenrollmentOpts) (InlineResponseDefault16, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponseDefault16
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/resubmitFbEnrollment"

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
	if localVarOptionals != nil && localVarOptionals.FbeStatusId.IsSet() {
		localVarFormParams.Add("fbeStatusId", parameterToString(localVarOptionals.FbeStatusId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RecipientType.IsSet() {
		localVarFormParams.Add("recipientType", parameterToString(localVarOptionals.RecipientType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FirstName.IsSet() {
		localVarFormParams.Add("firstName", parameterToString(localVarOptionals.FirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LastName.IsSet() {
		localVarFormParams.Add("lastName", parameterToString(localVarOptionals.LastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ssn.IsSet() {
		localVarFormParams.Add("ssn", parameterToString(localVarOptionals.Ssn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AgencyType.IsSet() {
		localVarFormParams.Add("agencyType", parameterToString(localVarOptionals.AgencyType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AgentId.IsSet() {
		localVarFormParams.Add("agentId", parameterToString(localVarOptionals.AgentId.Value(), ""))
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
			var v InlineResponseDefault16
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
FederalBenefitEnrollmentApiService Update Federal Benefit Enrollment
Use the Update Federal Benefits Enrollment endpoint to modify an existing federal benefits enrollment prior to re-submitting it with the &lt;a href&#x3D;\&quot;ref:post_resubmitfbenrollment\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Resubmit Federal Benefit Enrollment&lt;/a&gt; endpoint.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *FederalBenefitEnrollmentApiPostUpdatefbenrollmentOpts - Optional Parameters:
     * @param "ApiLogin" (optional.String) - 
     * @param "ApiTransKey" (optional.String) - 
     * @param "ProviderId" (optional.Int32) - 
     * @param "TransactionId" (optional.String) - 
     * @param "FbeStatusId" (optional.Int32) - 
     * @param "RecipientType" (optional.Int32) - 
     * @param "FirstName" (optional.String) - 
     * @param "LastName" (optional.String) - 
     * @param "Ssn" (optional.String) - 
     * @param "AgencyType" (optional.String) - 
     * @param "Status" (optional.String) - 
     * @param "SignedForm" (optional.String) - 
     * @param "EarlyAccess" (optional.String) - 
     * @param "AgentId" (optional.Int32) - 
     * @param "ResponseContentType" (optional.String) -  Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header.
@return InlineResponseDefault16
*/

type FederalBenefitEnrollmentApiPostUpdatefbenrollmentOpts struct {
    ApiLogin optional.String
    ApiTransKey optional.String
    ProviderId optional.Int32
    TransactionId optional.String
    FbeStatusId optional.Int32
    RecipientType optional.Int32
    FirstName optional.String
    LastName optional.String
    Ssn optional.String
    AgencyType optional.String
    Status optional.String
    SignedForm optional.String
    EarlyAccess optional.String
    AgentId optional.Int32
    ResponseContentType optional.String
}

func (a *FederalBenefitEnrollmentApiService) PostUpdatefbenrollment(ctx context.Context, localVarOptionals *FederalBenefitEnrollmentApiPostUpdatefbenrollmentOpts) (InlineResponseDefault16, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponseDefault16
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/updateFbEnrollment"

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
	if localVarOptionals != nil && localVarOptionals.FbeStatusId.IsSet() {
		localVarFormParams.Add("fbeStatusId", parameterToString(localVarOptionals.FbeStatusId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RecipientType.IsSet() {
		localVarFormParams.Add("recipientType", parameterToString(localVarOptionals.RecipientType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FirstName.IsSet() {
		localVarFormParams.Add("firstName", parameterToString(localVarOptionals.FirstName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LastName.IsSet() {
		localVarFormParams.Add("lastName", parameterToString(localVarOptionals.LastName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ssn.IsSet() {
		localVarFormParams.Add("ssn", parameterToString(localVarOptionals.Ssn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AgencyType.IsSet() {
		localVarFormParams.Add("agencyType", parameterToString(localVarOptionals.AgencyType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Status.IsSet() {
		localVarFormParams.Add("status", parameterToString(localVarOptionals.Status.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SignedForm.IsSet() {
		localVarFormParams.Add("signedForm", parameterToString(localVarOptionals.SignedForm.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EarlyAccess.IsSet() {
		localVarFormParams.Add("earlyAccess", parameterToString(localVarOptionals.EarlyAccess.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AgentId.IsSet() {
		localVarFormParams.Add("agentId", parameterToString(localVarOptionals.AgentId.Value(), ""))
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
			var v InlineResponseDefault16
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