# Biller

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The account number of the biller entity | [default to null]
**Address1** | **string** | Street and residence number on the account | [optional] [default to null]
**Address2** | **string** | Additional address information on the account | [optional] [default to null]
**BillerId** | **string** | An ID assigned to the biller | [default to null]
**City** | **string** | City for address information | [optional] [default to null]
**Name** | **string** | A name used to identify a biller entity | [default to null]
**Nickname** | **string** | A substring of the biller &#x60;name&#x60;. | [optional] [default to null]
**Phone** | **string** | The main phone number on the biller account | [optional] [default to null]
**PostalCode** | **string** | A postal code for the address information | [optional] [default to null]
**StateProvince** | **string** | State for address information | [optional] [default to null]
**Type_** | **string** | The type of biller: &#x60;P&#x60; (paper) or &#x60;E&#x60; (electronic, RPPS) | [default to null]
**FrequencyType** | **string** | Frequency of the bill payment | [optional] [default to null]
**NextDate** | **string** | The next date that the payment is scheduled | [optional] [default to null]
**EndDate** | **string** | The last date that the payment is scheduled. Can be up to five years in the future | [optional] [default to null]
**Amount** | **float32** | Currency amount as a whole or decimal amount | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

