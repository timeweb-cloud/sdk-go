# SkillScriptInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя файла: латиница, цифры, &#x60;_&#x60;, &#x60;.&#x60; и &#x60;-&#x60;, без путей, с расширением &#x60;.py&#x60;, &#x60;.js&#x60;, &#x60;.mjs&#x60; или &#x60;.cjs&#x60;. Имена в навыке не повторяются | 
**Description** | **string** | Что делает скрипт. По этому тексту модель решает, запускать ли его, поэтому описывайте результат, а не реализацию | 
**ArgsHint** | Pointer to **string** | Какие аргументы принимает скрипт | [optional] 
**Content** | Pointer to **string** | Текст скрипта. Обязателен для нового файла. Для файла, который уже есть в текущей версии навыка, можно не передавать: содержимое возьмётся из неё по имени. Общий размер скриптов навыка — до 10 МБ (10 485 760 байт в UTF-8) | [optional] 

## Methods

### NewSkillScriptInput

`func NewSkillScriptInput(name string, description string, ) *SkillScriptInput`

NewSkillScriptInput instantiates a new SkillScriptInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillScriptInputWithDefaults

`func NewSkillScriptInputWithDefaults() *SkillScriptInput`

NewSkillScriptInputWithDefaults instantiates a new SkillScriptInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SkillScriptInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkillScriptInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkillScriptInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SkillScriptInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkillScriptInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkillScriptInput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArgsHint

`func (o *SkillScriptInput) GetArgsHint() string`

GetArgsHint returns the ArgsHint field if non-nil, zero value otherwise.

### GetArgsHintOk

`func (o *SkillScriptInput) GetArgsHintOk() (*string, bool)`

GetArgsHintOk returns a tuple with the ArgsHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgsHint

`func (o *SkillScriptInput) SetArgsHint(v string)`

SetArgsHint sets ArgsHint field to given value.

### HasArgsHint

`func (o *SkillScriptInput) HasArgsHint() bool`

HasArgsHint returns a boolean if a field has been set.

### GetContent

`func (o *SkillScriptInput) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SkillScriptInput) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SkillScriptInput) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SkillScriptInput) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


