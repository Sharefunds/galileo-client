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
type ModifyStatusResponseResponseData struct {
	// A Galileo generated account number
	PmtRefNo string `json:"pmt_ref_no"`
	// The condition of the account after a modification
	AccountStatus string `json:"account_status"`
	// UUID for a new_emboss record that has been sent to the embosser
	NewEmbossUuid string `json:"new_emboss_uuid,omitempty"`
}
