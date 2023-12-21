# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Уникальный идентификатор сети. | 
**FloatingIp** | Pointer to **string** | Плавающий IP-адрес | [optional] 
**Ip** | Pointer to **string** | IP-адрес в сети. | [optional] 

## Methods

### NewNetwork

`func NewNetwork(id string, ) *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Network) SetId(v string)`

SetId sets Id field to given value.


### GetFloatingIp

`func (o *Network) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *Network) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *Network) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *Network) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetIp

`func (o *Network) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Network) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Network) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Network) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


