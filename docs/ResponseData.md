# ResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | [**time.Time**](time.Time.md) | The start date for the transaction range | [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The end date for the transaction range | [default to null]
**NumberOfPages** | **int32** | The total number of pages available in the paginated response | [default to null]
**Page** | **int32** | The current page being retunred in the paginated response | [default to null]
**TotalRecordCount** | **int32** | Number of records in the accounts list display | [default to null]
**TransactionCount** | **int32** | The number of transactions listed in the response | [default to null]
**Transactions** | [**[]ResponseDataTransactions**](ResponseData_transactions.md) | List of transactions | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

