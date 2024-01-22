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

type Hold struct {
	// A unique ID for the hold to be expired
	HoldId string `json:"hold_id"`
	// The date the hold was created
	CreateDt time.Time `json:"create_dt"`
	// The date the hold will expire
	ExpiryDt time.Time `json:"expiry_dt"`
	// A code unique to the source of the activity (such as fees, adjustments, etc.)
	SourceId string `json:"source_id"`
	// The date the hold was last modified or created
	ChangeDt time.Time `json:"change_dt"`
	// The type of hold
	HoldType string `json:"hold_type"`
	// An external identifier associate with the hold
	ExtId string `json:"ext_id"`
	// The description associated with the hold
	Dscr string `json:"dscr"`
	// The process that created the hold
	OriginatingSystemId string `json:"originating_system_id"`
	// The id for the agent that created the hold
	AgentId string `json:"agent_id"`
	// The financial sum being held
	Amount float64 `json:"amount"`
	// The transaction ID associated with the hold
	Xid string `json:"xid"`
	// The process that expired the hold
	ExpiringSystemId string `json:"expiring_system_id"`
	// The id for the agent that expired the hold
	ExpiringAgentId string `json:"expiring_agent_id"`
}
