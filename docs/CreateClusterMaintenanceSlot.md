# CreateClusterMaintenanceSlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Режим обслуживания. | 
**From** | Pointer to **string** | Начало периода обслуживания в формате &#x60;HH:mmZ&#x60; (UTC). Учитывается только при &#x60;type: fixed_time&#x60;. | [optional] 
**To** | Pointer to **string** | Конец периода обслуживания в формате &#x60;HH:mmZ&#x60; (UTC). Учитывается только при &#x60;type: fixed_time&#x60;. | [optional] 

## Methods

### NewCreateClusterMaintenanceSlot

`func NewCreateClusterMaintenanceSlot(type_ string, ) *CreateClusterMaintenanceSlot`

NewCreateClusterMaintenanceSlot instantiates a new CreateClusterMaintenanceSlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterMaintenanceSlotWithDefaults

`func NewCreateClusterMaintenanceSlotWithDefaults() *CreateClusterMaintenanceSlot`

NewCreateClusterMaintenanceSlotWithDefaults instantiates a new CreateClusterMaintenanceSlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateClusterMaintenanceSlot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateClusterMaintenanceSlot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateClusterMaintenanceSlot) SetType(v string)`

SetType sets Type field to given value.


### GetFrom

`func (o *CreateClusterMaintenanceSlot) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateClusterMaintenanceSlot) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateClusterMaintenanceSlot) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CreateClusterMaintenanceSlot) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CreateClusterMaintenanceSlot) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateClusterMaintenanceSlot) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateClusterMaintenanceSlot) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CreateClusterMaintenanceSlot) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


