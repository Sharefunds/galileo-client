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

type GetAuthControl1 struct {
	// Whether this velocity control applies to domestic transactions, international transactions or both. **Y**—Domestic only; **N**—International only; **A**—Both
	DomesticFlag string `json:"domestic_flag,omitempty"`
	// Timespan for the velocity control. xD (x days), xM (x months), 1T (one transaction).
	Period string `json:"period,omitempty"`
	// Number of transactions used in the current period.
	TransactionCountUsed int32 `json:"transaction_count_used,omitempty"`
	// Ending date-time of the account-level velocity control. Not populated for product-level controls.
	EndDate time.Time `json:"end_date,omitempty"`
	// Type of transaction affected by the control: ATM, CAD, CBA, POS, VFT.
	TransType string `json:"trans_type,omitempty"`
	// Description of the velocity control.
	ControlDesc string `json:"control_desc,omitempty"`
	// Starting date-time of the account-level velocity control. Not populated for product-level controls.
	StartDate time.Time `json:"start_date,omitempty"`
	// Transaction amount used in the current period.
	AmountUsed float32 `json:"amount_used,omitempty"`
	// The identifier for the specific velocity control.
	ControlId int32 `json:"control_id,omitempty"`
	// Whether the limit applies to PIN transactions, signature transactions or both. **Y**—PIN only; **N**—Signature only; **A**—Both
	HasPin string `json:"has_pin,omitempty"`
	// Transaction amount available in the current period.
	AmountAvailable float32 `json:"amount_available,omitempty"`
	// Number of transactions available in the current period.
	TransactionCountAvailable int32 `json:"transaction_count_available,omitempty"`
	// Total transaction amount allowed per period.
	Amount float32 `json:"amount,omitempty"`
	// Number of allowed transactions per period.
	TransactionCount int32 `json:"transaction_count,omitempty"`
}
