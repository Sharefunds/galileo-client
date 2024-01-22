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
type GetIngoCustomerIdResponseResponseData struct {
	// ID for Ingo Money system
	IngoCustomerId string `json:"ingo_customer_id"`
	RegisteredCard *ResponseData66RegisteredCard `json:"registered_card"`
	// SSO Token used for SDK
	IngoSsoToken string `json:"ingo_sso_token"`
	// Additional error messages from Ingo system
	IngoMessage string `json:"ingo_message"`
}
