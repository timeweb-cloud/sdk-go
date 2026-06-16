# AvailableNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID сети | 
**Description** | **string** | Описание сети | 
**Name** | **string** | Имя сети | 
**SubnetV4** | **string** | Подсеть IPv4 | 
**Location** | **string** | Локация | 
**CreatedAt** | **time.Time** | Дата и время создания | 
**AvailabilityZone** | **string** | Зона доступности | 
**PublicIp** | **NullableString** | Публичный IP-адрес | 
**Type** | **string** | Тип сети | 
**BusyAddress** | **[]string** | Занятые IP-адреса | 

## Methods

### NewAvailableNetwork

`func NewAvailableNetwork(id string, description string, name string, subnetV4 string, location string, createdAt time.Time, availabilityZone string, publicIp NullableString, type_ string, busyAddress []string, ) *AvailableNetwork`

NewAvailableNetwork instantiates a new AvailableNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableNetworkWithDefaults

`func NewAvailableNetworkWithDefaults() *AvailableNetwork`

NewAvailableNetworkWithDefaults instantiates a new AvailableNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AvailableNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AvailableNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AvailableNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *AvailableNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AvailableNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AvailableNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *AvailableNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailableNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailableNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetSubnetV4

`func (o *AvailableNetwork) GetSubnetV4() string`

GetSubnetV4 returns the SubnetV4 field if non-nil, zero value otherwise.

### GetSubnetV4Ok

`func (o *AvailableNetwork) GetSubnetV4Ok() (*string, bool)`

GetSubnetV4Ok returns a tuple with the SubnetV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetV4

`func (o *AvailableNetwork) SetSubnetV4(v string)`

SetSubnetV4 sets SubnetV4 field to given value.


### GetLocation

`func (o *AvailableNetwork) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AvailableNetwork) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AvailableNetwork) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetCreatedAt

`func (o *AvailableNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AvailableNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AvailableNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetAvailabilityZone

`func (o *AvailableNetwork) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AvailableNetwork) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AvailableNetwork) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetPublicIp

`func (o *AvailableNetwork) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *AvailableNetwork) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *AvailableNetwork) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.


### SetPublicIpNil

`func (o *AvailableNetwork) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *AvailableNetwork) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetType

`func (o *AvailableNetwork) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AvailableNetwork) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AvailableNetwork) SetType(v string)`

SetType sets Type field to given value.


### GetBusyAddress

`func (o *AvailableNetwork) GetBusyAddress() []string`

GetBusyAddress returns the BusyAddress field if non-nil, zero value otherwise.

### GetBusyAddressOk

`func (o *AvailableNetwork) GetBusyAddressOk() (*[]string, bool)`

GetBusyAddressOk returns a tuple with the BusyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyAddress

`func (o *AvailableNetwork) SetBusyAddress(v []string)`

SetBusyAddress sets BusyAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


