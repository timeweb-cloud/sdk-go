# GetAppDeploys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Deploys** | Pointer to [**[]Deploy**](Deploy.md) |  | [optional] 

## Methods

### NewGetAppDeploys200Response

`func NewGetAppDeploys200Response(meta Meta, ) *GetAppDeploys200Response`

NewGetAppDeploys200Response instantiates a new GetAppDeploys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppDeploys200ResponseWithDefaults

`func NewGetAppDeploys200ResponseWithDefaults() *GetAppDeploys200Response`

NewGetAppDeploys200ResponseWithDefaults instantiates a new GetAppDeploys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetAppDeploys200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAppDeploys200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAppDeploys200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDeploys

`func (o *GetAppDeploys200Response) GetDeploys() []Deploy`

GetDeploys returns the Deploys field if non-nil, zero value otherwise.

### GetDeploysOk

`func (o *GetAppDeploys200Response) GetDeploysOk() (*[]Deploy, bool)`

GetDeploysOk returns a tuple with the Deploys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploys

`func (o *GetAppDeploys200Response) SetDeploys(v []Deploy)`

SetDeploys sets Deploys field to given value.

### HasDeploys

`func (o *GetAppDeploys200Response) HasDeploys() bool`

HasDeploys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


