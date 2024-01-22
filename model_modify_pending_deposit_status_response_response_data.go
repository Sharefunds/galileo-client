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
type ModifyPendingDepositStatusResponseResponseData struct {
	// A unique system generated ID number that identifies the deposit
	DepositTransactionId string `json:"deposit_transaction_id"`
	// 'P' and 'R' (Post and Return). Post = pending (ACH) deposit, Return = return the pending deposit
	ActionType string `json:"action_type"`
	// A = Approve; W = Watch; D = Decline
	CategoryType string `json:"category_type"`
	// COF=Questionable IAT Country; CRN= Unauthorized IAT Country; L=Large Xfer Amount > 4000; LRG=Large Xfer 4000; NAME=Name miss match; TAX=Tax
	CategoryCode string `json:"category_code"`
}
