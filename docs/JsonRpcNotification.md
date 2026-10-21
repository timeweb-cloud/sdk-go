# JsonRpcNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jsonrpc** | **string** | Версия протокола JSON-RPC, всегда &#x60;2.0&#x60; | 
**Method** | **string** | Метод уведомления | 
**Params** | Pointer to **map[string]interface{}** | Параметры уведомления | [optional] 

## Methods

### NewJsonRpcNotification

`func NewJsonRpcNotification(jsonrpc string, method string, ) *JsonRpcNotification`

NewJsonRpcNotification instantiates a new JsonRpcNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonRpcNotificationWithDefaults

`func NewJsonRpcNotificationWithDefaults() *JsonRpcNotification`

NewJsonRpcNotificationWithDefaults instantiates a new JsonRpcNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonrpc

`func (o *JsonRpcNotification) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *JsonRpcNotification) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *JsonRpcNotification) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.


### GetMethod

`func (o *JsonRpcNotification) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *JsonRpcNotification) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *JsonRpcNotification) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetParams

`func (o *JsonRpcNotification) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *JsonRpcNotification) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *JsonRpcNotification) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *JsonRpcNotification) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


