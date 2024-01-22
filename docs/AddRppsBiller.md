# AddRppsBiller

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN  Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**RppsBillerId** | **string** | The &#x60;rpps_biller_id&#x60; as returned by &lt;a href&#x3D;\&quot;ref:post_searchbillerdirectory\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Search Biller Directory&lt;/a&gt;. This value must be zero-padded on the left to be 10 digits, so you would pass &#x60;rpps_biller_id: 1234&#x60; as &#x60;rppsBillerId: 0000001234&#x60;. Pattern: Exactly 10 digits Example: &#x60;\&quot;0000001234\&quot;&#x60; | [default to 0000001234]
**BillerAccountNo** | **string** | The account number that the account holder has with the biller. When applicable, the account number is validated against the &#x60;biller_account_no_patterns&#x60; that were returned by &lt;a href&#x3D;\&quot;ref:post_searchbillerdirectory\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Search Biller Directory&lt;/a&gt;. You cannot edit this value with &lt;a href&#x3D;\&quot;ref:post_modifyrppsbiller\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Modify RPPS Biller&lt;/a&gt;. If the account holder submits an incorrect value, you must remove the biller and create the biller again with the correct account number. Pattern: Alphanumeric string including hyphens and spaces. Example: &#x60;\&quot;3333223323455555\&quot;&#x60; | [default to 3333223323455555]
**FrequencyType** | **string** | Frequency of the bill payment: * &#x60;O&#x60; &amp;mdash; One time * &#x60;W&#x60; &amp;mdash; Weekly * &#x60;M&#x60; &amp;mdash; Monthly * &#x60;Q&#x60; &amp;mdash; Quarterly * &#x60;Y&#x60; &amp;mdash; Yearly  If this value is not &#x60;O&#x60; then &#x60;nextDate&#x60; and &#x60;endDate&#x60; are **required**. Pattern: One letter Example: &#x60;\&quot;W\&quot;&#x60; | [optional] [default to null]
**NextDate** | [**time.Time**](time.Time.md) | The next date that the payment is scheduled. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2015-02-04\&quot;&#x60; | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The last date that the payment is scheduled. Can be up to five years in the future, so if today is 20 Jan 2020, this parameter can be no later than 20 Jan 2025. However, if today&#x27;s date is a leap day, such as 29 Feb 2020, the latest date can be 1 Mar 2025. Pattern: YYYY-MM-DD Example: &#x60;\&quot;2015-02-04\&quot;&#x60; | [optional] [default to null]
**Amount** | **float32** | Currency amount as a whole or decimal amount. Pattern: Positive integer or decimal number Example: &#x60;100.00&#x60;, &#x60;100&#x60;, or &#x60;100.73&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

