# ResponseData126PendingDeposits

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | **float32** | A fee or transaction charge | [default to null]
**InTs** | [**time.Time**](time.Time.md) | An initial timestamp for the creation of the record | [default to null]
**EffectiveDt** | [**time.Time**](time.Time.md) | A timestamp for an ACH record that specifies when the payment posts | [default to null]
**Name** | **string** | The name of the account receiving the flagged deposit | [default to null]
**Xid** | **string** | An account ID that can be used instead of the PAN or other restricted information | [default to null]
**ProgId** | **string** | An ID number unique to a program | [default to null]
**BatchHdr** | **string** | A record of a batch of transactions | [optional] [default to null]
**DestAcctNo** | **string** | The destination account for a flagged pending deposit | [default to null]
**SourceInstId** | **string** | An ID (Usually a bank routing number) for the institution that originated the deposit | [optional] [default to null]
**SourceInstName** | **string** | The name of the institution that originated the deposit | [optional] [default to null]
**Status** | **string** | The status of the deposit. See &lt;a href&#x3D;\&quot;ref:api-reference-deposit-status-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Deposit Status Codes&lt;/a&gt;. | [optional] [default to null]
**TransType** | **string** | Transaction type | [optional] [default to null]
**AddendaRec** | **string** | Supplemental information to identify a deposit | [optional] [default to null]
**Categories** | [**[]PendingDepositCategories**](PendingDeposit_categories.md) | List containing information about categories | [optional] [default to null]
**AchTransId** | **string** | A unique ID for an ACH transaction | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

