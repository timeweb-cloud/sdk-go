# OriginNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Subnet** | **NullableString** | IPv4-адрес или подсеть узлов. &#x60;null&#x60;, если у записи нет IPv4-подсети. | 
**Ipv6Subnet** | **NullableString** | IPv6-адрес или подсеть узлов. &#x60;null&#x60;, если у записи нет IPv6-подсети. | 

## Methods

### NewOriginNode

`func NewOriginNode(ipv4Subnet NullableString, ipv6Subnet NullableString, ) *OriginNode`

NewOriginNode instantiates a new OriginNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginNodeWithDefaults

`func NewOriginNodeWithDefaults() *OriginNode`

NewOriginNodeWithDefaults instantiates a new OriginNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Subnet

`func (o *OriginNode) GetIpv4Subnet() string`

GetIpv4Subnet returns the Ipv4Subnet field if non-nil, zero value otherwise.

### GetIpv4SubnetOk

`func (o *OriginNode) GetIpv4SubnetOk() (*string, bool)`

GetIpv4SubnetOk returns a tuple with the Ipv4Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Subnet

`func (o *OriginNode) SetIpv4Subnet(v string)`

SetIpv4Subnet sets Ipv4Subnet field to given value.


### SetIpv4SubnetNil

`func (o *OriginNode) SetIpv4SubnetNil(b bool)`

 SetIpv4SubnetNil sets the value for Ipv4Subnet to be an explicit nil

### UnsetIpv4Subnet
`func (o *OriginNode) UnsetIpv4Subnet()`

UnsetIpv4Subnet ensures that no value is present for Ipv4Subnet, not even an explicit nil
### GetIpv6Subnet

`func (o *OriginNode) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *OriginNode) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *OriginNode) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.


### SetIpv6SubnetNil

`func (o *OriginNode) SetIpv6SubnetNil(b bool)`

 SetIpv6SubnetNil sets the value for Ipv6Subnet to be an explicit nil

### UnsetIpv6Subnet
`func (o *OriginNode) UnsetIpv6Subnet()`

UnsetIpv6Subnet ensures that no value is present for Ipv6Subnet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


