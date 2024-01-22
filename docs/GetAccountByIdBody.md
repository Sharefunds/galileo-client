# GetAccountByIdBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**Id** | **string** | Unique identifier for a Cardholder related to the accepted ID type. Pattern: See &lt;a href&#x3D;\&quot;ref:api-reference-customer-id-types\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Customer ID Types&lt;/a&gt; Example: &#x60;\&quot;123456789\&quot;&#x60; | [default to 123456789]
**IdType** | **int32** | Specifies the type of primary identifier in the &#x60;id&#x60; parameter.  This parameter is **required** when &#x60;id&#x60; is populated. See the **ID** column in the &lt;a href&#x3D;\&quot;ref:api-reference-customer-id-types\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Customer ID Types&lt;/a&gt; table for valid values. Pattern: 1- or 2-digit integer Example: &#x60;2&#x60; | [default to 2]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

