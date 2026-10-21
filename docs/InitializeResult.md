# InitializeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtocolVersion** | **string** | Версия протокола: запрошенная клиентом, если сервер её поддерживает, иначе &#x60;2025-11-25&#x60; | 
**Capabilities** | **map[string]interface{}** | Возможности сервера | 
**ServerInfo** | [**InitializeResultServerInfo**](InitializeResultServerInfo.md) |  | 
**Instructions** | Pointer to **string** | Правила для модели: какой инструмент на какой вопрос отвечает. MCP-клиенты передают их модели автоматически | [optional] 

## Methods

### NewInitializeResult

`func NewInitializeResult(protocolVersion string, capabilities map[string]interface{}, serverInfo InitializeResultServerInfo, ) *InitializeResult`

NewInitializeResult instantiates a new InitializeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitializeResultWithDefaults

`func NewInitializeResultWithDefaults() *InitializeResult`

NewInitializeResultWithDefaults instantiates a new InitializeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocolVersion

`func (o *InitializeResult) GetProtocolVersion() string`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *InitializeResult) GetProtocolVersionOk() (*string, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *InitializeResult) SetProtocolVersion(v string)`

SetProtocolVersion sets ProtocolVersion field to given value.


### GetCapabilities

`func (o *InitializeResult) GetCapabilities() map[string]interface{}`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *InitializeResult) GetCapabilitiesOk() (*map[string]interface{}, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *InitializeResult) SetCapabilities(v map[string]interface{})`

SetCapabilities sets Capabilities field to given value.


### GetServerInfo

`func (o *InitializeResult) GetServerInfo() InitializeResultServerInfo`

GetServerInfo returns the ServerInfo field if non-nil, zero value otherwise.

### GetServerInfoOk

`func (o *InitializeResult) GetServerInfoOk() (*InitializeResultServerInfo, bool)`

GetServerInfoOk returns a tuple with the ServerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInfo

`func (o *InitializeResult) SetServerInfo(v InitializeResultServerInfo)`

SetServerInfo sets ServerInfo field to given value.


### GetInstructions

`func (o *InitializeResult) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *InitializeResult) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *InitializeResult) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *InitializeResult) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


