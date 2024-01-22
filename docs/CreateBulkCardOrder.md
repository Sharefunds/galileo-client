# CreateBulkCardOrder

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**ProdId** | **int32** | Galileo-generated identifier for a product. Pattern: Integer Example: &#x60;501&#x60; | [default to 501]
**NumberOfCards** | **int32** | Number of cards to be created. Pattern: Positive integer Example: &#x60;1000&#x60; | [default to 1000]
**EmbossWith** | **string** | Optional string to be printed on the first emboss line of the cards. Pattern: Max 20 alphanumeric characters Example: &#x60;\&quot;Bob&#x27;s cards\&quot;&#x60; | [optional] [default to null]
**ShipToName** | **string** | In-care-of name for the address where the bulk card order is to be shipped. Pattern: Max 40 alphanumeric characters Example: &#x60;\&quot;c\&quot;&#x60; | [default to Bob's cards]
**ShipToAddress** | **string** | Street address where the bulk order is to be shipped. Pattern: Max 40 alphanumeric characters Example: &#x60;\&quot;33 Maple Street\&quot;&#x60; | [default to 33 Maple Street]
**ShipToCity** | **string** | City where the bulk order is to be shipped. Pattern: Max 20 alphanumeric characters Example: &#x60;\&quot;Salt Lake City\&quot;&#x60; | [default to Salt Lake City]
**ShipToStateOrProvince** | **string** | State or province where the bulk order is to be shipped. Pattern: 2-character state or province code Example: &#x60;\&quot;UT\&quot;&#x60; | [default to UT]
**ShipToPostalCode** | **string** | Postal code where the bulk order is to be shipped Pattern: &#x60;12345&#x60;, &#x60;12345-6789&#x60;, or &#x60;K1A-1A1&#x60; Example: &#x60;\&quot;84121\&quot;&#x60; | [default to 84121]
**Location** | **string** | Location identifier as returned by &lt;a href&#x3D;\&quot;ref:post_createlocation\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Create Location&lt;/a&gt; (&#x60;location&#x60;) or &lt;a href&#x3D;\&quot;ref:post_getlocations\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Get Locations&lt;/a&gt; (&#x60;location_id&#x60; or &#x60;provider_specified_id&#x60;). Pattern: Integer if &#x60;location&#x60; or &#x60;location_id&#x60;; max 15 characters if &#x60;provider_specified_id&#x60; Example: &#x60;\&quot;a455-3483\&quot;&#x60; | [default to a455-3483]
**LocationType** | **int32** | Type of ID in &#x60;location&#x60;: * &#x60;0&#x60; &amp;mdash; Galileo location ID * &#x60;1&#x60; &amp;mdash; Partner location ID  Pattern: &#x60;0&#x60; or &#x60;1&#x60; Example: &#x60;1&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

