# Card1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | **string** | Unique identifier for a card | [default to null]
**EmbossUuid** | **string** | UUID for an emboss record that has been sent to the embosser. When this field is present outside the &#x60;embossed_cards&#x60; object, it contains the most recent emboss UUID. | [optional] [default to null]
**CardNumber** | **string** | A PAN or 16 digit card number (usually masked) | [default to null]
**CardStatus** | **string** | A single character code specifying the status of the card | [optional] [default to null]
**CardSecurityCode** | **string** | The card verification value (cvv) | [optional] [default to null]
**ExpiryDate** | **string** | Expiration date of the card | [optional] [default to null]
**ExternalCardId** | **string** | Partner specified external card identity | [default to null]
**PinFailCount** | **string** | Counts the number of times a PIN entry fails | [default to null]
**PinFailDate** | [**time.Time**](time.Time.md) | Date a PIN fail was recorded | [default to null]
**FreezeInfo** | [***Card1FreezeInfo**](Card1_freeze_info.md) |  | [default to null]
**EncryptedCardNumber** | **string** | Encrypted number for the card | [optional] [default to null]
**EncryptedExpiryDate** | **string** | Encrypted expiration date for the card | [optional] [default to null]
**EmbossedCards** | [**[]Card1EmbossedCards**](Card1_embossed_cards.md) | List of embossed card objects | [default to null]
**SpendControls** | [***ResponseData17SpendControls**](ResponseData17_spend_controls.md) |  | [optional] [default to null]
**FirstName** | **string** | First name | [default to null]
**MiddleName** | **string** | Middle name | [default to null]
**LastName** | **string** | Last name | [default to null]
**PmtRefNo** | **string** | A Galileo generated account number | [default to null]
**Status** | **string** | Flag showing whether a card is frozen or not | [optional] [default to null]
**Ssn** | **string** | The social security number associated with the account. This field is returned only if &#x60;CINFN&#x60; is set. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

