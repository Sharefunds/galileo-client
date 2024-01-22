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
type GetCardResponseResponseData struct {
	// A PAN or 16 digit card number
	CardNumber string `json:"card_number"`
	// Expiration date of the card
	ExpiryDate string `json:"expiry_date"`
	// The card verification value (cvv)
	CardSecurityCode string `json:"card_security_code"`
	// The status of the card
	Status string `json:"status"`
	// Integer identifier of the card as found in the raw data file (RDF). Unique identifier for a PAN.
	CardId string `json:"card_id"`
	// External card identity
	ExternalCardId string `json:"external_card_id"`
	// A Galileo generated account number
	PmtRefNo string `json:"pmt_ref_no"`
	// A person's first name as listed on the account
	FirstName string `json:"first_name"`
	// A person's middle name as listed on the account
	MiddleName string `json:"middle_name"`
	// A person's last name as listed on the account
	LastName string `json:"last_name"`
	// Encrypted number for the card
	EncryptedCardNumber string `json:"encrypted_card_number"`
	// Encrypted expiration date for the card
	EncryptedExpiryDate string `json:"encrypted_expiry_date"`
	// UUID for an emboss record that has been sent to the embosser. When this field is present outside the `embossed_cards` object, it contains the most recent emboss UUID.
	EmbossUuid string `json:"emboss_uuid,omitempty"`
	// List of embossed card objects
	EmbossedCards []ResponseData17EmbossedCards `json:"embossed_cards"`
	FreezeInfo *CardFreezeInfo `json:"freeze_info"`
	// A flag for the PIN fail count (0 or -1)
	PinFailCount string `json:"pin_fail_count"`
	// Date a PIN fail was recorded
	PinFailDate time.Time `json:"pin_fail_date"`
	SpendControls *ResponseData17SpendControls `json:"spend_controls"`
	// The social security number associated with the account. This field is returned only if `CINFN` is set.
	Ssn string `json:"ssn,omitempty"`
}