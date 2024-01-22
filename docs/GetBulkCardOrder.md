# GetBulkCardOrder

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**OrderId** | **int32** | The order ID (&#x60;order_id&#x60;) as returned by the &lt;a href&#x3D;\&quot;ref:post_createbulkcardorder\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Create Bulk Card Order&lt;/a&gt; endpoint. Pattern: Positive integer Example: &#x60;5068&#x60; | [default to 5068]
**LocationType** | **int32** | Type of ID in &#x60;location&#x60;: * &#x60;0&#x60; &amp;mdash; Galileo location ID * &#x60;1&#x60; &amp;mdash; Partner location ID * &#x60;2&#x60; &amp;mdash; Don&#x27;t validate  Pattern: &#x60;0&#x60;, &#x60;1&#x60;, or &#x60;2&#x60; Example: &#x60;1&#x60; | [optional] [default to null]
**ReturnAllCards** | **int32** | Specifies whether to return a list of all cards in the order, in the &#x60;cards&#x60; object: * &#x60;0&#x60; &amp;mdash; Do not return cards * &#x60;1&#x60; &amp;mdash; Return all cards  Pattern: &#x60;0&#x60; or &#x60;1&#x60; Example: &#x60;0&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

