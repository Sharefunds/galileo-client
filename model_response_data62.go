/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData62 struct {
	// A dictionary with keys being an account number (prn) and values that are balance data objects
	Balance map[string][]interface{} `json:"balance"`
}
