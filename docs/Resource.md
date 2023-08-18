# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requested** | Pointer to **float32** | Запрошенное количество ресурса | [optional] [default to 0]
**Allocatable** | Pointer to **float32** | Доступное количество | [optional] [default to 0]
**Capacity** | Pointer to **float32** | Общее количество | [optional] [default to 0]
**Used** | Pointer to **float32** | Используемое количество | [optional] [default to 0]

## Methods

### NewResource

`func NewResource() *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequested

`func (o *Resource) GetRequested() float32`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *Resource) GetRequestedOk() (*float32, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *Resource) SetRequested(v float32)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *Resource) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetAllocatable

`func (o *Resource) GetAllocatable() float32`

GetAllocatable returns the Allocatable field if non-nil, zero value otherwise.

### GetAllocatableOk

`func (o *Resource) GetAllocatableOk() (*float32, bool)`

GetAllocatableOk returns a tuple with the Allocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatable

`func (o *Resource) SetAllocatable(v float32)`

SetAllocatable sets Allocatable field to given value.

### HasAllocatable

`func (o *Resource) HasAllocatable() bool`

HasAllocatable returns a boolean if a field has been set.

### GetCapacity

`func (o *Resource) GetCapacity() float32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *Resource) GetCapacityOk() (*float32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *Resource) SetCapacity(v float32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *Resource) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetUsed

`func (o *Resource) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Resource) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Resource) SetUsed(v float32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Resource) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


