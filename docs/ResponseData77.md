# ResponseData77

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the card | [default to null]
**PmtRefNo** | **string** | A Galileo generated account number | [default to null]
**ProdId** | **string** | Galileo generated product ID of the account | [default to null]
**MaxLoadAmount** | **string** | The maximum amount that can be loaded on a card | [default to null]
**CardId** | **string** | Integer identifier of the card as found in the raw dta file (RDF). Unique identifier for a PAN. | [default to null]
**CardNumber** | **string** | A PAN or 16 digit card number | [default to null]
**ExpiryDate** | **string** | Expiration date of the card | [default to null]
**CardSecurityCode** | **string** | The card verification value (cvv), is a card anti-fraud measure | [default to null]
**BatchId** | **string** | The ID of a request for a batch of instant issue cards to be created | [default to null]
**CaseId** | **string** | Physical cards are shipped in bundles, which are divided by boxes. A case_id is a tracking number on each box and bundle. | [default to null]
**BoxId** | **string** | A shipping tracking number on a box of shipped cards | [default to null]
**BundleId** | **string** | A shipping tracking number on the bundle of shipped cards | [default to null]
**GalileoLocationId** | **string** | A code for the location at which account was created, if applicable. Can be provided by client or Galileo | [default to null]
**PartnerLocationId** | **string** | A code for the partner location | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

