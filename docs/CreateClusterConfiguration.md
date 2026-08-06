# CreateClusterConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguratorId** | **int32** | ID конфигуратора | 
**Disk** | **int32** | Размер диска (в Мб) | 
**Cpu** | **int32** | Количество ядер процессора | 
**Ram** | **int32** | Объем оперативной памяти (в Мб) | 

## Methods

### NewCreateClusterConfiguration

`func NewCreateClusterConfiguration(configuratorId int32, disk int32, cpu int32, ram int32, ) *CreateClusterConfiguration`

NewCreateClusterConfiguration instantiates a new CreateClusterConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterConfigurationWithDefaults

`func NewCreateClusterConfigurationWithDefaults() *CreateClusterConfiguration`

NewCreateClusterConfigurationWithDefaults instantiates a new CreateClusterConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguratorId

`func (o *CreateClusterConfiguration) GetConfiguratorId() int32`

GetConfiguratorId returns the ConfiguratorId field if non-nil, zero value otherwise.

### GetConfiguratorIdOk

`func (o *CreateClusterConfiguration) GetConfiguratorIdOk() (*int32, bool)`

GetConfiguratorIdOk returns a tuple with the ConfiguratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguratorId

`func (o *CreateClusterConfiguration) SetConfiguratorId(v int32)`

SetConfiguratorId sets ConfiguratorId field to given value.


### GetDisk

`func (o *CreateClusterConfiguration) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *CreateClusterConfiguration) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *CreateClusterConfiguration) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetCpu

`func (o *CreateClusterConfiguration) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *CreateClusterConfiguration) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *CreateClusterConfiguration) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *CreateClusterConfiguration) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *CreateClusterConfiguration) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *CreateClusterConfiguration) SetRam(v int32)`

SetRam sets Ram field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


