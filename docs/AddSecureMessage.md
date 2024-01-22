# AddSecureMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to null]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to null]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to null]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to null]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN  Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to null]
**MsgSubject** | **string** | Secure message subject Pattern: Alphanumeric string including punctuation; maximum 255 characters Example: &#x60;\&quot;account login help pls\&quot;&#x60; | [optional] [default to null]
**MsgText** | **string** | Secure message body text Pattern: Alphanumeric string including punctuation; maximum 4000 characters Example: &#x60;\&quot;Please help me recover my password.\&quot;&#x60; | [optional] [default to null]
**ParentMsgId** | **string** | The msgId of the original cardholder message that you wish to respond to. Pattern: Positive integer value Example: &#x60;483383&#x60; | [optional] [default to null]
**Locale** | **string** | Sets customer language preference. Default is \&quot;en_US\&quot;. Pattern: Accepted values &#x3D; &#x60;en_US&#x60;, &#x60;es_US&#x60;, &#x60;fr_CA&#x60;, &#x60;en_CA&#x60;, &#x60;es_MX&#x60;, and &#x60;en_MX&#x60; Example: &#x60;\&quot;en_US\&quot;&#x60; | [optional] [default to LOCALE.EN_US]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

