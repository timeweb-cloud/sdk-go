# RouterOutIpsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP-адрес | 
**Nat** | [**NullableRouterOutIpsInnerNat**](RouterOutIpsInnerNat.md) |  | 

## Methods

### NewRouterOutIpsInner

`func NewRouterOutIpsInner(ip string, nat NullableRouterOutIpsInnerNat, ) *RouterOutIpsInner`

NewRouterOutIpsInner instantiates a new RouterOutIpsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterOutIpsInnerWithDefaults

`func NewRouterOutIpsInnerWithDefaults() *RouterOutIpsInner`

NewRouterOutIpsInnerWithDefaults instantiates a new RouterOutIpsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *RouterOutIpsInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RouterOutIpsInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RouterOutIpsInner) SetIp(v string)`

SetIp sets Ip field to given value.


### GetNat

`func (o *RouterOutIpsInner) GetNat() RouterOutIpsInnerNat`

GetNat returns the Nat field if non-nil, zero value otherwise.

### GetNatOk

`func (o *RouterOutIpsInner) GetNatOk() (*RouterOutIpsInnerNat, bool)`

GetNatOk returns a tuple with the Nat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNat

`func (o *RouterOutIpsInner) SetNat(v RouterOutIpsInnerNat)`

SetNat sets Nat field to given value.


### SetNatNil

`func (o *RouterOutIpsInner) SetNatNil(b bool)`

 SetNatNil sets the value for Nat to be an explicit nil

### UnsetNat
`func (o *RouterOutIpsInner) UnsetNat()`

UnsetNat ensures that no value is present for Nat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


