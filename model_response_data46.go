/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData46 struct {
	// An associative array with keys that identify the alert type.
	Alerts map[string]ResponseData46Alerts `json:"alerts"`
}
