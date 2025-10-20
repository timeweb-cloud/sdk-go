# UpdateAgentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to [**AgentModelSettings**](AgentModelSettings.md) |  | [optional] 
**SystemPrompt** | Pointer to **string** | Системный промпт | [optional] 
**RefineQuery** | Pointer to **bool** | Уточнять ли запрос перед обработкой | [optional] 
**Widget** | Pointer to [**AgentSettingsWidget**](AgentSettingsWidget.md) |  | [optional] 

## Methods

### NewUpdateAgentSettings

`func NewUpdateAgentSettings() *UpdateAgentSettings`

NewUpdateAgentSettings instantiates a new UpdateAgentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAgentSettingsWithDefaults

`func NewUpdateAgentSettingsWithDefaults() *UpdateAgentSettings`

NewUpdateAgentSettingsWithDefaults instantiates a new UpdateAgentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *UpdateAgentSettings) GetModel() AgentModelSettings`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UpdateAgentSettings) GetModelOk() (*AgentModelSettings, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UpdateAgentSettings) SetModel(v AgentModelSettings)`

SetModel sets Model field to given value.

### HasModel

`func (o *UpdateAgentSettings) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *UpdateAgentSettings) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *UpdateAgentSettings) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *UpdateAgentSettings) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *UpdateAgentSettings) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetRefineQuery

`func (o *UpdateAgentSettings) GetRefineQuery() bool`

GetRefineQuery returns the RefineQuery field if non-nil, zero value otherwise.

### GetRefineQueryOk

`func (o *UpdateAgentSettings) GetRefineQueryOk() (*bool, bool)`

GetRefineQueryOk returns a tuple with the RefineQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefineQuery

`func (o *UpdateAgentSettings) SetRefineQuery(v bool)`

SetRefineQuery sets RefineQuery field to given value.

### HasRefineQuery

`func (o *UpdateAgentSettings) HasRefineQuery() bool`

HasRefineQuery returns a boolean if a field has been set.

### GetWidget

`func (o *UpdateAgentSettings) GetWidget() AgentSettingsWidget`

GetWidget returns the Widget field if non-nil, zero value otherwise.

### GetWidgetOk

`func (o *UpdateAgentSettings) GetWidgetOk() (*AgentSettingsWidget, bool)`

GetWidgetOk returns a tuple with the Widget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidget

`func (o *UpdateAgentSettings) SetWidget(v AgentSettingsWidget)`

SetWidget sets Widget field to given value.

### HasWidget

`func (o *UpdateAgentSettings) HasWidget() bool`

HasWidget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


