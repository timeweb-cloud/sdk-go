# ModelV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Уникальный идентификатор модели | 
**ProviderId** | **float32** | ID провайдера, который предоставляет модель | 
**Name** | **string** | Название модели | 
**PublicName** | **string** | Публичное имя модели | 
**ModelName** | **string** | Идентификатор модели, который передаётся в поле &#x60;model&#x60; тела запроса | 
**Type** | **string** | Тип модели (llm — языковая модель, hosted-llm — self-hosted языковая модель, embedding — модель для эмбеддингов, image — генерация изображений, audio — работа с аудио, moderation — модерация контента) | 
**Version** | **string** | Версия модели | 
**IsDeprecated** | **bool** | Признак, что модель устарела | 
**IsStopped** | **bool** | Признак, что поддержка модели остановлена | 
**DeprecationDate** | Pointer to **NullableTime** | Дата отключения модели у провайдера. &#x60;null&#x60;, если отключение не запланировано | [optional] 
**CreatedAt** | **time.Time** | Дата и время добавления модели | 
**UpdatedAt** | **time.Time** | Дата и время последнего обновления модели | 
**ParamsInfo** | Pointer to **map[string]interface{}** | Информация о доступных параметрах модели с их ограничениями. Ключ — название параметра, значение — его тип, ограничения и значение по умолчанию | [optional] 
**ParameterValues** | **map[string]interface{}** | Характеристики модели. Ключ — код характеристики, значение — число, строка, булево значение или массив строк.  Основные коды: - &#x60;cost_in&#x60; — цена за 1 входной токен в рублях - &#x60;cost_out&#x60; — цена за 1 выходной токен в рублях - &#x60;context_window&#x60; — размер контекстного окна в токенах - &#x60;max_output_tokens&#x60; — максимальное количество токенов в ответе - &#x60;is_reasoning&#x60; — поддержка режима рассуждений - &#x60;supports_function_calling&#x60; — поддержка function calling - &#x60;supported_input_modalities&#x60; — поддерживаемые входные модальности (например, &#x60;text&#x60;, &#x60;image&#x60;) - &#x60;supported_output_modalities&#x60; — поддерживаемые выходные модальности - &#x60;avg_speed&#x60; — средняя скорость генерации в токенах в секунду - &#x60;aa_intelligence_index&#x60;, &#x60;aa_coding_index&#x60;, &#x60;aa_agentic_index&#x60; — индексы качества Artificial Analysis - &#x60;release_date&#x60; — дата релиза модели | 

## Methods

### NewModelV3

`func NewModelV3(id float32, providerId float32, name string, publicName string, modelName string, type_ string, version string, isDeprecated bool, isStopped bool, createdAt time.Time, updatedAt time.Time, parameterValues map[string]interface{}, ) *ModelV3`

NewModelV3 instantiates a new ModelV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelV3WithDefaults

`func NewModelV3WithDefaults() *ModelV3`

NewModelV3WithDefaults instantiates a new ModelV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelV3) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelV3) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelV3) SetId(v float32)`

SetId sets Id field to given value.


### GetProviderId

`func (o *ModelV3) GetProviderId() float32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ModelV3) GetProviderIdOk() (*float32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ModelV3) SetProviderId(v float32)`

SetProviderId sets ProviderId field to given value.


### GetName

`func (o *ModelV3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelV3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelV3) SetName(v string)`

SetName sets Name field to given value.


### GetPublicName

`func (o *ModelV3) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *ModelV3) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *ModelV3) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.


### GetModelName

`func (o *ModelV3) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ModelV3) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ModelV3) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetType

`func (o *ModelV3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelV3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelV3) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *ModelV3) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelV3) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelV3) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetIsDeprecated

`func (o *ModelV3) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *ModelV3) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *ModelV3) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.


### GetIsStopped

`func (o *ModelV3) GetIsStopped() bool`

GetIsStopped returns the IsStopped field if non-nil, zero value otherwise.

### GetIsStoppedOk

`func (o *ModelV3) GetIsStoppedOk() (*bool, bool)`

GetIsStoppedOk returns a tuple with the IsStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStopped

`func (o *ModelV3) SetIsStopped(v bool)`

SetIsStopped sets IsStopped field to given value.


### GetDeprecationDate

`func (o *ModelV3) GetDeprecationDate() time.Time`

GetDeprecationDate returns the DeprecationDate field if non-nil, zero value otherwise.

### GetDeprecationDateOk

`func (o *ModelV3) GetDeprecationDateOk() (*time.Time, bool)`

GetDeprecationDateOk returns a tuple with the DeprecationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationDate

`func (o *ModelV3) SetDeprecationDate(v time.Time)`

SetDeprecationDate sets DeprecationDate field to given value.

### HasDeprecationDate

`func (o *ModelV3) HasDeprecationDate() bool`

HasDeprecationDate returns a boolean if a field has been set.

### SetDeprecationDateNil

`func (o *ModelV3) SetDeprecationDateNil(b bool)`

 SetDeprecationDateNil sets the value for DeprecationDate to be an explicit nil

### UnsetDeprecationDate
`func (o *ModelV3) UnsetDeprecationDate()`

UnsetDeprecationDate ensures that no value is present for DeprecationDate, not even an explicit nil
### GetCreatedAt

`func (o *ModelV3) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelV3) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelV3) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ModelV3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelV3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelV3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetParamsInfo

`func (o *ModelV3) GetParamsInfo() map[string]interface{}`

GetParamsInfo returns the ParamsInfo field if non-nil, zero value otherwise.

### GetParamsInfoOk

`func (o *ModelV3) GetParamsInfoOk() (*map[string]interface{}, bool)`

GetParamsInfoOk returns a tuple with the ParamsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsInfo

`func (o *ModelV3) SetParamsInfo(v map[string]interface{})`

SetParamsInfo sets ParamsInfo field to given value.

### HasParamsInfo

`func (o *ModelV3) HasParamsInfo() bool`

HasParamsInfo returns a boolean if a field has been set.

### SetParamsInfoNil

`func (o *ModelV3) SetParamsInfoNil(b bool)`

 SetParamsInfoNil sets the value for ParamsInfo to be an explicit nil

### UnsetParamsInfo
`func (o *ModelV3) UnsetParamsInfo()`

UnsetParamsInfo ensures that no value is present for ParamsInfo, not even an explicit nil
### GetParameterValues

`func (o *ModelV3) GetParameterValues() map[string]interface{}`

GetParameterValues returns the ParameterValues field if non-nil, zero value otherwise.

### GetParameterValuesOk

`func (o *ModelV3) GetParameterValuesOk() (*map[string]interface{}, bool)`

GetParameterValuesOk returns a tuple with the ParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterValues

`func (o *ModelV3) SetParameterValues(v map[string]interface{})`

SetParameterValues sets ParameterValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


