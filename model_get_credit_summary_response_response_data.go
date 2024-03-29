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
type GetCreditSummaryResponseResponseData struct {
	// The credit amount currently available
	AvailableCredit string `json:"available_credit,omitempty"`
	// The outstanding balance on the account
	OutstandingBalance float32 `json:"outstanding_balance,omitempty"`
	Billing *ResponseData86Billing `json:"billing,omitempty"`
	CreditLimit *ResponseData86CreditLimit `json:"credit_limit,omitempty"`
}
