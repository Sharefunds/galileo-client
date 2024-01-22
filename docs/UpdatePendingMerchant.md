# UpdatePendingMerchant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**SettleId** | **string** | The &#x60;settle_id&#x60; as returned by &lt;a href&#x3D;\&quot;ref:post_getpendingmerchantcredits\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Get Pending Merchant Credits&lt;/a&gt;. Pattern: &#x60;/^[a-z A-Z]{1}-[0-9]{1,20}$/&#x60; Example: &#x60;\&quot;v-43843747\&quot;&#x60; | [default to v-43843747]
**Type_** | **string** | Type of update to perform on the merchant credit: * &#x60;1&#x60; &amp;mdash; Post * &#x60;2&#x60; &amp;mdash; Post and hold  Pattern: &#x60;1&#x60; or &#x60;2&#x60; Example: &#x60;\&quot;2\&quot;&#x60; | [default to TYPE_.2_]
**ProgramId** | **int32** | Galileo-generated identifier for the program. Pattern: Positive integer Example: &#x60;1032&#x60; | [default to 1032]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

