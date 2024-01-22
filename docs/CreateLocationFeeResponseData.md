# CreateLocationFeeResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxFee** | **string** | The minimum amount to charge, when &#x60;amount_type&#x60; is &#x60;percent&#x60;. | [default to null]
**Location** | **int32** | Galileo ID of the location, when &#x60;locationType&#x60; is &#x60;0&#x60; or &#x60;2&#x60;. | [default to null]
**ProviderSpecifiedId** | **string** | Provider ID of the location, when &#x60;locationType&#x60; is &#x60;1&#x60;. | [default to null]
**MinFee** | **string** | The minimum amount to charge, when &#x60;amount_type&#x60; is &#x60;percent&#x60;. | [default to null]
**FeeType** | **string** | Type of fee, &#x60;FBD&#x60; (FBE direct deposit fee) or &#x60;DEL&#x60; (ACH direct deposit fee). | [default to null]
**Amount** | **float32** | Amount of the fee. | [default to null]
**AmountType** | **string** | The type of the amount: &#x60;flat&#x60; or &#x60;percent&#x60;. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

