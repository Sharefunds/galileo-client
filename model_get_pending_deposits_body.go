/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GetPendingDepositsBody struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	// The <<glossary:PRN>> or <<glossary:PAN>> of the account. Pattern: PAN or PRN  Example: `\"074103447228\"`
	AccountNo string `json:"accountNo,omitempty"`
	// The beginning date for the date range. Pattern: YYYY-MM-DD Example: `\"2016-01-01\"`
	StartDate string `json:"startDate,omitempty"`
	// The end date for the date range. Must be equal to or later than `startDate`. Pattern: YYYY-MM-DD Example: `\"2016-01-01\"`
	EndDate string `json:"endDate,omitempty"`
	// The maximum number of records per page to be returned. Pattern: P Positive integer `1-99999` Example: `100`
	RecordCnt int32 `json:"recordCnt,omitempty"`
	// The number of the page to retrieve. Pattern: Integer value of `1` or greater Example: `3`
	Page int32 `json:"page,omitempty"`
}
