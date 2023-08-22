# GetDatabaseBackups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Backups** | [**[]Backup**](Backup.md) |  | 

## Methods

### NewGetDatabaseBackups200Response

`func NewGetDatabaseBackups200Response(meta Meta, backups []Backup, ) *GetDatabaseBackups200Response`

NewGetDatabaseBackups200Response instantiates a new GetDatabaseBackups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabaseBackups200ResponseWithDefaults

`func NewGetDatabaseBackups200ResponseWithDefaults() *GetDatabaseBackups200Response`

NewGetDatabaseBackups200ResponseWithDefaults instantiates a new GetDatabaseBackups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDatabaseBackups200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDatabaseBackups200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDatabaseBackups200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetBackups

`func (o *GetDatabaseBackups200Response) GetBackups() []Backup`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *GetDatabaseBackups200Response) GetBackupsOk() (*[]Backup, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *GetDatabaseBackups200Response) SetBackups(v []Backup)`

SetBackups sets Backups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


