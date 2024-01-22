/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResubmitFbEnrollmentBody struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	// The federal benefits enrollment ID (`fbe_status_id`) as returned by the <a href=\"ref:post_createfederalbenefitenrollment\" target=\"_blank\">Create Federal Benefit Enrollment</a> endpoint. Pattern: Integer Example: `1`
	FbeStatusId int32 `json:"fbeStatusId"`
	// Specifies whether the owner of the account in `accountNo` is a direct recipient or a beneficiary: * `0` &mdash; Direct recipient * `1` &mdash; Beneficiary  Pattern: `0` or `1` Example: `1`
	RecipientType int32 `json:"recipientType"`
	// Account holder's first name. <a href=\"ref:api-reference-special-characters\" target=\"_blank\">Special character support</a>. Pattern: 1&ndash;40 characters: letters (`A-Z`, `a-z`), spaces, hyphens (`-`) and single quotes (`'`) Example: `\"Ed\"`
	FirstName string `json:"firstName"`
	// Account holder's last name. <a href=\"ref:api-reference-special-characters\" target=\"_blank\">Special character support</a>. Pattern: 1&ndash;40 characters: letters (`A-Z`, `a-z`), spaces, hyphens (`-`) and single quotes (`'`) Example: `\"Harley\"`
	LastName string `json:"lastName"`
	// The Social Security number of the direct recipient of the disbursement. Pattern: 9-digit Social Security number, no hyphens Example: `\"123456789\"`
	Ssn string `json:"ssn"`
	// One of several defined agency types. Pattern: See the <a href=\"ref:api-reference-federal-benefit-enrollment-agency-types\" target=\"_blank\">Federal Benefit Enrollment Agency Types</a> table in Enumerations. Example: `\"SOCIAL SECURITY\"`
	AgencyType string `json:"agencyType"`
	// Identifier of the agent, for use in transaction reporting and tracking. Pattern: Integer Example: `109405`
	AgentId int32 `json:"agentId,omitempty"`
}
