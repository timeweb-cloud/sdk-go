# UpdateAutoBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включено ли автобэкапирование | 
**CopyCount** | Pointer to **float32** | Количество копий для хранения. Минимальное количество &#x60;1&#x60;, максимальное &#x60;99&#x60;. Обязательно при &#x60;is_enabled&#x60;: &#x60;true&#x60;. | [optional] 
**CreationStartAt** | Pointer to **time.Time** | Дата начала создания первого автобэкапа. Значение в формате &#x60;ISO8601&#x60;. Время не учитывается. Обязательно при &#x60;is_enabled&#x60;: &#x60;true&#x60;. | [optional] 
**Interval** | Pointer to **string** | Периодичность создания автобэкапов. Обязательно при &#x60;is_enabled&#x60;: &#x60;true&#x60;. | [optional] 
**DayOfWeek** | Pointer to **float32** | День недели, в который будут создаваться автобэкапы. Доступные значения от &#x60;1&#x60; до &#x60;7&#x60;. Обязательно при &#x60;is_enabled&#x60;: &#x60;true&#x60; при любой периодичности, но на расписание влияет только при значении &#x60;interval&#x60;: &#x60;week&#x60;. | [optional] 

## Methods

### NewUpdateAutoBackup

`func NewUpdateAutoBackup(isEnabled bool, ) *UpdateAutoBackup`

NewUpdateAutoBackup instantiates a new UpdateAutoBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAutoBackupWithDefaults

`func NewUpdateAutoBackupWithDefaults() *UpdateAutoBackup`

NewUpdateAutoBackupWithDefaults instantiates a new UpdateAutoBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *UpdateAutoBackup) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateAutoBackup) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateAutoBackup) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetCopyCount

`func (o *UpdateAutoBackup) GetCopyCount() float32`

GetCopyCount returns the CopyCount field if non-nil, zero value otherwise.

### GetCopyCountOk

`func (o *UpdateAutoBackup) GetCopyCountOk() (*float32, bool)`

GetCopyCountOk returns a tuple with the CopyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyCount

`func (o *UpdateAutoBackup) SetCopyCount(v float32)`

SetCopyCount sets CopyCount field to given value.

### HasCopyCount

`func (o *UpdateAutoBackup) HasCopyCount() bool`

HasCopyCount returns a boolean if a field has been set.

### GetCreationStartAt

`func (o *UpdateAutoBackup) GetCreationStartAt() time.Time`

GetCreationStartAt returns the CreationStartAt field if non-nil, zero value otherwise.

### GetCreationStartAtOk

`func (o *UpdateAutoBackup) GetCreationStartAtOk() (*time.Time, bool)`

GetCreationStartAtOk returns a tuple with the CreationStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStartAt

`func (o *UpdateAutoBackup) SetCreationStartAt(v time.Time)`

SetCreationStartAt sets CreationStartAt field to given value.

### HasCreationStartAt

`func (o *UpdateAutoBackup) HasCreationStartAt() bool`

HasCreationStartAt returns a boolean if a field has been set.

### GetInterval

`func (o *UpdateAutoBackup) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UpdateAutoBackup) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UpdateAutoBackup) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UpdateAutoBackup) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *UpdateAutoBackup) GetDayOfWeek() float32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *UpdateAutoBackup) GetDayOfWeekOk() (*float32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *UpdateAutoBackup) SetDayOfWeek(v float32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *UpdateAutoBackup) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


