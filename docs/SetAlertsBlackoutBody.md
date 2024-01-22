# SetAlertsBlackoutBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN  Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**BeginTime** | **string** | Starting time for alerts to be sent. Set to whole hours only. Pattern: HH&#x60;:00&#x60; Example: &#x60;\&quot;07:00\&quot;&#x60; | [default to 07:00]
**EndTime** | **string** | Time of day alerts must end. Set to whole hours only. Pattern: HH&#x60;:00&#x60; Example: &#x60;\&quot;22:00\&quot;&#x60; | [default to 22:00]
**TimeZone** | **string** | Time zone offset from Arizona time for &#x60;beginTime&#x60; and &#x60;endTime&#x60;: * &#x60;-3&#x60; &amp;mdash; Hawaii time zone * &#x60;-2&#x60; &amp;mdash; Alaska time zone * &#x60;-1&#x60; &amp;mdash; Pacific time zone * &#x60;0&#x60; &amp;mdash; Mountain time zone * &#x60;1&#x60; &amp;mdash; Central time zone * &#x60;2&#x60; &amp;mdash; Eastern time zone * &#x60;3&#x60; &amp;mdash; Atlantic time zone  Pattern:  Integer between &#x60;-3&#x60; and &#x60;3&#x60; Example: &#x60;-1&#x60; | [default to TIME_ZONE.1_]
**DaylightSavings** | **string** | Specifies whether to observe daylight savings time. Pattern: &#x60;Y&#x60; or &#x60;N&#x60; Example: &#x60;\&quot;Y\&quot;&#x60; | [default to DAYLIGHT_SAVINGS.Y]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

