# ComponentsSchemasBaseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** |  | 
**ErrorCode** | **string** |  | 
**Message** | **string** |  | 
**ResponseId** | Pointer to **string** |  | [optional] 

## Methods

### NewComponentsSchemasBaseError

`func NewComponentsSchemasBaseError(statusCode int32, errorCode string, message string, ) *ComponentsSchemasBaseError`

NewComponentsSchemasBaseError instantiates a new ComponentsSchemasBaseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentsSchemasBaseErrorWithDefaults

`func NewComponentsSchemasBaseErrorWithDefaults() *ComponentsSchemasBaseError`

NewComponentsSchemasBaseErrorWithDefaults instantiates a new ComponentsSchemasBaseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *ComponentsSchemasBaseError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ComponentsSchemasBaseError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ComponentsSchemasBaseError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetErrorCode

`func (o *ComponentsSchemasBaseError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ComponentsSchemasBaseError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ComponentsSchemasBaseError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetMessage

`func (o *ComponentsSchemasBaseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ComponentsSchemasBaseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ComponentsSchemasBaseError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResponseId

`func (o *ComponentsSchemasBaseError) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *ComponentsSchemasBaseError) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *ComponentsSchemasBaseError) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *ComponentsSchemasBaseError) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


