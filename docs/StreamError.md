# StreamError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | Код ошибки | [optional] 
**Message** | Pointer to **string** | Текст ошибки | [optional] 
**StatusCode** | Pointer to **float32** | Код HTTP-статуса (у ошибок валидации и внутренних ошибок) | [optional] 
**Error** | Pointer to **string** | Название HTTP-статуса (только у ошибок валидации) | [optional] 

## Methods

### NewStreamError

`func NewStreamError() *StreamError`

NewStreamError instantiates a new StreamError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamErrorWithDefaults

`func NewStreamErrorWithDefaults() *StreamError`

NewStreamErrorWithDefaults instantiates a new StreamError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *StreamError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *StreamError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *StreamError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *StreamError) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessage

`func (o *StreamError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StreamError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StreamError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StreamError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *StreamError) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *StreamError) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *StreamError) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *StreamError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *StreamError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StreamError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


