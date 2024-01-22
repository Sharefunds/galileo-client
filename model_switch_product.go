/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SwitchProduct struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin,omitempty"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey,omitempty"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId,omitempty"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	LogQueries int32 `json:"logQueries,omitempty"`
	// The <<glossary:PRN>>, <<glossary:PAN>> or <<glossary:CAD>> of the account. Pattern: PAN, PRN, or CAD Example: `\"074103447228\"`
	AccountNo string `json:"accountNo"`
	// The new product ID for the account. Pattern: Integer Example: `4444`
	ProdId int32 `json:"prodId"`
	// Controls whether to reissue the card. Default: `N`. Pattern: `Y` or `N` Example: `Y`
	DoReissue string `json:"doReissue,omitempty"`
	// Controls whether to generate a new PAN for the reissued card. Pattern: `Y` or `N` Example: `Y`
	NewPan string `json:"newPan,omitempty"`
	// Controls whether to generate a new expiry date for the reissued card. If `newPan: Y` then this value must be `Y`. Pattern: `Y` or `N` Example: `Y`
	NewExpiryDate string `json:"newExpiryDate,omitempty"`
	// Controls whether to emboss the reissued card. Pattern: `Y` or `N` Example: `Y`
	Emboss string `json:"emboss,omitempty"`
	// Specifies the status to set for the card that is being replaced. Status is set by the emboss process. Valid values are `C`, `D` or blank. Default: blank Pattern: `C`, `D` or blank Example: `D`
	OldCardStatus string `json:"oldCardStatus,omitempty"`
}