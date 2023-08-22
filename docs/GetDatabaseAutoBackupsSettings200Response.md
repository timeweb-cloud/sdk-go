# GetDatabaseAutoBackupsSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**AutoBackupsSettings** | [**[]AutoBackup**](AutoBackup.md) |  | 

## Methods

### NewGetDatabaseAutoBackupsSettings200Response

`func NewGetDatabaseAutoBackupsSettings200Response(meta Meta, autoBackupsSettings []AutoBackup, ) *GetDatabaseAutoBackupsSettings200Response`

NewGetDatabaseAutoBackupsSettings200Response instantiates a new GetDatabaseAutoBackupsSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabaseAutoBackupsSettings200ResponseWithDefaults

`func NewGetDatabaseAutoBackupsSettings200ResponseWithDefaults() *GetDatabaseAutoBackupsSettings200Response`

NewGetDatabaseAutoBackupsSettings200ResponseWithDefaults instantiates a new GetDatabaseAutoBackupsSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDatabaseAutoBackupsSettings200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDatabaseAutoBackupsSettings200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDatabaseAutoBackupsSettings200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetAutoBackupsSettings

`func (o *GetDatabaseAutoBackupsSettings200Response) GetAutoBackupsSettings() []AutoBackup`

GetAutoBackupsSettings returns the AutoBackupsSettings field if non-nil, zero value otherwise.

### GetAutoBackupsSettingsOk

`func (o *GetDatabaseAutoBackupsSettings200Response) GetAutoBackupsSettingsOk() (*[]AutoBackup, bool)`

GetAutoBackupsSettingsOk returns a tuple with the AutoBackupsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBackupsSettings

`func (o *GetDatabaseAutoBackupsSettings200Response) SetAutoBackupsSettings(v []AutoBackup)`

SetAutoBackupsSettings sets AutoBackupsSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


