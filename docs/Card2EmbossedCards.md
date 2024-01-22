# Card2EmbossedCards

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmbossUuid** | **string** | UUID for an emboss record that has been sent to the embosser. When this field is present outside the &#x60;embossed_cards&#x60; object, it contains the most recent emboss UUID. | [optional] [default to null]
**CreatedDate** | **string** | Date the card was created | [default to null]
**EmbossDate** | **string** | Date the card was embossed | [default to null]
**ExpiryDate** | [**time.Time**](time.Time.md) | Expiration date of the card | [optional] [default to null]
**Status** | **string** | The status of the card | [default to null]
**ShippingType** | **string** | The type of shipping used for the card | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

