# AddGit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | **string** | Тип провайдера. | 
**Url** | **string** | Внешний url адрес репозитория. | 
**Login** | Pointer to **string** | Логин провайдера. Используется при добавлении приватных репозиториев | [optional] 
**Password** | Pointer to **string** | Пароль или токен провайдера. Обязателен при добавлении приватных репозиториев. Подробнее в нашей &lt;a href&#x3D;&#39;https://timeweb.cloud/docs/apps/connecting-repositories#podkluchenie-repozitoriya-po-ssylke&#39;&gt;документации&lt;/a&gt;. | [optional] 

## Methods

### NewAddGit

`func NewAddGit(providerType string, url string, ) *AddGit`

NewAddGit instantiates a new AddGit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGitWithDefaults

`func NewAddGitWithDefaults() *AddGit`

NewAddGitWithDefaults instantiates a new AddGit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderType

`func (o *AddGit) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *AddGit) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *AddGit) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetUrl

`func (o *AddGit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddGit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddGit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLogin

`func (o *AddGit) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *AddGit) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *AddGit) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *AddGit) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetPassword

`func (o *AddGit) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddGit) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddGit) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddGit) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


