# NetworkOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID сети | 
**Name** | **string** | Имя сети | 
**Subnet** | **string** | Подсеть | 
**NatIp** | **NullableString** | IP-адрес NAT | 
**Gateway** | **NullableString** | Шлюз | 
**ReservedIps** | **[]string** | Зарезервированные IP-адреса | 
**Dhcp** | [**RouterNetworkMetaDhcp**](RouterNetworkMetaDhcp.md) |  | 
**BusyAddresses** | **[]string** | Занятые IP-адреса | 

## Methods

### NewNetworkOut

`func NewNetworkOut(id string, name string, subnet string, natIp NullableString, gateway NullableString, reservedIps []string, dhcp RouterNetworkMetaDhcp, busyAddresses []string, ) *NetworkOut`

NewNetworkOut instantiates a new NetworkOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkOutWithDefaults

`func NewNetworkOutWithDefaults() *NetworkOut`

NewNetworkOutWithDefaults instantiates a new NetworkOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NetworkOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkOut) SetName(v string)`

SetName sets Name field to given value.


### GetSubnet

`func (o *NetworkOut) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *NetworkOut) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *NetworkOut) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetNatIp

`func (o *NetworkOut) GetNatIp() string`

GetNatIp returns the NatIp field if non-nil, zero value otherwise.

### GetNatIpOk

`func (o *NetworkOut) GetNatIpOk() (*string, bool)`

GetNatIpOk returns a tuple with the NatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatIp

`func (o *NetworkOut) SetNatIp(v string)`

SetNatIp sets NatIp field to given value.


### SetNatIpNil

`func (o *NetworkOut) SetNatIpNil(b bool)`

 SetNatIpNil sets the value for NatIp to be an explicit nil

### UnsetNatIp
`func (o *NetworkOut) UnsetNatIp()`

UnsetNatIp ensures that no value is present for NatIp, not even an explicit nil
### GetGateway

`func (o *NetworkOut) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkOut) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkOut) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### SetGatewayNil

`func (o *NetworkOut) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *NetworkOut) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetReservedIps

`func (o *NetworkOut) GetReservedIps() []string`

GetReservedIps returns the ReservedIps field if non-nil, zero value otherwise.

### GetReservedIpsOk

`func (o *NetworkOut) GetReservedIpsOk() (*[]string, bool)`

GetReservedIpsOk returns a tuple with the ReservedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIps

`func (o *NetworkOut) SetReservedIps(v []string)`

SetReservedIps sets ReservedIps field to given value.


### GetDhcp

`func (o *NetworkOut) GetDhcp() RouterNetworkMetaDhcp`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *NetworkOut) GetDhcpOk() (*RouterNetworkMetaDhcp, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *NetworkOut) SetDhcp(v RouterNetworkMetaDhcp)`

SetDhcp sets Dhcp field to given value.


### GetBusyAddresses

`func (o *NetworkOut) GetBusyAddresses() []string`

GetBusyAddresses returns the BusyAddresses field if non-nil, zero value otherwise.

### GetBusyAddressesOk

`func (o *NetworkOut) GetBusyAddressesOk() (*[]string, bool)`

GetBusyAddressesOk returns a tuple with the BusyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyAddresses

`func (o *NetworkOut) SetBusyAddresses(v []string)`

SetBusyAddresses sets BusyAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


