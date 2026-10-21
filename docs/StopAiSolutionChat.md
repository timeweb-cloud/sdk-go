# StopAiSolutionChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | Pointer to **string** | ID разговора. Если не передать, останавливается текущий ход основного разговора решения. Буквы, цифры и дефисы, до 64 символов | [optional] 

## Methods

### NewStopAiSolutionChat

`func NewStopAiSolutionChat() *StopAiSolutionChat`

NewStopAiSolutionChat instantiates a new StopAiSolutionChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopAiSolutionChatWithDefaults

`func NewStopAiSolutionChatWithDefaults() *StopAiSolutionChat`

NewStopAiSolutionChatWithDefaults instantiates a new StopAiSolutionChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *StopAiSolutionChat) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *StopAiSolutionChat) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *StopAiSolutionChat) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *StopAiSolutionChat) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


