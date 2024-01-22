# SetAccountLevelMerchantControlBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; for the account. Do not use the &lt;&lt;glossary:CAD&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt;. Pattern: PRN (12-digit numeric string) Example: &#x60;\&quot;344101254935\&quot;&#x60; | [default to 344101254935]
**MerchantId** | **string** | The merchant ID. Not case-sensitive. Pattern: Max 15 alphanumeric characters Example: &#x60;\&quot;000534LECR08003\&quot;&#x60; | [default to 000534LECR08003]
**AllowDeny** | **string** | Whether the merchant ID is ALLOW (&#x60;a&#x60;) or DENY (&#x60;d&#x60;). This parameter is required for new controls but optional for updates. Pattern: &#x60;a&#x60; or &#x60;d&#x60; Example: &#x60;\&quot;d\&quot;&#x60; | [optional] [default to null]
**StartDate** | **string** | Start date-time for the control. Can be set up to six months in the future. Cannot be in the past. Must be earlier than &#x60;endDate&#x60;. Default: Current &lt;a href&#x3D;\&quot;ref:galileo-system-time\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Galileo system time&lt;/a&gt;  Pattern: YYYY-MM-DD hh:mm:ss or YYYY-MM-DD Example: &#x60;\&quot;2022-03-03 00:00:00\&quot;&#x60; | [optional] [default to null]
**EndDate** | **string** | End date-time for the control. Must be later than &#x60;startDate&#x60;. Cannot be in the past. Default: &#x60;3000-01-01 00:00:00&#x60;.  Pattern: YYYY-MM-DD hh:mm:ss or YYYY-MM-DD Example: &#x60;\&quot;2022-03-10 23:59:59\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

