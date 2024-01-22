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
type GetMerchantControlsResponseResponseData struct {
	// The product ID. Returned only when `prodId` was passed in the original call.
	ProductId int32 `json:"product_id,omitempty"`
	// List of MCC controls
	MerchantControls []ResponseData72MerchantControls `json:"merchant_controls"`
	// The PRN of the account. Returned only when `accountNo` was passed in the original call.
	PmtRefNo string `json:"pmt_ref_no,omitempty"`
}