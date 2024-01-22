/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AccountsResponse struct {
	// The identifier for the product that this account belongs to.
	ProductId string `json:"product_id"`
	// Whether the account is active: `Y` or `N`
	Active string `json:"active"`
	// Balance ID for the account.
	GalileoAccountNumber string `json:"galileo_account_number"`
	// The status of the account. See <a href=\"ref:api-reference-account-statuses\" target=\"_blank\">Account Statuses</a> for valid values.
	Status string `json:"status"`
	// The payment reference number of the account.
	PmtRefNo string `json:"pmt_ref_no,omitempty"`
}
