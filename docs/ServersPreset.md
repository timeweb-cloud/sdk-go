# ServersPreset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID тарифа сервера. | 
**Location** | **string** | Локация сервера. | 
**Price** | **float32** | Стоимость в рублях. | 
**Cpu** | **float32** | Количество ядер процессора. | 
**CpuFrequency** | **string** | Частота процессора. | 
**Ram** | **float32** | Количество (в Мб) оперативной памяти. | 
**Disk** | **float32** | Размер диска (в Мб). | 
**DiskType** | **string** | Тип диска. | 
**Bandwidth** | **float32** | Пропускная способность тарифа. | 
**Description** | **string** | Описание тарифа. | 
**DescriptionShort** | **string** | Короткое описание тарифа. | 
**IsAllowedLocalNetwork** | **bool** | Есть возможность подключения локальной сети | 
**Tags** | **[]string** | Список тегов тарифа. | 

## Methods

### NewServersPreset

`func NewServersPreset(id float32, location string, price float32, cpu float32, cpuFrequency string, ram float32, disk float32, diskType string, bandwidth float32, description string, descriptionShort string, isAllowedLocalNetwork bool, tags []string, ) *ServersPreset`

NewServersPreset instantiates a new ServersPreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServersPresetWithDefaults

`func NewServersPresetWithDefaults() *ServersPreset`

NewServersPresetWithDefaults instantiates a new ServersPreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServersPreset) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServersPreset) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServersPreset) SetId(v float32)`

SetId sets Id field to given value.


### GetLocation

`func (o *ServersPreset) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ServersPreset) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ServersPreset) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetPrice

`func (o *ServersPreset) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ServersPreset) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ServersPreset) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetCpu

`func (o *ServersPreset) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ServersPreset) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ServersPreset) SetCpu(v float32)`

SetCpu sets Cpu field to given value.


### GetCpuFrequency

`func (o *ServersPreset) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *ServersPreset) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *ServersPreset) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.


### GetRam

`func (o *ServersPreset) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ServersPreset) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ServersPreset) SetRam(v float32)`

SetRam sets Ram field to given value.


### GetDisk

`func (o *ServersPreset) GetDisk() float32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ServersPreset) GetDiskOk() (*float32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ServersPreset) SetDisk(v float32)`

SetDisk sets Disk field to given value.


### GetDiskType

`func (o *ServersPreset) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *ServersPreset) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *ServersPreset) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.


### GetBandwidth

`func (o *ServersPreset) GetBandwidth() float32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ServersPreset) GetBandwidthOk() (*float32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ServersPreset) SetBandwidth(v float32)`

SetBandwidth sets Bandwidth field to given value.


### GetDescription

`func (o *ServersPreset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServersPreset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServersPreset) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionShort

`func (o *ServersPreset) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *ServersPreset) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *ServersPreset) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.


### GetIsAllowedLocalNetwork

`func (o *ServersPreset) GetIsAllowedLocalNetwork() bool`

GetIsAllowedLocalNetwork returns the IsAllowedLocalNetwork field if non-nil, zero value otherwise.

### GetIsAllowedLocalNetworkOk

`func (o *ServersPreset) GetIsAllowedLocalNetworkOk() (*bool, bool)`

GetIsAllowedLocalNetworkOk returns a tuple with the IsAllowedLocalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowedLocalNetwork

`func (o *ServersPreset) SetIsAllowedLocalNetwork(v bool)`

SetIsAllowedLocalNetwork sets IsAllowedLocalNetwork field to given value.


### GetTags

`func (o *ServersPreset) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServersPreset) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServersPreset) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


