# AgentSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillId** | **int32** | ID навыка | 
**Name** | **string** | Название навыка | 
**WhenToUse** | **string** | Когда применять навык | 
**CurrentVersion** | **int32** | Номер текущей версии навыка | 
**IsConnected** | **bool** | Навык подключён к агенту | 
**IsEnabled** | **bool** | Навык включён у агента | 
**PinnedVersion** | **NullableInt32** | Версия, на которой закреплён агент. &#x60;null&#x60; — навык не подключён | 
**HasNewerVersion** | **bool** | Текущая версия навыка отличается от закреплённой | 
**IsFromCatalog** | **bool** | Навык установлен из каталога | 
**IsModified** | **bool** | Установленный из каталога навык изменялся владельцем | 

## Methods

### NewAgentSkill

`func NewAgentSkill(skillId int32, name string, whenToUse string, currentVersion int32, isConnected bool, isEnabled bool, pinnedVersion NullableInt32, hasNewerVersion bool, isFromCatalog bool, isModified bool, ) *AgentSkill`

NewAgentSkill instantiates a new AgentSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSkillWithDefaults

`func NewAgentSkillWithDefaults() *AgentSkill`

NewAgentSkillWithDefaults instantiates a new AgentSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillId

`func (o *AgentSkill) GetSkillId() int32`

GetSkillId returns the SkillId field if non-nil, zero value otherwise.

### GetSkillIdOk

`func (o *AgentSkill) GetSkillIdOk() (*int32, bool)`

GetSkillIdOk returns a tuple with the SkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillId

`func (o *AgentSkill) SetSkillId(v int32)`

SetSkillId sets SkillId field to given value.


### GetName

`func (o *AgentSkill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentSkill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentSkill) SetName(v string)`

SetName sets Name field to given value.


### GetWhenToUse

`func (o *AgentSkill) GetWhenToUse() string`

GetWhenToUse returns the WhenToUse field if non-nil, zero value otherwise.

### GetWhenToUseOk

`func (o *AgentSkill) GetWhenToUseOk() (*string, bool)`

GetWhenToUseOk returns a tuple with the WhenToUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenToUse

`func (o *AgentSkill) SetWhenToUse(v string)`

SetWhenToUse sets WhenToUse field to given value.


### GetCurrentVersion

`func (o *AgentSkill) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *AgentSkill) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *AgentSkill) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.


### GetIsConnected

`func (o *AgentSkill) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *AgentSkill) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *AgentSkill) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.


### GetIsEnabled

`func (o *AgentSkill) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AgentSkill) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AgentSkill) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetPinnedVersion

`func (o *AgentSkill) GetPinnedVersion() int32`

GetPinnedVersion returns the PinnedVersion field if non-nil, zero value otherwise.

### GetPinnedVersionOk

`func (o *AgentSkill) GetPinnedVersionOk() (*int32, bool)`

GetPinnedVersionOk returns a tuple with the PinnedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedVersion

`func (o *AgentSkill) SetPinnedVersion(v int32)`

SetPinnedVersion sets PinnedVersion field to given value.


### SetPinnedVersionNil

`func (o *AgentSkill) SetPinnedVersionNil(b bool)`

 SetPinnedVersionNil sets the value for PinnedVersion to be an explicit nil

### UnsetPinnedVersion
`func (o *AgentSkill) UnsetPinnedVersion()`

UnsetPinnedVersion ensures that no value is present for PinnedVersion, not even an explicit nil
### GetHasNewerVersion

`func (o *AgentSkill) GetHasNewerVersion() bool`

GetHasNewerVersion returns the HasNewerVersion field if non-nil, zero value otherwise.

### GetHasNewerVersionOk

`func (o *AgentSkill) GetHasNewerVersionOk() (*bool, bool)`

GetHasNewerVersionOk returns a tuple with the HasNewerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNewerVersion

`func (o *AgentSkill) SetHasNewerVersion(v bool)`

SetHasNewerVersion sets HasNewerVersion field to given value.


### GetIsFromCatalog

`func (o *AgentSkill) GetIsFromCatalog() bool`

GetIsFromCatalog returns the IsFromCatalog field if non-nil, zero value otherwise.

### GetIsFromCatalogOk

`func (o *AgentSkill) GetIsFromCatalogOk() (*bool, bool)`

GetIsFromCatalogOk returns a tuple with the IsFromCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromCatalog

`func (o *AgentSkill) SetIsFromCatalog(v bool)`

SetIsFromCatalog sets IsFromCatalog field to given value.


### GetIsModified

`func (o *AgentSkill) GetIsModified() bool`

GetIsModified returns the IsModified field if non-nil, zero value otherwise.

### GetIsModifiedOk

`func (o *AgentSkill) GetIsModifiedOk() (*bool, bool)`

GetIsModifiedOk returns a tuple with the IsModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModified

`func (o *AgentSkill) SetIsModified(v bool)`

SetIsModified sets IsModified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


