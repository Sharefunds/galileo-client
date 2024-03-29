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

// A structure for the response data. It can be empty but usually will contain information.
type VerifyEnrollmentResponseResponseData struct {
	// Timestamp for when the response data was processed
	ProcessedDate time.Time `json:"processed_date"`
	// Identifier for the enrollment application
	AppTransId time.Time `json:"app_trans_id"`
	// Maximum credit limit
	MaxCreditLimit float32 `json:"max_credit_limit,omitempty"`
	Application *ResponseData36Application `json:"application"`
}
