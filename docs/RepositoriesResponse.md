# RepositoriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**Meta** | [**Meta1**](Meta1.md) |  | 
**ContainerRegistriesRepositories** | [**[]ContainerRegistryRepositoriesInner**](ContainerRegistryRepositoriesInner.md) | Массив тарифов container registry | 

## Methods

### NewRepositoriesResponse

`func NewRepositoriesResponse(meta Meta1, containerRegistriesRepositories []ContainerRegistryRepositoriesInner, ) *RepositoriesResponse`

NewRepositoriesResponse instantiates a new RepositoriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoriesResponseWithDefaults

`func NewRepositoriesResponseWithDefaults() *RepositoriesResponse`

NewRepositoriesResponseWithDefaults instantiates a new RepositoriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *RepositoriesResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *RepositoriesResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *RepositoriesResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *RepositoriesResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *RepositoriesResponse) GetMeta() Meta1`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RepositoriesResponse) GetMetaOk() (*Meta1, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RepositoriesResponse) SetMeta(v Meta1)`

SetMeta sets Meta field to given value.


### GetContainerRegistriesRepositories

`func (o *RepositoriesResponse) GetContainerRegistriesRepositories() []ContainerRegistryRepositoriesInner`

GetContainerRegistriesRepositories returns the ContainerRegistriesRepositories field if non-nil, zero value otherwise.

### GetContainerRegistriesRepositoriesOk

`func (o *RepositoriesResponse) GetContainerRegistriesRepositoriesOk() (*[]ContainerRegistryRepositoriesInner, bool)`

GetContainerRegistriesRepositoriesOk returns a tuple with the ContainerRegistriesRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRegistriesRepositories

`func (o *RepositoriesResponse) SetContainerRegistriesRepositories(v []ContainerRegistryRepositoriesInner)`

SetContainerRegistriesRepositories sets ContainerRegistriesRepositories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


