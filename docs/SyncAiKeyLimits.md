# SyncAiKeyLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | [**[]SyncAiKeyLimitsLimitsInner**](SyncAiKeyLimitsLimitsInner.md) | Лимиты по моделям. &#x60;model_id&#x60; в списке не должны повторяться. Пустой список удаляет все лимиты периода | 
**LimitType** | **string** | Период лимита: &#x60;hour&#x60; — час, &#x60;day&#x60; — сутки, &#x60;week&#x60; — 7 дней, &#x60;month&#x60; — 30 дней | 
**EntityType** | Pointer to **string** | Тип сущности. По умолчанию — тип ключа (лимиты на сам ключ) | [optional] 
**EntityId** | Pointer to **NullableString** | ID сущности. Для лимитов на сам ключ не нужен | [optional] 

## Methods

### NewSyncAiKeyLimits

`func NewSyncAiKeyLimits(limits []SyncAiKeyLimitsLimitsInner, limitType string, ) *SyncAiKeyLimits`

NewSyncAiKeyLimits instantiates a new SyncAiKeyLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncAiKeyLimitsWithDefaults

`func NewSyncAiKeyLimitsWithDefaults() *SyncAiKeyLimits`

NewSyncAiKeyLimitsWithDefaults instantiates a new SyncAiKeyLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *SyncAiKeyLimits) GetLimits() []SyncAiKeyLimitsLimitsInner`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *SyncAiKeyLimits) GetLimitsOk() (*[]SyncAiKeyLimitsLimitsInner, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *SyncAiKeyLimits) SetLimits(v []SyncAiKeyLimitsLimitsInner)`

SetLimits sets Limits field to given value.


### GetLimitType

`func (o *SyncAiKeyLimits) GetLimitType() string`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *SyncAiKeyLimits) GetLimitTypeOk() (*string, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *SyncAiKeyLimits) SetLimitType(v string)`

SetLimitType sets LimitType field to given value.


### GetEntityType

`func (o *SyncAiKeyLimits) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *SyncAiKeyLimits) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *SyncAiKeyLimits) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *SyncAiKeyLimits) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEntityId

`func (o *SyncAiKeyLimits) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SyncAiKeyLimits) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SyncAiKeyLimits) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SyncAiKeyLimits) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *SyncAiKeyLimits) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *SyncAiKeyLimits) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


