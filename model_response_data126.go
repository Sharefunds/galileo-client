/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type ResponseData126 struct {
	// The number of pending deposits in the response
	PendingDepositCount int32 `json:"pending_deposit_count"`
	// The page number to be retrieved in the context of recordset paging
	Page int32 `json:"page"`
	// Number of records in the accounts list display
	TotalRecordCount int32 `json:"total_record_count"`
	// Total number of pages in the accounts list display
	NumberOfPages int32 `json:"number_of_pages"`
	// Start date of the pending deposits data range
	StartDate time.Time `json:"start_date"`
	// End date of the pending deposits data range
	EndDate time.Time `json:"end_date"`
	// List of information on pending deposits
	PendingDeposits []ResponseData126PendingDeposits `json:"pending_deposits"`
}