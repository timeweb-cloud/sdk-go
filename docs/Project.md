# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого проекта. Автоматически генерируется при создании. | 
**AccountId** | **string** | ID пользователя. | 
**AvatarId** | **NullableString** | ID аватара пользователя. Описание методов работы с аватарами появится позднее. | 
**Description** | **string** | Описание проекта. | 
**Name** | **string** | Удобочитаемое имя проекта. | 
**IsDefault** | **bool** | Это логическое значение, которое показывает, выбран ли проект по умолчанию для создания новых сущностей. | 

## Methods

### NewProject

`func NewProject(id float32, accountId string, avatarId NullableString, description string, name string, isDefault bool, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Project) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v float32)`

SetId sets Id field to given value.


### GetAccountId

`func (o *Project) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Project) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Project) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAvatarId

`func (o *Project) GetAvatarId() string`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *Project) GetAvatarIdOk() (*string, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *Project) SetAvatarId(v string)`

SetAvatarId sets AvatarId field to given value.


### SetAvatarIdNil

`func (o *Project) SetAvatarIdNil(b bool)`

 SetAvatarIdNil sets the value for AvatarId to be an explicit nil

### UnsetAvatarId
`func (o *Project) UnsetAvatarId()`

UnsetAvatarId ensures that no value is present for AvatarId, not even an explicit nil
### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *Project) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Project) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Project) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


