/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ModifyAchAccountBody struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	// The <<glossary:PRN>> or <<glossary:PAN>> of the account. Pattern: PAN or PRN  Example: `\"074103447228\"`
	AccountNo string `json:"accountNo"`
	// ACH account identifier (`ach_account_id`), as returned by the <a href=\"ref:post_addachaccount\" target=\"_blank\">Add ACH Account</a> or <a href=\"ref:post_getachaccounts\" target=\"_blank\">Get ACH Accounts</a> endpoint. Pattern: Positive integer Example: `354656`
	AchAccountId int32 `json:"achAccountId"`
	// The display name of the ACH account. Pattern: Alphanumeric string, max 22 characters Example: `\"Checking account\"`
	Name string `json:"name,omitempty"`
	// Type of the external bank account: * `C` &mdash; Checking * `S` &mdash; Savings * `M` &mdash; Money market  Pattern: `C`, `S`, or `M` Example: `\"C\"`
	Type_ string `json:"type,omitempty"`
	// External bank account number. Pattern: Max 22 digits Example: `\"4483434234348\"`
	AchAccountNo string `json:"achAccountNo,omitempty"`
	// Routing number for the external bank where `achAccountNo` is housed. Pattern: 9-digit routing number, including check digit Example: `\"124001545\"`
	AchRoutingNo string `json:"achRoutingNo,omitempty"`
	// Specifies the type of ACH account to modify: individual (`I`) or company (`C`). `C` is valid only when BTBPG is set at the program level and when the ACH account was created using the <a href=\"ref:post_addachaccountcorporate\" target=\"_blank\">Add ACH Account Corporate</a> endpoint. When this value is `I` then `firstName` and `lastName` are required. When this value is `C` then `companyName` is required. Pattern: `C` or `I` Example: `\"C\"`
	EntityType string `json:"entityType,omitempty"`
	// First name(s) of the external account holder. Required when `entityType: I`. This is a display name. Pattern: 1–40 alphanumeric characters Example: `\"Maricela Elena\"`
	FirstName string `json:"firstName,omitempty"`
	// Last name(s) of the external account holder. Required when `entityType: I`. This is a display name. Pattern: 1–40 alphanumeric characters Example: `\"Garcia Castro\"`
	LastName string `json:"lastName,omitempty"`
	// First name to include in the outgoing <<glossary:Nacha file>> instead of `firstName`. Pattern: 1–9 alphanumeric characters Example: `\"Maricela\"`
	FileFirstName string `json:"fileFirstName,omitempty"`
	// Last name to include in the outgoing Nacha file instead of `lastName`. Pattern: 1–12 alphanumeric characters Example: `\"Garcia\"`
	FileLastName string `json:"fileLastName,omitempty"`
	// Name of the external corporate account holder. Required when `entityType: C`.  This is a display name. Pattern: 1–40 alphanumeric characters Example: `\"Mountain Star Utilities\"`
	CompanyName string `json:"companyName,omitempty"`
	// Company name to include in the outgoing Nacha file instead of `companyName`. Pattern: 1–20 alphanumeric characters Example: `\"MountainStarUtil\"`
	FileCompanyName string `json:"fileCompanyName,omitempty"`
	// Unique location identifier (`location`) as returned by the <a href=\"ref:post_createlocation\" target=\"_blank\">Create Location</a> endpoint.  This value is also returned by the <a href=\"ref:post_getlocations\" target=\"_blank\">Get Locations</a> endpoint depending on the value of `locationType` when the location was created: * `0` or `2` &mdash; Returned in the `location_id` field * `1` &mdash; Returned in the `provider_specified_id` field  Pattern: Integer if `locationType: 0` or `locationType: 2`; max 15 characters if `locationType: 1` Example: `\"a455-3483\"`
	Location string `json:"location,omitempty"`
	// Type of ID in `location`: * `0` &mdash; Galileo location ID * `1` &mdash; Partner location ID * `2` &mdash; Don't validate  Pattern: `0` or `1` Example: `1`
	LocationType int32 `json:"locationType,omitempty"`
}
