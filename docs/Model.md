# Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID модели | 
**ProviderId** | **int32** | ID провайдера модели | 
**Name** | **string** | Семейство модели | 
**PublicName** | **string** | Отображаемое название модели | 
**ModelName** | **string** | Идентификатор модели с префиксом провайдера — передаётся в поле &#x60;model&#x60; запросов к AI Gateway | 
**Type** | **string** | Тип модели: - &#x60;llm&#x60; — языковая модель - &#x60;hosted-llm&#x60; — языковая модель, размещённая в Timeweb Cloud - &#x60;embedding&#x60; — модель эмбеддингов - &#x60;image&#x60; — генерация изображений - &#x60;audio&#x60; — работа с аудио - &#x60;video&#x60; — генерация видео - &#x60;moderation&#x60; — модерация контента - &#x60;system-one&#x60; — служебная модель с собственным API | 
**Version** | **string** | Версия модели | 
**IsDeprecated** | **bool** | Модель устарела и недоступна для выбора в новых агентах | 
**IsStopped** | **bool** | Поддержка модели остановлена: запросы к ней не выполняются | 
**DeprecationDate** | **NullableTime** | Дата отключения модели у провайдера. &#x60;null&#x60;, если отключение не запланировано | 
**IsReasoning** | **bool** | Модель поддерживает режим рассуждений | 
**ParamsInfo** | [**map[string]ModelParamInfo**](ModelParamInfo.md) | Параметры генерации, которые можно настроить у агента с этой моделью. Ключ — имя параметра (&#x60;temperature&#x60;, &#x60;top_p&#x60;, &#x60;max_tokens&#x60;, &#x60;reasoning_effort&#x60; и т.д.). &#x60;null&#x60;, если параметры не заданы | 

## Methods

### NewModel

`func NewModel(id int32, providerId int32, name string, publicName string, modelName string, type_ string, version string, isDeprecated bool, isStopped bool, deprecationDate NullableTime, isReasoning bool, paramsInfo map[string]ModelParamInfo, ) *Model`

NewModel instantiates a new Model object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelWithDefaults

`func NewModelWithDefaults() *Model`

NewModelWithDefaults instantiates a new Model object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Model) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Model) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Model) SetId(v int32)`

SetId sets Id field to given value.


### GetProviderId

`func (o *Model) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Model) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Model) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetName

`func (o *Model) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Model) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Model) SetName(v string)`

SetName sets Name field to given value.


### GetPublicName

`func (o *Model) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *Model) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *Model) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.


### GetModelName

`func (o *Model) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Model) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Model) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetType

`func (o *Model) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Model) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Model) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *Model) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Model) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Model) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetIsDeprecated

`func (o *Model) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *Model) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *Model) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.


### GetIsStopped

`func (o *Model) GetIsStopped() bool`

GetIsStopped returns the IsStopped field if non-nil, zero value otherwise.

### GetIsStoppedOk

`func (o *Model) GetIsStoppedOk() (*bool, bool)`

GetIsStoppedOk returns a tuple with the IsStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStopped

`func (o *Model) SetIsStopped(v bool)`

SetIsStopped sets IsStopped field to given value.


### GetDeprecationDate

`func (o *Model) GetDeprecationDate() time.Time`

GetDeprecationDate returns the DeprecationDate field if non-nil, zero value otherwise.

### GetDeprecationDateOk

`func (o *Model) GetDeprecationDateOk() (*time.Time, bool)`

GetDeprecationDateOk returns a tuple with the DeprecationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationDate

`func (o *Model) SetDeprecationDate(v time.Time)`

SetDeprecationDate sets DeprecationDate field to given value.


### SetDeprecationDateNil

`func (o *Model) SetDeprecationDateNil(b bool)`

 SetDeprecationDateNil sets the value for DeprecationDate to be an explicit nil

### UnsetDeprecationDate
`func (o *Model) UnsetDeprecationDate()`

UnsetDeprecationDate ensures that no value is present for DeprecationDate, not even an explicit nil
### GetIsReasoning

`func (o *Model) GetIsReasoning() bool`

GetIsReasoning returns the IsReasoning field if non-nil, zero value otherwise.

### GetIsReasoningOk

`func (o *Model) GetIsReasoningOk() (*bool, bool)`

GetIsReasoningOk returns a tuple with the IsReasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReasoning

`func (o *Model) SetIsReasoning(v bool)`

SetIsReasoning sets IsReasoning field to given value.


### GetParamsInfo

`func (o *Model) GetParamsInfo() map[string]ModelParamInfo`

GetParamsInfo returns the ParamsInfo field if non-nil, zero value otherwise.

### GetParamsInfoOk

`func (o *Model) GetParamsInfoOk() (*map[string]ModelParamInfo, bool)`

GetParamsInfoOk returns a tuple with the ParamsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsInfo

`func (o *Model) SetParamsInfo(v map[string]ModelParamInfo)`

SetParamsInfo sets ParamsInfo field to given value.


### SetParamsInfoNil

`func (o *Model) SetParamsInfoNil(b bool)`

 SetParamsInfoNil sets the value for ParamsInfo to be an explicit nil

### UnsetParamsInfo
`func (o *Model) UnsetParamsInfo()`

UnsetParamsInfo ensures that no value is present for ParamsInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


