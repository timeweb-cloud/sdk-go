# ServersConfiguratorRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMin** | **float32** | Минимальное количество ядер процессора. | 
**CpuStep** | **float32** | Размер шага ядер процессора. | 
**CpuMax** | **float32** | Максимальное количество ядер процессора. | 
**RamMin** | **float32** | Минимальное количество оперативной памяти (в Мб). | 
**RamStep** | **float32** | Размер шага оперативной памяти. | 
**RamMax** | **float32** | Максимальное количество оперативной памяти (в Мб). | 
**DiskMin** | **float32** | Минимальный размер диска (в Мб). | 
**DiskStep** | **float32** | Размер шага диска | 
**DiskMax** | **float32** | Максимальный размер диска (в Мб). | 
**NetworkBandwidthMin** | **float32** | Минимальныая пропускная способноть интернет-канала (в Мб) | 
**NetworkBandwidthStep** | **float32** | Размер шага пропускной способноти интернет-канала (в Мб) | 
**NetworkBandwidthMax** | **float32** | Максимальная пропускная способноть интернет-канала (в Мб) | 
**GpuMin** | **NullableFloat32** | Минимальное количество видеокарт | 
**GpuMax** | **NullableFloat32** | Максимальное количество видеокарт | 
**GpuStep** | **NullableFloat32** | Размер шага видеокарт | 

## Methods

### NewServersConfiguratorRequirements

`func NewServersConfiguratorRequirements(cpuMin float32, cpuStep float32, cpuMax float32, ramMin float32, ramStep float32, ramMax float32, diskMin float32, diskStep float32, diskMax float32, networkBandwidthMin float32, networkBandwidthStep float32, networkBandwidthMax float32, gpuMin NullableFloat32, gpuMax NullableFloat32, gpuStep NullableFloat32, ) *ServersConfiguratorRequirements`

NewServersConfiguratorRequirements instantiates a new ServersConfiguratorRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServersConfiguratorRequirementsWithDefaults

`func NewServersConfiguratorRequirementsWithDefaults() *ServersConfiguratorRequirements`

NewServersConfiguratorRequirementsWithDefaults instantiates a new ServersConfiguratorRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMin

`func (o *ServersConfiguratorRequirements) GetCpuMin() float32`

GetCpuMin returns the CpuMin field if non-nil, zero value otherwise.

### GetCpuMinOk

`func (o *ServersConfiguratorRequirements) GetCpuMinOk() (*float32, bool)`

GetCpuMinOk returns a tuple with the CpuMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMin

`func (o *ServersConfiguratorRequirements) SetCpuMin(v float32)`

SetCpuMin sets CpuMin field to given value.


### GetCpuStep

`func (o *ServersConfiguratorRequirements) GetCpuStep() float32`

GetCpuStep returns the CpuStep field if non-nil, zero value otherwise.

### GetCpuStepOk

`func (o *ServersConfiguratorRequirements) GetCpuStepOk() (*float32, bool)`

GetCpuStepOk returns a tuple with the CpuStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuStep

`func (o *ServersConfiguratorRequirements) SetCpuStep(v float32)`

SetCpuStep sets CpuStep field to given value.


### GetCpuMax

`func (o *ServersConfiguratorRequirements) GetCpuMax() float32`

GetCpuMax returns the CpuMax field if non-nil, zero value otherwise.

### GetCpuMaxOk

`func (o *ServersConfiguratorRequirements) GetCpuMaxOk() (*float32, bool)`

GetCpuMaxOk returns a tuple with the CpuMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMax

`func (o *ServersConfiguratorRequirements) SetCpuMax(v float32)`

SetCpuMax sets CpuMax field to given value.


### GetRamMin

`func (o *ServersConfiguratorRequirements) GetRamMin() float32`

GetRamMin returns the RamMin field if non-nil, zero value otherwise.

### GetRamMinOk

`func (o *ServersConfiguratorRequirements) GetRamMinOk() (*float32, bool)`

GetRamMinOk returns a tuple with the RamMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamMin

`func (o *ServersConfiguratorRequirements) SetRamMin(v float32)`

SetRamMin sets RamMin field to given value.


### GetRamStep

`func (o *ServersConfiguratorRequirements) GetRamStep() float32`

GetRamStep returns the RamStep field if non-nil, zero value otherwise.

### GetRamStepOk

`func (o *ServersConfiguratorRequirements) GetRamStepOk() (*float32, bool)`

GetRamStepOk returns a tuple with the RamStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamStep

`func (o *ServersConfiguratorRequirements) SetRamStep(v float32)`

SetRamStep sets RamStep field to given value.


### GetRamMax

`func (o *ServersConfiguratorRequirements) GetRamMax() float32`

GetRamMax returns the RamMax field if non-nil, zero value otherwise.

### GetRamMaxOk

`func (o *ServersConfiguratorRequirements) GetRamMaxOk() (*float32, bool)`

GetRamMaxOk returns a tuple with the RamMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamMax

`func (o *ServersConfiguratorRequirements) SetRamMax(v float32)`

SetRamMax sets RamMax field to given value.


### GetDiskMin

`func (o *ServersConfiguratorRequirements) GetDiskMin() float32`

GetDiskMin returns the DiskMin field if non-nil, zero value otherwise.

### GetDiskMinOk

`func (o *ServersConfiguratorRequirements) GetDiskMinOk() (*float32, bool)`

GetDiskMinOk returns a tuple with the DiskMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMin

`func (o *ServersConfiguratorRequirements) SetDiskMin(v float32)`

SetDiskMin sets DiskMin field to given value.


### GetDiskStep

`func (o *ServersConfiguratorRequirements) GetDiskStep() float32`

GetDiskStep returns the DiskStep field if non-nil, zero value otherwise.

### GetDiskStepOk

`func (o *ServersConfiguratorRequirements) GetDiskStepOk() (*float32, bool)`

GetDiskStepOk returns a tuple with the DiskStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStep

`func (o *ServersConfiguratorRequirements) SetDiskStep(v float32)`

SetDiskStep sets DiskStep field to given value.


### GetDiskMax

`func (o *ServersConfiguratorRequirements) GetDiskMax() float32`

GetDiskMax returns the DiskMax field if non-nil, zero value otherwise.

### GetDiskMaxOk

`func (o *ServersConfiguratorRequirements) GetDiskMaxOk() (*float32, bool)`

GetDiskMaxOk returns a tuple with the DiskMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMax

`func (o *ServersConfiguratorRequirements) SetDiskMax(v float32)`

SetDiskMax sets DiskMax field to given value.


### GetNetworkBandwidthMin

`func (o *ServersConfiguratorRequirements) GetNetworkBandwidthMin() float32`

GetNetworkBandwidthMin returns the NetworkBandwidthMin field if non-nil, zero value otherwise.

### GetNetworkBandwidthMinOk

`func (o *ServersConfiguratorRequirements) GetNetworkBandwidthMinOk() (*float32, bool)`

GetNetworkBandwidthMinOk returns a tuple with the NetworkBandwidthMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidthMin

`func (o *ServersConfiguratorRequirements) SetNetworkBandwidthMin(v float32)`

SetNetworkBandwidthMin sets NetworkBandwidthMin field to given value.


### GetNetworkBandwidthStep

`func (o *ServersConfiguratorRequirements) GetNetworkBandwidthStep() float32`

GetNetworkBandwidthStep returns the NetworkBandwidthStep field if non-nil, zero value otherwise.

### GetNetworkBandwidthStepOk

`func (o *ServersConfiguratorRequirements) GetNetworkBandwidthStepOk() (*float32, bool)`

GetNetworkBandwidthStepOk returns a tuple with the NetworkBandwidthStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidthStep

`func (o *ServersConfiguratorRequirements) SetNetworkBandwidthStep(v float32)`

SetNetworkBandwidthStep sets NetworkBandwidthStep field to given value.


### GetNetworkBandwidthMax

`func (o *ServersConfiguratorRequirements) GetNetworkBandwidthMax() float32`

GetNetworkBandwidthMax returns the NetworkBandwidthMax field if non-nil, zero value otherwise.

### GetNetworkBandwidthMaxOk

`func (o *ServersConfiguratorRequirements) GetNetworkBandwidthMaxOk() (*float32, bool)`

GetNetworkBandwidthMaxOk returns a tuple with the NetworkBandwidthMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidthMax

`func (o *ServersConfiguratorRequirements) SetNetworkBandwidthMax(v float32)`

SetNetworkBandwidthMax sets NetworkBandwidthMax field to given value.


### GetGpuMin

`func (o *ServersConfiguratorRequirements) GetGpuMin() float32`

GetGpuMin returns the GpuMin field if non-nil, zero value otherwise.

### GetGpuMinOk

`func (o *ServersConfiguratorRequirements) GetGpuMinOk() (*float32, bool)`

GetGpuMinOk returns a tuple with the GpuMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuMin

`func (o *ServersConfiguratorRequirements) SetGpuMin(v float32)`

SetGpuMin sets GpuMin field to given value.


### SetGpuMinNil

`func (o *ServersConfiguratorRequirements) SetGpuMinNil(b bool)`

 SetGpuMinNil sets the value for GpuMin to be an explicit nil

### UnsetGpuMin
`func (o *ServersConfiguratorRequirements) UnsetGpuMin()`

UnsetGpuMin ensures that no value is present for GpuMin, not even an explicit nil
### GetGpuMax

`func (o *ServersConfiguratorRequirements) GetGpuMax() float32`

GetGpuMax returns the GpuMax field if non-nil, zero value otherwise.

### GetGpuMaxOk

`func (o *ServersConfiguratorRequirements) GetGpuMaxOk() (*float32, bool)`

GetGpuMaxOk returns a tuple with the GpuMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuMax

`func (o *ServersConfiguratorRequirements) SetGpuMax(v float32)`

SetGpuMax sets GpuMax field to given value.


### SetGpuMaxNil

`func (o *ServersConfiguratorRequirements) SetGpuMaxNil(b bool)`

 SetGpuMaxNil sets the value for GpuMax to be an explicit nil

### UnsetGpuMax
`func (o *ServersConfiguratorRequirements) UnsetGpuMax()`

UnsetGpuMax ensures that no value is present for GpuMax, not even an explicit nil
### GetGpuStep

`func (o *ServersConfiguratorRequirements) GetGpuStep() float32`

GetGpuStep returns the GpuStep field if non-nil, zero value otherwise.

### GetGpuStepOk

`func (o *ServersConfiguratorRequirements) GetGpuStepOk() (*float32, bool)`

GetGpuStepOk returns a tuple with the GpuStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuStep

`func (o *ServersConfiguratorRequirements) SetGpuStep(v float32)`

SetGpuStep sets GpuStep field to given value.


### SetGpuStepNil

`func (o *ServersConfiguratorRequirements) SetGpuStepNil(b bool)`

 SetGpuStepNil sets the value for GpuStep to be an explicit nil

### UnsetGpuStep
`func (o *ServersConfiguratorRequirements) UnsetGpuStep()`

UnsetGpuStep ensures that no value is present for GpuStep, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


