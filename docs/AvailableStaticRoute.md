# AvailableStaticRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** | Имя сервиса | 
**ServiceType** | **string** | Тип сервиса | 
**Nexthop** | **string** | IP-адрес следующего узла | 
**Subnets** | Pointer to [**[]AvailableStaticRouteSubnetsInner**](AvailableStaticRouteSubnetsInner.md) | Доступные подсети | [optional] 

## Methods

### NewAvailableStaticRoute

`func NewAvailableStaticRoute(serviceName string, serviceType string, nexthop string, ) *AvailableStaticRoute`

NewAvailableStaticRoute instantiates a new AvailableStaticRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableStaticRouteWithDefaults

`func NewAvailableStaticRouteWithDefaults() *AvailableStaticRoute`

NewAvailableStaticRouteWithDefaults instantiates a new AvailableStaticRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *AvailableStaticRoute) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *AvailableStaticRoute) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *AvailableStaticRoute) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServiceType

`func (o *AvailableStaticRoute) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *AvailableStaticRoute) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *AvailableStaticRoute) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetNexthop

`func (o *AvailableStaticRoute) GetNexthop() string`

GetNexthop returns the Nexthop field if non-nil, zero value otherwise.

### GetNexthopOk

`func (o *AvailableStaticRoute) GetNexthopOk() (*string, bool)`

GetNexthopOk returns a tuple with the Nexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthop

`func (o *AvailableStaticRoute) SetNexthop(v string)`

SetNexthop sets Nexthop field to given value.


### GetSubnets

`func (o *AvailableStaticRoute) GetSubnets() []AvailableStaticRouteSubnetsInner`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *AvailableStaticRoute) GetSubnetsOk() (*[]AvailableStaticRouteSubnetsInner, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *AvailableStaticRoute) SetSubnets(v []AvailableStaticRouteSubnetsInner)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *AvailableStaticRoute) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


