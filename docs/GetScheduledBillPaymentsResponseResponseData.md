# GetScheduledBillPaymentsResponseResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledPayments** | [**[]ResponseData128ScheduledPayments**](ResponseData128_scheduled_payments.md) | List of scheduled payments | [default to null]
**Found** | **int32** | The number of scheduled payments found | [default to null]
**FutureScheduledPayments** | [**[]ResponseData128FutureScheduledPayments**](ResponseData128_future_scheduled_payments.md) | List of future scheduled payments | [default to null]
**Page** | **int32** | The page number to be retrieved in the context of recordset paging | [default to null]
**PageScheduled** | **int32** | The page number to be retrieved for scheduled bill payments | [optional] [default to null]
**PageFuture** | **int32** | The page number to be retrieved for future scheduled bill payments | [optional] [default to null]
**TotalRecordCount** | **int32** | Sum of scheduled payments and future scheduled payments | [default to null]
**NumberOfPages** | **int32** | Total number of pages for scheduled payments and future scheduled payments | [default to null]
**TotalRecordCountScheduled** | **int32** | Number of records for scheduled payments | [optional] [default to null]
**NumberOfPagesScheduled** | **int32** | Total number of pages for the scheduled payments | [optional] [default to null]
**TotalRecordCountFuture** | **int32** | Number of records for future scheduled payments | [optional] [default to null]
**NumberOfPagesFuture** | **int32** | Total number of pages for the future scheduled payments | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

