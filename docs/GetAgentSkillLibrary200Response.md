# GetAgentSkillLibrary200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skills** | [**[]Skill**](Skill.md) |  | 
**Meta** | [**GetAgentSkillLibrary200ResponseMeta**](GetAgentSkillLibrary200ResponseMeta.md) |  | 

## Methods

### NewGetAgentSkillLibrary200Response

`func NewGetAgentSkillLibrary200Response(skills []Skill, meta GetAgentSkillLibrary200ResponseMeta, ) *GetAgentSkillLibrary200Response`

NewGetAgentSkillLibrary200Response instantiates a new GetAgentSkillLibrary200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentSkillLibrary200ResponseWithDefaults

`func NewGetAgentSkillLibrary200ResponseWithDefaults() *GetAgentSkillLibrary200Response`

NewGetAgentSkillLibrary200ResponseWithDefaults instantiates a new GetAgentSkillLibrary200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkills

`func (o *GetAgentSkillLibrary200Response) GetSkills() []Skill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *GetAgentSkillLibrary200Response) GetSkillsOk() (*[]Skill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *GetAgentSkillLibrary200Response) SetSkills(v []Skill)`

SetSkills sets Skills field to given value.


### GetMeta

`func (o *GetAgentSkillLibrary200Response) GetMeta() GetAgentSkillLibrary200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentSkillLibrary200Response) GetMetaOk() (*GetAgentSkillLibrary200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentSkillLibrary200Response) SetMeta(v GetAgentSkillLibrary200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


