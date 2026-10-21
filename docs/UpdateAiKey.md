# UpdateAiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название ключа | [optional] 
**Description** | Pointer to **NullableString** | Описание ключа. &#x60;null&#x60; — очистить описание | [optional] 
**FallbackModels** | Pointer to **[]string** | Запасные модели в порядке приоритета, не больше двух и без повторов. Переданный список заменяет текущий целиком; &#x60;null&#x60; или &#x60;[]&#x60; — убрать запасные модели; поле не передано — не меняются. Запасные модели применяются к любой модели, запрошенной этим ключом, и только в &#x60;/v1/chat/completions&#x60; и &#x60;/v1/responses&#x60;: если запрошенная модель не ответила, запрос по порядку повторяется на запасных. Запасная модель, совпадающая с запрошенной, пропускается. Резерв средств и проверка лимитов делаются по цене запрошенной модели, списание и учёт расхода — по цене ответившей: если запасная модель дороже, итоговое списание может превысить резерв и денежный лимит ключа. | [optional] 

## Methods

### NewUpdateAiKey

`func NewUpdateAiKey() *UpdateAiKey`

NewUpdateAiKey instantiates a new UpdateAiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAiKeyWithDefaults

`func NewUpdateAiKeyWithDefaults() *UpdateAiKey`

NewUpdateAiKeyWithDefaults instantiates a new UpdateAiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAiKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAiKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAiKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAiKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAiKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAiKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateAiKey) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateAiKey) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFallbackModels

`func (o *UpdateAiKey) GetFallbackModels() []string`

GetFallbackModels returns the FallbackModels field if non-nil, zero value otherwise.

### GetFallbackModelsOk

`func (o *UpdateAiKey) GetFallbackModelsOk() (*[]string, bool)`

GetFallbackModelsOk returns a tuple with the FallbackModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackModels

`func (o *UpdateAiKey) SetFallbackModels(v []string)`

SetFallbackModels sets FallbackModels field to given value.

### HasFallbackModels

`func (o *UpdateAiKey) HasFallbackModels() bool`

HasFallbackModels returns a boolean if a field has been set.

### SetFallbackModelsNil

`func (o *UpdateAiKey) SetFallbackModelsNil(b bool)`

 SetFallbackModelsNil sets the value for FallbackModels to be an explicit nil

### UnsetFallbackModels
`func (o *UpdateAiKey) UnsetFallbackModels()`

UnsetFallbackModels ensures that no value is present for FallbackModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


