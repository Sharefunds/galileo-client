/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateVirtualCardResponseResponseData struct {
	// Galileo-generated account number. Also called the PRN.
	PmtRefNo string `json:"pmt_ref_no"`
	// Identifier for the product associated with the account
	ProductId string `json:"product_id"`
	// Galileo-generated balance ID
	GalileoAccountNumber string `json:"galileo_account_number"`
	// Galileo-generated card identifier (CAD). Use this ID instead of the PAN if not PCI-compliant. `None` means that no card was created
	CardId string `json:"card_id"`
	// The primary account number (PAN) printed on the card
	CardNumber string `json:"card_number"`
	// Date when the card expires
	ExpiryDate string `json:"expiry_date"`
	// The card verification value (CVV)
	CardSecurityCode string `json:"card_security_code"`
	// List of cardholder identification process (CIP) objects
	Cip []ResponseData4Cip `json:"cip"`
	// The day of the month the billing cycle starts for the account
	BillingCycleDay int32 `json:"billing_cycle_day"`
}
