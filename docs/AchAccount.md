# AchAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchAccountId** | **string** | A unique ID for an ACH account | [default to null]
**Status** | **string** | The status of the ACH account. See &lt;a href&#x3D;\&quot;ref:api-reference-ach-account-statuses\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ACH Account Statuses&lt;/a&gt;. | [default to null]
**RoutingNo** | **string** | A number used to associate a transaction with an institution | [default to null]
**AccountNo** | **string** | A number use to associate a transaction with an account | [default to null]
**Type_** | **string** | The type of record being displayed | [default to null]
**Name** | **string** | The name of the ACH account being displayed | [default to null]
**CompanyName** | **string** | Name of the company that owns the ACH account. Returned when &#x60;entity_type: C&#x60;. | [default to null]
**EntityType** | **string** | Type of account holder of the ACH account: individual (&#x60;I&#x60;) or company (&#x60;C&#x60;). | [default to null]
**FirstName** | **string** | First name of the ACH account holder. Returned when &#x60;entity_type: I&#x60;. | [default to null]
**LastName** | **string** | Last name of the ACH account holder. Returned when &#x60;entity_type: I&#x60; | [default to null]
**FileFirstName** | **string** | First name of the individual account holder that is used in the outgoing Nacha file. Returned when &#x60;entity_type: I&#x60; | [default to null]
**FileLastName** | **string** | Last name of the individual account holder that is used in the outgoing Nacha file. Returned when &#x60;entity_type: I&#x60; | [default to null]
**FileCompanyName** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

