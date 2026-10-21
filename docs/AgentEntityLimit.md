# AgentEntityLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID лимита | 
**AiKeyId** | **string** | ID служебного ключа агента, на котором учитывается лимит | 
**EntityType** | **string** | Тип сущности — всегда &#x60;agent&#x60; | 
**EntityId** | **string** | ID агента | 
**LimitType** | **string** | Период лимита | 
**ModelId** | **NullableInt32** | ID модели, на запросы к которой действует лимит. &#x60;null&#x60; — лимит на все модели агента. | 
**Limit** | **int32** | Лимит в токенах за период. Не действует, если задан &#x60;money_limit&#x60;. &#x60;0&#x60; — запрещены все запросы. | 
**MoneyLimit** | **NullableFloat32** | Лимит расходов за период в валюте аккаунта; &#x60;null&#x60; — не задан, действует &#x60;limit&#x60; | 
**Usage** | **int32** | Израсходовано токенов в текущем периоде | 
**MoneyUsage** | **float32** | Израсходовано денег в текущем периоде (полная стоимость запросов до бонусов) | 
**CreatedAt** | **time.Time** | Дата создания лимита | 
**ExpireAt** | **time.Time** | Окончание текущего периода — после него расход обнуляется | 

## Methods

### NewAgentEntityLimit

`func NewAgentEntityLimit(id string, aiKeyId string, entityType string, entityId string, limitType string, modelId NullableInt32, limit int32, moneyLimit NullableFloat32, usage int32, moneyUsage float32, createdAt time.Time, expireAt time.Time, ) *AgentEntityLimit`

NewAgentEntityLimit instantiates a new AgentEntityLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentEntityLimitWithDefaults

`func NewAgentEntityLimitWithDefaults() *AgentEntityLimit`

NewAgentEntityLimitWithDefaults instantiates a new AgentEntityLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentEntityLimit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentEntityLimit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentEntityLimit) SetId(v string)`

SetId sets Id field to given value.


### GetAiKeyId

`func (o *AgentEntityLimit) GetAiKeyId() string`

GetAiKeyId returns the AiKeyId field if non-nil, zero value otherwise.

### GetAiKeyIdOk

`func (o *AgentEntityLimit) GetAiKeyIdOk() (*string, bool)`

GetAiKeyIdOk returns a tuple with the AiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiKeyId

`func (o *AgentEntityLimit) SetAiKeyId(v string)`

SetAiKeyId sets AiKeyId field to given value.


### GetEntityType

`func (o *AgentEntityLimit) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AgentEntityLimit) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AgentEntityLimit) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetEntityId

`func (o *AgentEntityLimit) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AgentEntityLimit) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AgentEntityLimit) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetLimitType

`func (o *AgentEntityLimit) GetLimitType() string`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *AgentEntityLimit) GetLimitTypeOk() (*string, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *AgentEntityLimit) SetLimitType(v string)`

SetLimitType sets LimitType field to given value.


### GetModelId

`func (o *AgentEntityLimit) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AgentEntityLimit) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AgentEntityLimit) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### SetModelIdNil

`func (o *AgentEntityLimit) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *AgentEntityLimit) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetLimit

`func (o *AgentEntityLimit) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AgentEntityLimit) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AgentEntityLimit) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetMoneyLimit

`func (o *AgentEntityLimit) GetMoneyLimit() float32`

GetMoneyLimit returns the MoneyLimit field if non-nil, zero value otherwise.

### GetMoneyLimitOk

`func (o *AgentEntityLimit) GetMoneyLimitOk() (*float32, bool)`

GetMoneyLimitOk returns a tuple with the MoneyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyLimit

`func (o *AgentEntityLimit) SetMoneyLimit(v float32)`

SetMoneyLimit sets MoneyLimit field to given value.


### SetMoneyLimitNil

`func (o *AgentEntityLimit) SetMoneyLimitNil(b bool)`

 SetMoneyLimitNil sets the value for MoneyLimit to be an explicit nil

### UnsetMoneyLimit
`func (o *AgentEntityLimit) UnsetMoneyLimit()`

UnsetMoneyLimit ensures that no value is present for MoneyLimit, not even an explicit nil
### GetUsage

`func (o *AgentEntityLimit) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AgentEntityLimit) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AgentEntityLimit) SetUsage(v int32)`

SetUsage sets Usage field to given value.


### GetMoneyUsage

`func (o *AgentEntityLimit) GetMoneyUsage() float32`

GetMoneyUsage returns the MoneyUsage field if non-nil, zero value otherwise.

### GetMoneyUsageOk

`func (o *AgentEntityLimit) GetMoneyUsageOk() (*float32, bool)`

GetMoneyUsageOk returns a tuple with the MoneyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyUsage

`func (o *AgentEntityLimit) SetMoneyUsage(v float32)`

SetMoneyUsage sets MoneyUsage field to given value.


### GetCreatedAt

`func (o *AgentEntityLimit) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentEntityLimit) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentEntityLimit) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpireAt

`func (o *AgentEntityLimit) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *AgentEntityLimit) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *AgentEntityLimit) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


