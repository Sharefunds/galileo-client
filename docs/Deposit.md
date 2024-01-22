# Deposit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | **string** | A fee or transaction charge | [default to null]
**InTs** | [**time.Time**](time.Time.md) | Timestamp for the initial creation of the record | [default to null]
**EffectiveDt** | [**time.Time**](time.Time.md) | Date when the ACH payment posts | [default to null]
**Name** | **string** | The name of the account receiving the deposit | [default to null]
**Efname** | **string** | Encrypted cardholder first name | [optional] [default to null]
**Elname** | **string** | Encrypted cardholder last name | [optional] [default to null]
**Xid** | **string** | An account ID that can be used instead of the PAN | [optional] [default to null]
**ProgId** | **string** | Identifier for the program associated with the account | [default to null]
**BatchHdr** | **string** | A record of a batch of transactions | [optional] [default to null]
**DestAcctNo** | **string** | The destination account for a pending deposit | [optional] [default to null]
**SourceInstId** | **string** | An ID (usually a bank routing number) for the institution that originated the deposit | [default to null]
**SourceInstName** | **string** | The name of the institution that originated the deposit | [default to null]
**SourceAcctNo** | **string** | An account number of the institution that originated the deposit | [optional] [default to null]
**Status** | **string** | The status of the deposit. See &lt;a href&#x3D;\&quot;ref:api-reference-deposit-status-codes\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Deposit Status Codes&lt;/a&gt;. | [default to null]
**TransType** | **string** | Transaction type | [default to null]
**AddendaRec** | **string** | Supplemental information to identify a deposit | [optional] [default to null]
**AchTransId** | **string** | A unique ID for an ACH transaction | [default to null]
**PmtRefNo** | **string** | Galileo-generated account number. Also called the PRN. | [optional] [default to null]
**Categories** | [**[]DepositCategories**](Deposit_categories.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

