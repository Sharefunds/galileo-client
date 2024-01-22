# ResponseData37NewAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PmtRefNo** | **string** | Galileo-generated account number. Also called the PRN. | [default to null]
**ProductId** | **string** | Identifier for the product associated with the account | [default to null]
**GalileoAccountNumber** | **string** | Galileo-generated balance ID | [default to null]
**CardId** | **string** | Galileo-generated card identifier (CAD). Use this ID instead of the PAN if not PCI-compliant. &#x60;None&#x60; means that no card was created | [default to null]
**Cip** | [**[]NewAccountCip**](NewAccount_cip.md) | List of cardholder identification process (CIP) objects | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

