# AiGatewayChatCompletionChoicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Номер варианта | [optional] 
**Message** | Pointer to [**AiGatewayChatMessage**](AiGatewayChatMessage.md) |  | [optional] 
**FinishReason** | Pointer to **string** | Причина завершения: &#x60;stop&#x60;, &#x60;length&#x60; (достигнут лимит выходных токенов), &#x60;tool_calls&#x60; или &#x60;content_filter&#x60; | [optional] 

## Methods

### NewAiGatewayChatCompletionChoicesInner

`func NewAiGatewayChatCompletionChoicesInner() *AiGatewayChatCompletionChoicesInner`

NewAiGatewayChatCompletionChoicesInner instantiates a new AiGatewayChatCompletionChoicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayChatCompletionChoicesInnerWithDefaults

`func NewAiGatewayChatCompletionChoicesInnerWithDefaults() *AiGatewayChatCompletionChoicesInner`

NewAiGatewayChatCompletionChoicesInnerWithDefaults instantiates a new AiGatewayChatCompletionChoicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *AiGatewayChatCompletionChoicesInner) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AiGatewayChatCompletionChoicesInner) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AiGatewayChatCompletionChoicesInner) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *AiGatewayChatCompletionChoicesInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetMessage

`func (o *AiGatewayChatCompletionChoicesInner) GetMessage() AiGatewayChatMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiGatewayChatCompletionChoicesInner) GetMessageOk() (*AiGatewayChatMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiGatewayChatCompletionChoicesInner) SetMessage(v AiGatewayChatMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AiGatewayChatCompletionChoicesInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFinishReason

`func (o *AiGatewayChatCompletionChoicesInner) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *AiGatewayChatCompletionChoicesInner) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *AiGatewayChatCompletionChoicesInner) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *AiGatewayChatCompletionChoicesInner) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


