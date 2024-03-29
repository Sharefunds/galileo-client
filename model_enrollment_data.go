/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EnrollmentData struct {
	// Account holder's first name
	FirstName string `json:"first_name"`
	// Account holder's middle name
	MiddleName string `json:"middle_name"`
	// Account holder's last name
	LastName string `json:"last_name"`
	// First line of the account holder's address
	Address1 string `json:"address_1"`
	// Second line of the account holder's address
	Address2 string `json:"address_2"`
	// Account holder's city
	City string `json:"city"`
	// Account holder's state
	State string `json:"state"`
	// Account holder's postal code
	PostalCode string `json:"postal_code"`
	// The ISO 3166 international standard for country codes. Identifies the account holder's country.
	CountryCode string `json:"country_code"`
	// Account holder's primary phone number
	PrimaryPhone string `json:"primary_phone"`
	// Account holder's mobile phone number
	MobilePhone string `json:"mobile_phone"`
	// Account holder's other phone number
	OtherPhone string `json:"other_phone"`
	// Account holder's user name (if Galileo hosts the enrollment website)
	WebUid string `json:"web_uid"`
	// A secret question associated with the account for access verification
	SecretQuestion string `json:"secret_question"`
	// The customer ID type
	Id string `json:"id"`
	// Description of the ID type specified in the `id` response field
	IdType int32 `json:"idType"`
	// The customer ID type for a secondary ID
	Id2 string `json:"id2"`
	// Description of the ID type specified in the `id2` response field
	IdType2 int32 `json:"idType2"`
}
