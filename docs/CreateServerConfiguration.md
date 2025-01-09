# CreateServerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguratorId** | **float32** | ID конфигуратора сервера. | 
**Disk** | **float32** | Размер диска в МБ. | 
**Cpu** | **float32** | Количество ядер процессора. | 
**Ram** | **float32** | Размер ОЗУ сервера в МБ. | 
**Gpu** | Pointer to **float32** | Количество видеокарт. | [optional] 

## Methods

### NewCreateServerConfiguration

`func NewCreateServerConfiguration(configuratorId float32, disk float32, cpu float32, ram float32, ) *CreateServerConfiguration`

NewCreateServerConfiguration instantiates a new CreateServerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerConfigurationWithDefaults

`func NewCreateServerConfigurationWithDefaults() *CreateServerConfiguration`

NewCreateServerConfigurationWithDefaults instantiates a new CreateServerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguratorId

`func (o *CreateServerConfiguration) GetConfiguratorId() float32`

GetConfiguratorId returns the ConfiguratorId field if non-nil, zero value otherwise.

### GetConfiguratorIdOk

`func (o *CreateServerConfiguration) GetConfiguratorIdOk() (*float32, bool)`

GetConfiguratorIdOk returns a tuple with the ConfiguratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguratorId

`func (o *CreateServerConfiguration) SetConfiguratorId(v float32)`

SetConfiguratorId sets ConfiguratorId field to given value.


### GetDisk

`func (o *CreateServerConfiguration) GetDisk() float32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *CreateServerConfiguration) GetDiskOk() (*float32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *CreateServerConfiguration) SetDisk(v float32)`

SetDisk sets Disk field to given value.


### GetCpu

`func (o *CreateServerConfiguration) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *CreateServerConfiguration) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *CreateServerConfiguration) SetCpu(v float32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *CreateServerConfiguration) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *CreateServerConfiguration) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *CreateServerConfiguration) SetRam(v float32)`

SetRam sets Ram field to given value.


### GetGpu

`func (o *CreateServerConfiguration) GetGpu() float32`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *CreateServerConfiguration) GetGpuOk() (*float32, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *CreateServerConfiguration) SetGpu(v float32)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *CreateServerConfiguration) HasGpu() bool`

HasGpu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


