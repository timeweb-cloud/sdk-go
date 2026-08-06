# AddCdnCertificate422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **float32** | Короткий идентификатор, соответствующий возвращаемому коду состояния HTTP. | 
**Message** | Pointer to **string** | Сообщение, предоставляющее дополнительную информацию об ошибке, в том числе сведения, помогающие устранить ее, когда это возможно. | [optional] 
**ErrorCode** | **string** | Краткое описание ошибки HTTP на основе статуса. | 
**ResponseId** | **string** | ID запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewAddCdnCertificate422Response

`func NewAddCdnCertificate422Response(statusCode float32, errorCode string, responseId string, ) *AddCdnCertificate422Response`

NewAddCdnCertificate422Response instantiates a new AddCdnCertificate422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCdnCertificate422ResponseWithDefaults

`func NewAddCdnCertificate422ResponseWithDefaults() *AddCdnCertificate422Response`

NewAddCdnCertificate422ResponseWithDefaults instantiates a new AddCdnCertificate422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *AddCdnCertificate422Response) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *AddCdnCertificate422Response) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *AddCdnCertificate422Response) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetMessage

`func (o *AddCdnCertificate422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AddCdnCertificate422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AddCdnCertificate422Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AddCdnCertificate422Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *AddCdnCertificate422Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AddCdnCertificate422Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AddCdnCertificate422Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetResponseId

`func (o *AddCdnCertificate422Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *AddCdnCertificate422Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *AddCdnCertificate422Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


