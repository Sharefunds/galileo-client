# {{classname}}

All URIs are relative to *api-sandbox.cv.gpsrv.com/intserv/4.0/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostActivatecard**](AccountsAndCardsApi.md#PostActivatecard) | **Post** /activateCard | Activate Card
[**PostAddaccount**](AccountsAndCardsApi.md#PostAddaccount) | **Post** /addAccount | Add Account
[**PostAddcard**](AccountsAndCardsApi.md#PostAddcard) | **Post** /addCard | Add Card
[**PostAddcustomernote**](AccountsAndCardsApi.md#PostAddcustomernote) | **Post** /addCustomerNote | Add Customer Note
[**PostChargeoffaccount**](AccountsAndCardsApi.md#PostChargeoffaccount) | **Post** /chargeOffAccount | Charge Off Account
[**PostCommitcardpinchange**](AccountsAndCardsApi.md#PostCommitcardpinchange) | **Post** /commitCardPinChange | Commit Card PIN Change
[**PostCompleteenrollment**](AccountsAndCardsApi.md#PostCompleteenrollment) | **Post** /completeEnrollment | Complete Enrollment
[**PostCreateaccount**](AccountsAndCardsApi.md#PostCreateaccount) | **Post** /createAccount | Create Account
[**PostCreateprovisioningrequest**](AccountsAndCardsApi.md#PostCreateprovisioningrequest) | **Post** /createProvisioningRequest | Create Provisioning Request
[**PostCreatesingleusevirtualcard**](AccountsAndCardsApi.md#PostCreatesingleusevirtualcard) | **Post** /createSingleUseVirtualCard | Create Single-Use Virtual Card
[**PostCreatevirtualcard**](AccountsAndCardsApi.md#PostCreatevirtualcard) | **Post** /createVirtualCard | Create Virtual Card Account
[**PostForcepasscip**](AccountsAndCardsApi.md#PostForcepasscip) | **Post** /forcePassCip | Force Pass CIP
[**PostGetaccesstoken**](AccountsAndCardsApi.md#PostGetaccesstoken) | **Post** /getAccessToken | Get Access Token
[**PostGetaccountbyid**](AccountsAndCardsApi.md#PostGetaccountbyid) | **Post** /getAccountById | Get Account by ID
[**PostGetaccountcards**](AccountsAndCardsApi.md#PostGetaccountcards) | **Post** /getAccountCards | Get Account Cards
[**PostGetaccountfeatures**](AccountsAndCardsApi.md#PostGetaccountfeatures) | **Post** /getAccountFeatures | Get Account Features
[**PostGetbalance**](AccountsAndCardsApi.md#PostGetbalance) | **Post** /getBalance | Get Balance
[**PostGetcard**](AccountsAndCardsApi.md#PostGetcard) | **Post** /getCard | Get Card
[**PostGetcardpinchangekey**](AccountsAndCardsApi.md#PostGetcardpinchangekey) | **Post** /getCardPinChangeKey | Get Card PIN-Change Key
[**PostGetcustomernotehistory**](AccountsAndCardsApi.md#PostGetcustomernotehistory) | **Post** /getCustomerNoteHistory | Get Customer Note History
[**PostGetenrollmentinfo**](AccountsAndCardsApi.md#PostGetenrollmentinfo) | **Post** /getEnrollmentInfo | Get Enrollment Info
[**PostGetinterest**](AccountsAndCardsApi.md#PostGetinterest) | **Post** /getInterest | Get Interest
[**PostGetrelatedaccounts**](AccountsAndCardsApi.md#PostGetrelatedaccounts) | **Post** /getRelatedAccounts | Get Related Accounts
[**PostGetroundupaccounts**](AccountsAndCardsApi.md#PostGetroundupaccounts) | **Post** /getRoundupAccounts | Get Roundup Accounts
[**PostGetrtfaccountrelationship**](AccountsAndCardsApi.md#PostGetrtfaccountrelationship) | **Post** /getRtfAccountRelationship | Get RTF Account Relationship
[**PostGetsavingsinterest**](AccountsAndCardsApi.md#PostGetsavingsinterest) | **Post** /getSavingsInterest | Get Savings Interest
[**PostGetuserdefinedaccountfields**](AccountsAndCardsApi.md#PostGetuserdefinedaccountfields) | **Post** /getUserDefinedAccountFields | Get User-Defined Account Fields
[**PostModifystatus**](AccountsAndCardsApi.md#PostModifystatus) | **Post** /modifyStatus | Modify Status
[**PostRecoverchargedoffaccount**](AccountsAndCardsApi.md#PostRecoverchargedoffaccount) | **Post** /recoverChargedOffAccount | Recover Charged-Off Account
[**PostReissuecard**](AccountsAndCardsApi.md#PostReissuecard) | **Post** /reissueCard | Reissue Card
[**PostResetcardpinfailcount**](AccountsAndCardsApi.md#PostResetcardpinfailcount) | **Post** /resetCardPinFailCount | Reset Card PIN-Fail Count
[**PostRuncip**](AccountsAndCardsApi.md#PostRuncip) | **Post** /runCip | Run CIP
[**PostRunenrollmentcip**](AccountsAndCardsApi.md#PostRunenrollmentcip) | **Post** /runEnrollmentCip | Run Enrollment CIP
[**PostSearchaccounts**](AccountsAndCardsApi.md#PostSearchaccounts) | **Post** /searchAccounts | Search Accounts
[**PostSetaccountfeature**](AccountsAndCardsApi.md#PostSetaccountfeature) | **Post** /setAccountFeature | Set Account Feature
[**PostSetroundupaccounts**](AccountsAndCardsApi.md#PostSetroundupaccounts) | **Post** /setRoundupAccounts | Set Roundup Accounts
[**PostSetuserdefinedaccountfield**](AccountsAndCardsApi.md#PostSetuserdefinedaccountfield) | **Post** /setUserDefinedAccountField | Set User-Defined Account Field
[**PostStartenrollment**](AccountsAndCardsApi.md#PostStartenrollment) | **Post** /startEnrollment | Start Enrollment
[**PostSwitchproduct**](AccountsAndCardsApi.md#PostSwitchproduct) | **Post** /switchProduct | Switch Product
[**PostUpdateaccount**](AccountsAndCardsApi.md#PostUpdateaccount) | **Post** /updateAccount | Update Account
[**PostUpdateenrollment**](AccountsAndCardsApi.md#PostUpdateenrollment) | **Post** /updateEnrollment | Update Enrollment
[**PostVerifyaccount**](AccountsAndCardsApi.md#PostVerifyaccount) | **Post** /verifyAccount | Verify Account
[**PostVerifycardsecuritycode**](AccountsAndCardsApi.md#PostVerifycardsecuritycode) | **Post** /verifyCardSecurityCode | Verify Card Security Code
[**PostVerifyenrollment**](AccountsAndCardsApi.md#PostVerifyenrollment) | **Post** /verifyEnrollment | Verify Enrollment
[**PostVoidaddcard**](AccountsAndCardsApi.md#PostVoidaddcard) | **Post** /voidAddCard | Void Add Card

# **PostActivatecard**
> InlineResponseDefault67 PostActivatecard(ctx, optional)
Activate Card

Use the Activate Card endpoint to activate a physical card that has an emboss record in `status: Y`. Do not use this endpoint to activate a virtual card; instead, use the <a href=\"ref:post_modifystatus\" target=\"_blank\">Modify Status</a> endpoint with `type: 7`.  Consult the <a href=\"doc:activate-card-procedure\" target=\"_blank\">Activating a Card</a> procedure for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostActivatecardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostActivatecardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **cardExpiryDate** | **optional.**|  | 
 **cardSecurityCode** | **optional.**|  | 
 **cardNumberLastFour** | **optional.**|  | 
 **deactivateTemporaryCards** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault67**](inline_response_default_67.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAddaccount**
> InlineResponseDefault7 PostAddaccount(ctx, optional)
Add Account

Use the Add Account endpoint to add a secondary account to an existing customer account. The account types that may be added include savings, overdraft, account with no card, virtual card, or personalized card.  Consult the <a href=\"doc:add-account-procedure\" target=\"_blank\">Adding an Account</a> procedure for instructions on using this endpoint. Also see <a href=\"doc:about-accounts#create-account-vs-add-account\" target=\"_blank\">Create Account vs Add Account</a> in the *About Accounts* guide.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostAddaccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostAddaccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **location** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **sharedBalance** | **optional.**|  | 
 **fundingAccountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault7**](inline_response_default_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAddcard**
> InlineResponseDefault20 PostAddcard(ctx, optional)
Add Card

Use the Add Card endpoint to add a new card (<<glossary:PAN>> and <<glossary:CAD>>) to an existing <<glossary:PRN>>. Cards that are created with this endpoint are active upon creation, and Galileo recommends that you have only one active card per PRN. As needed, ensure that any previous card has been marked lost, stolen, canceled, or any status but `N`.  Do not use this endpoint to create multiple spending cards that share a central balance. For that use case, Galileo recommends <a href=\"doc:real-time-funding\" target=\"_blank\">Real-Time Funding</a> or <a href=\"doc:corporate-credit\" target=\"_blank\">Corporate Credit</a> for large numbers of cards, or for small numbers of cards you can create <a href=\"doc:about-accounts#joint-accounts-and-shared-balances\" target=\"_blank\">one primary and multiple secondary accounts</a> that share a balance.  Use this endpoint in scenarios such as these: - Replacing a virtual card that is lost, stolen, or expired - Adding another virtual card to a PRN - Replacing an instant-issue card at a storefront (for card programs with a retail integration) - Replacing a personalized card 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostAddcardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostAddcardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **newAccountNo** | **optional.**|  | 
 **creditLimit** | **optional.**|  | 
 **singleUse** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault20**](inline_response_default_20.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAddcustomernote**
> InlineResponseDefault52 PostAddcustomernote(ctx, optional)
Add Customer Note

Use the Add Customer Note endpoint to create a note that is accessible in the Galileo <<glossary:CST>>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostAddcustomernoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostAddcustomernoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **note** | **optional.**|  | 
 **sticky** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault52**](inline_response_default_52.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostChargeoffaccount**
> InlineResponseDefault16 PostChargeoffaccount(ctx, optional)
Charge Off Account

Use the Charge Off Account endpoint to charge off the specified account (sweep funds and set balance to `0.00`) and move the account to `status: R` (charged off).  These are the rules for using the `chargeOffDetails` parameter: * `chargeOffDetails` accepts a JSON-formatted list with `chargeOffAmount` and `chargeOffReason` as the two keys. * The only valid values for `chargeOffReason` are the numerals 1-13, as shown in the <a href=\"ref:post_chargeoffaccount#chargeoffreason\" target=\"_blank\">`chargeOffReason`</a> table. * `chargeOffAmount` is a float that must be greater than `0`. If you pass `0` the request will fail. * `chargeOffAmount` is required when `chargeOffReason` does not equal `1`.  Example: `[{\"chargeOffAmount\": 2.6, \"chargeOffReason\": 5}, {\"chargeOffAmount\": 4.75, \"chargeOffReason\": 2}]`  To recover a charged-off account use the <a href=\"ref:post_recoverchargedoffaccount\" target=\"_blank\">Recover Charged-Off Account</a> endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostChargeoffaccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostChargeoffaccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **chargeOffDetails** | **optional.**|  | 
 **closeAssociatedAccounts** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCommitcardpinchange**
> InlineResponseDefault16 PostCommitcardpinchange(ctx, optional)
Commit Card PIN Change

Use the Commit Card PIN Change endpoint to complete a PIN change that was staged by using either the direct POST or direct render PIN-set methods.  Consult the <a href=\"doc:direct-post-pin-set-procedure\" target=\"_blank\">Direct POST PIN-Set Procedure</a> or the <a href=\"doc:direct-render-pin-set-procedure\" target=\"_blank\">Direct Render PIN-Set Procedure</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostCommitcardpinchangeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostCommitcardpinchangeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCompleteenrollment**
> InlineResponseDefault39 PostCompleteenrollment(ctx, optional)
Complete Enrollment

Use the Complete Enrollment endpoint to finalize an incomplete enrollment that was begun by <a href=\"ref:post_startenrollment\" target=\"_blank\">Start Enrollment</a> or <a href=\"ref:post_createaccount\" target=\"_blank\">Create Account</a>. You must pass the same `transactionId` as the original Start Enrollment or Create Account request. Consult <a href=\"doc:digital-first-program#multi-step-onboarding\" target=\"_blank\">Multi-step onboarding and <a href=\"doc:customer-id-verification\" target=\"_blank\">Customer ID Verification</a> for instructions on using this endpoint with Galileo's integrated CIP solution.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"If this endpoint returns a status code that does not match a status code that is specific to this endpoint, it may be an enrollment status code. See <a href=\\\"ref:api-reference-enrollment-statuses\\\" target=\\\"_blank\\\">Enrollment Statuses</a> for more information.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostCompleteenrollmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostCompleteenrollmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **location** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **sharedBalance** | **optional.**|  | 
 **offline** | **optional.**|  | 
 **primaryAccount** | **optional.**|  | 
 **externalAccountId** | **optional.**|  | 
 **embossLine2** | **optional.**|  | 
 **fundingAccountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault39**](inline_response_default_39.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreateaccount**
> InlineResponseDefault4 PostCreateaccount(ctx, optional)
Create Account

Use the Create Account endpoint to create an account for a new customer. You can use this endpoint for personalized, instant issue, and secondary products.  This endpoint runs <<glossary:CIP>> if you are using Galileo's integrated CIP process. In contrast with the <a href=\"ref:post_startenrollment\" target=\"_blank\">Start Enrollment</a> endpoint, Create Account creates a customer record and at the same time creates an account. Depending on product settings, it also creates a card and loads funds onto it.  Consult the <a href=\"doc:create-account-procedure\" target=\"_blank\">Creating an Account</a> procedure for instructions on using this endpoint, and consult the <a href=\"doc:customer-id-verification\" target=\"_blank\">Customer ID Verification</a> guide for how to use this endpoint with Galileo's integrated CIP solution. The instructions include a flowchart to illustrate how Create Account works in the Galileo system. Also see <a href=\"doc:about-accounts#create-account-vs-add-account\" target=\"_blank\">Create Account vs Add Account</a> in the *About Accounts* guide.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"You can receive PCI-sensitive information only if your provider parameters permit it.\" } [/block]  #### Duplicate use of customer ID Galileo can configure your product parameters to allow or disallow the duplicate use of customer IDs such as SSNs across your programs. (If one of your prospective customers already has a product with another Galileo partner, no duplicate is detected.)  #### CIP response If you are using Galileo's integrated <<glossary:CIP>> solution, the Create Account response includes the verdict. Consult <a href=\"doc:customer-id-verification#create-account-and-create-virtual-card-account-process\" target=\"_blank\">Create Account and Create Virtual Account Process</a> in the *Customer ID Verification (CIP/KYC)* guide for more information.  #### Test names in the CV environment In the <<glossary:CV>> environment only, you can use specific names with the Create Account endpoint to trigger different <<glossary:CIP>> responses. For the `firstName`, `middleName`, and `lastName` parameters, use these values: * `John F Smith` &mdash; Triggers a CIP failure (F) * `John R Smith` &mdash; Triggers a CIP refer (R) * `John P Smith` &mdash; Triggers a CIP pass (P)  Do not use these test names in <<glossary:Production>>.  #### Required fields Most input parameters for this endpoint are not required, to accommodate various use cases. If you would like some of the parameters to be required, populate the PINED product parameter with the parameters to be required. For example: `idType,id`.   If this endpoint returns a status code that does not match a status code that is specific to this endpoint, it may be an enrollment status code. See <a href=\"ref:api-reference-enrollment-statuses\" target=\"_blank\">Enrollment Statuses</a> for more information.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostCreateaccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostCreateaccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **idType** | **optional.**|  | 
 **id** | **optional.**|  | 
 **idType2** | **optional.**|  | 
 **id2** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **location** | **optional.**|  | 
 **locale** | **optional.**|  | 
 **externalAccountId** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **middleName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **dateOfBirth** | **optional.**|  | 
 **address1** | **optional.**|  | 
 **address2** | **optional.**|  | 
 **address3** | **optional.**|  | 
 **address4** | **optional.**|  | 
 **address5** | **optional.**|  | 
 **city** | **optional.**|  | 
 **state** | **optional.**|  | 
 **postalCode** | **optional.**|  | 
 **countryCode** | **optional.**|  | 
 **primaryPhone** | **optional.**|  | 
 **otherPhone** | **optional.**|  | 
 **mobilePhone** | **optional.**|  | 
 **mobileCarrierId** | **optional.**|  | 
 **email** | **optional.**|  | 
 **webUid** | **optional.**|  | 
 **webPwd** | **optional.**|  | 
 **secretQuestion** | **optional.**|  | 
 **secretAnswer** | **optional.**|  | 
 **incomeSource** | **optional.**|  | 
 **occupation** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **loadAmount** | **optional.**|  | 
 **loadType** | **optional.**|  | 
 **primaryAccount** | **optional.**|  | 
 **fundingAccountNo** | **optional.**|  | 
 **sharedBalance** | **optional.**|  | 
 **userData** | **optional.**|  | 
 **offline** | **optional.**|  | 
 **verifyOnly** | **optional.**|  | 
 **cipStatus** | **optional.**|  | 
 **embossLine2** | **optional.**|  | 
 **providerAssessedFee** | **optional.**|  | 
 **loadFromAccountNo** | **optional.**|  | 
 **sweepDate** | **optional.**|  | 
 **expressMail** | **optional.**|  | 
 **shipToAddressPermanent** | **optional.**|  | 
 **shipToAddress1** | **optional.**|  | 
 **shipToAddress2** | **optional.**|  | 
 **shipToCity** | **optional.**|  | 
 **shipToState** | **optional.**|  | 
 **shipToPostalCode** | **optional.**|  | 
 **shipToCountryCode** | **optional.**|  | 
 **businessName** | **optional.**|  | 
 **mobilePhoneCountryCode** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault4**](inline_response_default_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreateprovisioningrequest**
> InlineResponseDefault21 PostCreateprovisioningrequest(ctx, optional)
Create Provisioning Request

Use the Create Provisioning Request endpoint to push-provision a virtual card to a mobile wallet. This endpoint is part of a multiple-step integration that must be completed with each mobile wallet partner. Consult the <a href=\"doc:creating-a-provisioning-request\" target=\"_blank\">Creating a Provisioning Request</a> guide for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostCreateprovisioningrequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostCreateprovisioningrequestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **walletProvider** | **optional.**|  | 
 **cert1** | **optional.**|  | 
 **cert2** | **optional.**|  | 
 **nonce** | **optional.**|  | 
 **nonceSignature** | **optional.**|  | 
 **clientWalletAccountID** | **optional.**|  | 
 **clientDeviceID** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault21**](inline_response_default_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreatesingleusevirtualcard**
> InlineResponseDefault22 PostCreatesingleusevirtualcard(ctx, optional)
Create Single-Use Virtual Card

Use the Create Single-Use Virtual Card endpoint to create a single-use virtual card for a BNPL customer. To use this endpoint with a product, the SUVC product parameter must be set to Y, indicating that the product is a single-use virtual card.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostCreatesingleusevirtualcardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostCreatesingleusevirtualcardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **primaryAccountNo** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **creditLimit** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault22**](inline_response_default_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreatevirtualcard**
> InlineResponseDefault42 PostCreatevirtualcard(ctx, optional)
Create Virtual Card Account

Use the Create Virtual Card Account endpoint to create a virtual card account for a new customer. To create any other type of account for a new customer use the <a href=\"ref:post_createaccount\" target=\"_blank\">Create Account</a> endpoint.  Create Virtual Account is similar to Create Account except that it accepts only virtual card product IDs for `prodId`. Consult the <a href=\"doc:create-account-procedure\" target=\"_blank\">Creating an Account</a> procedure for instructions on using this endpoint, and consult the <a href=\"doc:customer-id-verification\" target=\"_blank\">Customer ID Verification</a> guide for how to use this endpoint with Galileo's integrated CIP solution.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"You must be PCI-compliant to use this endpoint. If you are not PCI compliant, use the <a href=\\\"ref:post_createaccount\\\" target=\\\"_blank\\\">Create Account</a> endpoint and specify a virtual card product for `prodId`.\" } [/block]    [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"If this endpoint returns a status code that does not match a status code that is specific to this endpoint, it may be an enrollment status code. See <a href=\\\"ref:api-reference-enrollment-statuses\\\" target=\\\"_blank\\\">Enrollment Statuses</a> for more information.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostCreatevirtualcardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostCreatevirtualcardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **id** | **optional.**|  | 
 **idType** | **optional.**|  | 
 **id2** | **optional.**|  | 
 **idType2** | **optional.**|  | 
 **locale** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **middleName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **dateOfBirth** | **optional.**|  | 
 **address1** | **optional.**|  | 
 **address2** | **optional.**|  | 
 **address3** | **optional.**|  | 
 **address4** | **optional.**|  | 
 **address5** | **optional.**|  | 
 **city** | **optional.**|  | 
 **state** | **optional.**|  | 
 **postalCode** | **optional.**|  | 
 **countryCode** | **optional.**|  | 
 **primaryPhone** | **optional.**|  | 
 **otherPhone** | **optional.**|  | 
 **mobilePhone** | **optional.**|  | 
 **mobileCarrierId** | **optional.**|  | 
 **email** | **optional.**|  | 
 **webUid** | **optional.**|  | 
 **webPwd** | **optional.**|  | 
 **secretQuestion** | **optional.**|  | 
 **secretAnswer** | **optional.**|  | 
 **loadAmount** | **optional.**|  | 
 **loadType** | **optional.**|  | 
 **externalAccountId** | **optional.**|  | 
 **primaryAccount** | **optional.**|  | 
 **sharedBalance** | **optional.**|  | 
 **userData** | **optional.**|  | 
 **verifyOnly** | **optional.**|  | 
 **loadFromAccountNo** | **optional.**|  | 
 **sweepDate** | **optional.**|  | 
 **creditLimit** | **optional.**|  | 
 **singleUse** | **optional.**|  | 
 **businessName** | **optional.**|  | 
 **mobilePhoneCountryCode** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault42**](inline_response_default_42.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostForcepasscip**
> InlineResponseDefault16 PostForcepasscip(ctx, optional)
Force Pass CIP

Use the Force Pass CIP endpoint to move a customer account into `status: N` and `active: Y`. This endpoint is valid only if you are using Galileo's <<glossary:CIP>> and you verify customer documents when the customer fails CIP. Do not use this endpoint if you are using your own CIP provider or if Galileo verifies customer documents.  Consult the <a href=\"doc:create-account-procedure\" target=\"_blank\">Creating an Account</a> procedure for instructions on using this endpoint, and consult the <a href=\"doc:customer-id-verification\" target=\"_blank\">Customer ID Verification</a> guide for how to use this endpoint with Galileo's integrated CIP solution.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostForcepasscipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostForcepasscipOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetaccesstoken**
> InlineResponseDefault19 PostGetaccesstoken(ctx, optional)
Get Access Token

Use the Get Access Token endpoint to retrieve an access token for a card, account, or customer record. The expiry in seconds (default: 300) and usage count (default: 3) for the access token are configurable using the TSECV (seconds) and TUSEC (usage) parameters.  When `type: 0` always use CAD for `accountNo`.  Use this endpoint to send a customer a link to a one-time view of a dynamically generated image of a virtual card via HTTP, for example, or for other purposes as appropriate.  Consult <a href=\"doc:retrieving-card-information#digital-card-images\" target=\"_blank\">Digital card images</a> in the *Retrieving Card Information* guide for more information. For other uses contact Galileo.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetaccesstokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetaccesstokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **type_** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault19**](inline_response_default_19.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetaccountbyid**
> InlineResponseDefault56 PostGetaccountbyid(ctx, optional)
Get Account by ID

Use the Get Account by ID endpoint to retrieve account information by customer IDs (`id` or `id2` parameter) such as SSN or driver license number, depending on which ID types your program supports.  This endpoint returns customer data such as enrollment status and onboarding date as well as a list of all accounts associated with the customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetaccountbyidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetaccountbyidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **id** | **optional.**|  | 
 **idType** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault56**](inline_response_default_56.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetaccountcards**
> InlineResponseDefault5 PostGetaccountcards(ctx, optional)
Get Account Cards

Use the Get Account Cards endpoint to retrieve a customer's profile information along with all accounts and cards that are associated with that customer, regardless of status.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"You can receive PCI-sensitive information only if your provider parameters permit it.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetaccountcardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetaccountcardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **includeRelated** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault5**](inline_response_default_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetaccountfeatures**
> InlineResponseDefault45 PostGetaccountfeatures(ctx, optional)
Get Account Features

Use the Get Account Features endpoint to retrieve the account features that are set for the specified customer account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetaccountfeaturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetaccountfeaturesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault45**](inline_response_default_45.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetbalance**
> InlineResponseDefault1 PostGetbalance(ctx, optional)
Get Balance

Use the Get Balance endpoint to retrieve the specified account balance and its currency code. This endpoint returns the balance for all accounts that share the same balance ID (Galileo account number).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetbalanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetbalanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault1**](inline_response_default_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetcard**
> InlineResponseDefault18 PostGetcard(ctx, optional)
Get Card

Use the Get Card endpoint to retrieve information for the specified card.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"You can receive PCI-sensitive information only if your provider parameters permit it.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetcardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetcardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault18**](inline_response_default_18.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetcardpinchangekey**
> InlineResponseDefault68 PostGetcardpinchangekey(ctx, optional)
Get Card PIN-Change Key

Use the Get Card PIN-Change Key endpoint to retrieve a PIN-change token. Use the token as part of a PIN-change strategy wherein the cardholder submits a PIN change request via a form that either you or Galileo host.  Consult the <a href=\"doc:direct-post-pin-set-procedure\" target=\"_blank\">Direct POST PIN-Set Procedure</a> or the <a href=\"doc:direct-render-pin-set-procedure\" target=\"_blank\">Direct Render PIN-Set Procedure</a> for instructions on using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetcardpinchangekeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetcardpinchangekeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault68**](inline_response_default_68.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetcustomernotehistory**
> InlineResponseDefault48 PostGetcustomernotehistory(ctx, optional)
Get Customer Note History

Use the Get Customer Note History endpoint to retrieve the customer notes that are in the Galileo <<glossary:CST>>. This endpoint requires a date range: 365 days maximum.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetcustomernotehistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetcustomernotehistoryOpts struct
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

[**InlineResponseDefault48**](inline_response_default_48.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetenrollmentinfo**
> InlineResponseDefault36 PostGetenrollmentinfo(ctx, optional)
Get Enrollment Info

Use Get Enrollment Info to get the customer data that was submitted with <a href=\"ref:post_startenrollment\" target=\"_blank\">Start Enrollment</a> or with a <a href=\"ref:post_createaccount\" target=\"_blank\">Create Account</a> call that passed `cipStatus: 1` (send customer profile to <<glossary:CIP>> but do not create account). The endpoint returns customer data and any CIP-related information.  Pass the `transactionId` from the original Start Enrollment or Create Account call. No other parameters except the base form parameters are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetenrollmentinfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetenrollmentinfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault36**](inline_response_default_36.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetinterest**
> InlineResponseDefault30 PostGetinterest(ctx, optional)
Get Interest

Use the Get Interest endpoint to retrieve accrual interest, interest year-to-date, and annual percentage yield for any type of interest bearing account (savings or non-savings). The Get Interest endpoint is the preferred method for retrieving data on interest bearing accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetinterestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetinterestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **interestMonth** | **optional.**|  | 
 **includeRelated** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault30**](inline_response_default_30.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetrelatedaccounts**
> InlineResponseDefault57 PostGetrelatedaccounts(ctx, optional)
Get Related Accounts

Use the Get Related Accounts endpoint to retrieve all accounts that are related to the account specified in `accountNo`. This endpoint returns all accounts with the same owner. If a secondary account is passed, it returns the primary account and other secondaries associated with the same primary.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetrelatedaccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetrelatedaccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault57**](inline_response_default_57.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetroundupaccounts**
> InlineResponseDefault62 PostGetroundupaccounts(ctx, optional)
Get Roundup Accounts

Use the Get Roundup Accounts endpoint to retrieve a record of all accounts linked to a roundup account and the contribution percentage from each account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetroundupaccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetroundupaccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault62**](inline_response_default_62.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetrtfaccountrelationship**
> InlineResponseDefault66 PostGetrtfaccountrelationship(ctx, optional)
Get RTF Account Relationship

Use the Get <<glossary:RTF>> Account Relationship endpoint to retrieve either the RTF spending accounts that are associated with an RTF funding account, or the RTF funding account that is associated with an RTF spending account.  For instructions on using this endpoint see <a href=\"doc:creating-real-time-funding-accounts\" target=\"_blank\">Creating Real-Time Funding Accounts</a>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetrtfaccountrelationshipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetrtfaccountrelationshipOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **page** | **optional.**|  | 
 **recordCnt** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault66**](inline_response_default_66.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetsavingsinterest**
> InlineResponseDefault29 PostGetsavingsinterest(ctx, optional)
Get Savings Interest

Use the Get Savings Interest endpoint to retrieve accrual interest, interest year-to-date, and annual percentage yield earnedfor a savings account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetsavingsinterestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetsavingsinterestOpts struct
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

[**InlineResponseDefault29**](inline_response_default_29.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGetuserdefinedaccountfields**
> InlineResponseDefault55 PostGetuserdefinedaccountfields(ctx, optional)
Get User-Defined Account Fields

Use the Get User-Defined Account Fields endpoint to retrieve the user-defined fields and values that were created with the <a href=\"ref:post_setuserdefinedaccountfield\" target=\"_blank\">Set User-Defined Account Field</a> endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostGetuserdefinedaccountfieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostGetuserdefinedaccountfieldsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault55**](inline_response_default_55.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostModifystatus**
> InlineResponseDefault6 PostModifystatus(ctx, optional)
Modify Status

Use the Modify Status endpoint to change the status of an account, a card, or both. This endpoint does not change any other kind of status.  Consult the <a href=\"ref:api-reference-modify-status-types\" target=\"_blank\">Account and card status table</a> to see what is affected by each `type` value. When activating a card, use this endpoint to activate a virtual card, but use the <a href=\"ref:post_activatecard\" target=\"_blank\">Activate Card</a> endpoint for physical cards.  See the <a href=\"doc:lost-stolen-or-damaged-cards\" target=\"_blank\">Lost, Stolen, or Damaged Cards</a> guide for instructions on using types `3`, `4`, `7`, `8`, and `12`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostModifystatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostModifystatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **type_** | **optional.**|  | 
 **startDate** | **optional.**|  | 
 **endDate** | **optional.**|  | 
 **bypassRepFee** | **optional.**|  | 
 **cardNumberLastFour** | **optional.**|  | 
 **closureReason** | **optional.**|  | 
 **bypassMailFee** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault6**](inline_response_default_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRecoverchargedoffaccount**
> InlineResponseDefault16 PostRecoverchargedoffaccount(ctx, optional)
Recover Charged-Off Account

Use the Recover Charged-Off Account endpoint to change a charged-off account (`status: R`) to another status (active or canceled without refund) and to restore the account balance. Using this endpoint is valid only when there was an outstanding positive balance at the time of charge-off.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostRecoverchargedoffaccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostRecoverchargedoffaccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **chargeOffRecoveryAmt** | **optional.**|  | 
 **newAccountStatus** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostReissuecard**
> InlineResponseDefault70 PostReissuecard(ctx, optional)
Reissue Card

Use the Reissue Card endpoint to reissue a card, which means to send the card to the embosser again. Specify whether to keep the same <<glossary:PAN>> and expiry date as the original card. This endpoint also changes the status of the original card.  Galileo recommends that you use the <<glossary:CAD>> for `accountNo`; however, you can use the <<glossary:PRN>> if only one card has ever been associated with the account.  Galileo recommends that you use this endpoint instead of Modify Status with `type: 12` to reissue cards. For instructions see the <a href=\"doc:reissuing-cards\" target=\"_blank\">Reissuing Cards</a> guide.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"The `expiry_date` and `card_security_code` (CVV) that are returned by the endpoint are not the new values. The new values are generated later by the emboss process. Call <a href=\\\"ref:post_getcard\\\" target=\\\"_blank\\\">Get Card</a> to retrieve the new values after the emboss process has run.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostReissuecardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostReissuecardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **newPan** | **optional.**|  | 
 **newExpiryDate** | **optional.**|  | 
 **emboss** | **optional.**|  | 
 **oldCardStatus** | **optional.**|  | 
 **bypassMailFee** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault70**](inline_response_default_70.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostResetcardpinfailcount**
> InlineResponseDefault16 PostResetcardpinfailcount(ctx, optional)
Reset Card PIN-Fail Count

Use the Reset Card PIN-Fail Count endpoint to move the cardholder's PIN-failure count to zero. Optionally, you can also notate the customer account. You can retrieve the current PIN-fail count and fail date with the <a href=\"ref:post_getcard\" target=\"_blank\">Get Card</a> or <a href=\"ref:post_getaccountcard\" target=\"_blank\">Get Account Card</a> endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostResetcardpinfailcountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostResetcardpinfailcountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **notate** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRuncip**
> InlineResponseDefault41 PostRuncip(ctx, optional)
Run CIP

Use the Run CIP endpoint to run the Galileo <<glossary:CIP>> process on a customer who has already been enrolled. Do not use this endpoint if you are using your own CIP provider.  Consult the <a href=\"doc:customer-id-verification\" target=\"_blank\">Customer ID Verification</a> guide for how to use this endpoint with Galileo's integrated CIP solution.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostRuncipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostRuncipOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault41**](inline_response_default_41.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRunenrollmentcip**
> InlineResponseDefault40 PostRunenrollmentcip(ctx, optional)
Run Enrollment CIP

Use the Run Enrollment CIP endpoint to run or re-run the <<glossary:CIP>> process on an existing customer. Consult the <a href=\"doc:customer-id-verification\" target=\"_blank\">Customer ID Verification</a> guide for instructions on using this endpoint with Galileo's integrated CIP solution.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"Validation on an SSN is limited to a length check.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostRunenrollmentcipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostRunenrollmentcipOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **id** | **optional.**|  | 
 **idType** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault40**](inline_response_default_40.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSearchaccounts**
> InlineResponseDefault58 PostSearchaccounts(ctx, optional)
Search Accounts

Use the Search Accounts endpoint to find a customer account by searching on at least one of the parameters. The data in the response is similar to what the Galileo <<glossary:CST>> search provides.  For the `recordCnt` parameter use these values to get the desired number of records per page: * **No value** &mdash; 50 records per page * **`1` through `100`** &mdash; The specified number of records per page * **Values over `100`** &mdash; 100 records per page 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostSearchaccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostSearchaccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **middleName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **dateOfBirth** | **optional.**|  | 
 **postalCode** | **optional.**|  | 
 **primaryPhone** | **optional.**|  | 
 **otherPhone** | **optional.**|  | 
 **mobilePhone** | **optional.**|  | 
 **email** | **optional.**|  | 
 **userData** | **optional.**|  | 
 **eextid** | **optional.**|  | 
 **recordCnt** | **optional.**|  | 
 **page** | **optional.**|  | 
 **mobilePhoneCountryCode** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault58**](inline_response_default_58.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSetaccountfeature**
> InlineResponseDefault46 PostSetaccountfeature(ctx, optional)
Set Account Feature

Use the Set Account Feature endpoint to set or modify specified attributes of a customer account. * When enabling and disabling fraud rules (`featureType: 14`), keep in mind that if you pass `featureValue: Y` and do not pass `endDate`, the end date is set to 3000-01-01, meaning that fraud rules are suspended indefinitely. To update the timespan for fraud-rule suspension, pass `featureValue: Y` and the new start and/or end date-times. To immediately re-enable fraud rules, pass `featureValue: N`. When `featureValue: N`, the date fields are ignored. * When you change account features 20, 21 or 22, you can arrange with Galileo to receive the <a href=\"ref:api-reference-events-api-account-feature-change\" target=\"_blank\">`ACFC: account_feature_change`</a> event message. [block:callout] { \"type\": \"info\", \"title\": \"Note\", \"body\": \"Feature types 20, 21, and 22 are mutually exclusive. Only one can be set to `Y` at a time.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostSetaccountfeatureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostSetaccountfeatureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **featureType** | **optional.**|  | 
 **featureValue** | **optional.**|  | 
 **startDate** | **optional.**|  | 
 **endDate** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault46**](inline_response_default_46.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSetroundupaccounts**
> InlineResponseDefault62 PostSetroundupaccounts(ctx, optional)
Set Roundup Accounts

Use the Set Roundup Accounts endpoint to specify the roundup account for a consumer account, such as rounding up into a charity account. (To enable roundup for the account, call <a href=\"ref:post_setaccountfeature\" target=\"_blank\">Set Account Feature</a> with `type: 11`.) You can assign multiple roundup accounts to a single consumer account. The total percentage from all linked roundup accounts must equal 100.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostSetroundupaccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostSetroundupaccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **linkedAccounts** | **optional.**|  | 
 **expire** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault62**](inline_response_default_62.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSetuserdefinedaccountfield**
> InlineResponseDefault54 PostSetuserdefinedaccountfield(ctx, optional)
Set User-Defined Account Field

Use the Set User-Defined Account Field endpoint to create or update custom data elements on an account. There is no technical limit on the number of fields that you can add. These fields are visible in the <<glossary:CST>>.  * `fieldId` &mdash; Provide an identifier for the field. This identifier needs to be unique only to the account. * `fieldValue` &mdash; Provide the value for that field.  These fields are visible in the <<glossary:CST>> in the **Account** tab.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostSetuserdefinedaccountfieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostSetuserdefinedaccountfieldOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **fieldId** | **optional.**|  | 
 **fieldValue** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault54**](inline_response_default_54.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostStartenrollment**
> InlineResponseDefault35 PostStartenrollment(ctx, optional)
Start Enrollment

Use the Start Enrollment endpoint to initiate a new enrollment, run <<glossary:CIP>>, and calculate credit limit, in the case of a credit product. This endpoint is different from <a href=\"ref:post_createaccount\" target=\"_blank\">Create Account</a> in that it creates a customer record but does not create an account or card.  Consult the <a href=\"doc:customer-id-verification\" target=\"_blank\">Customer ID Verification</a> guide for instructions on using this endpoint with Galileo's integrated CIP solution.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"If this endpoint returns a status code that does not match a status code that is specific to this endpoint, it may be an enrollment status code. See <a href=\\\"ref:api-reference-enrollment-statuses\\\" target=\\\"_blank\\\">Enrollment Statuses</a> for more information.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostStartenrollmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostStartenrollmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **id** | **optional.**|  | 
 **idType** | **optional.**|  | 
 **id2** | **optional.**|  | 
 **idType2** | **optional.**|  | 
 **location** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **locale** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **middleName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **dateOfBirth** | **optional.**|  | 
 **address1** | **optional.**|  | 
 **address2** | **optional.**|  | 
 **address3** | **optional.**|  | 
 **address4** | **optional.**|  | 
 **address5** | **optional.**|  | 
 **city** | **optional.**|  | 
 **state** | **optional.**|  | 
 **postalCode** | **optional.**|  | 
 **countryCode** | **optional.**|  | 
 **expressMail** | **optional.**|  | 
 **primaryPhone** | **optional.**|  | 
 **otherPhone** | **optional.**|  | 
 **mobilePhone** | **optional.**|  | 
 **mobileCarrierId** | **optional.**|  | 
 **email** | **optional.**|  | 
 **webUid** | **optional.**|  | 
 **webPwd** | **optional.**|  | 
 **secretQuestion** | **optional.**|  | 
 **secretAnswer** | **optional.**|  | 
 **userData** | **optional.**|  | 
 **verifyOnly** | **optional.**|  | 
 **monthlyIncome** | **optional.**|  | 
 **monthlyLiab** | **optional.**|  | 
 **runCip** | **optional.**|  | 
 **businessName** | **optional.**|  | 
 **mobilePhoneCountryCode** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault35**](inline_response_default_35.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSwitchproduct**
> InlineResponseDefault61 PostSwitchproduct(ctx, optional)
Switch Product

Use the Switch Product endpoint to change the product ID on the specified account to a new product ID. To control whether the product switch triggers a card reissue, set `doReissue`.   Galileo recommends that you use the <<glossary:CAD>> for `accountNo`; however, you can use the <<glossary:PRN>> if only one card has ever been associated with the account.  For `prodId` pass the new product ID.   Galileo recommends that you use this endpoint instead of Set Account Feature with `type: 3` to switch products. For instructions see the <a href=\"doc:switching-products\" target=\"_blank\">Switching Products</a> guide.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"The `expiry_date` and `card_security_code` (CVV) that are returned by the endpoint are not the new values. The new values are generated later by the emboss process. Call <a href=\\\"ref:post_getcard\\\" target=\\\"_blank\\\">Get Card</a> to retrieve the new values after the emboss process has run.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostSwitchproductOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostSwitchproductOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **doReissue** | **optional.**|  | 
 **newPan** | **optional.**|  | 
 **newExpiryDate** | **optional.**|  | 
 **emboss** | **optional.**|  | 
 **oldCardStatus** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault61**](inline_response_default_61.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdateaccount**
> InlineResponseDefault8 PostUpdateaccount(ctx, optional)
Update Account

Use the Update Account endpoint to modify customer profile information in an existing customer record. Pass only the parameters to modify &mdash; other parameters should be left blank. Only active accounts (`status: N`) can be updated using this endpoint.  #### Nullifying data elements Pass the string `null` for any of these parameters to update that value to `null`.  |   |   |   |   | |:--|:--|:--|:--| | `firstName` | `middleName` | `lastName` | `address1` | | `city` | `state` | `postalCode` | `primaryPhone` | | `otherPhone` | `mobilePhone` | `shipToAddress1` | `webUid` | | `secretQuestion` | `secretAnswer` | `mailBounced` | `email`\\* | | `embossLine2` |  |  | | \\*Only if the provider parameter AENUL permits it.  [block:callout]  { \"type\": \"warning\", \"title\": \"Warning\", \"body\": \"Exercise caution when nullifying a primary address &mdash; for compliance in the United States you must maintain a physical address as an individual cardholder's primary address.\" } [/block]  #### Updating customer ID When you are using Galileo's integrated CIP process, you can update the primary customer ID (`id` and `idType`) as long as the customer has not passed <<glossary:CIP>>. If the customer has already passed CIP, status code 415-02 is returned. Secondary customer ID (`id2` and `idType2`) can be updated regardless of customer CIP status.  #### Updating the ship-to address When updating the ship-to address for a customer, the elements `shipToAddress1`, `shipToCity`, `shipToState`, and `shipToPostalCode` are required &ndash; `shipToAddress2` is optional. You can retrieve the current customer ship-to address data using the <a href=\"ref:post_getaccountcard\" target=\"_blank\">Get Account Card</a> endpoint.  By default, the embosser sends the card to the primary address. To ship to a different address than the primary, use the ship-to address fields. If you pass `shipToAddressPermanent: 1` the embosser will always send cards to the ship-to address; otherwise, the ship-to address is used only for the current order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostUpdateaccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostUpdateaccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **idType** | **optional.**|  | 
 **id** | **optional.**|  | 
 **idType2** | **optional.**|  | 
 **id2** | **optional.**|  | 
 **locationType** | **optional.**|  | 
 **location** | **optional.**|  | 
 **locale** | **optional.**|  | 
 **externalAccountId** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **middleName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **dateOfBirth** | **optional.**|  | 
 **address1** | **optional.**|  | 
 **address2** | **optional.**|  | 
 **address3** | **optional.**|  | 
 **address4** | **optional.**|  | 
 **address5** | **optional.**|  | 
 **city** | **optional.**|  | 
 **state** | **optional.**|  | 
 **postalCode** | **optional.**|  | 
 **countryCode** | **optional.**|  | 
 **primaryPhone** | **optional.**|  | 
 **otherPhone** | **optional.**|  | 
 **mobilePhone** | **optional.**|  | 
 **mobileCarrierId** | **optional.**|  | 
 **email** | **optional.**|  | 
 **webUid** | **optional.**|  | 
 **webPwd** | **optional.**|  | 
 **secretQuestion** | **optional.**|  | 
 **secretAnswer** | **optional.**|  | 
 **incomeSource** | **optional.**|  | 
 **occupation** | **optional.**|  | 
 **mailBounced** | **optional.**|  | 
 **shipToAddress1** | **optional.**|  | 
 **shipToAddress2** | **optional.**|  | 
 **shipToCity** | **optional.**|  | 
 **shipToState** | **optional.**|  | 
 **shipToPostalCode** | **optional.**|  | 
 **shipToCountryCode** | **optional.**|  | 
 **shipToAddressPermanent** | **optional.**|  | 
 **embossLine2** | **optional.**|  | 
 **businessName** | **optional.**|  | 
 **mobilePhoneCountryCode** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault8**](inline_response_default_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUpdateenrollment**
> InlineResponseDefault37 PostUpdateenrollment(ctx, optional)
Update Enrollment

Use the Update Enrollment endpoint to update customer profile information for an enrollment that was created using <a href=\"ref:post_startenrollment\" target=\"_blank\">Start Enrollment</a>.  [block:callout]  { \"type\": \"info\", \"title\": \"Note\", \"body\": \"Validation on an SSN is limited to a length check.\" } [/block]  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostUpdateenrollmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostUpdateenrollmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **prodId** | **optional.**|  | 
 **id** | **optional.**|  | 
 **idType** | **optional.**|  | 
 **id2** | **optional.**|  | 
 **idType2** | **optional.**|  | 
 **firstName** | **optional.**|  | 
 **middleName** | **optional.**|  | 
 **lastName** | **optional.**|  | 
 **dateOfBirth** | **optional.**|  | 
 **address1** | **optional.**|  | 
 **address2** | **optional.**|  | 
 **address3** | **optional.**|  | 
 **address4** | **optional.**|  | 
 **address5** | **optional.**|  | 
 **locale** | **optional.**|  | 
 **city** | **optional.**|  | 
 **state** | **optional.**|  | 
 **postalCode** | **optional.**|  | 
 **countryCode** | **optional.**|  | 
 **primaryPhone** | **optional.**|  | 
 **otherPhone** | **optional.**|  | 
 **mobilePhone** | **optional.**|  | 
 **mobileCarrierId** | **optional.**|  | 
 **email** | **optional.**|  | 
 **webUid** | **optional.**|  | 
 **webPwd** | **optional.**|  | 
 **secretQuestion** | **optional.**|  | 
 **secretAnswer** | **optional.**|  | 
 **monthlyIncome** | **optional.**|  | 
 **monthlyLiab** | **optional.**|  | 
 **businessName** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault37**](inline_response_default_37.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVerifyaccount**
> InlineResponseDefault47 PostVerifyaccount(ctx, optional)
Verify Account

Use the Verify Account endpoint to return information on an account. For example, you can use this endpoint to validate that you can call <a href=\"ref:post_createpayment\" target=\"_blank\">Create Payment</a> on the account without violating the load amount. Verify Account will also return secondary accounts that share the same balance ID and product ID.  This endpoint returns: * Current account balance &mdash; Both posted and authorized transactions are factored into the balance amount * Current maximum amount that can be loaded into the account, according to product parameters * External account ID * Account PRN * Galileo account number (balance ID) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostVerifyaccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostVerifyaccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **loadType** | **optional.**|  | 
 **includeRelated** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault47**](inline_response_default_47.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVerifycardsecuritycode**
> InlineResponseDefault16 PostVerifycardsecuritycode(ctx, optional)
Verify Card Security Code

Use the Verify Card Security Code endpoint to validate a <<glossary:CVV>> for the specified card that a cardholder inputs. Returns success on a match or failure if it does not match. This endpoint does not return the CVV.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostVerifycardsecuritycodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostVerifycardsecuritycodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **cardNumber** | **optional.**|  | 
 **cardSecurityCode** | **optional.**|  | 
 **cardExpiryDate** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVerifyenrollment**
> InlineResponseDefault38 PostVerifyenrollment(ctx, optional)
Verify Enrollment

Use the Verify Enrollment endpoint to look up an existing enrollment by `transactionId` or `id` to view information and status. The endpoint returns enrollment data, credit limit (if applicable), enrollment date, and `transactionID` of the initial enrollment, if the enrollment was initiated with the <a href=\"ref:post_startenrollment\" target=\"_blank\">Start Enrollment</a> or <a href=\"ref:post_createaccount\" target=\"_blank\">Create Account</a> endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostVerifyenrollmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostVerifyenrollmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **id** | **optional.**|  | 
 **idType** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault38**](inline_response_default_38.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVoidaddcard**
> InlineResponseDefault16 PostVoidaddcard(ctx, optional)
Void Add Card

Use the Void Add Card endpoint to cancel a card that was added to an account using the <a href=\"ref:post_addcard\" target=\"_blank\">Add Card</a> endpoint. For `transactionId`, pass the `transaction_id` from the original Add Card **response** (not request).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsAndCardsApiPostVoidaddcardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsAndCardsApiPostVoidaddcardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiLogin** | **optional.**|  | 
 **apiTransKey** | **optional.**|  | 
 **providerId** | **optional.**|  | 
 **transactionId** | **optional.**|  | 
 **accountNo** | **optional.**|  | 
 **responseContentType** | **optional.**| Use &#x60;xml&#x60; or &#x60;json&#x60; to specify the type of response. The default value is &#x60;xml&#x60;. Use this instead of the standard &#x60;accept&#x60; header. | [default to json]

### Return type

[**InlineResponseDefault16**](inline_response_default_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

