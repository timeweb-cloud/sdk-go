# SkillAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **int32** | ID агента | 
**AgentName** | **string** | Название агента | 
**IsEnabled** | **bool** | Навык включён у агента | 
**PinnedVersion** | **int32** | Версия навыка, на которой закреплён агент | 
**HasNewerVersion** | **bool** | Текущая версия навыка отличается от закреплённой | 

## Methods

### NewSkillAgent

`func NewSkillAgent(agentId int32, agentName string, isEnabled bool, pinnedVersion int32, hasNewerVersion bool, ) *SkillAgent`

NewSkillAgent instantiates a new SkillAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillAgentWithDefaults

`func NewSkillAgentWithDefaults() *SkillAgent`

NewSkillAgentWithDefaults instantiates a new SkillAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *SkillAgent) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *SkillAgent) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *SkillAgent) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.


### GetAgentName

`func (o *SkillAgent) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *SkillAgent) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *SkillAgent) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.


### GetIsEnabled

`func (o *SkillAgent) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SkillAgent) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SkillAgent) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetPinnedVersion

`func (o *SkillAgent) GetPinnedVersion() int32`

GetPinnedVersion returns the PinnedVersion field if non-nil, zero value otherwise.

### GetPinnedVersionOk

`func (o *SkillAgent) GetPinnedVersionOk() (*int32, bool)`

GetPinnedVersionOk returns a tuple with the PinnedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedVersion

`func (o *SkillAgent) SetPinnedVersion(v int32)`

SetPinnedVersion sets PinnedVersion field to given value.


### GetHasNewerVersion

`func (o *SkillAgent) GetHasNewerVersion() bool`

GetHasNewerVersion returns the HasNewerVersion field if non-nil, zero value otherwise.

### GetHasNewerVersionOk

`func (o *SkillAgent) GetHasNewerVersionOk() (*bool, bool)`

GetHasNewerVersionOk returns a tuple with the HasNewerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNewerVersion

`func (o *SkillAgent) SetHasNewerVersion(v bool)`

SetHasNewerVersion sets HasNewerVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


