# Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Уникальный идентификатор модели | 
**ProviderId** | **float32** | ID провайдера, который предоставляет модель | 
**Name** | **string** | Название модели | 
**PublicName** | **string** | Публичное имя модели | 
**Type** | **string** | Тип модели (llm - языковая модель, embedding - модель для эмбеддингов) | 
**IsDeprecated** | **bool** | Признак, что модель устарела | 
**IsReasoning** | **bool** | Признак поддержки режима рассуждения | 
**Version** | **string** | Версия модели | 
**ParamsInfo** | Pointer to [**NullableModelParamsInfo**](ModelParamsInfo.md) |  | [optional] 

## Methods

### NewModel

`func NewModel(id float32, providerId float32, name string, publicName string, type_ string, isDeprecated bool, isReasoning bool, version string, ) *Model`

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

`func (o *Model) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Model) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Model) SetId(v float32)`

SetId sets Id field to given value.


### GetProviderId

`func (o *Model) GetProviderId() float32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Model) GetProviderIdOk() (*float32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Model) SetProviderId(v float32)`

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


### GetParamsInfo

`func (o *Model) GetParamsInfo() ModelParamsInfo`

GetParamsInfo returns the ParamsInfo field if non-nil, zero value otherwise.

### GetParamsInfoOk

`func (o *Model) GetParamsInfoOk() (*ModelParamsInfo, bool)`

GetParamsInfoOk returns a tuple with the ParamsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsInfo

`func (o *Model) SetParamsInfo(v ModelParamsInfo)`

SetParamsInfo sets ParamsInfo field to given value.

### HasParamsInfo

`func (o *Model) HasParamsInfo() bool`

HasParamsInfo returns a boolean if a field has been set.

### SetParamsInfoNil

`func (o *Model) SetParamsInfoNil(b bool)`

 SetParamsInfoNil sets the value for ParamsInfo to be an explicit nil

### UnsetParamsInfo
`func (o *Model) UnsetParamsInfo()`

UnsetParamsInfo ensures that no value is present for ParamsInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


