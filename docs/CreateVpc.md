# CreateVpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя сети. | 
**SubnetV4** | **string** | Маска подсети. | 
**Location** | **string** | Локация сети. | 
**Description** | Pointer to **string** | Описание. | [optional] 
**AvailabilityZone** | Pointer to [**AvailabilityZone**](AvailabilityZone.md) |  | [optional] 

## Methods

### NewCreateVpc

`func NewCreateVpc(name string, subnetV4 string, location string, ) *CreateVpc`

NewCreateVpc instantiates a new CreateVpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVpcWithDefaults

`func NewCreateVpcWithDefaults() *CreateVpc`

NewCreateVpcWithDefaults instantiates a new CreateVpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVpc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVpc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVpc) SetName(v string)`

SetName sets Name field to given value.


### GetSubnetV4

`func (o *CreateVpc) GetSubnetV4() string`

GetSubnetV4 returns the SubnetV4 field if non-nil, zero value otherwise.

### GetSubnetV4Ok

`func (o *CreateVpc) GetSubnetV4Ok() (*string, bool)`

GetSubnetV4Ok returns a tuple with the SubnetV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetV4

`func (o *CreateVpc) SetSubnetV4(v string)`

SetSubnetV4 sets SubnetV4 field to given value.


### GetLocation

`func (o *CreateVpc) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateVpc) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateVpc) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetDescription

`func (o *CreateVpc) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVpc) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVpc) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVpc) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *CreateVpc) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateVpc) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateVpc) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateVpc) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


