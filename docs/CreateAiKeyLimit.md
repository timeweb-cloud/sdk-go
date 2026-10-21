# CreateAiKeyLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | Тип сущности, на которую действует лимит. Для лимита на сам ключ передайте его тип — поле &#x60;entity_type&#x60; ключа (&#x60;aikey&#x60; или &#x60;k8s&#x60;) | 
**EntityId** | Pointer to **NullableString** | ID сущности. Для лимита на сам ключ не нужен: подставляется ID ключа | [optional] 
**LimitType** | **string** | Период лимита: &#x60;hour&#x60; — час, &#x60;day&#x60; — сутки, &#x60;week&#x60; — 7 дней, &#x60;month&#x60; — 30 дней | 
**ModelId** | Pointer to **NullableInt32** | ID модели из каталога (&#x60;GET /api/v3/cloud-ai/models&#x60;), на которую действует лимит. Не передан — лимит на все модели | [optional] 
**Limit** | Pointer to **int32** | Лимит в токенах за период. Нужен &#x60;limit&#x60; или &#x60;money_limit&#x60;; если передан только &#x60;money_limit&#x60;, &#x60;limit&#x60; будет &#x60;0&#x60;. &#x60;0&#x60; без &#x60;money_limit&#x60; — запросы запрещены | [optional] 
**MoneyLimit** | Pointer to **NullableFloat32** | Лимит в деньгах за период, в валюте аккаунта, не больше 8 знаков после точки. Приоритетнее &#x60;limit&#x60;. &#x60;0&#x60; — запросы запрещены | [optional] 

## Methods

### NewCreateAiKeyLimit

`func NewCreateAiKeyLimit(entityType string, limitType string, ) *CreateAiKeyLimit`

NewCreateAiKeyLimit instantiates a new CreateAiKeyLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAiKeyLimitWithDefaults

`func NewCreateAiKeyLimitWithDefaults() *CreateAiKeyLimit`

NewCreateAiKeyLimitWithDefaults instantiates a new CreateAiKeyLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *CreateAiKeyLimit) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CreateAiKeyLimit) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CreateAiKeyLimit) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetEntityId

`func (o *CreateAiKeyLimit) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *CreateAiKeyLimit) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *CreateAiKeyLimit) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *CreateAiKeyLimit) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *CreateAiKeyLimit) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *CreateAiKeyLimit) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetLimitType

`func (o *CreateAiKeyLimit) GetLimitType() string`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *CreateAiKeyLimit) GetLimitTypeOk() (*string, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *CreateAiKeyLimit) SetLimitType(v string)`

SetLimitType sets LimitType field to given value.


### GetModelId

`func (o *CreateAiKeyLimit) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CreateAiKeyLimit) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CreateAiKeyLimit) SetModelId(v int32)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *CreateAiKeyLimit) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### SetModelIdNil

`func (o *CreateAiKeyLimit) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *CreateAiKeyLimit) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetLimit

`func (o *CreateAiKeyLimit) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CreateAiKeyLimit) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CreateAiKeyLimit) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CreateAiKeyLimit) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMoneyLimit

`func (o *CreateAiKeyLimit) GetMoneyLimit() float32`

GetMoneyLimit returns the MoneyLimit field if non-nil, zero value otherwise.

### GetMoneyLimitOk

`func (o *CreateAiKeyLimit) GetMoneyLimitOk() (*float32, bool)`

GetMoneyLimitOk returns a tuple with the MoneyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyLimit

`func (o *CreateAiKeyLimit) SetMoneyLimit(v float32)`

SetMoneyLimit sets MoneyLimit field to given value.

### HasMoneyLimit

`func (o *CreateAiKeyLimit) HasMoneyLimit() bool`

HasMoneyLimit returns a boolean if a field has been set.

### SetMoneyLimitNil

`func (o *CreateAiKeyLimit) SetMoneyLimitNil(b bool)`

 SetMoneyLimitNil sets the value for MoneyLimit to be an explicit nil

### UnsetMoneyLimit
`func (o *CreateAiKeyLimit) UnsetMoneyLimit()`

UnsetMoneyLimit ensures that no value is present for MoneyLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


