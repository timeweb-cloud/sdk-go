# VdsNetworksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Уникальный идентификатор сети. Есть только у приватных сетей. | [optional] 
**Type** | **string** | Тип сети. | 
**NatMode** | Pointer to **string** | Тип преобразования сетевых адресов. | [optional] 
**Bandwidth** | Pointer to **NullableFloat32** | Пропускная способность сети. | [optional] 
**Ips** | [**[]VdsNetworksInnerIpsInner**](VdsNetworksInnerIpsInner.md) | Список IP-адресов сети. | 
**IsDdosGuard** | Pointer to **bool** | Подключена ли DDoS-защита. Только для публичных сетей. | [optional] 

## Methods

### NewVdsNetworksInner

`func NewVdsNetworksInner(type_ string, ips []VdsNetworksInnerIpsInner, ) *VdsNetworksInner`

NewVdsNetworksInner instantiates a new VdsNetworksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdsNetworksInnerWithDefaults

`func NewVdsNetworksInnerWithDefaults() *VdsNetworksInner`

NewVdsNetworksInnerWithDefaults instantiates a new VdsNetworksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VdsNetworksInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VdsNetworksInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VdsNetworksInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VdsNetworksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *VdsNetworksInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VdsNetworksInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VdsNetworksInner) SetType(v string)`

SetType sets Type field to given value.


### GetNatMode

`func (o *VdsNetworksInner) GetNatMode() string`

GetNatMode returns the NatMode field if non-nil, zero value otherwise.

### GetNatModeOk

`func (o *VdsNetworksInner) GetNatModeOk() (*string, bool)`

GetNatModeOk returns a tuple with the NatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatMode

`func (o *VdsNetworksInner) SetNatMode(v string)`

SetNatMode sets NatMode field to given value.

### HasNatMode

`func (o *VdsNetworksInner) HasNatMode() bool`

HasNatMode returns a boolean if a field has been set.

### GetBandwidth

`func (o *VdsNetworksInner) GetBandwidth() float32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *VdsNetworksInner) GetBandwidthOk() (*float32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *VdsNetworksInner) SetBandwidth(v float32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *VdsNetworksInner) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### SetBandwidthNil

`func (o *VdsNetworksInner) SetBandwidthNil(b bool)`

 SetBandwidthNil sets the value for Bandwidth to be an explicit nil

### UnsetBandwidth
`func (o *VdsNetworksInner) UnsetBandwidth()`

UnsetBandwidth ensures that no value is present for Bandwidth, not even an explicit nil
### GetIps

`func (o *VdsNetworksInner) GetIps() []VdsNetworksInnerIpsInner`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *VdsNetworksInner) GetIpsOk() (*[]VdsNetworksInnerIpsInner, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *VdsNetworksInner) SetIps(v []VdsNetworksInnerIpsInner)`

SetIps sets Ips field to given value.


### SetIpsNil

`func (o *VdsNetworksInner) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *VdsNetworksInner) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil
### GetIsDdosGuard

`func (o *VdsNetworksInner) GetIsDdosGuard() bool`

GetIsDdosGuard returns the IsDdosGuard field if non-nil, zero value otherwise.

### GetIsDdosGuardOk

`func (o *VdsNetworksInner) GetIsDdosGuardOk() (*bool, bool)`

GetIsDdosGuardOk returns a tuple with the IsDdosGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDdosGuard

`func (o *VdsNetworksInner) SetIsDdosGuard(v bool)`

SetIsDdosGuard sets IsDdosGuard field to given value.

### HasIsDdosGuard

`func (o *VdsNetworksInner) HasIsDdosGuard() bool`

HasIsDdosGuard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


