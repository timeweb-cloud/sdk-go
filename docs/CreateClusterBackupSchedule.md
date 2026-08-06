# CreateClusterBackupSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyCount** | Pointer to **int32** | Количество хранимых резервных копий. | [optional] 
**Interval** | Pointer to **string** | Периодичность создания резервных копий. | [optional] 
**DayOfWeek** | Pointer to **int32** | День недели (от 1 до 7) для создания резервной копии. Учитывается только при &#x60;interval: week&#x60;. | [optional] 
**DayOfMonth** | Pointer to **int32** | День месяца (от 1 до 28) для создания резервной копии. Учитывается только при &#x60;interval: month&#x60;. | [optional] 

## Methods

### NewCreateClusterBackupSchedule

`func NewCreateClusterBackupSchedule() *CreateClusterBackupSchedule`

NewCreateClusterBackupSchedule instantiates a new CreateClusterBackupSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterBackupScheduleWithDefaults

`func NewCreateClusterBackupScheduleWithDefaults() *CreateClusterBackupSchedule`

NewCreateClusterBackupScheduleWithDefaults instantiates a new CreateClusterBackupSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyCount

`func (o *CreateClusterBackupSchedule) GetCopyCount() int32`

GetCopyCount returns the CopyCount field if non-nil, zero value otherwise.

### GetCopyCountOk

`func (o *CreateClusterBackupSchedule) GetCopyCountOk() (*int32, bool)`

GetCopyCountOk returns a tuple with the CopyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyCount

`func (o *CreateClusterBackupSchedule) SetCopyCount(v int32)`

SetCopyCount sets CopyCount field to given value.

### HasCopyCount

`func (o *CreateClusterBackupSchedule) HasCopyCount() bool`

HasCopyCount returns a boolean if a field has been set.

### GetInterval

`func (o *CreateClusterBackupSchedule) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CreateClusterBackupSchedule) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CreateClusterBackupSchedule) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CreateClusterBackupSchedule) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *CreateClusterBackupSchedule) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *CreateClusterBackupSchedule) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *CreateClusterBackupSchedule) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *CreateClusterBackupSchedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *CreateClusterBackupSchedule) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *CreateClusterBackupSchedule) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *CreateClusterBackupSchedule) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *CreateClusterBackupSchedule) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


