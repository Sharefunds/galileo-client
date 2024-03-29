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
type RunCipResponseResponseData struct {
	// A flag that indicates whether a CIP result exists in the response
	CipResultIndicator string `json:"cip_result_indicator"`
	CipResult []ResponseData39CipResult `json:"cip_result"`
}
