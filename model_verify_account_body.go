/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type VerifyAccountBody struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	// The <<glossary:PRN>> or <<glossary:PAN>> of the account. Pattern: PRN or PAN Example: `\"074103447228\"`
	AccountNo string `json:"accountNo"`
	// Load type (otype) of the load to verify. Use values from your Galileo-supplied list of load types. If no `loadType` is specified, the default type `RL` is used. Pattern: 2 alphanumeric characters Example: `\"RL\"`
	LoadType string `json:"loadType,omitempty"`
	// Whether to return data for related accounts. See <a href=\"doc:about-accounts#getting-related-account-information\" target=\"_blank\">Getting related account information</a> for instructions on using this parameter. Default is `1`. Pattern: `0` or `1` Example: `0`
	IncludeRelated int32 `json:"includeRelated,omitempty"`
}
