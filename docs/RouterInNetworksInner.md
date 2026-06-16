# RouterInNetworksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID локальной сети | 
**Gateway** | Pointer to **string** | IP-адрес шлюза. При отсутствии подставляется первый свободный IP в подсети | [optional] 
**ReservedIps** | Pointer to **[]string** | Зарезервированные IP-адреса. При отсутствии — пустой массив | [optional] 
**IsDhcpEnabled** | Pointer to **bool** | Включен ли DHCP | [optional] 

## Methods

### NewRouterInNetworksInner

`func NewRouterInNetworksInner(id string, ) *RouterInNetworksInner`

NewRouterInNetworksInner instantiates a new RouterInNetworksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterInNetworksInnerWithDefaults

`func NewRouterInNetworksInnerWithDefaults() *RouterInNetworksInner`

NewRouterInNetworksInnerWithDefaults instantiates a new RouterInNetworksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouterInNetworksInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouterInNetworksInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouterInNetworksInner) SetId(v string)`

SetId sets Id field to given value.


### GetGateway

`func (o *RouterInNetworksInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *RouterInNetworksInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *RouterInNetworksInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *RouterInNetworksInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetReservedIps

`func (o *RouterInNetworksInner) GetReservedIps() []string`

GetReservedIps returns the ReservedIps field if non-nil, zero value otherwise.

### GetReservedIpsOk

`func (o *RouterInNetworksInner) GetReservedIpsOk() (*[]string, bool)`

GetReservedIpsOk returns a tuple with the ReservedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIps

`func (o *RouterInNetworksInner) SetReservedIps(v []string)`

SetReservedIps sets ReservedIps field to given value.

### HasReservedIps

`func (o *RouterInNetworksInner) HasReservedIps() bool`

HasReservedIps returns a boolean if a field has been set.

### GetIsDhcpEnabled

`func (o *RouterInNetworksInner) GetIsDhcpEnabled() bool`

GetIsDhcpEnabled returns the IsDhcpEnabled field if non-nil, zero value otherwise.

### GetIsDhcpEnabledOk

`func (o *RouterInNetworksInner) GetIsDhcpEnabledOk() (*bool, bool)`

GetIsDhcpEnabledOk returns a tuple with the IsDhcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDhcpEnabled

`func (o *RouterInNetworksInner) SetIsDhcpEnabled(v bool)`

SetIsDhcpEnabled sets IsDhcpEnabled field to given value.

### HasIsDhcpEnabled

`func (o *RouterInNetworksInner) HasIsDhcpEnabled() bool`

HasIsDhcpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


