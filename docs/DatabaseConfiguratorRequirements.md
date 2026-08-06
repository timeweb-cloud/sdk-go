# DatabaseConfiguratorRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMin** | **float32** | Минимальное количество ядер процессора. | 
**CpuStep** | **float32** | Размер шага ядер процессора. | 
**CpuMax** | **float32** | Максимальное количество ядер процессора. | 
**RamMin** | **float32** | Минимальный объём оперативной памяти (в Мб). | 
**RamStep** | **float32** | Размер шага оперативной памяти (в Мб). | 
**RamMax** | **float32** | Максимальный объём оперативной памяти (в Мб). | 
**DiskMin** | **float32** | Минимальный размер диска (в Мб). | 
**DiskStep** | **float32** | Размер шага диска (в Мб). | 
**DiskMax** | **float32** | Максимальный размер диска (в Мб). | 
**NetworkBandwidthMin** | **float32** | Минимальная пропускная способность интернет-канала (в Мбит/с). | 
**NetworkBandwidthStep** | **float32** | Размер шага пропускной способности интернет-канала (в Мбит/с). | 
**NetworkBandwidthMax** | **float32** | Максимальная пропускная способность интернет-канала (в Мбит/с). | 

## Methods

### NewDatabaseConfiguratorRequirements

`func NewDatabaseConfiguratorRequirements(cpuMin float32, cpuStep float32, cpuMax float32, ramMin float32, ramStep float32, ramMax float32, diskMin float32, diskStep float32, diskMax float32, networkBandwidthMin float32, networkBandwidthStep float32, networkBandwidthMax float32, ) *DatabaseConfiguratorRequirements`

NewDatabaseConfiguratorRequirements instantiates a new DatabaseConfiguratorRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseConfiguratorRequirementsWithDefaults

`func NewDatabaseConfiguratorRequirementsWithDefaults() *DatabaseConfiguratorRequirements`

NewDatabaseConfiguratorRequirementsWithDefaults instantiates a new DatabaseConfiguratorRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMin

`func (o *DatabaseConfiguratorRequirements) GetCpuMin() float32`

GetCpuMin returns the CpuMin field if non-nil, zero value otherwise.

### GetCpuMinOk

`func (o *DatabaseConfiguratorRequirements) GetCpuMinOk() (*float32, bool)`

GetCpuMinOk returns a tuple with the CpuMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMin

`func (o *DatabaseConfiguratorRequirements) SetCpuMin(v float32)`

SetCpuMin sets CpuMin field to given value.


### GetCpuStep

`func (o *DatabaseConfiguratorRequirements) GetCpuStep() float32`

GetCpuStep returns the CpuStep field if non-nil, zero value otherwise.

### GetCpuStepOk

`func (o *DatabaseConfiguratorRequirements) GetCpuStepOk() (*float32, bool)`

GetCpuStepOk returns a tuple with the CpuStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuStep

`func (o *DatabaseConfiguratorRequirements) SetCpuStep(v float32)`

SetCpuStep sets CpuStep field to given value.


### GetCpuMax

`func (o *DatabaseConfiguratorRequirements) GetCpuMax() float32`

GetCpuMax returns the CpuMax field if non-nil, zero value otherwise.

### GetCpuMaxOk

`func (o *DatabaseConfiguratorRequirements) GetCpuMaxOk() (*float32, bool)`

GetCpuMaxOk returns a tuple with the CpuMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMax

`func (o *DatabaseConfiguratorRequirements) SetCpuMax(v float32)`

SetCpuMax sets CpuMax field to given value.


### GetRamMin

`func (o *DatabaseConfiguratorRequirements) GetRamMin() float32`

GetRamMin returns the RamMin field if non-nil, zero value otherwise.

### GetRamMinOk

`func (o *DatabaseConfiguratorRequirements) GetRamMinOk() (*float32, bool)`

GetRamMinOk returns a tuple with the RamMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamMin

`func (o *DatabaseConfiguratorRequirements) SetRamMin(v float32)`

SetRamMin sets RamMin field to given value.


### GetRamStep

`func (o *DatabaseConfiguratorRequirements) GetRamStep() float32`

GetRamStep returns the RamStep field if non-nil, zero value otherwise.

### GetRamStepOk

`func (o *DatabaseConfiguratorRequirements) GetRamStepOk() (*float32, bool)`

GetRamStepOk returns a tuple with the RamStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamStep

`func (o *DatabaseConfiguratorRequirements) SetRamStep(v float32)`

SetRamStep sets RamStep field to given value.


### GetRamMax

`func (o *DatabaseConfiguratorRequirements) GetRamMax() float32`

GetRamMax returns the RamMax field if non-nil, zero value otherwise.

### GetRamMaxOk

`func (o *DatabaseConfiguratorRequirements) GetRamMaxOk() (*float32, bool)`

GetRamMaxOk returns a tuple with the RamMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamMax

`func (o *DatabaseConfiguratorRequirements) SetRamMax(v float32)`

SetRamMax sets RamMax field to given value.


### GetDiskMin

`func (o *DatabaseConfiguratorRequirements) GetDiskMin() float32`

GetDiskMin returns the DiskMin field if non-nil, zero value otherwise.

### GetDiskMinOk

`func (o *DatabaseConfiguratorRequirements) GetDiskMinOk() (*float32, bool)`

GetDiskMinOk returns a tuple with the DiskMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMin

`func (o *DatabaseConfiguratorRequirements) SetDiskMin(v float32)`

SetDiskMin sets DiskMin field to given value.


### GetDiskStep

`func (o *DatabaseConfiguratorRequirements) GetDiskStep() float32`

GetDiskStep returns the DiskStep field if non-nil, zero value otherwise.

### GetDiskStepOk

`func (o *DatabaseConfiguratorRequirements) GetDiskStepOk() (*float32, bool)`

GetDiskStepOk returns a tuple with the DiskStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStep

`func (o *DatabaseConfiguratorRequirements) SetDiskStep(v float32)`

SetDiskStep sets DiskStep field to given value.


### GetDiskMax

`func (o *DatabaseConfiguratorRequirements) GetDiskMax() float32`

GetDiskMax returns the DiskMax field if non-nil, zero value otherwise.

### GetDiskMaxOk

`func (o *DatabaseConfiguratorRequirements) GetDiskMaxOk() (*float32, bool)`

GetDiskMaxOk returns a tuple with the DiskMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMax

`func (o *DatabaseConfiguratorRequirements) SetDiskMax(v float32)`

SetDiskMax sets DiskMax field to given value.


### GetNetworkBandwidthMin

`func (o *DatabaseConfiguratorRequirements) GetNetworkBandwidthMin() float32`

GetNetworkBandwidthMin returns the NetworkBandwidthMin field if non-nil, zero value otherwise.

### GetNetworkBandwidthMinOk

`func (o *DatabaseConfiguratorRequirements) GetNetworkBandwidthMinOk() (*float32, bool)`

GetNetworkBandwidthMinOk returns a tuple with the NetworkBandwidthMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidthMin

`func (o *DatabaseConfiguratorRequirements) SetNetworkBandwidthMin(v float32)`

SetNetworkBandwidthMin sets NetworkBandwidthMin field to given value.


### GetNetworkBandwidthStep

`func (o *DatabaseConfiguratorRequirements) GetNetworkBandwidthStep() float32`

GetNetworkBandwidthStep returns the NetworkBandwidthStep field if non-nil, zero value otherwise.

### GetNetworkBandwidthStepOk

`func (o *DatabaseConfiguratorRequirements) GetNetworkBandwidthStepOk() (*float32, bool)`

GetNetworkBandwidthStepOk returns a tuple with the NetworkBandwidthStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidthStep

`func (o *DatabaseConfiguratorRequirements) SetNetworkBandwidthStep(v float32)`

SetNetworkBandwidthStep sets NetworkBandwidthStep field to given value.


### GetNetworkBandwidthMax

`func (o *DatabaseConfiguratorRequirements) GetNetworkBandwidthMax() float32`

GetNetworkBandwidthMax returns the NetworkBandwidthMax field if non-nil, zero value otherwise.

### GetNetworkBandwidthMaxOk

`func (o *DatabaseConfiguratorRequirements) GetNetworkBandwidthMaxOk() (*float32, bool)`

GetNetworkBandwidthMaxOk returns a tuple with the NetworkBandwidthMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidthMax

`func (o *DatabaseConfiguratorRequirements) SetNetworkBandwidthMax(v float32)`

SetNetworkBandwidthMax sets NetworkBandwidthMax field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


