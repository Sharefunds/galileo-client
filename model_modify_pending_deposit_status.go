/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ModifyPendingDepositStatus struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin,omitempty"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey,omitempty"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId,omitempty"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	LogQueries int32 `json:"logQueries,omitempty"`
	// The <<glossary:PRN>> or <<glossary:PAN>> of the account. Pattern: PAN or PRN Example: `\"074103447228\"`
	AccountNo string `json:"accountNo,omitempty"`
	// Identifier supplied by the provider, which is not related to the Galileo system. This ID is stored in the Galileo system in association with this account and can be provided in the <<glossary:RDF>>s. Pattern: Max 30 alphanumeric characters Example: `\"553b45sbs\"`
	ExternalAccountId string `json:"externalAccountId,omitempty"`
	// The ACH transaction identifier (`ach_trans_id`), as returned by the <a href=\"ref:post_getpendingdeposits\" target=\"_blank\">Get Pending Deposits</a> endpoint. Pattern: Positive integer Example: `6844743`
	DepositTransactionId int32 `json:"depositTransactionId"`
	// Action to take on the pending deposit: * `P` &mdash; Post the deposit * `R` &mdash; Return the deposit to the sender. When returning the deposit, `retCode` is **required**  Pattern: Single character Example: `\"P\"`
	ActionType string `json:"actionType"`
	// Category to assign to the deposit. See <a href=\"ref:api-reference-deposit-category-codes\" target=\"_blank\">Deposit Category Codes</a> for valid values. Pattern: Three-letter code Example: `\"TAX\"`
	CategoryCode string `json:"categoryCode"`
	// Indicates the decision for future ACH deposits that match the program settings for the current deposit. Values are:   * `A` — Approve matching transactions. * `D` — Decline matching transactions. * `W` — Watch matching transactions and send for manual review.  Pattern: Must be `A`, `W` or `D` Example: `\"D\"`
	CategoryType string `json:"categoryType"`
	// Reason for returning the deposit. This parameter is **required** when `actionType: R`. See the <a href=\"ref:api-reference-modify-pending-deposit-status-return-codes\" target=\"_blank\">Return Codes</a> table for valid values. Pattern: 3-character code Example: `\"R02\"`
	RetCode string `json:"retCode,omitempty"`
}