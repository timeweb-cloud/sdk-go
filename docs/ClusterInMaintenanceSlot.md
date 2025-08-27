# ClusterInMaintenanceSlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | В любое время или в заданное время. При значении &#x60;fixed_time&#x60; поля &#x60;from&#x60; и &#x60;to&#x60; являются обязательными. Минимально допустимый временной интервал — 3 часа. Время задается в часовом поясе UTC. | 
**From** | Pointer to **string** | Интервал времени с. Время должно быть в формате HH:MM (24 часа) | [optional] 
**To** | Pointer to **string** | Интервал времени до. Время должно быть в формате HH:MM (24 часа) | [optional] 

## Methods

### NewClusterInMaintenanceSlot

`func NewClusterInMaintenanceSlot(type_ string, ) *ClusterInMaintenanceSlot`

NewClusterInMaintenanceSlot instantiates a new ClusterInMaintenanceSlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInMaintenanceSlotWithDefaults

`func NewClusterInMaintenanceSlotWithDefaults() *ClusterInMaintenanceSlot`

NewClusterInMaintenanceSlotWithDefaults instantiates a new ClusterInMaintenanceSlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClusterInMaintenanceSlot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterInMaintenanceSlot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterInMaintenanceSlot) SetType(v string)`

SetType sets Type field to given value.


### GetFrom

`func (o *ClusterInMaintenanceSlot) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ClusterInMaintenanceSlot) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ClusterInMaintenanceSlot) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ClusterInMaintenanceSlot) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *ClusterInMaintenanceSlot) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ClusterInMaintenanceSlot) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ClusterInMaintenanceSlot) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ClusterInMaintenanceSlot) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


