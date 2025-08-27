# ClusterInConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguratorId** | **int32** | ID конфигуратора мастер-ноды | 
**Disk** | **int32** | Размер диска в МБ | 
**Cpu** | **int32** | Количество ядер процессора | 
**Ram** | **int32** | Размер ОЗУ сервера в МБ | 

## Methods

### NewClusterInConfiguration

`func NewClusterInConfiguration(configuratorId int32, disk int32, cpu int32, ram int32, ) *ClusterInConfiguration`

NewClusterInConfiguration instantiates a new ClusterInConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInConfigurationWithDefaults

`func NewClusterInConfigurationWithDefaults() *ClusterInConfiguration`

NewClusterInConfigurationWithDefaults instantiates a new ClusterInConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguratorId

`func (o *ClusterInConfiguration) GetConfiguratorId() int32`

GetConfiguratorId returns the ConfiguratorId field if non-nil, zero value otherwise.

### GetConfiguratorIdOk

`func (o *ClusterInConfiguration) GetConfiguratorIdOk() (*int32, bool)`

GetConfiguratorIdOk returns a tuple with the ConfiguratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguratorId

`func (o *ClusterInConfiguration) SetConfiguratorId(v int32)`

SetConfiguratorId sets ConfiguratorId field to given value.


### GetDisk

`func (o *ClusterInConfiguration) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ClusterInConfiguration) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ClusterInConfiguration) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetCpu

`func (o *ClusterInConfiguration) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ClusterInConfiguration) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ClusterInConfiguration) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *ClusterInConfiguration) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ClusterInConfiguration) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ClusterInConfiguration) SetRam(v int32)`

SetRam sets Ram field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


