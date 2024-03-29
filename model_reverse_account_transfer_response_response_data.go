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
type ReverseAccountTransferResponseResponseData struct {
	// The balance on the account after the transaction posted
	NewBalance float64 `json:"new_balance"`
	// The balance on the account before the transaction posted
	OldBalance float64 `json:"old_balance"`
}
