# ResponseData70AuthControls

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomesticFlag** | **string** | Whether this velocity control applies to domestic transactions, international transactions or both. **Y**—Domestic only; **N**—International only; **A**—Both | [optional] [default to null]
**Period** | **string** | Timespan for the velocity control. xD (x days), xM (x months), 1T (one transaction). | [optional] [default to null]
**TransactionCountUsed** | **int32** | Number of transactions used in the current period. | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | Ending date-time of the account-level velocity control. Not populated for product-level controls. | [optional] [default to null]
**TransType** | **string** | Type of transaction affected by the control: ATM, CAD, CBA, POS, VFT. | [optional] [default to null]
**ControlDesc** | **string** | Description of the velocity control. | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Starting date-time of the account-level velocity control. Not populated for product-level controls. | [optional] [default to null]
**AmountUsed** | **float32** | Transaction amount used in the current period. | [optional] [default to null]
**ControlId** | **int32** | The identifier for the specific velocity control. | [optional] [default to null]
**HasPin** | **string** | Whether the limit applies to PIN transactions, signature transactions or both. **Y**—PIN only; **N**—Signature only; **A**—Both | [optional] [default to null]
**AmountAvailable** | **float32** | Transaction amount available in the current period. | [optional] [default to null]
**TransactionCountAvailable** | **int32** | Number of transactions available in the current period. | [optional] [default to null]
**Amount** | **float32** | Total transaction amount allowed per period. | [optional] [default to null]
**TransactionCount** | **int32** | Number of allowed transactions per period. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

