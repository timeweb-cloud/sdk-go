# ImageOutAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор образа | 
**Status** | [**ImageStatus**](ImageStatus.md) |  | 
**CreatedAt** | **time.Time** | Дата и время создания | 
**DeletedAt** | **time.Time** | Дата и время удаления | 
**Size** | **int32** | Размер в мегабайтах | 
**Name** | **string** | Имя образа | 
**Description** | **string** | Описание образа | 
**DiskId** | **int32** | Идентификатор связанного с образом диска | 
**Location** | Pointer to **string** | Локация, в которой создан образ | [optional] 
**Os** | [**OS**](OS.md) |  | 
**Progress** | **int32** | Процент создания образа | 
**IsCustom** | **bool** | Признак указывающий на то является ли образ кастомным | 

## Methods

### NewImageOutAPI

`func NewImageOutAPI(id string, status ImageStatus, createdAt time.Time, deletedAt time.Time, size int32, name string, description string, diskId int32, os OS, progress int32, isCustom bool, ) *ImageOutAPI`

NewImageOutAPI instantiates a new ImageOutAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageOutAPIWithDefaults

`func NewImageOutAPIWithDefaults() *ImageOutAPI`

NewImageOutAPIWithDefaults instantiates a new ImageOutAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageOutAPI) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageOutAPI) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageOutAPI) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *ImageOutAPI) GetStatus() ImageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageOutAPI) GetStatusOk() (*ImageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageOutAPI) SetStatus(v ImageStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ImageOutAPI) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImageOutAPI) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImageOutAPI) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeletedAt

`func (o *ImageOutAPI) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ImageOutAPI) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ImageOutAPI) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetSize

`func (o *ImageOutAPI) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImageOutAPI) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImageOutAPI) SetSize(v int32)`

SetSize sets Size field to given value.


### GetName

`func (o *ImageOutAPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageOutAPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageOutAPI) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ImageOutAPI) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageOutAPI) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageOutAPI) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDiskId

`func (o *ImageOutAPI) GetDiskId() int32`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *ImageOutAPI) GetDiskIdOk() (*int32, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *ImageOutAPI) SetDiskId(v int32)`

SetDiskId sets DiskId field to given value.


### GetLocation

`func (o *ImageOutAPI) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ImageOutAPI) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ImageOutAPI) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ImageOutAPI) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetOs

`func (o *ImageOutAPI) GetOs() OS`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ImageOutAPI) GetOsOk() (*OS, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ImageOutAPI) SetOs(v OS)`

SetOs sets Os field to given value.


### GetProgress

`func (o *ImageOutAPI) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ImageOutAPI) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ImageOutAPI) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetIsCustom

`func (o *ImageOutAPI) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *ImageOutAPI) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *ImageOutAPI) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


