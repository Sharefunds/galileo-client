# SearchAccountsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiLogin** | **string** | Web service username, as provided by Galileo. Pattern: Max 50 characters Example: &#x60;\&quot;AbC123-9999\&quot;&#x60; | [optional] [default to AbC123-9999]
**ApiTransKey** | **string** | Web service password, as provided by Galileo. Pattern: Max 15 characters Example: &#x60;\&quot;4sb62fh6w4h7w34g\&quot;&#x60; | [optional] [default to 4sb62fh6w4h7w34g]
**ProviderId** | **int32** | Galileo-issued provider identifier. Pattern: Max 10 digits Example: &#x60;9999&#x60; | [optional] [default to 9999]
**TransactionId** | **string** | A unique provider-generated ID to identify this API call. A UUID is preferred. This value is used for &lt;a href&#x3D;\&quot;ref:idempotency\&quot; target&#x3D;\&quot;_blank\&quot;&gt;idempotency&lt;/a&gt;. Pattern: 60 characters or less Example: &#x60;\&quot;9845dk-39fdk3fj3-4483483478\&quot;&#x60; | [default to 123e4567-e89b-12d3-a456-426614174000]
**LogQueries** | **int32** |  | [optional] [default to LOG_QUERIES.0_]
**AccountNo** | **string** | The &lt;&lt;glossary:PRN&gt;&gt; or &lt;&lt;glossary:PAN&gt;&gt; of the account. Pattern: PAN or PRN Example: &#x60;\&quot;074103447228\&quot;&#x60; | [optional] [default to null]
**FirstName** | **string** | Account holder&#x27;s first name. &lt;a href&#x3D;\&quot;ref:api-reference-special-characters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Special character support&lt;/a&gt;. Pattern: 1&amp;ndash;40 characters: letters (&#x60;A-Z&#x60;, &#x60;a-z&#x60;), spaces, hyphens (&#x60;-&#x60;) and single quotes (&#x60;&#x27;&#x60;) Example: &#x60;\&quot;Ed\&quot;&#x60; | [optional] [default to null]
**MiddleName** | **string** | Account holder&#x27;s middle name. &lt;a href&#x3D;\&quot;ref:api-reference-special-characters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Special character support&lt;/a&gt;. Pattern: 1&amp;ndash;40 characters: letters (&#x60;A-Z&#x60;, &#x60;a-z&#x60;), spaces, hyphens (&#x60;-&#x60;) and single quotes (&#x60;&#x27;&#x60;) Example: &#x60;\&quot;W\&quot;&#x60; | [optional] [default to null]
**LastName** | **string** | Account holder&#x27;s last name. &lt;a href&#x3D;\&quot;ref:api-reference-special-characters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Special character support&lt;/a&gt;. Pattern: 1&amp;ndash;40 characters: letters (&#x60;A-Z&#x60;, &#x60;a-z&#x60;), spaces, hyphens (&#x60;-&#x60;) and single quotes (&#x60;&#x27;&#x60;) Example: &#x60;\&quot;Harley\&quot;&#x60; | [optional] [default to null]
**DateOfBirth** | **string** | Account holder&#x27;s birth date. This date is compared with the DOB product parameter when there is a minimum age requirement on the account. Pattern: YYYY-MM-DD Example: &#x60;\&quot;1980-01-01\&quot;&#x60; | [optional] [default to null]
**PostalCode** | **string** | Account holder&#x27;s postal code (ZIP code). Pattern: &#x60;12345&#x60;, &#x60;12345-1234&#x60;, or &#x60;K1A-1A1&#x60; Example: &#x60;\&quot;84121\&quot;&#x60; | [optional] [default to null]
**PrimaryPhone** | **string** | Account holder&#x27;s primary phone number. Pattern: Exactly 10 digits, no hyphens or other characters Example: &#x60;\&quot;8013656050\&quot;&#x60; | [optional] [default to null]
**OtherPhone** | **string** | Account holder&#x27;s secondary phone number. Pattern: Exactly 10 digits, no hyphens or other characters Example: &#x60;\&quot;8013656050\&quot;&#x60; | [optional] [default to null]
**MobilePhone** | **string** | Account holder&#x27;s mobile phone number Pattern: Exactly 10 digits, no hyphens or other characters Example: &#x60;\&quot;8013656050\&quot;&#x60; | [optional] [default to null]
**Email** | **string** | Account holder&#x27;s email. The Python implementation uses Marshmallow validation for this field. For the PHP implementation, contact Galileo. Pattern: Email address, 3&amp;ndash;63 characters Example: &#x60;\&quot;user@fakedomain.com\&quot;&#x60; | [optional] [default to null]
**UserData** | **string** | Data supplied by the provider, from sources external to the Galileo system. This data will be stored in the Galileo system in association with this account. The most common use of this parameter is to track affiliate-marketing traffic. The data in this field can be added to an &lt;&lt;glossary:RDF&gt;&gt;. Pattern: Max 50 alphanumeric characters and spaces Example: &#x60;\&quot;a4434gg44\&quot;&#x60; | [optional] [default to null]
**Eextid** | **string** | The &#x60;externalAccountId&#x60; that was passed in the account-creation call. Pattern: Max 30 alphanumeric characters Example: &#x60;\&quot;553b45sbs\&quot;&#x60; | [optional] [default to null]
**RecordCnt** | **int32** | The maximum number of records per page to be returned. Pattern: Positive integer &#x60;1-99999&#x60; Example: &#x60;100&#x60; | [optional] [default to null]
**Page** | **int32** | The number of the page to retrieve. Pattern: Integer value of &#x60;1&#x60; or greater Example: &#x60;3&#x60; | [optional] [default to 1]
**MobilePhoneCountryCode** | **string** | Account holder&#x27;s mobile phone country code. See &lt;a href&#x3D;\&quot;ref:api-reference-phone-validation\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Phone Validation&lt;/a&gt; for valid values. Pattern: A plus sign followed by 1 - 3 digits  Example: &#x60;\&quot;+1\&quot;&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

