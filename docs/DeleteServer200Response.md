# DeleteServer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerDelete** | [**DeleteServiceResponse**](DeleteServiceResponse.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewDeleteServer200Response

`func NewDeleteServer200Response(serverDelete DeleteServiceResponse, responseId string, ) *DeleteServer200Response`

NewDeleteServer200Response instantiates a new DeleteServer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteServer200ResponseWithDefaults

`func NewDeleteServer200ResponseWithDefaults() *DeleteServer200Response`

NewDeleteServer200ResponseWithDefaults instantiates a new DeleteServer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerDelete

`func (o *DeleteServer200Response) GetServerDelete() DeleteServiceResponse`

GetServerDelete returns the ServerDelete field if non-nil, zero value otherwise.

### GetServerDeleteOk

`func (o *DeleteServer200Response) GetServerDeleteOk() (*DeleteServiceResponse, bool)`

GetServerDeleteOk returns a tuple with the ServerDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDelete

`func (o *DeleteServer200Response) SetServerDelete(v DeleteServiceResponse)`

SetServerDelete sets ServerDelete field to given value.


### GetResponseId

`func (o *DeleteServer200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *DeleteServer200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *DeleteServer200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


