# Transaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PmtRefNo** | **string** | Galileo-generated account number. Also called the PRN. | [default to null]
**ActId** | **string** | Transaction activity identifier used in Galileo&#x27;s system | [default to null]
**ActType** | **string** | Identifier for the transaction activity type. See the &lt;a href&#x3D;\&quot;ref:api-reference-activity-type\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Activity Type&lt;/a&gt; enumeration. | [default to null]
**PostTs** | [**time.Time**](time.Time.md) | Galileo system timestamp when the transaction posted to the customer account, in Mountain Standard Time (-0700) | [default to null]
**Amt** | **string** | The transaction amount in the currency of the account. A negative amount debits funds from the customer account. | [default to null]
**Details** | **string** | Description provided by the merchant about the transaction (DE043) | [default to null]
**Description** | **string** | Description of the activity type specified in the &#x60;act_type&#x60; field. See the &lt;a href&#x3D;\&quot;ref:api-reference-activity-type\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Activity Type&lt;/a&gt; enumeration. | [default to null]
**SourceId** | **string** | Galileo-generated identifier that maps to the original transaction, such as &#x60;auth_id&#x60;, &#x60;pmt_id&#x60;, or &#x60;adj_id&#x60; | [default to null]
**BalId** | **string** | Balance ID, a Galileo-generated identifier for the account on which the transaction occurs. Maps to &#x60;galileo_account_number&#x60;. | [default to null]
**ProdId** | **string** | Identifier for the product associated with the account | [default to null]
**AuthTs** | [**time.Time**](time.Time.md) | Galileo system timestamp when the transaction was authorized, in Mountain Standard Time (-0700) | [default to null]
**TransCode** | **string** | Reference your program&#x27;s &lt;a href&#x3D;\&quot;ref:api-reference-activity-type\&quot; target&#x3D;\&quot;_blank\&quot;&gt;activity&lt;/a&gt; and &lt;a href&#x3D;\&quot;ref:api-reference-transaction-types\&quot; target&#x3D;\&quot;_blank\&quot;&gt;transaction types&lt;/a&gt; for possible values. | [default to null]
**AchTransactionId** | **string** | Identifier for the ACH transaction, if applicable | [optional] [default to null]
**ExternalTransId** | **string** | Optional identifier for a transaction that you supply. External to the Galileo system. | [optional] [default to null]
**OriginalAuthId** | **string** | The &#x60;auth_id&#x60; of the previous transaction in the sequence, if any. Maps to &#x60;prior_id&#x60;. | [optional] [default to null]
**NetworkCode** | **string** | A Galileo-generated code to identify the subnetwork over which the transaction took place. Maps to &#x60;network_id&#x60;. | [optional] [default to null]
**LocalAmt** | **string** | Amount, in cents, of the authorization request, in the currency at the point of sale. 12-digit number including leading zeros. This amount does not include upcharges or program fees. (DE004) | [optional] [default to null]
**LocalCurrCode** | **string** | Currency code for &#x60;local_amt&#x60; (DE049) | [optional] [default to null]
**SettleAmt** | **string** | The transaction amount in the settlement currency (DE005) | [optional] [default to null]
**SettleCurrCode** | **string** | Currency code for &#x60;settle_amt&#x60; (DE050) | [optional] [default to null]
**BillingAmt** | **string** | The transaction amount in the currency of the account (DE006)  | [optional] [default to null]
**BillingCurrCode** | **string** | Currency code for &#x60;billing_amt&#x60; (DE051) | [optional] [default to null]
**Mcc** | **string** | Category code for the merchant that initiated the transaction (DE018) | [default to null]
**MerchantId** | **string** | Network-assigned identifier for a merchant (DE042) | [optional] [default to null]
**FormattedMerchantDesc** | **string** | The same information as in the &#x60;details&#x60; field, with formatting | [optional] [default to null]
**TerminalId** | **string** | Identifier for the card reader at the point of sale (DE041) | [optional] [default to null]
**CardId** | **string** | A Galileo-generated identifier for a card, which can be used instead of the PAN. Maps to &#x60;cad&#x60;. | [optional] [default to null]
**CreditInd** | **string** | Indicates whether a PIN was input at the point of sale. &#x60;Y&#x60; &#x3D; No PIN was input. &#x60;N&#x60; &#x3D; A PIN was input. &#x60;None&#x60; &#x3D; Not a card transaction. | [optional] [default to null]
**IacTax** | **float32** | Impuesto al Consumo. Colombian consumption tax. Required for the Mastercard Interchange Intracountry Calculation Program. | [optional] [default to null]
**IvaTax** | **float32** | Impuesto al Valor Agregado. Colombian value-added tax. Required for the Mastercard Interchange Intracountry Calculation Program. | [optional] [default to null]
**FundingAccountPrn** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; of the &lt;&lt;glossary:RTF&gt;&gt; funding account | [optional] [default to null]
**SpendingAccountPrn** | **string** | The PRN of the RTF spending account | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

