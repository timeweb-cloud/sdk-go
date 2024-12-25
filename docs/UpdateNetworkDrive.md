# UpdateNetworkDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название сетевого диска. | [optional] 
**Comment** | Pointer to **NullableString** | Комментарий | [optional] 
**Size** | Pointer to **float32** | Размер диска в Гб | [optional] 

## Methods

### NewUpdateNetworkDrive

`func NewUpdateNetworkDrive() *UpdateNetworkDrive`

NewUpdateNetworkDrive instantiates a new UpdateNetworkDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkDriveWithDefaults

`func NewUpdateNetworkDriveWithDefaults() *UpdateNetworkDrive`

NewUpdateNetworkDriveWithDefaults instantiates a new UpdateNetworkDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkDrive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkDrive) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkDrive) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkDrive) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *UpdateNetworkDrive) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateNetworkDrive) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateNetworkDrive) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateNetworkDrive) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *UpdateNetworkDrive) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *UpdateNetworkDrive) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetSize

`func (o *UpdateNetworkDrive) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UpdateNetworkDrive) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UpdateNetworkDrive) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *UpdateNetworkDrive) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


