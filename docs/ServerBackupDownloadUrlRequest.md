# ServerBackupDownloadUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | Имя файла бэкапа. Допустимы буквы латиницы и кириллицы, цифры, символ подчеркивания, а также &#x60;-&#x60;, &#x60;.&#x60;, &#x60;,&#x60;, &#x60;:&#x60; и пробел. | 

## Methods

### NewServerBackupDownloadUrlRequest

`func NewServerBackupDownloadUrlRequest(filename string, ) *ServerBackupDownloadUrlRequest`

NewServerBackupDownloadUrlRequest instantiates a new ServerBackupDownloadUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerBackupDownloadUrlRequestWithDefaults

`func NewServerBackupDownloadUrlRequestWithDefaults() *ServerBackupDownloadUrlRequest`

NewServerBackupDownloadUrlRequestWithDefaults instantiates a new ServerBackupDownloadUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *ServerBackupDownloadUrlRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ServerBackupDownloadUrlRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ServerBackupDownloadUrlRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


