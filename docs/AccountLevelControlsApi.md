# {{classname}}

All URIs are relative to *api-sandbox.cv.gpsrv.com/intserv/4.0/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostDeleteaccountlevelauthcontrol**](AccountLevelControlsApi.md#PostDeleteaccountlevelauthcontrol) | **Post** /deleteAccountLevelAuthControl | Delete Account-Level Auth Control 
[**PostDeleteaccountlevelmcccontrol**](AccountLevelControlsApi.md#PostDeleteaccountlevelmcccontrol) | **Post** /deleteAccountLevelMccControl | Delete Account-Level MCC Control 
[**PostDeleteaccountlevelmerchantcontrol**](AccountLevelControlsApi.md#PostDeleteaccountlevelmerchantcontrol) | **Post** /deleteAccountLevelMerchantControl | Delete Account-Level Merchant Control 
[**PostGetauthcontrol**](AccountLevelControlsApi.md#PostGetauthcontrol) | **Post** /getAuthControl | Get Auth Control
[**PostGetmcccontrols**](AccountLevelControlsApi.md#PostGetmcccontrols) | **Post** /getMccControls | Get MCC Controls
[**PostGetmerchantcontrols**](AccountLevelControlsApi.md#PostGetmerchantcontrols) | **Post** /getMerchantControls | Get Merchant Controls
[**PostSetaccountlevelauthcontrol**](AccountLevelControlsApi.md#PostSetaccountlevelauthcontrol) | **Post** /setAccountLevelAuthControl | Set Account-Level Auth Control 
[**PostSetaccountlevelmcccontrols**](AccountLevelControlsApi.md#PostSetaccountlevelmcccontrols) | **Post** /setAccountLevelMccControls | Set Account-Level MCC Controls
[**PostSetaccountlevelmerchantcontrol**](AccountLevelControlsApi.md#PostSetaccountlevelmerchantcontrol) | **Post** /setAccountLevelMerchantControl | Set Account-Level Merchant Control

# **PostDeleteaccountlevelauthcontrol**
> InlineResponseDefault16 PostDeleteaccountlevelauthcontrol(ctx, optional)
Delete Account-Level Auth Control 

Use the Delete Auth endpoint to delete one Auth Account Level Control. To delete an Auth control, follow the steps in <a href=\"doc:setting-account-level-auth-controls#disable-a-velocity-alc\" target=\"_blank\">Delete an Auth Control</a> in *Disable an account-level MCC control*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountLevelControlsApiPostDeleteaccountlevelauthcontrolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLevelControlsApiPostDeleteaccountlevelauthcontrolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **controlId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDeleteaccountlevelmcccontrol**
> InlineResponseDefault16 PostDeleteaccountlevelmcccontrol(ctx, optional)
Delete Account-Level MCC Control 

Use the Delete MCC endpoint to delete one MCC Account Level Control. To delete a MCC control, follow the steps in <a href=\"doc:setting-account-level-auth-controls#disable-an-account-level-mcc-control\" target=\"_blank\">Delete a MCC</a> in *Disable an account-level MCC control*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountLevelControlsApiPostDeleteaccountlevelmcccontrolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLevelControlsApiPostDeleteaccountlevelmcccontrolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **mccControlStart** | **optional.**|  | 
 **mccControlEnd** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDeleteaccountlevelmerchantcontrol**
> InlineResponseDefault16 PostDeleteaccountlevelmerchantcontrol(ctx, optional)
Delete Account-Level Merchant Control 

Use the Delete Merchant endpoint to delete one Merchant Account Level Control. To delete a Merchant control, follow the steps in <a href=\"doc:setting-account-level-auth-controls#disable-an-account-level-mid-control\" target=\"_blank\">Delete a Merchant</a> in *Disable an account-level Merchant control*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountLevelControlsApiPostDeleteaccountlevelmerchantcontrolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLevelControlsApiPostDeleteaccountlevelmerchantcontrolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **merchantId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetauthcontrol**
> InlineResponseDefault63 PostGetauthcontrol(ctx, optional)
Get Auth Control

Use the Get Auth Control endpoint to retrieve the velocity control details and their control IDs—at either the account level or product level. If you try to retrieve account-level settings that do not exist, this endpoint returns an error. See <a href=\"doc:setting-account-level-auth-controls\" target=\"_blank\">Setting Account-Level Authorization Controls</a> for instructions on using this endpoint.  | Desired Result | `controlId` | `prodId` | `accountNo` | | --- | :---: | :---: | :---: | | Get all product-level controls for a product. | *blank* | Product ID | *blank* | | Get all account-level controls for an account, including current usages and available balances of controls. | *blank* | *blank* | PRN | | Get the product-level control details for a particular control ID. | Control ID | Product ID | *blank* | | Get the account-level control details for a particular control ID, including current usage and available balance of the control. | Control ID | *blank* | PRN | 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountLevelControlsApiPostGetauthcontrolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLevelControlsApiPostGetauthcontrolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **controlId** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault63**](inline_response_default_63.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetmcccontrols**
> InlineResponseDefault64 PostGetmcccontrols(ctx, optional)
Get MCC Controls

Use the Get MCC Controls endpoint to retrieve the <<glossary:MCC>> controls for either a product ID or an account (PRN). See <a href=\"doc:setting-account-level-auth-controls\" target=\"_blank\">Setting Account-Level Authorization Controls</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountLevelControlsApiPostGetmcccontrolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLevelControlsApiPostGetmcccontrolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault64**](inline_response_default_64.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetmerchantcontrols**
> InlineResponseDefault65 PostGetmerchantcontrols(ctx, optional)
Get Merchant Controls

Use the Get Merchant Controls endpoint to retrieve merchant ID (MID) controls for either the product ID or the specified account. See <a href=\"doc:setting-account-level-auth-controls\" target=\"_blank\">Setting Account-Level Authorization Controls</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountLevelControlsApiPostGetmerchantcontrolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLevelControlsApiPostGetmerchantcontrolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault65**](inline_response_default_65.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSetaccountlevelauthcontrol**
> InlineResponseDefault16 PostSetaccountlevelauthcontrol(ctx, optional)
Set Account-Level Auth Control 

Use the Set Account-Level Auth Control endpoint to create or modify an account-level velocity control. Prior to calling this endpoint, call <a href=\"ref:post_getauthcontrol\" target=\"_blank\">Get Auth Control</a> to get the `controlId` for the control to set or modify. See <a href=\"doc:setting-account-level-auth-controls\" target=\"_blank\">Setting Account-Level Authorization Controls</a> for instructions on using this endpoint.  When updating an existing account-level velocity control, follow these rules: * **Positive value** — Set the control to this value * **Blank** — Do not change the value * `Null` — Nullify the existing value 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountLevelControlsApiPostSetaccountlevelauthcontrolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLevelControlsApiPostSetaccountlevelauthcontrolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **controlId** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **transactionCount** | **optional.**|  | 
 **startDate** | **optional.**|  | 
 **endDate** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSetaccountlevelmcccontrols**
> InlineResponseDefault16 PostSetaccountlevelmcccontrols(ctx, optional)
Set Account-Level MCC Controls

Use the Set Account-Level MCC Controls endpoint to create or modify account-level <<glossary:MCC>> controls. See <a href=\"doc:setting-account-level-auth-controls\" target=\"_blank\">Setting Account-Level Authorization Controls</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountLevelControlsApiPostSetaccountlevelmcccontrolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLevelControlsApiPostSetaccountlevelmcccontrolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **allowDeny** | **optional.**|  | 
 **onlineOnly** | **optional.**|  | 
 **startDate** | **optional.**|  | 
 **endDate** | **optional.**|  | 
 **mccControls** | [**optional.Interface of []string**](string.md)|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSetaccountlevelmerchantcontrol**
> InlineResponseDefault16 PostSetaccountlevelmerchantcontrol(ctx, optional)
Set Account-Level Merchant Control

Use the Set Account-Level Merchant Control endpoint to create or modify an account-level merchant ID (MID) control. See <a href=\"doc:setting-account-level-auth-controls\" target=\"_blank\">Setting Account-Level Authorization Controls</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountLevelControlsApiPostSetaccountlevelmerchantcontrolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountLevelControlsApiPostSetaccountlevelmerchantcontrolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **merchantId** | **optional.**|  | 
 **allowDeny** | **optional.**|  | 
 **startDate** | **optional.**|  | 
 **endDate** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

