# SkillCatalogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID навыка в каталоге | 
**Key** | **string** | Уникальный ключ навыка в каталоге | 
**Category** | **string** | Назначение: &#x60;communication&#x60; — манера общения, &#x60;support&#x60; — регламенты обслуживания, &#x60;sales&#x60; — продажи, &#x60;content&#x60; — документы и база знаний, &#x60;safety&#x60; — ограничения и безопасность | 
**Name** | **string** | Название навыка | 
**Description** | **NullableString** | Краткое описание навыка | 
**WhenToUse** | **string** | Когда применять навык | 
**Instruction** | **string** | Полная инструкция навыка | 
**Version** | **int32** | Версия навыка в каталоге | 
**IsInstalled** | **bool** | Навык уже установлен в аккаунте | 
**InstalledSkillId** | **NullableInt32** | ID установленного навыка в аккаунте | 
**IsModified** | **bool** | Установленный навык изменялся владельцем | 
**HasCatalogUpdate** | **bool** | В каталоге есть версия новее установленной | 

## Methods

### NewSkillCatalogItem

`func NewSkillCatalogItem(id int32, key string, category string, name string, description NullableString, whenToUse string, instruction string, version int32, isInstalled bool, installedSkillId NullableInt32, isModified bool, hasCatalogUpdate bool, ) *SkillCatalogItem`

NewSkillCatalogItem instantiates a new SkillCatalogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillCatalogItemWithDefaults

`func NewSkillCatalogItemWithDefaults() *SkillCatalogItem`

NewSkillCatalogItemWithDefaults instantiates a new SkillCatalogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SkillCatalogItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkillCatalogItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkillCatalogItem) SetId(v int32)`

SetId sets Id field to given value.


### GetKey

`func (o *SkillCatalogItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SkillCatalogItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SkillCatalogItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetCategory

`func (o *SkillCatalogItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SkillCatalogItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SkillCatalogItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetName

`func (o *SkillCatalogItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkillCatalogItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkillCatalogItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SkillCatalogItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkillCatalogItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkillCatalogItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *SkillCatalogItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SkillCatalogItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWhenToUse

`func (o *SkillCatalogItem) GetWhenToUse() string`

GetWhenToUse returns the WhenToUse field if non-nil, zero value otherwise.

### GetWhenToUseOk

`func (o *SkillCatalogItem) GetWhenToUseOk() (*string, bool)`

GetWhenToUseOk returns a tuple with the WhenToUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenToUse

`func (o *SkillCatalogItem) SetWhenToUse(v string)`

SetWhenToUse sets WhenToUse field to given value.


### GetInstruction

`func (o *SkillCatalogItem) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *SkillCatalogItem) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *SkillCatalogItem) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.


### GetVersion

`func (o *SkillCatalogItem) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SkillCatalogItem) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SkillCatalogItem) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetIsInstalled

`func (o *SkillCatalogItem) GetIsInstalled() bool`

GetIsInstalled returns the IsInstalled field if non-nil, zero value otherwise.

### GetIsInstalledOk

`func (o *SkillCatalogItem) GetIsInstalledOk() (*bool, bool)`

GetIsInstalledOk returns a tuple with the IsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInstalled

`func (o *SkillCatalogItem) SetIsInstalled(v bool)`

SetIsInstalled sets IsInstalled field to given value.


### GetInstalledSkillId

`func (o *SkillCatalogItem) GetInstalledSkillId() int32`

GetInstalledSkillId returns the InstalledSkillId field if non-nil, zero value otherwise.

### GetInstalledSkillIdOk

`func (o *SkillCatalogItem) GetInstalledSkillIdOk() (*int32, bool)`

GetInstalledSkillIdOk returns a tuple with the InstalledSkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledSkillId

`func (o *SkillCatalogItem) SetInstalledSkillId(v int32)`

SetInstalledSkillId sets InstalledSkillId field to given value.


### SetInstalledSkillIdNil

`func (o *SkillCatalogItem) SetInstalledSkillIdNil(b bool)`

 SetInstalledSkillIdNil sets the value for InstalledSkillId to be an explicit nil

### UnsetInstalledSkillId
`func (o *SkillCatalogItem) UnsetInstalledSkillId()`

UnsetInstalledSkillId ensures that no value is present for InstalledSkillId, not even an explicit nil
### GetIsModified

`func (o *SkillCatalogItem) GetIsModified() bool`

GetIsModified returns the IsModified field if non-nil, zero value otherwise.

### GetIsModifiedOk

`func (o *SkillCatalogItem) GetIsModifiedOk() (*bool, bool)`

GetIsModifiedOk returns a tuple with the IsModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModified

`func (o *SkillCatalogItem) SetIsModified(v bool)`

SetIsModified sets IsModified field to given value.


### GetHasCatalogUpdate

`func (o *SkillCatalogItem) GetHasCatalogUpdate() bool`

GetHasCatalogUpdate returns the HasCatalogUpdate field if non-nil, zero value otherwise.

### GetHasCatalogUpdateOk

`func (o *SkillCatalogItem) GetHasCatalogUpdateOk() (*bool, bool)`

GetHasCatalogUpdateOk returns a tuple with the HasCatalogUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCatalogUpdate

`func (o *SkillCatalogItem) SetHasCatalogUpdate(v bool)`

SetHasCatalogUpdate sets HasCatalogUpdate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


