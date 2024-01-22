# CreateSimulatedCardSettleBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN  Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**Amount** | **float32** | The amount of the transaction in the currency at the point of sale. Pattern: Positive integer or decimal number Example: &#x60;100.00&#x60;, &#x60;100&#x60;, or &#x60;100.73&#x60; | [optional] [default to null]
**Association** | **string** | Specifies which card network to simulate: &#x60;visa&#x60; for Visa or &#x60;star&#x60; for Star or &#x60;mc_auth&#x60; for Mastercard or &#x60;visa_baseii&#x60; for Visa Base II. Pattern: &#x60;visa&#x60; or &#x60;mc_auth&#x60; or &#x60;visa_baseii&#x60; Example: &#x60;\&quot;visa\&quot;&#x60; | [default to ASSOCIATION.VISA]
**MerchantName** | **string** | Simulated merchant name. Pass 40 characters for the complete string or pass up to 22 characters and the simulator will add 18 characters of a simulated address. Pattern: Max 40 alphanumeric characters Example: &#x60;\&quot;Fred Jones Bagels\&quot;&#x60; | [optional] [default to null]
**CurrencyCode** | **string** | The currency code for &#x60;amount&#x60;. Use ISO 4217 numeric currency codes. This parameter is required when using this endpoint outside the United States. Pattern: 3 digits Example: &#x60;\&quot;840\&quot;&#x60; | [optional] [default to null]
**SettleAmount** | **float32** | The amount to be transferred from merchant to issuer in the settlement currency. Required when &#x60;specialFunctions: make_multicurrency&#x60; Pattern: Positive integer or decimal number Example: &#x60;100.00&#x60;, &#x60;100&#x60;, or &#x60;100.73&#x60; | [optional] [default to null]
**SettleCurrencyCode** | **string** | The currency code for &#x60;settleAmount&#x60;. Use ISO 4217 numeric currency codes. Required when &#x60;specialFunctions: make_multicurrency&#x60;. Pattern: 3 digits Example: &#x60;\&quot;124\&quot;&#x60; | [optional] [default to null]
**CardBillingAmount** | **float32** | The amount that will be posted to the cardholder account in the currency of the account. Required when &#x60;specialFunctions: make_multicurrency&#x60;. Pattern: Positive integer or decimal number Example: &#x60;100.00&#x60;, &#x60;100&#x60;, or &#x60;100.73&#x60; | [optional] [default to null]
**CardBillingCurrencyCode** | **string** | The currency code for &#x60;cardBillingAmount&#x60;. Use ISO 4217 numeric currency codes. Required when &#x60;specialFunctions: make_multicurrency&#x60;. Pattern: 3 digits Example: &#x60;484&#x60; | [optional] [default to null]
**AuthId** | **string** | The authorization ID (&#x60;auth_id&#x60;) as returned by the &lt;a href&#x3D;\&quot;ref:post_createsimulatedcardauth\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Create Simulated Card Authorization&lt;/a&gt; endpoint. Pattern: Positive integer Example: &#x60;58344373&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

