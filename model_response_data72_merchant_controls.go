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

type ResponseData72MerchantControls struct {
	// Ending date-time of the account-level control. Not populated for product-level controls.
	EndDate time.Time `json:"end_date,omitempty"`
	// Whether the control is DENY (`d`) or ALLOW (`a`).
	DenyAllowFlag string `json:"deny_allow_flag"`
	// Starting date-time of the account-level control. Not populated for product-level controls.
	StartDate time.Time `json:"start_date,omitempty"`
	// Merchant ID. If the value passed has fewer than 15 characters, the Galileo system adds spaces at the end until it has 15 characters.
	MerchantId string `json:"merchant_id"`
}
