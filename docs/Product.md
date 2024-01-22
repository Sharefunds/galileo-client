# Product

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apy** | **string** | The Annual percentage yield | [optional] [default to null]
**Intmx** | **float32** | The maximum balance to which interest will be applied | [optional] [default to null]
**InterestTiers** | [**[]ProductInterestTiers**](Product_interest_tiers.md) | Interest tiers on the product, if any | [optional] [default to null]
**ProdId** | **string** | ID for the product | [default to null]
**Description** | **string** | Description of the product | [default to null]
**BaseBin** | **string** | The active base bin for the product | [default to null]
**BaseBins** | [**[]ProductBaseBins**](Product_base_bins.md) | A list of base bins for the product. Also shows, for each bin, whether it is active or not. | [default to null]
**HoldDays** | **int32** | Hold days for the product | [default to null]
**MaxPinRetries** | **int32** | Maximum number of pin retries | [default to null]
**ProgId** | **string** | The program ID for the product | [default to null]
**ProgName** | **string** | The name of the program | [default to null]
**Type_** | **string** | Type of product | [default to null]
**MaxBalance** | **float32** | The max balance allowed | [default to null]
**MccRestrictions** | [**[]ProductMccRestrictions**](Product_mcc_restrictions.md) |  | [default to null]
**AuthLimits** | [**[]ProductAuthLimits**](Product_auth_limits.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

