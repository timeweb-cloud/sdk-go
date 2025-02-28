# Vpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID сети. | 
**Name** | **string** | Имя сети. | 
**SubnetV4** | **string** | Маска подсети. | 
**Location** | **string** | Локация сети. | 
**CreatedAt** | **time.Time** | Дата создания сети. | 
**Description** | **string** | Описание. | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | 
**PublicIp** | **NullableString** | Публичный IP-адрес сети. | 
**Type** | **string** | Тип сети. | 
**BusyAddress** | **[]string** | Занятые адреса в сети | 

## Methods

### NewVpc

`func NewVpc(id string, name string, subnetV4 string, location string, createdAt time.Time, description string, availabilityZone AvailabilityZone, publicIp NullableString, type_ string, busyAddress []string, ) *Vpc`

NewVpc instantiates a new Vpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcWithDefaults

`func NewVpcWithDefaults() *Vpc`

NewVpcWithDefaults instantiates a new Vpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vpc) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vpc) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vpc) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Vpc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vpc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vpc) SetName(v string)`

SetName sets Name field to given value.


### GetSubnetV4

`func (o *Vpc) GetSubnetV4() string`

GetSubnetV4 returns the SubnetV4 field if non-nil, zero value otherwise.

### GetSubnetV4Ok

`func (o *Vpc) GetSubnetV4Ok() (*string, bool)`

GetSubnetV4Ok returns a tuple with the SubnetV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetV4

`func (o *Vpc) SetSubnetV4(v string)`

SetSubnetV4 sets SubnetV4 field to given value.


### GetLocation

`func (o *Vpc) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Vpc) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Vpc) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetCreatedAt

`func (o *Vpc) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Vpc) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Vpc) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *Vpc) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Vpc) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Vpc) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAvailabilityZone

`func (o *Vpc) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *Vpc) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *Vpc) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetPublicIp

`func (o *Vpc) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *Vpc) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *Vpc) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.


### SetPublicIpNil

`func (o *Vpc) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *Vpc) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetType

`func (o *Vpc) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Vpc) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Vpc) SetType(v string)`

SetType sets Type field to given value.


### GetBusyAddress

`func (o *Vpc) GetBusyAddress() []string`

GetBusyAddress returns the BusyAddress field if non-nil, zero value otherwise.

### GetBusyAddressOk

`func (o *Vpc) GetBusyAddressOk() (*[]string, bool)`

GetBusyAddressOk returns a tuple with the BusyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyAddress

`func (o *Vpc) SetBusyAddress(v []string)`

SetBusyAddress sets BusyAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


