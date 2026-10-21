# InitializeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtocolVersion** | **string** | Версия протокола MCP, которую поддерживает клиент | 
**Capabilities** | **map[string]interface{}** | Возможности клиента | 
**ClientInfo** | [**InitializeParamsClientInfo**](InitializeParamsClientInfo.md) |  | 

## Methods

### NewInitializeParams

`func NewInitializeParams(protocolVersion string, capabilities map[string]interface{}, clientInfo InitializeParamsClientInfo, ) *InitializeParams`

NewInitializeParams instantiates a new InitializeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitializeParamsWithDefaults

`func NewInitializeParamsWithDefaults() *InitializeParams`

NewInitializeParamsWithDefaults instantiates a new InitializeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocolVersion

`func (o *InitializeParams) GetProtocolVersion() string`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *InitializeParams) GetProtocolVersionOk() (*string, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *InitializeParams) SetProtocolVersion(v string)`

SetProtocolVersion sets ProtocolVersion field to given value.


### GetCapabilities

`func (o *InitializeParams) GetCapabilities() map[string]interface{}`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *InitializeParams) GetCapabilitiesOk() (*map[string]interface{}, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *InitializeParams) SetCapabilities(v map[string]interface{})`

SetCapabilities sets Capabilities field to given value.


### GetClientInfo

`func (o *InitializeParams) GetClientInfo() InitializeParamsClientInfo`

GetClientInfo returns the ClientInfo field if non-nil, zero value otherwise.

### GetClientInfoOk

`func (o *InitializeParams) GetClientInfoOk() (*InitializeParamsClientInfo, bool)`

GetClientInfoOk returns a tuple with the ClientInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientInfo

`func (o *InitializeParams) SetClientInfo(v InitializeParamsClientInfo)`

SetClientInfo sets ClientInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


