# ImageDownloadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса. | [optional] 
**Download** | [**ImageDownload**](ImageDownload.md) |  | 

## Methods

### NewImageDownloadResponse

`func NewImageDownloadResponse(download ImageDownload, ) *ImageDownloadResponse`

NewImageDownloadResponse instantiates a new ImageDownloadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDownloadResponseWithDefaults

`func NewImageDownloadResponseWithDefaults() *ImageDownloadResponse`

NewImageDownloadResponseWithDefaults instantiates a new ImageDownloadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *ImageDownloadResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *ImageDownloadResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *ImageDownloadResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *ImageDownloadResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetDownload

`func (o *ImageDownloadResponse) GetDownload() ImageDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ImageDownloadResponse) GetDownloadOk() (*ImageDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ImageDownloadResponse) SetDownload(v ImageDownload)`

SetDownload sets Download field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


