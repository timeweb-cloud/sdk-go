# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Уникальный идентификатор документа | 
**Name** | **string** | Название документа | 
**Size** | **float32** | Размер документа в байтах | 
**Status** | **string** | Статус индексации документа | 
**IndexingVersion** | Pointer to **NullableString** | Версия индексации | [optional] 
**StatusInfo** | Pointer to [**NullableDocumentStatusInfo**](DocumentStatusInfo.md) |  | [optional] 
**CreatedAt** | **time.Time** | Дата создания документа | 
**UpdatedAt** | **time.Time** | Дата обновления документа | 

## Methods

### NewDocument

`func NewDocument(id float32, name string, size float32, status string, createdAt time.Time, updatedAt time.Time, ) *Document`

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

`func (o *Document) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Document) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Document) SetId(v float32)`

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

`func (o *Document) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Document) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Document) SetSize(v float32)`

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

### HasIndexingVersion

`func (o *Document) HasIndexingVersion() bool`

HasIndexingVersion returns a boolean if a field has been set.

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

### HasStatusInfo

`func (o *Document) HasStatusInfo() bool`

HasStatusInfo returns a boolean if a field has been set.

### SetStatusInfoNil

`func (o *Document) SetStatusInfoNil(b bool)`

 SetStatusInfoNil sets the value for StatusInfo to be an explicit nil

### UnsetStatusInfo
`func (o *Document) UnsetStatusInfo()`

UnsetStatusInfo ensures that no value is present for StatusInfo, not even an explicit nil
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


