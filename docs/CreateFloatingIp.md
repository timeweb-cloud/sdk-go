# CreateFloatingIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDdosGuard** | **bool** | Это логическое значение, которое показывает, включена ли защита от DDoS. | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | 

## Methods

### NewCreateFloatingIp

`func NewCreateFloatingIp(isDdosGuard bool, availabilityZone AvailabilityZone, ) *CreateFloatingIp`

NewCreateFloatingIp instantiates a new CreateFloatingIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFloatingIpWithDefaults

`func NewCreateFloatingIpWithDefaults() *CreateFloatingIp`

NewCreateFloatingIpWithDefaults instantiates a new CreateFloatingIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDdosGuard

`func (o *CreateFloatingIp) GetIsDdosGuard() bool`

GetIsDdosGuard returns the IsDdosGuard field if non-nil, zero value otherwise.

### GetIsDdosGuardOk

`func (o *CreateFloatingIp) GetIsDdosGuardOk() (*bool, bool)`

GetIsDdosGuardOk returns a tuple with the IsDdosGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDdosGuard

`func (o *CreateFloatingIp) SetIsDdosGuard(v bool)`

SetIsDdosGuard sets IsDdosGuard field to given value.


### GetAvailabilityZone

`func (o *CreateFloatingIp) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateFloatingIp) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateFloatingIp) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


