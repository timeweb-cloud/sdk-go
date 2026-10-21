# Tool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID инструмента | 
**AccountId** | **string** | Логин аккаунта-владельца | 
**TemplateId** | **NullableInt32** | ID шаблона маркетплейса, из которого установлен инструмент. &#x60;null&#x60; — собственный MCP-сервер | 
**Key** | Pointer to **NullableString** | Ключ шаблона маркетплейса. Сейчас в ответах не возвращается | [optional] 
**Name** | **NullableString** | Название инструмента | 
**Description** | **NullableString** | Описание инструмента | 
**Type** | **NullableString** | Тип инструмента. У инструментов из шаблонов может быть &#x60;null&#x60; | 
**Protocol** | **NullableString** | Транспорт MCP: &#x60;http&#x60; — Streamable HTTP, &#x60;sse&#x60; — Server-Sent Events. У инструментов из шаблонов всегда &#x60;null&#x60; | 
**Url** | **NullableString** | Адрес MCP-сервера. У инструментов из шаблонов всегда &#x60;null&#x60; | 
**Methods** | [**[]ToolMethod**](ToolMethod.md) | Методы MCP-сервера. &#x60;null&#x60; или пустой массив — методы ещё не получены, например до завершения OAuth2-авторизации | 
**AuthData** | [**ToolAuthData**](ToolAuthData.md) |  | 
**Status** | **string** | Статус инструмента: &#x60;active&#x60; — работает, &#x60;inactive&#x60; — выключен вручную, &#x60;pending_auth&#x60; — ожидает OAuth2-авторизации, &#x60;error&#x60; — последнее обновление методов завершилось ошибкой (причина в &#x60;status_info&#x60;). Агент вызывает только инструменты в статусе &#x60;active&#x60; | 
**StatusInfo** | **NullableString** | Причина статуса &#x60;error&#x60; | 
**CreatedAt** | **time.Time** | Дата и время создания | 
**UpdatedAt** | **time.Time** | Дата и время последнего изменения | 

## Methods

### NewTool

`func NewTool(id int32, accountId string, templateId NullableInt32, name NullableString, description NullableString, type_ NullableString, protocol NullableString, url NullableString, methods []ToolMethod, authData ToolAuthData, status string, statusInfo NullableString, createdAt time.Time, updatedAt time.Time, ) *Tool`

NewTool instantiates a new Tool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolWithDefaults

`func NewToolWithDefaults() *Tool`

NewToolWithDefaults instantiates a new Tool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tool) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tool) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tool) SetId(v int32)`

SetId sets Id field to given value.


### GetAccountId

`func (o *Tool) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Tool) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Tool) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTemplateId

`func (o *Tool) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Tool) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Tool) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.


### SetTemplateIdNil

`func (o *Tool) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *Tool) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetKey

`func (o *Tool) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Tool) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Tool) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Tool) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *Tool) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *Tool) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetName

`func (o *Tool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tool) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Tool) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Tool) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Tool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tool) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Tool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Tool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *Tool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Tool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Tool) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *Tool) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Tool) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetProtocol

`func (o *Tool) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Tool) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Tool) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### SetProtocolNil

`func (o *Tool) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *Tool) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetUrl

`func (o *Tool) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Tool) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Tool) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *Tool) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Tool) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetMethods

`func (o *Tool) GetMethods() []ToolMethod`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *Tool) GetMethodsOk() (*[]ToolMethod, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *Tool) SetMethods(v []ToolMethod)`

SetMethods sets Methods field to given value.


### SetMethodsNil

`func (o *Tool) SetMethodsNil(b bool)`

 SetMethodsNil sets the value for Methods to be an explicit nil

### UnsetMethods
`func (o *Tool) UnsetMethods()`

UnsetMethods ensures that no value is present for Methods, not even an explicit nil
### GetAuthData

`func (o *Tool) GetAuthData() ToolAuthData`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *Tool) GetAuthDataOk() (*ToolAuthData, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *Tool) SetAuthData(v ToolAuthData)`

SetAuthData sets AuthData field to given value.


### GetStatus

`func (o *Tool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Tool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Tool) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusInfo

`func (o *Tool) GetStatusInfo() string`

GetStatusInfo returns the StatusInfo field if non-nil, zero value otherwise.

### GetStatusInfoOk

`func (o *Tool) GetStatusInfoOk() (*string, bool)`

GetStatusInfoOk returns a tuple with the StatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfo

`func (o *Tool) SetStatusInfo(v string)`

SetStatusInfo sets StatusInfo field to given value.


### SetStatusInfoNil

`func (o *Tool) SetStatusInfoNil(b bool)`

 SetStatusInfoNil sets the value for StatusInfo to be an explicit nil

### UnsetStatusInfo
`func (o *Tool) UnsetStatusInfo()`

UnsetStatusInfo ensures that no value is present for StatusInfo, not even an explicit nil
### GetCreatedAt

`func (o *Tool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Tool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Tool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Tool) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Tool) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Tool) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


