# AddAccountBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PAN&gt;&gt; or &lt;&lt;glossary:PRN&gt;&gt; of the primary account to associate with this secondary account. Pattern: PAN or PRN Example: &#x60;\&quot;112541764367\&quot;&#x60; | [default to 074103447228]
**ProdId** | **int32** | Galileo-generated identifier for a product. If this ID is for a business product then &#x60;businessName&#x60; is **required**. Pattern: Integer Example: &#x60;501&#x60; | [default to 501]
**Location** | **string** | Unique location identifier (&#x60;location&#x60;) as returned by the &lt;a href&#x3D;\&quot;ref:post_createlocation\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Create Location&lt;/a&gt; endpoint.  This value is also returned by the &lt;a href&#x3D;\&quot;ref:post_getlocations\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Get Locations&lt;/a&gt; endpoint depending on the value of &#x60;locationType&#x60; when the location was created: * &#x60;0&#x60; or &#x60;2&#x60; &amp;mdash; Returned in the &#x60;location_id&#x60; field * &#x60;1&#x60; &amp;mdash; Returned in the &#x60;provider_specified_id&#x60; field  Pattern: Integer if &#x60;locationType: 0&#x60; or &#x60;locationType: 2&#x60;; max 15 characters if &#x60;locationType: 1&#x60; Example: &#x60;\&quot;a455-3483\&quot;&#x60; | [optional] [default to null]
**LocationType** | **int32** | Type of ID in &#x60;location&#x60;: * &#x60;0&#x60; &amp;mdash; Galileo location ID * &#x60;1&#x60; &amp;mdash; Partner location ID * &#x60;2&#x60; &amp;mdash; Don&#x27;t validate  Pattern: &#x60;0&#x60; or &#x60;1&#x60; Example: &#x60;1&#x60; | [optional] [default to null]
**SharedBalance** | **int32** | Specifies whether the secondary account will share the balance of the primary account. * &#x60;0&#x60; &amp;mdash; Do not share the balance * &#x60;1&#x60; &amp;mdash; Share the balance  Pattern: One digit Example: &#x60;1&#x60; | [optional] [default to null]
**FundingAccountNo** | **string** | PRN of the &lt;&lt;glossary:RTF&gt;&gt; funding account. **Required** if &#x60;prodId&#x60; is an RTF spending product. Pattern: 12-digit numeric string Example: &#x60;\&quot;143101204201\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

