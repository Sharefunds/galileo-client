/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData24TransferToAccount struct {
	// The old balance on the destination account
	OldBalance float32 `json:"old_balance"`
	// The new balance on the destination account
	NewBalance float32 `json:"new_balance"`
}
