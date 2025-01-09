# GetFinances401Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **float32** | Короткий идентификатор, соответствующий возвращаемому коду состояния HTTP. | 
**Message** | Pointer to **string** | Сообщение, предоставляющее дополнительную информацию об ошибке, в том числе сведения, помогающие устранить ее, когда это возможно. | [optional] 
**ErrorCode** | **string** | Краткое описание ошибки HTTP на основе статуса. | 
**ResponseId** | **string** | ID запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetFinances401Response

`func NewGetFinances401Response(statusCode float32, errorCode string, responseId string, ) *GetFinances401Response`

NewGetFinances401Response instantiates a new GetFinances401Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinances401ResponseWithDefaults

`func NewGetFinances401ResponseWithDefaults() *GetFinances401Response`

NewGetFinances401ResponseWithDefaults instantiates a new GetFinances401Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *GetFinances401Response) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetFinances401Response) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetFinances401Response) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetMessage

`func (o *GetFinances401Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetFinances401Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetFinances401Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetFinances401Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetFinances401Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetFinances401Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetFinances401Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetResponseId

`func (o *GetFinances401Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetFinances401Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetFinances401Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


