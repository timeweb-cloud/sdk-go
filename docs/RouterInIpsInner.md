# RouterInIpsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP-адрес. Можно передать значение &#x60;create_ip&#x60; для добавления нового IP-адреса | 
**Nat** | Pointer to [**RouterInIpsInnerNat**](RouterInIpsInnerNat.md) |  | [optional] 

## Methods

### NewRouterInIpsInner

`func NewRouterInIpsInner(ip string, ) *RouterInIpsInner`

NewRouterInIpsInner instantiates a new RouterInIpsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterInIpsInnerWithDefaults

`func NewRouterInIpsInnerWithDefaults() *RouterInIpsInner`

NewRouterInIpsInnerWithDefaults instantiates a new RouterInIpsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *RouterInIpsInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RouterInIpsInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RouterInIpsInner) SetIp(v string)`

SetIp sets Ip field to given value.


### GetNat

`func (o *RouterInIpsInner) GetNat() RouterInIpsInnerNat`

GetNat returns the Nat field if non-nil, zero value otherwise.

### GetNatOk

`func (o *RouterInIpsInner) GetNatOk() (*RouterInIpsInnerNat, bool)`

GetNatOk returns a tuple with the Nat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNat

`func (o *RouterInIpsInner) SetNat(v RouterInIpsInnerNat)`

SetNat sets Nat field to given value.

### HasNat

`func (o *RouterInIpsInner) HasNat() bool`

HasNat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


