# AgentToolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolId** | **int32** | ID инструмента | 
**AllowedMethods** | **[]string** | Методы инструмента, которые разрешено вызывать агенту | 

## Methods

### NewAgentToolsInner

`func NewAgentToolsInner(toolId int32, allowedMethods []string, ) *AgentToolsInner`

NewAgentToolsInner instantiates a new AgentToolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentToolsInnerWithDefaults

`func NewAgentToolsInnerWithDefaults() *AgentToolsInner`

NewAgentToolsInnerWithDefaults instantiates a new AgentToolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolId

`func (o *AgentToolsInner) GetToolId() int32`

GetToolId returns the ToolId field if non-nil, zero value otherwise.

### GetToolIdOk

`func (o *AgentToolsInner) GetToolIdOk() (*int32, bool)`

GetToolIdOk returns a tuple with the ToolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolId

`func (o *AgentToolsInner) SetToolId(v int32)`

SetToolId sets ToolId field to given value.


### GetAllowedMethods

`func (o *AgentToolsInner) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *AgentToolsInner) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *AgentToolsInner) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


