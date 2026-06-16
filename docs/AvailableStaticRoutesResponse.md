# AvailableStaticRoutesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**AvailableNetworks** | [**[]AvailableStaticRoute**](AvailableStaticRoute.md) | Доступные подсети | 
**Meta** | [**ComponentsSchemasMeta**](ComponentsSchemasMeta.md) |  | 

## Methods

### NewAvailableStaticRoutesResponse

`func NewAvailableStaticRoutesResponse(availableNetworks []AvailableStaticRoute, meta ComponentsSchemasMeta, ) *AvailableStaticRoutesResponse`

NewAvailableStaticRoutesResponse instantiates a new AvailableStaticRoutesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableStaticRoutesResponseWithDefaults

`func NewAvailableStaticRoutesResponseWithDefaults() *AvailableStaticRoutesResponse`

NewAvailableStaticRoutesResponseWithDefaults instantiates a new AvailableStaticRoutesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *AvailableStaticRoutesResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *AvailableStaticRoutesResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *AvailableStaticRoutesResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *AvailableStaticRoutesResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetAvailableNetworks

`func (o *AvailableStaticRoutesResponse) GetAvailableNetworks() []AvailableStaticRoute`

GetAvailableNetworks returns the AvailableNetworks field if non-nil, zero value otherwise.

### GetAvailableNetworksOk

`func (o *AvailableStaticRoutesResponse) GetAvailableNetworksOk() (*[]AvailableStaticRoute, bool)`

GetAvailableNetworksOk returns a tuple with the AvailableNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNetworks

`func (o *AvailableStaticRoutesResponse) SetAvailableNetworks(v []AvailableStaticRoute)`

SetAvailableNetworks sets AvailableNetworks field to given value.


### GetMeta

`func (o *AvailableStaticRoutesResponse) GetMeta() ComponentsSchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AvailableStaticRoutesResponse) GetMetaOk() (*ComponentsSchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AvailableStaticRoutesResponse) SetMeta(v ComponentsSchemasMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


