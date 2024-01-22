/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData71MccControls struct {
	// Whether the control is DENY (`d`) or ALLOW (`a`).
	DenyAllowFlag string `json:"deny_allow_flag"`
	// Ending date-time of the account-level control. Not populated for product-level controls.
	EndDate string `json:"end_date,omitempty"`
	// The last MCC in the range. If the control contains only one MCC, this field is either blank or the same value as `begMcc`.
	EndMcc int32 `json:"end_mcc"`
	// Starting date-time of the account-level control. Not populated for product-level controls.
	StartDate string `json:"start_date,omitempty"`
	// Whether the control applies to online (card not present) transactions only. **Y**—Online only; **N**—All transactions
	OnlineOnly string `json:"online_only"`
	// The first MCC in the range.
	BeginningMcc int32 `json:"beginning_mcc"`
}
