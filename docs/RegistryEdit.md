# RegistryEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Новое описание реестра | [optional] 
**PresetId** | Pointer to **int32** | ID тарифа. Нельзя передавать вместе с &#x60;configuration&#x60; | [optional] 
**Configuration** | Pointer to [**RegistryInConfiguration**](RegistryInConfiguration.md) |  | [optional] 

## Methods

### NewRegistryEdit

`func NewRegistryEdit() *RegistryEdit`

NewRegistryEdit instantiates a new RegistryEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryEditWithDefaults

`func NewRegistryEditWithDefaults() *RegistryEdit`

NewRegistryEditWithDefaults instantiates a new RegistryEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RegistryEdit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegistryEdit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegistryEdit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegistryEdit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPresetId

`func (o *RegistryEdit) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *RegistryEdit) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *RegistryEdit) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *RegistryEdit) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetConfiguration

`func (o *RegistryEdit) GetConfiguration() RegistryInConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *RegistryEdit) GetConfigurationOk() (*RegistryInConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *RegistryEdit) SetConfiguration(v RegistryInConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *RegistryEdit) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


