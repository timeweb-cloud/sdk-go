# TrafficUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | **float32** | Количество запросов к ресурсу с начала расчетного месяца. | 
**OutgoingTraffic** | **float32** | Исходящий трафик ресурса с начала расчетного месяца, в байтах. | 
**CacheRatio** | **float32** | Доля запросов, отданных из кэша, в процентах. | 

## Methods

### NewTrafficUsage

`func NewTrafficUsage(requests float32, outgoingTraffic float32, cacheRatio float32, ) *TrafficUsage`

NewTrafficUsage instantiates a new TrafficUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficUsageWithDefaults

`func NewTrafficUsageWithDefaults() *TrafficUsage`

NewTrafficUsageWithDefaults instantiates a new TrafficUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *TrafficUsage) GetRequests() float32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TrafficUsage) GetRequestsOk() (*float32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TrafficUsage) SetRequests(v float32)`

SetRequests sets Requests field to given value.


### GetOutgoingTraffic

`func (o *TrafficUsage) GetOutgoingTraffic() float32`

GetOutgoingTraffic returns the OutgoingTraffic field if non-nil, zero value otherwise.

### GetOutgoingTrafficOk

`func (o *TrafficUsage) GetOutgoingTrafficOk() (*float32, bool)`

GetOutgoingTrafficOk returns a tuple with the OutgoingTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTraffic

`func (o *TrafficUsage) SetOutgoingTraffic(v float32)`

SetOutgoingTraffic sets OutgoingTraffic field to given value.


### GetCacheRatio

`func (o *TrafficUsage) GetCacheRatio() float32`

GetCacheRatio returns the CacheRatio field if non-nil, zero value otherwise.

### GetCacheRatioOk

`func (o *TrafficUsage) GetCacheRatioOk() (*float32, bool)`

GetCacheRatioOk returns a tuple with the CacheRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheRatio

`func (o *TrafficUsage) SetCacheRatio(v float32)`

SetCacheRatio sets CacheRatio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


