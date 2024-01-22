# {{classname}}

All URIs are relative to *api-sandbox.cv.gpsrv.com/intserv/4.0/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCreatebulkcardorder**](InstantIssueApi.md#PostCreatebulkcardorder) | **Post** /createBulkCardOrder | Create Bulk Card Order
[**PostGetbulkcardorder**](InstantIssueApi.md#PostGetbulkcardorder) | **Post** /getBulkCardOrder | Get Bulk Card Order
[**PostGetloadlocations**](InstantIssueApi.md#PostGetloadlocations) | **Post** /getLoadLocations | Get Load Locations
[**PostMovecard**](InstantIssueApi.md#PostMovecard) | **Post** /moveCard | Move Card
[**PostMovecardinventory**](InstantIssueApi.md#PostMovecardinventory) | **Post** /moveCardInventory | Move Card Inventory
[**PostVerifyinstantissuecard**](InstantIssueApi.md#PostVerifyinstantissuecard) | **Post** /verifyInstantIssueCard | Verify Instant-Issue Card
[**PostVoidcreateaccount**](InstantIssueApi.md#PostVoidcreateaccount) | **Post** /voidCreateAccount | Void Create Account

# **PostCreatebulkcardorder**
> InlineResponseDefault72 PostCreatebulkcardorder(ctx, optional)
Create Bulk Card Order

Use the Create Bulk Card Order endpoint to initiate a request for a batch of instant-issue cards. Consult the <a href=\"doc:instant-issue-cards\" target=\"_blank\">Instant-Issue Cards</a> guide for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstantIssueApiPostCreatebulkcardorderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstantIssueApiPostCreatebulkcardorderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **numberOfCards** | **optional.**|  | 
 **embossWith** | **optional.**|  | 
 **shipToName** | **optional.**|  | 
 **shipToAddress** | **optional.**|  | 
 **shipToCity** | **optional.**|  | 
 **shipToStateOrProvince** | **optional.**|  | 
 **shipToPostalCode** | **optional.**|  | 
 **location** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault72**](inline_response_default_72.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetbulkcardorder**
> InlineResponseDefault71 PostGetbulkcardorder(ctx, optional)
Get Bulk Card Order

Use Get Bulk Card Order to retrieve the details of an existing bulk card order that was created with the <a href=\"ref:post_createbulkcardorder\" target=\"_blank\">Create Bulk Card Order</a> endpoint. Consult the <a href=\"doc:instant-issue-cards\" target=\"_blank\">Instant-Issue Cards</a> guide for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstantIssueApiPostGetbulkcardorderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstantIssueApiPostGetbulkcardorderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **orderId** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **returnAllCards** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault71**](inline_response_default_71.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetloadlocations**
> InlineResponseDefault80 PostGetloadlocations(ctx, optional)
Get Load Locations

Use the Get Load Locations endpoint to retrieve a list of load locations that are nearest to the specified postal code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstantIssueApiPostGetloadlocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstantIssueApiPostGetloadlocationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **programId** | **optional.**|  | 
 **postalCode** | **optional.**|  | 
 **resultCount** | **optional.**|  | 
 **latLong** | **optional.**|  | 
 **newVersion** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault80**](inline_response_default_80.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMovecard**
> InlineResponseDefault16 PostMovecard(ctx, optional)
Move Card

Use the Move Card endpoint to move the specified instant-issue card to a different location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstantIssueApiPostMovecardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstantIssueApiPostMovecardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **moveToLocation** | **optional.**|  | 
 **moveToLocationType** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMovecardinventory**
> InlineResponseDefault73 PostMovecardinventory(ctx, optional)
Move Card Inventory

Use the Move Card Inventory endpoint to reallocate card inventory to other entities. Specify the type of move to make with the `type` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstantIssueApiPostMovecardinventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstantIssueApiPostMovecardinventoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **location** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **moveToLocation** | **optional.**|  | 
 **moveToLocationType** | **optional.**|  | 
 **type_** | **optional.**|  | 
 **rangeStart** | **optional.**|  | 
 **rangeEnd** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault73**](inline_response_default_73.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVerifyinstantissuecard**
> InlineResponseDefault69 PostVerifyinstantissuecard(ctx, optional)
Verify Instant-Issue Card

Use the Verify Instant-Issue Card endpoint to retrieve data related to the specified instant-issue card. Consult the <a href=\"doc:instant-issue-cards\" target=\"_blank\">Instant-Issue Cards</a> guide for instructions on using this endpoint.[block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"You can receive PCI-sensitive information only if your provider parameters permit it.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstantIssueApiPostVerifyinstantissuecardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstantIssueApiPostVerifyinstantissuecardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **loadType** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault69**](inline_response_default_69.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVoidcreateaccount**
> InlineResponseDefault16 PostVoidcreateaccount(ctx, optional)
Void Create Account

Use the Void Create Account endpoint to cancel an account (move the account into `status: B`, Voided account) that was created by <a href=\"ref:post_createaccount\" target=\"_blank\">Create Account</a>. Pass either the `transactionId` of the original Create Account call or the optional `accountNo` parameter to identify and void the created account. The `accountNo` parameter can be the PRN that was returned by Create Account or a PAN that you can retrieve using <a href=\"ref:post_getaccountcards\" target=\"_blank\">Get Account Cards</a>.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"Use Void Create Account only for instant-issue (prepaid) card accounts. To reverse an account creation for other account types, use the <<glossary:CST>>.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstantIssueApiPostVoidcreateaccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstantIssueApiPostVoidcreateaccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **type_** | **optional.**|  | 
 **verifyOnly** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

