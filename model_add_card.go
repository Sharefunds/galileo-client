/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AddCard struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin,omitempty"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey,omitempty"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId,omitempty"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	LogQueries int32 `json:"logQueries,omitempty"`
	// The <<glossary:PRN>> or <<glossary:PAN>> of the account. Do not use the <<glossary:CAD>>. Pattern: PAN or PRN Example: `\"074103447228\"`
	AccountNo string `json:"accountNo"`
	// Product ID of the card to add. Pattern: Integer Example: `\"501\"`
	ProdId int32 `json:"prodId"`
	// The <<glossary:PAN>>, <<glossary:PRN>> or package ID of the instant-issue card to add. Pass this parameter only for instant-issue cards. Pattern: PAN or PRN or Package ID Example: `\"123456789012\"`
	NewAccountNo string `json:"newAccountNo,omitempty"`
	// Specifies whether the secondary account will share the balance of the primary account. This parameter is **required** when `primaryAccount` is populated. Do not set this parameter to `1` when `primaryAccount` is not populated. * `0` &mdash; Do not share the balance * `1` &mdash; Share the balance  Pattern: One digit Example: `1`
	SharedBalance string `json:"sharedBalance,omitempty"`
	// Credit limit for card-based spending control. This is available only for credit programs. Pattern: Monetary amount greater than 0. Example: `500.25`
	CreditLimit float64 `json:"creditLimit,omitempty"`
	// Spending control to limit the card to one purchase transaction. This is available only for credit programs. Pattern: `Y` or `N` Example: `\"Y\"`
	SingleUse string `json:"singleUse,omitempty"`
}
