# AiGatewayCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID модели из &#x60;GET /v1/models&#x60; | 
**Prompt** | **string** | Текст или список текстов для продолжения | 
**MaxTokens** | Pointer to **int32** | Максимум выходных токенов | [optional] 
**Temperature** | Pointer to **float32** | Температура выборки | [optional] 
**TopP** | Pointer to **float32** | Порог nucleus sampling | [optional] 
**Stop** | Pointer to **string** | Строка или список строк, на которых генерация останавливается | [optional] 
**Stream** | Pointer to **bool** | Отправлять ответ потоком Server-Sent Events (&#x60;text/event-stream&#x60;) | [optional] [default to false]
**StreamOptions** | Pointer to [**AiGatewayChatCompletionRequestStreamOptions**](AiGatewayChatCompletionRequestStreamOptions.md) |  | [optional] 

## Methods

### NewAiGatewayCompletionRequest

`func NewAiGatewayCompletionRequest(model string, prompt string, ) *AiGatewayCompletionRequest`

NewAiGatewayCompletionRequest instantiates a new AiGatewayCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayCompletionRequestWithDefaults

`func NewAiGatewayCompletionRequestWithDefaults() *AiGatewayCompletionRequest`

NewAiGatewayCompletionRequestWithDefaults instantiates a new AiGatewayCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiGatewayCompletionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiGatewayCompletionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiGatewayCompletionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetPrompt

`func (o *AiGatewayCompletionRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiGatewayCompletionRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiGatewayCompletionRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetMaxTokens

`func (o *AiGatewayCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *AiGatewayCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *AiGatewayCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *AiGatewayCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetTemperature

`func (o *AiGatewayCompletionRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *AiGatewayCompletionRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *AiGatewayCompletionRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *AiGatewayCompletionRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTopP

`func (o *AiGatewayCompletionRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *AiGatewayCompletionRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *AiGatewayCompletionRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *AiGatewayCompletionRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetStop

`func (o *AiGatewayCompletionRequest) GetStop() string`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *AiGatewayCompletionRequest) GetStopOk() (*string, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *AiGatewayCompletionRequest) SetStop(v string)`

SetStop sets Stop field to given value.

### HasStop

`func (o *AiGatewayCompletionRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetStream

`func (o *AiGatewayCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AiGatewayCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AiGatewayCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AiGatewayCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStreamOptions

`func (o *AiGatewayCompletionRequest) GetStreamOptions() AiGatewayChatCompletionRequestStreamOptions`

GetStreamOptions returns the StreamOptions field if non-nil, zero value otherwise.

### GetStreamOptionsOk

`func (o *AiGatewayCompletionRequest) GetStreamOptionsOk() (*AiGatewayChatCompletionRequestStreamOptions, bool)`

GetStreamOptionsOk returns a tuple with the StreamOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamOptions

`func (o *AiGatewayCompletionRequest) SetStreamOptions(v AiGatewayChatCompletionRequestStreamOptions)`

SetStreamOptions sets StreamOptions field to given value.

### HasStreamOptions

`func (o *AiGatewayCompletionRequest) HasStreamOptions() bool`

HasStreamOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


