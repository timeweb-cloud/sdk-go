# SendAiSolutionMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Текст сообщения, до 32 000 символов. Если не передать, агент сам отправит приветствие — так можно получить первую реплику чата | [optional] 
**ConversationId** | Pointer to **string** | ID разговора. Если не передать, сообщение уйдёт в основной разговор решения (&#x60;solution-&lt;solution_id&gt;&#x60;). Чтобы начать отдельный разговор, передайте новый ID. Буквы, цифры и дефисы, до 64 символов | [optional] 

## Methods

### NewSendAiSolutionMessage

`func NewSendAiSolutionMessage() *SendAiSolutionMessage`

NewSendAiSolutionMessage instantiates a new SendAiSolutionMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendAiSolutionMessageWithDefaults

`func NewSendAiSolutionMessageWithDefaults() *SendAiSolutionMessage`

NewSendAiSolutionMessageWithDefaults instantiates a new SendAiSolutionMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SendAiSolutionMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendAiSolutionMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendAiSolutionMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SendAiSolutionMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetConversationId

`func (o *SendAiSolutionMessage) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *SendAiSolutionMessage) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *SendAiSolutionMessage) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *SendAiSolutionMessage) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


