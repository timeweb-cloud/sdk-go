# GetCdnResources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpResources** | [**[]HttpResource**](HttpResource.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetCdnResources200Response

`func NewGetCdnResources200Response(httpResources []HttpResource, meta Meta, ) *GetCdnResources200Response`

NewGetCdnResources200Response instantiates a new GetCdnResources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCdnResources200ResponseWithDefaults

`func NewGetCdnResources200ResponseWithDefaults() *GetCdnResources200Response`

NewGetCdnResources200ResponseWithDefaults instantiates a new GetCdnResources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpResources

`func (o *GetCdnResources200Response) GetHttpResources() []HttpResource`

GetHttpResources returns the HttpResources field if non-nil, zero value otherwise.

### GetHttpResourcesOk

`func (o *GetCdnResources200Response) GetHttpResourcesOk() (*[]HttpResource, bool)`

GetHttpResourcesOk returns a tuple with the HttpResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResources

`func (o *GetCdnResources200Response) SetHttpResources(v []HttpResource)`

SetHttpResources sets HttpResources field to given value.


### GetMeta

`func (o *GetCdnResources200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCdnResources200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCdnResources200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


