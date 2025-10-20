# DocumentStatusInfoDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | Pointer to **time.Time** | Время начала индексации | [optional] 
**CurrentStage** | Pointer to **string** | Текущий этап индексации | [optional] 
**Progress** | Pointer to **float32** | Прогресс в процентах (0-100) | [optional] 

## Methods

### NewDocumentStatusInfoDetails

`func NewDocumentStatusInfoDetails() *DocumentStatusInfoDetails`

NewDocumentStatusInfoDetails instantiates a new DocumentStatusInfoDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentStatusInfoDetailsWithDefaults

`func NewDocumentStatusInfoDetailsWithDefaults() *DocumentStatusInfoDetails`

NewDocumentStatusInfoDetailsWithDefaults instantiates a new DocumentStatusInfoDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartedAt

`func (o *DocumentStatusInfoDetails) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DocumentStatusInfoDetails) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DocumentStatusInfoDetails) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DocumentStatusInfoDetails) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCurrentStage

`func (o *DocumentStatusInfoDetails) GetCurrentStage() string`

GetCurrentStage returns the CurrentStage field if non-nil, zero value otherwise.

### GetCurrentStageOk

`func (o *DocumentStatusInfoDetails) GetCurrentStageOk() (*string, bool)`

GetCurrentStageOk returns a tuple with the CurrentStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStage

`func (o *DocumentStatusInfoDetails) SetCurrentStage(v string)`

SetCurrentStage sets CurrentStage field to given value.

### HasCurrentStage

`func (o *DocumentStatusInfoDetails) HasCurrentStage() bool`

HasCurrentStage returns a boolean if a field has been set.

### GetProgress

`func (o *DocumentStatusInfoDetails) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DocumentStatusInfoDetails) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DocumentStatusInfoDetails) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DocumentStatusInfoDetails) HasProgress() bool`

HasProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


