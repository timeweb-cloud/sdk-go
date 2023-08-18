# UpdateServerIPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP-адрес (IPv4 или IPv6) | 
**Ptr** | **string** | PTR-запись IP-адреса | 

## Methods

### NewUpdateServerIPRequest

`func NewUpdateServerIPRequest(ip string, ptr string, ) *UpdateServerIPRequest`

NewUpdateServerIPRequest instantiates a new UpdateServerIPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerIPRequestWithDefaults

`func NewUpdateServerIPRequestWithDefaults() *UpdateServerIPRequest`

NewUpdateServerIPRequestWithDefaults instantiates a new UpdateServerIPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *UpdateServerIPRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *UpdateServerIPRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *UpdateServerIPRequest) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPtr

`func (o *UpdateServerIPRequest) GetPtr() string`

GetPtr returns the Ptr field if non-nil, zero value otherwise.

### GetPtrOk

`func (o *UpdateServerIPRequest) GetPtrOk() (*string, bool)`

GetPtrOk returns a tuple with the Ptr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtr

`func (o *UpdateServerIPRequest) SetPtr(v string)`

SetPtr sets Ptr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


