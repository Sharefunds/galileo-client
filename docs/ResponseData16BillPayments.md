# ResponseData16BillPayments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PmtRefNo** | **string** | A Galileo generated account number | [default to null]
**BillpayTransactionId** | **string** | An ID assigned to a bill payment transaction | [default to null]
**Amount** | **string** | The amount of the bill payment | [default to null]
**ProcessDate** | [**time.Time**](time.Time.md) | The date the payment was processed | [default to null]
**BillerId** | **string** | Positive integer value of a customer configured biller | [default to null]
**Name** | **string** | The name of the biller | [default to null]
**Nickname** | **string** | A nickname used to identify a biller entity | [default to null]
**Status** | **string** | The status of the payment. See &lt;a href&#x3D;\&quot;ref:api-reference-bill-payment-statuses\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Bill Payment Statuses&lt;/a&gt;. | [default to null]
**Type_** | **string** | &#x60;P&#x60; for paper or &#x60;E&#x60; for electronic | [default to null]
**ExternalTransId** | **string** | The &#x60;transactionId&#x60; from the Create Bill Payment call that created the transaction. | [default to null]
**PrintedDate** | [**time.Time**](time.Time.md) | The date the paper check was mailed or the electronic bill payment was sent to Mastercard. | [default to null]
**ClearedDate** | [**time.Time**](time.Time.md) | The date the paper check cleared the bank. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

