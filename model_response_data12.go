/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseData12 struct {
	// Positive integer value used to identify the biller
	BillerId string `json:"biller_id"`
	// The RPPS biller name
	Name string `json:"name"`
}
