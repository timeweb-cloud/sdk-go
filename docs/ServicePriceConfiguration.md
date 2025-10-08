# ServicePriceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **float32** | Количество ядер процессора | [optional] 
**Ram** | Pointer to **float32** | Объем оперативной памяти в МБ | [optional] 
**Disk** | Pointer to **float32** | Объем дискового пространства в ГБ | [optional] 

## Methods

### NewServicePriceConfiguration

`func NewServicePriceConfiguration() *ServicePriceConfiguration`

NewServicePriceConfiguration instantiates a new ServicePriceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePriceConfigurationWithDefaults

`func NewServicePriceConfigurationWithDefaults() *ServicePriceConfiguration`

NewServicePriceConfigurationWithDefaults instantiates a new ServicePriceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ServicePriceConfiguration) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ServicePriceConfiguration) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ServicePriceConfiguration) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ServicePriceConfiguration) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *ServicePriceConfiguration) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ServicePriceConfiguration) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ServicePriceConfiguration) SetRam(v float32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *ServicePriceConfiguration) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetDisk

`func (o *ServicePriceConfiguration) GetDisk() float32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ServicePriceConfiguration) GetDiskOk() (*float32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ServicePriceConfiguration) SetDisk(v float32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *ServicePriceConfiguration) HasDisk() bool`

HasDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


