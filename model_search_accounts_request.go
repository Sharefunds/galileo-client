/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SearchAccountsRequest struct {
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
	// Account holder's first name. <a href=\"ref:api-reference-special-characters\" target=\"_blank\">Special character support</a>. Pattern: 1&ndash;40 characters: letters (`A-Z`, `a-z`), spaces, hyphens (`-`) and single quotes (`'`) Example: `\"Ed\"`
	FirstName string `json:"firstName,omitempty"`
	// Account holder's middle name. <a href=\"ref:api-reference-special-characters\" target=\"_blank\">Special character support</a>. Pattern: 1&ndash;40 characters: letters (`A-Z`, `a-z`), spaces, hyphens (`-`) and single quotes (`'`) Example: `\"W\"`
	MiddleName string `json:"middleName,omitempty"`
	// Account holder's last name. <a href=\"ref:api-reference-special-characters\" target=\"_blank\">Special character support</a>. Pattern: 1&ndash;40 characters: letters (`A-Z`, `a-z`), spaces, hyphens (`-`) and single quotes (`'`) Example: `\"Harley\"`
	LastName string `json:"lastName,omitempty"`
	// Account holder's birth date. This date is compared with the DOB product parameter when there is a minimum age requirement on the account. Pattern: YYYY-MM-DD Example: `\"1980-01-01\"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	// Account holder's postal code (ZIP code). Pattern: `12345`, `12345-1234`, or `K1A-1A1` Example: `\"84121\"`
	PostalCode string `json:"postalCode,omitempty"`
	// Account holder's primary phone number. Pattern: Exactly 10 digits, no hyphens or other characters Example: `\"8013656050\"`
	PrimaryPhone string `json:"primaryPhone,omitempty"`
	// Account holder's secondary phone number. Pattern: Exactly 10 digits, no hyphens or other characters Example: `\"8013656050\"`
	OtherPhone string `json:"otherPhone,omitempty"`
	// Account holder's mobile phone number Pattern: Exactly 10 digits, no hyphens or other characters Example: `\"8013656050\"`
	MobilePhone string `json:"mobilePhone,omitempty"`
	// Account holder's email. The Python implementation uses Marshmallow validation for this field. For the PHP implementation, contact Galileo. Pattern: Email address, 3&ndash;63 characters Example: `\"user@fakedomain.com\"`
	Email string `json:"email,omitempty"`
	// Data supplied by the provider, from sources external to the Galileo system. This data will be stored in the Galileo system in association with this account. The most common use of this parameter is to track affiliate-marketing traffic. The data in this field can be added to an <<glossary:RDF>>. Pattern: Max 50 alphanumeric characters and spaces Example: `\"a4434gg44\"`
	UserData string `json:"userData,omitempty"`
	// The `externalAccountId` that was passed in the account-creation call. Pattern: Max 30 alphanumeric characters Example: `\"553b45sbs\"`
	Eextid string `json:"eextid,omitempty"`
	// The maximum number of records per page to be returned. Pattern: Positive integer `1-99999` Example: `100`
	RecordCnt int32 `json:"recordCnt,omitempty"`
	// The number of the page to retrieve. Pattern: Integer value of `1` or greater Example: `3`
	Page int32 `json:"page,omitempty"`
	// Account holder's mobile phone country code. See <a href=\"ref:api-reference-phone-validation\" target=\"_blank\">Phone Validation</a> for valid values. Pattern: A plus sign followed by 1 - 3 digits  Example: `\"+1\"`
	MobilePhoneCountryCode string `json:"mobilePhoneCountryCode,omitempty"`
}
