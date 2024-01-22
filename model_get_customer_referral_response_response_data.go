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
type GetCustomerReferralResponseResponseData struct {
	// ID of the referral
	ReferralId float32 `json:"referral_id"`
	// The date of the referral
	ReferralDate time.Time `json:"referral_date"`
	// The status of the referral
	ReferralStatus string `json:"referral_status"`
	// The first name of the person being referred
	ReferralFirstName string `json:"referral_first_name"`
	// The last name of the person being referred
	ReferralLastName string `json:"referral_last_name"`
	// The email of the person being referred
	ReferralEmail string `json:"referral_email"`
	// Payment reference number of the account associated with the referrer
	ReferrerPmtRefNo string `json:"referrer_pmt_ref_no,omitempty"`
}