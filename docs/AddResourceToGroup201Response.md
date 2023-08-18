# AddResourceToGroup201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Resource** | [**FirewallGroupResourceOutAPI**](FirewallGroupResourceOutAPI.md) |  | 

## Methods

### NewAddResourceToGroup201Response

`func NewAddResourceToGroup201Response(responseId string, resource FirewallGroupResourceOutAPI, ) *AddResourceToGroup201Response`

NewAddResourceToGroup201Response instantiates a new AddResourceToGroup201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddResourceToGroup201ResponseWithDefaults

`func NewAddResourceToGroup201ResponseWithDefaults() *AddResourceToGroup201Response`

NewAddResourceToGroup201ResponseWithDefaults instantiates a new AddResourceToGroup201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *AddResourceToGroup201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *AddResourceToGroup201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *AddResourceToGroup201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetResource

`func (o *AddResourceToGroup201Response) GetResource() FirewallGroupResourceOutAPI`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AddResourceToGroup201Response) GetResourceOk() (*FirewallGroupResourceOutAPI, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AddResourceToGroup201Response) SetResource(v FirewallGroupResourceOutAPI)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


