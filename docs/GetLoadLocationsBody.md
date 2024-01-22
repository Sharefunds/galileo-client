# GetLoadLocationsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**ProgramId** | **int32** | Galileo-generated identifier for the program. Pattern: Positive integer Example: &#x60;1032&#x60; | [default to 1032]
**PostalCode** | **string** | Postal code for the locations to retrieve. The finder returns load locations in and near this postal code. Pattern: &#x60;12345&#x60;, &#x60;12345-6789&#x60; or &#x60;K1A-1A1&#x60; Example: &#x60;\&quot;84121\&quot;&#x60; | [optional] [default to null]
**ResultCount** | **int32** | The maximum number of records to return. Pattern:  Integer 1&amp;ndash;9999 Example: &#x60;100&#x60; | [optional] [default to 5]
**LatLong** | **string** | The latitude and longitude of the requester&#x27;s location. The finder returns load locations near this geographical point. Pattern: Latitude/longitude pair separated by a comma and space; no limit on decimal places Example: &#x60;40.57297941556069, -111.89920138637734&#x60; | [optional] [default to null]
**NewVersion** | **int32** | Specifies whether to use the new version of the load-location finder: * &#x60;0&#x60; &amp;mdash; Do not use the new version * &#x60;1&#x60; &amp;mdash; Use the new version  Pattern: &#x60;0&#x60; or &#x60;1&#x60; Example: &#x60;1&#x60;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

