# ResponseData2Authorizations

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthId** | **string** | A Galileo-generated identifier for an authorization. | [default to null]
**Details** | **string** | Description provided by the merchant about the transaction (DE043) | [default to null]
**DetailsFormatted** | **string** | The same information as in the &#x60;details&#x60; field, with formatting | [default to null]
**Amount** | **string** | The authorization amount, in the currency of the account | [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The Galileo system timestamp for the authorization, in Mountain Standard Time (-0700) | [default to null]
**Type_** | **string** | Reference your program&#x27;s authorization &lt;a href&#x3D;\&quot;ref:api-reference-transaction-types\&quot; target&#x3D;\&quot;_blank\&quot;&gt;transaction types&lt;/a&gt; for possible values. | [default to null]
**Mcc** | **string** | Category code for the merchant that initiated the transaction (DE018) | [default to null]
**MerchantId** | **string** | Network-assigned identifier for a merchant (DE042) | [default to null]
**AcqId** | **string** | The identifier for the acquirer (DE032) | [default to null]
**TerminalId** | **string** | Identifier for the card reader at the point of sale (DE041) | [default to null]
**CanBeExpired** | **string** | Whether to allow an expiration on this authorization: &#x60;1&#x60; &#x3D; Allow, or &#x60;0&#x60;  &#x3D; Do not allow | [default to null]
**OriginalAuthId** | **string** | The &#x60;auth_id&#x60; of the previous transaction in the sequence, if any. Maps to &#x60;prior_id&#x60; in other contexts. | [default to null]
**NetworkCode** | **string** | A Galileo-generated code to identify the network over which the transaction took place. Maps to &#x60;network_id&#x60;. | [default to null]
**LocalAmt** | **string** | Amount in cents of the transaction based on the currency at the point of sale (DE004). 12-digit number including leading zeros. | [default to null]
**LocalCurrCode** | **string** | Currency code for &#x60;local_amt&#x60; (DE049) | [default to null]
**SettleAmt** | **string** | The transaction amount, in cents, in the settlement currency (DE005). 12-digit number including leading zeros. | [default to null]
**SettleCurrCode** | **string** | Currency code for &#x60;settle_amt&#x60; (DE050) | [default to null]
**BillingAmt** | **string** | The billing amount in cents (DE006). 12-digit number including leading zeros. | [default to null]
**BillingCurrCode** | **string** | Currency code for &#x60;billing_amt&#x60; (DE051) | [default to null]
**IacTax** | **float32** | Impuesto al Consumo. Colombian consumption tax. Required for the Mastercard Interchange Intracountry Calculation Program. | [optional] [default to null]
**IvaTax** | **float32** | Impuesto al Valor Agregado. Colombian value-added tax. Required for the Mastercard Interchange Intracountry Calculation Program. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

