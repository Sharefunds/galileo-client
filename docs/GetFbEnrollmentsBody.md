# GetFbEnrollmentsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PRN or PAN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**FbeDda** | **string** | The PRN of the federal benefits enrollment &lt;&lt;glossary:DDA&gt;&gt;. Pattern: PRN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [optional] [default to null]
**AgencyType** | **string** | The type of agency that is providing the benefit. See the &lt;a href&#x3D;\&quot;ref:api-reference-federal-benefit-enrollment-agency-types\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Federal Benefit Enrollment Agency Types&lt;/a&gt; enumeration for valid values. Pattern: String Example: &#x60;\&quot;SOCIAL SECURITY\&quot;&#x60; | [optional] [default to null]
**Status** | **string** | Specifies the state of the federal benefits enrollment. See the &lt;a href&#x3D;\&quot;ref:api-reference-federal-benefit-enrollment-statuses\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Federal Benefit Enrollment Statuses&lt;/a&gt; enumeration for valid values. Pattern: One character Example: &#x60;\&quot;B\&quot;&#x60; | [optional] [default to null]
**BeneficiarySsn** | **string** | The Social Security number of the beneficiary of the disbursement. Pattern: 9-digit Social Security number, no hyphens Example: &#x60;\&quot;123456789\&quot;&#x60; | [optional] [default to null]
**BeneficiaryFirstName** | **string** | Account holder&#x27;s first name. &lt;a href&#x3D;\&quot;ref:api-reference-special-characters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Special character support&lt;/a&gt;. Pattern: 1&amp;ndash;40 characters: letters (&#x60;A-Z&#x60;, &#x60;a-z&#x60;), spaces, hyphens (&#x60;-&#x60;) and single quotes (&#x60;&#x27;&#x60;) Example: &#x60;\&quot;Ed\&quot;&#x60; | [optional] [default to null]
**BeneficiaryLastName** | **string** | Cardholder&#x27;s last name Account holder&#x27;s last name. &lt;a href&#x3D;\&quot;ref:api-reference-special-characters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Special character support&lt;/a&gt;. Pattern: 1&amp;ndash;40 characters: letters (&#x60;A-Z&#x60;, &#x60;a-z&#x60;), spaces, hyphens (&#x60;-&#x60;) and single quotes (&#x60;&#x27;&#x60;) Example: &#x60;\&quot;Harley\&quot;&#x60; | [optional] [default to null]
**EnrolledFrom** | **string** | Date enrollment started. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2016-01-01\&quot;&#x60; | [optional] [default to null]
**EnrolledTo** | **string** | Date enrolled to. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2016-01-01\&quot;&#x60; | [optional] [default to null]
**SubmittedFrom** | **string** | Date submitted from. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2016-01-01\&quot;&#x60; | [optional] [default to null]
**SubmittedTo** | **string** | Date submitted to. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2016-01-01\&quot;&#x60; | [optional] [default to null]
**NotedFrom** | **string** | Date noted from. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2016-01-01\&quot;&#x60; | [optional] [default to null]
**NotedTo** | **string** | Date noted to. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2016-01-01\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

