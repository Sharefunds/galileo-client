# ResponseData7

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PmtRefNo** | **string** | A Galileo generated account number | [default to null]
**ProductId** | **string** | Galileo generated integer ID for the product | [default to null]
**GalileoAccountNumber** | **string** | Galileo generated integer account number, also known as balance ID | [default to null]
**CardId** | **string** | Unique identifier for a card. None is returned if no card was created. | [optional] [default to null]
**CardNumber** | **string** | A PAN or 16 digit card number (usually masked) | [optional] [default to null]
**ExpiryDate** | **string** | Expiration date of the card | [optional] [default to null]
**CardSecurityCode** | **string** | Card security code (CVV) | [optional] [default to null]
**EmbossLine2** | **string** | Second embossed line on a card | [optional] [default to null]
**BillingCycleDay** | **int32** | The day of the month the billing cycle starts for the account | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

