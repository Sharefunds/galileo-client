/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RegisteredCardData struct {
	// Masked card number of card registered on Ingo System
	CardNumber string `json:"card_number"`
	// card_id of card registered on Ingo System
	CardId int32 `json:"card_id"`
	// status of card registered on Ingo System
	Status string `json:"status"`
}