# GetAgentToolsByAgent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentTools** | [**[]Tool**](Tool.md) |  | 
**Meta** | [**GetAgentToolsByAgent200ResponseMeta**](GetAgentToolsByAgent200ResponseMeta.md) |  | 

## Methods

### NewGetAgentToolsByAgent200Response

`func NewGetAgentToolsByAgent200Response(agentTools []Tool, meta GetAgentToolsByAgent200ResponseMeta, ) *GetAgentToolsByAgent200Response`

NewGetAgentToolsByAgent200Response instantiates a new GetAgentToolsByAgent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentToolsByAgent200ResponseWithDefaults

`func NewGetAgentToolsByAgent200ResponseWithDefaults() *GetAgentToolsByAgent200Response`

NewGetAgentToolsByAgent200ResponseWithDefaults instantiates a new GetAgentToolsByAgent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentTools

`func (o *GetAgentToolsByAgent200Response) GetAgentTools() []Tool`

GetAgentTools returns the AgentTools field if non-nil, zero value otherwise.

### GetAgentToolsOk

`func (o *GetAgentToolsByAgent200Response) GetAgentToolsOk() (*[]Tool, bool)`

GetAgentToolsOk returns a tuple with the AgentTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTools

`func (o *GetAgentToolsByAgent200Response) SetAgentTools(v []Tool)`

SetAgentTools sets AgentTools field to given value.


### GetMeta

`func (o *GetAgentToolsByAgent200Response) GetMeta() GetAgentToolsByAgent200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentToolsByAgent200Response) GetMetaOk() (*GetAgentToolsByAgent200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentToolsByAgent200Response) SetMeta(v GetAgentToolsByAgent200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


