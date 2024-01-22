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

type Conversation struct {
	// ID of the retrieved cardholder message
	MsgId string `json:"msg_id"`
	// The subject of a secure message
	MsgSubject string `json:"msg_subject"`
	// The date a secure message was received
	MsgDt time.Time `json:"msg_dt"`
	// The read status of a secure message
	Status string `json:"status"`
	// The date of a response to the secure message
	ResponseDt time.Time `json:"response_dt"`
	// Language code in use locally
	Locale string `json:"locale"`
	// The content of a secure message
	MsgText string `json:"msg_text"`
	// List of replies to a secure message
	Replies []ConversationReplies `json:"replies"`
}