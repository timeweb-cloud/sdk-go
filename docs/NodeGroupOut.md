# NodeGroupOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID группы | 
**Name** | **string** | Название группы | 
**CreatedAt** | **time.Time** | Дата и время создания группы в формате ISO8601 | 
**PresetId** | **int32** | ID тарифа мастер-ноды | 
**NodeCount** | **int32** | Количество нод в группе | 

## Methods

### NewNodeGroupOut

`func NewNodeGroupOut(id int32, name string, createdAt time.Time, presetId int32, nodeCount int32, ) *NodeGroupOut`

NewNodeGroupOut instantiates a new NodeGroupOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeGroupOutWithDefaults

`func NewNodeGroupOutWithDefaults() *NodeGroupOut`

NewNodeGroupOutWithDefaults instantiates a new NodeGroupOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeGroupOut) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeGroupOut) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeGroupOut) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *NodeGroupOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeGroupOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeGroupOut) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *NodeGroupOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NodeGroupOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NodeGroupOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPresetId

`func (o *NodeGroupOut) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *NodeGroupOut) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *NodeGroupOut) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetNodeCount

`func (o *NodeGroupOut) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *NodeGroupOut) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *NodeGroupOut) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


