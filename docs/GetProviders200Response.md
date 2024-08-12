# GetProviders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Providers** | [**[]Provider**](Provider.md) |  | 

## Methods

### NewGetProviders200Response

`func NewGetProviders200Response(meta Meta, providers []Provider, ) *GetProviders200Response`

NewGetProviders200Response instantiates a new GetProviders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProviders200ResponseWithDefaults

`func NewGetProviders200ResponseWithDefaults() *GetProviders200Response`

NewGetProviders200ResponseWithDefaults instantiates a new GetProviders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetProviders200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetProviders200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetProviders200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetProviders

`func (o *GetProviders200Response) GetProviders() []Provider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *GetProviders200Response) GetProvidersOk() (*[]Provider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *GetProviders200Response) SetProviders(v []Provider)`

SetProviders sets Providers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


