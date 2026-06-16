# NetworksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**RouterNetworks** | [**[]NetworkOut**](NetworkOut.md) | Сети роутера | 
**Meta** | [**ComponentsSchemasMeta**](ComponentsSchemasMeta.md) |  | 

## Methods

### NewNetworksResponse

`func NewNetworksResponse(routerNetworks []NetworkOut, meta ComponentsSchemasMeta, ) *NetworksResponse`

NewNetworksResponse instantiates a new NetworksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksResponseWithDefaults

`func NewNetworksResponseWithDefaults() *NetworksResponse`

NewNetworksResponseWithDefaults instantiates a new NetworksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *NetworksResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *NetworksResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *NetworksResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *NetworksResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetRouterNetworks

`func (o *NetworksResponse) GetRouterNetworks() []NetworkOut`

GetRouterNetworks returns the RouterNetworks field if non-nil, zero value otherwise.

### GetRouterNetworksOk

`func (o *NetworksResponse) GetRouterNetworksOk() (*[]NetworkOut, bool)`

GetRouterNetworksOk returns a tuple with the RouterNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterNetworks

`func (o *NetworksResponse) SetRouterNetworks(v []NetworkOut)`

SetRouterNetworks sets RouterNetworks field to given value.


### GetMeta

`func (o *NetworksResponse) GetMeta() ComponentsSchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworksResponse) GetMetaOk() (*ComponentsSchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworksResponse) SetMeta(v ComponentsSchemasMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


