# CreateSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название навыка, уникальное в аккаунте | 
**WhenToUse** | **string** | Когда применять навык. Описывайте ситуации и формулировки клиента. Если текст короче 40 символов, придёт предупреждение &#x60;when_to_use_too_short&#x60; | 
**Instruction** | **string** | Инструкция навыка | 
**Scripts** | Pointer to [**[]SkillScriptInput**](SkillScriptInput.md) | Скрипты навыка: до 20 файлов общим размером до 10 МБ | [optional] 

## Methods

### NewCreateSkill

`func NewCreateSkill(name string, whenToUse string, instruction string, ) *CreateSkill`

NewCreateSkill instantiates a new CreateSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSkillWithDefaults

`func NewCreateSkillWithDefaults() *CreateSkill`

NewCreateSkillWithDefaults instantiates a new CreateSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSkill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSkill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSkill) SetName(v string)`

SetName sets Name field to given value.


### GetWhenToUse

`func (o *CreateSkill) GetWhenToUse() string`

GetWhenToUse returns the WhenToUse field if non-nil, zero value otherwise.

### GetWhenToUseOk

`func (o *CreateSkill) GetWhenToUseOk() (*string, bool)`

GetWhenToUseOk returns a tuple with the WhenToUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenToUse

`func (o *CreateSkill) SetWhenToUse(v string)`

SetWhenToUse sets WhenToUse field to given value.


### GetInstruction

`func (o *CreateSkill) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *CreateSkill) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *CreateSkill) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.


### GetScripts

`func (o *CreateSkill) GetScripts() []SkillScriptInput`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *CreateSkill) GetScriptsOk() (*[]SkillScriptInput, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *CreateSkill) SetScripts(v []SkillScriptInput)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *CreateSkill) HasScripts() bool`

HasScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


