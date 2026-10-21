# JsonRpcError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | Код ошибки: &#x60;-32700&#x60; — ошибка разбора, &#x60;-32600&#x60; — некорректный запрос, &#x60;-32601&#x60; — метод не найден, &#x60;-32602&#x60; — некорректные параметры, &#x60;-32603&#x60; — внутренняя ошибка, &#x60;-32000&#x60; — ошибка транспорта, &#x60;-32001&#x60; — запрос не удалось связать с аккаунтом | 
**Message** | **string** | Текст ошибки | 
**Data** | Pointer to **map[string]interface{}** | Дополнительные сведения об ошибке | [optional] 

## Methods

### NewJsonRpcError

`func NewJsonRpcError(code int32, message string, ) *JsonRpcError`

NewJsonRpcError instantiates a new JsonRpcError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonRpcErrorWithDefaults

`func NewJsonRpcErrorWithDefaults() *JsonRpcError`

NewJsonRpcErrorWithDefaults instantiates a new JsonRpcError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *JsonRpcError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *JsonRpcError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *JsonRpcError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *JsonRpcError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JsonRpcError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JsonRpcError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *JsonRpcError) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JsonRpcError) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JsonRpcError) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *JsonRpcError) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


