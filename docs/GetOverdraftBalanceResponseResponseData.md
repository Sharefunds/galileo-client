# GetOverdraftBalanceResponseResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverdraftAccountNo** | **string** | PRN of the overdraft account | [default to null]
**Available** | **float32** | Available overdraft limit | [default to null]
**Payback** | **float32** | Payback amount | [default to null]
**PaybackWithPendingFees** | **float32** | The posted open to buy, including pending fees | [default to null]
**RawBalance** | **float32** | The actual open to buy of the overdraft account | [default to null]
**ActualWithPrimary** | **float32** | The actual balance of the primary account if overdraft was combined with it | [default to null]
**Limit** | **float32** | The available limit on the overdraft account (the most the cardholder could use on their overdraft if nothing was owed) | [default to null]
**RawLimit** | **float32** | The limit that was set on the overdraft account (the actual limit available to the cardholder may be less if an initial fee was charged) | [default to null]
**Frozen** | **string** | A flag indicating whether the overdraft account has been frozen or not (the value &#x27;1&#x27; indicating that it is frozen) | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

