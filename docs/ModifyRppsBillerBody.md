# ModifyRppsBillerBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN  Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**BillerId** | **int32** | Identifier for the biller (&#x60;biller_id&#x60;) as returned by the &lt;a href&#x3D;\&quot;ref:post_addrppsbiller\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Add RPPS Biller&lt;/a&gt; or &lt;a href&#x3D;\&quot;ref:post_getbillers\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Get Billers&lt;/a&gt; endpoint. Do not use the &#x60;rpps_biller_id&#x60; from Search Biller Directory. Pattern: Integer Example: &#x60;2982&#x60; | [default to 2982]
**FrequencyType** | **string** | Frequency of the bill payment: * &#x60;O&#x60; &amp;mdash; One time * &#x60;W&#x60; &amp;mdash; Weekly * &#x60;M&#x60; &amp;mdash; Monthly * &#x60;Q&#x60; &amp;mdash; Quarterly * &#x60;Y&#x60; &amp;mdash; Yearly  If this value is not &#x60;O&#x60; then &#x60;nextDate&#x60; and &#x60;endDate&#x60; are **required**. Pattern: One letter Example: &#x60;\&quot;W\&quot;&#x60; | [default to FREQUENCY_TYPE.W]
**NextDate** | [**time.Time**](time.Time.md) | The next date that the payment is scheduled. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2015-02-04\&quot;&#x60; | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The last date that the payment is scheduled. Can be up to five years in the future, so if today is 20 Jan 2020, this parameter can be no later than 20 Jan 2025. However, if today&#x27;s date is a leap day, such as 29 Feb 2020, the latest date can be 1 Mar 2025. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2015-02-04\&quot;&#x60; | [optional] [default to null]
**Amount** | **float32** | Currency amount as a whole or decimal amount. Pattern: Positive integer or decimal number Example: &#x60;100.00&#x60;, &#x60;100&#x60;, or &#x60;100.73&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

