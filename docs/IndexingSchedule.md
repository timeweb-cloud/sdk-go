# IndexingSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **[]string** | Дни недели, в которые выполняется переиндексация (минимум один) | 
**Time** | **string** | Время запуска в формате &#x60;HH:mm&#x60; (24 часа) | 

## Methods

### NewIndexingSchedule

`func NewIndexingSchedule(days []string, time string, ) *IndexingSchedule`

NewIndexingSchedule instantiates a new IndexingSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexingScheduleWithDefaults

`func NewIndexingScheduleWithDefaults() *IndexingSchedule`

NewIndexingScheduleWithDefaults instantiates a new IndexingSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *IndexingSchedule) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *IndexingSchedule) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *IndexingSchedule) SetDays(v []string)`

SetDays sets Days field to given value.


### GetTime

`func (o *IndexingSchedule) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *IndexingSchedule) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *IndexingSchedule) SetTime(v string)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


