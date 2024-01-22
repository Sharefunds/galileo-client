# Fee

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeEventId** | **string** | Galileo generated fee transaction integer ID | [default to null]
**Type_** | **string** | Three-letter fee code. This is not the transaction type (otype). | [default to null]
**TypeDescription** | **string** | A description of the type code | [default to null]
**Amt** | **string** | Amount of the fee charge | [default to null]
**FeeDate** | [**time.Time**](time.Time.md) | A timestamp for the time the fee was charged | [default to null]
**CardId** | **int32** | Integer identifier of the card as found in the raw data file (RDF). Unique identifier for a PAN. | [default to null]
**FeeDescription** | **string** | The description on a fee | [default to null]
**RelatedTransaction** | [***FeeRelatedTransaction**](Fee_related_transaction.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

