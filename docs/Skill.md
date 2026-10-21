# Skill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID навыка | 
**Name** | **string** | Название навыка, уникальное в аккаунте | 
**WhenToUse** | **string** | Когда применять навык. Это поле читает модель, когда выбирает навык | 
**Instruction** | **string** | Инструкция навыка. Попадает в контекст модели только при срабатывании | 
**CurrentVersion** | **int32** | Номер текущей версии | 
**AgentsCount** | **int32** | У скольких агентов навык включён | 
**FirstAgentName** | **NullableString** | Имя агента, если навык включён ровно у одного агента | 
**CatalogItemId** | **NullableInt32** | ID навыка в каталоге, из которого установлен этот навык | 
**CatalogVersion** | **NullableInt32** | Версия каталожного навыка, из которой установлен или обновлён этот навык | 
**IsFromCatalog** | **bool** | Навык установлен из каталога | 
**IsModified** | **bool** | Установленный из каталога навык изменялся владельцем | 
**CreatedAt** | **time.Time** | Дата и время создания | 
**UpdatedAt** | **time.Time** | Дата и время последнего изменения | 

## Methods

### NewSkill

`func NewSkill(id int32, name string, whenToUse string, instruction string, currentVersion int32, agentsCount int32, firstAgentName NullableString, catalogItemId NullableInt32, catalogVersion NullableInt32, isFromCatalog bool, isModified bool, createdAt time.Time, updatedAt time.Time, ) *Skill`

NewSkill instantiates a new Skill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillWithDefaults

`func NewSkillWithDefaults() *Skill`

NewSkillWithDefaults instantiates a new Skill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Skill) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Skill) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Skill) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Skill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Skill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Skill) SetName(v string)`

SetName sets Name field to given value.


### GetWhenToUse

`func (o *Skill) GetWhenToUse() string`

GetWhenToUse returns the WhenToUse field if non-nil, zero value otherwise.

### GetWhenToUseOk

`func (o *Skill) GetWhenToUseOk() (*string, bool)`

GetWhenToUseOk returns a tuple with the WhenToUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenToUse

`func (o *Skill) SetWhenToUse(v string)`

SetWhenToUse sets WhenToUse field to given value.


### GetInstruction

`func (o *Skill) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *Skill) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *Skill) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.


### GetCurrentVersion

`func (o *Skill) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *Skill) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *Skill) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.


### GetAgentsCount

`func (o *Skill) GetAgentsCount() int32`

GetAgentsCount returns the AgentsCount field if non-nil, zero value otherwise.

### GetAgentsCountOk

`func (o *Skill) GetAgentsCountOk() (*int32, bool)`

GetAgentsCountOk returns a tuple with the AgentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentsCount

`func (o *Skill) SetAgentsCount(v int32)`

SetAgentsCount sets AgentsCount field to given value.


### GetFirstAgentName

`func (o *Skill) GetFirstAgentName() string`

GetFirstAgentName returns the FirstAgentName field if non-nil, zero value otherwise.

### GetFirstAgentNameOk

`func (o *Skill) GetFirstAgentNameOk() (*string, bool)`

GetFirstAgentNameOk returns a tuple with the FirstAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAgentName

`func (o *Skill) SetFirstAgentName(v string)`

SetFirstAgentName sets FirstAgentName field to given value.


### SetFirstAgentNameNil

`func (o *Skill) SetFirstAgentNameNil(b bool)`

 SetFirstAgentNameNil sets the value for FirstAgentName to be an explicit nil

### UnsetFirstAgentName
`func (o *Skill) UnsetFirstAgentName()`

UnsetFirstAgentName ensures that no value is present for FirstAgentName, not even an explicit nil
### GetCatalogItemId

`func (o *Skill) GetCatalogItemId() int32`

GetCatalogItemId returns the CatalogItemId field if non-nil, zero value otherwise.

### GetCatalogItemIdOk

`func (o *Skill) GetCatalogItemIdOk() (*int32, bool)`

GetCatalogItemIdOk returns a tuple with the CatalogItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemId

`func (o *Skill) SetCatalogItemId(v int32)`

SetCatalogItemId sets CatalogItemId field to given value.


### SetCatalogItemIdNil

`func (o *Skill) SetCatalogItemIdNil(b bool)`

 SetCatalogItemIdNil sets the value for CatalogItemId to be an explicit nil

### UnsetCatalogItemId
`func (o *Skill) UnsetCatalogItemId()`

UnsetCatalogItemId ensures that no value is present for CatalogItemId, not even an explicit nil
### GetCatalogVersion

`func (o *Skill) GetCatalogVersion() int32`

GetCatalogVersion returns the CatalogVersion field if non-nil, zero value otherwise.

### GetCatalogVersionOk

`func (o *Skill) GetCatalogVersionOk() (*int32, bool)`

GetCatalogVersionOk returns a tuple with the CatalogVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogVersion

`func (o *Skill) SetCatalogVersion(v int32)`

SetCatalogVersion sets CatalogVersion field to given value.


### SetCatalogVersionNil

`func (o *Skill) SetCatalogVersionNil(b bool)`

 SetCatalogVersionNil sets the value for CatalogVersion to be an explicit nil

### UnsetCatalogVersion
`func (o *Skill) UnsetCatalogVersion()`

UnsetCatalogVersion ensures that no value is present for CatalogVersion, not even an explicit nil
### GetIsFromCatalog

`func (o *Skill) GetIsFromCatalog() bool`

GetIsFromCatalog returns the IsFromCatalog field if non-nil, zero value otherwise.

### GetIsFromCatalogOk

`func (o *Skill) GetIsFromCatalogOk() (*bool, bool)`

GetIsFromCatalogOk returns a tuple with the IsFromCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromCatalog

`func (o *Skill) SetIsFromCatalog(v bool)`

SetIsFromCatalog sets IsFromCatalog field to given value.


### GetIsModified

`func (o *Skill) GetIsModified() bool`

GetIsModified returns the IsModified field if non-nil, zero value otherwise.

### GetIsModifiedOk

`func (o *Skill) GetIsModifiedOk() (*bool, bool)`

GetIsModifiedOk returns a tuple with the IsModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModified

`func (o *Skill) SetIsModified(v bool)`

SetIsModified sets IsModified field to given value.


### GetCreatedAt

`func (o *Skill) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Skill) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Skill) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Skill) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Skill) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Skill) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


