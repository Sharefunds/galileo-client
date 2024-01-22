/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData85 struct {
	// The updated credit limit for the account
	NewCreditLimit string `json:"new_credit_limit,omitempty"`
	// The amount available to be charged to a credit card
	AvailableCredit string `json:"available_credit,omitempty"`
}