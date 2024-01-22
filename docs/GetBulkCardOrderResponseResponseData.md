# GetBulkCardOrderResponseResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the card order: &#x60;P&#x60; - Processed; &#x60;E&#x60; - Error | [default to null]
**ProdId** | **string** | Identifier for the product associated with the card order | [default to null]
**OrderId** | **string** | Identifier for the card order | [default to null]
**NumberOfCards** | **int32** | Number of cards in the order | [default to null]
**EmbossWith** | **string** | First line to be embossed on cards | [default to null]
**ShipToName** | **string** | Shipping name the for card order | [default to null]
**ShipToAddress** | **string** | Address for shipping card order | [default to null]
**ShipToCity** | **string** | City for shipping address | [default to null]
**ShipToStateOrProvince** | **string** | State or province for the shipping address | [default to null]
**ShipToPostalCode** | **string** | Postal code for the shipping address | [default to null]
**Location** | **string** | Location for cards | [default to null]
**LocationName** | **string** | Location name for cards | [default to null]
**LocationType** | **int32** | Location type for cards | [default to null]
**CardIdRangeStart** | **int32** | This field returns zeros only | [default to null]
**CardIdRangeEnd** | **int32** | This field returns zeros only | [default to null]
**Cards** | [**[]ResponseData81Cards**](ResponseData81_cards.md) | List of transfer accounts | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

