# GetImageDownloadURLs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Meta** | [**Meta**](Meta.md) |  | 
**Downloads** | [**[]ImageDownloadAPI**](ImageDownloadAPI.md) | Массив объектов \&quot;Ссылка на загрузку\&quot; | 

## Methods

### NewGetImageDownloadURLs200Response

`func NewGetImageDownloadURLs200Response(responseId string, meta Meta, downloads []ImageDownloadAPI, ) *GetImageDownloadURLs200Response`

NewGetImageDownloadURLs200Response instantiates a new GetImageDownloadURLs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageDownloadURLs200ResponseWithDefaults

`func NewGetImageDownloadURLs200ResponseWithDefaults() *GetImageDownloadURLs200Response`

NewGetImageDownloadURLs200ResponseWithDefaults instantiates a new GetImageDownloadURLs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *GetImageDownloadURLs200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetImageDownloadURLs200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetImageDownloadURLs200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetMeta

`func (o *GetImageDownloadURLs200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetImageDownloadURLs200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetImageDownloadURLs200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDownloads

`func (o *GetImageDownloadURLs200Response) GetDownloads() []ImageDownloadAPI`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *GetImageDownloadURLs200Response) GetDownloadsOk() (*[]ImageDownloadAPI, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *GetImageDownloadURLs200Response) SetDownloads(v []ImageDownloadAPI)`

SetDownloads sets Downloads field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


