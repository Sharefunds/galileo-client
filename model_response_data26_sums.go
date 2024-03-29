/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// sums
type ResponseData26Sums struct {
	// The total sum of transactions that have not been settled yet
	Unsettled float32 `json:"unsettled"`
	// The amount of a singled transaction that has been settled
	Settled float32 `json:"settled"`
	// The sum amount of adjustments
	Adjustment float32 `json:"adjustment"`
	// The sum amount charged for a service
	Fee float32 `json:"fee"`
	// The sum amount of payments
	Payment float32 `json:"payment"`
}
