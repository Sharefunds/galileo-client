# AccountsCards

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | **string** | Unique identifier for a card | [default to null]
**CardNumber** | **string** | A PAN or 16 digit card number (usually masked) | [default to null]
**EmbossUuid** | **string** | UUID for an emboss record that has been sent to the embosser. When this field is present outside the &#x60;embossed_cards&#x60; object, it contains the most recent emboss UUID. | [optional] [default to null]
**CardStatus** | **string** | A single character code specifying the status of the card | [default to null]
**ExpiryDate** | [**time.Time**](time.Time.md) | Expiration date of the card | [default to null]
**ExternalCardId** | **string** | Partner specified external card identity | [default to null]
**PinFailCount** | **string** | Counts the number of times a PIN entry fails | [default to null]
**PinFailDate** | [**time.Time**](time.Time.md) | Date a PIN fail was recorded | [default to null]
**FreezeInfo** | [***CardFreezeInfo**](Card_freeze_info.md) |  | [default to null]
**EmbossedCards** | [**[]CardEmbossedCards**](Card_embossed_cards.md) | List of embossed card objects | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

