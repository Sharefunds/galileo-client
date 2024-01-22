# SetAccountFeatureBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt;, &lt;&lt;glossary:PAN&gt;&gt; or &lt;&lt;glossary:CAD&gt;&gt; of the account. Pattern: PAN, PRN, or CAD Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**FeatureType** | **int32** | Consult the **featureType** column in the &lt;a href&#x3D;\&quot;ref:api-reference-set-account-feature-values\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Set Account Feature Values&lt;/a&gt; table for valid values. Pattern: Integer Example: &#x60;3&#x60; | [default to 3]
**FeatureValue** | **string** | Consult the **featureValue** column in the &lt;a href&#x3D;\&quot;ref:api-reference-set-account-feature-values\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Set Account Feature Values&lt;/a&gt; table for valid values. Pattern: As specified in the **featureValue** column Example: &#x60;Y&#x60; | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The starting date-time for the date range. Default: current date-time Pattern: YYYY-MM-DD hh:mm:ss Example: &#x60;\&quot;2016-01-01 10:10:15\&quot;&#x60; | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The ending date-time for the date range; must be equal to or later than &#x60;startDate&#x60;. Default: 3000-01-01 00:01:00 Pattern: YYYY-MM-DD hh:mm:ss Example: &#x60;\&quot;2016-02-01 10:10:10\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

