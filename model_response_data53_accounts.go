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

// Acts as a logical separator between accounts if more than one is returned
type ResponseData53Accounts struct {
	// The payment reference number
	Prn string `json:"prn"`
	// Unique system generated product id associated with the account
	ProdId string `json:"prod_id"`
	// The date the account's application was submitted
	AppDate time.Time `json:"app_date"`
	// Indicates the status of an account
	Status string `json:"status"`
	// Indicates if an account is active or not
	ActiveFlag string `json:"active_flag"`
	// The balance (open to buy) of an account
	Balance float32 `json:"balance,omitempty"`
}
