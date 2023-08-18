# ImageDownloadAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор ссылки | 
**CreatedAt** | **time.Time** | Дата и время создания ссылки | 
**Image** | **string** | Идентификатор образа | 
**Type** | [**URLType**](URLType.md) |  | 
**Url** | Pointer to **string** | Ссылка на скачивание | [optional] 
**Status** | [**UrlStatus**](UrlStatus.md) |  | 
**Progress** | **int32** | Прогресс загрузки образа | 

## Methods

### NewImageDownloadAPI

`func NewImageDownloadAPI(id string, createdAt time.Time, image string, type_ URLType, status UrlStatus, progress int32, ) *ImageDownloadAPI`

NewImageDownloadAPI instantiates a new ImageDownloadAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDownloadAPIWithDefaults

`func NewImageDownloadAPIWithDefaults() *ImageDownloadAPI`

NewImageDownloadAPIWithDefaults instantiates a new ImageDownloadAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageDownloadAPI) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageDownloadAPI) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageDownloadAPI) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ImageDownloadAPI) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImageDownloadAPI) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImageDownloadAPI) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetImage

`func (o *ImageDownloadAPI) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ImageDownloadAPI) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ImageDownloadAPI) SetImage(v string)`

SetImage sets Image field to given value.


### GetType

`func (o *ImageDownloadAPI) GetType() URLType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageDownloadAPI) GetTypeOk() (*URLType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageDownloadAPI) SetType(v URLType)`

SetType sets Type field to given value.


### GetUrl

`func (o *ImageDownloadAPI) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageDownloadAPI) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageDownloadAPI) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ImageDownloadAPI) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *ImageDownloadAPI) GetStatus() UrlStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageDownloadAPI) GetStatusOk() (*UrlStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageDownloadAPI) SetStatus(v UrlStatus)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *ImageDownloadAPI) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ImageDownloadAPI) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ImageDownloadAPI) SetProgress(v int32)`

SetProgress sets Progress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


