# ConnectAgentTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolId** | **int32** | ID инструмента | 
**AllowedMethods** | **[]string** | Методы, которые агенту разрешено вызывать: имена из &#x60;methods[].name&#x60; инструмента. Пустой массив тоже допустим | 

## Methods

### NewConnectAgentTool

`func NewConnectAgentTool(toolId int32, allowedMethods []string, ) *ConnectAgentTool`

NewConnectAgentTool instantiates a new ConnectAgentTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectAgentToolWithDefaults

`func NewConnectAgentToolWithDefaults() *ConnectAgentTool`

NewConnectAgentToolWithDefaults instantiates a new ConnectAgentTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolId

`func (o *ConnectAgentTool) GetToolId() int32`

GetToolId returns the ToolId field if non-nil, zero value otherwise.

### GetToolIdOk

`func (o *ConnectAgentTool) GetToolIdOk() (*int32, bool)`

GetToolIdOk returns a tuple with the ToolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolId

`func (o *ConnectAgentTool) SetToolId(v int32)`

SetToolId sets ToolId field to given value.


### GetAllowedMethods

`func (o *ConnectAgentTool) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *ConnectAgentTool) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *ConnectAgentTool) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


