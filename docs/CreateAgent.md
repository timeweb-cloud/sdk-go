# CreateAgent

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
**TokenPackageId** | **int32** | ID пакета токенов. Обязателен для совместимости, но не учитывается: агент всегда создаётся с поресурсной оплатой. | 

## Methods

### NewCreateAgent

`func NewCreateAgent(name string, accessType string, modelId int32, settings AgentSettings, tokenPackageId int32, ) *CreateAgent`

NewCreateAgent instantiates a new CreateAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAgentWithDefaults

`func NewCreateAgentWithDefaults() *CreateAgent`

NewCreateAgentWithDefaults instantiates a new CreateAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAgent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAgent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAgent) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAgent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAgent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAgent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAgent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessType

`func (o *CreateAgent) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CreateAgent) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CreateAgent) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetModelId

`func (o *CreateAgent) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CreateAgent) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CreateAgent) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetSettings

`func (o *CreateAgent) GetSettings() AgentSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateAgent) GetSettingsOk() (*AgentSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateAgent) SetSettings(v AgentSettings)`

SetSettings sets Settings field to given value.


### GetProjectId

`func (o *CreateAgent) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateAgent) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateAgent) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateAgent) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetIsWebSearchEnabled

`func (o *CreateAgent) GetIsWebSearchEnabled() bool`

GetIsWebSearchEnabled returns the IsWebSearchEnabled field if non-nil, zero value otherwise.

### GetIsWebSearchEnabledOk

`func (o *CreateAgent) GetIsWebSearchEnabledOk() (*bool, bool)`

GetIsWebSearchEnabledOk returns a tuple with the IsWebSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebSearchEnabled

`func (o *CreateAgent) SetIsWebSearchEnabled(v bool)`

SetIsWebSearchEnabled sets IsWebSearchEnabled field to given value.

### HasIsWebSearchEnabled

`func (o *CreateAgent) HasIsWebSearchEnabled() bool`

HasIsWebSearchEnabled returns a boolean if a field has been set.

### GetTokenPackageId

`func (o *CreateAgent) GetTokenPackageId() int32`

GetTokenPackageId returns the TokenPackageId field if non-nil, zero value otherwise.

### GetTokenPackageIdOk

`func (o *CreateAgent) GetTokenPackageIdOk() (*int32, bool)`

GetTokenPackageIdOk returns a tuple with the TokenPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackageId

`func (o *CreateAgent) SetTokenPackageId(v int32)`

SetTokenPackageId sets TokenPackageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


