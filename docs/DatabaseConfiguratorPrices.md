# DatabaseConfiguratorPrices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **float32** | Стоимость одного ядра процессора. | 
**Ram** | **float32** | Стоимость одной единицы оперативной памяти. | 
**Disk** | **float32** | Стоимость одной единицы дискового пространства. | 
**Bandwidth** | **NullableFloat32** | Стоимость одной единицы пропускной способности интернет-канала. | 

## Methods

### NewDatabaseConfiguratorPrices

`func NewDatabaseConfiguratorPrices(cpu float32, ram float32, disk float32, bandwidth NullableFloat32, ) *DatabaseConfiguratorPrices`

NewDatabaseConfiguratorPrices instantiates a new DatabaseConfiguratorPrices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseConfiguratorPricesWithDefaults

`func NewDatabaseConfiguratorPricesWithDefaults() *DatabaseConfiguratorPrices`

NewDatabaseConfiguratorPricesWithDefaults instantiates a new DatabaseConfiguratorPrices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *DatabaseConfiguratorPrices) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DatabaseConfiguratorPrices) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DatabaseConfiguratorPrices) SetCpu(v float32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *DatabaseConfiguratorPrices) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *DatabaseConfiguratorPrices) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *DatabaseConfiguratorPrices) SetRam(v float32)`

SetRam sets Ram field to given value.


### GetDisk

`func (o *DatabaseConfiguratorPrices) GetDisk() float32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *DatabaseConfiguratorPrices) GetDiskOk() (*float32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *DatabaseConfiguratorPrices) SetDisk(v float32)`

SetDisk sets Disk field to given value.


### GetBandwidth

`func (o *DatabaseConfiguratorPrices) GetBandwidth() float32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *DatabaseConfiguratorPrices) GetBandwidthOk() (*float32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *DatabaseConfiguratorPrices) SetBandwidth(v float32)`

SetBandwidth sets Bandwidth field to given value.


### SetBandwidthNil

`func (o *DatabaseConfiguratorPrices) SetBandwidthNil(b bool)`

 SetBandwidthNil sets the value for Bandwidth to be an explicit nil

### UnsetBandwidth
`func (o *DatabaseConfiguratorPrices) UnsetBandwidth()`

UnsetBandwidth ensures that no value is present for Bandwidth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


