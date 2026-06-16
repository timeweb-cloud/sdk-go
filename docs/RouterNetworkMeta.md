# RouterNetworkMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID сети | 
**Name** | **string** | Имя сети | 
**NatIp** | **NullableString** | IP-адрес NAT | 
**Gateway** | **NullableString** | Шлюз | 
**Dhcp** | [**RouterNetworkMetaDhcp**](RouterNetworkMetaDhcp.md) |  | 

## Methods

### NewRouterNetworkMeta

`func NewRouterNetworkMeta(id string, name string, natIp NullableString, gateway NullableString, dhcp RouterNetworkMetaDhcp, ) *RouterNetworkMeta`

NewRouterNetworkMeta instantiates a new RouterNetworkMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterNetworkMetaWithDefaults

`func NewRouterNetworkMetaWithDefaults() *RouterNetworkMeta`

NewRouterNetworkMetaWithDefaults instantiates a new RouterNetworkMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouterNetworkMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouterNetworkMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouterNetworkMeta) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RouterNetworkMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouterNetworkMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouterNetworkMeta) SetName(v string)`

SetName sets Name field to given value.


### GetNatIp

`func (o *RouterNetworkMeta) GetNatIp() string`

GetNatIp returns the NatIp field if non-nil, zero value otherwise.

### GetNatIpOk

`func (o *RouterNetworkMeta) GetNatIpOk() (*string, bool)`

GetNatIpOk returns a tuple with the NatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatIp

`func (o *RouterNetworkMeta) SetNatIp(v string)`

SetNatIp sets NatIp field to given value.


### SetNatIpNil

`func (o *RouterNetworkMeta) SetNatIpNil(b bool)`

 SetNatIpNil sets the value for NatIp to be an explicit nil

### UnsetNatIp
`func (o *RouterNetworkMeta) UnsetNatIp()`

UnsetNatIp ensures that no value is present for NatIp, not even an explicit nil
### GetGateway

`func (o *RouterNetworkMeta) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *RouterNetworkMeta) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *RouterNetworkMeta) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### SetGatewayNil

`func (o *RouterNetworkMeta) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *RouterNetworkMeta) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetDhcp

`func (o *RouterNetworkMeta) GetDhcp() RouterNetworkMetaDhcp`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *RouterNetworkMeta) GetDhcpOk() (*RouterNetworkMetaDhcp, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *RouterNetworkMeta) SetDhcp(v RouterNetworkMetaDhcp)`

SetDhcp sets Dhcp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


