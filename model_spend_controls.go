/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SpendControls struct {
	// The amount available to be charged to a credit card
	AvailableCredit string `json:"available_credit"`
	// A flag that indicated an alias credit card number, exhausted after one use
	SingleUse string `json:"single_use"`
	// The maximum outstanding balance allowed on a credit card
	CreditLimit string `json:"credit_limit"`
}
