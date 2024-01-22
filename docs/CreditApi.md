# {{classname}}

All URIs are relative to *api-sandbox.cv.gpsrv.com/intserv/4.0/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostGetcreditlimitchangehistory**](CreditApi.md#PostGetcreditlimitchangehistory) | **Post** /getCreditLimitChangeHistory | Get Credit Limit Change History
[**PostGetcreditsummary**](CreditApi.md#PostGetcreditsummary) | **Post** /getCreditSummary | Get Credit Summary
[**PostSetcreditlimit**](CreditApi.md#PostSetcreditlimit) | **Post** /setCreditLimit | Set Credit Limit

# **PostGetcreditlimitchangehistory**
> InlineResponseDefault74 PostGetcreditlimitchangehistory(ctx, optional)
Get Credit Limit Change History

Use the Get Credit Limit Change History endpoint to retrieve a list of credit-limit changes for the specified account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreditApiPostGetcreditlimitchangehistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreditApiPostGetcreditlimitchangehistoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault74**](inline_response_default_74.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetcreditsummary**
> InlineResponseDefault76 PostGetcreditsummary(ctx, optional)
Get Credit Summary

Use the Get Credit Summary endpoint to retrieve billing information for the specified account, from the most recent calculated bill cycle, and available credit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreditApiPostGetcreditsummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreditApiPostGetcreditsummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault76**](inline_response_default_76.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSetcreditlimit**
> InlineResponseDefault75 PostSetcreditlimit(ctx, optional)
Set Credit Limit

Use the Set Credit Limit endpoint to modify the credit limit for the specified account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreditApiPostSetcreditlimitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreditApiPostSetcreditlimitOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault75**](inline_response_default_75.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

