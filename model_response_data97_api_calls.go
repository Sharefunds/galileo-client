/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type ResponseData97ApiCalls struct {
	// An ID used for internal tracking of the transaction
	GalTransId string `json:"gal_trans_id"`
	// An ID used for internal tracking of the transaction
	ResultTransId string `json:"result_trans_id"`
	// The status of the original request
	StatusCode string `json:"status_code"`
	// The datetime the original request
	InTimestamp time.Time `json:"in_timestamp"`
	// An ID used for internal tracking of the transaction
	Rtoken string `json:"rtoken"`
}
