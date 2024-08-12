# AppConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **float32** | Количество ядер процессора. | [optional] 
**Ram** | Pointer to **float32** | Объем оперативной памяти (в МБ). | [optional] 
**NetworkBandwidth** | Pointer to **float32** | Скорость сетевого канала (в Мбит/с). | [optional] 
**CpuFrequency** | Pointer to **string** | Частота процессора (в ГГц). | [optional] 
**DiskType** | Pointer to **string** | Тип диска. | [optional] 

## Methods

### NewAppConfiguration

`func NewAppConfiguration() *AppConfiguration`

NewAppConfiguration instantiates a new AppConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConfigurationWithDefaults

`func NewAppConfigurationWithDefaults() *AppConfiguration`

NewAppConfigurationWithDefaults instantiates a new AppConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *AppConfiguration) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *AppConfiguration) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *AppConfiguration) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *AppConfiguration) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *AppConfiguration) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *AppConfiguration) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *AppConfiguration) SetRam(v float32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *AppConfiguration) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetNetworkBandwidth

`func (o *AppConfiguration) GetNetworkBandwidth() float32`

GetNetworkBandwidth returns the NetworkBandwidth field if non-nil, zero value otherwise.

### GetNetworkBandwidthOk

`func (o *AppConfiguration) GetNetworkBandwidthOk() (*float32, bool)`

GetNetworkBandwidthOk returns a tuple with the NetworkBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidth

`func (o *AppConfiguration) SetNetworkBandwidth(v float32)`

SetNetworkBandwidth sets NetworkBandwidth field to given value.

### HasNetworkBandwidth

`func (o *AppConfiguration) HasNetworkBandwidth() bool`

HasNetworkBandwidth returns a boolean if a field has been set.

### GetCpuFrequency

`func (o *AppConfiguration) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *AppConfiguration) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *AppConfiguration) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.

### HasCpuFrequency

`func (o *AppConfiguration) HasCpuFrequency() bool`

HasCpuFrequency returns a boolean if a field has been set.

### GetDiskType

`func (o *AppConfiguration) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *AppConfiguration) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *AppConfiguration) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *AppConfiguration) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


