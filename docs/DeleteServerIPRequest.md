# DeleteServerIPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP-адрес (IPv4 или IPv6) | 

## Methods

### NewDeleteServerIPRequest

`func NewDeleteServerIPRequest(ip string, ) *DeleteServerIPRequest`

NewDeleteServerIPRequest instantiates a new DeleteServerIPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteServerIPRequestWithDefaults

`func NewDeleteServerIPRequestWithDefaults() *DeleteServerIPRequest`

NewDeleteServerIPRequestWithDefaults instantiates a new DeleteServerIPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *DeleteServerIPRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DeleteServerIPRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DeleteServerIPRequest) SetIp(v string)`

SetIp sets Ip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


