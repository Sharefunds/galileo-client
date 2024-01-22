# CreateGroupRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**GroupName** | **string** | Name of the group.  Pattern: 1–40 alphanumeric characters, including spaces, hyphens, and single quotes Example: &#x60;\&quot;CDE Corp\&quot;&#x60; | [default to CDE Corp]
**ParentGroupId** | **string** | The identifier of the next group up in the hierarchy. **Required** when creating a non-root group; not valid when &#x60;maxLevel&#x60; is populated. This value cannot be changed after it is set .Pattern: Max 10 numeric characters Example: &#x60;\&quot;456\&quot;&#x60; | [optional] [default to null]
**MaxLevel** | **int32** | Maximum number of subgroup levels that can be created below this root group. **Required** when creating a root group. Not valid when &#x60;parentGroupId&#x60; is populated. This value cannot be changed to a smaller number after it is set. Pattern: Numerals 0–5 Example: &#x60;3&#x60; | [optional] [default to null]
**ExternalId** | **string** | Provider-specified identifier for the group.  Pattern: Max 30 alphanumeric characters Example: &#x60;\&quot;L0000\&quot;&#x60; | [optional] [default to null]
**DoingBusinessAs** | **string** | Operating name of the company, if different from &#x60;businessLegalName&#x60;. Pattern: 2–150 alphanumeric characters, as shown in &lt;a href&#x3D;\&quot;ref:api-reference-special-characters#businessname-parameter\&quot; target&#x3D;\&quot;_blank\&quot;&gt;&#x60;businessName&#x60; parameter&lt;/a&gt; Example: &#x60;\&quot;SportsBall Enterprises\&quot;&#x60; | [optional] [default to null]
**BusinessLegalName** | **string** | Legal name of the business. Pattern: 2–150 alphanumeric characters, as shown in &lt;a href&#x3D;\&quot;ref:api-reference-special-characters#businessname-parameter\&quot; target&#x3D;\&quot;_blank\&quot;&gt;&#x60;businessName&#x60; parameter&lt;/a&gt; Example: &#x60;\&quot;CDE Corporation, Inc.\&quot;&#x60; | [optional] [default to null]
**PhoneCountryCode** | **string** | Country code prefix for &#x60;phone&#x60;. Pattern: Plus sign (&#x60;+&#x60;) with 1–3 digits Example: &#x60;\&quot;+011\&quot;&#x60; | [optional] [default to null]
**Phone** | **string** | Phone number for the &#x60;primaryContactName&#x60;. Pattern: Phone number for the &#x60;primaryContactName&#x60;. Example: &#x60;\&quot;18015551456\&quot;&#x60; | [optional] [default to null]
**PrimaryContactName** | **string** | The name of the primary point of contact at &#x60;businessLegalName&#x60;.  Pattern: 2–150 Latin-9 characters Example: &#x60;\&quot;Maxina Seward\&quot;&#x60; | [optional] [default to null]
**PrimaryContactEmail** | **string** | Email address of &#x60;primaryContactName&#x60;. Pattern: Email address, 3–63 characters Example: &#x60;\&quot;mseward@cde-corp.com\&quot;&#x60; | [optional] [default to null]
**Smartdata** | **string** | Specifies whether the root group is enrolled for Mastercard Smart Data. Pattern: &#x60;Y&#x60; or &#x60;N&#x60;  Example: &#x60;\&quot;Y\&quot;&#x60; | [optional] [default to null]
**EmployerId** | **string** | Specifies employerId data in group creation or update | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

