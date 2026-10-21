# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID документа | 
**Name** | **string** | Название документа | 
**Size** | **int32** | Размер в байтах. Для документа по ссылке — значение заголовка &#x60;Content-Length&#x60; на момент добавления, &#x60;0&#x60;, если сервер его не вернул | 
**Status** | **string** | Статус индексации: - &#x60;new&#x60; — документ добавлен и ждёт индексации - &#x60;indexing&#x60; — индексируется - &#x60;indexed&#x60; — проиндексирован, используется в поиске - &#x60;error&#x60; — ошибка индексации, подробности в &#x60;status_info&#x60; | 
**Type** | **string** | Тип документа: &#x60;file&#x60; — загруженный файл, &#x60;link&#x60; — документ по ссылке | 
**IndexingSchedule** | [**IndexingSchedule**](IndexingSchedule.md) |  | 
**IndexingVersion** | **NullableString** | Версия индексации документа | 
**StatusInfo** | [**NullableDocumentStatusInfo**](DocumentStatusInfo.md) |  | 
**Link** | Pointer to **string** | URL документа. Только для документов типа &#x60;link&#x60; | [optional] 
**CreatedAt** | **time.Time** | Дата и время добавления документа | 
**UpdatedAt** | **time.Time** | Дата и время последнего изменения документа | 

## Methods

### NewDocument

`func NewDocument(id int32, name string, size int32, status string, type_ string, indexingSchedule IndexingSchedule, indexingVersion NullableString, statusInfo NullableDocumentStatusInfo, createdAt time.Time, updatedAt time.Time, ) *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Document) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Document) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Document) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Document) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Document) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Document) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *Document) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Document) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Document) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStatus

`func (o *Document) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Document) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Document) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *Document) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Document) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Document) SetType(v string)`

SetType sets Type field to given value.


### GetIndexingSchedule

`func (o *Document) GetIndexingSchedule() IndexingSchedule`

GetIndexingSchedule returns the IndexingSchedule field if non-nil, zero value otherwise.

### GetIndexingScheduleOk

`func (o *Document) GetIndexingScheduleOk() (*IndexingSchedule, bool)`

GetIndexingScheduleOk returns a tuple with the IndexingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingSchedule

`func (o *Document) SetIndexingSchedule(v IndexingSchedule)`

SetIndexingSchedule sets IndexingSchedule field to given value.


### GetIndexingVersion

`func (o *Document) GetIndexingVersion() string`

GetIndexingVersion returns the IndexingVersion field if non-nil, zero value otherwise.

### GetIndexingVersionOk

`func (o *Document) GetIndexingVersionOk() (*string, bool)`

GetIndexingVersionOk returns a tuple with the IndexingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingVersion

`func (o *Document) SetIndexingVersion(v string)`

SetIndexingVersion sets IndexingVersion field to given value.


### SetIndexingVersionNil

`func (o *Document) SetIndexingVersionNil(b bool)`

 SetIndexingVersionNil sets the value for IndexingVersion to be an explicit nil

### UnsetIndexingVersion
`func (o *Document) UnsetIndexingVersion()`

UnsetIndexingVersion ensures that no value is present for IndexingVersion, not even an explicit nil
### GetStatusInfo

`func (o *Document) GetStatusInfo() DocumentStatusInfo`

GetStatusInfo returns the StatusInfo field if non-nil, zero value otherwise.

### GetStatusInfoOk

`func (o *Document) GetStatusInfoOk() (*DocumentStatusInfo, bool)`

GetStatusInfoOk returns a tuple with the StatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfo

`func (o *Document) SetStatusInfo(v DocumentStatusInfo)`

SetStatusInfo sets StatusInfo field to given value.


### SetStatusInfoNil

`func (o *Document) SetStatusInfoNil(b bool)`

 SetStatusInfoNil sets the value for StatusInfo to be an explicit nil

### UnsetStatusInfo
`func (o *Document) UnsetStatusInfo()`

UnsetStatusInfo ensures that no value is present for StatusInfo, not even an explicit nil
### GetLink

`func (o *Document) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Document) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Document) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Document) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Document) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Document) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Document) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Document) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Document) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Document) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


