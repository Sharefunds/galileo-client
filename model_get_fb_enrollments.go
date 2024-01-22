/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GetFbEnrollments struct {
	// Web service username, as provided by Galileo. Pattern: Max 50 characters Example: `\"AbC123-9999\"`
	ApiLogin string `json:"apiLogin,omitempty"`
	// Web service password, as provided by Galileo. Pattern: Max 15 characters Example: `\"4sb62fh6w4h7w34g\"`
	ApiTransKey string `json:"apiTransKey,omitempty"`
	// Galileo-issued provider identifier. Pattern: Max 10 digits Example: `9999`
	ProviderId int32 `json:"providerId,omitempty"`
	// A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for <a href=\"ref:idempotency\" target=\"_blank\">idempotency</a>. Pattern: 60 characters or less Example: `\"9845dk-39fdk3fj3-4483483478\"`
	TransactionId string `json:"transactionId"`
	LogQueries int32 `json:"logQueries,omitempty"`
	// The <<glossary:PRN>> or <<glossary:PAN>> of the account. Pattern: PRN or PAN Example: `\"074103447228\"`
	AccountNo string `json:"accountNo"`
	// The PRN of the federal benefits enrollment <<glossary:DDA>>. Pattern: PRN Example: `\"074103447228\"`
	FbeDda string `json:"fbeDda,omitempty"`
	// The type of agency that is providing the benefit. See the <a href=\"ref:api-reference-federal-benefit-enrollment-agency-types\" target=\"_blank\">Federal Benefit Enrollment Agency Types</a> enumeration for valid values. Pattern: String Example: `\"SOCIAL SECURITY\"`
	AgencyType string `json:"agencyType,omitempty"`
	// Specifies the state of the federal benefits enrollment. See the <a href=\"ref:api-reference-federal-benefit-enrollment-statuses\" target=\"_blank\">Federal Benefit Enrollment Statuses</a> enumeration for valid values. Pattern: One character Example: `\"B\"`
	Status string `json:"status,omitempty"`
	// The Social Security number of the beneficiary of the disbursement. Pattern: 9-digit Social Security number, no hyphens Example: `\"123456789\"`
	BeneficiarySsn string `json:"beneficiarySsn,omitempty"`
	// Account holder's first name. <a href=\"ref:api-reference-special-characters\" target=\"_blank\">Special character support</a>. Pattern: 1&ndash;40 characters: letters (`A-Z`, `a-z`), spaces, hyphens (`-`) and single quotes (`'`) Example: `\"Ed\"`
	BeneficiaryFirstName string `json:"beneficiaryFirstName,omitempty"`
	// Cardholder's last name Account holder's last name. <a href=\"ref:api-reference-special-characters\" target=\"_blank\">Special character support</a>. Pattern: 1&ndash;40 characters: letters (`A-Z`, `a-z`), spaces, hyphens (`-`) and single quotes (`'`) Example: `\"Harley\"`
	BeneficiaryLastName string `json:"beneficiaryLastName,omitempty"`
	// Date enrollment started. Pattern: YYYY-MM-DD Example: `\"2016-01-01\"`
	EnrolledFrom string `json:"enrolledFrom,omitempty"`
	// Date enrolled to. Pattern: YYYY-MM-DD Example: `\"2016-01-01\"`
	EnrolledTo string `json:"enrolledTo,omitempty"`
	// Date submitted from. Pattern: YYYY-MM-DD Example: `\"2016-01-01\"`
	SubmittedFrom string `json:"submittedFrom,omitempty"`
	// Date submitted to. Pattern: YYYY-MM-DD Example: `\"2016-01-01\"`
	SubmittedTo string `json:"submittedTo,omitempty"`
	// Date noted from. Pattern: YYYY-MM-DD Example: `\"2016-01-01\"`
	NotedFrom string `json:"notedFrom,omitempty"`
	// Date noted to. Pattern: YYYY-MM-DD Example: `\"2016-01-01\"`
	NotedTo string `json:"notedTo,omitempty"`
}
