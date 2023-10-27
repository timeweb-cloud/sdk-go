# VdsNetworksInnerIpsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип IP-адреса сети | 
**Ip** | **string** | IP-адрес сети. | 
**Ptr** | Pointer to **string** | Запись имени узла. | [optional] 
**IsMain** | **bool** | Является ли сеть основной. | 

## Methods

### NewVdsNetworksInnerIpsInner

`func NewVdsNetworksInnerIpsInner(type_ string, ip string, isMain bool, ) *VdsNetworksInnerIpsInner`

NewVdsNetworksInnerIpsInner instantiates a new VdsNetworksInnerIpsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdsNetworksInnerIpsInnerWithDefaults

`func NewVdsNetworksInnerIpsInnerWithDefaults() *VdsNetworksInnerIpsInner`

NewVdsNetworksInnerIpsInnerWithDefaults instantiates a new VdsNetworksInnerIpsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VdsNetworksInnerIpsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VdsNetworksInnerIpsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VdsNetworksInnerIpsInner) SetType(v string)`

SetType sets Type field to given value.


### GetIp

`func (o *VdsNetworksInnerIpsInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *VdsNetworksInnerIpsInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *VdsNetworksInnerIpsInner) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPtr

`func (o *VdsNetworksInnerIpsInner) GetPtr() string`

GetPtr returns the Ptr field if non-nil, zero value otherwise.

### GetPtrOk

`func (o *VdsNetworksInnerIpsInner) GetPtrOk() (*string, bool)`

GetPtrOk returns a tuple with the Ptr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtr

`func (o *VdsNetworksInnerIpsInner) SetPtr(v string)`

SetPtr sets Ptr field to given value.

### HasPtr

`func (o *VdsNetworksInnerIpsInner) HasPtr() bool`

HasPtr returns a boolean if a field has been set.

### GetIsMain

`func (o *VdsNetworksInnerIpsInner) GetIsMain() bool`

GetIsMain returns the IsMain field if non-nil, zero value otherwise.

### GetIsMainOk

`func (o *VdsNetworksInnerIpsInner) GetIsMainOk() (*bool, bool)`

GetIsMainOk returns a tuple with the IsMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMain

`func (o *VdsNetworksInnerIpsInner) SetIsMain(v bool)`

SetIsMain sets IsMain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


