# SkillScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя файла | 
**Description** | **string** | Что делает скрипт | 
**ArgsHint** | **NullableString** | Какие аргументы принимает скрипт | 
**SizeBytes** | **int32** | Размер содержимого в байтах | 

## Methods

### NewSkillScript

`func NewSkillScript(name string, description string, argsHint NullableString, sizeBytes int32, ) *SkillScript`

NewSkillScript instantiates a new SkillScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillScriptWithDefaults

`func NewSkillScriptWithDefaults() *SkillScript`

NewSkillScriptWithDefaults instantiates a new SkillScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SkillScript) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkillScript) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkillScript) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SkillScript) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkillScript) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkillScript) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArgsHint

`func (o *SkillScript) GetArgsHint() string`

GetArgsHint returns the ArgsHint field if non-nil, zero value otherwise.

### GetArgsHintOk

`func (o *SkillScript) GetArgsHintOk() (*string, bool)`

GetArgsHintOk returns a tuple with the ArgsHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgsHint

`func (o *SkillScript) SetArgsHint(v string)`

SetArgsHint sets ArgsHint field to given value.


### SetArgsHintNil

`func (o *SkillScript) SetArgsHintNil(b bool)`

 SetArgsHintNil sets the value for ArgsHint to be an explicit nil

### UnsetArgsHint
`func (o *SkillScript) UnsetArgsHint()`

UnsetArgsHint ensures that no value is present for ArgsHint, not even an explicit nil
### GetSizeBytes

`func (o *SkillScript) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *SkillScript) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *SkillScript) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


