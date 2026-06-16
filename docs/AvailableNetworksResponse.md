# AvailableNetworksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**AvailableNetworks** | [**[]AvailableNetwork**](AvailableNetwork.md) | Доступные сети | 
**Meta** | [**ComponentsSchemasMeta**](ComponentsSchemasMeta.md) |  | 

## Methods

### NewAvailableNetworksResponse

`func NewAvailableNetworksResponse(availableNetworks []AvailableNetwork, meta ComponentsSchemasMeta, ) *AvailableNetworksResponse`

NewAvailableNetworksResponse instantiates a new AvailableNetworksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableNetworksResponseWithDefaults

`func NewAvailableNetworksResponseWithDefaults() *AvailableNetworksResponse`

NewAvailableNetworksResponseWithDefaults instantiates a new AvailableNetworksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *AvailableNetworksResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *AvailableNetworksResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *AvailableNetworksResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *AvailableNetworksResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetAvailableNetworks

`func (o *AvailableNetworksResponse) GetAvailableNetworks() []AvailableNetwork`

GetAvailableNetworks returns the AvailableNetworks field if non-nil, zero value otherwise.

### GetAvailableNetworksOk

`func (o *AvailableNetworksResponse) GetAvailableNetworksOk() (*[]AvailableNetwork, bool)`

GetAvailableNetworksOk returns a tuple with the AvailableNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNetworks

`func (o *AvailableNetworksResponse) SetAvailableNetworks(v []AvailableNetwork)`

SetAvailableNetworks sets AvailableNetworks field to given value.


### GetMeta

`func (o *AvailableNetworksResponse) GetMeta() ComponentsSchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AvailableNetworksResponse) GetMetaOk() (*ComponentsSchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AvailableNetworksResponse) SetMeta(v ComponentsSchemasMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


