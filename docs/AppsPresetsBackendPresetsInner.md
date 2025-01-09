# AppsPresetsBackendPresetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID пресета. | 
**Description** | Pointer to **string** | Описание пресета. | [optional] 
**DescriptionShort** | **string** | Краткое описание пресета. | 
**Price** | **int32** | Цена пресета. | 
**Cpu** | **int32** | Количество CPU. | 
**Ram** | **int32** | Объем RAM в МБ. | 
**Disk** | **int32** | Объем диска в МБ. | 
**Location** | **string** | Локация. | 
**CpuFrequency** | Pointer to **string** | Частота CPU. | [optional] 

## Methods

### NewAppsPresetsBackendPresetsInner

`func NewAppsPresetsBackendPresetsInner(id int32, descriptionShort string, price int32, cpu int32, ram int32, disk int32, location string, ) *AppsPresetsBackendPresetsInner`

NewAppsPresetsBackendPresetsInner instantiates a new AppsPresetsBackendPresetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsPresetsBackendPresetsInnerWithDefaults

`func NewAppsPresetsBackendPresetsInnerWithDefaults() *AppsPresetsBackendPresetsInner`

NewAppsPresetsBackendPresetsInnerWithDefaults instantiates a new AppsPresetsBackendPresetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppsPresetsBackendPresetsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppsPresetsBackendPresetsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppsPresetsBackendPresetsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetDescription

`func (o *AppsPresetsBackendPresetsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppsPresetsBackendPresetsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppsPresetsBackendPresetsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppsPresetsBackendPresetsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionShort

`func (o *AppsPresetsBackendPresetsInner) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *AppsPresetsBackendPresetsInner) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *AppsPresetsBackendPresetsInner) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.


### GetPrice

`func (o *AppsPresetsBackendPresetsInner) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AppsPresetsBackendPresetsInner) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AppsPresetsBackendPresetsInner) SetPrice(v int32)`

SetPrice sets Price field to given value.


### GetCpu

`func (o *AppsPresetsBackendPresetsInner) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *AppsPresetsBackendPresetsInner) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *AppsPresetsBackendPresetsInner) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *AppsPresetsBackendPresetsInner) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *AppsPresetsBackendPresetsInner) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *AppsPresetsBackendPresetsInner) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetDisk

`func (o *AppsPresetsBackendPresetsInner) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *AppsPresetsBackendPresetsInner) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *AppsPresetsBackendPresetsInner) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetLocation

`func (o *AppsPresetsBackendPresetsInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AppsPresetsBackendPresetsInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AppsPresetsBackendPresetsInner) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetCpuFrequency

`func (o *AppsPresetsBackendPresetsInner) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *AppsPresetsBackendPresetsInner) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *AppsPresetsBackendPresetsInner) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.

### HasCpuFrequency

`func (o *AppsPresetsBackendPresetsInner) HasCpuFrequency() bool`

HasCpuFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


