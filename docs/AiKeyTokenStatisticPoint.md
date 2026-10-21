# AiKeyTokenStatisticPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | Начало интервала | 
**IngoingTokens** | **int32** | Входные токены (запрос) | 
**OutgoingTokens** | **int32** | Выходные токены (ответ модели) | 
**CacheReadTokens** | **int32** | Входные токены, прочитанные из кэша | 
**CacheWriteTokens** | **int32** | Входные токены, записанные в кэш | 
**TotalCost** | **NullableFloat32** | Стоимость запросов за интервал с учётом скидок аккаунта, до списания бонусов, в валюте аккаунта. &#x60;null&#x60; — стоимость части запросов интервала определить не удалось | 

## Methods

### NewAiKeyTokenStatisticPoint

`func NewAiKeyTokenStatisticPoint(time time.Time, ingoingTokens int32, outgoingTokens int32, cacheReadTokens int32, cacheWriteTokens int32, totalCost NullableFloat32, ) *AiKeyTokenStatisticPoint`

NewAiKeyTokenStatisticPoint instantiates a new AiKeyTokenStatisticPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiKeyTokenStatisticPointWithDefaults

`func NewAiKeyTokenStatisticPointWithDefaults() *AiKeyTokenStatisticPoint`

NewAiKeyTokenStatisticPointWithDefaults instantiates a new AiKeyTokenStatisticPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *AiKeyTokenStatisticPoint) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AiKeyTokenStatisticPoint) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AiKeyTokenStatisticPoint) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetIngoingTokens

`func (o *AiKeyTokenStatisticPoint) GetIngoingTokens() int32`

GetIngoingTokens returns the IngoingTokens field if non-nil, zero value otherwise.

### GetIngoingTokensOk

`func (o *AiKeyTokenStatisticPoint) GetIngoingTokensOk() (*int32, bool)`

GetIngoingTokensOk returns a tuple with the IngoingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngoingTokens

`func (o *AiKeyTokenStatisticPoint) SetIngoingTokens(v int32)`

SetIngoingTokens sets IngoingTokens field to given value.


### GetOutgoingTokens

`func (o *AiKeyTokenStatisticPoint) GetOutgoingTokens() int32`

GetOutgoingTokens returns the OutgoingTokens field if non-nil, zero value otherwise.

### GetOutgoingTokensOk

`func (o *AiKeyTokenStatisticPoint) GetOutgoingTokensOk() (*int32, bool)`

GetOutgoingTokensOk returns a tuple with the OutgoingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTokens

`func (o *AiKeyTokenStatisticPoint) SetOutgoingTokens(v int32)`

SetOutgoingTokens sets OutgoingTokens field to given value.


### GetCacheReadTokens

`func (o *AiKeyTokenStatisticPoint) GetCacheReadTokens() int32`

GetCacheReadTokens returns the CacheReadTokens field if non-nil, zero value otherwise.

### GetCacheReadTokensOk

`func (o *AiKeyTokenStatisticPoint) GetCacheReadTokensOk() (*int32, bool)`

GetCacheReadTokensOk returns a tuple with the CacheReadTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadTokens

`func (o *AiKeyTokenStatisticPoint) SetCacheReadTokens(v int32)`

SetCacheReadTokens sets CacheReadTokens field to given value.


### GetCacheWriteTokens

`func (o *AiKeyTokenStatisticPoint) GetCacheWriteTokens() int32`

GetCacheWriteTokens returns the CacheWriteTokens field if non-nil, zero value otherwise.

### GetCacheWriteTokensOk

`func (o *AiKeyTokenStatisticPoint) GetCacheWriteTokensOk() (*int32, bool)`

GetCacheWriteTokensOk returns a tuple with the CacheWriteTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWriteTokens

`func (o *AiKeyTokenStatisticPoint) SetCacheWriteTokens(v int32)`

SetCacheWriteTokens sets CacheWriteTokens field to given value.


### GetTotalCost

`func (o *AiKeyTokenStatisticPoint) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *AiKeyTokenStatisticPoint) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *AiKeyTokenStatisticPoint) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.


### SetTotalCostNil

`func (o *AiKeyTokenStatisticPoint) SetTotalCostNil(b bool)`

 SetTotalCostNil sets the value for TotalCost to be an explicit nil

### UnsetTotalCost
`func (o *AiKeyTokenStatisticPoint) UnsetTotalCost()`

UnsetTotalCost ensures that no value is present for TotalCost, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


