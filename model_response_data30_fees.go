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

type ResponseData30Fees struct {
	// A Galileo generated account number
	PmtRefNo string `json:"pmt_ref_no"`
	// The ID of the fee record
	FeeId string `json:"fee_id"`
	// A timestamp for the time the fee was charged
	FeeDate time.Time `json:"fee_date"`
	// A fee or transaction amount
	Amt string `json:"amt"`
	// The status of a fee event record return in this method
	Status string `json:"status"`
	// A spelled-out status of the fee
	StatusDescription string `json:"status_description"`
	// Three-letter fee code. This is not the transaction type (otype).
	Type_ string `json:"type"`
	// A description of the type code
	TypeDescription string `json:"type_description"`
	// An ID for the fee event
	FeeEventId string `json:"fee_event_id"`
}
