# VpcPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID порта. | 
**NatMode** | **string** | Тип преобразования сетевых адресов. | 
**Mac** | **string** | MAC адрес. | 
**Ipv4** | **string** | Внутренний адрес. | 
**Service** | [**VpcPortService**](VpcPortService.md) |  | 

## Methods

### NewVpcPort

`func NewVpcPort(id string, natMode string, mac string, ipv4 string, service VpcPortService, ) *VpcPort`

NewVpcPort instantiates a new VpcPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcPortWithDefaults

`func NewVpcPortWithDefaults() *VpcPort`

NewVpcPortWithDefaults instantiates a new VpcPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpcPort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpcPort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpcPort) SetId(v string)`

SetId sets Id field to given value.


### GetNatMode

`func (o *VpcPort) GetNatMode() string`

GetNatMode returns the NatMode field if non-nil, zero value otherwise.

### GetNatModeOk

`func (o *VpcPort) GetNatModeOk() (*string, bool)`

GetNatModeOk returns a tuple with the NatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatMode

`func (o *VpcPort) SetNatMode(v string)`

SetNatMode sets NatMode field to given value.


### GetMac

`func (o *VpcPort) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VpcPort) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VpcPort) SetMac(v string)`

SetMac sets Mac field to given value.


### GetIpv4

`func (o *VpcPort) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *VpcPort) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *VpcPort) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.


### GetService

`func (o *VpcPort) GetService() VpcPortService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *VpcPort) GetServiceOk() (*VpcPortService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *VpcPort) SetService(v VpcPortService)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


