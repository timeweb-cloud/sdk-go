# AiSolutionChatReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | **string** | ID разговора. У основного разговора решения — &#x60;solution-&lt;solution_id&gt;&#x60; | 
**Reply** | **string** | Ответ агента | 
**Model** | **NullableString** | Модель, сгенерировавшая ответ. &#x60;null&#x60;, если агент её не сообщил | 

## Methods

### NewAiSolutionChatReply

`func NewAiSolutionChatReply(conversationId string, reply string, model NullableString, ) *AiSolutionChatReply`

NewAiSolutionChatReply instantiates a new AiSolutionChatReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSolutionChatReplyWithDefaults

`func NewAiSolutionChatReplyWithDefaults() *AiSolutionChatReply`

NewAiSolutionChatReplyWithDefaults instantiates a new AiSolutionChatReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *AiSolutionChatReply) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *AiSolutionChatReply) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *AiSolutionChatReply) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetReply

`func (o *AiSolutionChatReply) GetReply() string`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *AiSolutionChatReply) GetReplyOk() (*string, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *AiSolutionChatReply) SetReply(v string)`

SetReply sets Reply field to given value.


### GetModel

`func (o *AiSolutionChatReply) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiSolutionChatReply) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiSolutionChatReply) SetModel(v string)`

SetModel sets Model field to given value.


### SetModelNil

`func (o *AiSolutionChatReply) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *AiSolutionChatReply) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


