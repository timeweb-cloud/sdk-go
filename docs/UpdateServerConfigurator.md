# UpdateServerConfigurator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguratorId** | Pointer to **float32** | ID конфигуратора сервера. | [optional] 
**Disk** | Pointer to **float32** | Размер диска в МБ. | [optional] 
**Cpu** | Pointer to **float32** | Количество ядер процессора. | [optional] 
**Ram** | Pointer to **float32** | Размер ОЗУ сервера в МБ. | [optional] 
**Gpu** | Pointer to **float32** | Количество видеокарт. | [optional] 

## Methods

### NewUpdateServerConfigurator

`func NewUpdateServerConfigurator() *UpdateServerConfigurator`

NewUpdateServerConfigurator instantiates a new UpdateServerConfigurator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerConfiguratorWithDefaults

`func NewUpdateServerConfiguratorWithDefaults() *UpdateServerConfigurator`

NewUpdateServerConfiguratorWithDefaults instantiates a new UpdateServerConfigurator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguratorId

`func (o *UpdateServerConfigurator) GetConfiguratorId() float32`

GetConfiguratorId returns the ConfiguratorId field if non-nil, zero value otherwise.

### GetConfiguratorIdOk

`func (o *UpdateServerConfigurator) GetConfiguratorIdOk() (*float32, bool)`

GetConfiguratorIdOk returns a tuple with the ConfiguratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguratorId

`func (o *UpdateServerConfigurator) SetConfiguratorId(v float32)`

SetConfiguratorId sets ConfiguratorId field to given value.

### HasConfiguratorId

`func (o *UpdateServerConfigurator) HasConfiguratorId() bool`

HasConfiguratorId returns a boolean if a field has been set.

### GetDisk

`func (o *UpdateServerConfigurator) GetDisk() float32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *UpdateServerConfigurator) GetDiskOk() (*float32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *UpdateServerConfigurator) SetDisk(v float32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *UpdateServerConfigurator) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetCpu

`func (o *UpdateServerConfigurator) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *UpdateServerConfigurator) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *UpdateServerConfigurator) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *UpdateServerConfigurator) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *UpdateServerConfigurator) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *UpdateServerConfigurator) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *UpdateServerConfigurator) SetRam(v float32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *UpdateServerConfigurator) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetGpu

`func (o *UpdateServerConfigurator) GetGpu() float32`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *UpdateServerConfigurator) GetGpuOk() (*float32, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *UpdateServerConfigurator) SetGpu(v float32)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *UpdateServerConfigurator) HasGpu() bool`

HasGpu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


