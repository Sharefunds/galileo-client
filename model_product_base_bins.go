/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ProductBaseBins struct {
	// Whether the base bin is active or not. `Y` or `N`.
	Active string `json:"active,omitempty"`
	// Base bin
	BaseBin string `json:"base_bin"`
}
