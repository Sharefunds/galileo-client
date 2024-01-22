# {{classname}}

All URIs are relative to *api-sandbox.cv.gpsrv.com/intserv/4.0/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCreategroup**](CorporateHierarchyApi.md#PostCreategroup) | **Post** /createGroup | Create Group
[**PostDeletegroups**](CorporateHierarchyApi.md#PostDeletegroups) | **Post** /deleteGroups | Delete Groups
[**PostGetaccountgrouprelationships**](CorporateHierarchyApi.md#PostGetaccountgrouprelationships) | **Post** /getAccountGroupRelationships | Get Account Group Relationships
[**PostGetgrouphierarchy**](CorporateHierarchyApi.md#PostGetgrouphierarchy) | **Post** /getGroupHierarchy | Get Group Hierarchy
[**PostGetgroupsinfo**](CorporateHierarchyApi.md#PostGetgroupsinfo) | **Post** /getGroupsInfo | Get Groups Info
[**PostGetrootgroups**](CorporateHierarchyApi.md#PostGetrootgroups) | **Post** /getRootGroups | Get Root Groups
[**PostRemoveaccountgrouprelationship**](CorporateHierarchyApi.md#PostRemoveaccountgrouprelationship) | **Post** /removeAccountGroupRelationship | Remove Account Group Relationship
[**PostSetaccountgrouprelationships**](CorporateHierarchyApi.md#PostSetaccountgrouprelationships) | **Post** /setAccountGroupRelationships | Set Account Group Relationships
[**PostUpdategroup**](CorporateHierarchyApi.md#PostUpdategroup) | **Post** /updateGroup | Update Group

# **PostCreategroup**
> InlineResponseDefault109 PostCreategroup(ctx, optional)
Create Group

Use the Create Group endpoint to create a root group or a non-root group. When creating a root group, populate these parameters to create the business profile: * `businessLegalName` * `doingBusinessAs` * `phoneCountryCode` * `phone` * `primaryContactEmail` * `primaryContactName`  When creating a non-root group, you can populate the business-profile parameters as desired, except for `phone` and `phoneCountryCode`, which are required. All of the name parameters support the Latin-9 (ISO-8859-15) character set. See <a href=\"doc:creating-a-corporate-hierarchy\" target=\"_blank\">Creating a Corporate Hierarchy</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CorporateHierarchyApiPostCreategroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CorporateHierarchyApiPostCreategroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **groupName** | **optional.**|  | 
 **parentGroupId** | **optional.**|  | 
 **maxLevel** | **optional.**|  | 
 **externalId** | **optional.**|  | 
 **doingBusinessAs** | **optional.**|  | 
 **businessLegalName** | **optional.**|  | 
 **phoneCountryCode** | **optional.**|  | 
 **phone** | **optional.**|  | 
 **primaryContactName** | **optional.**|  | 
 **primaryContactEmail** | **optional.**|  | 
 **smartdata** | **optional.**|  | 
 **employerId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault109**](inline_response_default_109.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDeletegroups**
> InlineResponseDefault16 PostDeletegroups(ctx, optional)
Delete Groups

Use the Delete Groups endpoint to delete one or more groups. You cannot delete a group that has subgroups or linked accounts. To delete a group that has subgroups or linked accounts, follow the steps in <a href=\"doc:creating-a-corporate-hierarchy#delete-a-group\" target=\"_blank\">Delete a Group</a> in *Creating a Corporate Hierarchy*.   See <a href=\"doc:creating-a-corporate-hierarchy\" target=\"_blank\">Creating a Corporate Hierarchy</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CorporateHierarchyApiPostDeletegroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CorporateHierarchyApiPostDeletegroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **groupIds** | [**optional.Interface of []string**](string.md)|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetaccountgrouprelationships**
> InlineResponseDefault113 PostGetaccountgrouprelationships(ctx, optional)
Get Account Group Relationships

Use the Get Account Group Relationships endpoint to retrieve all accounts that are in a group, either in a single group or in all associated subgroups. See <a href=\"doc:creating-a-corporate-hierarchy\" target=\"_blank\">Creating a Corporate Hierarchy</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CorporateHierarchyApiPostGetaccountgrouprelationshipsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CorporateHierarchyApiPostGetaccountgrouprelationshipsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **groupId** | **optional.**|  | 
 **includeSubGroups** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault113**](inline_response_default_113.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetgrouphierarchy**
> InlineResponseDefault112 PostGetgrouphierarchy(ctx, optional)
Get Group Hierarchy

Use the Get Group Hierarchy endpoint to retrieve a group's hierarchy, either a root group or a non-root group. This endpoint returns only the groups, not the associated accounts. By default, the endpoint returns all of the subgroups that are downstream of the specified group. Pass a value for `subGroupLevel` to specify how many subgroup levels to retrieve. See <a href=\"doc:creating-a-corporate-hierarchy\" target=\"_blank\">Creating a Corporate Hierarchy</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CorporateHierarchyApiPostGetgrouphierarchyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CorporateHierarchyApiPostGetgrouphierarchyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **groupId** | **optional.**|  | 
 **subGroupLevel** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault112**](inline_response_default_112.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetgroupsinfo**
> InlineResponseDefault111 PostGetgroupsinfo(ctx, optional)
Get Groups Info

Use the Get Groups Info endpoint to retrieve information about one or more groups. See <a href=\"doc:creating-a-corporate-hierarchy\" target=\"_blank\">Creating a Corporate Hierarchy</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CorporateHierarchyApiPostGetgroupsinfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CorporateHierarchyApiPostGetgroupsinfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **groupIds** | [**optional.Interface of []string**](string.md)|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault111**](inline_response_default_111.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetrootgroups**
> InlineResponseDefault110 PostGetrootgroups(ctx, optional)
Get Root Groups

Use the Get Root Groups endpoint to retrieve all of the root groups within a core. See <a href=\"doc:creating-a-corporate-hierarchy\" target=\"_blank\">Creating a Corporate Hierarchy</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CorporateHierarchyApiPostGetrootgroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CorporateHierarchyApiPostGetrootgroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **page** | **optional.**|  | 
 **recordCnt** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault110**](inline_response_default_110.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRemoveaccountgrouprelationship**
> InlineResponseDefault16 PostRemoveaccountgrouprelationship(ctx, optional)
Remove Account Group Relationship

Use the Remove Account Group Relationships endpoint to remove an account or accounts from a group. See <a href=\"doc:creating-a-corporate-hierarchy\" target=\"_blank\">Creating a Corporate Hierarchy</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CorporateHierarchyApiPostRemoveaccountgrouprelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CorporateHierarchyApiPostRemoveaccountgrouprelationshipOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNos** | [**optional.Interface of []string**](string.md)|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSetaccountgrouprelationships**
> InlineResponseDefault16 PostSetaccountgrouprelationships(ctx, optional)
Set Account Group Relationships

Use the Set Account Group Relationships endpoint to associate one or more accounts with a group. This endpoint can update up to 100 accounts in one request. You can also use this endpoint to reassign an account to a different group without first deleting the existing relationship. See <a href=\"doc:creating-a-corporate-hierarchy\" target=\"_blank\">Creating a Corporate Hierarchy</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CorporateHierarchyApiPostSetaccountgrouprelationshipsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CorporateHierarchyApiPostSetaccountgrouprelationshipsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNos** | [**optional.Interface of []string**](string.md)|  | 
 **groupId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdategroup**
> InlineResponseDefault16 PostUpdategroup(ctx, optional)
Update Group

Use the Update Group endpoint to update an existing group, either a root group or a non-root group. Pass only the parameter(s) to update. All of the name parameters support the Latin-9 (ISO-8859-15) character set. You cannot change the parent ID of a group. See <a href=\"doc:creating-a-corporate-hierarchy\" target=\"_blank\">Creating a Corporate Hierarchy</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CorporateHierarchyApiPostUpdategroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CorporateHierarchyApiPostUpdategroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **groupId** | **optional.**|  | 
 **groupName** | **optional.**|  | 
 **maxLevel** | **optional.**|  | 
 **externalId** | **optional.**|  | 
 **doingBusinessAs** | **optional.**|  | 
 **businessLegalName** | **optional.**|  | 
 **phoneCountryCode** | **optional.**|  | 
 **phone** | **optional.**|  | 
 **primaryContactName** | **optional.**|  | 
 **primaryContactEmail** | **optional.**|  | 
 **smartdata** | **optional.**|  | 
 **employerId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

