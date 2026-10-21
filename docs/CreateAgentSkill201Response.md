# CreateAgentSkill201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skill** | [**Skill**](Skill.md) |  | 
**Meta** | [**SkillSaveMeta**](SkillSaveMeta.md) |  | 

## Methods

### NewCreateAgentSkill201Response

`func NewCreateAgentSkill201Response(skill Skill, meta SkillSaveMeta, ) *CreateAgentSkill201Response`

NewCreateAgentSkill201Response instantiates a new CreateAgentSkill201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAgentSkill201ResponseWithDefaults

`func NewCreateAgentSkill201ResponseWithDefaults() *CreateAgentSkill201Response`

NewCreateAgentSkill201ResponseWithDefaults instantiates a new CreateAgentSkill201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkill

`func (o *CreateAgentSkill201Response) GetSkill() Skill`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *CreateAgentSkill201Response) GetSkillOk() (*Skill, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *CreateAgentSkill201Response) SetSkill(v Skill)`

SetSkill sets Skill field to given value.


### GetMeta

`func (o *CreateAgentSkill201Response) GetMeta() SkillSaveMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateAgentSkill201Response) GetMetaOk() (*SkillSaveMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateAgentSkill201Response) SetMeta(v SkillSaveMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


