# GetAuthControlBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; for the account. Do not use the &lt;&lt;glossary:CAD&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt;. Pass this value to retrieve all of the account-level velocity controls for this account. If &#x60;controlId&#x60; is also populated then return only the limits for that &#x60;controlId&#x60;. If this parameter is not populated then &#x60;prodId&#x60; must be populated. Pattern: PRN (12-digit numeric string) Example: &#x60;\&quot;344101254935\&quot;&#x60; | [optional] [default to null]
**ControlId** | **int32** | The Galileo-generated identifier for the auth-level velocity control, if known. When this parameter is populated then either &#x60;prodId&#x60; or &#x60;accountNo&#x60; must be populated. Pattern: Integer, maximum 18 digits Example: 102 | [optional] [default to null]
**ProdId** | **int32** | The Galileo-generated identifier for the product. Pass this value to retrieve all of the product-level velocity controls for this product. If &#x60;controlId&#x60; is also populated then return only the product-level controls for that &#x60;controlId&#x60;. If this parameter is not populated then &#x60;accountNo&#x60; must be populated.  Pattern: Integer Example: &#x60;15293&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

