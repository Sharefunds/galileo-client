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

type ResponseData3Holds struct {
	// Identifier for the hold
	HoldId string `json:"hold_id"`
	// Timestamp when the hold was created
	CreateDt time.Time `json:"create_dt"`
	// Date the hold expires
	ExpiryDt time.Time `json:"expiry_dt"`
	// Source ID for the hold
	SourceId string `json:"source_id"`
	// Timestamp when the hold was changed
	ChangeTs time.Time `json:"change_ts,omitempty"`
	// Identifier for the hold type
	HoldType string `json:"hold_type"`
	// External identifier for the hold
	ExtId string `json:"ext_id"`
	// Description of the hold
	Dscr string `json:"dscr"`
	// Identifier that specifies the source system for the hold
	OriginatingSystemId string `json:"originating_system_id"`
	// Identifier for the agent that created the hold
	AgentId string `json:"agent_id"`
	// The amount for the hold
	Amount float32 `json:"amount"`
	// The transaction ID associated with the hold
	Xid string `json:"xid"`
	// Identifier for the process that expired the hold
	ExpiringSystemId string `json:"expiring_system_id"`
	// Identifier for the agent that expired the hold
	ExpiringAgentId string `json:"expiring_agent_id"`
}
