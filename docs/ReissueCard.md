# ReissueCard

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt;, &lt;&lt;glossary:PAN&gt;&gt; or &lt;&lt;glossary:CAD&gt;&gt; of the account. Pattern: PAN, PRN, or CAD Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**NewPan** | **string** | Controls whether to generate a new PAN for the reissued card. Default: &#x60;N&#x60;. Pattern: &#x60;Y&#x60; or &#x60;N&#x60; Example: &#x60;\&quot;Y\&quot;&#x60; | [optional] [default to NEW_PAN.N]
**NewExpiryDate** | **string** | Controls whether to generate a new expiration date for the reissued card. Always set this value to &#x60;Y&#x60;. Pattern: &#x60;Y&#x60; Example: &#x60;\&quot;Y\&quot;&#x60; | [optional] [default to null]
**Emboss** | **string** | Controls whether to send the reissued card to the embosser. Because virtual cards cannot be reissued, the value &#x60;N&#x60; is not supported.  Pattern: &#x60;Y&#x60; or blank Example: &#x60;\&quot;Y\&quot;&#x60; | [optional] [default to null]
**OldCardStatus** | **string** | Specifies the status to set for the card that is being replaced. Status is set by the emboss process. Valid values are &#x60;C&#x60;, &#x60;D&#x60; or blank. Default: blank Pattern: &#x60;C&#x60;, &#x60;D&#x60; or blank Example: &#x60;D&#x60; | [optional] [default to null]
**BypassMailFee** | **string** | Controls whether to charge the express mail fee for the reissued card. Default: blank. Pattern: &#x60;Y&#x60; or blank Example: &#x60;Y&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

