# GetGroupsInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**GroupIds** | **[]string** | A list of group identifiers to get information for. To input a JSON list as the value, you must set the &#x60;Content-Type&#x60; in the header to &#x60;application/json&#x60;. If the &#x60;Content-Type&#x60; is not &#x60;json&#x60;, then pass one parameter/value pair for each value. See &lt;a href&#x3D;\&quot;doc:creating-a-corporate-hierarchy#inputting-multiple-values\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Inputting multiple values&lt;/a&gt; for examples. Pattern: List of identifiers Example: &#x60;[\&quot;2222\&quot;, \&quot;3333\&quot;, \&quot;4444\&quot;]&#x60; or &#x60;[\&quot;5555\&quot;]&#x60; | [default to ["0","4"]]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

