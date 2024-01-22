# ResponseData122ActiveHolds

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoldId** | **string** | A unique ID for the hold to be expired | [default to null]
**CreateDt** | [**time.Time**](time.Time.md) | The date the hold was created | [default to null]
**ExpiryDt** | [**time.Time**](time.Time.md) | The date the hold will expire | [default to null]
**SourceId** | **string** | A code unique to the source of the activity (such as fees, adjustments, etc.) | [default to null]
**ChangeDt** | [**time.Time**](time.Time.md) | The date the hold was last modified or created | [default to null]
**HoldType** | **string** | The type of hold | [default to null]
**ExtId** | **string** | An external identifier associate with the hold | [default to null]
**Dscr** | **string** | The description associated with the hold | [default to null]
**OriginatingSystemId** | **string** | The process that created the hold | [default to null]
**AgentId** | **string** | The id for the agent that created the hold | [default to null]
**Amount** | **float64** | The financial sum being held | [default to null]
**Xid** | **string** | The transaction ID associated with the hold | [default to null]
**ExpiringSystemId** | **string** | The process that expired the hold | [default to null]
**ExpiringAgentId** | **string** | The id for the agent that expired the hold | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

