# AiGatewayFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID файла | [optional] 
**Object** | Pointer to **string** | Тип объекта | [optional] 
**Bytes** | Pointer to **int32** | Размер файла в байтах | [optional] 
**CreatedAt** | Pointer to **int32** | Время загрузки, Unix-время в секундах | [optional] 
**Filename** | Pointer to **string** | Имя файла | [optional] 
**Purpose** | Pointer to **string** | Назначение файла | [optional] 
**Status** | Pointer to **string** | Статус обработки файла, например &#x60;uploaded&#x60; или &#x60;processed&#x60; | [optional] 
**ExpiresAt** | Pointer to **NullableInt32** | Когда файл будет удалён, Unix-время в секундах. &#x60;null&#x60; — не ограничено | [optional] 

## Methods

### NewAiGatewayFile

`func NewAiGatewayFile() *AiGatewayFile`

NewAiGatewayFile instantiates a new AiGatewayFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayFileWithDefaults

`func NewAiGatewayFileWithDefaults() *AiGatewayFile`

NewAiGatewayFileWithDefaults instantiates a new AiGatewayFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiGatewayFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiGatewayFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiGatewayFile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiGatewayFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *AiGatewayFile) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiGatewayFile) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiGatewayFile) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiGatewayFile) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetBytes

`func (o *AiGatewayFile) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *AiGatewayFile) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *AiGatewayFile) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *AiGatewayFile) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AiGatewayFile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiGatewayFile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiGatewayFile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AiGatewayFile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFilename

`func (o *AiGatewayFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AiGatewayFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AiGatewayFile) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *AiGatewayFile) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetPurpose

`func (o *AiGatewayFile) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *AiGatewayFile) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *AiGatewayFile) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *AiGatewayFile) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetStatus

`func (o *AiGatewayFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AiGatewayFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AiGatewayFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AiGatewayFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AiGatewayFile) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AiGatewayFile) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AiGatewayFile) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AiGatewayFile) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *AiGatewayFile) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *AiGatewayFile) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


