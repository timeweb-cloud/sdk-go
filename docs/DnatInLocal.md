# DnatInLocal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP-адрес | 
**Port** | Pointer to **string** | Порт или диапазон портов | [optional] 

## Methods

### NewDnatInLocal

`func NewDnatInLocal(ip string, ) *DnatInLocal`

NewDnatInLocal instantiates a new DnatInLocal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnatInLocalWithDefaults

`func NewDnatInLocalWithDefaults() *DnatInLocal`

NewDnatInLocalWithDefaults instantiates a new DnatInLocal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *DnatInLocal) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DnatInLocal) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DnatInLocal) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPort

`func (o *DnatInLocal) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DnatInLocal) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DnatInLocal) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *DnatInLocal) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


