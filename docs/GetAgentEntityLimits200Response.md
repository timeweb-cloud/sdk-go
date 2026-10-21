# GetAgentEntityLimits200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentEntityLimits** | [**[]AgentEntityLimit**](AgentEntityLimit.md) |  | 
**Meta** | [**GetAgentEntityLimits200ResponseMeta**](GetAgentEntityLimits200ResponseMeta.md) |  | 

## Methods

### NewGetAgentEntityLimits200Response

`func NewGetAgentEntityLimits200Response(agentEntityLimits []AgentEntityLimit, meta GetAgentEntityLimits200ResponseMeta, ) *GetAgentEntityLimits200Response`

NewGetAgentEntityLimits200Response instantiates a new GetAgentEntityLimits200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentEntityLimits200ResponseWithDefaults

`func NewGetAgentEntityLimits200ResponseWithDefaults() *GetAgentEntityLimits200Response`

NewGetAgentEntityLimits200ResponseWithDefaults instantiates a new GetAgentEntityLimits200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentEntityLimits

`func (o *GetAgentEntityLimits200Response) GetAgentEntityLimits() []AgentEntityLimit`

GetAgentEntityLimits returns the AgentEntityLimits field if non-nil, zero value otherwise.

### GetAgentEntityLimitsOk

`func (o *GetAgentEntityLimits200Response) GetAgentEntityLimitsOk() (*[]AgentEntityLimit, bool)`

GetAgentEntityLimitsOk returns a tuple with the AgentEntityLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentEntityLimits

`func (o *GetAgentEntityLimits200Response) SetAgentEntityLimits(v []AgentEntityLimit)`

SetAgentEntityLimits sets AgentEntityLimits field to given value.


### GetMeta

`func (o *GetAgentEntityLimits200Response) GetMeta() GetAgentEntityLimits200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentEntityLimits200Response) GetMetaOk() (*GetAgentEntityLimits200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentEntityLimits200Response) SetMeta(v GetAgentEntityLimits200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


