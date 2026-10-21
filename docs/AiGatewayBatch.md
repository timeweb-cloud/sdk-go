# AiGatewayBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID батча | [optional] 
**Object** | Pointer to **string** | Тип объекта | [optional] 
**Endpoint** | Pointer to **string** | Эндпоинт запросов батча | [optional] 
**Model** | Pointer to **NullableString** | Модель батча. На этапе &#x60;validating&#x60; может быть &#x60;null&#x60; | [optional] 
**InputFileId** | Pointer to **string** | ID входного файла | [optional] 
**OutputFileId** | Pointer to **NullableString** | ID файла результатов — когда батч выполнен | [optional] 
**ErrorFileId** | Pointer to **NullableString** | ID файла с ошибками отдельных запросов, если они были | [optional] 
**CompletionWindow** | Pointer to **string** | Срок выполнения | [optional] 
**Status** | Pointer to **string** | Статус батча. &#x60;cancelling&#x60; — отмена выполняется, может занять несколько минут | [optional] 
**CreatedAt** | Pointer to **int32** | Время создания, Unix-время в секундах | [optional] 
**CompletedAt** | Pointer to **NullableInt32** | Время завершения, Unix-время в секундах | [optional] 
**ExpiresAt** | Pointer to **NullableInt32** | Крайний срок выполнения, Unix-время в секундах | [optional] 
**RequestCounts** | Pointer to [**AiGatewayBatchRequestCounts**](AiGatewayBatchRequestCounts.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Метки батча из запроса | [optional] 
**Usage** | Pointer to **map[string]interface{}** | Расход токенов батча | [optional] 

## Methods

### NewAiGatewayBatch

`func NewAiGatewayBatch() *AiGatewayBatch`

NewAiGatewayBatch instantiates a new AiGatewayBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayBatchWithDefaults

`func NewAiGatewayBatchWithDefaults() *AiGatewayBatch`

NewAiGatewayBatchWithDefaults instantiates a new AiGatewayBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiGatewayBatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiGatewayBatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiGatewayBatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiGatewayBatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *AiGatewayBatch) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiGatewayBatch) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiGatewayBatch) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiGatewayBatch) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetEndpoint

`func (o *AiGatewayBatch) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AiGatewayBatch) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AiGatewayBatch) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AiGatewayBatch) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetModel

`func (o *AiGatewayBatch) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiGatewayBatch) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiGatewayBatch) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiGatewayBatch) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *AiGatewayBatch) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *AiGatewayBatch) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetInputFileId

`func (o *AiGatewayBatch) GetInputFileId() string`

GetInputFileId returns the InputFileId field if non-nil, zero value otherwise.

### GetInputFileIdOk

`func (o *AiGatewayBatch) GetInputFileIdOk() (*string, bool)`

GetInputFileIdOk returns a tuple with the InputFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFileId

`func (o *AiGatewayBatch) SetInputFileId(v string)`

SetInputFileId sets InputFileId field to given value.

### HasInputFileId

`func (o *AiGatewayBatch) HasInputFileId() bool`

HasInputFileId returns a boolean if a field has been set.

### GetOutputFileId

`func (o *AiGatewayBatch) GetOutputFileId() string`

GetOutputFileId returns the OutputFileId field if non-nil, zero value otherwise.

### GetOutputFileIdOk

`func (o *AiGatewayBatch) GetOutputFileIdOk() (*string, bool)`

GetOutputFileIdOk returns a tuple with the OutputFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFileId

`func (o *AiGatewayBatch) SetOutputFileId(v string)`

SetOutputFileId sets OutputFileId field to given value.

### HasOutputFileId

`func (o *AiGatewayBatch) HasOutputFileId() bool`

HasOutputFileId returns a boolean if a field has been set.

### SetOutputFileIdNil

`func (o *AiGatewayBatch) SetOutputFileIdNil(b bool)`

 SetOutputFileIdNil sets the value for OutputFileId to be an explicit nil

### UnsetOutputFileId
`func (o *AiGatewayBatch) UnsetOutputFileId()`

UnsetOutputFileId ensures that no value is present for OutputFileId, not even an explicit nil
### GetErrorFileId

`func (o *AiGatewayBatch) GetErrorFileId() string`

GetErrorFileId returns the ErrorFileId field if non-nil, zero value otherwise.

### GetErrorFileIdOk

`func (o *AiGatewayBatch) GetErrorFileIdOk() (*string, bool)`

GetErrorFileIdOk returns a tuple with the ErrorFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFileId

`func (o *AiGatewayBatch) SetErrorFileId(v string)`

SetErrorFileId sets ErrorFileId field to given value.

### HasErrorFileId

`func (o *AiGatewayBatch) HasErrorFileId() bool`

HasErrorFileId returns a boolean if a field has been set.

### SetErrorFileIdNil

`func (o *AiGatewayBatch) SetErrorFileIdNil(b bool)`

 SetErrorFileIdNil sets the value for ErrorFileId to be an explicit nil

### UnsetErrorFileId
`func (o *AiGatewayBatch) UnsetErrorFileId()`

UnsetErrorFileId ensures that no value is present for ErrorFileId, not even an explicit nil
### GetCompletionWindow

`func (o *AiGatewayBatch) GetCompletionWindow() string`

GetCompletionWindow returns the CompletionWindow field if non-nil, zero value otherwise.

### GetCompletionWindowOk

`func (o *AiGatewayBatch) GetCompletionWindowOk() (*string, bool)`

GetCompletionWindowOk returns a tuple with the CompletionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionWindow

`func (o *AiGatewayBatch) SetCompletionWindow(v string)`

SetCompletionWindow sets CompletionWindow field to given value.

### HasCompletionWindow

`func (o *AiGatewayBatch) HasCompletionWindow() bool`

HasCompletionWindow returns a boolean if a field has been set.

### GetStatus

`func (o *AiGatewayBatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AiGatewayBatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AiGatewayBatch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AiGatewayBatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AiGatewayBatch) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiGatewayBatch) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiGatewayBatch) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AiGatewayBatch) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *AiGatewayBatch) GetCompletedAt() int32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *AiGatewayBatch) GetCompletedAtOk() (*int32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *AiGatewayBatch) SetCompletedAt(v int32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *AiGatewayBatch) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *AiGatewayBatch) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *AiGatewayBatch) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetExpiresAt

`func (o *AiGatewayBatch) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AiGatewayBatch) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AiGatewayBatch) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AiGatewayBatch) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *AiGatewayBatch) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *AiGatewayBatch) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetRequestCounts

`func (o *AiGatewayBatch) GetRequestCounts() AiGatewayBatchRequestCounts`

GetRequestCounts returns the RequestCounts field if non-nil, zero value otherwise.

### GetRequestCountsOk

`func (o *AiGatewayBatch) GetRequestCountsOk() (*AiGatewayBatchRequestCounts, bool)`

GetRequestCountsOk returns a tuple with the RequestCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCounts

`func (o *AiGatewayBatch) SetRequestCounts(v AiGatewayBatchRequestCounts)`

SetRequestCounts sets RequestCounts field to given value.

### HasRequestCounts

`func (o *AiGatewayBatch) HasRequestCounts() bool`

HasRequestCounts returns a boolean if a field has been set.

### GetMetadata

`func (o *AiGatewayBatch) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiGatewayBatch) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiGatewayBatch) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiGatewayBatch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AiGatewayBatch) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AiGatewayBatch) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetUsage

`func (o *AiGatewayBatch) GetUsage() map[string]interface{}`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiGatewayBatch) GetUsageOk() (*map[string]interface{}, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiGatewayBatch) SetUsage(v map[string]interface{})`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiGatewayBatch) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *AiGatewayBatch) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *AiGatewayBatch) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


