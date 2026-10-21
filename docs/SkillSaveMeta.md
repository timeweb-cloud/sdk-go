# SkillSaveMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionCreated** | **bool** | Создана ли новая версия. &#x60;false&#x60; — изменился только заголовок или изменений не было | 
**AffectedAgentsCount** | **int32** | У скольких агентов навык включён. Их поведение не меняется, пока их не переведут на новую версию | 
**Warnings** | **[]string** | Предупреждения, которые не мешают сохранению. &#x60;when_to_use_too_short&#x60; — текст «когда применять» короче 40 символов: модель не сможет отличить подходящий запрос от постороннего | 
**PrunedVersions** | **[]int32** | Номера версий, удалённых из истории из-за лимита в 50 версий | 

## Methods

### NewSkillSaveMeta

`func NewSkillSaveMeta(versionCreated bool, affectedAgentsCount int32, warnings []string, prunedVersions []int32, ) *SkillSaveMeta`

NewSkillSaveMeta instantiates a new SkillSaveMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillSaveMetaWithDefaults

`func NewSkillSaveMetaWithDefaults() *SkillSaveMeta`

NewSkillSaveMetaWithDefaults instantiates a new SkillSaveMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionCreated

`func (o *SkillSaveMeta) GetVersionCreated() bool`

GetVersionCreated returns the VersionCreated field if non-nil, zero value otherwise.

### GetVersionCreatedOk

`func (o *SkillSaveMeta) GetVersionCreatedOk() (*bool, bool)`

GetVersionCreatedOk returns a tuple with the VersionCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCreated

`func (o *SkillSaveMeta) SetVersionCreated(v bool)`

SetVersionCreated sets VersionCreated field to given value.


### GetAffectedAgentsCount

`func (o *SkillSaveMeta) GetAffectedAgentsCount() int32`

GetAffectedAgentsCount returns the AffectedAgentsCount field if non-nil, zero value otherwise.

### GetAffectedAgentsCountOk

`func (o *SkillSaveMeta) GetAffectedAgentsCountOk() (*int32, bool)`

GetAffectedAgentsCountOk returns a tuple with the AffectedAgentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedAgentsCount

`func (o *SkillSaveMeta) SetAffectedAgentsCount(v int32)`

SetAffectedAgentsCount sets AffectedAgentsCount field to given value.


### GetWarnings

`func (o *SkillSaveMeta) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SkillSaveMeta) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SkillSaveMeta) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.


### GetPrunedVersions

`func (o *SkillSaveMeta) GetPrunedVersions() []int32`

GetPrunedVersions returns the PrunedVersions field if non-nil, zero value otherwise.

### GetPrunedVersionsOk

`func (o *SkillSaveMeta) GetPrunedVersionsOk() (*[]int32, bool)`

GetPrunedVersionsOk returns a tuple with the PrunedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrunedVersions

`func (o *SkillSaveMeta) SetPrunedVersions(v []int32)`

SetPrunedVersions sets PrunedVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


