# ResponseData3

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PmtRefNo** | **string** | Galileo-generated account number. Also called the PRN. | [default to null]
**Status** | **string** | Status of the account | [default to null]
**ApplicationDate** | **string** | Date the account application was received | [default to null]
**Balance** | **float32** | The current account balance | [default to null]
**CurrencyCode** | **string** | A three-character code to represent global currencies | [default to null]
**Profile** | [***ResponseData3Profile**](ResponseData3_profile.md) |  | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The start date for the range that account information is displayed | [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The end date for the range that account information is displayed | [default to null]
**TransactionCount** | **int32** | The number of transactions listed in the response | [default to null]
**Transactions** | [**[]ResponseData3Transactions**](ResponseData3_transactions.md) | List of transactions | [default to null]
**AuthorizationCount** | **int32** | The number of authorizations listed in the response | [default to null]
**Authorizations** | [**[]ResponseData2Authorizations**](ResponseData2_authorizations.md) | List of authorizations | [default to null]
**PendingFees** | [**[]ResponseData3PendingFees**](ResponseData3_pending_fees.md) | List of fees | [default to null]
**SavingsInterest** | [***ResponseData3SavingsInterest**](ResponseData3_savings_interest.md) |  | [default to null]
**Holds** | [**[]ResponseData3Holds**](ResponseData3_holds.md) | List of holds. Holds being returned are dependant on product parameters | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

