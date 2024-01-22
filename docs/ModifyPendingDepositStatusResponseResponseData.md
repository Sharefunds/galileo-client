# ModifyPendingDepositStatusResponseResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepositTransactionId** | **string** | A unique system generated ID number that identifies the deposit | [default to null]
**ActionType** | **string** | &#x27;P&#x27; and &#x27;R&#x27; (Post and Return). Post &#x3D; pending (ACH) deposit, Return &#x3D; return the pending deposit | [default to null]
**CategoryType** | **string** | A &#x3D; Approve; W &#x3D; Watch; D &#x3D; Decline | [default to null]
**CategoryCode** | **string** | COF&#x3D;Questionable IAT Country; CRN&#x3D; Unauthorized IAT Country; L&#x3D;Large Xfer Amount &gt; 4000; LRG&#x3D;Large Xfer 4000; NAME&#x3D;Name miss match; TAX&#x3D;Tax | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

