/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Category information for an ACH request to move funds into or out of the customer's account
type DepositCategories struct {
	// ACH source identifier used in Galileo's system
	AchSourceId string `json:"ach_source_id"`
	// Category code for the deposit. See <a href=\"ref:api-reference-deposit-category-codes\" target=\"_blank\">Deposit Category Codes</a> for valid values.
	CategoryCode string `json:"category_code"`
	// Description of the record
	Description string `json:"description"`
	// Indicates when the `source_status` was last updated
	LastUpdated string `json:"last_updated,omitempty"`
	// Indicates who last updated the `source_status`
	LastUpdatedBy string `json:"last_updated_by,omitempty"`
	// The source status for the ACH request. Possible values are `APPROVE`, `WATCH`, or `DECLINE`
	SourceStatus string `json:"source_status,omitempty"`
	// The deposit status. See <a href=\"ref:api-reference-deposit-status-codes\">Deposit Status Codes</a> for possible values.
	Status string `json:"status,omitempty"`
}
