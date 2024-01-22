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
type VerifyChLoginResponseResponseData struct {
	AccountId int32 `json:"account_id,omitempty"`
	ProgramId int32 `json:"program_id,omitempty"`
	// The validation status
	ValidationStatus int32 `json:"validation_status"`
}
