# AchTrans

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of account associated with the transaction | [default to null]
**Description** | **string** | Description for the transaction | [default to null]
**AchTransactionId** | **string** | A unique ID for an ACH transaction | [default to null]
**RejectCode** | **string** | The &lt;a href&#x3D;\&quot;ref:api-reference-events-ach-return-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ACH return code&lt;/a&gt;. | [default to null]
**Date** | **string** | The datetime for the transaction. In the format: &#x60;yyyy-mm-dd hh:mm:ss&#x60;. | [default to null]
**Status** | **string** | Status of transaction. See &lt;a href&#x3D;\&quot;ref:api-reference-ach-transaction-statuses\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ACH Transaction Statuses&lt;/a&gt;. | [default to null]
**AchRoutingNo** | **string** | ACH routing number | [default to null]
**AchAccountNo** | **string** | The ACH account number. May be masked depending on configuration. | [default to null]
**DebitCreditIndicator** | **string** | Indicator of whether the transaction was a credit or debit. &#x60;D&#x60; or &#x60;C&#x60;. | [default to null]
**Amount** | **string** | Transaction amount | [default to null]
**PmtRefNo** | **string** | A Galileo generated account number | [default to null]
**AchAccountId** | **string** | ID for the ACH account | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

