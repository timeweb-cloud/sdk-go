# ToolAuthDataParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | Токен для заголовка &#x60;Authorization: Bearer&#x60; (для &#x60;bearer&#x60;). В ответах заменён на &#x60;********&#x60; | [optional] 
**Headers** | Pointer to **map[string]string** | HTTP-заголовки, которые передаются MCP-серверу (для &#x60;custom_headers&#x60;). В ответах значения заменены на &#x60;********&#x60; | [optional] 
**ClientId** | Pointer to **string** | ID OAuth2-клиента (для &#x60;oauth2&#x60;). Не нужен, если сервер поддерживает динамическую регистрацию клиентов | [optional] 
**ClientSecret** | Pointer to **string** | Секрет OAuth2-клиента (для &#x60;oauth2&#x60;). В ответах заменён на &#x60;********&#x60; | [optional] 
**AuthorizationUrl** | Pointer to **string** | Адрес страницы авторизации провайдера (для &#x60;oauth2&#x60;). Если не задан, определяется по метаданным сервера | [optional] 
**TokenUrl** | Pointer to **string** | Адрес получения токенов провайдера (для &#x60;oauth2&#x60;). Если не задан, определяется по метаданным сервера | [optional] 
**RegistrationUrl** | Pointer to **string** | Адрес динамической регистрации клиента (для &#x60;oauth2&#x60;) | [optional] 
**Scopes** | Pointer to **[]string** | Запрашиваемые права (для &#x60;oauth2&#x60;). Если не заданы, берутся из метаданных сервера | [optional] 

## Methods

### NewToolAuthDataParams

`func NewToolAuthDataParams() *ToolAuthDataParams`

NewToolAuthDataParams instantiates a new ToolAuthDataParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolAuthDataParamsWithDefaults

`func NewToolAuthDataParamsWithDefaults() *ToolAuthDataParams`

NewToolAuthDataParamsWithDefaults instantiates a new ToolAuthDataParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ToolAuthDataParams) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ToolAuthDataParams) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ToolAuthDataParams) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ToolAuthDataParams) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetHeaders

`func (o *ToolAuthDataParams) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ToolAuthDataParams) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ToolAuthDataParams) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ToolAuthDataParams) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetClientId

`func (o *ToolAuthDataParams) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ToolAuthDataParams) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ToolAuthDataParams) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ToolAuthDataParams) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ToolAuthDataParams) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ToolAuthDataParams) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ToolAuthDataParams) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ToolAuthDataParams) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetAuthorizationUrl

`func (o *ToolAuthDataParams) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *ToolAuthDataParams) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *ToolAuthDataParams) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.

### HasAuthorizationUrl

`func (o *ToolAuthDataParams) HasAuthorizationUrl() bool`

HasAuthorizationUrl returns a boolean if a field has been set.

### GetTokenUrl

`func (o *ToolAuthDataParams) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *ToolAuthDataParams) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *ToolAuthDataParams) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *ToolAuthDataParams) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### GetRegistrationUrl

`func (o *ToolAuthDataParams) GetRegistrationUrl() string`

GetRegistrationUrl returns the RegistrationUrl field if non-nil, zero value otherwise.

### GetRegistrationUrlOk

`func (o *ToolAuthDataParams) GetRegistrationUrlOk() (*string, bool)`

GetRegistrationUrlOk returns a tuple with the RegistrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationUrl

`func (o *ToolAuthDataParams) SetRegistrationUrl(v string)`

SetRegistrationUrl sets RegistrationUrl field to given value.

### HasRegistrationUrl

`func (o *ToolAuthDataParams) HasRegistrationUrl() bool`

HasRegistrationUrl returns a boolean if a field has been set.

### GetScopes

`func (o *ToolAuthDataParams) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ToolAuthDataParams) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ToolAuthDataParams) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ToolAuthDataParams) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


