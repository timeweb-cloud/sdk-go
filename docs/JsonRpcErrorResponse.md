# JsonRpcErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jsonrpc** | **string** | Версия протокола JSON-RPC, всегда &#x60;2.0&#x60; | 
**Id** | **NullableString** | ID запроса; &#x60;null&#x60;, если запрос не удалось прочитать | 
**Error** | [**JsonRpcError**](JsonRpcError.md) |  | 

## Methods

### NewJsonRpcErrorResponse

`func NewJsonRpcErrorResponse(jsonrpc string, id NullableString, error_ JsonRpcError, ) *JsonRpcErrorResponse`

NewJsonRpcErrorResponse instantiates a new JsonRpcErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonRpcErrorResponseWithDefaults

`func NewJsonRpcErrorResponseWithDefaults() *JsonRpcErrorResponse`

NewJsonRpcErrorResponseWithDefaults instantiates a new JsonRpcErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonrpc

`func (o *JsonRpcErrorResponse) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *JsonRpcErrorResponse) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *JsonRpcErrorResponse) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.


### GetId

`func (o *JsonRpcErrorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonRpcErrorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonRpcErrorResponse) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *JsonRpcErrorResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JsonRpcErrorResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetError

`func (o *JsonRpcErrorResponse) GetError() JsonRpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *JsonRpcErrorResponse) GetErrorOk() (*JsonRpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *JsonRpcErrorResponse) SetError(v JsonRpcError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


