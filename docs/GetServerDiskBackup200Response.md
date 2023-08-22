# GetServerDiskBackup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | [**ServerBackup**](ServerBackup.md) |  | 

## Methods

### NewGetServerDiskBackup200Response

`func NewGetServerDiskBackup200Response(backup ServerBackup, ) *GetServerDiskBackup200Response`

NewGetServerDiskBackup200Response instantiates a new GetServerDiskBackup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerDiskBackup200ResponseWithDefaults

`func NewGetServerDiskBackup200ResponseWithDefaults() *GetServerDiskBackup200Response`

NewGetServerDiskBackup200ResponseWithDefaults instantiates a new GetServerDiskBackup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *GetServerDiskBackup200Response) GetBackup() ServerBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *GetServerDiskBackup200Response) GetBackupOk() (*ServerBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *GetServerDiskBackup200Response) SetBackup(v ServerBackup)`

SetBackup sets Backup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


