# AiGatewayResponseCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID модели из &#x60;GET /v1/models&#x60; | 
**Input** | **string** | Текст или массив входных элементов (сообщения, результаты вызова функций и т. п.) | 
**Instructions** | Pointer to **string** | Системные инструкции | [optional] 
**PreviousResponseId** | Pointer to **string** | ID сохранённого ответа, чтобы продолжить диалог | [optional] 
**Stream** | Pointer to **bool** | Отправлять ответ потоком типизированных событий Server-Sent Events | [optional] [default to false]
**Background** | Pointer to **bool** | Выполнить запрос в фоне: результат получайте через &#x60;GET /v1/responses/{response_id}&#x60; | [optional] 
**Store** | Pointer to **bool** | Сохранить ответ, чтобы получить его позже или продолжить диалог | [optional] 
**MaxOutputTokens** | Pointer to **int32** | Максимум выходных токенов | [optional] 
**Tools** | Pointer to **[]map[string]interface{}** | Инструменты, которые может вызвать модель | [optional] 
**Reasoning** | Pointer to [**AiGatewayResponseCreateRequestReasoning**](AiGatewayResponseCreateRequestReasoning.md) |  | [optional] 
**Text** | Pointer to **map[string]interface{}** | Настройки формата текста ответа | [optional] 
**Temperature** | Pointer to **float32** | Температура выборки | [optional] 

## Methods

### NewAiGatewayResponseCreateRequest

`func NewAiGatewayResponseCreateRequest(model string, input string, ) *AiGatewayResponseCreateRequest`

NewAiGatewayResponseCreateRequest instantiates a new AiGatewayResponseCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayResponseCreateRequestWithDefaults

`func NewAiGatewayResponseCreateRequestWithDefaults() *AiGatewayResponseCreateRequest`

NewAiGatewayResponseCreateRequestWithDefaults instantiates a new AiGatewayResponseCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiGatewayResponseCreateRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiGatewayResponseCreateRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiGatewayResponseCreateRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *AiGatewayResponseCreateRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AiGatewayResponseCreateRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AiGatewayResponseCreateRequest) SetInput(v string)`

SetInput sets Input field to given value.


### GetInstructions

`func (o *AiGatewayResponseCreateRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AiGatewayResponseCreateRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AiGatewayResponseCreateRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AiGatewayResponseCreateRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetPreviousResponseId

`func (o *AiGatewayResponseCreateRequest) GetPreviousResponseId() string`

GetPreviousResponseId returns the PreviousResponseId field if non-nil, zero value otherwise.

### GetPreviousResponseIdOk

`func (o *AiGatewayResponseCreateRequest) GetPreviousResponseIdOk() (*string, bool)`

GetPreviousResponseIdOk returns a tuple with the PreviousResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousResponseId

`func (o *AiGatewayResponseCreateRequest) SetPreviousResponseId(v string)`

SetPreviousResponseId sets PreviousResponseId field to given value.

### HasPreviousResponseId

`func (o *AiGatewayResponseCreateRequest) HasPreviousResponseId() bool`

HasPreviousResponseId returns a boolean if a field has been set.

### GetStream

`func (o *AiGatewayResponseCreateRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AiGatewayResponseCreateRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AiGatewayResponseCreateRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AiGatewayResponseCreateRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetBackground

`func (o *AiGatewayResponseCreateRequest) GetBackground() bool`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *AiGatewayResponseCreateRequest) GetBackgroundOk() (*bool, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *AiGatewayResponseCreateRequest) SetBackground(v bool)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *AiGatewayResponseCreateRequest) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetStore

`func (o *AiGatewayResponseCreateRequest) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *AiGatewayResponseCreateRequest) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *AiGatewayResponseCreateRequest) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *AiGatewayResponseCreateRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetMaxOutputTokens

`func (o *AiGatewayResponseCreateRequest) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *AiGatewayResponseCreateRequest) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *AiGatewayResponseCreateRequest) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *AiGatewayResponseCreateRequest) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### GetTools

`func (o *AiGatewayResponseCreateRequest) GetTools() []map[string]interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AiGatewayResponseCreateRequest) GetToolsOk() (*[]map[string]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AiGatewayResponseCreateRequest) SetTools(v []map[string]interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *AiGatewayResponseCreateRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetReasoning

`func (o *AiGatewayResponseCreateRequest) GetReasoning() AiGatewayResponseCreateRequestReasoning`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *AiGatewayResponseCreateRequest) GetReasoningOk() (*AiGatewayResponseCreateRequestReasoning, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *AiGatewayResponseCreateRequest) SetReasoning(v AiGatewayResponseCreateRequestReasoning)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *AiGatewayResponseCreateRequest) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### GetText

`func (o *AiGatewayResponseCreateRequest) GetText() map[string]interface{}`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AiGatewayResponseCreateRequest) GetTextOk() (*map[string]interface{}, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AiGatewayResponseCreateRequest) SetText(v map[string]interface{})`

SetText sets Text field to given value.

### HasText

`func (o *AiGatewayResponseCreateRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTemperature

`func (o *AiGatewayResponseCreateRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *AiGatewayResponseCreateRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *AiGatewayResponseCreateRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *AiGatewayResponseCreateRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


