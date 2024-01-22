# CreateAccountResponseResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PmtRefNo** | **string** | Galileo-generated account number. Also called the PRN. | [default to null]
**ProductId** | **string** | Identifier for the product associated with the account | [default to null]
**GalileoAccountNumber** | **string** | Galileo-generated balance ID | [default to null]
**Cip** | [**[]ResponseData4Cip**](ResponseData4_cip.md) | List of cardholder identification process (CIP) objects | [default to null]
**CardId** | **string** | Galileo-generated card identifier (CAD). Use this ID instead of the PAN if not PCI-compliant. &#x60;None&#x60; means that no card was created | [optional] [default to null]
**NewEmbossUuid** | **string** | UUID for an emboss record that has not yet been sent to the embosser. | [optional] [default to null]
**CardNumber** | **string** | The primary account number (PAN) printed on the card | [default to null]
**ExpiryDate** | **string** | Date when the card expires | [default to null]
**CardSecurityCode** | **string** | The card verification value (CVV) | [default to null]
**EmbossLine2** | **string** | Second line to emboss on the card | [default to null]
**BillingCycleDay** | **int32** | The day of the month the billing cycle starts for the account | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

