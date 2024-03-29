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

type ResponseData3PendingFees struct {
	// Galileo generated fee transaction integer ID
	FeeEventId string `json:"fee_event_id"`
	// Three-letter fee code. This is not the transaction type (otype).
	Type_ string `json:"type"`
	// A description of the type code
	TypeDescription string `json:"type_description"`
	// Amount of the fee charge
	Amt string `json:"amt"`
	// A timestamp for the time the fee was charged
	FeeDate time.Time `json:"fee_date"`
	// Integer identifier of the card as found in the raw data file (RDF). Unique identifier for a PAN.
	CardId int32 `json:"card_id"`
	// The description on a fee
	FeeDescription string `json:"fee_description"`
	RelatedTransaction *FeeRelatedTransaction `json:"related_transaction"`
}
