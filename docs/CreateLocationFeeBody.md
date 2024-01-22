# CreateLocationFeeBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**Location** | **string** | Unique location identifier (&#x60;location&#x60;) as returned by the &lt;a href&#x3D;\&quot;ref:post_createlocation\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Create Location&lt;/a&gt; endpoint.  This value is also returned by the &lt;a href&#x3D;\&quot;ref:post_getlocations\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Get Locations&lt;/a&gt; endpoint depending on the value of &#x60;locationType&#x60; when the location was created: * &#x60;0&#x60; or &#x60;2&#x60; &amp;mdash; Returned in the &#x60;location_id&#x60; field * &#x60;1&#x60; &amp;mdash; Returned in the &#x60;provider_specified_id&#x60; field  Pattern: Integer if &#x60;locationType: 0&#x60; or &#x60;locationType: 2&#x60;; max 15 characters if &#x60;locationType: 1&#x60; Example: &#x60;\&quot;a455-3483\&quot;&#x60; | [default to a455-3483]
**LocationType** | **int32** | Type of ID in &#x60;location&#x60;: * &#x60;0&#x60; &amp;mdash; Galileo location ID * &#x60;1&#x60; &amp;mdash; Partner location ID * &#x60;2&#x60; &amp;mdash; Don&#x27;t validate  Pattern: &#x60;0&#x60;, &#x60;1&#x60;, or &#x60;2&#x60; Example: &#x60;0&#x60; | [default to null]
**FeeType** | **string** | Type of fee to assess at this location. To set up both fee types at the same location, call this endpoint twice: * &#x60;FBD&#x60; &amp;mdash; Federal benefit enrollment direct-deposit fee * &#x60;DEL&#x60; &amp;mdash; ACH direct-deposit fee  Pattern: 3-character code Example: &#x60;\&quot;FBD\&quot;&#x60; | [default to FEE_TYPE.FBD]
**Amount** | **float32** | Specify either a flat amount or a percentage. For a flat amount the valid range is between &#x60;.01&#x60; and &#x60;999&#x60;. For a percentage the valid range is &#x60;.01&#x60; (1%) to &#x60;1&#x60; (100%). When this parameter is populated then &#x60;amountType&#x60; is **required**. Pattern: Positive integer or decimal number Example: &#x60;1.23&#x60; | [default to null]
**AmountType** | **string** | Specify which type of value is in &#x60;amount&#x60;: * &#x60;flat&#x60; &amp;mdash; Flat rate * &#x60;percent&#x60; &amp;mdash; Percentage  Pattern: String Example: &#x60;\&quot;flat\&quot;&#x60; | [default to AMOUNT_TYPE.FLAT]
**MinFee** | **float32** | The minimum amount for the fee, when &#x60;amountType: percent&#x60;. Pattern: Positive integer or decimal number Example: &#x60;1.10&#x60; | [optional] [default to null]
**MaxFee** | **float32** | The maximum amount for the fee, when &#x60;amountType: percent&#x60;. Pattern: Positive integer or decimal number Example: &#x60;1.10&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

