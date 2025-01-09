# NetworkDriversResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**Meta** | [**SchemasMeta**](SchemasMeta.md) |  | 
**NetworkDrivers** | **[]string** | Массив сетевых драйверов k8s | 

## Methods

### NewNetworkDriversResponse

`func NewNetworkDriversResponse(meta SchemasMeta, networkDrivers []string, ) *NetworkDriversResponse`

NewNetworkDriversResponse instantiates a new NetworkDriversResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDriversResponseWithDefaults

`func NewNetworkDriversResponseWithDefaults() *NetworkDriversResponse`

NewNetworkDriversResponseWithDefaults instantiates a new NetworkDriversResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *NetworkDriversResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *NetworkDriversResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *NetworkDriversResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *NetworkDriversResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *NetworkDriversResponse) GetMeta() SchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkDriversResponse) GetMetaOk() (*SchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkDriversResponse) SetMeta(v SchemasMeta)`

SetMeta sets Meta field to given value.


### GetNetworkDrivers

`func (o *NetworkDriversResponse) GetNetworkDrivers() []string`

GetNetworkDrivers returns the NetworkDrivers field if non-nil, zero value otherwise.

### GetNetworkDriversOk

`func (o *NetworkDriversResponse) GetNetworkDriversOk() (*[]string, bool)`

GetNetworkDriversOk returns a tuple with the NetworkDrivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDrivers

`func (o *NetworkDriversResponse) SetNetworkDrivers(v []string)`

SetNetworkDrivers sets NetworkDrivers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


