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
type AssessFeeResponseResponseData struct {
	// The balance on the account before the reversal is processed
	OldBalance float64 `json:"old_balance"`
	// The balance on the account after the reversal is processed
	NewBalance float64 `json:"new_balance"`
	// Numeric amount of the fee
	FeeAmount float64 `json:"fee_amount"`
	// A number that represents transaction
	TransactionId string `json:"transaction_id"`
	// Galileo generated fee transaction integer ID
	FeeTransId int32 `json:"fee_trans_id"`
}