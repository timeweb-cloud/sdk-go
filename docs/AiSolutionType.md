# AiSolutionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип решения | 
**Name** | **string** | Название решения | 
**Description** | **string** | Описание решения в HTML | 
**OsIds** | **[]int32** | ID операционной системы, на которой работает решение (один элемент). Пустой массив, если этой ОС нет в каталоге ОС аккаунта | 
**Installations** | **int32** | Количество установок. Всегда &#x60;0&#x60;, оставлено для обратной совместимости | 
**Requirements** | Pointer to [**AiSolutionTypeRequirements**](AiSolutionTypeRequirements.md) |  | [optional] 

## Methods

### NewAiSolutionType

`func NewAiSolutionType(type_ string, name string, description string, osIds []int32, installations int32, ) *AiSolutionType`

NewAiSolutionType instantiates a new AiSolutionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSolutionTypeWithDefaults

`func NewAiSolutionTypeWithDefaults() *AiSolutionType`

NewAiSolutionTypeWithDefaults instantiates a new AiSolutionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AiSolutionType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiSolutionType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiSolutionType) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AiSolutionType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiSolutionType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiSolutionType) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AiSolutionType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AiSolutionType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AiSolutionType) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOsIds

`func (o *AiSolutionType) GetOsIds() []int32`

GetOsIds returns the OsIds field if non-nil, zero value otherwise.

### GetOsIdsOk

`func (o *AiSolutionType) GetOsIdsOk() (*[]int32, bool)`

GetOsIdsOk returns a tuple with the OsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsIds

`func (o *AiSolutionType) SetOsIds(v []int32)`

SetOsIds sets OsIds field to given value.


### GetInstallations

`func (o *AiSolutionType) GetInstallations() int32`

GetInstallations returns the Installations field if non-nil, zero value otherwise.

### GetInstallationsOk

`func (o *AiSolutionType) GetInstallationsOk() (*int32, bool)`

GetInstallationsOk returns a tuple with the Installations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallations

`func (o *AiSolutionType) SetInstallations(v int32)`

SetInstallations sets Installations field to given value.


### GetRequirements

`func (o *AiSolutionType) GetRequirements() AiSolutionTypeRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *AiSolutionType) GetRequirementsOk() (*AiSolutionTypeRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *AiSolutionType) SetRequirements(v AiSolutionTypeRequirements)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *AiSolutionType) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


