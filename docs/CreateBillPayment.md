# CreateBillPayment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN  Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**Amount** | **float32** | Currency amount as a whole or decimal amount. Pattern: Positive integer or decimal number Example: &#x60;100.00&#x60;, &#x60;100&#x60;, or &#x60;100.73&#x60; | [default to 25.5]
**BillerId** | **int32** | Identifier for the biller (&#x60;biller_id&#x60;) as returned by the &lt;a href&#x3D;\&quot;ref:post_addrppsbiller\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Add RPPS Biller&lt;/a&gt;, &lt;a href&#x3D;\&quot;ref:post_addpaperbiller\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Add Paper Biller&lt;/a&gt; or &lt;a href&#x3D;\&quot;ref:post_getbillers\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Get Billers&lt;/a&gt; endpoint. Pattern: Positive integer Example: &#x60;37323&#x60; | [default to 37323]
**ProcessDate** | **string** | Date in the future to process the payment. Leave blank to process the payment immediately. Pattern: YYYY-MM-DD Example: &#x60;2025-01-01&#x60; | [optional] [default to null]
**Memo** | **string** | String to print in the memo field of printed paper checks. Pattern: Max 50 alphanumeric characters including punctuation Example: &#x60;For babysitting&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

