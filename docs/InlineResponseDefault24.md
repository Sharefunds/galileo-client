# InlineResponseDefault24

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** | The response status code. May return a string for some statuses. | [default to null]
**Status** | **string** | The condition of a process or response | [default to null]
**ProcessingTime** | **float32** | The time elapsed in processing the transaction | [default to null]
**Echo** | [***GetTransHistoryResponseEcho**](GetTransHistoryResponse_echo.md) |  | [default to null]
**SystemTimestamp** | [**time.Time**](time.Time.md) | A system generated timestamp | [default to null]
**Rtoken** | **string** | A Galileo-generated ID used for tracking | [default to null]
**Errors** | **[]string** | A list of errors generated while the request was processed | [optional] [default to null]
**ResponseData** | [***CreateSimulatedCardSettleResponseResponseData**](CreateSimulatedCardSettleResponse_response_data.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

