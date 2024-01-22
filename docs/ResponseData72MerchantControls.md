# ResponseData72MerchantControls

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | [**time.Time**](time.Time.md) | Ending date-time of the account-level control. Not populated for product-level controls. | [optional] [default to null]
**DenyAllowFlag** | **string** | Whether the control is DENY (&#x60;d&#x60;) or ALLOW (&#x60;a&#x60;). | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Starting date-time of the account-level control. Not populated for product-level controls. | [optional] [default to null]
**MerchantId** | **string** | Merchant ID. If the value passed has fewer than 15 characters, the Galileo system adds spaces at the end until it has 15 characters. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

