# FloatingIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID IP. | 
**Ip** | **NullableString** | IP-адрес | 
**IsDdosGuard** | **bool** | Это логическое значение, которое показывает, включена ли защита от DDoS. | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | 
**ResourceType** | **NullableString** | Тип ресурса. | 
**ResourceId** | **NullableFloat32** | Id ресурса. | 
**Comment** | **NullableString** | Комментарий | 
**Ptr** | **NullableString** | Запись имени узла. | 

## Methods

### NewFloatingIp

`func NewFloatingIp(id string, ip NullableString, isDdosGuard bool, availabilityZone AvailabilityZone, resourceType NullableString, resourceId NullableFloat32, comment NullableString, ptr NullableString, ) *FloatingIp`

NewFloatingIp instantiates a new FloatingIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatingIpWithDefaults

`func NewFloatingIpWithDefaults() *FloatingIp`

NewFloatingIpWithDefaults instantiates a new FloatingIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FloatingIp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FloatingIp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FloatingIp) SetId(v string)`

SetId sets Id field to given value.


### GetIp

`func (o *FloatingIp) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *FloatingIp) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *FloatingIp) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *FloatingIp) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *FloatingIp) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetIsDdosGuard

`func (o *FloatingIp) GetIsDdosGuard() bool`

GetIsDdosGuard returns the IsDdosGuard field if non-nil, zero value otherwise.

### GetIsDdosGuardOk

`func (o *FloatingIp) GetIsDdosGuardOk() (*bool, bool)`

GetIsDdosGuardOk returns a tuple with the IsDdosGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDdosGuard

`func (o *FloatingIp) SetIsDdosGuard(v bool)`

SetIsDdosGuard sets IsDdosGuard field to given value.


### GetAvailabilityZone

`func (o *FloatingIp) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *FloatingIp) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *FloatingIp) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetResourceType

`func (o *FloatingIp) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *FloatingIp) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *FloatingIp) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### SetResourceTypeNil

`func (o *FloatingIp) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *FloatingIp) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetResourceId

`func (o *FloatingIp) GetResourceId() float32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FloatingIp) GetResourceIdOk() (*float32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FloatingIp) SetResourceId(v float32)`

SetResourceId sets ResourceId field to given value.


### SetResourceIdNil

`func (o *FloatingIp) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *FloatingIp) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetComment

`func (o *FloatingIp) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FloatingIp) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FloatingIp) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *FloatingIp) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FloatingIp) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetPtr

`func (o *FloatingIp) GetPtr() string`

GetPtr returns the Ptr field if non-nil, zero value otherwise.

### GetPtrOk

`func (o *FloatingIp) GetPtrOk() (*string, bool)`

GetPtrOk returns a tuple with the Ptr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtr

`func (o *FloatingIp) SetPtr(v string)`

SetPtr sets Ptr field to given value.


### SetPtrNil

`func (o *FloatingIp) SetPtrNil(b bool)`

 SetPtrNil sets the value for Ptr to be an explicit nil

### UnsetPtr
`func (o *FloatingIp) UnsetPtr()`

UnsetPtr ensures that no value is present for Ptr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


