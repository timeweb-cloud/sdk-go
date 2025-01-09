# DedicatedServerPreset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID тарифа выделенного сервера. | 
**Description** | **string** | Описание характеристик тарифа выделенного сервера. | 
**IsIpmiEnabled** | **bool** | Это логическое значение, которое показывает, доступен ли IPMI у данного тарифа. | 
**IsPreInstalled** | **bool** | Это логическое значение, которое показывает, готов ли выделенный сервер к моментальной выдаче. | 
**Cpu** | [**DedicatedServerPresetCpu**](DedicatedServerPresetCpu.md) |  | 
**Disk** | [**DedicatedServerPresetDisk**](DedicatedServerPresetDisk.md) |  | 
**Price** | Pointer to **float32** | Стоимость тарифа выделенного сервера | [optional] 
**Memory** | [**DedicatedServerPresetMemory**](DedicatedServerPresetMemory.md) |  | 
**Location** | **string** | Локация. | 

## Methods

### NewDedicatedServerPreset

`func NewDedicatedServerPreset(id float32, description string, isIpmiEnabled bool, isPreInstalled bool, cpu DedicatedServerPresetCpu, disk DedicatedServerPresetDisk, memory DedicatedServerPresetMemory, location string, ) *DedicatedServerPreset`

NewDedicatedServerPreset instantiates a new DedicatedServerPreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerPresetWithDefaults

`func NewDedicatedServerPresetWithDefaults() *DedicatedServerPreset`

NewDedicatedServerPresetWithDefaults instantiates a new DedicatedServerPreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DedicatedServerPreset) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DedicatedServerPreset) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DedicatedServerPreset) SetId(v float32)`

SetId sets Id field to given value.


### GetDescription

`func (o *DedicatedServerPreset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DedicatedServerPreset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DedicatedServerPreset) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsIpmiEnabled

`func (o *DedicatedServerPreset) GetIsIpmiEnabled() bool`

GetIsIpmiEnabled returns the IsIpmiEnabled field if non-nil, zero value otherwise.

### GetIsIpmiEnabledOk

`func (o *DedicatedServerPreset) GetIsIpmiEnabledOk() (*bool, bool)`

GetIsIpmiEnabledOk returns a tuple with the IsIpmiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIpmiEnabled

`func (o *DedicatedServerPreset) SetIsIpmiEnabled(v bool)`

SetIsIpmiEnabled sets IsIpmiEnabled field to given value.


### GetIsPreInstalled

`func (o *DedicatedServerPreset) GetIsPreInstalled() bool`

GetIsPreInstalled returns the IsPreInstalled field if non-nil, zero value otherwise.

### GetIsPreInstalledOk

`func (o *DedicatedServerPreset) GetIsPreInstalledOk() (*bool, bool)`

GetIsPreInstalledOk returns a tuple with the IsPreInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreInstalled

`func (o *DedicatedServerPreset) SetIsPreInstalled(v bool)`

SetIsPreInstalled sets IsPreInstalled field to given value.


### GetCpu

`func (o *DedicatedServerPreset) GetCpu() DedicatedServerPresetCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DedicatedServerPreset) GetCpuOk() (*DedicatedServerPresetCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DedicatedServerPreset) SetCpu(v DedicatedServerPresetCpu)`

SetCpu sets Cpu field to given value.


### GetDisk

`func (o *DedicatedServerPreset) GetDisk() DedicatedServerPresetDisk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *DedicatedServerPreset) GetDiskOk() (*DedicatedServerPresetDisk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *DedicatedServerPreset) SetDisk(v DedicatedServerPresetDisk)`

SetDisk sets Disk field to given value.


### GetPrice

`func (o *DedicatedServerPreset) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DedicatedServerPreset) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DedicatedServerPreset) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DedicatedServerPreset) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetMemory

`func (o *DedicatedServerPreset) GetMemory() DedicatedServerPresetMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *DedicatedServerPreset) GetMemoryOk() (*DedicatedServerPresetMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *DedicatedServerPreset) SetMemory(v DedicatedServerPresetMemory)`

SetMemory sets Memory field to given value.


### GetLocation

`func (o *DedicatedServerPreset) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DedicatedServerPreset) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DedicatedServerPreset) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


