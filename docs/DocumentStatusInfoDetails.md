# DocumentStatusInfoDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | Pointer to **time.Time** | &#x60;indexing&#x60;: время начала индексации | [optional] 
**CurrentStage** | Pointer to **string** | &#x60;indexing&#x60;: текущий этап | [optional] 
**Progress** | Pointer to **float32** | &#x60;indexing&#x60;: прогресс в процентах | [optional] 
**IndexingDurationMs** | Pointer to **float32** | &#x60;indexed&#x60;: длительность индексации в миллисекундах | [optional] 
**TokensUsed** | Pointer to **float32** | &#x60;indexed&#x60;: израсходовано токенов | [optional] 
**TokensEstimated** | Pointer to **float32** | &#x60;indexed&#x60;: оценка количества токенов перед индексацией | [optional] 
**ChunksCount** | Pointer to **float32** | &#x60;indexed&#x60;: количество фрагментов, на которые разбит документ | [optional] 
**Needed** | Pointer to **float32** | &#x60;INSUFFICIENT_TOKENS&#x60;: сколько токенов нужно | [optional] 
**Available** | Pointer to **float32** | &#x60;INSUFFICIENT_TOKENS&#x60;: сколько токенов доступно | [optional] 
**Estimated** | Pointer to **float32** | &#x60;INSUFFICIENT_TOKENS&#x60;: оценка без запаса | [optional] 
**Balance** | Pointer to **float32** | &#x60;CRITICAL_BALANCE&#x60;, &#x60;NEGATIVE_BALANCE&#x60;: текущий остаток токенов | [optional] 
**MinRequired** | Pointer to **float32** | &#x60;CRITICAL_BALANCE&#x60;: минимально необходимый остаток | [optional] 

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

### GetIndexingDurationMs

`func (o *DocumentStatusInfoDetails) GetIndexingDurationMs() float32`

GetIndexingDurationMs returns the IndexingDurationMs field if non-nil, zero value otherwise.

### GetIndexingDurationMsOk

`func (o *DocumentStatusInfoDetails) GetIndexingDurationMsOk() (*float32, bool)`

GetIndexingDurationMsOk returns a tuple with the IndexingDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingDurationMs

`func (o *DocumentStatusInfoDetails) SetIndexingDurationMs(v float32)`

SetIndexingDurationMs sets IndexingDurationMs field to given value.

### HasIndexingDurationMs

`func (o *DocumentStatusInfoDetails) HasIndexingDurationMs() bool`

HasIndexingDurationMs returns a boolean if a field has been set.

### GetTokensUsed

`func (o *DocumentStatusInfoDetails) GetTokensUsed() float32`

GetTokensUsed returns the TokensUsed field if non-nil, zero value otherwise.

### GetTokensUsedOk

`func (o *DocumentStatusInfoDetails) GetTokensUsedOk() (*float32, bool)`

GetTokensUsedOk returns a tuple with the TokensUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensUsed

`func (o *DocumentStatusInfoDetails) SetTokensUsed(v float32)`

SetTokensUsed sets TokensUsed field to given value.

### HasTokensUsed

`func (o *DocumentStatusInfoDetails) HasTokensUsed() bool`

HasTokensUsed returns a boolean if a field has been set.

### GetTokensEstimated

`func (o *DocumentStatusInfoDetails) GetTokensEstimated() float32`

GetTokensEstimated returns the TokensEstimated field if non-nil, zero value otherwise.

### GetTokensEstimatedOk

`func (o *DocumentStatusInfoDetails) GetTokensEstimatedOk() (*float32, bool)`

GetTokensEstimatedOk returns a tuple with the TokensEstimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensEstimated

`func (o *DocumentStatusInfoDetails) SetTokensEstimated(v float32)`

SetTokensEstimated sets TokensEstimated field to given value.

### HasTokensEstimated

`func (o *DocumentStatusInfoDetails) HasTokensEstimated() bool`

HasTokensEstimated returns a boolean if a field has been set.

### GetChunksCount

`func (o *DocumentStatusInfoDetails) GetChunksCount() float32`

GetChunksCount returns the ChunksCount field if non-nil, zero value otherwise.

### GetChunksCountOk

`func (o *DocumentStatusInfoDetails) GetChunksCountOk() (*float32, bool)`

GetChunksCountOk returns a tuple with the ChunksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunksCount

`func (o *DocumentStatusInfoDetails) SetChunksCount(v float32)`

SetChunksCount sets ChunksCount field to given value.

### HasChunksCount

`func (o *DocumentStatusInfoDetails) HasChunksCount() bool`

HasChunksCount returns a boolean if a field has been set.

### GetNeeded

`func (o *DocumentStatusInfoDetails) GetNeeded() float32`

GetNeeded returns the Needed field if non-nil, zero value otherwise.

### GetNeededOk

`func (o *DocumentStatusInfoDetails) GetNeededOk() (*float32, bool)`

GetNeededOk returns a tuple with the Needed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeeded

`func (o *DocumentStatusInfoDetails) SetNeeded(v float32)`

SetNeeded sets Needed field to given value.

### HasNeeded

`func (o *DocumentStatusInfoDetails) HasNeeded() bool`

HasNeeded returns a boolean if a field has been set.

### GetAvailable

`func (o *DocumentStatusInfoDetails) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *DocumentStatusInfoDetails) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *DocumentStatusInfoDetails) SetAvailable(v float32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *DocumentStatusInfoDetails) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetEstimated

`func (o *DocumentStatusInfoDetails) GetEstimated() float32`

GetEstimated returns the Estimated field if non-nil, zero value otherwise.

### GetEstimatedOk

`func (o *DocumentStatusInfoDetails) GetEstimatedOk() (*float32, bool)`

GetEstimatedOk returns a tuple with the Estimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimated

`func (o *DocumentStatusInfoDetails) SetEstimated(v float32)`

SetEstimated sets Estimated field to given value.

### HasEstimated

`func (o *DocumentStatusInfoDetails) HasEstimated() bool`

HasEstimated returns a boolean if a field has been set.

### GetBalance

`func (o *DocumentStatusInfoDetails) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *DocumentStatusInfoDetails) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *DocumentStatusInfoDetails) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *DocumentStatusInfoDetails) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetMinRequired

`func (o *DocumentStatusInfoDetails) GetMinRequired() float32`

GetMinRequired returns the MinRequired field if non-nil, zero value otherwise.

### GetMinRequiredOk

`func (o *DocumentStatusInfoDetails) GetMinRequiredOk() (*float32, bool)`

GetMinRequiredOk returns a tuple with the MinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRequired

`func (o *DocumentStatusInfoDetails) SetMinRequired(v float32)`

SetMinRequired sets MinRequired field to given value.

### HasMinRequired

`func (o *DocumentStatusInfoDetails) HasMinRequired() bool`

HasMinRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


