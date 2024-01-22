/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type NewAccount struct {
	// Galileo-generated account number. Also called the PRN.
	PmtRefNo string `json:"pmt_ref_no"`
	// Identifier for the product associated with the account
	ProductId string `json:"product_id"`
	// Galileo-generated balance ID
	GalileoAccountNumber string `json:"galileo_account_number"`
	// Galileo-generated card identifier (CAD). Use this ID instead of the PAN if not PCI-compliant. `None` means that no card was created
	CardId string `json:"card_id"`
	// List of cardholder identification process (CIP) objects
	Cip []NewAccountCip `json:"cip"`
}
