# UpdateProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Удобочитаемое имя проекта. Максимальная длина — 255 символов. | [optional] 
**Description** | Pointer to **NullableString** | Описание проекта. Максимальная длина — 255 символов. | [optional] 
**AvatarId** | Pointer to **NullableString** | ID аватара пользователя. Описание методов работы с аватарами появится позднее. | [optional] 

## Methods

### NewUpdateProject

`func NewUpdateProject() *UpdateProject`

NewUpdateProject instantiates a new UpdateProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectWithDefaults

`func NewUpdateProjectWithDefaults() *UpdateProject`

NewUpdateProjectWithDefaults instantiates a new UpdateProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateProject) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateProject) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAvatarId

`func (o *UpdateProject) GetAvatarId() string`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *UpdateProject) GetAvatarIdOk() (*string, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *UpdateProject) SetAvatarId(v string)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *UpdateProject) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### SetAvatarIdNil

`func (o *UpdateProject) SetAvatarIdNil(b bool)`

 SetAvatarIdNil sets the value for AvatarId to be an explicit nil

### UnsetAvatarId
`func (o *UpdateProject) UnsetAvatarId()`

UnsetAvatarId ensures that no value is present for AvatarId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


