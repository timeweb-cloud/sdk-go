# AgentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | [**AgentModelSettings**](AgentModelSettings.md) |  | 
**SystemPrompt** | **string** | Системный промпт (инструкция) агента | 
**RefineQuery** | **bool** | Уточнять (переформулировать) запрос пользователя перед поиском по базам знаний | 
**MemoryEnabled** | Pointer to **bool** | Запоминать факты о каждом авторизованном пользователе и учитывать их в ответах. В виджете не работает. | [optional] [default to false]
**Widget** | Pointer to [**AgentWidgetSettings**](AgentWidgetSettings.md) |  | [optional] 
**WebSearchAllowedSites** | Pointer to **[]string** | Домены, которыми ограничивается веб-поиск агента (не более 5). Каждый элемент — домен (&#x60;timeweb.ru&#x60;) или домен с путём (&#x60;timeweb.ru/docs&#x60;). Если поле не передано при обновлении, сохранённый список не меняется. | [optional] 
**AdditionalModels** | Pointer to [**[]AgentAdditionalModel**](AgentAdditionalModel.md) | Настройки дополнительных моделей. Заполняется автоматически из поля &#x60;additional_models&#x60; запроса создания или обновления агента, передавать здесь не нужно. | [optional] 

## Methods

### NewAgentSettings

`func NewAgentSettings(model AgentModelSettings, systemPrompt string, refineQuery bool, ) *AgentSettings`

NewAgentSettings instantiates a new AgentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSettingsWithDefaults

`func NewAgentSettingsWithDefaults() *AgentSettings`

NewAgentSettingsWithDefaults instantiates a new AgentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AgentSettings) GetModel() AgentModelSettings`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentSettings) GetModelOk() (*AgentModelSettings, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentSettings) SetModel(v AgentModelSettings)`

SetModel sets Model field to given value.


### GetSystemPrompt

`func (o *AgentSettings) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *AgentSettings) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *AgentSettings) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.


### GetRefineQuery

`func (o *AgentSettings) GetRefineQuery() bool`

GetRefineQuery returns the RefineQuery field if non-nil, zero value otherwise.

### GetRefineQueryOk

`func (o *AgentSettings) GetRefineQueryOk() (*bool, bool)`

GetRefineQueryOk returns a tuple with the RefineQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefineQuery

`func (o *AgentSettings) SetRefineQuery(v bool)`

SetRefineQuery sets RefineQuery field to given value.


### GetMemoryEnabled

`func (o *AgentSettings) GetMemoryEnabled() bool`

GetMemoryEnabled returns the MemoryEnabled field if non-nil, zero value otherwise.

### GetMemoryEnabledOk

`func (o *AgentSettings) GetMemoryEnabledOk() (*bool, bool)`

GetMemoryEnabledOk returns a tuple with the MemoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryEnabled

`func (o *AgentSettings) SetMemoryEnabled(v bool)`

SetMemoryEnabled sets MemoryEnabled field to given value.

### HasMemoryEnabled

`func (o *AgentSettings) HasMemoryEnabled() bool`

HasMemoryEnabled returns a boolean if a field has been set.

### GetWidget

`func (o *AgentSettings) GetWidget() AgentWidgetSettings`

GetWidget returns the Widget field if non-nil, zero value otherwise.

### GetWidgetOk

`func (o *AgentSettings) GetWidgetOk() (*AgentWidgetSettings, bool)`

GetWidgetOk returns a tuple with the Widget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidget

`func (o *AgentSettings) SetWidget(v AgentWidgetSettings)`

SetWidget sets Widget field to given value.

### HasWidget

`func (o *AgentSettings) HasWidget() bool`

HasWidget returns a boolean if a field has been set.

### GetWebSearchAllowedSites

`func (o *AgentSettings) GetWebSearchAllowedSites() []string`

GetWebSearchAllowedSites returns the WebSearchAllowedSites field if non-nil, zero value otherwise.

### GetWebSearchAllowedSitesOk

`func (o *AgentSettings) GetWebSearchAllowedSitesOk() (*[]string, bool)`

GetWebSearchAllowedSitesOk returns a tuple with the WebSearchAllowedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchAllowedSites

`func (o *AgentSettings) SetWebSearchAllowedSites(v []string)`

SetWebSearchAllowedSites sets WebSearchAllowedSites field to given value.

### HasWebSearchAllowedSites

`func (o *AgentSettings) HasWebSearchAllowedSites() bool`

HasWebSearchAllowedSites returns a boolean if a field has been set.

### GetAdditionalModels

`func (o *AgentSettings) GetAdditionalModels() []AgentAdditionalModel`

GetAdditionalModels returns the AdditionalModels field if non-nil, zero value otherwise.

### GetAdditionalModelsOk

`func (o *AgentSettings) GetAdditionalModelsOk() (*[]AgentAdditionalModel, bool)`

GetAdditionalModelsOk returns a tuple with the AdditionalModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalModels

`func (o *AgentSettings) SetAdditionalModels(v []AgentAdditionalModel)`

SetAdditionalModels sets AdditionalModels field to given value.

### HasAdditionalModels

`func (o *AgentSettings) HasAdditionalModels() bool`

HasAdditionalModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


