# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID агента | 
**Name** | **string** | Название агента | 
**Description** | **NullableString** | Описание агента | 
**ProjectId** | Pointer to **int32** | ID проекта, в котором находится агент | [optional] 
**ModelId** | **int32** | ID основной модели агента | 
**ProviderId** | **int32** | ID провайдера модели | 
**Settings** | [**AgentSettings**](AgentSettings.md) |  | 
**Status** | **string** | Статус агента: - &#x60;active&#x60; — работает; - &#x60;suspended&#x60; — приостановлен владельцем; - &#x60;blocked&#x60; — заблокирован (для агентов с пакетом токенов — закончились токены или подписка); - &#x60;deleted&#x60; — удалён. | 
**AccessType** | **string** | Тип доступа: &#x60;public&#x60; — агент отвечает анонимным посетителям сайта через виджет; &#x60;private&#x60; — только по токену аккаунта. | 
**AccessId** | **string** | Идентификатор доступа агента. Используется для вызова агента из своих приложений через OpenAI-совместимое API (см. https://agent.timeweb.cloud/docs) и для встраивания виджета. | 
**TokenLimit** | **int32** | Устарело. Дневной лимит агента в токенах без привязки к модели, если он задан; &#x60;0&#x60; — такого лимита нет. Лимиты в деньгах, по моделям и за другие периоды здесь не отражаются — используйте лимиты агента (&#x60;/api/v1/cloud-ai/agents/{id}/entity-limits&#x60;). | 
**TotalTokens** | Pointer to **int32** | Всего токенов в пакете. Только для агентов с пакетом токенов (устаревшая схема оплаты). | [optional] 
**UsedTokens** | **NullableInt32** | Использовано токенов пакета. Для агентов с поресурсной оплатой — &#x60;null&#x60;. | 
**RemainingTokens** | Pointer to **int32** | Осталось токенов в пакете. Только для агентов с пакетом токенов. | [optional] 
**TokenPackageId** | Pointer to **int32** | ID пакета токенов. Только для агентов с пакетом токенов; у агентов с поресурсной оплатой поле отсутствует. | [optional] 
**SubscriptionRenewalDate** | Pointer to **time.Time** | Дата продления подписки на пакет токенов. Только для агентов с пакетом токенов. | [optional] 
**AutoBuyAdditionalTokens** | **bool** | Включена ли автоматическая докупка пакетов токенов | 
**Tools** | [**[]AgentToolsInner**](AgentToolsInner.md) | Инструменты, подключённые к агенту | 
**AdditionalModelIds** | **[]int32** | ID дополнительных моделей агента | 
**AdditionalModels** | [**[]AgentAdditionalModel**](AgentAdditionalModel.md) | Дополнительные модели агента с настройками | 
**IsWebSearchEnabled** | **bool** | Включён ли веб-поиск | 
**AiGatewayKeyId** | Pointer to **NullableString** | ID ключа AI Gateway, к которому привязан агент. &#x60;null&#x60; — агента можно привязать к ключу (см. обновление агента); поле отсутствует — привязка недоступна. | [optional] 
**CreatedAt** | **time.Time** | Дата создания агента | 
**KnowledgeBasesIds** | **[]int32** | ID баз знаний, подключённых к агенту | 

## Methods

### NewAgent

`func NewAgent(id int32, name string, description NullableString, modelId int32, providerId int32, settings AgentSettings, status string, accessType string, accessId string, tokenLimit int32, usedTokens NullableInt32, autoBuyAdditionalTokens bool, tools []AgentToolsInner, additionalModelIds []int32, additionalModels []AgentAdditionalModel, isWebSearchEnabled bool, createdAt time.Time, knowledgeBasesIds []int32, ) *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Agent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agent) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agent) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Agent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Agent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Agent) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Agent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Agent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectId

`func (o *Agent) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Agent) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Agent) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Agent) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetModelId

`func (o *Agent) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *Agent) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *Agent) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetProviderId

`func (o *Agent) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Agent) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Agent) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetSettings

`func (o *Agent) GetSettings() AgentSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Agent) GetSettingsOk() (*AgentSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Agent) SetSettings(v AgentSettings)`

SetSettings sets Settings field to given value.


### GetStatus

`func (o *Agent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Agent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Agent) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccessType

`func (o *Agent) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *Agent) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *Agent) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetAccessId

`func (o *Agent) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *Agent) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *Agent) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetTokenLimit

`func (o *Agent) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *Agent) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *Agent) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.


### GetTotalTokens

`func (o *Agent) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *Agent) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *Agent) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *Agent) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsedTokens

`func (o *Agent) GetUsedTokens() int32`

GetUsedTokens returns the UsedTokens field if non-nil, zero value otherwise.

### GetUsedTokensOk

`func (o *Agent) GetUsedTokensOk() (*int32, bool)`

GetUsedTokensOk returns a tuple with the UsedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTokens

`func (o *Agent) SetUsedTokens(v int32)`

SetUsedTokens sets UsedTokens field to given value.


### SetUsedTokensNil

`func (o *Agent) SetUsedTokensNil(b bool)`

 SetUsedTokensNil sets the value for UsedTokens to be an explicit nil

### UnsetUsedTokens
`func (o *Agent) UnsetUsedTokens()`

UnsetUsedTokens ensures that no value is present for UsedTokens, not even an explicit nil
### GetRemainingTokens

`func (o *Agent) GetRemainingTokens() int32`

GetRemainingTokens returns the RemainingTokens field if non-nil, zero value otherwise.

### GetRemainingTokensOk

`func (o *Agent) GetRemainingTokensOk() (*int32, bool)`

GetRemainingTokensOk returns a tuple with the RemainingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTokens

`func (o *Agent) SetRemainingTokens(v int32)`

SetRemainingTokens sets RemainingTokens field to given value.

### HasRemainingTokens

`func (o *Agent) HasRemainingTokens() bool`

HasRemainingTokens returns a boolean if a field has been set.

### GetTokenPackageId

`func (o *Agent) GetTokenPackageId() int32`

GetTokenPackageId returns the TokenPackageId field if non-nil, zero value otherwise.

### GetTokenPackageIdOk

`func (o *Agent) GetTokenPackageIdOk() (*int32, bool)`

GetTokenPackageIdOk returns a tuple with the TokenPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackageId

`func (o *Agent) SetTokenPackageId(v int32)`

SetTokenPackageId sets TokenPackageId field to given value.

### HasTokenPackageId

`func (o *Agent) HasTokenPackageId() bool`

HasTokenPackageId returns a boolean if a field has been set.

### GetSubscriptionRenewalDate

`func (o *Agent) GetSubscriptionRenewalDate() time.Time`

GetSubscriptionRenewalDate returns the SubscriptionRenewalDate field if non-nil, zero value otherwise.

### GetSubscriptionRenewalDateOk

`func (o *Agent) GetSubscriptionRenewalDateOk() (*time.Time, bool)`

GetSubscriptionRenewalDateOk returns a tuple with the SubscriptionRenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRenewalDate

`func (o *Agent) SetSubscriptionRenewalDate(v time.Time)`

SetSubscriptionRenewalDate sets SubscriptionRenewalDate field to given value.

### HasSubscriptionRenewalDate

`func (o *Agent) HasSubscriptionRenewalDate() bool`

HasSubscriptionRenewalDate returns a boolean if a field has been set.

### GetAutoBuyAdditionalTokens

`func (o *Agent) GetAutoBuyAdditionalTokens() bool`

GetAutoBuyAdditionalTokens returns the AutoBuyAdditionalTokens field if non-nil, zero value otherwise.

### GetAutoBuyAdditionalTokensOk

`func (o *Agent) GetAutoBuyAdditionalTokensOk() (*bool, bool)`

GetAutoBuyAdditionalTokensOk returns a tuple with the AutoBuyAdditionalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBuyAdditionalTokens

`func (o *Agent) SetAutoBuyAdditionalTokens(v bool)`

SetAutoBuyAdditionalTokens sets AutoBuyAdditionalTokens field to given value.


### GetTools

`func (o *Agent) GetTools() []AgentToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *Agent) GetToolsOk() (*[]AgentToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *Agent) SetTools(v []AgentToolsInner)`

SetTools sets Tools field to given value.


### GetAdditionalModelIds

`func (o *Agent) GetAdditionalModelIds() []int32`

GetAdditionalModelIds returns the AdditionalModelIds field if non-nil, zero value otherwise.

### GetAdditionalModelIdsOk

`func (o *Agent) GetAdditionalModelIdsOk() (*[]int32, bool)`

GetAdditionalModelIdsOk returns a tuple with the AdditionalModelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalModelIds

`func (o *Agent) SetAdditionalModelIds(v []int32)`

SetAdditionalModelIds sets AdditionalModelIds field to given value.


### GetAdditionalModels

`func (o *Agent) GetAdditionalModels() []AgentAdditionalModel`

GetAdditionalModels returns the AdditionalModels field if non-nil, zero value otherwise.

### GetAdditionalModelsOk

`func (o *Agent) GetAdditionalModelsOk() (*[]AgentAdditionalModel, bool)`

GetAdditionalModelsOk returns a tuple with the AdditionalModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalModels

`func (o *Agent) SetAdditionalModels(v []AgentAdditionalModel)`

SetAdditionalModels sets AdditionalModels field to given value.


### GetIsWebSearchEnabled

`func (o *Agent) GetIsWebSearchEnabled() bool`

GetIsWebSearchEnabled returns the IsWebSearchEnabled field if non-nil, zero value otherwise.

### GetIsWebSearchEnabledOk

`func (o *Agent) GetIsWebSearchEnabledOk() (*bool, bool)`

GetIsWebSearchEnabledOk returns a tuple with the IsWebSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebSearchEnabled

`func (o *Agent) SetIsWebSearchEnabled(v bool)`

SetIsWebSearchEnabled sets IsWebSearchEnabled field to given value.


### GetAiGatewayKeyId

`func (o *Agent) GetAiGatewayKeyId() string`

GetAiGatewayKeyId returns the AiGatewayKeyId field if non-nil, zero value otherwise.

### GetAiGatewayKeyIdOk

`func (o *Agent) GetAiGatewayKeyIdOk() (*string, bool)`

GetAiGatewayKeyIdOk returns a tuple with the AiGatewayKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiGatewayKeyId

`func (o *Agent) SetAiGatewayKeyId(v string)`

SetAiGatewayKeyId sets AiGatewayKeyId field to given value.

### HasAiGatewayKeyId

`func (o *Agent) HasAiGatewayKeyId() bool`

HasAiGatewayKeyId returns a boolean if a field has been set.

### SetAiGatewayKeyIdNil

`func (o *Agent) SetAiGatewayKeyIdNil(b bool)`

 SetAiGatewayKeyIdNil sets the value for AiGatewayKeyId to be an explicit nil

### UnsetAiGatewayKeyId
`func (o *Agent) UnsetAiGatewayKeyId()`

UnsetAiGatewayKeyId ensures that no value is present for AiGatewayKeyId, not even an explicit nil
### GetCreatedAt

`func (o *Agent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Agent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Agent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetKnowledgeBasesIds

`func (o *Agent) GetKnowledgeBasesIds() []int32`

GetKnowledgeBasesIds returns the KnowledgeBasesIds field if non-nil, zero value otherwise.

### GetKnowledgeBasesIdsOk

`func (o *Agent) GetKnowledgeBasesIdsOk() (*[]int32, bool)`

GetKnowledgeBasesIdsOk returns a tuple with the KnowledgeBasesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBasesIds

`func (o *Agent) SetKnowledgeBasesIds(v []int32)`

SetKnowledgeBasesIds sets KnowledgeBasesIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


