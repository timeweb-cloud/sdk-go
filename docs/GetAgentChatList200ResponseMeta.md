# GetAgentChatList200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Количество чатов с учётом фильтров | 
**Metrics** | Pointer to [**ChatListMetrics**](ChatListMetrics.md) |  | [optional] 

## Methods

### NewGetAgentChatList200ResponseMeta

`func NewGetAgentChatList200ResponseMeta(total int32, ) *GetAgentChatList200ResponseMeta`

NewGetAgentChatList200ResponseMeta instantiates a new GetAgentChatList200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentChatList200ResponseMetaWithDefaults

`func NewGetAgentChatList200ResponseMetaWithDefaults() *GetAgentChatList200ResponseMeta`

NewGetAgentChatList200ResponseMetaWithDefaults instantiates a new GetAgentChatList200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetAgentChatList200ResponseMeta) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAgentChatList200ResponseMeta) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAgentChatList200ResponseMeta) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetMetrics

`func (o *GetAgentChatList200ResponseMeta) GetMetrics() ChatListMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetAgentChatList200ResponseMeta) GetMetricsOk() (*ChatListMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetAgentChatList200ResponseMeta) SetMetrics(v ChatListMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetAgentChatList200ResponseMeta) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


