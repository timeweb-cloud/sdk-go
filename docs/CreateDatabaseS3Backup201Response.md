# CreateDatabaseS3Backup201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | [**S3Backup**](S3Backup.md) |  | 

## Methods

### NewCreateDatabaseS3Backup201Response

`func NewCreateDatabaseS3Backup201Response(backup S3Backup, ) *CreateDatabaseS3Backup201Response`

NewCreateDatabaseS3Backup201Response instantiates a new CreateDatabaseS3Backup201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatabaseS3Backup201ResponseWithDefaults

`func NewCreateDatabaseS3Backup201ResponseWithDefaults() *CreateDatabaseS3Backup201Response`

NewCreateDatabaseS3Backup201ResponseWithDefaults instantiates a new CreateDatabaseS3Backup201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *CreateDatabaseS3Backup201Response) GetBackup() S3Backup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *CreateDatabaseS3Backup201Response) GetBackupOk() (*S3Backup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *CreateDatabaseS3Backup201Response) SetBackup(v S3Backup)`

SetBackup sets Backup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


