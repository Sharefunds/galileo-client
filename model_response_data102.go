/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData102 struct {
	// The number of ACH accounts found
	Found int32 `json:"found"`
	// List of ACH accounts
	AchAccounts []ResponseData102AchAccounts `json:"ach_accounts"`
}
