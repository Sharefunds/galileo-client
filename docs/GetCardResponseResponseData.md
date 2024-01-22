# GetCardResponseResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardNumber** | **string** | A PAN or 16 digit card number | [default to null]
**ExpiryDate** | **string** | Expiration date of the card | [default to null]
**CardSecurityCode** | **string** | The card verification value (cvv) | [default to null]
**Status** | **string** | The status of the card | [default to null]
**CardId** | **string** | Integer identifier of the card as found in the raw data file (RDF). Unique identifier for a PAN. | [default to null]
**ExternalCardId** | **string** | External card identity | [default to null]
**PmtRefNo** | **string** | A Galileo generated account number | [default to null]
**FirstName** | **string** | A person&#x27;s first name as listed on the account | [default to null]
**MiddleName** | **string** | A person&#x27;s middle name as listed on the account | [default to null]
**LastName** | **string** | A person&#x27;s last name as listed on the account | [default to null]
**EncryptedCardNumber** | **string** | Encrypted number for the card | [default to null]
**EncryptedExpiryDate** | **string** | Encrypted expiration date for the card | [default to null]
**EmbossUuid** | **string** | UUID for an emboss record that has been sent to the embosser. When this field is present outside the &#x60;embossed_cards&#x60; object, it contains the most recent emboss UUID. | [optional] [default to null]
**EmbossedCards** | [**[]ResponseData17EmbossedCards**](ResponseData17_embossed_cards.md) | List of embossed card objects | [default to null]
**FreezeInfo** | [***CardFreezeInfo**](Card_freeze_info.md) |  | [default to null]
**PinFailCount** | **string** | A flag for the PIN fail count (0 or -1) | [default to null]
**PinFailDate** | [**time.Time**](time.Time.md) | Date a PIN fail was recorded | [default to null]
**SpendControls** | [***ResponseData17SpendControls**](ResponseData17_spend_controls.md) |  | [default to null]
**Ssn** | **string** | The social security number associated with the account. This field is returned only if &#x60;CINFN&#x60; is set. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

