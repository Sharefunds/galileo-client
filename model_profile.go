/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Profile struct {
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
	// Business name on the account
	BusinessName string `json:"business_name"`
	// Account holder's city
	City string `json:"city"`
	// Account holder's state
	State string `json:"state"`
	// Account holder's postal code
	PostalCode string `json:"postal_code"`
	// The ISO 3166 international standard for country codes. Identifies the account holder's country.
	CountryCode string `json:"country_code"`
	// Account holder's home phone number
	HomePhone string `json:"home_phone"`
	// Account holder's mobile phone number
	MobilePhone string `json:"mobile_phone"`
	// Account holder's email address
	Email string `json:"email"`
	// Account holder's date of birth
	Dob string `json:"dob"`
	ShipToAddress *ProfileShipToAddress `json:"ship_to_address"`
	// Indicates whether to use express mail for shipping
	ExpressMail string `json:"express_mail,omitempty"`
	// Account holder's occupation
	Occupation string `json:"occupation,omitempty"`
	// Account holder's income source
	IncomeSource string `json:"income_source,omitempty"`
}
