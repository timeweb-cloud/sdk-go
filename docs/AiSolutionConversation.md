# AiSolutionConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | **string** | ID разговора | 
**StartedAt** | **time.Time** | Дата и время первого сообщения | 
**LastMessageAt** | **time.Time** | Дата и время последнего сообщения | 
**MessagesCount** | **int32** | Количество сообщений в разговоре | 

## Methods

### NewAiSolutionConversation

`func NewAiSolutionConversation(conversationId string, startedAt time.Time, lastMessageAt time.Time, messagesCount int32, ) *AiSolutionConversation`

NewAiSolutionConversation instantiates a new AiSolutionConversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSolutionConversationWithDefaults

`func NewAiSolutionConversationWithDefaults() *AiSolutionConversation`

NewAiSolutionConversationWithDefaults instantiates a new AiSolutionConversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *AiSolutionConversation) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *AiSolutionConversation) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *AiSolutionConversation) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetStartedAt

`func (o *AiSolutionConversation) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AiSolutionConversation) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AiSolutionConversation) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLastMessageAt

`func (o *AiSolutionConversation) GetLastMessageAt() time.Time`

GetLastMessageAt returns the LastMessageAt field if non-nil, zero value otherwise.

### GetLastMessageAtOk

`func (o *AiSolutionConversation) GetLastMessageAtOk() (*time.Time, bool)`

GetLastMessageAtOk returns a tuple with the LastMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageAt

`func (o *AiSolutionConversation) SetLastMessageAt(v time.Time)`

SetLastMessageAt sets LastMessageAt field to given value.


### GetMessagesCount

`func (o *AiSolutionConversation) GetMessagesCount() int32`

GetMessagesCount returns the MessagesCount field if non-nil, zero value otherwise.

### GetMessagesCountOk

`func (o *AiSolutionConversation) GetMessagesCountOk() (*int32, bool)`

GetMessagesCountOk returns a tuple with the MessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesCount

`func (o *AiSolutionConversation) SetMessagesCount(v int32)`

SetMessagesCount sets MessagesCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


