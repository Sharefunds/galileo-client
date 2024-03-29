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

type Application struct {
	// Identifier for the enrollment application
	AppId string `json:"appId"`
	// Timestamp for the application transaction
	AppDate time.Time `json:"appDate"`
	// Identifier for the product associated with the account
	ProdId int32 `json:"prodId"`
	// The customer ID type
	Id string `json:"id"`
	// Description of the ID type specified in the `id` response field
	IdType int32 `json:"idType"`
	// The customer ID type for a secondary ID
	Id2 string `json:"id2"`
	// Description of the ID type specified in the `id2` response field
	IdType2 int32 `json:"idType2"`
	// The account holder's language and localization preference
	Locale string `json:"locale"`
	// Account holder's first name
	FirstName string `json:"firstName"`
	// Account holder's middle name
	MiddleName string `json:"middleName"`
	// Account holder's last name
	LastName string `json:"lastName"`
	// Account holder's date of birth
	DateOfBirth string `json:"dateOfBirth"`
	// First line of the account holder's address
	Address1 string `json:"address1"`
	// Second line of the account holder's address
	Address2 string `json:"address2"`
	// Account holder's city
	City string `json:"city"`
	// Account holder's state
	State string `json:"state"`
	// Account holder's postal code
	PostalCode string `json:"postalCode"`
	// Indicates whether to use express mail when shipping the card for this account
	ExpressMail string `json:"expressMail"`
	// Account holder's primary phone number
	PrimaryPhone string `json:"primaryPhone"`
	// Account holder's mobile phone number
	MobilePhone string `json:"mobilePhone"`
	// Identifier for the mobile phone service provider
	MobileCarrierId string `json:"mobileCarrierId"`
	// Account holder's other phone number
	OtherPhone string `json:"otherPhone"`
	// Account holder's email address
	Email string `json:"email"`
	// Data associated with the account that you supply. This data is external to the Galileo system.
	UserData string `json:"userData"`
	// Account holder's gross income for one month
	MonthlyIncome string `json:"monthlyIncome"`
	// Account holder's total liabilities for one month
	MonthlyLiab string `json:"monthlyLiab"`
}
