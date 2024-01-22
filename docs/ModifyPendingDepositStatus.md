# ModifyPendingDepositStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [optional] [default to null]
**ExternalAccountId** | **string** | Identifier supplied by the provider, which is not related to the Galileo system. This ID is stored in the Galileo system in association with this account and can be provided in the &lt;&lt;glossary:RDF&gt;&gt;s. Pattern: Max 30 alphanumeric characters Example: &#x60;\&quot;553b45sbs\&quot;&#x60; | [optional] [default to null]
**DepositTransactionId** | **int32** | The ACH transaction identifier (&#x60;ach_trans_id&#x60;), as returned by the &lt;a href&#x3D;\&quot;ref:post_getpendingdeposits\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Get Pending Deposits&lt;/a&gt; endpoint. Pattern: Positive integer Example: &#x60;6844743&#x60; | [default to 6844743]
**ActionType** | **string** | Action to take on the pending deposit: * &#x60;P&#x60; &amp;mdash; Post the deposit * &#x60;R&#x60; &amp;mdash; Return the deposit to the sender. When returning the deposit, &#x60;retCode&#x60; is **required**  Pattern: Single character Example: &#x60;\&quot;P\&quot;&#x60; | [default to ACTION_TYPE.P]
**CategoryCode** | **string** | Category to assign to the deposit. See &lt;a href&#x3D;\&quot;ref:api-reference-deposit-category-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Deposit Category Codes&lt;/a&gt; for valid values. Pattern: Three-letter code Example: &#x60;\&quot;TAX\&quot;&#x60; | [default to TAX]
**CategoryType** | **string** | Indicates the decision for future ACH deposits that match the program settings for the current deposit. Values are:   * &#x60;A&#x60; — Approve matching transactions. * &#x60;D&#x60; — Decline matching transactions. * &#x60;W&#x60; — Watch matching transactions and send for manual review.  Pattern: Must be &#x60;A&#x60;, &#x60;W&#x60; or &#x60;D&#x60; Example: &#x60;\&quot;D\&quot;&#x60; | [default to CATEGORY_TYPE.D]
**RetCode** | **string** | Reason for returning the deposit. This parameter is **required** when &#x60;actionType: R&#x60;. See the &lt;a href&#x3D;\&quot;ref:api-reference-modify-pending-deposit-status-return-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Return Codes&lt;/a&gt; table for valid values. Pattern: 3-character code Example: &#x60;\&quot;R02\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

