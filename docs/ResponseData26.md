# ResponseData26

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionCount** | **int32** | The number of transactions listed in the response | [default to null]
**Page** | **int32** | The page number to be retrieved in the context of recordset paging | [default to null]
**TotalRecordCount** | **int32** | Number of records in the accounts list display | [default to null]
**NumberOfPages** | **int32** | Total number of pages in the accounts list display | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Start date of the response | [default to null]
**EndDate** | [**time.Time**](time.Time.md) | End date of the response | [default to null]
**Sums** | [***ResponseData26Sums**](ResponseData26_sums.md) |  | [default to null]
**Transactions** | [**[]ResponseData26Transactions**](ResponseData26_transactions.md) | List of transactions | [default to null]
**BeginningBalance** | **float32** | The account balance as of the &#x60;startDate&#x60;. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

