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

type ExpireHoldBody struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	// Hold ID (`hold_id`) as returned by the <a href=\"ref:post_createhold\" target=\"_blank\">Create Hold</a> or <a href=\"ref:post_getholdhistory\" target=\"_blank\">Get Hold History</a> endpoint. Pattern: Positive integer Example: `453434`
	HoldId int32 `json:"holdId,omitempty"`
	// Date and time to expire the hold. Must be a date-time in the future. Leave this parameter empty to expire the hold immediately. Pattern: YYYY-MM-DD hh:mm:ss Example: `\"2017-01-01 13:00:00\"`
	ExpirationDateTime time.Time `json:"expirationDateTime,omitempty"`
}
