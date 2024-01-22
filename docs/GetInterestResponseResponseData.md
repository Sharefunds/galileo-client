# GetInterestResponseResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **int32** | Product Id | [default to null]
**ProductDescription** | **string** | Product Description | [default to null]
**AccountNo** | **int32** | Account No | [default to null]
**InterestPayout** | **float32** | Interest accrued on account to two decimal places. This includes carryover from the previous month and is what the account has been paid or will be paid. | [default to null]
**AccrualInterestWithoutCarryover** | **float32** | Interest accrued on the account without carryover from the previous month | [default to null]
**CarryoverPreviousMonth** | **float32** | The remainder interest from the previous month that is carried over to the requested month | [default to null]
**InterestPaidYtd** | **float32** | The interest paid on the account for today&#x27;s year to date. The interestMonth argument does not affect this value. | [default to null]
**Apye** | **float32** | The Annual Percentage Yield Earned (APYE) for the requested month | [default to null]
**StartDate** | **string** | The start date of the requested month | [default to null]
**EndDate** | **string** | The end date of the requested month | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

