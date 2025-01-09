# CreateProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Удобочитаемое имя проекта. Максимальная длина — 255 символов. | 
**Description** | Pointer to **NullableString** | Описание проекта. Максимальная длина — 255 символов. | [optional] 
**AvatarId** | Pointer to **NullableString** | ID аватара пользователя. Описание методов работы с аватарами появится позднее. | [optional] 

## Methods

### NewCreateProject

`func NewCreateProject(name string, ) *CreateProject`

NewCreateProject instantiates a new CreateProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectWithDefaults

`func NewCreateProjectWithDefaults() *CreateProject`

NewCreateProjectWithDefaults instantiates a new CreateProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProject) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateProject) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateProject) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAvatarId

`func (o *CreateProject) GetAvatarId() string`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *CreateProject) GetAvatarIdOk() (*string, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *CreateProject) SetAvatarId(v string)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *CreateProject) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### SetAvatarIdNil

`func (o *CreateProject) SetAvatarIdNil(b bool)`

 SetAvatarIdNil sets the value for AvatarId to be an explicit nil

### UnsetAvatarId
`func (o *CreateProject) UnsetAvatarId()`

UnsetAvatarId ensures that no value is present for AvatarId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


