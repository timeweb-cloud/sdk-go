# NetworkDriveAvailableResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **float32** | ID сервиса. | 
**ResourceType** | **string** | Тип ресурса. | 
**Ip** | Pointer to **NullableString** | IP-адрес сервиса. | [optional] 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | 

## Methods

### NewNetworkDriveAvailableResource

`func NewNetworkDriveAvailableResource(resourceId float32, resourceType string, availabilityZone AvailabilityZone, ) *NetworkDriveAvailableResource`

NewNetworkDriveAvailableResource instantiates a new NetworkDriveAvailableResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDriveAvailableResourceWithDefaults

`func NewNetworkDriveAvailableResourceWithDefaults() *NetworkDriveAvailableResource`

NewNetworkDriveAvailableResourceWithDefaults instantiates a new NetworkDriveAvailableResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *NetworkDriveAvailableResource) GetResourceId() float32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *NetworkDriveAvailableResource) GetResourceIdOk() (*float32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *NetworkDriveAvailableResource) SetResourceId(v float32)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *NetworkDriveAvailableResource) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *NetworkDriveAvailableResource) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *NetworkDriveAvailableResource) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetIp

`func (o *NetworkDriveAvailableResource) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NetworkDriveAvailableResource) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NetworkDriveAvailableResource) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NetworkDriveAvailableResource) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *NetworkDriveAvailableResource) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *NetworkDriveAvailableResource) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetAvailabilityZone

`func (o *NetworkDriveAvailableResource) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *NetworkDriveAvailableResource) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *NetworkDriveAvailableResource) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


