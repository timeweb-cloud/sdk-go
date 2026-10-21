# AiGatewayChatCompletionChunkChoicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Номер варианта | [optional] 
**Delta** | Pointer to **map[string]interface{}** | Очередная часть сообщения | [optional] 
**FinishReason** | Pointer to **NullableString** | Причина завершения: &#x60;stop&#x60;, &#x60;length&#x60; (достигнут лимит выходных токенов), &#x60;tool_calls&#x60; или &#x60;content_filter&#x60;. &#x60;null&#x60; до последнего чанка варианта | [optional] 

## Methods

### NewAiGatewayChatCompletionChunkChoicesInner

`func NewAiGatewayChatCompletionChunkChoicesInner() *AiGatewayChatCompletionChunkChoicesInner`

NewAiGatewayChatCompletionChunkChoicesInner instantiates a new AiGatewayChatCompletionChunkChoicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayChatCompletionChunkChoicesInnerWithDefaults

`func NewAiGatewayChatCompletionChunkChoicesInnerWithDefaults() *AiGatewayChatCompletionChunkChoicesInner`

NewAiGatewayChatCompletionChunkChoicesInnerWithDefaults instantiates a new AiGatewayChatCompletionChunkChoicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *AiGatewayChatCompletionChunkChoicesInner) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AiGatewayChatCompletionChunkChoicesInner) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AiGatewayChatCompletionChunkChoicesInner) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *AiGatewayChatCompletionChunkChoicesInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetDelta

`func (o *AiGatewayChatCompletionChunkChoicesInner) GetDelta() map[string]interface{}`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *AiGatewayChatCompletionChunkChoicesInner) GetDeltaOk() (*map[string]interface{}, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *AiGatewayChatCompletionChunkChoicesInner) SetDelta(v map[string]interface{})`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *AiGatewayChatCompletionChunkChoicesInner) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetFinishReason

`func (o *AiGatewayChatCompletionChunkChoicesInner) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *AiGatewayChatCompletionChunkChoicesInner) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *AiGatewayChatCompletionChunkChoicesInner) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *AiGatewayChatCompletionChunkChoicesInner) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.

### SetFinishReasonNil

`func (o *AiGatewayChatCompletionChunkChoicesInner) SetFinishReasonNil(b bool)`

 SetFinishReasonNil sets the value for FinishReason to be an explicit nil

### UnsetFinishReason
`func (o *AiGatewayChatCompletionChunkChoicesInner) UnsetFinishReason()`

UnsetFinishReason ensures that no value is present for FinishReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


