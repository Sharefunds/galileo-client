# UpdatePayment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN  Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**PmtId** | **int32** | The payment ID (&#x60;payment_trans_id&#x60;) as returned by the &lt;a href&#x3D;\&quot;ref:post_createpayment\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Create Payment&lt;/a&gt; endpoint or &#x60;pmt_id&#x60;  as returned by the &lt;a href&#x3D;\&quot;ref:post_getpaymenthistory\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Get Payment History&lt;/a&gt; endpoint. Pattern: Positive integer Example: &#x60;4234888&#x60; | [default to 4234888]
**HoldDays** | **int32** | Number of days to hold a payment before processing. If set to &#x60;0&#x60;, the payment will be posted the next time the internal payment process runs. Pattern: Integer value of &#x60;0&#x60; or greater Example: &#x60;0&#x60; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

