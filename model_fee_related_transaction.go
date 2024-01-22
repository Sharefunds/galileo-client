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

// A data structure that contains information on transactions related to a fee
type FeeRelatedTransaction struct {
	// Information on a transaction or authorization
	Details string `json:"details"`
	// Amount of a fee or transaction charge
	Amt float32 `json:"amt"`
	// The time stamp of a posted transaction
	PostTs time.Time `json:"post_ts"`
}