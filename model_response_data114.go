/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData114 struct {
	// The balance on the account after the transaction
	NewBalance float64 `json:"new_balance"`
	// The ID associated with the hold
	HoldId int32 `json:"hold_id"`
}
