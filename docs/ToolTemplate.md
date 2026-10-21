# ToolTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID шаблона. Передаётся в &#x60;template_id&#x60; при установке | 
**Name** | **string** | Название | 
**Key** | Pointer to **NullableString** | Уникальный ключ шаблона | [optional] 
**Description** | **string** | Описание | 
**Type** | **string** | Тип инструмента | 
**Version** | **string** | Версия шаблона | 
**IsApproved** | **int32** | Шаблон одобрен для маркетплейса: &#x60;1&#x60; — да, &#x60;0&#x60; — нет | 
**AuthorAccountId** | Pointer to **NullableString** | Автор шаблона. &#x60;null&#x60; — шаблон Timeweb Cloud | [optional] 
**IsSystem** | **int32** | Системный шаблон платформы: &#x60;1&#x60; — да, &#x60;0&#x60; — нет | 
**InstallCount** | **int32** | Сколько раз шаблон установлен | 
**Category** | Pointer to **NullableString** | Категория шаблона | [optional] 
**ConfigSchema** | Pointer to [**NullableToolTemplateConfigSchema**](ToolTemplateConfigSchema.md) |  | [optional] 
**AuthType** | **string** | Способ авторизации на MCP-сервере | 
**DeploymentType** | **string** | Где работает MCP-сервер: &#x60;managed&#x60; — на стороне Timeweb Cloud, &#x60;external&#x60; — у стороннего сервиса | 
**CreatedAt** | **time.Time** | Дата и время создания | 
**UpdatedAt** | **time.Time** | Дата и время последнего изменения | 

## Methods

### NewToolTemplate

`func NewToolTemplate(id int32, name string, description string, type_ string, version string, isApproved int32, isSystem int32, installCount int32, authType string, deploymentType string, createdAt time.Time, updatedAt time.Time, ) *ToolTemplate`

NewToolTemplate instantiates a new ToolTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolTemplateWithDefaults

`func NewToolTemplateWithDefaults() *ToolTemplate`

NewToolTemplateWithDefaults instantiates a new ToolTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ToolTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ToolTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ToolTemplate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ToolTemplate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ToolTemplate) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ToolTemplate) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ToolTemplate) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ToolTemplate) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetDescription

`func (o *ToolTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ToolTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolTemplate) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *ToolTemplate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ToolTemplate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ToolTemplate) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetIsApproved

`func (o *ToolTemplate) GetIsApproved() int32`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ToolTemplate) GetIsApprovedOk() (*int32, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ToolTemplate) SetIsApproved(v int32)`

SetIsApproved sets IsApproved field to given value.


### GetAuthorAccountId

`func (o *ToolTemplate) GetAuthorAccountId() string`

GetAuthorAccountId returns the AuthorAccountId field if non-nil, zero value otherwise.

### GetAuthorAccountIdOk

`func (o *ToolTemplate) GetAuthorAccountIdOk() (*string, bool)`

GetAuthorAccountIdOk returns a tuple with the AuthorAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAccountId

`func (o *ToolTemplate) SetAuthorAccountId(v string)`

SetAuthorAccountId sets AuthorAccountId field to given value.

### HasAuthorAccountId

`func (o *ToolTemplate) HasAuthorAccountId() bool`

HasAuthorAccountId returns a boolean if a field has been set.

### SetAuthorAccountIdNil

`func (o *ToolTemplate) SetAuthorAccountIdNil(b bool)`

 SetAuthorAccountIdNil sets the value for AuthorAccountId to be an explicit nil

### UnsetAuthorAccountId
`func (o *ToolTemplate) UnsetAuthorAccountId()`

UnsetAuthorAccountId ensures that no value is present for AuthorAccountId, not even an explicit nil
### GetIsSystem

`func (o *ToolTemplate) GetIsSystem() int32`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *ToolTemplate) GetIsSystemOk() (*int32, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *ToolTemplate) SetIsSystem(v int32)`

SetIsSystem sets IsSystem field to given value.


### GetInstallCount

`func (o *ToolTemplate) GetInstallCount() int32`

GetInstallCount returns the InstallCount field if non-nil, zero value otherwise.

### GetInstallCountOk

`func (o *ToolTemplate) GetInstallCountOk() (*int32, bool)`

GetInstallCountOk returns a tuple with the InstallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallCount

`func (o *ToolTemplate) SetInstallCount(v int32)`

SetInstallCount sets InstallCount field to given value.


### GetCategory

`func (o *ToolTemplate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ToolTemplate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ToolTemplate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ToolTemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ToolTemplate) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ToolTemplate) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetConfigSchema

`func (o *ToolTemplate) GetConfigSchema() ToolTemplateConfigSchema`

GetConfigSchema returns the ConfigSchema field if non-nil, zero value otherwise.

### GetConfigSchemaOk

`func (o *ToolTemplate) GetConfigSchemaOk() (*ToolTemplateConfigSchema, bool)`

GetConfigSchemaOk returns a tuple with the ConfigSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSchema

`func (o *ToolTemplate) SetConfigSchema(v ToolTemplateConfigSchema)`

SetConfigSchema sets ConfigSchema field to given value.

### HasConfigSchema

`func (o *ToolTemplate) HasConfigSchema() bool`

HasConfigSchema returns a boolean if a field has been set.

### SetConfigSchemaNil

`func (o *ToolTemplate) SetConfigSchemaNil(b bool)`

 SetConfigSchemaNil sets the value for ConfigSchema to be an explicit nil

### UnsetConfigSchema
`func (o *ToolTemplate) UnsetConfigSchema()`

UnsetConfigSchema ensures that no value is present for ConfigSchema, not even an explicit nil
### GetAuthType

`func (o *ToolTemplate) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ToolTemplate) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ToolTemplate) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetDeploymentType

`func (o *ToolTemplate) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *ToolTemplate) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *ToolTemplate) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.


### GetCreatedAt

`func (o *ToolTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ToolTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ToolTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ToolTemplate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ToolTemplate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ToolTemplate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


