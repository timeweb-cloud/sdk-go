# CreateAgentV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название агента | 
**Description** | Pointer to **string** | Описание агента | [optional] 
**AccessType** | **string** | Тип доступа: &#x60;public&#x60; — агент отвечает анонимным посетителям сайта через виджет; &#x60;private&#x60; — только по токену аккаунта. | 
**ModelId** | **int32** | ID основной модели агента | 
**Settings** | [**AgentSettings**](AgentSettings.md) |  | 
**ProjectId** | Pointer to **int32** | ID проекта, в который нужно добавить агента. По умолчанию — проект по умолчанию. | [optional] 
**IsWebSearchEnabled** | Pointer to **bool** | Включить веб-поиск | [optional] [default to false]
**MoneyLimit** | Pointer to **float32** | Дневной лимит расходов агента в валюте аккаунта. &#x60;0&#x60; или не передан — без лимита. Создаётся как лимит агента с периодом &#x60;day&#x60;; изменить его можно через лимиты агента. | [optional] 
**TokenLimit** | Pointer to **int32** | Устарело: дневной лимит агента в токенах. &#x60;0&#x60; или не передан — без лимита. Не учитывается, если передан &#x60;money_limit&#x60;. | [optional] 
**TokenPackageId** | Pointer to **int32** | Устарело и не учитывается: агент всегда создаётся с поресурсной оплатой. | [optional] 
**AdditionalModelIds** | Pointer to **[]int32** | ID дополнительных моделей агента (например, для генерации изображений). Список заменяет текущий целиком. | [optional] 
**AdditionalModels** | Pointer to [**[]AgentAdditionalModel**](AgentAdditionalModel.md) | Дополнительные модели вместе с настройками (для моделей генерации изображений — &#x60;width&#x60;/&#x60;height&#x60;). Если переданы и &#x60;additional_models&#x60;, и &#x60;additional_model_ids&#x60;, используется &#x60;additional_models&#x60;. | [optional] 
**SkillIds** | Pointer to **[]int32** | ID навыков, которые нужно подключить к агенту сразу после создания (текущие версии навыков). Неизвестные ID пропускаются; ошибка подключения навыков не отменяет создание агента. | [optional] 

## Methods

### NewCreateAgentV2

`func NewCreateAgentV2(name string, accessType string, modelId int32, settings AgentSettings, ) *CreateAgentV2`

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

`func (o *CreateAgentV2) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CreateAgentV2) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CreateAgentV2) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


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

`func (o *CreateAgentV2) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateAgentV2) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateAgentV2) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateAgentV2) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

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

### GetMoneyLimit

`func (o *CreateAgentV2) GetMoneyLimit() float32`

GetMoneyLimit returns the MoneyLimit field if non-nil, zero value otherwise.

### GetMoneyLimitOk

`func (o *CreateAgentV2) GetMoneyLimitOk() (*float32, bool)`

GetMoneyLimitOk returns a tuple with the MoneyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyLimit

`func (o *CreateAgentV2) SetMoneyLimit(v float32)`

SetMoneyLimit sets MoneyLimit field to given value.

### HasMoneyLimit

`func (o *CreateAgentV2) HasMoneyLimit() bool`

HasMoneyLimit returns a boolean if a field has been set.

### GetTokenLimit

`func (o *CreateAgentV2) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *CreateAgentV2) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *CreateAgentV2) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.

### HasTokenLimit

`func (o *CreateAgentV2) HasTokenLimit() bool`

HasTokenLimit returns a boolean if a field has been set.

### GetTokenPackageId

`func (o *CreateAgentV2) GetTokenPackageId() int32`

GetTokenPackageId returns the TokenPackageId field if non-nil, zero value otherwise.

### GetTokenPackageIdOk

`func (o *CreateAgentV2) GetTokenPackageIdOk() (*int32, bool)`

GetTokenPackageIdOk returns a tuple with the TokenPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackageId

`func (o *CreateAgentV2) SetTokenPackageId(v int32)`

SetTokenPackageId sets TokenPackageId field to given value.

### HasTokenPackageId

`func (o *CreateAgentV2) HasTokenPackageId() bool`

HasTokenPackageId returns a boolean if a field has been set.

### GetAdditionalModelIds

`func (o *CreateAgentV2) GetAdditionalModelIds() []int32`

GetAdditionalModelIds returns the AdditionalModelIds field if non-nil, zero value otherwise.

### GetAdditionalModelIdsOk

`func (o *CreateAgentV2) GetAdditionalModelIdsOk() (*[]int32, bool)`

GetAdditionalModelIdsOk returns a tuple with the AdditionalModelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalModelIds

`func (o *CreateAgentV2) SetAdditionalModelIds(v []int32)`

SetAdditionalModelIds sets AdditionalModelIds field to given value.

### HasAdditionalModelIds

`func (o *CreateAgentV2) HasAdditionalModelIds() bool`

HasAdditionalModelIds returns a boolean if a field has been set.

### GetAdditionalModels

`func (o *CreateAgentV2) GetAdditionalModels() []AgentAdditionalModel`

GetAdditionalModels returns the AdditionalModels field if non-nil, zero value otherwise.

### GetAdditionalModelsOk

`func (o *CreateAgentV2) GetAdditionalModelsOk() (*[]AgentAdditionalModel, bool)`

GetAdditionalModelsOk returns a tuple with the AdditionalModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalModels

`func (o *CreateAgentV2) SetAdditionalModels(v []AgentAdditionalModel)`

SetAdditionalModels sets AdditionalModels field to given value.

### HasAdditionalModels

`func (o *CreateAgentV2) HasAdditionalModels() bool`

HasAdditionalModels returns a boolean if a field has been set.

### GetSkillIds

`func (o *CreateAgentV2) GetSkillIds() []int32`

GetSkillIds returns the SkillIds field if non-nil, zero value otherwise.

### GetSkillIdsOk

`func (o *CreateAgentV2) GetSkillIdsOk() (*[]int32, bool)`

GetSkillIdsOk returns a tuple with the SkillIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillIds

`func (o *CreateAgentV2) SetSkillIds(v []int32)`

SetSkillIds sets SkillIds field to given value.

### HasSkillIds

`func (o *CreateAgentV2) HasSkillIds() bool`

HasSkillIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


