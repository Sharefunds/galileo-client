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
type GetOverdraftBalanceResponseResponseData struct {
	// PRN of the overdraft account
	OverdraftAccountNo string `json:"OverdraftAccountNo"`
	// Available overdraft limit
	Available float32 `json:"Available"`
	// Payback amount
	Payback float32 `json:"Payback"`
	// The posted open to buy, including pending fees
	PaybackWithPendingFees float32 `json:"PaybackWithPendingFees"`
	// The actual open to buy of the overdraft account
	RawBalance float32 `json:"RawBalance"`
	// The actual balance of the primary account if overdraft was combined with it
	ActualWithPrimary float32 `json:"ActualWithPrimary"`
	// The available limit on the overdraft account (the most the cardholder could use on their overdraft if nothing was owed)
	Limit float32 `json:"Limit"`
	// The limit that was set on the overdraft account (the actual limit available to the cardholder may be less if an initial fee was charged)
	RawLimit float32 `json:"RawLimit"`
	// A flag indicating whether the overdraft account has been frozen or not (the value '1' indicating that it is frozen)
	Frozen string `json:"Frozen"`
}