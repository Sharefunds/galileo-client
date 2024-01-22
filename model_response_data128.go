/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData128 struct {
	// List of scheduled payments
	ScheduledPayments []ResponseData128ScheduledPayments `json:"scheduled_payments"`
	// The number of scheduled payments found
	Found int32 `json:"found"`
	// List of future scheduled payments
	FutureScheduledPayments []ResponseData128FutureScheduledPayments `json:"future_scheduled_payments"`
	// The page number to be retrieved in the context of recordset paging
	Page int32 `json:"page"`
	// The page number to be retrieved for scheduled bill payments
	PageScheduled int32 `json:"page_scheduled,omitempty"`
	// The page number to be retrieved for future scheduled bill payments
	PageFuture int32 `json:"page_future,omitempty"`
	// Sum of scheduled payments and future scheduled payments
	TotalRecordCount int32 `json:"total_record_count"`
	// Total number of pages for scheduled payments and future scheduled payments
	NumberOfPages int32 `json:"number_of_pages"`
	// Number of records for scheduled payments
	TotalRecordCountScheduled int32 `json:"total_record_count_scheduled,omitempty"`
	// Total number of pages for the scheduled payments
	NumberOfPagesScheduled int32 `json:"number_of_pages_scheduled,omitempty"`
	// Number of records for future scheduled payments
	TotalRecordCountFuture int32 `json:"total_record_count_future,omitempty"`
	// Total number of pages for the future scheduled payments
	NumberOfPagesFuture int32 `json:"number_of_pages_future,omitempty"`
}
