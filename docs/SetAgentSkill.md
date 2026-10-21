# SetAgentSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включить (&#x60;true&#x60;) или выключить (&#x60;false&#x60;) навык у агента | 
**Version** | Pointer to **int32** | Закрепить агента на этой версии из истории навыка. Если не передано, при первом включении закрепляется текущая версия, а у уже подключённого навыка закреплённая версия не меняется | [optional] 

## Methods

### NewSetAgentSkill

`func NewSetAgentSkill(isEnabled bool, ) *SetAgentSkill`

NewSetAgentSkill instantiates a new SetAgentSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAgentSkillWithDefaults

`func NewSetAgentSkillWithDefaults() *SetAgentSkill`

NewSetAgentSkillWithDefaults instantiates a new SetAgentSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *SetAgentSkill) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SetAgentSkill) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SetAgentSkill) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetVersion

`func (o *SetAgentSkill) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SetAgentSkill) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SetAgentSkill) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SetAgentSkill) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


