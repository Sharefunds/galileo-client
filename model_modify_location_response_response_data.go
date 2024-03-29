/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A structure for the response data. It can be empty but usually will contain information.
type ModifyLocationResponseResponseData struct {
	// Name of the location
	Name string `json:"name"`
	// Address 1 of the location
	Address1 string `json:"address1"`
	// Address 2 of the location
	Address2 string `json:"address2"`
	// City of the location
	City string `json:"city"`
	// State of the location
	State string `json:"state"`
	// Postal code of the location
	PostalCode string `json:"postalCode"`
	// The phone number associated with the locations
	Phone string `json:"phone"`
	// The partner or program identity for a location
	ProviderSpecifiedId string `json:"providerSpecifiedId"`
	// Indicates a central corporate credit aggregate billing location
	CentralBillFlag int32 `json:"centralBillFlag,omitempty"`
	// The credit limit for this location
	CreditLimit float32 `json:"creditLimit,omitempty"`
	// The type of time period
	CycleType string `json:"cycleType,omitempty"`
	// If cycleType is weekly, the value must be a day of the week with 0 meaning Sunday. If cycleType is monthly, the value must be a day of the month from 1 to 28. If daily, the value is to be 1.
	CycleInvoice int32 `json:"cycleInvoice,omitempty"`
	// Represents a day of the week or day of the month that the invoice cycle begins (1=Monday, 7=Sunday).
	CycleStart int32 `json:"cycleStart,omitempty"`
}
