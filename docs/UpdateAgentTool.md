# UpdateAgentTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedMethods** | **[]string** | Методы, которые агенту разрешено вызывать: имена из &#x60;methods[].name&#x60; инструмента. Список заменяется целиком | 

## Methods

### NewUpdateAgentTool

`func NewUpdateAgentTool(allowedMethods []string, ) *UpdateAgentTool`

NewUpdateAgentTool instantiates a new UpdateAgentTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAgentToolWithDefaults

`func NewUpdateAgentToolWithDefaults() *UpdateAgentTool`

NewUpdateAgentToolWithDefaults instantiates a new UpdateAgentTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedMethods

`func (o *UpdateAgentTool) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *UpdateAgentTool) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *UpdateAgentTool) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


