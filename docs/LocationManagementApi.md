# {{classname}}

All URIs are relative to *api-sandbox.cv.gpsrv.com/intserv/4.0/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCreatelocation**](LocationManagementApi.md#PostCreatelocation) | **Post** /createLocation | Create Location
[**PostCreatelocationfee**](LocationManagementApi.md#PostCreatelocationfee) | **Post** /createLocationFee | Create Location Fee
[**PostGetlocations**](LocationManagementApi.md#PostGetlocations) | **Post** /getLocations | Get Locations
[**PostModifylocation**](LocationManagementApi.md#PostModifylocation) | **Post** /modifyLocation | Modify Location
[**PostModifylocationfee**](LocationManagementApi.md#PostModifylocationfee) | **Post** /modifyLocationFee | Modify Location Fee

# **PostCreatelocation**
> InlineResponseDefault77 PostCreatelocation(ctx, optional)
Create Location

Use the Create Location endpoint to create a new store location, associated with a provider-specific ID, and nest it below a parent location (store). The first location you create will be designated location 0, the top-level location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationManagementApiPostCreatelocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationManagementApiPostCreatelocationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **name** | **optional.**|  | 
 **address1** | **optional.**|  | 
 **address2** | **optional.**|  | 
 **city** | **optional.**|  | 
 **state** | **optional.**|  | 
 **postalCode** | **optional.**|  | 
 **countryCode** | **optional.**|  | 
 **phone** | **optional.**|  | 
 **parentLocation** | **optional.**|  | 
 **parentLocationType** | **optional.**|  | 
 **providerSpecifiedId** | **optional.**|  | 
 **status** | **optional.**|  | 
 **storeType** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault77**](inline_response_default_77.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreatelocationfee**
> InlineResponseDefault81 PostCreatelocationfee(ctx, optional)
Create Location Fee

Use the Create Location Fee endpoint to create a fee to assess at a location. A fee can be assessed as a flat rate or a percentage. When the fee is a percentage, you can specify minimum and maximum amounts. The fee becomes active upon creation and can be deactivated using the Modify Location Fee endpoint.  These are the fee types: * `FBD` &mdash; FBE direct deposit fee * `DEL` &mdash; ACH direct deposit fee

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationManagementApiPostCreatelocationfeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationManagementApiPostCreatelocationfeeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **location** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **feeType** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **amountType** | **optional.**|  | 
 **minFee** | **optional.**|  | 
 **maxFee** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault81**](inline_response_default_81.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetlocations**
> InlineResponseDefault79 PostGetlocations(ctx, optional)
Get Locations

Use the Get Locations endpoint to retrieve a list of locations according to the search criteria. You can search by location ID, location name, and location street address. You can use a partial match for location name and location street address.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationManagementApiPostGetlocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationManagementApiPostGetlocationsOpts struct
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

[**InlineResponseDefault79**](inline_response_default_79.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostModifylocation**
> InlineResponseDefault78 PostModifylocation(ctx, optional)
Modify Location

Use the Modify Location endpoint to update data elements that are related to the location. Pass only the parameters to be modified &mdash; leave the rest blank.  #### Nullifying data elements You can pass `null` for the following parameters to set the value in the database to `null`.  |   |   | |:--|:--| | `name` | `address1` | | `address2` | `city` | | `phone` | `providerSpecifiedId`|  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationManagementApiPostModifylocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationManagementApiPostModifylocationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **name** | **optional.**|  | 
 **address1** | **optional.**|  | 
 **address2** | **optional.**|  | 
 **city** | **optional.**|  | 
 **state** | **optional.**|  | 
 **postalCode** | **optional.**|  | 
 **countryCode** | **optional.**|  | 
 **phone** | **optional.**|  | 
 **location** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **providerSpecifiedId** | **optional.**|  | 
 **status** | **optional.**|  | 
 **storeType** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault78**](inline_response_default_78.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostModifylocationfee**
> InlineResponseDefault82 PostModifylocationfee(ctx, optional)
Modify Location Fee

Use the Modify Location Fee endpoint to modify a fee that was created with the <a href=\"ref:post_createlocationfee\" target=\"_blank\">Create Location Fee</a> endpoint. You cannot change `feeType` (`FBD` or `DEL`) with this endpoint. Instead, you must create a new fee with Create Location Fee and use this endpoint to deactivate the old fee.  For this endpoint, `amount` and `amountType` are not required, but if one of them is populated, then both are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationManagementApiPostModifylocationfeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationManagementApiPostModifylocationfeeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **location** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **feeType** | **optional.**|  | 
 **amount** | **optional.**|  | 
 **amountType** | **optional.**|  | 
 **minFee** | **optional.**|  | 
 **maxFee** | **optional.**|  | 
 **active** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault82**](inline_response_default_82.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

