# BackupDownloadUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | Имя файла бэкапа. Допустимы буквы латиницы и кириллицы, цифры, символ подчеркивания, а также &#x60;-&#x60;, &#x60;.&#x60;, &#x60;,&#x60;, &#x60;:&#x60; и пробел. | 

## Methods

### NewBackupDownloadUrlRequest

`func NewBackupDownloadUrlRequest(filename string, ) *BackupDownloadUrlRequest`

NewBackupDownloadUrlRequest instantiates a new BackupDownloadUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDownloadUrlRequestWithDefaults

`func NewBackupDownloadUrlRequestWithDefaults() *BackupDownloadUrlRequest`

NewBackupDownloadUrlRequestWithDefaults instantiates a new BackupDownloadUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *BackupDownloadUrlRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *BackupDownloadUrlRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *BackupDownloadUrlRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


