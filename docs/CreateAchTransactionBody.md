# CreateAchTransactionBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN  Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**AchAccountId** | **int32** | ACH account identifier (&#x60;ach_account_id&#x60;), as returned by the &lt;a href&#x3D;\&quot;ref:post_addachaccount\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Add ACH Account&lt;/a&gt; or &lt;a href&#x3D;\&quot;ref:post_getachaccounts\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Get ACH Accounts&lt;/a&gt; endpoint. Pattern: Positive integer Example: &#x60;354656&#x60; | [default to 354656]
**Amount** | **float32** | Currency amount as a whole or decimal amount. Pattern: Positive integer or decimal number Example: &#x60;100.00&#x60;, &#x60;100&#x60;, or &#x60;100.73&#x60; | [default to 25.5]
**Description** | **string** | Description for the ACH transfer. Pattern: Alphanumeric and punctuation, max 30 characters Example: &#x60;One time payroll load.&#x60; | [optional] [default to null]
**DebitCreditIndicator** | **string** | Specifies whether the recipient account is to be credited or debited. Default: &#x60;D&#x60;. * &#x60;C&#x60; &amp;mdash; Credit * &#x60;D&#x60; &amp;mdash; Debit  Pattern: &#x60;D&#x60; or &#x60;C&#x60; Example: &#x60;D&#x60; | [default to DEBIT_CREDIT_INDICATOR.D]
**SameDay** | **string** | Specifies whether this is a same-day transaction. Pattern: &#x60;Y&#x60; or &#x60;N&#x60; Example: &#x60;Y&#x60; | [optional] [default to null]
**IdentNumber** | **string** | Provider-supplied identifier, external to the Galileo system. This value is only required for the &#x60;CIE&#x60; and &#x60;WEB&#x60; &lt;a href&#x3D;\&quot;page:sec-codes-requirements-for-ach\&quot; target&#x3D;\&quot;_blank\&quot;&gt;SEC codes&lt;/a&gt;. Pattern: Max 22 characters Example: &#x60;\&quot;123456789\&quot;&#x60; | [optional] [default to null]
**ProcessorToken** | **string** | Obtained from Plaid if using Plaid integration. Checks the balance of the ACH account to verify that there are sufficient funds for an ACH debit. Pattern: String Example: &#x60;\&quot;processor-production-35cd43b-adfc-d8aa-b331-c9ba0fdha881\&quot;&#x60; | [optional] [default to null]
**HoldDays** | **int32** | The number of hold days to apply to this transaction, which overrides the hold days in product settings. This value applies only to outgoing ACH debits (&#x60;debitCreditIndicator: D&#x60;). This parameter is available only to clients who have obtained bank approval to use it. The ACOHD parameter must be set to use this parameter. Pattern: Positive integer Example: 2 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

