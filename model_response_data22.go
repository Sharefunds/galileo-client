/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData22 struct {
	// A code used to give the status of the simulated auth
	AuthResponseCode string `json:"auth_response_code"`
	// A description of the authorization response
	AuthResponseDescription string `json:"auth_response_description"`
	// An ID associated with the simulated auth
	AuthId string `json:"auth_id"`
}
