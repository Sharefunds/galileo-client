# ModifyLocationFeeBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**Location** | **string** | Identifier for the location that was created with Create Location. This value is returned by Get Locations in the &#x60;location_id&#x60; field when &#x60;locationType: 0&#x60; or &#x60;locationType: 2&#x60;, or in the &#x60;provider_specified_id&#x60; field when &#x60;locationType: 1&#x60;. Pattern: Must be a number if type is 0; Must be less than or equal to 15 characters if type is 1. Example: &#x60;\&quot;a455-3483\&quot;&#x60; | [default to a455-3483]
**LocationType** | **int32** | Location type, as specified by Create Location. 0&#x3D;Galileo location ID, 1&#x3D;Partner location ID, 2&#x3D;Don&#x27;t validate. Pattern: &#x60;0&#x60;, &#x60;1&#x60; or &#x60;2&#x60; Example: &#x60;\&quot;0\&quot;&#x60; | [default to null]
**FeeType** | **string** | Type of fee as specified by Create Location Fee. This value must be the original value provided at the time of the fee creation: &#x60;FBD&#x60; (FBE direct deposit fee) or &#x60;DEL&#x60; (ACH direct deposit fee) Pattern: &#x60;FBD&#x60; or &#x60;DEL&#x60; Example: &#x60;\&quot;FBD\&quot;&#x60; | [default to FEE_TYPE.FBD]
**Amount** | **float32** | Specify either a flat amount or a percentage. For a flat amount the valid range is between &#x60;.01&#x60; and &#x60;999&#x60;. For a percentage the valid range is &#x60;.01&#x60; (1%) to &#x60;1&#x60; (100%). When this parameter is populated then &#x60;amountType&#x60; is required. Pattern: 1, 1.10, 0.5 Example: 1.23 | [optional] [default to null]
**AmountType** | **string** | Specify whether the value in &#x60;amount&#x60; is a flat rate (&#x60;flat&#x60;) or a percentage (&#x60;percent&#x60;). When this parameter is populated then &#x60;amount&#x60; is required. Pattern: &#x60;flat&#x60; or &#x60;percent&#x60; Example: &#x27;flat&#x27; | [optional] [default to null]
**MinFee** | **float32** | The minimum amount for the fee, if the amount type is a percentage. Pattern: 10, 20.5, 2.33 Example: 1.10 | [optional] [default to null]
**MaxFee** | **float32** | The maximum amount for the fee, if the amount type is a percentage. Pattern: 10, 20.5, 2.33 Example: 1.10 | [optional] [default to null]
**Active** | **string** | Specify whether to activate the fee (&#x60;Y&#x60;) or deactivate the fee (&#x60;N&#x60;). Pattern: &#x60;Y&#x60; or &#x60;N&#x60; Example: &#x60;\&quot;Y\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

