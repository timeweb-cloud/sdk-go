# NodeOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID ноды | 
**CreatedAt** | **time.Time** | Дата и время создания ноды в формате ISO8601 | 
**Type** | **string** | Тип ноды | 
**GroupId** | **int32** | ID группы нод | 
**Status** | **string** | Статус | 
**PresetId** | **int32** | ID тарифа ноды | 
**Cpu** | **int32** | Количество ядер | 
**Ram** | **int32** | Количество памяти | 
**Disk** | **int32** | Количество пространства | 
**Network** | **int32** | Пропускная способность сети | 
**NodeIp** | **string** | Ip-адрес ноды | 

## Methods

### NewNodeOut

`func NewNodeOut(id int32, createdAt time.Time, type_ string, groupId int32, status string, presetId int32, cpu int32, ram int32, disk int32, network int32, nodeIp string, ) *NodeOut`

NewNodeOut instantiates a new NodeOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeOutWithDefaults

`func NewNodeOutWithDefaults() *NodeOut`

NewNodeOutWithDefaults instantiates a new NodeOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeOut) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeOut) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeOut) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *NodeOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NodeOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NodeOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetType

`func (o *NodeOut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeOut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeOut) SetType(v string)`

SetType sets Type field to given value.


### GetGroupId

`func (o *NodeOut) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NodeOut) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NodeOut) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetStatus

`func (o *NodeOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPresetId

`func (o *NodeOut) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *NodeOut) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *NodeOut) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetCpu

`func (o *NodeOut) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NodeOut) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NodeOut) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *NodeOut) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *NodeOut) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *NodeOut) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetDisk

`func (o *NodeOut) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *NodeOut) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *NodeOut) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetNetwork

`func (o *NodeOut) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NodeOut) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NodeOut) SetNetwork(v int32)`

SetNetwork sets Network field to given value.


### GetNodeIp

`func (o *NodeOut) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *NodeOut) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *NodeOut) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


