# ModelParamInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип значения параметра | 
**Description** | **string** | Описание параметра (на английском) | 
**Min** | Pointer to **float32** | Минимальное значение, для &#x60;type: number&#x60; | [optional] 
**Max** | Pointer to **float32** | Максимальное значение, для &#x60;type: number&#x60; | [optional] 
**Enum** | Pointer to **[]string** | Допустимые значения, для &#x60;type: enum&#x60; | [optional] 
**Default** | **float32** | Значение по умолчанию | 

## Methods

### NewModelParamInfo

`func NewModelParamInfo(type_ string, description string, default_ float32, ) *ModelParamInfo`

NewModelParamInfo instantiates a new ModelParamInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelParamInfoWithDefaults

`func NewModelParamInfoWithDefaults() *ModelParamInfo`

NewModelParamInfoWithDefaults instantiates a new ModelParamInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModelParamInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelParamInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelParamInfo) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *ModelParamInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelParamInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelParamInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMin

`func (o *ModelParamInfo) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ModelParamInfo) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ModelParamInfo) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *ModelParamInfo) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *ModelParamInfo) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ModelParamInfo) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ModelParamInfo) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *ModelParamInfo) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetEnum

`func (o *ModelParamInfo) GetEnum() []string`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *ModelParamInfo) GetEnumOk() (*[]string, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *ModelParamInfo) SetEnum(v []string)`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *ModelParamInfo) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### GetDefault

`func (o *ModelParamInfo) GetDefault() float32`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ModelParamInfo) GetDefaultOk() (*float32, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ModelParamInfo) SetDefault(v float32)`

SetDefault sets Default field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


