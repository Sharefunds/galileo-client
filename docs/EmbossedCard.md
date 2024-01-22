# EmbossedCard

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | [**time.Time**](time.Time.md) | Date the card was created. To receive a datetime in this field, set &#x60;CXPTM&#x60;. | [optional] [default to null]
**EmbossUuid** | **string** | UUID for an emboss record that has been sent to the embosser. When this field is present outside the &#x60;embossed_cards&#x60; object, it contains the most recent emboss UUID. | [optional] [default to null]
**EmbossDate** | **string** | Date the card was embossed | [optional] [default to null]
**ExpiryDate** | [**time.Time**](time.Time.md) | Expiration date of the card | [optional] [default to null]
**ProductId** | **string** | Product ID of card at the time of embossing | [optional] [default to null]
**Status** | **string** | Emboss status. For valid values see the &lt;a href&#x3D;\&quot;ref:api-reference-emboss-statuses\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Emboss Statuses&lt;/a&gt; enumeration. | [default to null]
**ShippingType** | **string** | The type of shipping used for the card | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

