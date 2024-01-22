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
type CreateBillPaymentResponseResponseData struct {
	// An ID assigned to a bill payment transaction
	BillpayTransactionId string `json:"billpay_transaction_id"`
	// The balance of the account before the payment is processed
	OldBalance float32 `json:"old_balance"`
	// The balance on the account after the payment is processed
	NewBalance float32 `json:"new_balance"`
	// The date the payment was processed
	ProcessDate string `json:"process_date"`
	// The fee amount assessed for this bill payment record
	FeeAmount float32 `json:"fee_amount"`
	// A dollar limit on each bill-pay transaction
	MaximumAmount float32 `json:"maximum_amount"`
}
