# AgentTokenStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | Начало интервала | 
**ModelName** | Pointer to **string** | Модель, на которую пришлись запросы. В интервалах без запросов поле отсутствует. | [optional] 
**IngoingTokens** | **int32** | Входящие токены (промпты). Если для аккаунта включена тарификация кэша, токены из кэша считаются отдельно. | 
**OutgoingTokens** | **int32** | Исходящие токены (ответы) | 
**CacheReadTokens** | **int32** | Входящие токены, прочитанные из кэша | 
**CacheWriteTokens** | **int32** | Входящие токены, записанные в кэш | 
**TotalCost** | **NullableFloat32** | Стоимость токенов интервала до учёта бонусов, в валюте аккаунта. &#x60;null&#x60; — стоимость посчитать не удалось. | 

## Methods

### NewAgentTokenStatistic

`func NewAgentTokenStatistic(time time.Time, ingoingTokens int32, outgoingTokens int32, cacheReadTokens int32, cacheWriteTokens int32, totalCost NullableFloat32, ) *AgentTokenStatistic`

NewAgentTokenStatistic instantiates a new AgentTokenStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTokenStatisticWithDefaults

`func NewAgentTokenStatisticWithDefaults() *AgentTokenStatistic`

NewAgentTokenStatisticWithDefaults instantiates a new AgentTokenStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *AgentTokenStatistic) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AgentTokenStatistic) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AgentTokenStatistic) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetModelName

`func (o *AgentTokenStatistic) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *AgentTokenStatistic) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *AgentTokenStatistic) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *AgentTokenStatistic) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetIngoingTokens

`func (o *AgentTokenStatistic) GetIngoingTokens() int32`

GetIngoingTokens returns the IngoingTokens field if non-nil, zero value otherwise.

### GetIngoingTokensOk

`func (o *AgentTokenStatistic) GetIngoingTokensOk() (*int32, bool)`

GetIngoingTokensOk returns a tuple with the IngoingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngoingTokens

`func (o *AgentTokenStatistic) SetIngoingTokens(v int32)`

SetIngoingTokens sets IngoingTokens field to given value.


### GetOutgoingTokens

`func (o *AgentTokenStatistic) GetOutgoingTokens() int32`

GetOutgoingTokens returns the OutgoingTokens field if non-nil, zero value otherwise.

### GetOutgoingTokensOk

`func (o *AgentTokenStatistic) GetOutgoingTokensOk() (*int32, bool)`

GetOutgoingTokensOk returns a tuple with the OutgoingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTokens

`func (o *AgentTokenStatistic) SetOutgoingTokens(v int32)`

SetOutgoingTokens sets OutgoingTokens field to given value.


### GetCacheReadTokens

`func (o *AgentTokenStatistic) GetCacheReadTokens() int32`

GetCacheReadTokens returns the CacheReadTokens field if non-nil, zero value otherwise.

### GetCacheReadTokensOk

`func (o *AgentTokenStatistic) GetCacheReadTokensOk() (*int32, bool)`

GetCacheReadTokensOk returns a tuple with the CacheReadTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadTokens

`func (o *AgentTokenStatistic) SetCacheReadTokens(v int32)`

SetCacheReadTokens sets CacheReadTokens field to given value.


### GetCacheWriteTokens

`func (o *AgentTokenStatistic) GetCacheWriteTokens() int32`

GetCacheWriteTokens returns the CacheWriteTokens field if non-nil, zero value otherwise.

### GetCacheWriteTokensOk

`func (o *AgentTokenStatistic) GetCacheWriteTokensOk() (*int32, bool)`

GetCacheWriteTokensOk returns a tuple with the CacheWriteTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWriteTokens

`func (o *AgentTokenStatistic) SetCacheWriteTokens(v int32)`

SetCacheWriteTokens sets CacheWriteTokens field to given value.


### GetTotalCost

`func (o *AgentTokenStatistic) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *AgentTokenStatistic) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *AgentTokenStatistic) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.


### SetTotalCostNil

`func (o *AgentTokenStatistic) SetTotalCostNil(b bool)`

 SetTotalCostNil sets the value for TotalCost to be an explicit nil

### UnsetTotalCost
`func (o *AgentTokenStatistic) UnsetTotalCost()`

UnsetTotalCost ensures that no value is present for TotalCost, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


