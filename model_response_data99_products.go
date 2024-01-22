/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData99Products struct {
	// The Annual percentage yield
	Apy string `json:"apy,omitempty"`
	// The maximum balance to which interest will be applied
	Intmx float32 `json:"intmx,omitempty"`
	// Interest tiers on the product, if any
	InterestTiers []ProductInterestTiers `json:"interest_tiers,omitempty"`
	// ID for the product
	ProdId string `json:"prod_id"`
	// Description of the product
	Description string `json:"description"`
	// The active base bin for the product
	BaseBin string `json:"base_bin"`
	// A list of base bins for the product. Also shows, for each bin, whether it is active or not.
	BaseBins []ProductBaseBins `json:"base_bins"`
	// Hold days for the product
	HoldDays int32 `json:"hold_days"`
	// Maximum number of pin retries
	MaxPinRetries int32 `json:"max_pin_retries"`
	// The program ID for the product
	ProgId string `json:"prog_id"`
	// The name of the program
	ProgName string `json:"prog_name"`
	// Type of product
	Type_ string `json:"type"`
	// The max balance allowed
	MaxBalance float32 `json:"max_balance"`
	MccRestrictions []ProductMccRestrictions `json:"mcc_restrictions"`
	AuthLimits []ProductAuthLimits `json:"auth_limits"`
}
