# ImageDownloadsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса. | [optional] 
**Meta** | [**Meta**](Meta.md) |  | 
**Downloads** | [**[]ImageDownload**](ImageDownload.md) |  | 

## Methods

### NewImageDownloadsResponse

`func NewImageDownloadsResponse(meta Meta, downloads []ImageDownload, ) *ImageDownloadsResponse`

NewImageDownloadsResponse instantiates a new ImageDownloadsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDownloadsResponseWithDefaults

`func NewImageDownloadsResponseWithDefaults() *ImageDownloadsResponse`

NewImageDownloadsResponseWithDefaults instantiates a new ImageDownloadsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *ImageDownloadsResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *ImageDownloadsResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *ImageDownloadsResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *ImageDownloadsResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *ImageDownloadsResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ImageDownloadsResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ImageDownloadsResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDownloads

`func (o *ImageDownloadsResponse) GetDownloads() []ImageDownload`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *ImageDownloadsResponse) GetDownloadsOk() (*[]ImageDownload, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *ImageDownloadsResponse) SetDownloads(v []ImageDownload)`

SetDownloads sets Downloads field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


