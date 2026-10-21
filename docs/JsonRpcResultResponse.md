# JsonRpcResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jsonrpc** | **string** | Версия протокола JSON-RPC, всегда &#x60;2.0&#x60; | 
**Id** | **string** | ID запроса, выбранный клиентом; ответ приходит с тем же ID | 
**Result** | [**InitializeResult**](InitializeResult.md) |  | 

## Methods

### NewJsonRpcResultResponse

`func NewJsonRpcResultResponse(jsonrpc string, id string, result InitializeResult, ) *JsonRpcResultResponse`

NewJsonRpcResultResponse instantiates a new JsonRpcResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonRpcResultResponseWithDefaults

`func NewJsonRpcResultResponseWithDefaults() *JsonRpcResultResponse`

NewJsonRpcResultResponseWithDefaults instantiates a new JsonRpcResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonrpc

`func (o *JsonRpcResultResponse) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *JsonRpcResultResponse) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *JsonRpcResultResponse) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.


### GetId

`func (o *JsonRpcResultResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonRpcResultResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonRpcResultResponse) SetId(v string)`

SetId sets Id field to given value.


### GetResult

`func (o *JsonRpcResultResponse) GetResult() InitializeResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *JsonRpcResultResponse) GetResultOk() (*InitializeResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *JsonRpcResultResponse) SetResult(v InitializeResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


