# Location

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **int32** | The numerical value assigned to the location | [default to null]
**Name** | **string** | The name of the location | [default to null]
**Address1** | **string** | Address 1 of the location | [default to null]
**Address2** | **string** | Address 2 of the location | [default to null]
**City** | **string** | City of the location | [default to null]
**State** | **string** | State of the location | [default to null]
**PostalCode** | **string** | Postal code of the location | [default to null]
**ParentLocationId** | **string** | The numerical ID of the parent location | [default to null]
**ProviderSpecifiedId** | **string** | The partner or program identity for a location | [default to null]
**CentralBillFlag** | **int32** | Indicates a central corporate credit aggregate billing location. Allows one centralized billing location for all invoice records attached to the location. | [optional] [default to null]
**CreditLimit** | **float32** | The credit limit for this location | [optional] [default to null]
**StoreType** | **string** | The store or location type. Values are &#x60;Corporate&#x60;, &#x60;Chain&#x60;, &#x60;Reseller&#x60;, &#x60;Region&#x60;, or &#x60;Store&#x60;. | [default to null]
**Status** | **string** | Current status of the location. Values are &#x60;Closed&#x60;, &#x60;Active&#x60;, &#x60;Suspended&#x60;, or &#x60;New&#x60;. | [default to null]
**Fees** | [**[]LocationFees**](Location_fees.md) | List of fees attached to the location | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

