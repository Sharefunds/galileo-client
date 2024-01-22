# GetBalanceResponseResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **float32** | For debit accounts, this is the available balance to spend, meaning the total of all posted transactions with all pending transactions subtracted out. For credit accounts this is the unpaid balance. | [default to null]
**BalanceWithoutPending** | **string** | For debit accounts, the total of all posted transactions with only pending authorizations subtracted out. For credit accounts this is the same as &#x60;balance&#x60;.  | [default to null]
**CurrencyCode** | **string** | A three-digit ISO 4217 code for the currency of the account. All amounts in this response are in this currency. | [default to null]
**PendingAdjustments** | **float32** | Total of all adjustments created by a batch process. This value will almost always be &#x60;0&#x60;. | [default to null]
**PendingBillpay** | **float32** | The total amount of pending bill payments. A bill payment is pending between the time it is created and when the amount is adjusted out of the account. | [default to null]
**PendingPurchase** | **float32** | Amounts from a deprecated feature. This value will always be &#x60;0&#x60;. | [default to null]
**BalanceWithoutAuths** | **float32** | The &#x60;balance&#x60; without subtracting out authorization holds. This field is populated only when the BWOOA parameter is set. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

