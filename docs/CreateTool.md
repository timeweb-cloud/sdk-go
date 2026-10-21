# CreateTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название инструмента | 
**Description** | Pointer to **string** | Описание инструмента | [optional] 
**Type** | **string** | Тип инструмента. Поддерживается только &#x60;MCP&#x60; | 
**Protocol** | **string** | Транспорт MCP: &#x60;http&#x60; — Streamable HTTP, &#x60;sse&#x60; — Server-Sent Events | 
**Url** | **string** | Адрес MCP-сервера. Сервер должен быть доступен из интернета | 
**AuthData** | Pointer to [**ToolAuthData**](ToolAuthData.md) |  | [optional] 

## Methods

### NewCreateTool

`func NewCreateTool(name string, type_ string, protocol string, url string, ) *CreateTool`

NewCreateTool instantiates a new CreateTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateToolWithDefaults

`func NewCreateToolWithDefaults() *CreateTool`

NewCreateToolWithDefaults instantiates a new CreateTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTool) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateTool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTool) SetType(v string)`

SetType sets Type field to given value.


### GetProtocol

`func (o *CreateTool) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateTool) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateTool) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetUrl

`func (o *CreateTool) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateTool) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateTool) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthData

`func (o *CreateTool) GetAuthData() ToolAuthData`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *CreateTool) GetAuthDataOk() (*ToolAuthData, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *CreateTool) SetAuthData(v ToolAuthData)`

SetAuthData sets AuthData field to given value.

### HasAuthData

`func (o *CreateTool) HasAuthData() bool`

HasAuthData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


