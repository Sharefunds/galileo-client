/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The provisioning request data to be directly sent to the wallet-provider
type ResponseData20ProvisioningRequestData struct {
	// Apple Pay only. The request's activation data. This field is always Base64-encoded when PPALE is set to Y; otherwise, it could be Base64-encoded or hex-encoded.
	ActivationData string `json:"activationData,omitempty"`
	// Apple Pay only. An encrypted JSON file containing the sensitive information needed to add a card to a wallet. This field is always Base64-encoded when PPALE is set to Y; otherwise, it could be Base64-encoded or hex-encoded.
	EncryptedPassData string `json:"encryptedPassData,omitempty"`
	// Apple Pay only. A generated key that is combined with a private key. This field is always Base64-encoded when PPALE is set to Y; otherwise, it could be Base64-encoded or hex-encoded.
	EphemeralPublicKey string `json:"ephemeralPublicKey,omitempty"`
	// Google Pay only. A Base64-encoded or Base64url-encoded opaque string.
	OpaquePaymentCard string `json:"opaquePaymentCard,omitempty"`
	// Samsung Pay only. A Base64-encoded or Base64url-encoded opaque string.
	Payload string `json:"payload,omitempty"`
}