/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SetAlertsBlackout struct {
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
	// Starting time for alerts to be sent. Set to whole hours only. Pattern: HH`:00` Example: `\"07:00\"`
	BeginTime string `json:"beginTime"`
	// Time of day alerts must end. Set to whole hours only. Pattern: HH`:00` Example: `\"22:00\"`
	EndTime string `json:"endTime"`
	// Time zone offset from Arizona time for `beginTime` and `endTime`: * `-3` &mdash; Hawaii time zone * `-2` &mdash; Alaska time zone * `-1` &mdash; Pacific time zone * `0` &mdash; Mountain time zone * `1` &mdash; Central time zone * `2` &mdash; Eastern time zone * `3` &mdash; Atlantic time zone  Pattern:  Integer between `-3` and `3` Example: `-1`
	TimeZone string `json:"timeZone"`
	// Specifies whether to observe daylight savings time. Pattern: `Y` or `N` Example: `\"Y\"`
	DaylightSavings string `json:"daylightSavings"`
}