# ResponseData3Holds

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoldId** | **string** | Identifier for the hold | [default to null]
**CreateDt** | [**time.Time**](time.Time.md) | Timestamp when the hold was created | [default to null]
**ExpiryDt** | [**time.Time**](time.Time.md) | Date the hold expires | [default to null]
**SourceId** | **string** | Source ID for the hold | [default to null]
**ChangeTs** | [**time.Time**](time.Time.md) | Timestamp when the hold was changed | [optional] [default to null]
**HoldType** | **string** | Identifier for the hold type | [default to null]
**ExtId** | **string** | External identifier for the hold | [default to null]
**Dscr** | **string** | Description of the hold | [default to null]
**OriginatingSystemId** | **string** | Identifier that specifies the source system for the hold | [default to null]
**AgentId** | **string** | Identifier for the agent that created the hold | [default to null]
**Amount** | **float32** | The amount for the hold | [default to null]
**Xid** | **string** | The transaction ID associated with the hold | [default to null]
**ExpiringSystemId** | **string** | Identifier for the process that expired the hold | [default to null]
**ExpiringAgentId** | **string** | Identifier for the agent that expired the hold | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

