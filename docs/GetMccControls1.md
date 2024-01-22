# GetMccControls1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DenyAllowFlag** | **string** | Whether the control is DENY (&#x60;d&#x60;) or ALLOW (&#x60;a&#x60;). | [default to null]
**EndDate** | **string** | Ending date-time of the account-level control. Not populated for product-level controls. | [optional] [default to null]
**EndMcc** | **int32** | The last MCC in the range. If the control contains only one MCC, this field is either blank or the same value as &#x60;begMcc&#x60;. | [default to null]
**StartDate** | **string** | Starting date-time of the account-level control. Not populated for product-level controls. | [optional] [default to null]
**OnlineOnly** | **string** | Whether the control applies to online (card not present) transactions only. **Y**—Online only; **N**—All transactions | [default to null]
**BeginningMcc** | **int32** | The first MCC in the range. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

