# SyncAiKeyLimitsLimitsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelId** | **NullableInt32** | ID модели из каталога. &#x60;null&#x60; — лимит на все модели | 
**Limit** | Pointer to **int32** | Лимит в токенах за период. Не передан — у существующего лимита не меняется, новый получает &#x60;0&#x60;. Нужен &#x60;limit&#x60; или &#x60;money_limit&#x60; | [optional] 
**MoneyLimit** | Pointer to **NullableFloat32** | Лимит в деньгах за период, в валюте аккаунта. Приоритетнее &#x60;limit&#x60;. Не передан — у существующего лимита не меняется; &#x60;null&#x60; — снять денежный лимит (только вместе с &#x60;limit&#x60;); &#x60;0&#x60; — запросы запрещены | [optional] 

## Methods

### NewSyncAiKeyLimitsLimitsInner

`func NewSyncAiKeyLimitsLimitsInner(modelId NullableInt32, ) *SyncAiKeyLimitsLimitsInner`

NewSyncAiKeyLimitsLimitsInner instantiates a new SyncAiKeyLimitsLimitsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncAiKeyLimitsLimitsInnerWithDefaults

`func NewSyncAiKeyLimitsLimitsInnerWithDefaults() *SyncAiKeyLimitsLimitsInner`

NewSyncAiKeyLimitsLimitsInnerWithDefaults instantiates a new SyncAiKeyLimitsLimitsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelId

`func (o *SyncAiKeyLimitsLimitsInner) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *SyncAiKeyLimitsLimitsInner) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *SyncAiKeyLimitsLimitsInner) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### SetModelIdNil

`func (o *SyncAiKeyLimitsLimitsInner) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *SyncAiKeyLimitsLimitsInner) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetLimit

`func (o *SyncAiKeyLimitsLimitsInner) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SyncAiKeyLimitsLimitsInner) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SyncAiKeyLimitsLimitsInner) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SyncAiKeyLimitsLimitsInner) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMoneyLimit

`func (o *SyncAiKeyLimitsLimitsInner) GetMoneyLimit() float32`

GetMoneyLimit returns the MoneyLimit field if non-nil, zero value otherwise.

### GetMoneyLimitOk

`func (o *SyncAiKeyLimitsLimitsInner) GetMoneyLimitOk() (*float32, bool)`

GetMoneyLimitOk returns a tuple with the MoneyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyLimit

`func (o *SyncAiKeyLimitsLimitsInner) SetMoneyLimit(v float32)`

SetMoneyLimit sets MoneyLimit field to given value.

### HasMoneyLimit

`func (o *SyncAiKeyLimitsLimitsInner) HasMoneyLimit() bool`

HasMoneyLimit returns a boolean if a field has been set.

### SetMoneyLimitNil

`func (o *SyncAiKeyLimitsLimitsInner) SetMoneyLimitNil(b bool)`

 SetMoneyLimitNil sets the value for MoneyLimit to be an explicit nil

### UnsetMoneyLimit
`func (o *SyncAiKeyLimitsLimitsInner) UnsetMoneyLimit()`

UnsetMoneyLimit ensures that no value is present for MoneyLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


