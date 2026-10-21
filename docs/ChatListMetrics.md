# ChatListMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalReactions** | **int32** | Количество чатов с учётом фильтра &#x60;response_time_ranges&#x60; (без учёта &#x60;reaction_type&#x60;) | 
**TotalTime** | **int32** | Количество чатов с учётом фильтра &#x60;reaction_type&#x60; (без учёта &#x60;response_time_ranges&#x60;) | 
**WithLikes** | **int32** | Чатов хотя бы с одной положительной оценкой | 
**WithDislikes** | **int32** | Чатов хотя бы с одной отрицательной оценкой | 
**WithAny** | **int32** | Чатов хотя бы с одной оценкой | 
**ResponseTimeGroupsReactions** | **map[string]int32** | Количество чатов по среднему времени ответа (в секундах) с учётом фильтра &#x60;response_time_ranges&#x60; | 
**ResponseTimeGroupsTime** | **map[string]int32** | Количество чатов по среднему времени ответа (в секундах) с учётом фильтра &#x60;reaction_type&#x60; | 

## Methods

### NewChatListMetrics

`func NewChatListMetrics(totalReactions int32, totalTime int32, withLikes int32, withDislikes int32, withAny int32, responseTimeGroupsReactions map[string]int32, responseTimeGroupsTime map[string]int32, ) *ChatListMetrics`

NewChatListMetrics instantiates a new ChatListMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatListMetricsWithDefaults

`func NewChatListMetricsWithDefaults() *ChatListMetrics`

NewChatListMetricsWithDefaults instantiates a new ChatListMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalReactions

`func (o *ChatListMetrics) GetTotalReactions() int32`

GetTotalReactions returns the TotalReactions field if non-nil, zero value otherwise.

### GetTotalReactionsOk

`func (o *ChatListMetrics) GetTotalReactionsOk() (*int32, bool)`

GetTotalReactionsOk returns a tuple with the TotalReactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReactions

`func (o *ChatListMetrics) SetTotalReactions(v int32)`

SetTotalReactions sets TotalReactions field to given value.


### GetTotalTime

`func (o *ChatListMetrics) GetTotalTime() int32`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *ChatListMetrics) GetTotalTimeOk() (*int32, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *ChatListMetrics) SetTotalTime(v int32)`

SetTotalTime sets TotalTime field to given value.


### GetWithLikes

`func (o *ChatListMetrics) GetWithLikes() int32`

GetWithLikes returns the WithLikes field if non-nil, zero value otherwise.

### GetWithLikesOk

`func (o *ChatListMetrics) GetWithLikesOk() (*int32, bool)`

GetWithLikesOk returns a tuple with the WithLikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithLikes

`func (o *ChatListMetrics) SetWithLikes(v int32)`

SetWithLikes sets WithLikes field to given value.


### GetWithDislikes

`func (o *ChatListMetrics) GetWithDislikes() int32`

GetWithDislikes returns the WithDislikes field if non-nil, zero value otherwise.

### GetWithDislikesOk

`func (o *ChatListMetrics) GetWithDislikesOk() (*int32, bool)`

GetWithDislikesOk returns a tuple with the WithDislikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithDislikes

`func (o *ChatListMetrics) SetWithDislikes(v int32)`

SetWithDislikes sets WithDislikes field to given value.


### GetWithAny

`func (o *ChatListMetrics) GetWithAny() int32`

GetWithAny returns the WithAny field if non-nil, zero value otherwise.

### GetWithAnyOk

`func (o *ChatListMetrics) GetWithAnyOk() (*int32, bool)`

GetWithAnyOk returns a tuple with the WithAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAny

`func (o *ChatListMetrics) SetWithAny(v int32)`

SetWithAny sets WithAny field to given value.


### GetResponseTimeGroupsReactions

`func (o *ChatListMetrics) GetResponseTimeGroupsReactions() map[string]int32`

GetResponseTimeGroupsReactions returns the ResponseTimeGroupsReactions field if non-nil, zero value otherwise.

### GetResponseTimeGroupsReactionsOk

`func (o *ChatListMetrics) GetResponseTimeGroupsReactionsOk() (*map[string]int32, bool)`

GetResponseTimeGroupsReactionsOk returns a tuple with the ResponseTimeGroupsReactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeGroupsReactions

`func (o *ChatListMetrics) SetResponseTimeGroupsReactions(v map[string]int32)`

SetResponseTimeGroupsReactions sets ResponseTimeGroupsReactions field to given value.


### GetResponseTimeGroupsTime

`func (o *ChatListMetrics) GetResponseTimeGroupsTime() map[string]int32`

GetResponseTimeGroupsTime returns the ResponseTimeGroupsTime field if non-nil, zero value otherwise.

### GetResponseTimeGroupsTimeOk

`func (o *ChatListMetrics) GetResponseTimeGroupsTimeOk() (*map[string]int32, bool)`

GetResponseTimeGroupsTimeOk returns a tuple with the ResponseTimeGroupsTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeGroupsTime

`func (o *ChatListMetrics) SetResponseTimeGroupsTime(v map[string]int32)`

SetResponseTimeGroupsTime sets ResponseTimeGroupsTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


