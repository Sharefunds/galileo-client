/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateAccountRequest struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin,omitempty"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey,omitempty"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId,omitempty"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	LogQueries    int32  `json:"logQueries,omitempty"`
	// The <<glossary:PRN>> or <<glossary:PAN>> of the account. Pattern: PAN or PRN Example: `\"074103447228\"`
	AccountNo string `json:"accountNo,omitempty"`
	// Specifies the type of primary identifier in the `id` parameter.  This parameter is **required** when `id` is populated. See the **ID** column in the <a href=\"ref:api-reference-customer-id-types\" target=\"_blank\">Customer ID Types</a> table for valid values. Pattern: 1- or 2-digit integer Example: `2`
	IdType int32 `json:"idType,omitempty"`
	// The value of the primary identifier that is specified in `idType`. Product settings determine whether this ID should be unique across the program or product. This parameter is **required** when using Galileo's integrated <<glossary:CIP>> solution. Pattern: As specified in the **Layout** column of the <a href=\"ref:api-reference-customer-id-types\" target=\"_blank\">Customer ID Types</a> enumeration Example: `\"123456789\"`
	Id string `json:"id,omitempty"`
	// Specifies the type of secondary identifier in the `id2` parameter. This parameter is **required** when `id2` is populated. See the **ID** column in the <a href=\"ref:api-reference-customer-id-types\" target=\"_blank\">Customer ID Types</a> table for valid values. Pattern: 1- or 2-digit integer Example: `1`
	IdType2 int32 `json:"idType2,omitempty"`
	// The value of the secondary identifier that is specified in `idType2`. Pattern: As specified in the **Layout** column of the <a href=\"ref:api-reference-customer-id-types\" target=\"_blank\">Customer ID Types</a> enumeration Example: `\"123456789012|UT|12/25/2020\"`
	Id2 string `json:"id2,omitempty"`
	// Type of ID in `location`: * `0` &mdash; Galileo location ID * `1` &mdash; Partner location ID * `2` &mdash; Don't validate  Pattern: `0`, `1`, or `2` Example: `0`
	LocationType int32 `json:"locationType,omitempty"`
	// Unique location identifier (`location`) as returned by the <a href=\"ref:post_createlocation\" target=\"_blank\">Create Location</a> endpoint.  This value is also returned by the <a href=\"ref:post_getlocations\" target=\"_blank\">Get Locations</a> endpoint depending on the value of `locationType` when the location was created: * `0` or `2` &mdash; Returned in the `location_id` field * `1` &mdash; Returned in the `provider_specified_id` field  Pattern: Integer if `locationType: 0` or `locationType: 2`; max 15 characters if `locationType: 1` Example: `\"a455-3483\"`
	Location string `json:"location,omitempty"`
	// The account holder language and localization preference. Default is `en_US`. Valid values are `en_US`, `es_US`, `fr_CA`, `en_CA`, `es_MX`, `en_MX`. If the account holder address is outside the U.S., pass a non `_US` value for this parameter to disable U.S. address validation. Pattern: Letters with underscore Example: `\"en_US\"`
	Locale string `json:"locale,omitempty"`
	// Identifier supplied by the provider, which is not related to the Galileo system. This ID is stored in the Galileo system in association with this account and can be provided in the <<glossary:RDF>>s. Pattern: Max 30 alphanumeric characters Example: `\"553b45sbs\"`
	ExternalAccountId string `json:"externalAccountId,omitempty"`
	// Account holder's first name. <a href=\"ref:api-reference-special-characters\" target=\"_blank\">Special character support</a>. Pattern: 1&ndash;40 characters: letters (`A-Z`, `a-z`), spaces, hyphens (`-`) and single quotes (`'`) Example: `\"Ed\"`
	FirstName string `json:"firstName,omitempty"`
	// Account holder's middle name. <a href=\"ref:api-reference-special-characters\" target=\"_blank\">Special character support</a>. Pattern: 1&ndash;40 characters: letters (`A-Z`, `a-z`), spaces, hyphens (`-`) and single quotes (`'`) Example: `\"W\"`
	MiddleName string `json:"middleName,omitempty"`
	// Account holder's last name. <a href=\"ref:api-reference-special-characters\" target=\"_blank\">Special character support</a>. Pattern: 1&ndash;40 characters: letters (`A-Z`, `a-z`), spaces, hyphens (`-`) and single quotes (`'`) Example: `\"Harley\"`
	LastName string `json:"lastName,omitempty"`
	// Account holder's birth date. This date is compared with the DOB product parameter when there is a minimum age requirement on the account. Pattern: YYYY-MM-DD Example: `\"1980-01-01\"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	// Account holder's first address line. Cannot be a P.O. box. <a href=\"ref:api-reference-address1-validation\" target=\"_blank\">Validation rules for `address1`</a>. Pattern: 4&ndash;40 alphanumeric characters Example: `\"33 Maple Street\"`
	Address1 string `json:"address1,omitempty"`
	// Account holder's second address line. <a href=\"ref:api-reference-special-characters#address-characters\" target=\"_blank\">Character set</a> for address lines.  Pattern: Max 40 characters Example: `\"#4B\"`
	Address2 string `json:"address2,omitempty"`
	// Account holder's third address line. <a href=\"ref:api-reference-special-characters#address-characters\" target=\"_blank\">Character set</a> for address lines.  Pattern: Max 30 characters Example: `\"Carrera Central\"`
	Address3 string `json:"address3,omitempty"`
	// Account holder's fourth address line. <a href=\"ref:api-reference-special-characters#address-characters\" target=\"_blank\">Character set</a> for address lines.  Pattern: Max 30 characters Example: `\"Barrio El Alhambra\"`
	Address4 string `json:"address4,omitempty"`
	// Account holder's fifth address line. <a href=\"ref:api-reference-special-characters#address-characters\" target=\"_blank\">Character set</a> for address lines.  Pattern: Max 30 characters Example: `\"Piso 3\"`
	Address5 string `json:"address5,omitempty"`
	// Account holder's city. Pattern: Max 30 characters: letters, spaces, hyphen and period Example: `\"Salt Lake City\"`
	City string `json:"city,omitempty"`
	// Account holder's state or province. Pattern: 2,3-character state or province abbreviation Example: `\"UT\"`
	State string `json:"state,omitempty"`
	// Account holder's postal code (ZIP code). Pattern: `1234`, `12345`, `12345-1234`, or `K1A-1A1` Example: `\"84121\"`
	PostalCode string `json:"postalCode,omitempty"`
	// Account holder's country. Three-digit UN M49 code, such as `840` for USA, `124` for Canada, `484` for Mexico, `170` for Colombia. Pattern: 3-digit number Example: `\"840\"`
	CountryCode string `json:"countryCode,omitempty"`
	// Account holder's primary phone number. Pattern: Exactly 10 digits, no hyphens or other characters Example: `\"8013656060\"`
	PrimaryPhone string `json:"primaryPhone,omitempty"`
	// Account holder's primary phone number. Pattern: Exactly 10 digits, no hyphens or other characters Example: `\"8013656060\"`
	OtherPhone string `json:"otherPhone,omitempty"`
	// Account holder's primary phone number. Pattern: Exactly 10 digits, no hyphens or other characters Example: `\"8013656060\"`
	MobilePhone string `json:"mobilePhone,omitempty"`
	// Account holder's mobile carrier. Configurable list. Pattern: Configurable list Example: `\"8\"`
	MobileCarrierId string `json:"mobileCarrierId,omitempty"`
	// Account holder's email. The Python implementation uses Marshmallow validation for this field. For the PHP implementation, contact Galileo. Pattern: Email address, 3&ndash;63 characters Example: `\"user@fakedomain.com\"`
	Email string `json:"email,omitempty"`
	// Account holder's username for the Galileo-hosted website. Must be unique across the program. Pattern: Alphanumeric, with first character a letter; case insensitive Example: `\"eharley\"`
	WebUid string `json:"webUid,omitempty"`
	// Account holder's password for the Galileo-hosted website. Pattern: Min 8 characters: must include upper-case character, lower-case character, and a number Example: `\"4j9KH3kkdh\"`
	WebPwd string `json:"webPwd,omitempty"`
	// Account holder's secret question for the Galileo-hosted website. Pattern: Max 50 characters: letters, spaces, `?` Example: `\"What was the name of your first pet?\"`
	SecretQuestion string `json:"secretQuestion,omitempty"`
	// Account holder's answer to `secretQuestion`. Pattern: Max 50 characters: letters, spaces, `?` Example: `\"Larry\"`
	SecretAnswer string `json:"secretAnswer,omitempty"`
	// Account holder's employer name or income source. Pattern: Max 60 characters: alphanumeric, space, underscore, hyphen, period, `@`, `&` and comma. Example: `\"Kroger Food & Drug\"`
	IncomeSource string `json:"incomeSource,omitempty"`
	// Account holder's occupation. Pattern: Max 60 characters: alphanumeric, space, underscore, hyphen, period, `@`, `&` and comma. Example: `\"Project Manager\"`
	Occupation string `json:"occupation,omitempty"`
	// `Y` indicates that mail has been returned to sender. Pattern: `Y` or `N` Example: `\"Y\"`
	MailBounced string `json:"mailBounced,omitempty"`
	// First ship-to address line. Pattern: Max 40 characters Example: `\"33 business parkway\"`
	ShipToAddress1 string `json:"shipToAddress1,omitempty"`
	// Second ship-to address line. Pattern: Max 30 characters Example: `\"Suite 400\"`
	ShipToAddress2 string `json:"shipToAddress2,omitempty"`
	// Ship-to city. Pattern: Max 30 characters Example: `\"Salt Lake City\"`
	ShipToCity string `json:"shipToCity,omitempty"`
	// Ship-to state. Pattern: 2,3-character state or province abbreviation Example: `\"UT\"`
	ShipToState string `json:"shipToState,omitempty"`
	// Ship-to ZIP or postal code. Pattern: `1234`, `12345`, `12345-1234`, `K1A-1A1` Example: `\"841213333\"`
	ShipToPostalCode string `json:"shipToPostalCode,omitempty"`
	// Ship-to country. Three-digit UN M49 country code, such as `840` for USA, `124` for Canada, `484` for Mexico, `170` for Colombia. Pattern: 3-digit number Example: `\"840\"`
	ShipToCountryCode string `json:"shipToCountryCode,omitempty"`
	// Pass `1` to always ship cards to the ship-to address instead of one time only. Pattern: `0` or `1` Example: `\"1\"`
	ShipToAddressPermanent string `json:"shipToAddressPermanent,omitempty"`
	// Second line to emboss on the card. Pattern: Max 28 alphanumeric characters, including hyphen (`-`) Example: `\"second emboss line\"`
	EmbossLine2 string `json:"embossLine2,omitempty"`
	// Account holder's business name. This field is **required** if `prodId` is a business product. Pattern: 2-150 characters. See the <a href=\"ref:api-reference-special-characters#businessname-parameter\" target=\"_blank\">supported character set</a>. Example: `\"Ed%26Ted\"`
	BusinessName string `json:"businessName,omitempty"`
	// Account holder's mobile phone country code. Pattern: A plus sign followed by 1–3 digits. See <a href=\"ref:api-reference-phone-validation\" target=\"_blank\">Phone Validation</a> for valid values.  Example: `\"+1\"`
	MobilePhoneCountryCode string `json:"mobilePhoneCountryCode,omitempty"`
}
