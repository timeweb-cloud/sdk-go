# ImageInAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Имя образа. | [optional] 
**Description** | Pointer to **string** | Описание образа. | [optional] 
**DiskId** | Pointer to **int32** | ID диска, для которого создается образ. | [optional] 
**UploadUrl** | Pointer to **string** | Ссылка для загрузки образа. | [optional] 
**Location** | [**Location**](Location.md) |  | 
**Os** | [**OS**](OS.md) |  | 

## Methods

### NewImageInAPI

`func NewImageInAPI(location Location, os OS, ) *ImageInAPI`

NewImageInAPI instantiates a new ImageInAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageInAPIWithDefaults

`func NewImageInAPIWithDefaults() *ImageInAPI`

NewImageInAPIWithDefaults instantiates a new ImageInAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImageInAPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageInAPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageInAPI) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageInAPI) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ImageInAPI) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageInAPI) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageInAPI) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageInAPI) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskId

`func (o *ImageInAPI) GetDiskId() int32`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *ImageInAPI) GetDiskIdOk() (*int32, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *ImageInAPI) SetDiskId(v int32)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *ImageInAPI) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### GetUploadUrl

`func (o *ImageInAPI) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *ImageInAPI) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *ImageInAPI) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *ImageInAPI) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetLocation

`func (o *ImageInAPI) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ImageInAPI) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ImageInAPI) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetOs

`func (o *ImageInAPI) GetOs() OS`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ImageInAPI) GetOsOk() (*OS, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ImageInAPI) SetOs(v OS)`

SetOs sets Os field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


