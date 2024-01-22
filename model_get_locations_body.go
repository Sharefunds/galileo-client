/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GetLocationsBody struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	// String containing the search criteria. Must be of the type specified in `searchCriteriaType`. Pattern: Alphanumeric string Example: `\"5544\"`
	SearchCriteria string `json:"searchCriteria"`
	// Type of criteria that is specified in `searchCriteria`: * `1` &mdash; Location ID * `2` &mdash; Location name, full or partial * `3` &mdash; Location street address, full or partial  Pattern: One digit Example: `\"1\"`
	SearchCriteriaType string `json:"searchCriteriaType"`
}
