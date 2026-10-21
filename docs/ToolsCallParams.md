# ToolsCallParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя инструмента | 
**Arguments** | Pointer to **map[string]interface{}** | Аргументы по &#x60;inputSchema&#x60; инструмента. Во втором вызове изменяющего инструмента добавьте &#x60;confirm_token&#x60; | [optional] 

## Methods

### NewToolsCallParams

`func NewToolsCallParams(name string, ) *ToolsCallParams`

NewToolsCallParams instantiates a new ToolsCallParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolsCallParamsWithDefaults

`func NewToolsCallParamsWithDefaults() *ToolsCallParams`

NewToolsCallParamsWithDefaults instantiates a new ToolsCallParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ToolsCallParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolsCallParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolsCallParams) SetName(v string)`

SetName sets Name field to given value.


### GetArguments

`func (o *ToolsCallParams) GetArguments() map[string]interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ToolsCallParams) GetArgumentsOk() (*map[string]interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ToolsCallParams) SetArguments(v map[string]interface{})`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ToolsCallParams) HasArguments() bool`

HasArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


