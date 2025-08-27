# RegistryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID реестра | 
**Name** | **string** | Название реестра | 
**Description** | **string** | Описание | 
**PresetId** | **int32** | ID тарифа | 
**ConfiguratorId** | **int32** | ID конфигуратора | 
**ProjectId** | **int32** | ID проекта | 
**CreatedAt** | **time.Time** | Дата и время создания реестра в формате ISO8601 | 
**UpdatedAt** | **time.Time** | Дата и время обновления реестра в формате ISO8601 | 
**DiskStats** | [**RegistryOutDiskStats**](RegistryOutDiskStats.md) |  | 

## Methods

### NewRegistryOut

`func NewRegistryOut(id int32, name string, description string, presetId int32, configuratorId int32, projectId int32, createdAt time.Time, updatedAt time.Time, diskStats RegistryOutDiskStats, ) *RegistryOut`

NewRegistryOut instantiates a new RegistryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryOutWithDefaults

`func NewRegistryOutWithDefaults() *RegistryOut`

NewRegistryOutWithDefaults instantiates a new RegistryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegistryOut) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryOut) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryOut) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *RegistryOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RegistryOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegistryOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegistryOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPresetId

`func (o *RegistryOut) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *RegistryOut) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *RegistryOut) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetConfiguratorId

`func (o *RegistryOut) GetConfiguratorId() int32`

GetConfiguratorId returns the ConfiguratorId field if non-nil, zero value otherwise.

### GetConfiguratorIdOk

`func (o *RegistryOut) GetConfiguratorIdOk() (*int32, bool)`

GetConfiguratorIdOk returns a tuple with the ConfiguratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguratorId

`func (o *RegistryOut) SetConfiguratorId(v int32)`

SetConfiguratorId sets ConfiguratorId field to given value.


### GetProjectId

`func (o *RegistryOut) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RegistryOut) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RegistryOut) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *RegistryOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RegistryOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RegistryOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RegistryOut) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RegistryOut) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RegistryOut) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDiskStats

`func (o *RegistryOut) GetDiskStats() RegistryOutDiskStats`

GetDiskStats returns the DiskStats field if non-nil, zero value otherwise.

### GetDiskStatsOk

`func (o *RegistryOut) GetDiskStatsOk() (*RegistryOutDiskStats, bool)`

GetDiskStatsOk returns a tuple with the DiskStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStats

`func (o *RegistryOut) SetDiskStats(v RegistryOutDiskStats)`

SetDiskStats sets DiskStats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


