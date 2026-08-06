# UpdateDatabaseInstance409Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **float32** | Короткий идентификатор, соответствующий возвращаемому коду состояния HTTP. | 
**Message** | Pointer to **string** | Сообщение, предоставляющее дополнительную информацию об ошибке, в том числе сведения, помогающие устранить ее, когда это возможно. | [optional] 
**ErrorCode** | **string** | Краткое описание ошибки HTTP на основе статуса. | 
**ResponseId** | **string** | ID запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewUpdateDatabaseInstance409Response

`func NewUpdateDatabaseInstance409Response(statusCode float32, errorCode string, responseId string, ) *UpdateDatabaseInstance409Response`

NewUpdateDatabaseInstance409Response instantiates a new UpdateDatabaseInstance409Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatabaseInstance409ResponseWithDefaults

`func NewUpdateDatabaseInstance409ResponseWithDefaults() *UpdateDatabaseInstance409Response`

NewUpdateDatabaseInstance409ResponseWithDefaults instantiates a new UpdateDatabaseInstance409Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *UpdateDatabaseInstance409Response) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *UpdateDatabaseInstance409Response) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *UpdateDatabaseInstance409Response) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetMessage

`func (o *UpdateDatabaseInstance409Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateDatabaseInstance409Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateDatabaseInstance409Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateDatabaseInstance409Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *UpdateDatabaseInstance409Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *UpdateDatabaseInstance409Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *UpdateDatabaseInstance409Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetResponseId

`func (o *UpdateDatabaseInstance409Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *UpdateDatabaseInstance409Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *UpdateDatabaseInstance409Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


