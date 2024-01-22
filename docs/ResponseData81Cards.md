# ResponseData81Cards

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | **string** | Galileo-generated card identifier (CAD). Use this ID instead of the PAN if not PCI-compliant. &#x60;None&#x60; means that no card was created. | [default to null]
**PmtRefNo** | **string** | Galileo-generated account number. Also called the PRN. | [default to null]
**Status** | **string** | The status of the card | [default to null]
**CardNumber** | **string** | The card number, also called the PAN (usually masked) | [default to null]
**ExpiryDate** | [**time.Time**](time.Time.md) | Date when the card expires | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

