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
type SetOverdraftLimitResponseResponseData struct {
	// After an overdraft limit change, lists the previous customer overdraft limit
	OldOverdraftLimit string `json:"old_overdraft_limit"`
	// Lists the changed customer overdraft limit
	NewOverdraftLimit string `json:"new_overdraft_limit"`
}
