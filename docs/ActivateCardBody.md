# ActivateCardBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt;, &lt;&lt;glossary:PAN&gt;&gt; or &lt;&lt;glossary:CAD&gt;&gt; of the account. For card-specific endpoints such as this one, the CAD is preferred. Do not use the PRN if more than one card has ever been associated with this account. Pattern: PAN, PRN, or CAD Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**CardExpiryDate** | **string** | The expiry date provided by the cardholder. This value is compared to the expiry date in the Galileo system. Pattern: YYYY-MM Example: &#x60;\&quot;2019-01\&quot;&#x60; | [default to 2019-01]
**CardSecurityCode** | **string** | The &lt;&lt;glossary:CVV&gt;&gt; value provided by the cardholder. This value is compared to the CVV in the Galileo system. Pattern: 3 digits Example: &#x60;\&quot;123\&quot;&#x60; | [optional] [default to null]
**CardNumberLastFour** | **string** | If &#x60;accountNo&#x60; contains a PRN, and multiple cards are associated with this account, use this value to ensure that the correct card is selected. Pattern: 4 digits Example: &#x60;\&quot;0789\&quot;&#x60; | [optional] [default to null]
**DeactivateTemporaryCards** | **int32** | If temporary cards were issued, this controls whether to deactivate the temporary cards: * &#x60;1&#x60; &amp;mdash; Deactivate * &#x60;0&#x60; &amp;mdash; Do not deactivate  Pattern: &#x60;1&#x60; or &#x60;0&#x60; Example: &#x60;1&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

