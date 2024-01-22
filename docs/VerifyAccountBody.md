# VerifyAccountBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PRN or PAN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**LoadType** | **string** | Load type (otype) of the load to verify. Use values from your Galileo-supplied list of load types. If no &#x60;loadType&#x60; is specified, the default type &#x60;RL&#x60; is used. Pattern: 2 alphanumeric characters Example: &#x60;\&quot;RL\&quot;&#x60; | [optional] [default to null]
**IncludeRelated** | **int32** | Whether to return data for related accounts. See &lt;a href&#x3D;\&quot;doc:about-accounts#getting-related-account-information\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Getting related account information&lt;/a&gt; for instructions on using this parameter. Default is &#x60;1&#x60;. Pattern: &#x60;0&#x60; or &#x60;1&#x60; Example: &#x60;0&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

