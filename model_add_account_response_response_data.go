/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A structure for the response data. It can be empty but usually will contain information.
type AddAccountResponseResponseData struct {
	// A Galileo generated account number
	PmtRefNo string `json:"pmt_ref_no"`
	// Galileo generated integer ID for the product
	ProductId string `json:"product_id"`
	// Galileo generated integer account number, also known as balance ID
	GalileoAccountNumber string `json:"galileo_account_number"`
	// Unique identifier for a card. None is returned if no card was created.
	CardId string `json:"card_id,omitempty"`
	// A PAN or 16 digit card number (usually masked)
	CardNumber string `json:"card_number,omitempty"`
	// Expiration date of the card
	ExpiryDate string `json:"expiry_date,omitempty"`
	// Card security code (CVV)
	CardSecurityCode string `json:"card_security_code,omitempty"`
	// Second embossed line on a card
	EmbossLine2 string `json:"emboss_line_2,omitempty"`
	// The day of the month the billing cycle starts for the account
	BillingCycleDay int32 `json:"billing_cycle_day"`
}