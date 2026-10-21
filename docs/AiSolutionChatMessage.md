# AiSolutionChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** | ID сообщения | 
**ConversationId** | **string** | ID разговора | 
**Role** | **string** | Автор сообщения: &#x60;user&#x60;, &#x60;assistant&#x60; или &#x60;context-reset&#x60; — служебная отметка с пустым &#x60;content&#x60;, обозначающая место сброса контекста разговора | 
**Content** | **string** | Текст сообщения | 
**CreatedAt** | **time.Time** | Дата и время сообщения | 

## Methods

### NewAiSolutionChatMessage

`func NewAiSolutionChatMessage(entryId string, conversationId string, role string, content string, createdAt time.Time, ) *AiSolutionChatMessage`

NewAiSolutionChatMessage instantiates a new AiSolutionChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSolutionChatMessageWithDefaults

`func NewAiSolutionChatMessageWithDefaults() *AiSolutionChatMessage`

NewAiSolutionChatMessageWithDefaults instantiates a new AiSolutionChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *AiSolutionChatMessage) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *AiSolutionChatMessage) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *AiSolutionChatMessage) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetConversationId

`func (o *AiSolutionChatMessage) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *AiSolutionChatMessage) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *AiSolutionChatMessage) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetRole

`func (o *AiSolutionChatMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AiSolutionChatMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AiSolutionChatMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetContent

`func (o *AiSolutionChatMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AiSolutionChatMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AiSolutionChatMessage) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreatedAt

`func (o *AiSolutionChatMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiSolutionChatMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiSolutionChatMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


