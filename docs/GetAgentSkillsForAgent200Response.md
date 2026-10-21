# GetAgentSkillsForAgent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentSkills** | [**[]AgentSkill**](AgentSkill.md) |  | 
**Meta** | [**GetAgentSkillsForAgent200ResponseMeta**](GetAgentSkillsForAgent200ResponseMeta.md) |  | 

## Methods

### NewGetAgentSkillsForAgent200Response

`func NewGetAgentSkillsForAgent200Response(agentSkills []AgentSkill, meta GetAgentSkillsForAgent200ResponseMeta, ) *GetAgentSkillsForAgent200Response`

NewGetAgentSkillsForAgent200Response instantiates a new GetAgentSkillsForAgent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentSkillsForAgent200ResponseWithDefaults

`func NewGetAgentSkillsForAgent200ResponseWithDefaults() *GetAgentSkillsForAgent200Response`

NewGetAgentSkillsForAgent200ResponseWithDefaults instantiates a new GetAgentSkillsForAgent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentSkills

`func (o *GetAgentSkillsForAgent200Response) GetAgentSkills() []AgentSkill`

GetAgentSkills returns the AgentSkills field if non-nil, zero value otherwise.

### GetAgentSkillsOk

`func (o *GetAgentSkillsForAgent200Response) GetAgentSkillsOk() (*[]AgentSkill, bool)`

GetAgentSkillsOk returns a tuple with the AgentSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSkills

`func (o *GetAgentSkillsForAgent200Response) SetAgentSkills(v []AgentSkill)`

SetAgentSkills sets AgentSkills field to given value.


### GetMeta

`func (o *GetAgentSkillsForAgent200Response) GetMeta() GetAgentSkillsForAgent200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentSkillsForAgent200Response) GetMetaOk() (*GetAgentSkillsForAgent200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentSkillsForAgent200Response) SetMeta(v GetAgentSkillsForAgent200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


