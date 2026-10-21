# AiGatewayChatCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID модели из &#x60;GET /v1/models&#x60; | 
**Messages** | [**[]AiGatewayChatMessage**](AiGatewayChatMessage.md) | История диалога целиком | 
**Stream** | Pointer to **bool** | Отправлять ответ потоком Server-Sent Events (&#x60;text/event-stream&#x60;) | [optional] [default to false]
**StreamOptions** | Pointer to [**AiGatewayChatCompletionRequestStreamOptions**](AiGatewayChatCompletionRequestStreamOptions.md) |  | [optional] 
**MaxCompletionTokens** | Pointer to **int32** | Максимум выходных токенов, включая рассуждения | [optional] 
**MaxTokens** | Pointer to **int32** | Устаревший синоним &#x60;max_completion_tokens&#x60; | [optional] 
**Temperature** | Pointer to **float32** | Температура выборки | [optional] 
**TopP** | Pointer to **float32** | Порог nucleus sampling | [optional] 
**N** | Pointer to **int32** | Количество вариантов ответа | [optional] [default to 1]
**Stop** | Pointer to **string** | Строка или список строк, на которых генерация останавливается | [optional] 
**Tools** | Pointer to **[]map[string]interface{}** | Функции, которые может вызвать модель | [optional] 
**ToolChoice** | Pointer to **string** | Режим вызова функций: &#x60;none&#x60;, &#x60;auto&#x60;, &#x60;required&#x60; или объект с конкретной функцией | [optional] 
**ResponseFormat** | Pointer to **map[string]interface{}** | Формат ответа: &#x60;{\&quot;type\&quot;: \&quot;json_object\&quot;}&#x60; или &#x60;{\&quot;type\&quot;: \&quot;json_schema\&quot;, \&quot;json_schema\&quot;: {...}}&#x60; | [optional] 
**ReasoningEffort** | Pointer to **string** | Уровень рассуждений для моделей с рассуждениями, например &#x60;low&#x60;, &#x60;high&#x60; | [optional] 
**User** | Pointer to **string** | Идентификатор вашего конечного пользователя | [optional] 

## Methods

### NewAiGatewayChatCompletionRequest

`func NewAiGatewayChatCompletionRequest(model string, messages []AiGatewayChatMessage, ) *AiGatewayChatCompletionRequest`

NewAiGatewayChatCompletionRequest instantiates a new AiGatewayChatCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayChatCompletionRequestWithDefaults

`func NewAiGatewayChatCompletionRequestWithDefaults() *AiGatewayChatCompletionRequest`

NewAiGatewayChatCompletionRequestWithDefaults instantiates a new AiGatewayChatCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiGatewayChatCompletionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiGatewayChatCompletionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiGatewayChatCompletionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetMessages

`func (o *AiGatewayChatCompletionRequest) GetMessages() []AiGatewayChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AiGatewayChatCompletionRequest) GetMessagesOk() (*[]AiGatewayChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AiGatewayChatCompletionRequest) SetMessages(v []AiGatewayChatMessage)`

SetMessages sets Messages field to given value.


### GetStream

`func (o *AiGatewayChatCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AiGatewayChatCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AiGatewayChatCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AiGatewayChatCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStreamOptions

`func (o *AiGatewayChatCompletionRequest) GetStreamOptions() AiGatewayChatCompletionRequestStreamOptions`

GetStreamOptions returns the StreamOptions field if non-nil, zero value otherwise.

### GetStreamOptionsOk

`func (o *AiGatewayChatCompletionRequest) GetStreamOptionsOk() (*AiGatewayChatCompletionRequestStreamOptions, bool)`

GetStreamOptionsOk returns a tuple with the StreamOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamOptions

`func (o *AiGatewayChatCompletionRequest) SetStreamOptions(v AiGatewayChatCompletionRequestStreamOptions)`

SetStreamOptions sets StreamOptions field to given value.

### HasStreamOptions

`func (o *AiGatewayChatCompletionRequest) HasStreamOptions() bool`

HasStreamOptions returns a boolean if a field has been set.

### GetMaxCompletionTokens

`func (o *AiGatewayChatCompletionRequest) GetMaxCompletionTokens() int32`

GetMaxCompletionTokens returns the MaxCompletionTokens field if non-nil, zero value otherwise.

### GetMaxCompletionTokensOk

`func (o *AiGatewayChatCompletionRequest) GetMaxCompletionTokensOk() (*int32, bool)`

GetMaxCompletionTokensOk returns a tuple with the MaxCompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCompletionTokens

`func (o *AiGatewayChatCompletionRequest) SetMaxCompletionTokens(v int32)`

SetMaxCompletionTokens sets MaxCompletionTokens field to given value.

### HasMaxCompletionTokens

`func (o *AiGatewayChatCompletionRequest) HasMaxCompletionTokens() bool`

HasMaxCompletionTokens returns a boolean if a field has been set.

### GetMaxTokens

`func (o *AiGatewayChatCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *AiGatewayChatCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *AiGatewayChatCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *AiGatewayChatCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetTemperature

`func (o *AiGatewayChatCompletionRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *AiGatewayChatCompletionRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *AiGatewayChatCompletionRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *AiGatewayChatCompletionRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTopP

`func (o *AiGatewayChatCompletionRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *AiGatewayChatCompletionRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *AiGatewayChatCompletionRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *AiGatewayChatCompletionRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetN

`func (o *AiGatewayChatCompletionRequest) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AiGatewayChatCompletionRequest) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AiGatewayChatCompletionRequest) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *AiGatewayChatCompletionRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### GetStop

`func (o *AiGatewayChatCompletionRequest) GetStop() string`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *AiGatewayChatCompletionRequest) GetStopOk() (*string, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *AiGatewayChatCompletionRequest) SetStop(v string)`

SetStop sets Stop field to given value.

### HasStop

`func (o *AiGatewayChatCompletionRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetTools

`func (o *AiGatewayChatCompletionRequest) GetTools() []map[string]interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AiGatewayChatCompletionRequest) GetToolsOk() (*[]map[string]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AiGatewayChatCompletionRequest) SetTools(v []map[string]interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *AiGatewayChatCompletionRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetToolChoice

`func (o *AiGatewayChatCompletionRequest) GetToolChoice() string`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *AiGatewayChatCompletionRequest) GetToolChoiceOk() (*string, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *AiGatewayChatCompletionRequest) SetToolChoice(v string)`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *AiGatewayChatCompletionRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### GetResponseFormat

`func (o *AiGatewayChatCompletionRequest) GetResponseFormat() map[string]interface{}`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *AiGatewayChatCompletionRequest) GetResponseFormatOk() (*map[string]interface{}, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *AiGatewayChatCompletionRequest) SetResponseFormat(v map[string]interface{})`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *AiGatewayChatCompletionRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetReasoningEffort

`func (o *AiGatewayChatCompletionRequest) GetReasoningEffort() string`

GetReasoningEffort returns the ReasoningEffort field if non-nil, zero value otherwise.

### GetReasoningEffortOk

`func (o *AiGatewayChatCompletionRequest) GetReasoningEffortOk() (*string, bool)`

GetReasoningEffortOk returns a tuple with the ReasoningEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningEffort

`func (o *AiGatewayChatCompletionRequest) SetReasoningEffort(v string)`

SetReasoningEffort sets ReasoningEffort field to given value.

### HasReasoningEffort

`func (o *AiGatewayChatCompletionRequest) HasReasoningEffort() bool`

HasReasoningEffort returns a boolean if a field has been set.

### GetUser

`func (o *AiGatewayChatCompletionRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AiGatewayChatCompletionRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AiGatewayChatCompletionRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AiGatewayChatCompletionRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


