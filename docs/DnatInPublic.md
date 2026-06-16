# DnatInPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP-адрес | 
**Port** | Pointer to **string** | Порт или диапазон портов | [optional] 

## Methods

### NewDnatInPublic

`func NewDnatInPublic(ip string, ) *DnatInPublic`

NewDnatInPublic instantiates a new DnatInPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnatInPublicWithDefaults

`func NewDnatInPublicWithDefaults() *DnatInPublic`

NewDnatInPublicWithDefaults instantiates a new DnatInPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *DnatInPublic) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DnatInPublic) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DnatInPublic) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPort

`func (o *DnatInPublic) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DnatInPublic) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DnatInPublic) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *DnatInPublic) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


