# GetCustomerNoteHistoryBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**StartDate** | **string** | The beginning date for the date range, either a date or a date-time. Pattern: YYYY-MM-DD or YYYY-MM-DD hh:mm:ss Example: &#x60;\&quot;2016-01-01\&quot;&#x60; | [default to Fri Jan 01 00:00:00 GMT 2016]
**EndDate** | **string** | The end date for the date range, either a date or a date-time. Must be equal to or later than &#x60;startDate&#x60;. Pattern: YYYY-MM-DD or YYYY-MM-DD hh:mm:ss Example: &#x60;\&quot;2016-01-01\&quot;&#x60; | [default to Fri Jan 01 00:00:00 GMT 2016]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

