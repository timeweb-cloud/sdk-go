# CreateAgentV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название агента | 
**Description** | Pointer to **string** | Описание агента | [optional] 
**AccessType** | **string** | Тип доступа к агенту | 
**ModelId** | **float32** | ID основной модели | 
**TokenLimit** | Pointer to **float32** | Дневной лимит токенов для агента (0 — без лимита) | [optional] 
**Settings** | [**AgentSettings**](AgentSettings.md) |  | 
**ProjectId** | Pointer to **float32** | ID проекта | [optional] 
**AdditionalModelIds** | Pointer to **[]float32** | Список ID дополнительных моделей агента | [optional] 
**IsWebSearchEnabled** | Pointer to **bool** | Признак использования веб-поиска агентом | [optional] 

## Methods

### NewCreateAgentV2

`func NewCreateAgentV2(name string, accessType string, modelId float32, settings AgentSettings, ) *CreateAgentV2`

NewCreateAgentV2 instantiates a new CreateAgentV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAgentV2WithDefaults

`func NewCreateAgentV2WithDefaults() *CreateAgentV2`

NewCreateAgentV2WithDefaults instantiates a new CreateAgentV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAgentV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAgentV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAgentV2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAgentV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAgentV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAgentV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAgentV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessType

`func (o *CreateAgentV2) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CreateAgentV2) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CreateAgentV2) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetModelId

`func (o *CreateAgentV2) GetModelId() float32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CreateAgentV2) GetModelIdOk() (*float32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CreateAgentV2) SetModelId(v float32)`

SetModelId sets ModelId field to given value.


### GetTokenLimit

`func (o *CreateAgentV2) GetTokenLimit() float32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *CreateAgentV2) GetTokenLimitOk() (*float32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *CreateAgentV2) SetTokenLimit(v float32)`

SetTokenLimit sets TokenLimit field to given value.

### HasTokenLimit

`func (o *CreateAgentV2) HasTokenLimit() bool`

HasTokenLimit returns a boolean if a field has been set.

### GetSettings

`func (o *CreateAgentV2) GetSettings() AgentSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateAgentV2) GetSettingsOk() (*AgentSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateAgentV2) SetSettings(v AgentSettings)`

SetSettings sets Settings field to given value.


### GetProjectId

`func (o *CreateAgentV2) GetProjectId() float32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateAgentV2) GetProjectIdOk() (*float32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateAgentV2) SetProjectId(v float32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateAgentV2) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetAdditionalModelIds

`func (o *CreateAgentV2) GetAdditionalModelIds() []float32`

GetAdditionalModelIds returns the AdditionalModelIds field if non-nil, zero value otherwise.

### GetAdditionalModelIdsOk

`func (o *CreateAgentV2) GetAdditionalModelIdsOk() (*[]float32, bool)`

GetAdditionalModelIdsOk returns a tuple with the AdditionalModelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalModelIds

`func (o *CreateAgentV2) SetAdditionalModelIds(v []float32)`

SetAdditionalModelIds sets AdditionalModelIds field to given value.

### HasAdditionalModelIds

`func (o *CreateAgentV2) HasAdditionalModelIds() bool`

HasAdditionalModelIds returns a boolean if a field has been set.

### GetIsWebSearchEnabled

`func (o *CreateAgentV2) GetIsWebSearchEnabled() bool`

GetIsWebSearchEnabled returns the IsWebSearchEnabled field if non-nil, zero value otherwise.

### GetIsWebSearchEnabledOk

`func (o *CreateAgentV2) GetIsWebSearchEnabledOk() (*bool, bool)`

GetIsWebSearchEnabledOk returns a tuple with the IsWebSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebSearchEnabled

`func (o *CreateAgentV2) SetIsWebSearchEnabled(v bool)`

SetIsWebSearchEnabled sets IsWebSearchEnabled field to given value.

### HasIsWebSearchEnabled

`func (o *CreateAgentV2) HasIsWebSearchEnabled() bool`

HasIsWebSearchEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


