# AddCardBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Do not use the &lt;&lt;glossary:CAD&gt;&gt;. Pattern: PAN or PRN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**ProdId** | **int32** | Product ID of the card to add. Pattern: Integer Example: &#x60;\&quot;501\&quot;&#x60; | [default to null]
**NewAccountNo** | **string** | The &lt;&lt;glossary:PAN&gt;&gt;, &lt;&lt;glossary:PRN&gt;&gt; or package ID of the instant-issue card to add. Pass this parameter only for instant-issue cards. Pattern: PAN or PRN or Package ID Example: &#x60;\&quot;123456789012\&quot;&#x60; | [optional] [default to null]
**CreditLimit** | **float64** | Credit limit for card-based spending control. This is available only for credit programs. Pattern: Monetary amount greater than 0. Example: &#x60;500.25&#x60; | [optional] [default to null]
**SingleUse** | **string** | Spending control to limit the card to one purchase transaction. This is available only for credit programs. Pattern: &#x60;Y&#x60; or &#x60;N&#x60; Example: &#x60;\&quot;Y\&quot;&#x60; | [optional] [default to SINGLE_USE.N]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

