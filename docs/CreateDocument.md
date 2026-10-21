# CreateDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название документа | 
**Type** | **string** | Тип документа. Поддерживается только &#x60;link&#x60;; файлы загружайте через &#x60;POST /api/v1/cloud-ai/knowledge-bases/{id}/upload&#x60; | 
**Url** | **string** | Публично доступный HTTP(S)-адрес документа | 
**IndexingSchedule** | Pointer to [**IndexingSchedule**](IndexingSchedule.md) |  | [optional] 

## Methods

### NewCreateDocument

`func NewCreateDocument(name string, type_ string, url string, ) *CreateDocument`

NewCreateDocument instantiates a new CreateDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDocumentWithDefaults

`func NewCreateDocumentWithDefaults() *CreateDocument`

NewCreateDocumentWithDefaults instantiates a new CreateDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDocument) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDocument) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *CreateDocument) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateDocument) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateDocument) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetIndexingSchedule

`func (o *CreateDocument) GetIndexingSchedule() IndexingSchedule`

GetIndexingSchedule returns the IndexingSchedule field if non-nil, zero value otherwise.

### GetIndexingScheduleOk

`func (o *CreateDocument) GetIndexingScheduleOk() (*IndexingSchedule, bool)`

GetIndexingScheduleOk returns a tuple with the IndexingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingSchedule

`func (o *CreateDocument) SetIndexingSchedule(v IndexingSchedule)`

SetIndexingSchedule sets IndexingSchedule field to given value.

### HasIndexingSchedule

`func (o *CreateDocument) HasIndexingSchedule() bool`

HasIndexingSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


