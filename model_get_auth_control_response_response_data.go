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
type GetAuthControlResponseResponseData struct {
	// PRN. This field is populated when returning account-level velocity controls.
	PmtRefNo string `json:"pmt_ref_no,omitempty"`
	// List of auth controls
	AuthControls []ResponseData70AuthControls `json:"auth_controls"`
	// This field is populated when returning product-level velocity controls.
	ProdId string `json:"prod_id,omitempty"`
}
