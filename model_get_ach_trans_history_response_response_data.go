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
type GetAchTransHistoryResponseResponseData struct {
	// The number of transactions found
	Found int32 `json:"found"`
	// List of transactions
	Transactions []ResponseData107Transactions `json:"transactions"`
}
