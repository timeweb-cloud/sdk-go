# UpdateDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Новое название документа | [optional] 
**Url** | Pointer to **string** | Новый HTTP(S)-адрес. Только для документов типа &#x60;link&#x60;; после смены адреса документ переиндексируется | [optional] 
**IndexingSchedule** | Pointer to [**IndexingSchedule**](IndexingSchedule.md) |  | [optional] 

## Methods

### NewUpdateDocument

`func NewUpdateDocument() *UpdateDocument`

NewUpdateDocument instantiates a new UpdateDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDocumentWithDefaults

`func NewUpdateDocumentWithDefaults() *UpdateDocument`

NewUpdateDocumentWithDefaults instantiates a new UpdateDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDocument) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDocument) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateDocument) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateDocument) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateDocument) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateDocument) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIndexingSchedule

`func (o *UpdateDocument) GetIndexingSchedule() IndexingSchedule`

GetIndexingSchedule returns the IndexingSchedule field if non-nil, zero value otherwise.

### GetIndexingScheduleOk

`func (o *UpdateDocument) GetIndexingScheduleOk() (*IndexingSchedule, bool)`

GetIndexingScheduleOk returns a tuple with the IndexingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingSchedule

`func (o *UpdateDocument) SetIndexingSchedule(v IndexingSchedule)`

SetIndexingSchedule sets IndexingSchedule field to given value.

### HasIndexingSchedule

`func (o *UpdateDocument) HasIndexingSchedule() bool`

HasIndexingSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


