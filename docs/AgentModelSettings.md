# AgentModelSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Temperature** | Pointer to **float32** | Контролирует случайность вывода модели. | [optional] 
**TopP** | Pointer to **float32** | Контролирует разнообразие вывода модели. | [optional] 
**MaxTokens** | Pointer to **float32** | Максимальное количество токенов в выводе. | [optional] 
**FrequencyPenalty** | Pointer to **float32** | Положительные значения штрафуют новые токены на основе их существующей частоты. | [optional] 
**PresencePenalty** | Pointer to **float32** | Положительные значения штрафуют новые токены на основе того, встречались ли они в тексте ранее. | [optional] 

## Methods

### NewAgentModelSettings

`func NewAgentModelSettings() *AgentModelSettings`

NewAgentModelSettings instantiates a new AgentModelSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentModelSettingsWithDefaults

`func NewAgentModelSettingsWithDefaults() *AgentModelSettings`

NewAgentModelSettingsWithDefaults instantiates a new AgentModelSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemperature

`func (o *AgentModelSettings) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *AgentModelSettings) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *AgentModelSettings) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *AgentModelSettings) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTopP

`func (o *AgentModelSettings) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *AgentModelSettings) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *AgentModelSettings) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *AgentModelSettings) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetMaxTokens

`func (o *AgentModelSettings) GetMaxTokens() float32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *AgentModelSettings) GetMaxTokensOk() (*float32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *AgentModelSettings) SetMaxTokens(v float32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *AgentModelSettings) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetFrequencyPenalty

`func (o *AgentModelSettings) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *AgentModelSettings) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *AgentModelSettings) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *AgentModelSettings) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### GetPresencePenalty

`func (o *AgentModelSettings) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *AgentModelSettings) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *AgentModelSettings) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *AgentModelSettings) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


