# StatisticsTrafficInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoggedAt** | **time.Time** | Начало часового интервала в UTC. | 
**OutgoingBytes** | **float32** | Исходящий трафик за интервал, в байтах. | 

## Methods

### NewStatisticsTrafficInner

`func NewStatisticsTrafficInner(loggedAt time.Time, outgoingBytes float32, ) *StatisticsTrafficInner`

NewStatisticsTrafficInner instantiates a new StatisticsTrafficInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsTrafficInnerWithDefaults

`func NewStatisticsTrafficInnerWithDefaults() *StatisticsTrafficInner`

NewStatisticsTrafficInnerWithDefaults instantiates a new StatisticsTrafficInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoggedAt

`func (o *StatisticsTrafficInner) GetLoggedAt() time.Time`

GetLoggedAt returns the LoggedAt field if non-nil, zero value otherwise.

### GetLoggedAtOk

`func (o *StatisticsTrafficInner) GetLoggedAtOk() (*time.Time, bool)`

GetLoggedAtOk returns a tuple with the LoggedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedAt

`func (o *StatisticsTrafficInner) SetLoggedAt(v time.Time)`

SetLoggedAt sets LoggedAt field to given value.


### GetOutgoingBytes

`func (o *StatisticsTrafficInner) GetOutgoingBytes() float32`

GetOutgoingBytes returns the OutgoingBytes field if non-nil, zero value otherwise.

### GetOutgoingBytesOk

`func (o *StatisticsTrafficInner) GetOutgoingBytesOk() (*float32, bool)`

GetOutgoingBytesOk returns a tuple with the OutgoingBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingBytes

`func (o *StatisticsTrafficInner) SetOutgoingBytes(v float32)`

SetOutgoingBytes sets OutgoingBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


