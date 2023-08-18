# CreateServer201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | [**Vds**](Vds.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCreateServer201Response

`func NewCreateServer201Response(server Vds, responseId string, ) *CreateServer201Response`

NewCreateServer201Response instantiates a new CreateServer201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServer201ResponseWithDefaults

`func NewCreateServer201ResponseWithDefaults() *CreateServer201Response`

NewCreateServer201ResponseWithDefaults instantiates a new CreateServer201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *CreateServer201Response) GetServer() Vds`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CreateServer201Response) GetServerOk() (*Vds, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CreateServer201Response) SetServer(v Vds)`

SetServer sets Server field to given value.


### GetResponseId

`func (o *CreateServer201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateServer201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateServer201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


