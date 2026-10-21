# SyncAgentEntityLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitType** | **string** | Период лимита | 
**Limits** | [**[]SyncAgentEntityLimitsLimitsInner**](SyncAgentEntityLimitsLimitsInner.md) | Лимиты за период — по одному на модель (&#x60;model_id&#x60; не должен повторяться). Лимиты этого периода, которых нет в списке, удаляются; пустой массив удаляет все лимиты периода. | 

## Methods

### NewSyncAgentEntityLimits

`func NewSyncAgentEntityLimits(limitType string, limits []SyncAgentEntityLimitsLimitsInner, ) *SyncAgentEntityLimits`

NewSyncAgentEntityLimits instantiates a new SyncAgentEntityLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncAgentEntityLimitsWithDefaults

`func NewSyncAgentEntityLimitsWithDefaults() *SyncAgentEntityLimits`

NewSyncAgentEntityLimitsWithDefaults instantiates a new SyncAgentEntityLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitType

`func (o *SyncAgentEntityLimits) GetLimitType() string`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *SyncAgentEntityLimits) GetLimitTypeOk() (*string, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *SyncAgentEntityLimits) SetLimitType(v string)`

SetLimitType sets LimitType field to given value.


### GetLimits

`func (o *SyncAgentEntityLimits) GetLimits() []SyncAgentEntityLimitsLimitsInner`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *SyncAgentEntityLimits) GetLimitsOk() (*[]SyncAgentEntityLimitsLimitsInner, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *SyncAgentEntityLimits) SetLimits(v []SyncAgentEntityLimitsLimitsInner)`

SetLimits sets Limits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


