# CreateNetworkDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название сетевого диска. | 
**Comment** | Pointer to **NullableString** | Комментарий | [optional] 
**Size** | **float32** | Размер диска в Гб | 
**PresetId** | **int32** | ID сетевого диска. | 

## Methods

### NewCreateNetworkDrive

`func NewCreateNetworkDrive(name string, size float32, presetId int32, ) *CreateNetworkDrive`

NewCreateNetworkDrive instantiates a new CreateNetworkDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDriveWithDefaults

`func NewCreateNetworkDriveWithDefaults() *CreateNetworkDrive`

NewCreateNetworkDriveWithDefaults instantiates a new CreateNetworkDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkDrive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkDrive) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkDrive) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *CreateNetworkDrive) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateNetworkDrive) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateNetworkDrive) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateNetworkDrive) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *CreateNetworkDrive) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *CreateNetworkDrive) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetSize

`func (o *CreateNetworkDrive) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateNetworkDrive) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateNetworkDrive) SetSize(v float32)`

SetSize sets Size field to given value.


### GetPresetId

`func (o *CreateNetworkDrive) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreateNetworkDrive) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreateNetworkDrive) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


