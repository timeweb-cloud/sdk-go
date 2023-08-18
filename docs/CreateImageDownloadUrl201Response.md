# CreateImageDownloadUrl201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Download** | [**ImageDownloadAPI**](ImageDownloadAPI.md) |  | 

## Methods

### NewCreateImageDownloadUrl201Response

`func NewCreateImageDownloadUrl201Response(responseId string, download ImageDownloadAPI, ) *CreateImageDownloadUrl201Response`

NewCreateImageDownloadUrl201Response instantiates a new CreateImageDownloadUrl201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageDownloadUrl201ResponseWithDefaults

`func NewCreateImageDownloadUrl201ResponseWithDefaults() *CreateImageDownloadUrl201Response`

NewCreateImageDownloadUrl201ResponseWithDefaults instantiates a new CreateImageDownloadUrl201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *CreateImageDownloadUrl201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateImageDownloadUrl201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateImageDownloadUrl201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetDownload

`func (o *CreateImageDownloadUrl201Response) GetDownload() ImageDownloadAPI`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *CreateImageDownloadUrl201Response) GetDownloadOk() (*ImageDownloadAPI, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *CreateImageDownloadUrl201Response) SetDownload(v ImageDownloadAPI)`

SetDownload sets Download field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


