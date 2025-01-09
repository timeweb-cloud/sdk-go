# ServerBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID бэкапа сервера. | 
**Name** | **string** | Название бэкапа. | 
**Comment** | **NullableString** | Комментарий к бэкапу. | 
**CreatedAt** | **string** | Дата создания бэкапа. | 
**Status** | **string** | Статус бэкапа. | 
**Size** | **float32** | Размер бэкапа (в Мб). | 
**Type** | **string** | Тип бэкапа. | 
**Progress** | **float32** | Прогресс создания бэкапа. Значение будет меняться в статусе бэкапа &#x60;create&#x60; от 0 до 99, для остальных статусов всегда будет возвращаться 0. | 

## Methods

### NewServerBackup

`func NewServerBackup(id float32, name string, comment NullableString, createdAt string, status string, size float32, type_ string, progress float32, ) *ServerBackup`

NewServerBackup instantiates a new ServerBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerBackupWithDefaults

`func NewServerBackupWithDefaults() *ServerBackup`

NewServerBackupWithDefaults instantiates a new ServerBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerBackup) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerBackup) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerBackup) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *ServerBackup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerBackup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerBackup) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *ServerBackup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ServerBackup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ServerBackup) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *ServerBackup) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ServerBackup) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCreatedAt

`func (o *ServerBackup) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerBackup) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerBackup) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetStatus

`func (o *ServerBackup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerBackup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerBackup) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSize

`func (o *ServerBackup) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ServerBackup) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ServerBackup) SetSize(v float32)`

SetSize sets Size field to given value.


### GetType

`func (o *ServerBackup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerBackup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerBackup) SetType(v string)`

SetType sets Type field to given value.


### GetProgress

`func (o *ServerBackup) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ServerBackup) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ServerBackup) SetProgress(v float32)`

SetProgress sets Progress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


