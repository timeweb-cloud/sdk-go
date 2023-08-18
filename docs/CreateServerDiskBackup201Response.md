# CreateServerDiskBackup201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to [**ServerBackup**](ServerBackup.md) |  | [optional] 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCreateServerDiskBackup201Response

`func NewCreateServerDiskBackup201Response(responseId string, ) *CreateServerDiskBackup201Response`

NewCreateServerDiskBackup201Response instantiates a new CreateServerDiskBackup201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerDiskBackup201ResponseWithDefaults

`func NewCreateServerDiskBackup201ResponseWithDefaults() *CreateServerDiskBackup201Response`

NewCreateServerDiskBackup201ResponseWithDefaults instantiates a new CreateServerDiskBackup201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *CreateServerDiskBackup201Response) GetBackup() ServerBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *CreateServerDiskBackup201Response) GetBackupOk() (*ServerBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *CreateServerDiskBackup201Response) SetBackup(v ServerBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *CreateServerDiskBackup201Response) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetResponseId

`func (o *CreateServerDiskBackup201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateServerDiskBackup201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateServerDiskBackup201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


