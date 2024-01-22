# CreateAccountTransferBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the sending account. Pattern: PAN or PRN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**Amount** | **float32** | Currency amount as a whole or decimal amount.  Pattern: Positive integer or decimal number Example: &#x60;100.00&#x60;, &#x60;100&#x60;, or &#x60;100.73&#x60; | [default to 25.5]
**TransferToAccountNo** | **string** | The &lt;&lt;glossary:PAN&gt;&gt; or &lt;&lt;glossary:PRN&gt;&gt; of the receiving account. Pattern: PAN or PRN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**Message** | **string** | A message for the transfer recipient. The C2DSC provider parameter must be set to &#x60;1&#x60; to populate this parameter. If you do not populate this parameter, it defaults to the value in the BMNAM product parameter. Pattern: Max 2000 alphanumeric characters, including punctuation, special characters and &lt;a href&#x3D;\&quot;ref:api-reference-special-characters#international-characters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;international characters&lt;/a&gt; Example: &#x60;\&quot;Thanks again for lunch!\&quot;&#x60; | [optional] [default to null]
**SenderMessage** | **string** | A message for the transfer sender. The C2DSC provider parameter must be set to &#x60;1&#x60; to populate this parameter. If you do not populate this parameter, it defaults to the value in &#x60;message&#x60;. Pattern: Max 2000 alphanumeric characters, including punctuation, special characters and &lt;a href&#x3D;\&quot;ref:api-reference-special-characters#international-characters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;international characters&lt;/a&gt; Example: &#x60;\&quot;Thanks again for lunch!\&quot;&#x60; | [optional] [default to null]
**Type_** | **string** | Transaction type for the transfer. Use the values provided by Galileo for your program. Pattern: Max 3 alphanumeric characters, case-sensitive Example: &#x60;\&quot;MRM\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

