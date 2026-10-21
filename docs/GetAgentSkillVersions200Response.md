# GetAgentSkillVersions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillVersions** | [**[]SkillVersion**](SkillVersion.md) |  | 
**Meta** | [**GetAgentSkillVersions200ResponseMeta**](GetAgentSkillVersions200ResponseMeta.md) |  | 

## Methods

### NewGetAgentSkillVersions200Response

`func NewGetAgentSkillVersions200Response(skillVersions []SkillVersion, meta GetAgentSkillVersions200ResponseMeta, ) *GetAgentSkillVersions200Response`

NewGetAgentSkillVersions200Response instantiates a new GetAgentSkillVersions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentSkillVersions200ResponseWithDefaults

`func NewGetAgentSkillVersions200ResponseWithDefaults() *GetAgentSkillVersions200Response`

NewGetAgentSkillVersions200ResponseWithDefaults instantiates a new GetAgentSkillVersions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillVersions

`func (o *GetAgentSkillVersions200Response) GetSkillVersions() []SkillVersion`

GetSkillVersions returns the SkillVersions field if non-nil, zero value otherwise.

### GetSkillVersionsOk

`func (o *GetAgentSkillVersions200Response) GetSkillVersionsOk() (*[]SkillVersion, bool)`

GetSkillVersionsOk returns a tuple with the SkillVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillVersions

`func (o *GetAgentSkillVersions200Response) SetSkillVersions(v []SkillVersion)`

SetSkillVersions sets SkillVersions field to given value.


### GetMeta

`func (o *GetAgentSkillVersions200Response) GetMeta() GetAgentSkillVersions200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentSkillVersions200Response) GetMetaOk() (*GetAgentSkillVersions200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentSkillVersions200Response) SetMeta(v GetAgentSkillVersions200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


