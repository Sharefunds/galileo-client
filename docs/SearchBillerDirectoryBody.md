# SearchBillerDirectoryBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**BillerAccountNo** | **string** | Valid biller account number. As applicable, this value is validated against a biller-supplied mask. Pattern: A Max 60 alphanumeric characters including space and hyphen Example: &#x60;\&quot;3333223323455555\&quot;&#x60; | [optional] [default to null]
**BillerName** | **string** | Pass a complete or partial name of a biller to begin the search. When a single term returns multiple entries, the account holder must select from among them. Pass &#x60;billerState&#x60; and/or &#x60;billerAccountNo&#x60; to filter the results. Pattern: Max 50 alphanumeric characters, no punctuation Example: &#x60;\&quot;Netflix\&quot;&#x60; | [default to Netflix]
**BillerState** | **string** | Biller state or province. Pattern: 2-character state or provincial abbreviation Example: &#x60;\&quot;UT\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

