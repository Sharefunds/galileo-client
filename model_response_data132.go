/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData132 struct {
	// Identifier for the group
	GroupId string `json:"group_id"`
	// Name of the group
	GroupName string `json:"group_name"`
	// Identifier for the parent group
	ParentGroupId string `json:"parent_group_id"`
	// Maximum number of levels allowed below the root group
	MaxLevel int32 `json:"max_level"`
	// User-provided ID
	ExternalId string `json:"external_id"`
	// Operating name of the business, if different from `business_legal_name`
	DoingBusinessAs string `json:"doing_business_as"`
	// Legal name of the business
	BusinessLegalName string `json:"business_legal_name"`
	// Business phone number
	Phone string `json:"phone"`
	// Name of the primary contact
	PrimaryContactName string `json:"primary_contact_name"`
	// Email of the primary contact at the business
	PrimaryContactEmail string `json:"primary_contact_email"`
	// change time stamp
	ChangeTs string `json:"change_ts"`
	// Whether the root group supports Mastercard Smart Data.
	Smartdata string `json:"smartdata"`
	// phone country code
	PhoneCountryCode string `json:"phone_country_code"`
	// Employer id
	EmployerId string `json:"employer_id,omitempty"`
}
