# AddGitlab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | **string** | Тип провайдера. | 
**ProviderToken** | **string** | Токен доступа. &lt;br&gt; Для GitLab необходимо использовать персональный токен доступа. Инструкции по созданию можно найти в &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://docs.GitLab.com/ee/user/profile/personal_access_tokens.html#create-a-personal-access-token&#39;&gt;документации GitLab&lt;/a&gt;. &lt;br&gt; Установите следующие разрешения: &#x60;api&#x60; | 

## Methods

### NewAddGitlab

`func NewAddGitlab(providerType string, providerToken string, ) *AddGitlab`

NewAddGitlab instantiates a new AddGitlab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGitlabWithDefaults

`func NewAddGitlabWithDefaults() *AddGitlab`

NewAddGitlabWithDefaults instantiates a new AddGitlab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderType

`func (o *AddGitlab) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *AddGitlab) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *AddGitlab) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetProviderToken

`func (o *AddGitlab) GetProviderToken() string`

GetProviderToken returns the ProviderToken field if non-nil, zero value otherwise.

### GetProviderTokenOk

`func (o *AddGitlab) GetProviderTokenOk() (*string, bool)`

GetProviderTokenOk returns a tuple with the ProviderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderToken

`func (o *AddGitlab) SetProviderToken(v string)`

SetProviderToken sets ProviderToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


