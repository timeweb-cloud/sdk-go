# CreateStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название хранилища. | 
**Type** | **string** | Тип хранилища. | 
**PresetId** | **float32** | ID тарифа. | 

## Methods

### NewCreateStorageRequest

`func NewCreateStorageRequest(name string, type_ string, presetId float32, ) *CreateStorageRequest`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


