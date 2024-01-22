/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdatePaymentBody struct {
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
	// The payment ID (`payment_trans_id`) as returned by the <a href=\"ref:post_createpayment\" target=\"_blank\">Create Payment</a> endpoint or `pmt_id`  as returned by the <a href=\"ref:post_getpaymenthistory\" target=\"_blank\">Get Payment History</a> endpoint. Pattern: Positive integer Example: `4234888`
	PmtId int32 `json:"pmtId"`
	// Number of days to hold a payment before processing. If set to `0`, the payment will be posted the next time the internal payment process runs. Pattern: Integer value of `0` or greater Example: `0`
	HoldDays int32 `json:"holdDays"`
}