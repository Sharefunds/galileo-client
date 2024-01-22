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

type Change struct {
	// The time the credit limit change was made
	Timestamp time.Time `json:"timestamp"`
	// The monetary amount the credit limit was changed to
	CreditLimitAmount float32 `json:"credit_limit_amount"`
}