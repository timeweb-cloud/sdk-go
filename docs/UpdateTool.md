# UpdateTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название инструмента | [optional] 
**Description** | Pointer to **string** | Описание инструмента | [optional] 
**Type** | Pointer to **string** | Тип инструмента | [optional] 
**Protocol** | Pointer to **string** | Транспорт MCP | [optional] 
**Url** | Pointer to **string** | Адрес MCP-сервера. Список методов при смене адреса не обновляется — вызовите обновление методов | [optional] 
**AuthData** | Pointer to [**ToolAuthData**](ToolAuthData.md) |  | [optional] 
**Status** | Pointer to **string** | Статус инструмента. Чтобы временно выключить инструмент у всех агентов, передайте &#x60;inactive&#x60;, чтобы включить — &#x60;active&#x60; | [optional] 

## Methods

### NewUpdateTool

`func NewUpdateTool() *UpdateTool`

NewUpdateTool instantiates a new UpdateTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateToolWithDefaults

`func NewUpdateToolWithDefaults() *UpdateTool`

NewUpdateToolWithDefaults instantiates a new UpdateTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *UpdateTool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateTool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateTool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateTool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProtocol

`func (o *UpdateTool) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateTool) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateTool) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *UpdateTool) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateTool) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateTool) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateTool) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateTool) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAuthData

`func (o *UpdateTool) GetAuthData() ToolAuthData`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *UpdateTool) GetAuthDataOk() (*ToolAuthData, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *UpdateTool) SetAuthData(v ToolAuthData)`

SetAuthData sets AuthData field to given value.

### HasAuthData

`func (o *UpdateTool) HasAuthData() bool`

HasAuthData returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateTool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateTool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateTool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateTool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


