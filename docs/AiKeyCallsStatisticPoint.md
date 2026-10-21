# AiKeyCallsStatisticPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | Начало интервала | 
**CallCount** | **int32** | Количество запросов за интервал | 

## Methods

### NewAiKeyCallsStatisticPoint

`func NewAiKeyCallsStatisticPoint(time time.Time, callCount int32, ) *AiKeyCallsStatisticPoint`

NewAiKeyCallsStatisticPoint instantiates a new AiKeyCallsStatisticPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiKeyCallsStatisticPointWithDefaults

`func NewAiKeyCallsStatisticPointWithDefaults() *AiKeyCallsStatisticPoint`

NewAiKeyCallsStatisticPointWithDefaults instantiates a new AiKeyCallsStatisticPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *AiKeyCallsStatisticPoint) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AiKeyCallsStatisticPoint) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AiKeyCallsStatisticPoint) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetCallCount

`func (o *AiKeyCallsStatisticPoint) GetCallCount() int32`

GetCallCount returns the CallCount field if non-nil, zero value otherwise.

### GetCallCountOk

`func (o *AiKeyCallsStatisticPoint) GetCallCountOk() (*int32, bool)`

GetCallCountOk returns a tuple with the CallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallCount

`func (o *AiKeyCallsStatisticPoint) SetCallCount(v int32)`

SetCallCount sets CallCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


