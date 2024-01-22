# AssessFee

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN  Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**Type_** | **string** | Three-character fee code. Consult the curated list of fees that Galileo provided you. Do not use the numeric otype. Pattern: 1-4 alphanumeric characters Example: &#x60;\&quot;C2C\&quot;&#x60; | [default to 2]
**TransAmount** | **float32** | The amount of the transaction on which to assess the fee, if the fee is a percentage of the transaction. Pass a currency amount as a whole or decimal amount. Pattern: Positive integer or decimal amount. Example: &#x60;100.00&#x60;, &#x60;100&#x60;  or &#x60;100.73&#x60; | [optional] [default to null]
**Amount** | **float32** | Amount of the fee to assess, if the fee is a flat fee. Pass a currency amount as a whole or decimal amount. Pattern: Positive integer or decimal amount Example: &#x60;100.00&#x60;, &#x60;100&#x60;  or &#x60;100.73&#x60; | [optional] [default to null]
**VerifyOnly** | **int32** | Pass &#x60;1&#x60; to test the validity of the parameter data without committing the information to the Galileo system. Pattern: &#x60;0&#x60; or &#x60;1&#x60; Example: &#x60;0&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

