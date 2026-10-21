# UpdateAiSolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название решения | [optional] 
**Description** | Pointer to **NullableString** | Описание решения. &#x60;null&#x60; очищает описание | [optional] 
**ModelId** | Pointer to **int32** | ID новой модели. Изменение записывается в конфигурацию агента на сервере решения, поэтому сервер должен быть включён и полностью развёрнут | [optional] 

## Methods

### NewUpdateAiSolution

`func NewUpdateAiSolution() *UpdateAiSolution`

NewUpdateAiSolution instantiates a new UpdateAiSolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAiSolutionWithDefaults

`func NewUpdateAiSolutionWithDefaults() *UpdateAiSolution`

NewUpdateAiSolutionWithDefaults instantiates a new UpdateAiSolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAiSolution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAiSolution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAiSolution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAiSolution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAiSolution) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAiSolution) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAiSolution) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAiSolution) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateAiSolution) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateAiSolution) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetModelId

`func (o *UpdateAiSolution) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *UpdateAiSolution) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *UpdateAiSolution) SetModelId(v int32)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *UpdateAiSolution) HasModelId() bool`

HasModelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


