# GetServerDiskBackups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Backups** | [**[]ServerBackup**](ServerBackup.md) |  | 

## Methods

### NewGetServerDiskBackups200Response

`func NewGetServerDiskBackups200Response(meta Meta, backups []ServerBackup, ) *GetServerDiskBackups200Response`

NewGetServerDiskBackups200Response instantiates a new GetServerDiskBackups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerDiskBackups200ResponseWithDefaults

`func NewGetServerDiskBackups200ResponseWithDefaults() *GetServerDiskBackups200Response`

NewGetServerDiskBackups200ResponseWithDefaults instantiates a new GetServerDiskBackups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetServerDiskBackups200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetServerDiskBackups200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetServerDiskBackups200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetBackups

`func (o *GetServerDiskBackups200Response) GetBackups() []ServerBackup`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *GetServerDiskBackups200Response) GetBackupsOk() (*[]ServerBackup, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *GetServerDiskBackups200Response) SetBackups(v []ServerBackup)`

SetBackups sets Backups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


