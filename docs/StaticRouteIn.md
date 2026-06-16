# StaticRouteIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | **string** | Сеть назначения в формате CIDR | 
**Nexthop** | **string** | IP-адрес следующего узла (IPv4) | 

## Methods

### NewStaticRouteIn

`func NewStaticRouteIn(subnet string, nexthop string, ) *StaticRouteIn`

NewStaticRouteIn instantiates a new StaticRouteIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticRouteInWithDefaults

`func NewStaticRouteInWithDefaults() *StaticRouteIn`

NewStaticRouteInWithDefaults instantiates a new StaticRouteIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *StaticRouteIn) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *StaticRouteIn) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *StaticRouteIn) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetNexthop

`func (o *StaticRouteIn) GetNexthop() string`

GetNexthop returns the Nexthop field if non-nil, zero value otherwise.

### GetNexthopOk

`func (o *StaticRouteIn) GetNexthopOk() (*string, bool)`

GetNexthopOk returns a tuple with the Nexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthop

`func (o *StaticRouteIn) SetNexthop(v string)`

SetNexthop sets Nexthop field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


