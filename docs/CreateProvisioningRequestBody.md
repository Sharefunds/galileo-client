# CreateProvisioningRequestBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt;, &lt;&lt;glossary:PAN&gt;&gt; or &lt;&lt;glossary:CAD&gt;&gt; of the account. For card-specific endpoints such as this one, the CAD is preferred. Do not use the PRN if more than one card has ever been associated with this account. Pattern: PAN, PRN, or CAD Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**WalletProvider** | **int32** | The type of wallet to provision: * &#x60;1&#x60; &amp;mdash; Apple Pay * &#x60;2&#x60; &amp;mdash; Google Pay * &#x60;3&#x60; &amp;mdash; Samsung Pay  Pattern: One digit Example: &#x60;2&#x60; | [default to WALLET_PROVIDER.2_]
**Cert1** | **string** | The leaf certificate returned by the wallet provider that was signed using &#x60;cert2&#x60;. Should be the hexlified (case insensitive) binary data of the certificate. See &lt;a href&#x3D;\&quot;doc:creating-a-provisioning-request#certificate-formatting\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Certificate formatting&lt;/a&gt; for more information. Required for Apple Pay; not required for Samsung Pay or Google Pay. | [optional] [default to null]
**Cert2** | **string** | The subordinate certificate returned by the wallet provider that was signed using the wallet provider’s Certificate Authority (CA) certificate. Should be the hexlified (case insensitive) binary data of the certificate. See &lt;a href&#x3D;\&quot;doc:creating-a-provisioning-request#certificate-formatting\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Certificate formatting&lt;/a&gt; for more information. Required for Apple Pay; not required for Samsung Pay or Google Pay. | [optional] [default to null]
**Nonce** | **string** | The hexlified (case insensitive) nonce value returned by Apple Pay. **Required** for Apple Pay; not required for Google Pay or Samsung Pay. Pattern: 8 character nonce value Example: &#x60;\&quot;9c023092\&quot;&#x60; | [optional] [default to null]
**NonceSignature** | **string** | The hexlified (case insensitive) nonce signature value returned by Apple Pay. **Required** for Apple Pay; not required for Google Pay or Samsung Pay. Pattern: Nonce signature value Example: &#x60;\&quot;4082f883ae62...\&quot;&#x60; | [optional] [default to null]
**ClientWalletAccountID** | **string** | Value returned by Google Pay or Samsung Pay for use with Visa. **Required** for Visa + Samsung Pay or Google Pay; not required for Visa + Apple Pay. | [optional] [default to null]
**ClientDeviceID** | **string** | Value returned by Google Pay or Samsung Pay for use with Visa. **Required** for Visa + Samsung Pay or Google Pay; not required for Visa + Apple Pay. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

