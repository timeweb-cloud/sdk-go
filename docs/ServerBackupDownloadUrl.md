# ServerBackupDownloadUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **string** | ID бэкапа. | 
**DownloadUrl** | **string** | Ссылка для скачивания бэкапа. Действует 7 дней. | 

## Methods

### NewServerBackupDownloadUrl

`func NewServerBackupDownloadUrl(backupId string, downloadUrl string, ) *ServerBackupDownloadUrl`

NewServerBackupDownloadUrl instantiates a new ServerBackupDownloadUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerBackupDownloadUrlWithDefaults

`func NewServerBackupDownloadUrlWithDefaults() *ServerBackupDownloadUrl`

NewServerBackupDownloadUrlWithDefaults instantiates a new ServerBackupDownloadUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *ServerBackupDownloadUrl) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *ServerBackupDownloadUrl) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *ServerBackupDownloadUrl) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### GetDownloadUrl

`func (o *ServerBackupDownloadUrl) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ServerBackupDownloadUrl) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ServerBackupDownloadUrl) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


