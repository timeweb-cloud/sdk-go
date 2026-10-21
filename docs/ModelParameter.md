# ModelParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Код характеристики — ключ в &#x60;parameter_values&#x60; модели | 
**Name** | Pointer to **string** | Название характеристики | [optional] 
**Description** | Pointer to **NullableString** | Описание характеристики | [optional] 
**Type** | **string** | Тип значения | 
**Unit** | **NullableString** | Единица измерения (&#x60;tokens&#x60;, &#x60;tokens_per_second&#x60;, &#x60;currency_per_million&#x60;, ...). &#x60;null&#x60;, если нет | 
**Group** | **string** | Группа характеристик: &#x60;pricing&#x60;, &#x60;quality&#x60;, &#x60;performance&#x60;, &#x60;capabilities&#x60;, &#x60;limits&#x60;, &#x60;meta&#x60; | 
**SortOrder** | **int32** | Порядок отображения | 
**ComparisonDirection** | **string** | Какое значение лучше при сравнении моделей | 
**Format** | **string** | Подсказка для отображения значения. &#x60;currency_per_million_tokens&#x60; — значение в &#x60;parameter_values&#x60; указано за 1 токен, показывать его удобно за 1 млн токенов | 

## Methods

### NewModelParameter

`func NewModelParameter(code string, type_ string, unit NullableString, group string, sortOrder int32, comparisonDirection string, format string, ) *ModelParameter`

NewModelParameter instantiates a new ModelParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelParameterWithDefaults

`func NewModelParameterWithDefaults() *ModelParameter`

NewModelParameterWithDefaults instantiates a new ModelParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ModelParameter) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ModelParameter) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ModelParameter) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *ModelParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ModelParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelParameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelParameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ModelParameter) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ModelParameter) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *ModelParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelParameter) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *ModelParameter) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ModelParameter) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ModelParameter) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *ModelParameter) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *ModelParameter) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetGroup

`func (o *ModelParameter) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ModelParameter) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ModelParameter) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetSortOrder

`func (o *ModelParameter) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ModelParameter) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ModelParameter) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetComparisonDirection

`func (o *ModelParameter) GetComparisonDirection() string`

GetComparisonDirection returns the ComparisonDirection field if non-nil, zero value otherwise.

### GetComparisonDirectionOk

`func (o *ModelParameter) GetComparisonDirectionOk() (*string, bool)`

GetComparisonDirectionOk returns a tuple with the ComparisonDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonDirection

`func (o *ModelParameter) SetComparisonDirection(v string)`

SetComparisonDirection sets ComparisonDirection field to given value.


### GetFormat

`func (o *ModelParameter) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ModelParameter) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ModelParameter) SetFormat(v string)`

SetFormat sets Format field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


