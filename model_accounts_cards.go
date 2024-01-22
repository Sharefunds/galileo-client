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

type AccountsCards struct {
	// Unique identifier for a card
	CardId string `json:"card_id"`
	// A PAN or 16 digit card number (usually masked)
	CardNumber string `json:"card_number"`
	// UUID for an emboss record that has been sent to the embosser. When this field is present outside the `embossed_cards` object, it contains the most recent emboss UUID.
	EmbossUuid string `json:"emboss_uuid,omitempty"`
	// A single character code specifying the status of the card
	CardStatus string `json:"card_status"`
	// Expiration date of the card
	ExpiryDate time.Time `json:"expiry_date"`
	// Partner specified external card identity
	ExternalCardId string `json:"external_card_id"`
	// Counts the number of times a PIN entry fails
	PinFailCount string `json:"pin_fail_count"`
	// Date a PIN fail was recorded
	PinFailDate time.Time `json:"pin_fail_date"`
	FreezeInfo *CardFreezeInfo `json:"freeze_info"`
	// List of embossed card objects
	EmbossedCards []CardEmbossedCards `json:"embossed_cards"`
}
