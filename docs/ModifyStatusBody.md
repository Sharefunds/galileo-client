# ModifyStatusBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt;, &lt;&lt;glossary:PAN&gt;&gt; or &lt;&lt;glossary:CAD&gt;&gt; of the account. Pattern: PAN, PRN, or CAD Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**Type_** | **int32** | Specifies the account and/or card status to set. See the &lt;a href&#x3D;\&quot;ref:api-reference-modify-status-types\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Modify Status Types&lt;/a&gt; enumeration for valid values. Pattern: Predefined integer value range Example: &#x60;2&#x60; | [default to 2]
**StartDate** | [**time.Time**](time.Time.md) | Valid only with &#x60;type: 17&#x60; (freeze card). If you pass a value for &#x60;startDate&#x60;, you must also populate &#x60;endDate&#x60;. Default is current date-time. Pattern: &#x60;YYYY-MM-DD hh:mm:ss&#x60; Example: &#x60;\&quot;2020-02-02 12:22:02\&quot;&#x60; | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | Valid only with &#x60;type: 17&#x60; (freeze card). If you pass a value for &#x60;endDate&#x60;, you must also populate &#x60;startDate&#x60;. Must be later than &#x60;startDate&#x60;. Default is current date-time plus 24 hours. Pattern: &#x60;YYYY-MM-DD hh:mm:ss&#x60; Example: &#x60;\&quot;2020-02-12 10:00:00\&quot;&#x60; | [optional] [default to null]
**BypassRepFee** | **string** | Valid only when &#x60;type: 19&#x60;. Pass &#x60;Y&#x60; to bypass the card replacement fee (REP). Pattern: &#x60;Y&#x60; or &#x60;N&#x60; Example: &#x60;\&quot;Y\&quot;&#x60; | [optional] [default to null]
**CardNumberLastFour** | **string** | If multiple cards are associated with this account, use this value to ensure that the correct card is selected. Pattern: 4 digits Example: &#x60;\&quot;0789\&quot;&#x60; | [optional] [default to null]
**ClosureReason** | **string** | Specifies the account closure reason to be reported to credit bureaus: * &#x60;chargeoff&#x60; &amp;mdash; Unpaid, closed by issuer, and charged off * &#x60;collection&#x60; &amp;mdash; Unpaid and closed by issuer * &#x60;deleted_fraud&#x60; &amp;mdash; Account takeover fraud, closed by issuer * &#x60;inactivity&#x60; &amp;mdash; Paid and closed by issuer due to inactivity * &#x60;paid_user&#x60; &amp;mdash; Paid and closed at the customer&#x27;s request  Only valid when &#x60;type&#x60; contains &#x60;2&#x60;, &#x60;13&#x60;, &#x60;16&#x60; or &#x60;20&#x60;. Default: &#x60;paid_user&#x60;. Pattern: Strings shown above Example: &#x60;inactivity&#x60; | [optional] [default to null]
**BypassMailFee** | **string** | Controls whether to charge the express mail fee for the reissued card. Default: blank. Pattern: &#x60;Y&#x60; or blank Example: &#x60;Y&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

