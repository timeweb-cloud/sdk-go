# AutoBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyCount** | Pointer to **float32** | Количество копий для хранения. Минимальное количество &#x60;1&#x60;, максимальное &#x60;99&#x60; | [optional] 
**CreationStartAt** | Pointer to **time.Time** | Дата начала создания первого автобэкапа. Значение в формате &#x60;ISO8601&#x60;. Время не учитывается. | [optional] 
**IsEnabled** | **bool** | Включено ли автобэкапирование | 
**Interval** | Pointer to **string** | Периодичность создания автобэкапов | [optional] 
**DayOfWeek** | Pointer to **float32** | День недели, в который будут создаваться автобэкапы. Работает только со значением &#x60;interval&#x60;: &#x60;week&#x60;. Доступные значение от &#x60;1 &#x60;до &#x60;7&#x60;. | [optional] 

## Methods

### NewAutoBackup

`func NewAutoBackup(isEnabled bool, ) *AutoBackup`

NewAutoBackup instantiates a new AutoBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoBackupWithDefaults

`func NewAutoBackupWithDefaults() *AutoBackup`

NewAutoBackupWithDefaults instantiates a new AutoBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyCount

`func (o *AutoBackup) GetCopyCount() float32`

GetCopyCount returns the CopyCount field if non-nil, zero value otherwise.

### GetCopyCountOk

`func (o *AutoBackup) GetCopyCountOk() (*float32, bool)`

GetCopyCountOk returns a tuple with the CopyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyCount

`func (o *AutoBackup) SetCopyCount(v float32)`

SetCopyCount sets CopyCount field to given value.

### HasCopyCount

`func (o *AutoBackup) HasCopyCount() bool`

HasCopyCount returns a boolean if a field has been set.

### GetCreationStartAt

`func (o *AutoBackup) GetCreationStartAt() time.Time`

GetCreationStartAt returns the CreationStartAt field if non-nil, zero value otherwise.

### GetCreationStartAtOk

`func (o *AutoBackup) GetCreationStartAtOk() (*time.Time, bool)`

GetCreationStartAtOk returns a tuple with the CreationStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStartAt

`func (o *AutoBackup) SetCreationStartAt(v time.Time)`

SetCreationStartAt sets CreationStartAt field to given value.

### HasCreationStartAt

`func (o *AutoBackup) HasCreationStartAt() bool`

HasCreationStartAt returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AutoBackup) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AutoBackup) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AutoBackup) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetInterval

`func (o *AutoBackup) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AutoBackup) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AutoBackup) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *AutoBackup) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *AutoBackup) GetDayOfWeek() float32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *AutoBackup) GetDayOfWeekOk() (*float32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *AutoBackup) SetDayOfWeek(v float32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *AutoBackup) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


