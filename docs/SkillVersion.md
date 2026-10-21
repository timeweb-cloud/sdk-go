# SkillVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID версии | 
**Version** | **int32** | Номер версии | 
**Name** | **string** | Название навыка на момент создания версии | 
**WhenToUse** | **string** | Текст «когда применять» в этой версии | 
**Instruction** | **string** | Инструкция в этой версии | 
**RolledBackFromVersion** | **NullableInt32** | Номер версии, с которой эта версия создана откатом | 
**IsCurrent** | **bool** | Это текущая версия навыка | 
**CreatedAt** | **time.Time** | Дата и время создания версии | 

## Methods

### NewSkillVersion

`func NewSkillVersion(id int32, version int32, name string, whenToUse string, instruction string, rolledBackFromVersion NullableInt32, isCurrent bool, createdAt time.Time, ) *SkillVersion`

NewSkillVersion instantiates a new SkillVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillVersionWithDefaults

`func NewSkillVersionWithDefaults() *SkillVersion`

NewSkillVersionWithDefaults instantiates a new SkillVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SkillVersion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkillVersion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkillVersion) SetId(v int32)`

SetId sets Id field to given value.


### GetVersion

`func (o *SkillVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SkillVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SkillVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetName

`func (o *SkillVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkillVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkillVersion) SetName(v string)`

SetName sets Name field to given value.


### GetWhenToUse

`func (o *SkillVersion) GetWhenToUse() string`

GetWhenToUse returns the WhenToUse field if non-nil, zero value otherwise.

### GetWhenToUseOk

`func (o *SkillVersion) GetWhenToUseOk() (*string, bool)`

GetWhenToUseOk returns a tuple with the WhenToUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenToUse

`func (o *SkillVersion) SetWhenToUse(v string)`

SetWhenToUse sets WhenToUse field to given value.


### GetInstruction

`func (o *SkillVersion) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *SkillVersion) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *SkillVersion) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.


### GetRolledBackFromVersion

`func (o *SkillVersion) GetRolledBackFromVersion() int32`

GetRolledBackFromVersion returns the RolledBackFromVersion field if non-nil, zero value otherwise.

### GetRolledBackFromVersionOk

`func (o *SkillVersion) GetRolledBackFromVersionOk() (*int32, bool)`

GetRolledBackFromVersionOk returns a tuple with the RolledBackFromVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolledBackFromVersion

`func (o *SkillVersion) SetRolledBackFromVersion(v int32)`

SetRolledBackFromVersion sets RolledBackFromVersion field to given value.


### SetRolledBackFromVersionNil

`func (o *SkillVersion) SetRolledBackFromVersionNil(b bool)`

 SetRolledBackFromVersionNil sets the value for RolledBackFromVersion to be an explicit nil

### UnsetRolledBackFromVersion
`func (o *SkillVersion) UnsetRolledBackFromVersion()`

UnsetRolledBackFromVersion ensures that no value is present for RolledBackFromVersion, not even an explicit nil
### GetIsCurrent

`func (o *SkillVersion) GetIsCurrent() bool`

GetIsCurrent returns the IsCurrent field if non-nil, zero value otherwise.

### GetIsCurrentOk

`func (o *SkillVersion) GetIsCurrentOk() (*bool, bool)`

GetIsCurrentOk returns a tuple with the IsCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrent

`func (o *SkillVersion) SetIsCurrent(v bool)`

SetIsCurrent sets IsCurrent field to given value.


### GetCreatedAt

`func (o *SkillVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SkillVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SkillVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


