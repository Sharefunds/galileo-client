# {{classname}}

All URIs are relative to *api-sandbox.cv.gpsrv.com/intserv/4.0/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCreatefbenrollment**](FederalBenefitEnrollmentApi.md#PostCreatefbenrollment) | **Post** /createFbEnrollment | Create Federal Benefit Enrollment
[**PostGetfbenrollments**](FederalBenefitEnrollmentApi.md#PostGetfbenrollments) | **Post** /getFbEnrollments | Get Federal Benefit Enrollments
[**PostResubmitfbenrollment**](FederalBenefitEnrollmentApi.md#PostResubmitfbenrollment) | **Post** /resubmitFbEnrollment | Resubmit Federal Benefit Enrollment
[**PostUpdatefbenrollment**](FederalBenefitEnrollmentApi.md#PostUpdatefbenrollment) | **Post** /updateFbEnrollment | Update Federal Benefit Enrollment

# **PostCreatefbenrollment**
> InlineResponseDefault43 PostCreatefbenrollment(ctx, optional)
Create Federal Benefit Enrollment

Use the Create Federal Benefit Enrollment endpoint to initiate the process for the specified customer to receive U.S. federal benefit funds via ACH deposit. All enrollments that you initiate with this endpoint are queued into a batch file that is sent regularly (usually daily) to the federal government for processing.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"This endpoint automatically creates a secondary account for benefits deposits that shares its balance with the primary account.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FederalBenefitEnrollmentApiPostCreatefbenrollmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederalBenefitEnrollmentApiPostCreatefbenrollmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **location** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **recipientType** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **ssn** | **optional.**|  | 
 **agencyType** | **optional.**|  | 
 **alerts** | **optional.**|  | 
 **agentId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault43**](inline_response_default_43.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetfbenrollments**
> InlineResponseDefault44 PostGetfbenrollments(ctx, optional)
Get Federal Benefit Enrollments

Use the Get Federal Benefit Enrollments endpoint to retrieve information on the federal benefit enrollments of the specified account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FederalBenefitEnrollmentApiPostGetfbenrollmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederalBenefitEnrollmentApiPostGetfbenrollmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **fbeDda** | **optional.**|  | 
 **agencyType** | **optional.**|  | 
 **status** | **optional.**|  | 
 **beneficiarySsn** | **optional.**|  | 
 **beneficiaryFirstName** | **optional.**|  | 
 **beneficiaryLastName** | **optional.**|  | 
 **enrolledFrom** | **optional.**|  | 
 **enrolledTo** | **optional.**|  | 
 **submittedFrom** | **optional.**|  | 
 **submittedTo** | **optional.**|  | 
 **notedFrom** | **optional.**|  | 
 **notedTo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault44**](inline_response_default_44.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostResubmitfbenrollment**
> InlineResponseDefault16 PostResubmitfbenrollment(ctx, optional)
Resubmit Federal Benefit Enrollment

Use the Resubmit Federal Benefit Enrollment endpoint to submit an amended federal-benefit enrollment record and send it in for processing. You can modify the name and SSN parameters as part of resubmittal.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FederalBenefitEnrollmentApiPostResubmitfbenrollmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederalBenefitEnrollmentApiPostResubmitfbenrollmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **fbeStatusId** | **optional.**|  | 
 **recipientType** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **ssn** | **optional.**|  | 
 **agencyType** | **optional.**|  | 
 **agentId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdatefbenrollment**
> InlineResponseDefault16 PostUpdatefbenrollment(ctx, optional)
Update Federal Benefit Enrollment

Use the Update Federal Benefits Enrollment endpoint to modify an existing federal benefits enrollment prior to re-submitting it with the <a href=\"ref:post_resubmitfbenrollment\" target=\"_blank\">Resubmit Federal Benefit Enrollment</a> endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FederalBenefitEnrollmentApiPostUpdatefbenrollmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederalBenefitEnrollmentApiPostUpdatefbenrollmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **fbeStatusId** | **optional.**|  | 
 **recipientType** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **ssn** | **optional.**|  | 
 **agencyType** | **optional.**|  | 
 **status** | **optional.**|  | 
 **signedForm** | **optional.**|  | 
 **earlyAccess** | **optional.**|  | 
 **agentId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

