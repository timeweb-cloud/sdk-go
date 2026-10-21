# AiGatewayModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID модели — передавайте его в поле &#x60;model&#x60; | [optional] 
**Object** | Pointer to **string** | Тип объекта | [optional] 
**Created** | Pointer to **int32** | Время добавления, Unix-время в секундах | [optional] 
**OwnedBy** | Pointer to **string** | Владелец модели | [optional] 
**Mode** | Pointer to **string** | Назначение модели: &#x60;chat&#x60;, &#x60;responses&#x60;, &#x60;embedding&#x60;, &#x60;image_generation&#x60;, &#x60;audio_speech&#x60;, &#x60;audio_transcription&#x60; и др. У части моделей не задано | [optional] 
**MaxInputTokens** | Pointer to **int32** | Контекстное окно, если известно | [optional] 
**MaxOutputTokens** | Pointer to **int32** | Максимум выходных токенов, если известно | [optional] 

## Methods

### NewAiGatewayModel

`func NewAiGatewayModel() *AiGatewayModel`

NewAiGatewayModel instantiates a new AiGatewayModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayModelWithDefaults

`func NewAiGatewayModelWithDefaults() *AiGatewayModel`

NewAiGatewayModelWithDefaults instantiates a new AiGatewayModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiGatewayModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiGatewayModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiGatewayModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiGatewayModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *AiGatewayModel) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiGatewayModel) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiGatewayModel) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiGatewayModel) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *AiGatewayModel) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AiGatewayModel) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AiGatewayModel) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AiGatewayModel) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetOwnedBy

`func (o *AiGatewayModel) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *AiGatewayModel) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *AiGatewayModel) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *AiGatewayModel) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetMode

`func (o *AiGatewayModel) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AiGatewayModel) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AiGatewayModel) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AiGatewayModel) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMaxInputTokens

`func (o *AiGatewayModel) GetMaxInputTokens() int32`

GetMaxInputTokens returns the MaxInputTokens field if non-nil, zero value otherwise.

### GetMaxInputTokensOk

`func (o *AiGatewayModel) GetMaxInputTokensOk() (*int32, bool)`

GetMaxInputTokensOk returns a tuple with the MaxInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInputTokens

`func (o *AiGatewayModel) SetMaxInputTokens(v int32)`

SetMaxInputTokens sets MaxInputTokens field to given value.

### HasMaxInputTokens

`func (o *AiGatewayModel) HasMaxInputTokens() bool`

HasMaxInputTokens returns a boolean if a field has been set.

### GetMaxOutputTokens

`func (o *AiGatewayModel) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *AiGatewayModel) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *AiGatewayModel) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *AiGatewayModel) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


