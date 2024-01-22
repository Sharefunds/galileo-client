# ChargeOffAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Do not use the PRN if more than one card has ever been associated with this account. Pattern: PAN or PRN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**ChargeOffDetails** | **string** | A JSON object containing a list of charge-off details. See the endpoint description above for details. Pattern: JSON format (list of dictionaries) Example: &#x60;\&quot;[{\&quot;chargeOffAmount\&quot;: 2.6, \&quot;chargeOffReason\&quot;: 5}, {\&quot;chargeOffAmount\&quot;: 4.75, \&quot;chargeOffReason\&quot;: 2}]\&quot;&#x60; | [optional] [default to null]
**CloseAssociatedAccounts** | **int32** | If &#x60;accountNo&#x60; is a secondary account, pass &#x60;1&#x60; to close all accounts with the same balance ID. Pattern: &#x60;0&#x60; or &#x60;1&#x60; Example: &#x60;1&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

