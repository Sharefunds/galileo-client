/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ProfileShipToAddress struct {
	// First line of the shipping address
	Address1 string `json:"address_1"`
	// Second line of the shipping address
	Address2 string `json:"address_2"`
	// City for the shipping address
	City string `json:"city"`
	// State for the shipping address
	State string `json:"state"`
	// Postal code for the shipping address
	PostalCode string `json:"postal_code"`
	// ISO 3166 country code indicating the country for the shipping address
	CountryCode string `json:"country_code"`
}
