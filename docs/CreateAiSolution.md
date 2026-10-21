# CreateAiSolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип решения. Доступные типы возвращает метод получения типов AI-решений | 
**Name** | **string** | Название решения. Используется и как имя сервера | 
**Description** | Pointer to **string** | Описание решения. Используется и как комментарий к серверу | [optional] 
**ModelId** | **int32** | ID модели, с которой будет работать решение. Подходят только действующие текстовые модели (&#x60;type: llm&#x60; в &#x60;GET /api/v3/cloud-ai/models&#x60;) | 
**PresetId** | **int32** | ID тарифа сервера из &#x60;GET /api/v1/presets/servers&#x60;. Тариф должен быть не ниже требований типа решения (при создании это не проверяется) и из той же локации, что и зона доступности. Баланса аккаунта должно хватать хотя бы на час работы тарифа | 
**ProjectId** | Pointer to **int32** | ID проекта, в который добавляются сервер и ключ AI Gateway решения | [optional] 
**AvailabilityZone** | Pointer to **string** | Зона доступности сервера. Если не указана, используется зона по умолчанию для региона аккаунта. Зона должна относиться к локации тарифа, а образ решения должен быть доступен в этой зоне | [optional] 

## Methods

### NewCreateAiSolution

`func NewCreateAiSolution(type_ string, name string, modelId int32, presetId int32, ) *CreateAiSolution`

NewCreateAiSolution instantiates a new CreateAiSolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAiSolutionWithDefaults

`func NewCreateAiSolutionWithDefaults() *CreateAiSolution`

NewCreateAiSolutionWithDefaults instantiates a new CreateAiSolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateAiSolution) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAiSolution) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAiSolution) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CreateAiSolution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAiSolution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAiSolution) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAiSolution) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAiSolution) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAiSolution) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAiSolution) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModelId

`func (o *CreateAiSolution) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CreateAiSolution) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CreateAiSolution) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetPresetId

`func (o *CreateAiSolution) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreateAiSolution) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreateAiSolution) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetProjectId

`func (o *CreateAiSolution) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateAiSolution) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateAiSolution) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateAiSolution) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *CreateAiSolution) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateAiSolution) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateAiSolution) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateAiSolution) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


