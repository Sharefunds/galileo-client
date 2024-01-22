# CreateIvrCallBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**CallType** | **string** | The type of call. &#x60;CPIN&#x60; is the only valid value. Pattern: Alphabetic string Example: &#x60;\&quot;CPIN\&quot;&#x60; | [default to CPIN]
**CallParams** | **string** | JSON-encoded array of call parameters. Pattern: JSON string Example: &#x60;\&quot;{\&quot;panLastFour\&quot;:\&quot;1234\&quot;, \&quot;passCode\&quot;:\&quot;0123\&quot;, \&quot;expiryDate\&quot;:\&quot;2011-01-01\&quot;}\&quot;&#x60; | [default to {"panLastFour":"1234", "passCode":"0123", "expiryDate":"2011-01-01"}]
**Phone** | **string** | Phone number where to route outbound calls. This number supersedes any account-related phone numbers. Pattern: Exactly 10 digits, no hyphens or other characters Example: &#x60;\&quot;8013656060\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

