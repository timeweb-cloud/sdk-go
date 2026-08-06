# StatisticsRequestsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoggedAt** | **time.Time** | Начало часового интервала в UTC. | 
**Status2xx** | **float32** | Количество ответов со статусом 2xx. | 
**Status3xx** | **float32** | Количество ответов со статусом 3xx. | 
**Status4xx** | **float32** | Количество ответов со статусом 4xx. | 
**Status5xx** | **float32** | Количество ответов со статусом 5xx. | 

## Methods

### NewStatisticsRequestsInner

`func NewStatisticsRequestsInner(loggedAt time.Time, status2xx float32, status3xx float32, status4xx float32, status5xx float32, ) *StatisticsRequestsInner`

NewStatisticsRequestsInner instantiates a new StatisticsRequestsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsRequestsInnerWithDefaults

`func NewStatisticsRequestsInnerWithDefaults() *StatisticsRequestsInner`

NewStatisticsRequestsInnerWithDefaults instantiates a new StatisticsRequestsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoggedAt

`func (o *StatisticsRequestsInner) GetLoggedAt() time.Time`

GetLoggedAt returns the LoggedAt field if non-nil, zero value otherwise.

### GetLoggedAtOk

`func (o *StatisticsRequestsInner) GetLoggedAtOk() (*time.Time, bool)`

GetLoggedAtOk returns a tuple with the LoggedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedAt

`func (o *StatisticsRequestsInner) SetLoggedAt(v time.Time)`

SetLoggedAt sets LoggedAt field to given value.


### GetStatus2xx

`func (o *StatisticsRequestsInner) GetStatus2xx() float32`

GetStatus2xx returns the Status2xx field if non-nil, zero value otherwise.

### GetStatus2xxOk

`func (o *StatisticsRequestsInner) GetStatus2xxOk() (*float32, bool)`

GetStatus2xxOk returns a tuple with the Status2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus2xx

`func (o *StatisticsRequestsInner) SetStatus2xx(v float32)`

SetStatus2xx sets Status2xx field to given value.


### GetStatus3xx

`func (o *StatisticsRequestsInner) GetStatus3xx() float32`

GetStatus3xx returns the Status3xx field if non-nil, zero value otherwise.

### GetStatus3xxOk

`func (o *StatisticsRequestsInner) GetStatus3xxOk() (*float32, bool)`

GetStatus3xxOk returns a tuple with the Status3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus3xx

`func (o *StatisticsRequestsInner) SetStatus3xx(v float32)`

SetStatus3xx sets Status3xx field to given value.


### GetStatus4xx

`func (o *StatisticsRequestsInner) GetStatus4xx() float32`

GetStatus4xx returns the Status4xx field if non-nil, zero value otherwise.

### GetStatus4xxOk

`func (o *StatisticsRequestsInner) GetStatus4xxOk() (*float32, bool)`

GetStatus4xxOk returns a tuple with the Status4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus4xx

`func (o *StatisticsRequestsInner) SetStatus4xx(v float32)`

SetStatus4xx sets Status4xx field to given value.


### GetStatus5xx

`func (o *StatisticsRequestsInner) GetStatus5xx() float32`

GetStatus5xx returns the Status5xx field if non-nil, zero value otherwise.

### GetStatus5xxOk

`func (o *StatisticsRequestsInner) GetStatus5xxOk() (*float32, bool)`

GetStatus5xxOk returns a tuple with the Status5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus5xx

`func (o *StatisticsRequestsInner) SetStatus5xx(v float32)`

SetStatus5xx sets Status5xx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


