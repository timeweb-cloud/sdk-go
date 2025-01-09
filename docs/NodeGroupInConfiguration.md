# NodeGroupInConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguratorId** | **int32** | ID конфигуратора кластера | 
**Disk** | **int32** | Размер диска в МБ | 
**Cpu** | **int32** | Количество ядер процессора | 
**Ram** | **int32** | Размер ОЗУ сервера в МБ | 

## Methods

### NewNodeGroupInConfiguration

`func NewNodeGroupInConfiguration(configuratorId int32, disk int32, cpu int32, ram int32, ) *NodeGroupInConfiguration`

NewNodeGroupInConfiguration instantiates a new NodeGroupInConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeGroupInConfigurationWithDefaults

`func NewNodeGroupInConfigurationWithDefaults() *NodeGroupInConfiguration`

NewNodeGroupInConfigurationWithDefaults instantiates a new NodeGroupInConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguratorId

`func (o *NodeGroupInConfiguration) GetConfiguratorId() int32`

GetConfiguratorId returns the ConfiguratorId field if non-nil, zero value otherwise.

### GetConfiguratorIdOk

`func (o *NodeGroupInConfiguration) GetConfiguratorIdOk() (*int32, bool)`

GetConfiguratorIdOk returns a tuple with the ConfiguratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguratorId

`func (o *NodeGroupInConfiguration) SetConfiguratorId(v int32)`

SetConfiguratorId sets ConfiguratorId field to given value.


### GetDisk

`func (o *NodeGroupInConfiguration) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *NodeGroupInConfiguration) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *NodeGroupInConfiguration) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetCpu

`func (o *NodeGroupInConfiguration) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NodeGroupInConfiguration) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NodeGroupInConfiguration) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *NodeGroupInConfiguration) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *NodeGroupInConfiguration) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *NodeGroupInConfiguration) SetRam(v int32)`

SetRam sets Ram field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


