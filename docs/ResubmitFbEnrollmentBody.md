# ResubmitFbEnrollmentBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**FbeStatusId** | **int32** | The federal benefits enrollment ID (&#x60;fbe_status_id&#x60;) as returned by the &lt;a href&#x3D;\&quot;ref:post_createfederalbenefitenrollment\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Create Federal Benefit Enrollment&lt;/a&gt; endpoint. Pattern: Integer Example: &#x60;1&#x60; | [default to 1]
**RecipientType** | **int32** | Specifies whether the owner of the account in &#x60;accountNo&#x60; is a direct recipient or a beneficiary: * &#x60;0&#x60; &amp;mdash; Direct recipient * &#x60;1&#x60; &amp;mdash; Beneficiary  Pattern: &#x60;0&#x60; or &#x60;1&#x60; Example: &#x60;1&#x60; | [default to null]
**FirstName** | **string** | Account holder&#x27;s first name. &lt;a href&#x3D;\&quot;ref:api-reference-special-characters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Special character support&lt;/a&gt;. Pattern: 1&amp;ndash;40 characters: letters (&#x60;A-Z&#x60;, &#x60;a-z&#x60;), spaces, hyphens (&#x60;-&#x60;) and single quotes (&#x60;&#x27;&#x60;) Example: &#x60;\&quot;Ed\&quot;&#x60; | [default to Ed]
**LastName** | **string** | Account holder&#x27;s last name. &lt;a href&#x3D;\&quot;ref:api-reference-special-characters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Special character support&lt;/a&gt;. Pattern: 1&amp;ndash;40 characters: letters (&#x60;A-Z&#x60;, &#x60;a-z&#x60;), spaces, hyphens (&#x60;-&#x60;) and single quotes (&#x60;&#x27;&#x60;) Example: &#x60;\&quot;Harley\&quot;&#x60; | [default to Harley]
**Ssn** | **string** | The Social Security number of the direct recipient of the disbursement. Pattern: 9-digit Social Security number, no hyphens Example: &#x60;\&quot;123456789\&quot;&#x60; | [default to 123456789]
**AgencyType** | **string** | One of several defined agency types. Pattern: See the &lt;a href&#x3D;\&quot;ref:api-reference-federal-benefit-enrollment-agency-types\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Federal Benefit Enrollment Agency Types&lt;/a&gt; table in Enumerations. Example: &#x60;\&quot;SOCIAL SECURITY\&quot;&#x60; | [default to SOCIAL SECURITY]
**AgentId** | **int32** | Identifier of the agent, for use in transaction reporting and tracking. Pattern: Integer Example: &#x60;109405&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

