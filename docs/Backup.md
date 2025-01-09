# Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID резервной копии. | 
**Name** | **string** | Название резервной копии. | 
**Comment** | **NullableString** | Комментарий. | 
**CreatedAt** | **time.Time** | Дата создания. | 
**Status** | **string** | Статус бэкапа. | 
**Size** | **int32** | Размер резервной копии (Мб). | 
**Type** | **string** | Тип бэкапа. | 
**Progress** | **float32** | Прогресс создания бэкапа. Значение будет меняться в статусе бэкапа &#x60;create&#x60; от 0 до 99, для остальных статусов всегда будет возвращаться 0. | 

## Methods

### NewBackup

`func NewBackup(id int32, name string, comment NullableString, createdAt time.Time, status string, size int32, type_ string, progress float32, ) *Backup`

NewBackup instantiates a new Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWithDefaults

`func NewBackupWithDefaults() *Backup`

NewBackupWithDefaults instantiates a new Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Backup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Backup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Backup) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Backup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Backup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Backup) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *Backup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Backup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Backup) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *Backup) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Backup) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCreatedAt

`func (o *Backup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Backup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Backup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStatus

`func (o *Backup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Backup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Backup) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSize

`func (o *Backup) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Backup) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Backup) SetSize(v int32)`

SetSize sets Size field to given value.


### GetType

`func (o *Backup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Backup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Backup) SetType(v string)`

SetType sets Type field to given value.


### GetProgress

`func (o *Backup) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Backup) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Backup) SetProgress(v float32)`

SetProgress sets Progress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


