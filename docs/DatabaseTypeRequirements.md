# DatabaseTypeRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMin** | Pointer to **float32** | Минимальное количество CPU. | [optional] 
**RamMin** | Pointer to **float32** | Минимальный объем оперативной памяти. | [optional] 
**DiskMin** | Pointer to **float32** | Минимальный объем дискового пространства. | [optional] 

## Methods

### NewDatabaseTypeRequirements

`func NewDatabaseTypeRequirements() *DatabaseTypeRequirements`

NewDatabaseTypeRequirements instantiates a new DatabaseTypeRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseTypeRequirementsWithDefaults

`func NewDatabaseTypeRequirementsWithDefaults() *DatabaseTypeRequirements`

NewDatabaseTypeRequirementsWithDefaults instantiates a new DatabaseTypeRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMin

`func (o *DatabaseTypeRequirements) GetCpuMin() float32`

GetCpuMin returns the CpuMin field if non-nil, zero value otherwise.

### GetCpuMinOk

`func (o *DatabaseTypeRequirements) GetCpuMinOk() (*float32, bool)`

GetCpuMinOk returns a tuple with the CpuMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMin

`func (o *DatabaseTypeRequirements) SetCpuMin(v float32)`

SetCpuMin sets CpuMin field to given value.

### HasCpuMin

`func (o *DatabaseTypeRequirements) HasCpuMin() bool`

HasCpuMin returns a boolean if a field has been set.

### GetRamMin

`func (o *DatabaseTypeRequirements) GetRamMin() float32`

GetRamMin returns the RamMin field if non-nil, zero value otherwise.

### GetRamMinOk

`func (o *DatabaseTypeRequirements) GetRamMinOk() (*float32, bool)`

GetRamMinOk returns a tuple with the RamMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamMin

`func (o *DatabaseTypeRequirements) SetRamMin(v float32)`

SetRamMin sets RamMin field to given value.

### HasRamMin

`func (o *DatabaseTypeRequirements) HasRamMin() bool`

HasRamMin returns a boolean if a field has been set.

### GetDiskMin

`func (o *DatabaseTypeRequirements) GetDiskMin() float32`

GetDiskMin returns the DiskMin field if non-nil, zero value otherwise.

### GetDiskMinOk

`func (o *DatabaseTypeRequirements) GetDiskMinOk() (*float32, bool)`

GetDiskMinOk returns a tuple with the DiskMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMin

`func (o *DatabaseTypeRequirements) SetDiskMin(v float32)`

SetDiskMin sets DiskMin field to given value.

### HasDiskMin

`func (o *DatabaseTypeRequirements) HasDiskMin() bool`

HasDiskMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


