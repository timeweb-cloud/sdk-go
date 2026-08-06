# Statistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Traffic** | [**[]StatisticsTrafficInner**](StatisticsTrafficInner.md) | Исходящий трафик по часовым интервалам. | 
**Requests** | [**[]StatisticsRequestsInner**](StatisticsRequestsInner.md) | Количество запросов по часовым интервалам с разбивкой по классам HTTP-статуса ответа. | 

## Methods

### NewStatistics

`func NewStatistics(traffic []StatisticsTrafficInner, requests []StatisticsRequestsInner, ) *Statistics`

NewStatistics instantiates a new Statistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsWithDefaults

`func NewStatisticsWithDefaults() *Statistics`

NewStatisticsWithDefaults instantiates a new Statistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraffic

`func (o *Statistics) GetTraffic() []StatisticsTrafficInner`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *Statistics) GetTrafficOk() (*[]StatisticsTrafficInner, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *Statistics) SetTraffic(v []StatisticsTrafficInner)`

SetTraffic sets Traffic field to given value.


### GetRequests

`func (o *Statistics) GetRequests() []StatisticsRequestsInner`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *Statistics) GetRequestsOk() (*[]StatisticsRequestsInner, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *Statistics) SetRequests(v []StatisticsRequestsInner)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


