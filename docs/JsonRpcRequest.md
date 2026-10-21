# JsonRpcRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jsonrpc** | **string** | Версия протокола JSON-RPC, всегда &#x60;2.0&#x60; | 
**Id** | **string** | ID запроса, выбранный клиентом; ответ приходит с тем же ID | 
**Method** | **string** | Метод MCP | 
**Params** | Pointer to [**InitializeParams**](InitializeParams.md) |  | [optional] 

## Methods

### NewJsonRpcRequest

`func NewJsonRpcRequest(jsonrpc string, id string, method string, ) *JsonRpcRequest`

NewJsonRpcRequest instantiates a new JsonRpcRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonRpcRequestWithDefaults

`func NewJsonRpcRequestWithDefaults() *JsonRpcRequest`

NewJsonRpcRequestWithDefaults instantiates a new JsonRpcRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonrpc

`func (o *JsonRpcRequest) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *JsonRpcRequest) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *JsonRpcRequest) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.


### GetId

`func (o *JsonRpcRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonRpcRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonRpcRequest) SetId(v string)`

SetId sets Id field to given value.


### GetMethod

`func (o *JsonRpcRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *JsonRpcRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *JsonRpcRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetParams

`func (o *JsonRpcRequest) GetParams() InitializeParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *JsonRpcRequest) GetParamsOk() (*InitializeParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *JsonRpcRequest) SetParams(v InitializeParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *JsonRpcRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


