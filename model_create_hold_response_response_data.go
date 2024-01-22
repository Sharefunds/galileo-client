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
type CreateHoldResponseResponseData struct {
	// The balance on the account after the transaction
	NewBalance float64 `json:"new_balance"`
	// The ID associated with the hold
	HoldId int32 `json:"hold_id"`
}
