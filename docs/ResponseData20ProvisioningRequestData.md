# ResponseData20ProvisioningRequestData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationData** | **string** | Apple Pay only. The request&#x27;s activation data. This field is always Base64-encoded when PPALE is set to Y; otherwise, it could be Base64-encoded or hex-encoded. | [optional] [default to null]
**EncryptedPassData** | **string** | Apple Pay only. An encrypted JSON file containing the sensitive information needed to add a card to a wallet. This field is always Base64-encoded when PPALE is set to Y; otherwise, it could be Base64-encoded or hex-encoded. | [optional] [default to null]
**EphemeralPublicKey** | **string** | Apple Pay only. A generated key that is combined with a private key. This field is always Base64-encoded when PPALE is set to Y; otherwise, it could be Base64-encoded or hex-encoded. | [optional] [default to null]
**OpaquePaymentCard** | **string** | Google Pay only. A Base64-encoded or Base64url-encoded opaque string. | [optional] [default to null]
**Payload** | **string** | Samsung Pay only. A Base64-encoded or Base64url-encoded opaque string. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

