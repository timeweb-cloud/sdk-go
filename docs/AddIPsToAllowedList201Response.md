# AddIPsToAllowedList201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | [**AddIps**](AddIps.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewAddIPsToAllowedList201Response

`func NewAddIPsToAllowedList201Response(ips AddIps, responseId string, ) *AddIPsToAllowedList201Response`

NewAddIPsToAllowedList201Response instantiates a new AddIPsToAllowedList201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIPsToAllowedList201ResponseWithDefaults

`func NewAddIPsToAllowedList201ResponseWithDefaults() *AddIPsToAllowedList201Response`

NewAddIPsToAllowedList201ResponseWithDefaults instantiates a new AddIPsToAllowedList201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *AddIPsToAllowedList201Response) GetIps() AddIps`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *AddIPsToAllowedList201Response) GetIpsOk() (*AddIps, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *AddIPsToAllowedList201Response) SetIps(v AddIps)`

SetIps sets Ips field to given value.


### GetResponseId

`func (o *AddIPsToAllowedList201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *AddIPsToAllowedList201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *AddIPsToAllowedList201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


