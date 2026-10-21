# ModelPriceTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AboveContextTokens** | **int32** | Порог входного контекста в токенах | 
**CostIn** | **float32** | Цена 1 входного токена в рублях | 
**CostOut** | **float32** | Цена 1 выходного токена в рублях | 

## Methods

### NewModelPriceTier

`func NewModelPriceTier(aboveContextTokens int32, costIn float32, costOut float32, ) *ModelPriceTier`

NewModelPriceTier instantiates a new ModelPriceTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPriceTierWithDefaults

`func NewModelPriceTierWithDefaults() *ModelPriceTier`

NewModelPriceTierWithDefaults instantiates a new ModelPriceTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAboveContextTokens

`func (o *ModelPriceTier) GetAboveContextTokens() int32`

GetAboveContextTokens returns the AboveContextTokens field if non-nil, zero value otherwise.

### GetAboveContextTokensOk

`func (o *ModelPriceTier) GetAboveContextTokensOk() (*int32, bool)`

GetAboveContextTokensOk returns a tuple with the AboveContextTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboveContextTokens

`func (o *ModelPriceTier) SetAboveContextTokens(v int32)`

SetAboveContextTokens sets AboveContextTokens field to given value.


### GetCostIn

`func (o *ModelPriceTier) GetCostIn() float32`

GetCostIn returns the CostIn field if non-nil, zero value otherwise.

### GetCostInOk

`func (o *ModelPriceTier) GetCostInOk() (*float32, bool)`

GetCostInOk returns a tuple with the CostIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostIn

`func (o *ModelPriceTier) SetCostIn(v float32)`

SetCostIn sets CostIn field to given value.


### GetCostOut

`func (o *ModelPriceTier) GetCostOut() float32`

GetCostOut returns the CostOut field if non-nil, zero value otherwise.

### GetCostOutOk

`func (o *ModelPriceTier) GetCostOutOk() (*float32, bool)`

GetCostOutOk returns a tuple with the CostOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOut

`func (o *ModelPriceTier) SetCostOut(v float32)`

SetCostOut sets CostOut field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


