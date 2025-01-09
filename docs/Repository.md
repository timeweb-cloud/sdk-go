# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID репозитория. | 
**Name** | **string** | Имя репозитория. | 
**FullName** | **string** | Полное имя репозитория. | 
**Url** | **string** | Ссылка на репозиторий. | 
**IsPrivate** | **bool** | Является ли репозиторий приватным. | 
**IsAllowedWebhook** | **bool** | Доступно ли создание вебхука для функции автодеплойя. | 

## Methods

### NewRepository

`func NewRepository(id string, name string, fullName string, url string, isPrivate bool, isAllowedWebhook bool, ) *Repository`

NewRepository instantiates a new Repository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWithDefaults

`func NewRepositoryWithDefaults() *Repository`

NewRepositoryWithDefaults instantiates a new Repository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Repository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Repository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Repository) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Repository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Repository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Repository) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *Repository) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Repository) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Repository) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetUrl

`func (o *Repository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Repository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Repository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetIsPrivate

`func (o *Repository) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *Repository) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *Repository) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.


### GetIsAllowedWebhook

`func (o *Repository) GetIsAllowedWebhook() bool`

GetIsAllowedWebhook returns the IsAllowedWebhook field if non-nil, zero value otherwise.

### GetIsAllowedWebhookOk

`func (o *Repository) GetIsAllowedWebhookOk() (*bool, bool)`

GetIsAllowedWebhookOk returns a tuple with the IsAllowedWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowedWebhook

`func (o *Repository) SetIsAllowedWebhook(v bool)`

SetIsAllowedWebhook sets IsAllowedWebhook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


