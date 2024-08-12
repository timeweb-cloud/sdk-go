# GetRepositories200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Repositories** | [**[]Repository**](Repository.md) |  | 

## Methods

### NewGetRepositories200Response

`func NewGetRepositories200Response(meta Meta, repositories []Repository, ) *GetRepositories200Response`

NewGetRepositories200Response instantiates a new GetRepositories200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRepositories200ResponseWithDefaults

`func NewGetRepositories200ResponseWithDefaults() *GetRepositories200Response`

NewGetRepositories200ResponseWithDefaults instantiates a new GetRepositories200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetRepositories200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRepositories200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRepositories200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetRepositories

`func (o *GetRepositories200Response) GetRepositories() []Repository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *GetRepositories200Response) GetRepositoriesOk() (*[]Repository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *GetRepositories200Response) SetRepositories(v []Repository)`

SetRepositories sets Repositories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


