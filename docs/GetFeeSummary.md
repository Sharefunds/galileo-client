# GetFeeSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**YtdStart** | **string** | Start date for the year to date. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2022-01-01\&quot;&#x60; | [default to Sat Jan 01 00:00:00 GMT 2022]
**YtdEnd** | **string** | End date of the year to date. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2022-11-05\&quot;&#x60; | [default to Sat Nov 05 00:00:00 GMT 2022]
**MtdStart** | **string** | Start date of the month to date. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2022-03-01\&quot;&#x60; | [default to Tue Mar 01 00:00:00 GMT 2022]
**MtdEnd** | **string** | End date of the month to date. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2022-03-15\&quot;&#x60; | [default to Tue Mar 15 00:00:00 GMT 2022]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

