# CreateAdjustment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **int64** | A unique integer ID for the transaction.  Pattern: 64-byte integer Example: &#x60;164736451002&#x60; | [default to null]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN  Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**Amount** | **float32** | Currency amount as a whole or decimal amount. Pattern: Positive integer or decimal number Example: &#x60;100.00&#x60;, &#x60;100&#x60;, or &#x60;100.73&#x60; | [default to 25.5]
**Type_** | **string** | The transaction type for the adjustment. Use the values provided by Galileo for your program. Pattern: 1- or 2-character alphanumeric, case-sensitive Example: &#x60;\&quot;le\&quot;&#x60; | [default to null]
**Description** | **string** | Description for the transaction. Pattern: 1&amp;ndash;40 alphanumeric characters, including punctuation Example: &#x60;\&quot;One time payroll load.\&quot;&#x60; | [optional] [default to null]
**DebitCreditIndicator** | **string** | Specifies whether this transaction credits or debits the account in &#x60;accountNo&#x60;: * &#x60;C&#x60; &amp;mdash; Credit * &#x60;D&#x60; &amp;mdash; Debit  Pattern: &#x60;D&#x60; or &#x60;C&#x60; Example: &#x60;\&quot;D\&quot;&#x60; | [default to DEBIT_CREDIT_INDICATOR.D]
**Location** | **string** | Unique location identifier (&#x60;location&#x60;) as returned by the &lt;a href&#x3D;\&quot;ref:post_createlocation\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Create Location&lt;/a&gt; endpoint.  This value is also returned by the &lt;a href&#x3D;\&quot;ref:post_getlocations\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Get Locations&lt;/a&gt; endpoint depending on the value of &#x60;locationType&#x60; when the location was created: * &#x60;0&#x60; or &#x60;2&#x60; &amp;mdash; Returned in the &#x60;location_id&#x60; field * &#x60;1&#x60; &amp;mdash; Returned in the &#x60;provider_specified_id&#x60; field  Pattern: Integer if &#x60;locationType: 0&#x60; or &#x60;locationType: 2&#x60;; max 15 characters if &#x60;locationType: 1&#x60; Example: &#x60;\&quot;a455-3483\&quot;&#x60; | [optional] [default to null]
**LocationType** | **int32** | Type of ID in &#x60;location&#x60;: * &#x60;0&#x60; &amp;mdash; Galileo location ID * &#x60;1&#x60; &amp;mdash; Partner location ID * &#x60;2&#x60; &amp;mdash; Don&#x27;t validate  Pattern: &#x60;0&#x60;, &#x60;1&#x60;, or &#x60;2&#x60; Example: &#x60;0&#x60; | [optional] [default to null]
**VerifyOnly** | **string** | Pass &#x60;1&#x60; to test the validity of the parameter data without committing the information to the Galileo system. Pattern: &#x60;0&#x60; or &#x60;1&#x60; Example: &#x60;0&#x60; | [optional] [default to null]
**IncludeRtfTransfer** | **string** | Specifies whether the &#x60;amount&#x60; in this transaction should be transferred to or from the RTF funding account that is associated with this RTF spending account. Default: &#x60;1&#x60; * &#x60;0&#x60; — Do not perform an RTF transfer for this amount * &#x60;1&#x60; — Perform an RTF transfer for this amount.  Pattern: &#x60;0&#x60; or &#x60;1&#x60;  Example: &#x60;\&quot;1\&quot;&#x60;  | [optional] [default to null]
**DisputeId** | **string** | External dispute ID, to be used by dispute providers only. Pattern: 1-16 alphanumeric characters Example: \&quot;21010000000D\&quot; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

