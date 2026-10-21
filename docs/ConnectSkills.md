# ConnectSkills

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillIds** | Pointer to **[]int32** | ID навыков. Агент закрепляется на их текущих версиях. Обязателен, если не передан &#x60;skills&#x60; | [optional] 
**Skills** | Pointer to [**[]ConnectSkillItem**](ConnectSkillItem.md) | Навыки с версиями. Если передан, &#x60;skill_ids&#x60; не используется | [optional] 

## Methods

### NewConnectSkills

`func NewConnectSkills() *ConnectSkills`

NewConnectSkills instantiates a new ConnectSkills object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectSkillsWithDefaults

`func NewConnectSkillsWithDefaults() *ConnectSkills`

NewConnectSkillsWithDefaults instantiates a new ConnectSkills object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillIds

`func (o *ConnectSkills) GetSkillIds() []int32`

GetSkillIds returns the SkillIds field if non-nil, zero value otherwise.

### GetSkillIdsOk

`func (o *ConnectSkills) GetSkillIdsOk() (*[]int32, bool)`

GetSkillIdsOk returns a tuple with the SkillIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillIds

`func (o *ConnectSkills) SetSkillIds(v []int32)`

SetSkillIds sets SkillIds field to given value.

### HasSkillIds

`func (o *ConnectSkills) HasSkillIds() bool`

HasSkillIds returns a boolean if a field has been set.

### GetSkills

`func (o *ConnectSkills) GetSkills() []ConnectSkillItem`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ConnectSkills) GetSkillsOk() (*[]ConnectSkillItem, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ConnectSkills) SetSkills(v []ConnectSkillItem)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ConnectSkills) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


