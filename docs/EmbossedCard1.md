# EmbossedCard1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmbossUuid** | **string** | UUID for an emboss record that has been sent to the embosser. When this field is present outside the &#x60;embossed_cards&#x60; object, it contains the most recent emboss UUID. | [optional] [default to null]
**Status** | **string** | Emboss status. For valid values see the &lt;a href&#x3D;\&quot;ref:api-reference-emboss-statuses\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Emboss Statuses&lt;/a&gt; enumeration. | [default to null]
**CreatedDate** | **string** | Date the card was created. To receive a datetime in this field, set &#x60;CXPTM&#x60;. | [default to null]
**EmbossDate** | **string** | The date the card was sent to be manufactured | [default to null]
**ExpiryDate** | [**time.Time**](time.Time.md) | Expiration date of the card | [default to null]
**ProductId** | **string** | Product ID of card at time of embossing | [optional] [default to null]
**ShippingType** | **string** | The shipping type used to ship the cards to the cardholder, Standard or Express Mail | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

