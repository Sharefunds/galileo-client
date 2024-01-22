# {{classname}}

All URIs are relative to *api-sandbox.cv.gpsrv.com/intserv/4.0/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostGetcallstatus**](MiscellaneousApi.md#PostGetcallstatus) | **Post** /getCallStatus | Get Call Status
[**PostGetingocustomerid**](MiscellaneousApi.md#PostGetingocustomerid) | **Post** /getIngoCustomerId | Get Ingo Customer ID
[**PostGetproductinfo**](MiscellaneousApi.md#PostGetproductinfo) | **Post** /getProductInfo | Get Product Info
[**PostGetpromontorybanklist**](MiscellaneousApi.md#PostGetpromontorybanklist) | **Post** /getPromontoryBankList | Get Promontory Bank List
[**PostPing**](MiscellaneousApi.md#PostPing) | **Post** /ping | Ping

# **PostGetcallstatus**
> InlineResponseDefault85 PostGetcallstatus(ctx, optional)
Get Call Status

Use the Get Call Status endpoint to retrieve the response code of a specified API call.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MiscellaneousApiPostGetcallstatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MiscellaneousApiPostGetcallstatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **requestTransactionId** | **optional.**|  | 
 **requestMethodName** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault85**](inline_response_default_85.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetingocustomerid**
> InlineResponseDefault60 PostGetingocustomerid(ctx, optional)
Get Ingo Customer ID

Use the Get Ingo Customer ID to retrieve customer identity information from third-party provider Ingo Money. Use the Ingo customer ID as needed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MiscellaneousApiPostGetingocustomeridOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MiscellaneousApiPostGetingocustomeridOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **mobileDeviceId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault60**](inline_response_default_60.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetproductinfo**
> InlineResponseDefault87 PostGetproductinfo(ctx, optional)
Get Product Info

Use the Get Product Info endpoint to search existing products and retrieve some of the related <<glossary:Production>> configurations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MiscellaneousApiPostGetproductinfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MiscellaneousApiPostGetproductinfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **searchCriteria** | **optional.**|  | 
 **searchCriteriaType** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault87**](inline_response_default_87.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetpromontorybanklist**
> InlineResponseDefault59 PostGetpromontorybanklist(ctx, optional)
Get Promontory Bank List

Use the Get Promontory Bank List endpoint to retrieve a list of Promontory Bank (now Intrafi) balances. This endpoint is related to Intrafi high yield accounts for sweeps and shows the banks that are included in the sweep. Omit `startDate` and `endDate` parameters to retrieve a list of the most recent bank balances; otherwise, you receive all bank-balance records within the specified date range.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MiscellaneousApiPostGetpromontorybanklistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MiscellaneousApiPostGetpromontorybanklistOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **startDate** | **optional.**|  | 
 **endDate** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault59**](inline_response_default_59.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPing**
> InlineResponseDefault16 PostPing(ctx, optional)
Ping

Use the Ping endpoint to test API connectivity and permissions. Ping does not perform any functions on the Galileo system. Some use cases for Ping are: * Testing access from an automated monitoring service * Integrating with an API call to confirm assumptions and validate access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MiscellaneousApiPostPingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MiscellaneousApiPostPingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

