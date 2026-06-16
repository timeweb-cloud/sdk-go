# StaticRouteOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID статического маршрута | 
**Nexthop** | **string** | IP-адрес следующего узла | 
**Subnet** | **string** | Сеть назначения в формате CIDR | 

## Methods

### NewStaticRouteOut

`func NewStaticRouteOut(id string, nexthop string, subnet string, ) *StaticRouteOut`

NewStaticRouteOut instantiates a new StaticRouteOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticRouteOutWithDefaults

`func NewStaticRouteOutWithDefaults() *StaticRouteOut`

NewStaticRouteOutWithDefaults instantiates a new StaticRouteOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StaticRouteOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StaticRouteOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StaticRouteOut) SetId(v string)`

SetId sets Id field to given value.


### GetNexthop

`func (o *StaticRouteOut) GetNexthop() string`

GetNexthop returns the Nexthop field if non-nil, zero value otherwise.

### GetNexthopOk

`func (o *StaticRouteOut) GetNexthopOk() (*string, bool)`

GetNexthopOk returns a tuple with the Nexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthop

`func (o *StaticRouteOut) SetNexthop(v string)`

SetNexthop sets Nexthop field to given value.


### GetSubnet

`func (o *StaticRouteOut) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *StaticRouteOut) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *StaticRouteOut) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


