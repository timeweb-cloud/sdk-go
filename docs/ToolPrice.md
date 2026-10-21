# ToolPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID шаблона инструмента | 
**Name** | **string** | Название инструмента | 
**Cost** | **float32** | Цена за единицу использования инструмента: в рублях для &#x60;locale&#x3D;ru&#x60;, в долларах США для &#x60;locale&#x3D;en&#x60; | 

## Methods

### NewToolPrice

`func NewToolPrice(id int32, name string, cost float32, ) *ToolPrice`

NewToolPrice instantiates a new ToolPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolPriceWithDefaults

`func NewToolPriceWithDefaults() *ToolPrice`

NewToolPriceWithDefaults instantiates a new ToolPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ToolPrice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolPrice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolPrice) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ToolPrice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolPrice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolPrice) SetName(v string)`

SetName sets Name field to given value.


### GetCost

`func (o *ToolPrice) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ToolPrice) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ToolPrice) SetCost(v float32)`

SetCost sets Cost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


