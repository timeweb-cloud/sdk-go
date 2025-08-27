# CreateStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название хранилища. | 
**Description** | Pointer to **string** | Комментарий к хранилищу. | [optional] 
**Type** | **string** | Тип хранилища. | 
**PresetId** | Pointer to **float32** | ID тарифа. Нельзя передавать вместе с &#x60;configurator&#x60;. | [optional] 
**Configurator** | Pointer to [**CreateStorageRequestConfigurator**](CreateStorageRequestConfigurator.md) |  | [optional] 
**ProjectId** | Pointer to **float32** | ID проекта. | [optional] 

## Methods

### NewCreateStorageRequest

`func NewCreateStorageRequest(name string, type_ string, ) *CreateStorageRequest`

NewCreateStorageRequest instantiates a new CreateStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageRequestWithDefaults

`func NewCreateStorageRequestWithDefaults() *CreateStorageRequest`

NewCreateStorageRequestWithDefaults instantiates a new CreateStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateStorageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStorageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStorageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateStorageRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStorageRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStorageRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateStorageRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateStorageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateStorageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateStorageRequest) SetType(v string)`

SetType sets Type field to given value.


### GetPresetId

`func (o *CreateStorageRequest) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreateStorageRequest) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreateStorageRequest) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *CreateStorageRequest) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetConfigurator

`func (o *CreateStorageRequest) GetConfigurator() CreateStorageRequestConfigurator`

GetConfigurator returns the Configurator field if non-nil, zero value otherwise.

### GetConfiguratorOk

`func (o *CreateStorageRequest) GetConfiguratorOk() (*CreateStorageRequestConfigurator, bool)`

GetConfiguratorOk returns a tuple with the Configurator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurator

`func (o *CreateStorageRequest) SetConfigurator(v CreateStorageRequestConfigurator)`

SetConfigurator sets Configurator field to given value.

### HasConfigurator

`func (o *CreateStorageRequest) HasConfigurator() bool`

HasConfigurator returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateStorageRequest) GetProjectId() float32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateStorageRequest) GetProjectIdOk() (*float32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateStorageRequest) SetProjectId(v float32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateStorageRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


