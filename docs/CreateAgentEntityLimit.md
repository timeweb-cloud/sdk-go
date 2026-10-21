# CreateAgentEntityLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitType** | **string** | Период лимита | 
**ModelId** | Pointer to **NullableInt32** | ID модели, на запросы к которой действует лимит. Не передан или &#x60;null&#x60; — лимит на все модели агента. | [optional] 
**Limit** | Pointer to **int32** | Лимит в токенах за период. Если не передан — &#x60;0&#x60;, тогда обязателен &#x60;money_limit&#x60;. Не действует, если задан &#x60;money_limit&#x60;. &#x60;0&#x60; при отсутствии &#x60;money_limit&#x60; — запрещены все запросы. | [optional] 
**MoneyLimit** | Pointer to **NullableFloat32** | Лимит расходов за период в валюте аккаунта, до 8 знаков после запятой. Приоритетнее &#x60;limit&#x60;. &#x60;0&#x60; — запрещены все запросы. | [optional] 

## Methods

### NewCreateAgentEntityLimit

`func NewCreateAgentEntityLimit(limitType string, ) *CreateAgentEntityLimit`

NewCreateAgentEntityLimit instantiates a new CreateAgentEntityLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAgentEntityLimitWithDefaults

`func NewCreateAgentEntityLimitWithDefaults() *CreateAgentEntityLimit`

NewCreateAgentEntityLimitWithDefaults instantiates a new CreateAgentEntityLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitType

`func (o *CreateAgentEntityLimit) GetLimitType() string`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *CreateAgentEntityLimit) GetLimitTypeOk() (*string, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *CreateAgentEntityLimit) SetLimitType(v string)`

SetLimitType sets LimitType field to given value.


### GetModelId

`func (o *CreateAgentEntityLimit) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CreateAgentEntityLimit) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CreateAgentEntityLimit) SetModelId(v int32)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *CreateAgentEntityLimit) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### SetModelIdNil

`func (o *CreateAgentEntityLimit) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *CreateAgentEntityLimit) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetLimit

`func (o *CreateAgentEntityLimit) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CreateAgentEntityLimit) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CreateAgentEntityLimit) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CreateAgentEntityLimit) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMoneyLimit

`func (o *CreateAgentEntityLimit) GetMoneyLimit() float32`

GetMoneyLimit returns the MoneyLimit field if non-nil, zero value otherwise.

### GetMoneyLimitOk

`func (o *CreateAgentEntityLimit) GetMoneyLimitOk() (*float32, bool)`

GetMoneyLimitOk returns a tuple with the MoneyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyLimit

`func (o *CreateAgentEntityLimit) SetMoneyLimit(v float32)`

SetMoneyLimit sets MoneyLimit field to given value.

### HasMoneyLimit

`func (o *CreateAgentEntityLimit) HasMoneyLimit() bool`

HasMoneyLimit returns a boolean if a field has been set.

### SetMoneyLimitNil

`func (o *CreateAgentEntityLimit) SetMoneyLimitNil(b bool)`

 SetMoneyLimitNil sets the value for MoneyLimit to be an explicit nil

### UnsetMoneyLimit
`func (o *CreateAgentEntityLimit) UnsetMoneyLimit()`

UnsetMoneyLimit ensures that no value is present for MoneyLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


