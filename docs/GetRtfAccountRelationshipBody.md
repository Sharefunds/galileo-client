# GetRtfAccountRelationshipBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; for either an RTF funding or RTF spending account. Do not use the &lt;&lt;glossary:CAD&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt;. Pattern: 12-digit numeric string Example: &#x60;\&quot;344101254935\&quot;&#x60; | [default to 074103447228]
**Page** | **int32** | The number of the page to retrieve. Pattern: Integer value of &#x60;1&#x60; or greater Example: &#x60;3&#x60; | [optional] [default to 1]
**RecordCnt** | **int32** | The maximum number of records per page to be returned. Pattern: Positive integer &#x60;1-99999&#x60; Example: &#x60;100&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

