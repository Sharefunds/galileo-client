# Payment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PmtId** | **string** | ID assigned to the specific payment | [default to null]
**Details** | **string** | Description of the payment | [default to null]
**Amount** | **string** | An amount of a payment | [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The date and time of the payment | [default to null]
**SourceId** | **string** | A code unique to the source of the payment | [default to null]
**AchTransactionId** | **string** | A unique ID for an ACH transaction | [default to null]
**ExternalTransId** | **string** | Non-Galileo ID for a transaction | [default to null]
**HoldDays** | **string** | The number of days on a hold. Specific to a pending payment, usually due to load limt violations | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

