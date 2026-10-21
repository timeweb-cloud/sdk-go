# GetAgentsV3200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | [**[]AgentV3**](AgentV3.md) |  | 
**Meta** | [**GetAgents200ResponseMeta**](GetAgents200ResponseMeta.md) |  | 

## Methods

### NewGetAgentsV3200Response

`func NewGetAgentsV3200Response(agents []AgentV3, meta GetAgents200ResponseMeta, ) *GetAgentsV3200Response`

NewGetAgentsV3200Response instantiates a new GetAgentsV3200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentsV3200ResponseWithDefaults

`func NewGetAgentsV3200ResponseWithDefaults() *GetAgentsV3200Response`

NewGetAgentsV3200ResponseWithDefaults instantiates a new GetAgentsV3200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *GetAgentsV3200Response) GetAgents() []AgentV3`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *GetAgentsV3200Response) GetAgentsOk() (*[]AgentV3, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *GetAgentsV3200Response) SetAgents(v []AgentV3)`

SetAgents sets Agents field to given value.


### GetMeta

`func (o *GetAgentsV3200Response) GetMeta() GetAgents200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentsV3200Response) GetMetaOk() (*GetAgents200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentsV3200Response) SetMeta(v GetAgents200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


