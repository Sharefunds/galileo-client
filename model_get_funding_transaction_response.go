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

type GetFundingTransactionResponse struct {
	// The response status code. May return a string for some statuses.
	StatusCode int32 `json:"status_code"`
	// The condition of a process or response
	Status string `json:"status"`
	// The time elapsed in processing the transaction
	ProcessingTime float32 `json:"processing_time"`
	Echo *GetTransHistoryResponseEcho `json:"echo"`
	// A system generated timestamp
	SystemTimestamp time.Time `json:"system_timestamp"`
	// A Galileo-generated ID used for tracking
	Rtoken string `json:"rtoken"`
	// A list of errors generated while the request was processed
	Errors []string `json:"errors,omitempty"`
	ResponseData *GetFundingTransactionResponseResponseData `json:"response_data"`
}
