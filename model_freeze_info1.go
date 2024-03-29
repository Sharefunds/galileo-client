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

type FreezeInfo1 struct {
	// Whether the card is Frozen or Unfrozen.
	Status string `json:"status"`
	// Start datetime for the card freeze.
	StartDate time.Time `json:"start_date"`
	// End datetime for the card freeze.
	EndDate time.Time `json:"end_date"`
}
