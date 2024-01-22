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

type Transaction1 struct {
	// Specifies whether the transaction is made on a savings account.
	IsSavings bool `json:"is_savings"`
	// Two-digit code when an authorization request is denied.
	DenyCode string `json:"deny_code"`
	// Specifies whether the transaction can be disputed in a chargeback.
	Disputable bool `json:"disputable"`
	// Description provided by the merchant about the transaction (DE043).
	Details string `json:"details"`
	// Activity type. See the <a href=\"ref:api-reference-activity-type\" target=\"_blank\">Activity Type</a> enumeration.
	ActType string `json:"act_type"`
	// Description of `act_type`. See the <a href=\"ref:api-reference-activity-type\" target=\"_blank\">Activity Type</a> enumeration.
	ActTypeDescription string `json:"act_type_description"`
	// The Galileo system date/time when a transaction posted to the customer account, in Mountain Standard Time (-0700).
	PostTs time.Time `json:"post_ts"`
	// The transaction amount in the currency of the account. A negative amount debits funds from the customer account.
	Amt float32 `json:"amt"`
	// A Galileo-generated integer that maps back to the original transaction, such as `auth_id`, `pmt_id`, or `adj_id`.
	SourceId string `json:"source_id"`
	// Reference your program's authorization <a href=\"ref:api-reference-transaction-types\" target=\"_blank\">transaction types</a> for possible values.
	Type_ string `json:"type"`
	// The description of the transaction type.
	TypeDescription string `json:"type_description"`
	// A concatenation of the activity type (`act_type`) and transaction type (`type`).
	TransCode string `json:"trans_code"`
	// Acquirer reference number. An identifier for the acquiring processor.
	Arn string `json:"arn"`
	// Network-assigned identifier for a merchant (DE042).
	MerchantId string `json:"merchant_id"`
	// A non-Galileo identifier for a transaction, if any.
	ExternalTransId string `json:"external_trans_id"`
	// Available balance (including this transaction) at the time the API was called.
	CalculatedBalance float32 `json:"calculated_balance"`
	// A unique identifier for the ACH transaction, if applicable.
	AchTransId string `json:"ach_trans_id"`
	// The Galileo system date/time a transaction was authorized, in Mountain Standard Time (-0700).
	AuthTs time.Time `json:"auth_ts"`
	// The `auth_id` of the previous transaction in the sequence, if any.  Maps to `original_auth_id`.
	PriorId string `json:"prior_id"`
	// A Galileo-generated identifier for a card, which can be used instead of the PAN. Maps to `cad`.
	CardId string `json:"card_id"`
	// The same information as in the `details` field, with formatting.
	FormattedMerchantDesc string `json:"formatted_merchant_desc"`
	// A Galileo-generated code to identify the subnetwork over which the transaction took place. Maps to `network_id`.
	NetworkCode string `json:"network_code"`
	// A Galileo-generated identifier for an authorization and its settlement.
	AuthId string `json:"auth_id"`
	// Amount, in cents, of the authorization request, in the currency at the point of sale. 12-digit number including leading zeros. This amount does not include upcharges or program fees. (DE004)
	LocalAmt float32 `json:"local_amt"`
	// Currency code for `local_amt` (DE049).
	LocalCurrCode string `json:"local_curr_code"`
	// The transaction amount, in the settlement currency (DE005)
	SettleAmt float32 `json:"settle_amt"`
	// Currency code for `settle_amt` (DE050).
	SettleCurrCode string `json:"settle_curr_code"`
	// The transaction amount, in the currency of the account (DE006) 
	BillingAmt float32 `json:"billing_amt"`
	// Currency code for `billing_amt` (DE051).
	BillingCurrCode string `json:"billing_curr_code"`
	// A Galileo-generated number to identify the customer account. Maps to `PRN` and `prn`.
	PmtRefNo string `json:"pmt_ref_no"`
	// Category code for the merchant that initiated the transaction (DE018).
	MccCode string `json:"mcc_code"`
	// Indicates whether a PIN was input at the point of sale. `Y` = No PIN was input. `N` = A PIN was input. `None` = Not a card transaction, or not processed as a card transaction.
	CreditInd string `json:"credit_ind"`
	// Impuesto al Consumo. Colombian consumption tax. Required for the Mastercard Interchange Intracountry Calculation Program.
	IacTax float32 `json:"iac_tax,omitempty"`
	// Impuesto al Valor Agregado. Colombian value-added tax. Required for the Mastercard Interchange Intracountry Calculation Program.
	IvaTax float32 `json:"iva_tax,omitempty"`
	// The <<glossary:PRN>> of the <<glossary:RTF>> funding account
	FundingAccountPrn string `json:"funding_account_prn,omitempty"`
	// The PRN of the RTF spending account
	SpendingAccountPrn string `json:"spending_account_prn,omitempty"`
}