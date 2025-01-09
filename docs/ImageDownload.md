# ImageDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID ссылки. | 
**CreatedAt** | **string** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда была создана ссылка. | 
**Image** | **string** | ID образа. | 
**Type** | [**URLType**](URLType.md) |  | 
**Url** | Pointer to **string** | Ссылка на скачивание. | [optional] 
**Status** | [**UrlStatus**](UrlStatus.md) |  | 
**Progress** | **int32** | Прогресс загрузки образа. | 

## Methods

### NewImageDownload

`func NewImageDownload(id string, createdAt string, image string, type_ URLType, status UrlStatus, progress int32, ) *ImageDownload`

NewImageDownload instantiates a new ImageDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDownloadWithDefaults

`func NewImageDownloadWithDefaults() *ImageDownload`

NewImageDownloadWithDefaults instantiates a new ImageDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageDownload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageDownload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageDownload) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ImageDownload) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImageDownload) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImageDownload) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetImage

`func (o *ImageDownload) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ImageDownload) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ImageDownload) SetImage(v string)`

SetImage sets Image field to given value.


### GetType

`func (o *ImageDownload) GetType() URLType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageDownload) GetTypeOk() (*URLType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageDownload) SetType(v URLType)`

SetType sets Type field to given value.


### GetUrl

`func (o *ImageDownload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageDownload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageDownload) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ImageDownload) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *ImageDownload) GetStatus() UrlStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageDownload) GetStatusOk() (*UrlStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageDownload) SetStatus(v UrlStatus)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *ImageDownload) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ImageDownload) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ImageDownload) SetProgress(v int32)`

SetProgress sets Progress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


