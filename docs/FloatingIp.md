# FloatingIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор IP. | 
**Ip** | **string** | IP-адрес | 
**IsDdosGuard** | **bool** | Это логическое значение, которое показывает, включена ли защита от DDoS. | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | 
**ResourceType** | Pointer to **string** | Тип ресурса. | [optional] 
**ResourceId** | Pointer to **float32** | Id ресурса. | [optional] 
**Comment** | Pointer to **string** | Комментарий | [optional] 
**Ptr** | Pointer to **string** | Запись имени узла. | [optional] 

## Methods

### NewFloatingIp

`func NewFloatingIp(id string, ip string, isDdosGuard bool, availabilityZone AvailabilityZone, ) *FloatingIp`

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

### HasResourceType

`func (o *FloatingIp) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

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

### HasResourceId

`func (o *FloatingIp) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

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

### HasComment

`func (o *FloatingIp) HasComment() bool`

HasComment returns a boolean if a field has been set.

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

### HasPtr

`func (o *FloatingIp) HasPtr() bool`

HasPtr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


