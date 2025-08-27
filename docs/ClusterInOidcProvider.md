# ClusterInOidcProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название создаваемого подключения. Используется только для идентификации и не влияет на остальные параметры | 
**IssuerUrl** | **string** | Адрес OIDC-провайдера, используемый для аутентификации пользователей, запрашивающих доступ к кластеру | 
**ClientId** | **string** | Идентификатор сервиса, выданный OIDC-провайдером, от имени которого осуществляется запрос к ресурсам | 
**UsernameClaim** | Pointer to **string** | Поле в JSON Web Token (JWT), используемое для идентификации пользователя | [optional] 
**GroupsClaim** | Pointer to **string** | Поле в JSON Web Token (JWT), содержащее названии группы, к которой принадлежит пользователь | [optional] 

## Methods

### NewClusterInOidcProvider

`func NewClusterInOidcProvider(name string, issuerUrl string, clientId string, ) *ClusterInOidcProvider`

NewClusterInOidcProvider instantiates a new ClusterInOidcProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInOidcProviderWithDefaults

`func NewClusterInOidcProviderWithDefaults() *ClusterInOidcProvider`

NewClusterInOidcProviderWithDefaults instantiates a new ClusterInOidcProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterInOidcProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterInOidcProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterInOidcProvider) SetName(v string)`

SetName sets Name field to given value.


### GetIssuerUrl

`func (o *ClusterInOidcProvider) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *ClusterInOidcProvider) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *ClusterInOidcProvider) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.


### GetClientId

`func (o *ClusterInOidcProvider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClusterInOidcProvider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClusterInOidcProvider) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetUsernameClaim

`func (o *ClusterInOidcProvider) GetUsernameClaim() string`

GetUsernameClaim returns the UsernameClaim field if non-nil, zero value otherwise.

### GetUsernameClaimOk

`func (o *ClusterInOidcProvider) GetUsernameClaimOk() (*string, bool)`

GetUsernameClaimOk returns a tuple with the UsernameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameClaim

`func (o *ClusterInOidcProvider) SetUsernameClaim(v string)`

SetUsernameClaim sets UsernameClaim field to given value.

### HasUsernameClaim

`func (o *ClusterInOidcProvider) HasUsernameClaim() bool`

HasUsernameClaim returns a boolean if a field has been set.

### GetGroupsClaim

`func (o *ClusterInOidcProvider) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *ClusterInOidcProvider) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *ClusterInOidcProvider) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *ClusterInOidcProvider) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


