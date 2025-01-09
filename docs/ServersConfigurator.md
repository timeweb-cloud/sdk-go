# ServersConfigurator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID конфигуратора сервера. | 
**Location** | **string** | Локация сервера. | 
**DiskType** | **string** | Тип диска. | 
**IsAllowedLocalNetwork** | **bool** | Есть возможность подключения локальной сети | 
**CpuFrequency** | **string** | Частота процессора. | 
**Requirements** | [**ServersConfiguratorRequirements**](ServersConfiguratorRequirements.md) |  | 

## Methods

### NewServersConfigurator

`func NewServersConfigurator(id float32, location string, diskType string, isAllowedLocalNetwork bool, cpuFrequency string, requirements ServersConfiguratorRequirements, ) *ServersConfigurator`

NewServersConfigurator instantiates a new ServersConfigurator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServersConfiguratorWithDefaults

`func NewServersConfiguratorWithDefaults() *ServersConfigurator`

NewServersConfiguratorWithDefaults instantiates a new ServersConfigurator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServersConfigurator) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServersConfigurator) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServersConfigurator) SetId(v float32)`

SetId sets Id field to given value.


### GetLocation

`func (o *ServersConfigurator) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ServersConfigurator) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ServersConfigurator) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetDiskType

`func (o *ServersConfigurator) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *ServersConfigurator) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *ServersConfigurator) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.


### GetIsAllowedLocalNetwork

`func (o *ServersConfigurator) GetIsAllowedLocalNetwork() bool`

GetIsAllowedLocalNetwork returns the IsAllowedLocalNetwork field if non-nil, zero value otherwise.

### GetIsAllowedLocalNetworkOk

`func (o *ServersConfigurator) GetIsAllowedLocalNetworkOk() (*bool, bool)`

GetIsAllowedLocalNetworkOk returns a tuple with the IsAllowedLocalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowedLocalNetwork

`func (o *ServersConfigurator) SetIsAllowedLocalNetwork(v bool)`

SetIsAllowedLocalNetwork sets IsAllowedLocalNetwork field to given value.


### GetCpuFrequency

`func (o *ServersConfigurator) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *ServersConfigurator) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *ServersConfigurator) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.


### GetRequirements

`func (o *ServersConfigurator) GetRequirements() ServersConfiguratorRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ServersConfigurator) GetRequirementsOk() (*ServersConfiguratorRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ServersConfigurator) SetRequirements(v ServersConfiguratorRequirements)`

SetRequirements sets Requirements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


