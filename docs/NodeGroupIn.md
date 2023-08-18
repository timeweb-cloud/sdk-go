# NodeGroupIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название группы | 
**PresetId** | **int32** | Идентификатор тарифа воркер-ноды | 
**NodeCount** | **int32** | Количество нод в группе | 

## Methods

### NewNodeGroupIn

`func NewNodeGroupIn(name string, presetId int32, nodeCount int32, ) *NodeGroupIn`

NewNodeGroupIn instantiates a new NodeGroupIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeGroupInWithDefaults

`func NewNodeGroupInWithDefaults() *NodeGroupIn`

NewNodeGroupInWithDefaults instantiates a new NodeGroupIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeGroupIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeGroupIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeGroupIn) SetName(v string)`

SetName sets Name field to given value.


### GetPresetId

`func (o *NodeGroupIn) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *NodeGroupIn) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *NodeGroupIn) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetNodeCount

`func (o *NodeGroupIn) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *NodeGroupIn) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *NodeGroupIn) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


