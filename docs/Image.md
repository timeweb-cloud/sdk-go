# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID образа. | 
**Status** | [**ImageStatus**](ImageStatus.md) |  | 
**CreatedAt** | **string** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был создан образ. | 
**DeletedAt** | **string** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был удален образ. | 
**Size** | **int32** | Размер физического диска в мегабайтах. | 
**VirtualSize** | **int32** | Размер виртуального диска в мегабайтах. | 
**Name** | **string** | Имя образа. | 
**Description** | **string** | Описание образа. | 
**DiskId** | **int32** | ID связанного с образом диска. | 
**Location** | **string** | Локация образа. | 
**Os** | [**OS**](OS.md) |  | 
**Progress** | **int32** | Процент создания образа. | 
**IsCustom** | **bool** | Логическое значение, которое показывает, является ли образ кастомным. | 
**Type** | **string** | Тип образа. | 

## Methods

### NewImage

`func NewImage(id string, status ImageStatus, createdAt string, deletedAt string, size int32, virtualSize int32, name string, description string, diskId int32, location string, os OS, progress int32, isCustom bool, type_ string, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Image) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *Image) GetStatus() ImageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Image) GetStatusOk() (*ImageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Image) SetStatus(v ImageStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *Image) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Image) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Image) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeletedAt

`func (o *Image) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Image) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Image) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.


### GetSize

`func (o *Image) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Image) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Image) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVirtualSize

`func (o *Image) GetVirtualSize() int32`

GetVirtualSize returns the VirtualSize field if non-nil, zero value otherwise.

### GetVirtualSizeOk

`func (o *Image) GetVirtualSizeOk() (*int32, bool)`

GetVirtualSizeOk returns a tuple with the VirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSize

`func (o *Image) SetVirtualSize(v int32)`

SetVirtualSize sets VirtualSize field to given value.


### GetName

`func (o *Image) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Image) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Image) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Image) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Image) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Image) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDiskId

`func (o *Image) GetDiskId() int32`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *Image) GetDiskIdOk() (*int32, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *Image) SetDiskId(v int32)`

SetDiskId sets DiskId field to given value.


### GetLocation

`func (o *Image) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Image) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Image) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetOs

`func (o *Image) GetOs() OS`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Image) GetOsOk() (*OS, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Image) SetOs(v OS)`

SetOs sets Os field to given value.


### GetProgress

`func (o *Image) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Image) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Image) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetIsCustom

`func (o *Image) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *Image) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *Image) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.


### GetType

`func (o *Image) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Image) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Image) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


