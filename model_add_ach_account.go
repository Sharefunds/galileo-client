/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AddAchAccount struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin,omitempty"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey,omitempty"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId,omitempty"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	LogQueries int32 `json:"logQueries,omitempty"`
	// The <<glossary:PRN>> or <<glossary:PAN>> of the account. Pattern: PAN or PRN  Example: `\"074103447228\"`
	AccountNo string `json:"accountNo"`
	// Description of the external account. **Required** when `processorToken` is not populated. Pattern: Max 22 alphanumeric characters Example: `\"Checking account\"`
	Name string `json:"name,omitempty"`
	// Type of the external bank account. **Required** when `processorToken` is not populated: * `C` &mdash; Checking * `S` &mdash; Savings * `M` &mdash; Money market  Pattern: `C`, `S`, or `M` Example: `\"C\"`
	Type_ string `json:"type,omitempty"`
	// External bank account number. **Required** when `processorToken` is not populated. Pattern: Max 22 digits Example: `\"4483434234348\"`
	AchAccountNo string `json:"achAccountNo,omitempty"`
	// Routing number for the external bank where `achAccountNo` is housed. **Required** when `processorToken` is not populated. Pattern: 9-digit routing number, including check digit Example: `\"124001545\"`
	AchRoutingNo string `json:"achRoutingNo,omitempty"`
	// Unique location identifier (`location`) as returned by the <a href=\"ref:post_createlocation\" target=\"_blank\">Create Location</a> endpoint.  This value is also returned by the <a href=\"ref:post_getlocations\" target=\"_blank\">Get Locations</a> endpoint depending on the value of `locationType` when the location was created: * `0` or `2` &mdash; Returned in the `location_id` field * `1` &mdash; Returned in the `provider_specified_id` field  Pattern: Integer if `locationType: 0` or `locationType: 2`; max 15 characters if `locationType: 1` Example: `\"a455-3483\"`
	Location string `json:"location,omitempty"`
	// Type of ID in `location`: * `0` &mdash; Galileo location ID * `1` &mdash; Partner location ID * `2` &mdash; Don't validate  Pattern: `0` or `1` Example: `1`
	LocationType int32 `json:"locationType,omitempty"`
	// Obtained from Plaid if using Plaid integration. When passing this token it is not necessary to pass `achAccountNumber`, `achRoutingNumber`, `name`, or `type`. If you are not using Plaid, do not send this parameter with an empty or `null` value. Pattern: String Example: `\"processor-production-35cd43b-adfc-d8aa-b331-c9ba0fdha881\"`
	ProcessorToken string `json:"processorToken,omitempty"`
	// If `processorToken` is populated, pass `1` to call the Plaid identity endpoint for account-holder verification. Pattern: `0` or `1` Example: `0`
	VerifyIdentity int32 `json:"verifyIdentity,omitempty"`
	// First name(s) of the external account holder. This is a display name. Pattern: 1–40 alphanumeric characters Example: `\"Maricela Elena\"`
	FirstName string `json:"firstName,omitempty"`
	// Last name(s) of the external account holder. This is a display name. Pattern: 1–40 alphanumeric characters Example: `\"Garcia Castro\"`
	LastName string `json:"lastName,omitempty"`
	// First name to include in the outgoing <<glossary:Nacha file>> instead of `firstName`. Pattern: 1–9 alphanumeric characters Example: `\"Maricela\"`
	FileFirstName string `json:"fileFirstName,omitempty"`
	// Last name to include in the outgoing Nacha file instead of `lastName`. Pattern: 1–12 alphanumeric characters  Example: `\"Garcia\"`
	FileLastName string `json:"fileLastName,omitempty"`
}
