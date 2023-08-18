# GetServerDiskAutoBackupSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoBackupsSettings** | [**AutoBackup**](AutoBackup.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetServerDiskAutoBackupSettings200Response

`func NewGetServerDiskAutoBackupSettings200Response(autoBackupsSettings AutoBackup, responseId string, ) *GetServerDiskAutoBackupSettings200Response`

NewGetServerDiskAutoBackupSettings200Response instantiates a new GetServerDiskAutoBackupSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerDiskAutoBackupSettings200ResponseWithDefaults

`func NewGetServerDiskAutoBackupSettings200ResponseWithDefaults() *GetServerDiskAutoBackupSettings200Response`

NewGetServerDiskAutoBackupSettings200ResponseWithDefaults instantiates a new GetServerDiskAutoBackupSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoBackupsSettings

`func (o *GetServerDiskAutoBackupSettings200Response) GetAutoBackupsSettings() AutoBackup`

GetAutoBackupsSettings returns the AutoBackupsSettings field if non-nil, zero value otherwise.

### GetAutoBackupsSettingsOk

`func (o *GetServerDiskAutoBackupSettings200Response) GetAutoBackupsSettingsOk() (*AutoBackup, bool)`

GetAutoBackupsSettingsOk returns a tuple with the AutoBackupsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBackupsSettings

`func (o *GetServerDiskAutoBackupSettings200Response) SetAutoBackupsSettings(v AutoBackup)`

SetAutoBackupsSettings sets AutoBackupsSettings field to given value.


### GetResponseId

`func (o *GetServerDiskAutoBackupSettings200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetServerDiskAutoBackupSettings200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetServerDiskAutoBackupSettings200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


