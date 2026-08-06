# DatabaseConfigurator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID конфигуратора. Передаётся при создании кластера в поле &#x60;configurator_id&#x60;. | 
**DiskType** | **string** | Тип диска. | 
**CpuFrequency** | **string** | Частота процессора (в ГГц). | 
**IsAllowedLocalNetwork** | **bool** | Есть возможность подключения локальной сети. | 
**Location** | **string** | Географическое расположение конфигуратора. | 
**Requirements** | [**DatabaseConfiguratorRequirements**](DatabaseConfiguratorRequirements.md) |  | 
**Prices** | Pointer to [**DatabaseConfiguratorPrices**](DatabaseConfiguratorPrices.md) |  | [optional] 
**Tags** | **[]string** | Теги конфигуратора, в том числе тег группы, в пределах которой доступна смена конфигурации кластера. | 

## Methods

### NewDatabaseConfigurator

`func NewDatabaseConfigurator(id float32, diskType string, cpuFrequency string, isAllowedLocalNetwork bool, location string, requirements DatabaseConfiguratorRequirements, tags []string, ) *DatabaseConfigurator`

NewDatabaseConfigurator instantiates a new DatabaseConfigurator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseConfiguratorWithDefaults

`func NewDatabaseConfiguratorWithDefaults() *DatabaseConfigurator`

NewDatabaseConfiguratorWithDefaults instantiates a new DatabaseConfigurator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseConfigurator) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseConfigurator) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseConfigurator) SetId(v float32)`

SetId sets Id field to given value.


### GetDiskType

`func (o *DatabaseConfigurator) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *DatabaseConfigurator) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *DatabaseConfigurator) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.


### GetCpuFrequency

`func (o *DatabaseConfigurator) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *DatabaseConfigurator) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *DatabaseConfigurator) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.


### GetIsAllowedLocalNetwork

`func (o *DatabaseConfigurator) GetIsAllowedLocalNetwork() bool`

GetIsAllowedLocalNetwork returns the IsAllowedLocalNetwork field if non-nil, zero value otherwise.

### GetIsAllowedLocalNetworkOk

`func (o *DatabaseConfigurator) GetIsAllowedLocalNetworkOk() (*bool, bool)`

GetIsAllowedLocalNetworkOk returns a tuple with the IsAllowedLocalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowedLocalNetwork

`func (o *DatabaseConfigurator) SetIsAllowedLocalNetwork(v bool)`

SetIsAllowedLocalNetwork sets IsAllowedLocalNetwork field to given value.


### GetLocation

`func (o *DatabaseConfigurator) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DatabaseConfigurator) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DatabaseConfigurator) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetRequirements

`func (o *DatabaseConfigurator) GetRequirements() DatabaseConfiguratorRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *DatabaseConfigurator) GetRequirementsOk() (*DatabaseConfiguratorRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *DatabaseConfigurator) SetRequirements(v DatabaseConfiguratorRequirements)`

SetRequirements sets Requirements field to given value.


### GetPrices

`func (o *DatabaseConfigurator) GetPrices() DatabaseConfiguratorPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *DatabaseConfigurator) GetPricesOk() (*DatabaseConfiguratorPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *DatabaseConfigurator) SetPrices(v DatabaseConfiguratorPrices)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *DatabaseConfigurator) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetTags

`func (o *DatabaseConfigurator) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DatabaseConfigurator) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DatabaseConfigurator) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


