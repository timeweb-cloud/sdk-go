# AiSolutionAiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiKeyId** | **string** | ID ключа AI Gateway решения. При перевыпуске не меняется | 
**Key** | **string** | Новое значение ключа. Возвращается один раз: ключ уже записан в конфигурацию агента на сервере решения, повторно получить его нельзя | 

## Methods

### NewAiSolutionAiKey

`func NewAiSolutionAiKey(aiKeyId string, key string, ) *AiSolutionAiKey`

NewAiSolutionAiKey instantiates a new AiSolutionAiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSolutionAiKeyWithDefaults

`func NewAiSolutionAiKeyWithDefaults() *AiSolutionAiKey`

NewAiSolutionAiKeyWithDefaults instantiates a new AiSolutionAiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiKeyId

`func (o *AiSolutionAiKey) GetAiKeyId() string`

GetAiKeyId returns the AiKeyId field if non-nil, zero value otherwise.

### GetAiKeyIdOk

`func (o *AiSolutionAiKey) GetAiKeyIdOk() (*string, bool)`

GetAiKeyIdOk returns a tuple with the AiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiKeyId

`func (o *AiSolutionAiKey) SetAiKeyId(v string)`

SetAiKeyId sets AiKeyId field to given value.


### GetKey

`func (o *AiSolutionAiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AiSolutionAiKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AiSolutionAiKey) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


