# DnatRuleOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID правила | 
**LocalIp** | **string** | Приватный IP-адрес | 
**LocalPort** | **NullableString** | Внутренний порт или диапазон | 
**PublicIp** | **string** | Публичный IP-адрес | 
**PublicPort** | **NullableString** | Внешний порт или диапазон | 
**Protocol** | **string** | Протокол | 

## Methods

### NewDnatRuleOut

`func NewDnatRuleOut(id string, localIp string, localPort NullableString, publicIp string, publicPort NullableString, protocol string, ) *DnatRuleOut`

NewDnatRuleOut instantiates a new DnatRuleOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnatRuleOutWithDefaults

`func NewDnatRuleOutWithDefaults() *DnatRuleOut`

NewDnatRuleOutWithDefaults instantiates a new DnatRuleOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnatRuleOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnatRuleOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnatRuleOut) SetId(v string)`

SetId sets Id field to given value.


### GetLocalIp

`func (o *DnatRuleOut) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *DnatRuleOut) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *DnatRuleOut) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.


### GetLocalPort

`func (o *DnatRuleOut) GetLocalPort() string`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *DnatRuleOut) GetLocalPortOk() (*string, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *DnatRuleOut) SetLocalPort(v string)`

SetLocalPort sets LocalPort field to given value.


### SetLocalPortNil

`func (o *DnatRuleOut) SetLocalPortNil(b bool)`

 SetLocalPortNil sets the value for LocalPort to be an explicit nil

### UnsetLocalPort
`func (o *DnatRuleOut) UnsetLocalPort()`

UnsetLocalPort ensures that no value is present for LocalPort, not even an explicit nil
### GetPublicIp

`func (o *DnatRuleOut) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *DnatRuleOut) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *DnatRuleOut) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.


### GetPublicPort

`func (o *DnatRuleOut) GetPublicPort() string`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *DnatRuleOut) GetPublicPortOk() (*string, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *DnatRuleOut) SetPublicPort(v string)`

SetPublicPort sets PublicPort field to given value.


### SetPublicPortNil

`func (o *DnatRuleOut) SetPublicPortNil(b bool)`

 SetPublicPortNil sets the value for PublicPort to be an explicit nil

### UnsetPublicPort
`func (o *DnatRuleOut) UnsetPublicPort()`

UnsetPublicPort ensures that no value is present for PublicPort, not even an explicit nil
### GetProtocol

`func (o *DnatRuleOut) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DnatRuleOut) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DnatRuleOut) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


