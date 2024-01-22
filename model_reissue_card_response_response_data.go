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
type ReissueCardResponseResponseData struct {
	// The payment reference number
	Prn string `json:"prn"`
	// Unique system generated product id associated with the account
	ProdId int32 `json:"prod_id"`
	// The date the account's application was submitted. `Formatted yyyy-mm-dd`.
	AppDate string `json:"app_date"`
	// Indicates the status of an account
	Status string `json:"status"`
	// Indicates if an account is active or not
	ActiveFlag string `json:"active_flag"`
	// Day of the month that the customer is billed
	BillCycleDay string `json:"bill_cycle_day"`
	// Integer value used for tracking the location (store, entity, etc...) where the customer was acquired
	GroupId string `json:"group_id"`
	// Date the account was activated
	StartDate string `json:"start_date"`
	// Galileo generated integer account number, also known as balance ID
	GalileoAccountNumber string `json:"galileo_account_number"`
	// UUID for an emboss record that has not yet been sent to the embosser.
	NewEmbossUuid string `json:"new_emboss_uuid,omitempty"`
	// List of cards associated with the account
	Cards []ResponseData80Cards `json:"cards"`
}