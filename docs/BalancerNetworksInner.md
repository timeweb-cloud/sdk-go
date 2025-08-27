# BalancerNetworksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID сети. Есть только у приватных сетей. | [optional] 
**Type** | **string** | Тип сети. | 
**NatMode** | Pointer to **string** | Тип преобразования сетевых адресов. | [optional] 
**PortId** | Pointer to **string** | ID порта. | [optional] 
**Ips** | [**[]BalancerNetworksInnerIpsInner**](BalancerNetworksInnerIpsInner.md) | Список IP-адресов сети. | 

## Methods

### NewBalancerNetworksInner

`func NewBalancerNetworksInner(type_ string, ips []BalancerNetworksInnerIpsInner, ) *BalancerNetworksInner`

NewBalancerNetworksInner instantiates a new BalancerNetworksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalancerNetworksInnerWithDefaults

`func NewBalancerNetworksInnerWithDefaults() *BalancerNetworksInner`

NewBalancerNetworksInnerWithDefaults instantiates a new BalancerNetworksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BalancerNetworksInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BalancerNetworksInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BalancerNetworksInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BalancerNetworksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BalancerNetworksInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BalancerNetworksInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BalancerNetworksInner) SetType(v string)`

SetType sets Type field to given value.


### GetNatMode

`func (o *BalancerNetworksInner) GetNatMode() string`

GetNatMode returns the NatMode field if non-nil, zero value otherwise.

### GetNatModeOk

`func (o *BalancerNetworksInner) GetNatModeOk() (*string, bool)`

GetNatModeOk returns a tuple with the NatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatMode

`func (o *BalancerNetworksInner) SetNatMode(v string)`

SetNatMode sets NatMode field to given value.

### HasNatMode

`func (o *BalancerNetworksInner) HasNatMode() bool`

HasNatMode returns a boolean if a field has been set.

### GetPortId

`func (o *BalancerNetworksInner) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *BalancerNetworksInner) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *BalancerNetworksInner) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *BalancerNetworksInner) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetIps

`func (o *BalancerNetworksInner) GetIps() []BalancerNetworksInnerIpsInner`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *BalancerNetworksInner) GetIpsOk() (*[]BalancerNetworksInnerIpsInner, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *BalancerNetworksInner) SetIps(v []BalancerNetworksInnerIpsInner)`

SetIps sets Ips field to given value.


### SetIpsNil

`func (o *BalancerNetworksInner) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *BalancerNetworksInner) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


