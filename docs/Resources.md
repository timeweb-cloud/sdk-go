# Resources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to **int32** | Количество нод | [optional] [default to 0]
**Cores** | Pointer to [**Resource**](Resource.md) |  | [optional] 
**Memory** | Pointer to [**Resource**](Resource.md) |  | [optional] 
**Pods** | Pointer to [**Resource**](Resource.md) |  | [optional] 

## Methods

### NewResources

`func NewResources() *Resources`

NewResources instantiates a new Resources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesWithDefaults

`func NewResourcesWithDefaults() *Resources`

NewResourcesWithDefaults instantiates a new Resources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *Resources) GetNodes() int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Resources) GetNodesOk() (*int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Resources) SetNodes(v int32)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Resources) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetCores

`func (o *Resources) GetCores() Resource`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *Resources) GetCoresOk() (*Resource, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *Resources) SetCores(v Resource)`

SetCores sets Cores field to given value.

### HasCores

`func (o *Resources) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetMemory

`func (o *Resources) GetMemory() Resource`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Resources) GetMemoryOk() (*Resource, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Resources) SetMemory(v Resource)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Resources) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetPods

`func (o *Resources) GetPods() Resource`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *Resources) GetPodsOk() (*Resource, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *Resources) SetPods(v Resource)`

SetPods sets Pods field to given value.

### HasPods

`func (o *Resources) HasPods() bool`

HasPods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


