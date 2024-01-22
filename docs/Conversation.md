# Conversation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgId** | **string** | ID of the retrieved cardholder message | [default to null]
**MsgSubject** | **string** | The subject of a secure message | [default to null]
**MsgDt** | [**time.Time**](time.Time.md) | The date a secure message was received | [default to null]
**Status** | **string** | The read status of a secure message | [default to null]
**ResponseDt** | [**time.Time**](time.Time.md) | The date of a response to the secure message | [default to null]
**Locale** | **string** | Language code in use locally | [default to null]
**MsgText** | **string** | The content of a secure message | [default to null]
**Replies** | [**[]ConversationReplies**](Conversation_replies.md) | List of replies to a secure message | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

