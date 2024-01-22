/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CustomerProfileShipToAddress struct {
	// The ship to address on the account
	Address1 string `json:"address_1"`
	// Additional ship to address information
	Address2 string `json:"address_2"`
	// City for the ship to address
	City string `json:"city"`
	// State for ship to address
	State string `json:"state"`
	// Postal code for ship to address
	PostalCode string `json:"postal_code"`
	// ISO 3166 country code for ship to address
	CountryCode string `json:"country_code"`
}
