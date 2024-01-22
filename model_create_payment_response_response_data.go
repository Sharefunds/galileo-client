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
type CreatePaymentResponseResponseData struct {
	// The balance on the account before the transaction
	OldBalance float64 `json:"old_balance"`
	// The balance on the account after the transaction
	NewBalance float64 `json:"new_balance"`
	// Payment fee amount. May return an integer if the fee amount is 0.
	FeeAmount string `json:"fee_amount"`
	// Galileo generated card payment or load transaction integer ID
	PaymentTransId int32 `json:"payment_trans_id"`
	// An identifier that represents a transaction
	TransactionId string `json:"transaction_id"`
	// Identifier for the hold. Use the <a href=\"ref:post_expirehold\" target=\"_blank\">Expire Hold</a> endpoint to expire the hold before the time in `holdExpirationDateTime`.
	HoldId int32 `json:"hold_id"`
}
