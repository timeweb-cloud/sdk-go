# CreateAiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название ключа | 
**Description** | Pointer to **string** | Описание ключа | [optional] 
**ProjectId** | Pointer to **int32** | ID проекта, к которому привязать ключ | [optional] 
**Expire** | **NullableTime** | Дата и время окончания действия ключа в формате ISO 8601. Поле обязательно: передайте &#x60;null&#x60;, чтобы ключ был бессрочным | 
**FallbackModels** | Pointer to **[]string** | Запасные модели в порядке приоритета, не больше двух и без повторов. Подходят только действующие текстовые модели каталога. Запасные модели применяются к любой модели, запрошенной этим ключом, и только в &#x60;/v1/chat/completions&#x60; и &#x60;/v1/responses&#x60;: если запрошенная модель не ответила, запрос по порядку повторяется на запасных. Запасная модель, совпадающая с запрошенной, пропускается. Резерв средств и проверка лимитов делаются по цене запрошенной модели, списание и учёт расхода — по цене ответившей: если запасная модель дороже, итоговое списание может превысить резерв и денежный лимит ключа. | [optional] 

## Methods

### NewCreateAiKey

`func NewCreateAiKey(name string, expire NullableTime, ) *CreateAiKey`

NewCreateAiKey instantiates a new CreateAiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAiKeyWithDefaults

`func NewCreateAiKeyWithDefaults() *CreateAiKey`

NewCreateAiKeyWithDefaults instantiates a new CreateAiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAiKey) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAiKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAiKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAiKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAiKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateAiKey) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateAiKey) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateAiKey) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateAiKey) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetExpire

`func (o *CreateAiKey) GetExpire() time.Time`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *CreateAiKey) GetExpireOk() (*time.Time, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *CreateAiKey) SetExpire(v time.Time)`

SetExpire sets Expire field to given value.


### SetExpireNil

`func (o *CreateAiKey) SetExpireNil(b bool)`

 SetExpireNil sets the value for Expire to be an explicit nil

### UnsetExpire
`func (o *CreateAiKey) UnsetExpire()`

UnsetExpire ensures that no value is present for Expire, not even an explicit nil
### GetFallbackModels

`func (o *CreateAiKey) GetFallbackModels() []string`

GetFallbackModels returns the FallbackModels field if non-nil, zero value otherwise.

### GetFallbackModelsOk

`func (o *CreateAiKey) GetFallbackModelsOk() (*[]string, bool)`

GetFallbackModelsOk returns a tuple with the FallbackModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackModels

`func (o *CreateAiKey) SetFallbackModels(v []string)`

SetFallbackModels sets FallbackModels field to given value.

### HasFallbackModels

`func (o *CreateAiKey) HasFallbackModels() bool`

HasFallbackModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


