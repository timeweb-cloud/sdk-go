# RegistryIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя реестра. Должно быть уникальным, от 3 до 48 символов, начинаться и заканчиваться буквой или числом, содержать только латинские буквы в нижнем регистре, цифры и символ «-», без пробелов | 
**Description** | Pointer to **string** | Описание реестра | [optional] 
**PresetId** | Pointer to **int32** | ID тарифа. Нельзя передавать вместе с &#x60;configuration&#x60; | [optional] 
**Configuration** | Pointer to [**RegistryInConfiguration**](RegistryInConfiguration.md) |  | [optional] 
**ProjectId** | Pointer to **int32** | ID проекта | [optional] 

## Methods

### NewRegistryIn

`func NewRegistryIn(name string, ) *RegistryIn`

NewRegistryIn instantiates a new RegistryIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryInWithDefaults

`func NewRegistryInWithDefaults() *RegistryIn`

NewRegistryInWithDefaults instantiates a new RegistryIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegistryIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryIn) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RegistryIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegistryIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegistryIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegistryIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPresetId

`func (o *RegistryIn) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *RegistryIn) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *RegistryIn) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *RegistryIn) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetConfiguration

`func (o *RegistryIn) GetConfiguration() RegistryInConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *RegistryIn) GetConfigurationOk() (*RegistryInConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *RegistryIn) SetConfiguration(v RegistryInConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *RegistryIn) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetProjectId

`func (o *RegistryIn) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RegistryIn) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RegistryIn) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RegistryIn) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


