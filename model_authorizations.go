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

type Authorizations struct {
	// A Galileo-generated identifier for an authorization.
	AuthId string `json:"auth_id"`
	// Description provided by the merchant about the transaction (DE043)
	Details string `json:"details"`
	// The same information as in the `details` field, with formatting
	DetailsFormatted string `json:"details_formatted"`
	// The authorization amount, in the currency of the account
	Amount string `json:"amount"`
	// The Galileo system timestamp for the authorization, in Mountain Standard Time (-0700)
	Timestamp time.Time `json:"timestamp"`
	// Reference your program's authorization <a href=\"ref:api-reference-transaction-types\" target=\"_blank\">transaction types</a> for possible values.
	Type_ string `json:"type"`
	// Category code for the merchant that initiated the transaction (DE018)
	Mcc string `json:"mcc"`
	// Network-assigned identifier for a merchant (DE042)
	MerchantId string `json:"merchant_id"`
	// The identifier for the acquirer (DE032)
	AcqId string `json:"acq_id"`
	// Identifier for the card reader at the point of sale (DE041)
	TerminalId string `json:"terminal_id"`
	// Whether to allow an expiration on this authorization: `1` = Allow, or `0`  = Do not allow
	CanBeExpired string `json:"can_be_expired"`
	// The `auth_id` of the previous transaction in the sequence, if any. Maps to `prior_id` in other contexts.
	OriginalAuthId string `json:"original_auth_id"`
	// A Galileo-generated code to identify the network over which the transaction took place. Maps to `network_id`.
	NetworkCode string `json:"network_code"`
	// Amount in cents of the transaction based on the currency at the point of sale (DE004). 12-digit number including leading zeros.
	LocalAmt string `json:"local_amt"`
	// Currency code for `local_amt` (DE049)
	LocalCurrCode string `json:"local_curr_code"`
	// The transaction amount, in cents, in the settlement currency (DE005). 12-digit number including leading zeros.
	SettleAmt string `json:"settle_amt"`
	// Currency code for `settle_amt` (DE050)
	SettleCurrCode string `json:"settle_curr_code"`
	// The billing amount in cents (DE006). 12-digit number including leading zeros.
	BillingAmt string `json:"billing_amt"`
	// Currency code for `billing_amt` (DE051)
	BillingCurrCode string `json:"billing_curr_code"`
	// Impuesto al Consumo. Colombian consumption tax. Required for the Mastercard Interchange Intracountry Calculation Program.
	IacTax float32 `json:"iac_tax,omitempty"`
	// Impuesto al Valor Agregado. Colombian value-added tax. Required for the Mastercard Interchange Intracountry Calculation Program.
	IvaTax float32 `json:"iva_tax,omitempty"`
}
