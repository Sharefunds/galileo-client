/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type VerifyChLogin struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin,omitempty"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey,omitempty"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId,omitempty"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	LogQueries int32 `json:"logQueries,omitempty"`
	// Cardholder website username Pattern: Must be unique. Start with a letter. Only letters and numbers. Case insensitive. Example: `\"4j9KH3kkdh\"`
	WebUid string `json:"webUid"`
	// Cardholder website password Pattern: At least 8 characters and have an uppercase character, lower case character and a number. Example: `\"eharley\"`
	WebPwd string `json:"webPwd"`
}
