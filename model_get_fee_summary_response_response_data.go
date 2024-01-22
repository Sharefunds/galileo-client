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
type GetFeeSummaryResponseResponseData struct {
	// Fees year-to-date
	FeesYtd float32 `json:"fees_ytd"`
	// Fees month-to-date
	FeesMtd float32 `json:"fees_mtd"`
}