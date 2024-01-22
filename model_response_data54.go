/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData54 struct {
	// If `accountNo` contains a secondary account, these are other secondary accounts that are associated with the primary account in `parent_accounts`.
	SiblingAccounts []ResponseData54SiblingAccounts `json:"sibling_accounts"`
	// If `accountNo` contains a secondary account, this is the primary account that it is associated with.
	ParentAccounts []ResponseData54SiblingAccounts `json:"parent_accounts"`
	// Accounts that belong to the same customer.
	SelfAccounts []ResponseData54SiblingAccounts `json:"self_accounts"`
	// Accounts that are secondary accounts to the primary account in `accountNo`.
	ChildAccounts []ResponseData54SiblingAccounts `json:"child_accounts"`
}