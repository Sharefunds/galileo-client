/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData11 struct {
	// The ID assigned to the biller
	BillerId string `json:"biller_id"`
	// First address line of the biller
	Address1 string `json:"address_1,omitempty"`
	// Second address line of the biller
	Address2 string `json:"address_2,omitempty"`
	// City of the biller
	City string `json:"city,omitempty"`
	// State of the biller
	StateProvince string `json:"state_province,omitempty"`
	// A postal code for the biller
	PostalCode string `json:"postal_code,omitempty"`
	// The main phone number on the biller account
	Phone string `json:"phone,omitempty"`
	// Frequency of the bill payment: `O` (one time) `W` (weekly) `M` (monthly) `Q` (quarterly) `Y` (yearly)
	FrequencyType string `json:"frequency_type,omitempty"`
	// The next date that the payment is scheduled
	NextDate string `json:"next_date,omitempty"`
	// The last date that the payment is scheduled. Can be up to five years in the future
	EndDate string `json:"end_date,omitempty"`
	// Amount of the bill payment
	Amount float32 `json:"amount,omitempty"`
}