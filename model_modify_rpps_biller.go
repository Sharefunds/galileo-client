/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type ModifyRppsBiller struct {
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
	// Identifier for the biller (`biller_id`) as returned by the <a href=\"ref:post_addrppsbiller\" target=\"_blank\">Add RPPS Biller</a> or <a href=\"ref:post_getbillers\" target=\"_blank\">Get Billers</a> endpoint. Do not use the `rpps_biller_id` from Search Biller Directory. Pattern: Integer Example: `2982`
	BillerId int32 `json:"billerId"`
	// Frequency of the bill payment: * `O` &mdash; One time * `W` &mdash; Weekly * `M` &mdash; Monthly * `Q` &mdash; Quarterly * `Y` &mdash; Yearly  If this value is not `O` then `nextDate` and `endDate` are **required**. Pattern: One letter Example: `\"W\"`
	FrequencyType string `json:"frequencyType"`
	// The next date that the payment is scheduled. Pattern: YYYY-MM-DD Example: `\"2015-02-04\"`
	NextDate time.Time `json:"nextDate,omitempty"`
	// The last date that the payment is scheduled. Can be up to five years in the future, so if today is 20 Jan 2020, this parameter can be no later than 20 Jan 2025. However, if today's date is a leap day, such as 29 Feb 2020, the latest date can be 1 Mar 2025. Pattern: YYYY-MM-DD Example: `\"2015-02-04\"`
	EndDate time.Time `json:"endDate,omitempty"`
	// Currency amount as a whole or decimal amount. Pattern: Positive integer or decimal number Example: `100.00`, `100`, or `100.73`
	Amount float32 `json:"amount,omitempty"`
}
