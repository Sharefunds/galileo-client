/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateLocationBody struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	// The name of the location. Pattern: Max 85 characters: letters, numbers, spaces, `_` (underscore), `'` (single quote), `-` (hyphen), `.` (period). All other characters will be removed. Example: `\"ABC Store 5\"`
	Name string `json:"name"`
	// First address line of the location. Pattern: 4&ndash;40 alphanumeric characters Example: `\"33 Maple Street\"`
	Address1 string `json:"address1"`
	// Second address line of the location. Pattern: Max 30 alphanumeric characters Example: `\"#4B\"`
	Address2 string `json:"address2,omitempty"`
	// City for the location. Pattern: Max 30 characters: letters, spaces, hyphen and period Example: `\"Salt Lake City\"`
	City string `json:"city"`
	// State or province for the location. Pattern: 2-character state or province abbreviation Example: `\"UT\"`
	State string `json:"state"`
	// Postal (ZIP) code for the location. Pattern: `12345`, `12345-1234`, `K1A-1A1` Example: `\"84121\"`
	PostalCode string `json:"postalCode"`
	// TCountry for the location. Three-digit UN M49 code, such as `840` for USA, `124` for Canada, `484` for Mexico, `170` for Colombia. Pattern: Three-digit country code Example: `\"840\"`
	CountryCode string `json:"countryCode,omitempty"`
	// Phone number of the location. Pattern: Exactly 10 digits, no hyphens or other characters Example: `\"8013656060\"`
	Phone string `json:"phone,omitempty"`
	// Identifier for the parent location. The first time you create a location, pass `0`. Pattern: Integer if `parentLocationType: 0`; max 15 characters if `parentLocationType: 1` Example: `\"5\"`
	ParentLocation string `json:"parentLocation"`
	// The type of location specified in `parentLocation`. If `parentLocation` is the top-level value from Galileo, then pass `0` for this parameter. * `0` &mdash; Galileo location ID * `1` &mdash; Partner location ID * `2` &mdash; Don't validate  Pattern: `0`, `1` or `2` Example: `0`
	ParentLocationType int32 `json:"parentLocationType"`
	// The provider-supplied identifier for a location, when `parentLocationType: 1`. Pattern: Max 15 alphanumeric characters Example: `\"abc-123\"`
	ProviderSpecifiedId string `json:"providerSpecifiedId,omitempty"`
	// Location status. Default `N`: * `C` &mdash; Closed * `N` &mdash; Active * `S` &mdash; Suspended * `n` &mdash; New  Pattern: Single character, case-sensitive Example: `\"N\"`
	Status string `json:"status,omitempty"`
	// The store or location type. Default `5`: * `1` &mdash; Corporate * `2` &mdash; Chain * `3` &mdash; Reseller * `4` &mdash; Region * `5` &mdash; Store  Pattern: Single digit or `null` Example: `\"5\"`
	StoreType string `json:"store_type,omitempty"`
}
