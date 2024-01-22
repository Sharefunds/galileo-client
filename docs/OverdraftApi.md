# {{classname}}

All URIs are relative to *api-sandbox.cv.gpsrv.com/intserv/4.0/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostGetoverdraftbalance**](OverdraftApi.md#PostGetoverdraftbalance) | **Post** /getOverdraftBalance | Get Overdraft Balance

# **PostGetoverdraftbalance**
> InlineResponseDefault34 PostGetoverdraftbalance(ctx, optional)
Get Overdraft Balance

Use the Get Overdraft Balance endpoint to retrieve overdraft-related information such as available overdraft limit, payback amount, and <<glossary:DDA>> balance including overdraft.  For `accountNo` you can use the <<glossary:PRN>> or <<glossary:PAN>> of the core account that is associated with the overdraft account or the PRN of the overdraft account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OverdraftApiPostGetoverdraftbalanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OverdraftApiPostGetoverdraftbalanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault34**](inline_response_default_34.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

