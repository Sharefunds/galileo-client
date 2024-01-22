# SavingsInterest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | [**time.Time**](time.Time.md) | The start date for the period | [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The end date for the period | [default to null]
**AccrualInterest** | [***Object**](.md) | Interest accrued on the account. May return an integer if the interest accrued is 0. | [default to null]
**InterestYtd** | [***Object**](.md) | The year-to-date interest paid on a savings account. May return an integer if the interest paid is 0. | [default to null]
**Apy** | [***Object**](.md) | The amount earned with compound interest over the course of one year. May return an integer if the amount earned is 0. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

