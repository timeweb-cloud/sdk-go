# AddGithub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | **string** | Тип провайдера. | 
**ProviderToken** | **string** | Токен доступа. &lt;br&gt; Для GitHub необходимо использовать &#39;Fine-grained personal access token&#39;. Инструкции по созданию можно найти в &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.GitHub.com/ru/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token&#39;&gt;документации GitHub&lt;/a&gt;. &lt;br&gt; Выберите репозитории, к которым хотите предоставить доступ, и установите следующие разрешения: &#x60;Webhooks: read and write&#x60;, &#x60;Contents: read-only&#x60;. | 

## Methods

### NewAddGithub

`func NewAddGithub(providerType string, providerToken string, ) *AddGithub`

NewAddGithub instantiates a new AddGithub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGithubWithDefaults

`func NewAddGithubWithDefaults() *AddGithub`

NewAddGithubWithDefaults instantiates a new AddGithub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderType

`func (o *AddGithub) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *AddGithub) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *AddGithub) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetProviderToken

`func (o *AddGithub) GetProviderToken() string`

GetProviderToken returns the ProviderToken field if non-nil, zero value otherwise.

### GetProviderTokenOk

`func (o *AddGithub) GetProviderTokenOk() (*string, bool)`

GetProviderTokenOk returns a tuple with the ProviderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderToken

`func (o *AddGithub) SetProviderToken(v string)`

SetProviderToken sets ProviderToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


