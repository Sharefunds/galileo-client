# GetDirectDepositSwitchToken

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the local account. Do not use the &lt;&lt;glossary:CAD&gt;&gt;. Pattern: 12-digit or 16-digit numeric string Example: &#x60;\&quot;344101254935\&quot;&#x60; | [default to 074103447228]
**DdAccountNo** | **string** | PAN or PRN of the account that will receive the direct deposit. Do not use the CAD. Can be the same account number as &#x60;accountNo&#x60;.  Pattern: 12-digit or 16-digit numeric stringExample: &#x60;\&quot;722844300741\&quot;&#x60; | [default to 722844300741]
**DdRoutingNo** | **string** | Routing number for the account in &#x60;ddSwitchAccountNo&#x60;. Pattern: 9-digit routing number, including check digit Example: &#x60;\&quot;124001545\&quot;&#x60; | [default to 124001545]
**DdAccountType** | **string** | Type of account in &#x60;ddSwitchAccountNo&#x60;. Pattern: &#x60;checking&#x60; or &#x60;savings&#x60; Example: &#x60;\&quot;checking\&quot;&#x60; | [default to DD_ACCOUNT_TYPE.CHECKING]
**DdAccountDescription** | **string** | Description for the direct deposit account. Pattern: Max 50 characters: letters, numbers, spaces, hyphens, underscores. Example: &#x60;\&quot;Galileo Plus Checking account\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

