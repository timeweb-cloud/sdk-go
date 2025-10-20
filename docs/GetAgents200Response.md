# GetAgents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | [**[]Agent**](Agent.md) |  | 
**Meta** | [**GetAgents200ResponseMeta**](GetAgents200ResponseMeta.md) |  | 

## Methods

### NewGetAgents200Response

`func NewGetAgents200Response(agents []Agent, meta GetAgents200ResponseMeta, ) *GetAgents200Response`

NewGetAgents200Response instantiates a new GetAgents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgents200ResponseWithDefaults

`func NewGetAgents200ResponseWithDefaults() *GetAgents200Response`

NewGetAgents200ResponseWithDefaults instantiates a new GetAgents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *GetAgents200Response) GetAgents() []Agent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *GetAgents200Response) GetAgentsOk() (*[]Agent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *GetAgents200Response) SetAgents(v []Agent)`

SetAgents sets Agents field to given value.


### GetMeta

`func (o *GetAgents200Response) GetMeta() GetAgents200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgents200Response) GetMetaOk() (*GetAgents200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgents200Response) SetMeta(v GetAgents200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


