# S3Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID резервной копии в формате UUID. | 
**Name** | **string** | Название резервной копии. В текущей реализации всегда совпадает со значением &#x60;id&#x60;. | 
**Status** | **string** | Статус резервной копии: &#x60;running&#x60; — копия создаётся, &#x60;success&#x60; — копия создана, &#x60;failed&#x60; — создать копию не удалось. Восстановить кластер можно только из копии со статусом &#x60;success&#x60;. | 
**Size** | **float32** | Размер резервной копии в байтах. | 
**Progress** | **float32** | Прогресс создания резервной копии, от 0 до 100. | 
**Type** | **string** | Тип резервной копии: &#x60;manual&#x60; — копия создана вручную, &#x60;auto&#x60; — копия создана по расписанию. | 
**Comment** | Pointer to **NullableString** | Комментарий к резервной копии. Если комментарий не задавался, возвращается &#x60;null&#x60;. | [optional] 
**CreatedAt** | **time.Time** | Дата и время создания резервной копии. | 

## Methods

### NewS3Backup

`func NewS3Backup(id string, name string, status string, size float32, progress float32, type_ string, createdAt time.Time, ) *S3Backup`

NewS3Backup instantiates a new S3Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BackupWithDefaults

`func NewS3BackupWithDefaults() *S3Backup`

NewS3BackupWithDefaults instantiates a new S3Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3Backup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3Backup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3Backup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *S3Backup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3Backup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3Backup) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *S3Backup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *S3Backup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *S3Backup) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSize

`func (o *S3Backup) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *S3Backup) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *S3Backup) SetSize(v float32)`

SetSize sets Size field to given value.


### GetProgress

`func (o *S3Backup) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *S3Backup) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *S3Backup) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetType

`func (o *S3Backup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *S3Backup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *S3Backup) SetType(v string)`

SetType sets Type field to given value.


### GetComment

`func (o *S3Backup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *S3Backup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *S3Backup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *S3Backup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *S3Backup) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *S3Backup) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCreatedAt

`func (o *S3Backup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *S3Backup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *S3Backup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


