/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData5 struct {
	Customer *ResponseData5Customer `json:"customer"`
	// List of account objects
	Accounts []ResponseData5Accounts `json:"accounts"`
}
