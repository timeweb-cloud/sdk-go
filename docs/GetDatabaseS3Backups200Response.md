# GetDatabaseS3Backups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backups** | [**[]S3Backup**](S3Backup.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetDatabaseS3Backups200Response

`func NewGetDatabaseS3Backups200Response(backups []S3Backup, meta Meta, ) *GetDatabaseS3Backups200Response`

NewGetDatabaseS3Backups200Response instantiates a new GetDatabaseS3Backups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabaseS3Backups200ResponseWithDefaults

`func NewGetDatabaseS3Backups200ResponseWithDefaults() *GetDatabaseS3Backups200Response`

NewGetDatabaseS3Backups200ResponseWithDefaults instantiates a new GetDatabaseS3Backups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackups

`func (o *GetDatabaseS3Backups200Response) GetBackups() []S3Backup`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *GetDatabaseS3Backups200Response) GetBackupsOk() (*[]S3Backup, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *GetDatabaseS3Backups200Response) SetBackups(v []S3Backup)`

SetBackups sets Backups field to given value.


### GetMeta

`func (o *GetDatabaseS3Backups200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDatabaseS3Backups200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDatabaseS3Backups200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


