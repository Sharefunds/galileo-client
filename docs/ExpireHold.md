# ExpireHold

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**HoldId** | **int32** | Hold ID (&#x60;hold_id&#x60;) as returned by the &lt;a href&#x3D;\&quot;ref:post_createhold\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Create Hold&lt;/a&gt; or &lt;a href&#x3D;\&quot;ref:post_getholdhistory\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Get Hold History&lt;/a&gt; endpoint. Pattern: Positive integer Example: &#x60;453434&#x60; | [optional] [default to null]
**ExpirationDateTime** | [**time.Time**](time.Time.md) | Date and time to expire the hold. Must be a date-time in the future. Leave this parameter empty to expire the hold immediately. Pattern: YYYY-MM-DD hh:mm:ss Example: &#x60;\&quot;2017-01-01 13:00:00\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

