/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData97 struct {
	// The transaction_id of the original request
	RequestTransactionId string `json:"request_transaction_id"`
	// List of API calls
	ApiCalls []ResponseData97ApiCalls `json:"api_calls"`
}
