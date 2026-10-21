# AiGatewayCompletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID ответа | [optional] 
**Object** | Pointer to **string** | Тип объекта | [optional] 
**Created** | Pointer to **int32** | Время создания, Unix-время в секундах | [optional] 
**Model** | Pointer to **string** | Модель, которая ответила | [optional] 
**Choices** | Pointer to [**[]AiGatewayCompletionChoicesInner**](AiGatewayCompletionChoicesInner.md) | Варианты продолжения | [optional] 
**Usage** | Pointer to [**AiGatewayUsage**](AiGatewayUsage.md) |  | [optional] 

## Methods

### NewAiGatewayCompletion

`func NewAiGatewayCompletion() *AiGatewayCompletion`

NewAiGatewayCompletion instantiates a new AiGatewayCompletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayCompletionWithDefaults

`func NewAiGatewayCompletionWithDefaults() *AiGatewayCompletion`

NewAiGatewayCompletionWithDefaults instantiates a new AiGatewayCompletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiGatewayCompletion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiGatewayCompletion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiGatewayCompletion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiGatewayCompletion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *AiGatewayCompletion) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiGatewayCompletion) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiGatewayCompletion) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiGatewayCompletion) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *AiGatewayCompletion) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AiGatewayCompletion) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AiGatewayCompletion) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AiGatewayCompletion) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModel

`func (o *AiGatewayCompletion) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiGatewayCompletion) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiGatewayCompletion) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiGatewayCompletion) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetChoices

`func (o *AiGatewayCompletion) GetChoices() []AiGatewayCompletionChoicesInner`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *AiGatewayCompletion) GetChoicesOk() (*[]AiGatewayCompletionChoicesInner, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *AiGatewayCompletion) SetChoices(v []AiGatewayCompletionChoicesInner)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *AiGatewayCompletion) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetUsage

`func (o *AiGatewayCompletion) GetUsage() AiGatewayUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiGatewayCompletion) GetUsageOk() (*AiGatewayUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiGatewayCompletion) SetUsage(v AiGatewayUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiGatewayCompletion) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


