/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData9 struct {
	// Reflects the pre-transaction amount on the account. For outgoing credits the balance adjustment occurs almost instantly but the response shows the balance before the transaction. For outgoing debits, the balance update is delayed until the hold period expires.
	Balance float32 `json:"balance"`
	// Unique identifier for the ACH transaction
	AchTransactionId string `json:"ach_transaction_id"`
}
