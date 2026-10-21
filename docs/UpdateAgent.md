# UpdateAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название агента | [optional] 
**Description** | Pointer to **string** | Описание агента | [optional] 
**AccessType** | Pointer to **string** | Тип доступа | [optional] 
**Status** | Pointer to **string** | Статус агента: &#x60;suspended&#x60; — приостановить, &#x60;active&#x60; — возобновить. Статус заблокированного или удалённого агента изменить нельзя. | [optional] 
**Settings** | Pointer to [**AgentSettings**](AgentSettings.md) |  | [optional] 
**ProjectId** | Pointer to **int32** | Не используется: перенести агента в другой проект можно методами проектов. | [optional] 
**AdditionalModelIds** | Pointer to **[]int32** | ID дополнительных моделей агента (например, для генерации изображений). Список заменяет текущий целиком. | [optional] 
**AdditionalModels** | Pointer to [**[]AgentAdditionalModel**](AgentAdditionalModel.md) | Дополнительные модели вместе с настройками (для моделей генерации изображений — &#x60;width&#x60;/&#x60;height&#x60;). Если переданы и &#x60;additional_models&#x60;, и &#x60;additional_model_ids&#x60;, используется &#x60;additional_models&#x60;. | [optional] 
**IsWebSearchEnabled** | Pointer to **bool** | Включить или выключить веб-поиск | [optional] 
**AiGatewayKeyId** | Pointer to **string** | ID ключа AI Gateway, к которому нужно привязать агента. Привязка доступна, только если в ответе агента &#x60;ai_gateway_key_id&#x60; равен &#x60;null&#x60; (агент без баз знаний, инструментов и интеграций, созданный до 9 октября 2026 года), и делается один раз. При первой привязке на аккаунт однократно начисляется временная скидка 10% на запросы через ключи AI Gateway. | [optional] 
**TokenPackageId** | Pointer to **int32** | Новый пакет токенов. Только для агентов с пакетом токенов (устаревшая схема оплаты): пакет должен относиться к модели агента и быть не меньше текущего, его стоимость сразу списывается с баланса. | [optional] 

## Methods

### NewUpdateAgent

`func NewUpdateAgent() *UpdateAgent`

NewUpdateAgent instantiates a new UpdateAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAgentWithDefaults

`func NewUpdateAgentWithDefaults() *UpdateAgent`

NewUpdateAgentWithDefaults instantiates a new UpdateAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAgent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAgent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAgent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAgent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAgent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAgent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAgent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAgent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessType

`func (o *UpdateAgent) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UpdateAgent) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UpdateAgent) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *UpdateAgent) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateAgent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAgent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAgent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateAgent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateAgent) GetSettings() AgentSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateAgent) GetSettingsOk() (*AgentSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateAgent) SetSettings(v AgentSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateAgent) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetProjectId

`func (o *UpdateAgent) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateAgent) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateAgent) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *UpdateAgent) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetAdditionalModelIds

`func (o *UpdateAgent) GetAdditionalModelIds() []int32`

GetAdditionalModelIds returns the AdditionalModelIds field if non-nil, zero value otherwise.

### GetAdditionalModelIdsOk

`func (o *UpdateAgent) GetAdditionalModelIdsOk() (*[]int32, bool)`

GetAdditionalModelIdsOk returns a tuple with the AdditionalModelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalModelIds

`func (o *UpdateAgent) SetAdditionalModelIds(v []int32)`

SetAdditionalModelIds sets AdditionalModelIds field to given value.

### HasAdditionalModelIds

`func (o *UpdateAgent) HasAdditionalModelIds() bool`

HasAdditionalModelIds returns a boolean if a field has been set.

### GetAdditionalModels

`func (o *UpdateAgent) GetAdditionalModels() []AgentAdditionalModel`

GetAdditionalModels returns the AdditionalModels field if non-nil, zero value otherwise.

### GetAdditionalModelsOk

`func (o *UpdateAgent) GetAdditionalModelsOk() (*[]AgentAdditionalModel, bool)`

GetAdditionalModelsOk returns a tuple with the AdditionalModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalModels

`func (o *UpdateAgent) SetAdditionalModels(v []AgentAdditionalModel)`

SetAdditionalModels sets AdditionalModels field to given value.

### HasAdditionalModels

`func (o *UpdateAgent) HasAdditionalModels() bool`

HasAdditionalModels returns a boolean if a field has been set.

### GetIsWebSearchEnabled

`func (o *UpdateAgent) GetIsWebSearchEnabled() bool`

GetIsWebSearchEnabled returns the IsWebSearchEnabled field if non-nil, zero value otherwise.

### GetIsWebSearchEnabledOk

`func (o *UpdateAgent) GetIsWebSearchEnabledOk() (*bool, bool)`

GetIsWebSearchEnabledOk returns a tuple with the IsWebSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebSearchEnabled

`func (o *UpdateAgent) SetIsWebSearchEnabled(v bool)`

SetIsWebSearchEnabled sets IsWebSearchEnabled field to given value.

### HasIsWebSearchEnabled

`func (o *UpdateAgent) HasIsWebSearchEnabled() bool`

HasIsWebSearchEnabled returns a boolean if a field has been set.

### GetAiGatewayKeyId

`func (o *UpdateAgent) GetAiGatewayKeyId() string`

GetAiGatewayKeyId returns the AiGatewayKeyId field if non-nil, zero value otherwise.

### GetAiGatewayKeyIdOk

`func (o *UpdateAgent) GetAiGatewayKeyIdOk() (*string, bool)`

GetAiGatewayKeyIdOk returns a tuple with the AiGatewayKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiGatewayKeyId

`func (o *UpdateAgent) SetAiGatewayKeyId(v string)`

SetAiGatewayKeyId sets AiGatewayKeyId field to given value.

### HasAiGatewayKeyId

`func (o *UpdateAgent) HasAiGatewayKeyId() bool`

HasAiGatewayKeyId returns a boolean if a field has been set.

### GetTokenPackageId

`func (o *UpdateAgent) GetTokenPackageId() int32`

GetTokenPackageId returns the TokenPackageId field if non-nil, zero value otherwise.

### GetTokenPackageIdOk

`func (o *UpdateAgent) GetTokenPackageIdOk() (*int32, bool)`

GetTokenPackageIdOk returns a tuple with the TokenPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackageId

`func (o *UpdateAgent) SetTokenPackageId(v int32)`

SetTokenPackageId sets TokenPackageId field to given value.

### HasTokenPackageId

`func (o *UpdateAgent) HasTokenPackageId() bool`

HasTokenPackageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


