# GetRtfAccountRelationshipResponseResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** | The page number to be retrieved in the context of recordset paging | [optional] [default to null]
**SpendingAccountPrn** | **string** | The PRN of the spending account. | [optional] [default to null]
**SpendingAccountCount** | **int32** | Number of spending accounts returned. | [optional] [default to null]
**TotalRecordCount** | **int32** | Number of records in the accounts list display | [optional] [default to null]
**SpendingAccounts** | [**[]ResponseData73SpendingAccounts**](ResponseData73_spending_accounts.md) | Contains a list of RTF spending accounts associated with the RTF funding account. | [optional] [default to null]
**NumberOfPages** | **int32** | Total number of pages in the accounts list display | [optional] [default to null]
**FundingAccountPrn** | **string** | The PRN of the funding account. | [optional] [default to null]
**FundingAccount** | [***ResponseData73FundingAccount**](ResponseData73_funding_account.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

