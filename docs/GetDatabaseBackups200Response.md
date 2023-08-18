# GetDatabaseBackups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Backups** | [**[]Backup**](Backup.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetDatabaseBackups200Response

`func NewGetDatabaseBackups200Response(meta Meta, backups []Backup, responseId string, ) *GetDatabaseBackups200Response`

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


### GetResponseId

`func (o *GetDatabaseBackups200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetDatabaseBackups200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetDatabaseBackups200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


