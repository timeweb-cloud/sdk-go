# AddBitbucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | **string** | Тип провайдера. | 
**ProviderToken** | **string** | Токен доступа. &lt;br&gt; Для Bitbucket необходимо использовать &#39;App password&#39;. Инструкции по созданию можно найти в &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://support.atlassian.com/bitbucket-cloud/docs/create-an-app-password/&#39;&gt;документации Bitbucket&lt;/a&gt;. &lt;br&gt; Установите следующие разрешения: &#x60;Account: Read&#x60;, &#x60;Projects: Read&#x60;, &#x60;Repositories: Read&#x60;, &#x60;Webhooks: Read and write&#x60; | 
**Login** | **string** | Логин пользователя Bitbucket. | 

## Methods

### NewAddBitbucket

`func NewAddBitbucket(providerType string, providerToken string, login string, ) *AddBitbucket`

NewAddBitbucket instantiates a new AddBitbucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBitbucketWithDefaults

`func NewAddBitbucketWithDefaults() *AddBitbucket`

NewAddBitbucketWithDefaults instantiates a new AddBitbucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderType

`func (o *AddBitbucket) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *AddBitbucket) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *AddBitbucket) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetProviderToken

`func (o *AddBitbucket) GetProviderToken() string`

GetProviderToken returns the ProviderToken field if non-nil, zero value otherwise.

### GetProviderTokenOk

`func (o *AddBitbucket) GetProviderTokenOk() (*string, bool)`

GetProviderTokenOk returns a tuple with the ProviderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderToken

`func (o *AddBitbucket) SetProviderToken(v string)`

SetProviderToken sets ProviderToken field to given value.


### GetLogin

`func (o *AddBitbucket) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *AddBitbucket) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *AddBitbucket) SetLogin(v string)`

SetLogin sets Login field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


