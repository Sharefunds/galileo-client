/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Category struct {
	// Category code for the deposit. See <a href=\"ref:api-reference-deposit-category-codes\" target=\"_blank\">Deposit Category Codes</a> for valid values.
	CategoryCode string `json:"category_code"`
	// A plain text description of a pending deposit
	Description string `json:"description"`
	// Status of the deposit
	Status string `json:"status"`
	// A document or record that supports the transaction
	AchSourceId string `json:"ach_source_id"`
}