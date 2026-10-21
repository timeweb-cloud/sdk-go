# InstallTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **int32** | ID шаблона маркетплейса | 
**Name** | Pointer to **string** | Своё название инструмента. Если не задано, используется название шаблона | [optional] 
**Description** | Pointer to **string** | Своё описание инструмента. Если не задано, используется описание шаблона | [optional] 
**ConfigValues** | Pointer to **map[string]string** | Значения полей из &#x60;config_schema&#x60; шаблона: ключ — &#x60;fields[].key&#x60;, значение — строка. Поля с &#x60;required: true&#x60; обязательны | [optional] 
**AuthData** | Pointer to [**ToolAuthData**](ToolAuthData.md) |  | [optional] 

## Methods

### NewInstallTool

`func NewInstallTool(templateId int32, ) *InstallTool`

NewInstallTool instantiates a new InstallTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallToolWithDefaults

`func NewInstallToolWithDefaults() *InstallTool`

NewInstallToolWithDefaults instantiates a new InstallTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *InstallTool) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *InstallTool) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *InstallTool) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.


### GetName

`func (o *InstallTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstallTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstallTool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstallTool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InstallTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstallTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstallTool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstallTool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfigValues

`func (o *InstallTool) GetConfigValues() map[string]string`

GetConfigValues returns the ConfigValues field if non-nil, zero value otherwise.

### GetConfigValuesOk

`func (o *InstallTool) GetConfigValuesOk() (*map[string]string, bool)`

GetConfigValuesOk returns a tuple with the ConfigValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigValues

`func (o *InstallTool) SetConfigValues(v map[string]string)`

SetConfigValues sets ConfigValues field to given value.

### HasConfigValues

`func (o *InstallTool) HasConfigValues() bool`

HasConfigValues returns a boolean if a field has been set.

### GetAuthData

`func (o *InstallTool) GetAuthData() ToolAuthData`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *InstallTool) GetAuthDataOk() (*ToolAuthData, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *InstallTool) SetAuthData(v ToolAuthData)`

SetAuthData sets AuthData field to given value.

### HasAuthData

`func (o *InstallTool) HasAuthData() bool`

HasAuthData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


