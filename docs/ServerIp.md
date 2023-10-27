# ServerIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип IP-адреса сети. | 
**Ip** | **string** | IP-адрес сети. | 
**Ptr** | **string** | Запись имени узла. | 
**IsMain** | **bool** | Является ли сеть основной. | 

## Methods

### NewServerIp

`func NewServerIp(type_ string, ip string, ptr string, isMain bool, ) *ServerIp`

NewServerIp instantiates a new ServerIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerIpWithDefaults

`func NewServerIpWithDefaults() *ServerIp`

NewServerIpWithDefaults instantiates a new ServerIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServerIp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerIp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerIp) SetType(v string)`

SetType sets Type field to given value.


### GetIp

`func (o *ServerIp) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ServerIp) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ServerIp) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPtr

`func (o *ServerIp) GetPtr() string`

GetPtr returns the Ptr field if non-nil, zero value otherwise.

### GetPtrOk

`func (o *ServerIp) GetPtrOk() (*string, bool)`

GetPtrOk returns a tuple with the Ptr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtr

`func (o *ServerIp) SetPtr(v string)`

SetPtr sets Ptr field to given value.


### GetIsMain

`func (o *ServerIp) GetIsMain() bool`

GetIsMain returns the IsMain field if non-nil, zero value otherwise.

### GetIsMainOk

`func (o *ServerIp) GetIsMainOk() (*bool, bool)`

GetIsMainOk returns a tuple with the IsMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMain

`func (o *ServerIp) SetIsMain(v bool)`

SetIsMain sets IsMain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


