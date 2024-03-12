# CreateDbAutoBackups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyCount** | **float32** | Количество копий для хранения. Минимальное количество &#x60;1&#x60;, максимальное &#x60;99&#x60; | 
**CreationStartAt** | **time.Time** | Дата начала создания первого автобэкапа. Значение в формате &#x60;ISO8601&#x60;. Время не учитывается. | 
**Interval** | **string** | Периодичность создания автобэкапов | 
**DayOfWeek** | **float32** | День недели, в который будут создаваться автобэкапы. Работает только со значением &#x60;interval&#x60;: &#x60;week&#x60;. Доступные значение от &#x60;1 &#x60;до &#x60;7&#x60;. | 

## Methods

### NewCreateDbAutoBackups

`func NewCreateDbAutoBackups(copyCount float32, creationStartAt time.Time, interval string, dayOfWeek float32, ) *CreateDbAutoBackups`

NewCreateDbAutoBackups instantiates a new CreateDbAutoBackups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDbAutoBackupsWithDefaults

`func NewCreateDbAutoBackupsWithDefaults() *CreateDbAutoBackups`

NewCreateDbAutoBackupsWithDefaults instantiates a new CreateDbAutoBackups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyCount

`func (o *CreateDbAutoBackups) GetCopyCount() float32`

GetCopyCount returns the CopyCount field if non-nil, zero value otherwise.

### GetCopyCountOk

`func (o *CreateDbAutoBackups) GetCopyCountOk() (*float32, bool)`

GetCopyCountOk returns a tuple with the CopyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyCount

`func (o *CreateDbAutoBackups) SetCopyCount(v float32)`

SetCopyCount sets CopyCount field to given value.


### GetCreationStartAt

`func (o *CreateDbAutoBackups) GetCreationStartAt() time.Time`

GetCreationStartAt returns the CreationStartAt field if non-nil, zero value otherwise.

### GetCreationStartAtOk

`func (o *CreateDbAutoBackups) GetCreationStartAtOk() (*time.Time, bool)`

GetCreationStartAtOk returns a tuple with the CreationStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStartAt

`func (o *CreateDbAutoBackups) SetCreationStartAt(v time.Time)`

SetCreationStartAt sets CreationStartAt field to given value.


### GetInterval

`func (o *CreateDbAutoBackups) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CreateDbAutoBackups) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CreateDbAutoBackups) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetDayOfWeek

`func (o *CreateDbAutoBackups) GetDayOfWeek() float32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *CreateDbAutoBackups) GetDayOfWeekOk() (*float32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *CreateDbAutoBackups) SetDayOfWeek(v float32)`

SetDayOfWeek sets DayOfWeek field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


