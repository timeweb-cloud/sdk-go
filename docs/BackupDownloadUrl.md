# BackupDownloadUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **string** | ID резервной копии. | 
**DownloadUrl** | **string** | Ссылка для скачивания бэкапа. | 

## Methods

### NewBackupDownloadUrl

`func NewBackupDownloadUrl(backupId string, downloadUrl string, ) *BackupDownloadUrl`

NewBackupDownloadUrl instantiates a new BackupDownloadUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDownloadUrlWithDefaults

`func NewBackupDownloadUrlWithDefaults() *BackupDownloadUrl`

NewBackupDownloadUrlWithDefaults instantiates a new BackupDownloadUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *BackupDownloadUrl) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *BackupDownloadUrl) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *BackupDownloadUrl) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### GetDownloadUrl

`func (o *BackupDownloadUrl) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *BackupDownloadUrl) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *BackupDownloadUrl) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


