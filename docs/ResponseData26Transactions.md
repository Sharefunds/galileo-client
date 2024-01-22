# ResponseData26Transactions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSavings** | **bool** | Specifies whether the transaction is made on a savings account. | [default to null]
**DenyCode** | **string** | Two-digit code when an authorization request is denied. | [default to null]
**Disputable** | **bool** | Specifies whether the transaction can be disputed in a chargeback. | [default to null]
**Details** | **string** | Description provided by the merchant about the transaction (DE043). | [default to null]
**ActType** | **string** | Activity type. See the &lt;a href&#x3D;\&quot;ref:api-reference-activity-type\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Activity Type&lt;/a&gt; enumeration. | [default to null]
**ActTypeDescription** | **string** | Description of &#x60;act_type&#x60;. See the &lt;a href&#x3D;\&quot;ref:api-reference-activity-type\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Activity Type&lt;/a&gt; enumeration. | [default to null]
**PostTs** | [**time.Time**](time.Time.md) | The Galileo system date/time when a transaction posted to the customer account, in Mountain Standard Time (-0700). | [default to null]
**Amt** | **float32** | The transaction amount in the currency of the account. A negative amount debits funds from the customer account. | [default to null]
**SourceId** | **string** | A Galileo-generated integer that maps back to the original transaction, such as &#x60;auth_id&#x60;, &#x60;pmt_id&#x60;, or &#x60;adj_id&#x60;. | [default to null]
**Type_** | **string** | Reference your program&#x27;s authorization &lt;a href&#x3D;\&quot;ref:api-reference-transaction-types\&quot; target&#x3D;\&quot;_blank\&quot;&gt;transaction types&lt;/a&gt; for possible values. | [default to null]
**TypeDescription** | **string** | The description of the transaction type. | [default to null]
**TransCode** | **string** | A concatenation of the activity type (&#x60;act_type&#x60;) and transaction type (&#x60;type&#x60;). | [default to null]
**Arn** | **string** | Acquirer reference number. An identifier for the acquiring processor. | [default to null]
**MerchantId** | **string** | Network-assigned identifier for a merchant (DE042). | [default to null]
**ExternalTransId** | **string** | A non-Galileo identifier for a transaction, if any. | [default to null]
**CalculatedBalance** | **float32** | Available balance (including this transaction) at the time the API was called. | [default to null]
**AchTransId** | **string** | A unique identifier for the ACH transaction, if applicable. | [default to null]
**AuthTs** | [**time.Time**](time.Time.md) | The Galileo system date/time a transaction was authorized, in Mountain Standard Time (-0700). | [default to null]
**PriorId** | **string** | The &#x60;auth_id&#x60; of the previous transaction in the sequence, if any.  Maps to &#x60;original_auth_id&#x60;. | [default to null]
**CardId** | **string** | A Galileo-generated identifier for a card, which can be used instead of the PAN. Maps to &#x60;cad&#x60;. | [default to null]
**FormattedMerchantDesc** | **string** | The same information as in the &#x60;details&#x60; field, with formatting. | [default to null]
**NetworkCode** | **string** | A Galileo-generated code to identify the subnetwork over which the transaction took place. Maps to &#x60;network_id&#x60;. | [default to null]
**AuthId** | **string** | A Galileo-generated identifier for an authorization and its settlement. | [default to null]
**LocalAmt** | **float32** | Amount, in cents, of the authorization request, in the currency at the point of sale. 12-digit number including leading zeros. This amount does not include upcharges or program fees. (DE004) | [default to null]
**LocalCurrCode** | **string** | Currency code for &#x60;local_amt&#x60; (DE049). | [default to null]
**SettleAmt** | **float32** | The transaction amount, in the settlement currency (DE005) | [default to null]
**SettleCurrCode** | **string** | Currency code for &#x60;settle_amt&#x60; (DE050). | [default to null]
**BillingAmt** | **float32** | The transaction amount, in the currency of the account (DE006)  | [default to null]
**BillingCurrCode** | **string** | Currency code for &#x60;billing_amt&#x60; (DE051). | [default to null]
**PmtRefNo** | **string** | A Galileo-generated number to identify the customer account. Maps to &#x60;PRN&#x60; and &#x60;prn&#x60;. | [default to null]
**MccCode** | **string** | Category code for the merchant that initiated the transaction (DE018). | [default to null]
**CreditInd** | **string** | Indicates whether a PIN was input at the point of sale. &#x60;Y&#x60; &#x3D; No PIN was input. &#x60;N&#x60; &#x3D; A PIN was input. &#x60;None&#x60; &#x3D; Not a card transaction, or not processed as a card transaction. | [default to null]
**IacTax** | **float32** | Impuesto al Consumo. Colombian consumption tax. Required for the Mastercard Interchange Intracountry Calculation Program. | [optional] [default to null]
**IvaTax** | **float32** | Impuesto al Valor Agregado. Colombian value-added tax. Required for the Mastercard Interchange Intracountry Calculation Program. | [optional] [default to null]
**FundingAccountPrn** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; of the &lt;&lt;glossary:RTF&gt;&gt; funding account | [optional] [default to null]
**SpendingAccountPrn** | **string** | The PRN of the RTF spending account | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

