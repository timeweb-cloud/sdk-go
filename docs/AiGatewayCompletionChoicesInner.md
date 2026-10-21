# AiGatewayCompletionChoicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Номер варианта | [optional] 
**Text** | Pointer to **string** | Сгенерированный текст | [optional] 
**FinishReason** | Pointer to **string** | Причина завершения: &#x60;stop&#x60; или &#x60;length&#x60; | [optional] 

## Methods

### NewAiGatewayCompletionChoicesInner

`func NewAiGatewayCompletionChoicesInner() *AiGatewayCompletionChoicesInner`

NewAiGatewayCompletionChoicesInner instantiates a new AiGatewayCompletionChoicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayCompletionChoicesInnerWithDefaults

`func NewAiGatewayCompletionChoicesInnerWithDefaults() *AiGatewayCompletionChoicesInner`

NewAiGatewayCompletionChoicesInnerWithDefaults instantiates a new AiGatewayCompletionChoicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *AiGatewayCompletionChoicesInner) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AiGatewayCompletionChoicesInner) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AiGatewayCompletionChoicesInner) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *AiGatewayCompletionChoicesInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetText

`func (o *AiGatewayCompletionChoicesInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AiGatewayCompletionChoicesInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AiGatewayCompletionChoicesInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *AiGatewayCompletionChoicesInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetFinishReason

`func (o *AiGatewayCompletionChoicesInner) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *AiGatewayCompletionChoicesInner) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *AiGatewayCompletionChoicesInner) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *AiGatewayCompletionChoicesInner) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


