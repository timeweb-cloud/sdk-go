# ToolMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя метода. Именно его указывают в &#x60;allowed_methods&#x60; при подключении к агенту | 
**Description** | Pointer to **string** | Описание метода, полученное от MCP-сервера | [optional] 
**PublicName** | Pointer to **string** | Отображаемое название метода. Задаётся у некоторых шаблонов маркетплейса | [optional] 

## Methods

### NewToolMethod

`func NewToolMethod(name string, ) *ToolMethod`

NewToolMethod instantiates a new ToolMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolMethodWithDefaults

`func NewToolMethodWithDefaults() *ToolMethod`

NewToolMethodWithDefaults instantiates a new ToolMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ToolMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolMethod) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ToolMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublicName

`func (o *ToolMethod) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *ToolMethod) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *ToolMethod) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *ToolMethod) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


