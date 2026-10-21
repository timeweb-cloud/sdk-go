# AiKeyLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID лимита (UUID) | 
**AiKeyId** | **string** | ID ключа, к которому относится лимит | 
**EntityType** | **string** | Тип сущности, на которую действует лимит. Для лимита на сам ключ — тип ключа (&#x60;aikey&#x60; или &#x60;k8s&#x60;) | 
**EntityId** | **string** | ID сущности, на которую действует лимит. Для лимита на сам ключ — ID ключа | 
**Limit** | **int32** | Лимит в токенах за период. Не действует, если задан &#x60;money_limit&#x60;. &#x60;0&#x60; — запросы запрещены | 
**MoneyLimit** | **NullableFloat32** | Лимит в деньгах за период, в валюте аккаунта. Приоритетнее &#x60;limit&#x60;. &#x60;null&#x60; — не задан, действует &#x60;limit&#x60;; &#x60;0&#x60; — запросы запрещены | 
**LimitType** | **string** | Период лимита: &#x60;hour&#x60; — час, &#x60;day&#x60; — сутки, &#x60;week&#x60; — 7 дней, &#x60;month&#x60; — 30 дней | 
**ModelId** | **NullableInt32** | ID модели, на которую действует лимит. &#x60;null&#x60; — лимит на все модели | 
**Usage** | **int32** | Израсходовано токенов за текущий период | 
**MoneyUsage** | **float32** | Израсходовано денег за текущий период — полная стоимость запросов до списания бонусов, в валюте аккаунта | 
**CreatedAt** | **time.Time** | Дата и время создания лимита | 
**ExpireAt** | **time.Time** | Дата и время окончания текущего периода: после него расход обнуляется и начинается новый период | 

## Methods

### NewAiKeyLimit

`func NewAiKeyLimit(id string, aiKeyId string, entityType string, entityId string, limit int32, moneyLimit NullableFloat32, limitType string, modelId NullableInt32, usage int32, moneyUsage float32, createdAt time.Time, expireAt time.Time, ) *AiKeyLimit`

NewAiKeyLimit instantiates a new AiKeyLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiKeyLimitWithDefaults

`func NewAiKeyLimitWithDefaults() *AiKeyLimit`

NewAiKeyLimitWithDefaults instantiates a new AiKeyLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiKeyLimit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiKeyLimit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiKeyLimit) SetId(v string)`

SetId sets Id field to given value.


### GetAiKeyId

`func (o *AiKeyLimit) GetAiKeyId() string`

GetAiKeyId returns the AiKeyId field if non-nil, zero value otherwise.

### GetAiKeyIdOk

`func (o *AiKeyLimit) GetAiKeyIdOk() (*string, bool)`

GetAiKeyIdOk returns a tuple with the AiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiKeyId

`func (o *AiKeyLimit) SetAiKeyId(v string)`

SetAiKeyId sets AiKeyId field to given value.


### GetEntityType

`func (o *AiKeyLimit) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AiKeyLimit) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AiKeyLimit) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetEntityId

`func (o *AiKeyLimit) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiKeyLimit) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiKeyLimit) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetLimit

`func (o *AiKeyLimit) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AiKeyLimit) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AiKeyLimit) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetMoneyLimit

`func (o *AiKeyLimit) GetMoneyLimit() float32`

GetMoneyLimit returns the MoneyLimit field if non-nil, zero value otherwise.

### GetMoneyLimitOk

`func (o *AiKeyLimit) GetMoneyLimitOk() (*float32, bool)`

GetMoneyLimitOk returns a tuple with the MoneyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyLimit

`func (o *AiKeyLimit) SetMoneyLimit(v float32)`

SetMoneyLimit sets MoneyLimit field to given value.


### SetMoneyLimitNil

`func (o *AiKeyLimit) SetMoneyLimitNil(b bool)`

 SetMoneyLimitNil sets the value for MoneyLimit to be an explicit nil

### UnsetMoneyLimit
`func (o *AiKeyLimit) UnsetMoneyLimit()`

UnsetMoneyLimit ensures that no value is present for MoneyLimit, not even an explicit nil
### GetLimitType

`func (o *AiKeyLimit) GetLimitType() string`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *AiKeyLimit) GetLimitTypeOk() (*string, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *AiKeyLimit) SetLimitType(v string)`

SetLimitType sets LimitType field to given value.


### GetModelId

`func (o *AiKeyLimit) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AiKeyLimit) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AiKeyLimit) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### SetModelIdNil

`func (o *AiKeyLimit) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *AiKeyLimit) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetUsage

`func (o *AiKeyLimit) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiKeyLimit) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiKeyLimit) SetUsage(v int32)`

SetUsage sets Usage field to given value.


### GetMoneyUsage

`func (o *AiKeyLimit) GetMoneyUsage() float32`

GetMoneyUsage returns the MoneyUsage field if non-nil, zero value otherwise.

### GetMoneyUsageOk

`func (o *AiKeyLimit) GetMoneyUsageOk() (*float32, bool)`

GetMoneyUsageOk returns a tuple with the MoneyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyUsage

`func (o *AiKeyLimit) SetMoneyUsage(v float32)`

SetMoneyUsage sets MoneyUsage field to given value.


### GetCreatedAt

`func (o *AiKeyLimit) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiKeyLimit) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiKeyLimit) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpireAt

`func (o *AiKeyLimit) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *AiKeyLimit) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *AiKeyLimit) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


