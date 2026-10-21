# SyncAgentSkillVersionsForAgent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentSkills** | [**[]AgentSkill**](AgentSkill.md) |  | 
**Meta** | [**GetAgentSkillsForAgent200ResponseMeta**](GetAgentSkillsForAgent200ResponseMeta.md) |  | 
**UpdatedCount** | **int32** | Сколько навыков переведено на текущую версию | 

## Methods

### NewSyncAgentSkillVersionsForAgent200Response

`func NewSyncAgentSkillVersionsForAgent200Response(agentSkills []AgentSkill, meta GetAgentSkillsForAgent200ResponseMeta, updatedCount int32, ) *SyncAgentSkillVersionsForAgent200Response`

NewSyncAgentSkillVersionsForAgent200Response instantiates a new SyncAgentSkillVersionsForAgent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncAgentSkillVersionsForAgent200ResponseWithDefaults

`func NewSyncAgentSkillVersionsForAgent200ResponseWithDefaults() *SyncAgentSkillVersionsForAgent200Response`

NewSyncAgentSkillVersionsForAgent200ResponseWithDefaults instantiates a new SyncAgentSkillVersionsForAgent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentSkills

`func (o *SyncAgentSkillVersionsForAgent200Response) GetAgentSkills() []AgentSkill`

GetAgentSkills returns the AgentSkills field if non-nil, zero value otherwise.

### GetAgentSkillsOk

`func (o *SyncAgentSkillVersionsForAgent200Response) GetAgentSkillsOk() (*[]AgentSkill, bool)`

GetAgentSkillsOk returns a tuple with the AgentSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSkills

`func (o *SyncAgentSkillVersionsForAgent200Response) SetAgentSkills(v []AgentSkill)`

SetAgentSkills sets AgentSkills field to given value.


### GetMeta

`func (o *SyncAgentSkillVersionsForAgent200Response) GetMeta() GetAgentSkillsForAgent200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SyncAgentSkillVersionsForAgent200Response) GetMetaOk() (*GetAgentSkillsForAgent200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SyncAgentSkillVersionsForAgent200Response) SetMeta(v GetAgentSkillsForAgent200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetUpdatedCount

`func (o *SyncAgentSkillVersionsForAgent200Response) GetUpdatedCount() int32`

GetUpdatedCount returns the UpdatedCount field if non-nil, zero value otherwise.

### GetUpdatedCountOk

`func (o *SyncAgentSkillVersionsForAgent200Response) GetUpdatedCountOk() (*int32, bool)`

GetUpdatedCountOk returns a tuple with the UpdatedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedCount

`func (o *SyncAgentSkillVersionsForAgent200Response) SetUpdatedCount(v int32)`

SetUpdatedCount sets UpdatedCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


