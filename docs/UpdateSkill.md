# UpdateSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название навыка. Изменение названия не создаёт новую версию | [optional] 
**WhenToUse** | Pointer to **string** | Когда применять навык | [optional] 
**Instruction** | Pointer to **string** | Инструкция навыка | [optional] 
**Scripts** | Pointer to [**[]SkillScriptInput**](SkillScriptInput.md) | Новый набор скриптов, заменяет прежний целиком. Пустой массив удаляет все скрипты. Если поле не передано, скрипты не меняются | [optional] 

## Methods

### NewUpdateSkill

`func NewUpdateSkill() *UpdateSkill`

NewUpdateSkill instantiates a new UpdateSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSkillWithDefaults

`func NewUpdateSkillWithDefaults() *UpdateSkill`

NewUpdateSkillWithDefaults instantiates a new UpdateSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSkill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSkill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSkill) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSkill) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWhenToUse

`func (o *UpdateSkill) GetWhenToUse() string`

GetWhenToUse returns the WhenToUse field if non-nil, zero value otherwise.

### GetWhenToUseOk

`func (o *UpdateSkill) GetWhenToUseOk() (*string, bool)`

GetWhenToUseOk returns a tuple with the WhenToUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenToUse

`func (o *UpdateSkill) SetWhenToUse(v string)`

SetWhenToUse sets WhenToUse field to given value.

### HasWhenToUse

`func (o *UpdateSkill) HasWhenToUse() bool`

HasWhenToUse returns a boolean if a field has been set.

### GetInstruction

`func (o *UpdateSkill) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *UpdateSkill) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *UpdateSkill) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *UpdateSkill) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### GetScripts

`func (o *UpdateSkill) GetScripts() []SkillScriptInput`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *UpdateSkill) GetScriptsOk() (*[]SkillScriptInput, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *UpdateSkill) SetScripts(v []SkillScriptInput)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *UpdateSkill) HasScripts() bool`

HasScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


