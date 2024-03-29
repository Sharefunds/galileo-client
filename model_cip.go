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

type Cip struct {
	// Status of the CIP process
	Status string `json:"status"`
	// List of model objects
	ModelResults []CipModelResults `json:"model_results"`
	// The date and time CIP was run
	DateTime time.Time `json:"date_time"`
	// Account holder's first name
	FirstName string `json:"first_name,omitempty"`
	// Account holder's middle name
	MiddleName string `json:"middle_name,omitempty"`
	// Account holder's last name
	LastName string `json:"last_name,omitempty"`
	// First line of the account holder's address
	Address1 string `json:"address_1,omitempty"`
	// Second line of the account holder's address
	Address2 string `json:"address_2,omitempty"`
	// Account holder's city
	City string `json:"city,omitempty"`
	// Account holder's state
	State string `json:"state,omitempty"`
	// Account holder's primary phone number
	PrimaryPhone string `json:"primary_phone,omitempty"`
	// Account holder's other phone number
	OtherPhone string `json:"other_phone,omitempty"`
}
