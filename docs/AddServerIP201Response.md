# AddServerIP201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerIp** | [**ServerIp**](ServerIp.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewAddServerIP201Response

`func NewAddServerIP201Response(serverIp ServerIp, responseId string, ) *AddServerIP201Response`

NewAddServerIP201Response instantiates a new AddServerIP201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServerIP201ResponseWithDefaults

`func NewAddServerIP201ResponseWithDefaults() *AddServerIP201Response`

NewAddServerIP201ResponseWithDefaults instantiates a new AddServerIP201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerIp

`func (o *AddServerIP201Response) GetServerIp() ServerIp`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *AddServerIP201Response) GetServerIpOk() (*ServerIp, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *AddServerIP201Response) SetServerIp(v ServerIp)`

SetServerIp sets ServerIp field to given value.


### GetResponseId

`func (o *AddServerIP201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *AddServerIP201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *AddServerIP201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


