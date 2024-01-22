# ResponseData80

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prn** | **string** | The payment reference number | [default to null]
**ProdId** | **int32** | Unique system generated product id associated with the account | [default to null]
**AppDate** | **string** | The date the account&#x27;s application was submitted. &#x60;Formatted yyyy-mm-dd&#x60;. | [default to null]
**Status** | **string** | Indicates the status of an account | [default to null]
**ActiveFlag** | **string** | Indicates if an account is active or not | [default to null]
**BillCycleDay** | **string** | Day of the month that the customer is billed | [default to null]
**GroupId** | **string** | Integer value used for tracking the location (store, entity, etc...) where the customer was acquired | [default to null]
**StartDate** | **string** | Date the account was activated | [default to null]
**GalileoAccountNumber** | **string** | Galileo generated integer account number, also known as balance ID | [default to null]
**NewEmbossUuid** | **string** | UUID for an emboss record that has not yet been sent to the embosser. | [optional] [default to null]
**Cards** | [**[]ResponseData80Cards**](ResponseData80_cards.md) | List of cards associated with the account | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

