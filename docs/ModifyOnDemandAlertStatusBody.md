# ModifyOnDemandAlertStatusBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN  Example: &#x60;\&quot;074103447228\&quot;&#x60; | [default to 074103447228]
**ProdId** | **int32** | Galileo-generated identifier for a product. Pattern: Integer Example: &#x60;501&#x60; | [default to 501]
**MobilePhone** | **string** | Account holder&#x27;s mobile phone number. Pattern: Exactly 10 digits, no hyphens or other characters Example: &#x60;\&quot;8013656050\&quot;&#x60; | [default to 8013656050]
**MobileCarrierId** | **int32** | Account holder&#x27;s mobile carrier. Configurable list. Pattern: Configurable list Example: &#x60;8&#x60; | [default to 8]
**ActionType** | **string** | Type of action to perform. See the &lt;a href&#x3D;\&quot;ref:api-reference-on-demand-alert-statuses\&quot; target&#x3D;\&quot;_blank\&quot;&gt;On-Demand Alert Statuses&lt;/a&gt; enumeration for valid values. Pattern: One character Example: &#x60;\&quot;E\&quot;&#x60; | [default to ACTION_TYPE.E]
**VerificationPin** | **string** | When &#x60;actionType: C&#x60;, this value will be validated against what was sent to the cardholder via text message. Pattern: 4-digit string Example: &#x60;\&quot;1234\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

