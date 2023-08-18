# CreateServerDisk201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerDisk** | [**ServerDisk**](ServerDisk.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCreateServerDisk201Response

`func NewCreateServerDisk201Response(serverDisk ServerDisk, responseId string, ) *CreateServerDisk201Response`

NewCreateServerDisk201Response instantiates a new CreateServerDisk201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerDisk201ResponseWithDefaults

`func NewCreateServerDisk201ResponseWithDefaults() *CreateServerDisk201Response`

NewCreateServerDisk201ResponseWithDefaults instantiates a new CreateServerDisk201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerDisk

`func (o *CreateServerDisk201Response) GetServerDisk() ServerDisk`

GetServerDisk returns the ServerDisk field if non-nil, zero value otherwise.

### GetServerDiskOk

`func (o *CreateServerDisk201Response) GetServerDiskOk() (*ServerDisk, bool)`

GetServerDiskOk returns a tuple with the ServerDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDisk

`func (o *CreateServerDisk201Response) SetServerDisk(v ServerDisk)`

SetServerDisk sets ServerDisk field to given value.


### GetResponseId

`func (o *CreateServerDisk201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateServerDisk201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateServerDisk201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


