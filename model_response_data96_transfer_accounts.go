/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData96TransferAccounts struct {
	// Transfer account ID
	TransferAccountId string `json:"transfer_account_id"`
	// The first name on the transfer account
	FirstName string `json:"first_name"`
	// The last name on the transfer account
	LastName string `json:"last_name"`
	// The email associated with the transfer account
	Email string `json:"email"`
	// A nickname for the account
	Nickname string `json:"nickname,omitempty"`
	// Payment reference number for the transfer account
	PmtRefNo string `json:"pmt_ref_no,omitempty"`
	// The status of the transfer account
	Status string `json:"status"`
}
