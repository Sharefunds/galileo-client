# GetRelatedAccountsResponseResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiblingAccounts** | [**[]ResponseData54SiblingAccounts**](ResponseData54_sibling_accounts.md) | If &#x60;accountNo&#x60; contains a secondary account, these are other secondary accounts that are associated with the primary account in &#x60;parent_accounts&#x60;. | [default to null]
**ParentAccounts** | [**[]ResponseData54SiblingAccounts**](ResponseData54_sibling_accounts.md) | If &#x60;accountNo&#x60; contains a secondary account, this is the primary account that it is associated with. | [default to null]
**SelfAccounts** | [**[]ResponseData54SiblingAccounts**](ResponseData54_sibling_accounts.md) | Accounts that belong to the same customer. | [default to null]
**ChildAccounts** | [**[]ResponseData54SiblingAccounts**](ResponseData54_sibling_accounts.md) | Accounts that are secondary accounts to the primary account in &#x60;accountNo&#x60;. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

