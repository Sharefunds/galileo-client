# CreatePaymentResponseResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldBalance** | **float64** | The balance on the account before the transaction | [default to null]
**NewBalance** | **float64** | The balance on the account after the transaction | [default to null]
**FeeAmount** | **string** | Payment fee amount. May return an integer if the fee amount is 0. | [default to null]
**PaymentTransId** | **int32** | Galileo generated card payment or load transaction integer ID | [default to null]
**TransactionId** | **string** | An identifier that represents a transaction | [default to null]
**HoldId** | **int32** | Identifier for the hold. Use the &lt;a href&#x3D;\&quot;ref:post_expirehold\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Expire Hold&lt;/a&gt; endpoint to expire the hold before the time in &#x60;holdExpirationDateTime&#x60;. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

