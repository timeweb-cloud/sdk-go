# NodeGroupIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название группы | 
**PresetId** | Pointer to **int32** | ID тарифа воркер-ноды. Нельзя передавать вместе с &#x60;configuration&#x60;. Локация воркер-нод должна совпадать с локацией кластера | [optional] 
**Configuration** | Pointer to [**NodeGroupInConfiguration**](NodeGroupInConfiguration.md) |  | [optional] 
**NodeCount** | **int32** | Количество нод в группе | 
**Labels** | Pointer to [**[]SetLabels**](SetLabels.md) | Лейблы для группы нод | [optional] 
**IsAutoscaling** | Pointer to **bool** | Автомасштабирование. Автоматическое увеличение и уменьшение количества нод в группе в зависимости от текущей нагрузки | [optional] 
**MinSize** | Pointer to **int32** | Минимальное количество нод. Передавать в связке с параметрами &#x60;is_autoscaling&#x60; и &#x60;max_size&#x60; | [optional] 
**MaxSize** | Pointer to **int32** | Максимальное количество нод. Передавать в связке с параметрами &#x60;is_autoscaling&#x60; и &#x60;min_size&#x60;. Максимальное количество нод ограничено тарифом кластера | [optional] 

## Methods

### NewNodeGroupIn

`func NewNodeGroupIn(name string, nodeCount int32, ) *NodeGroupIn`

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

### HasPresetId

`func (o *NodeGroupIn) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetConfiguration

`func (o *NodeGroupIn) GetConfiguration() NodeGroupInConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *NodeGroupIn) GetConfigurationOk() (*NodeGroupInConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *NodeGroupIn) SetConfiguration(v NodeGroupInConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *NodeGroupIn) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

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


### GetLabels

`func (o *NodeGroupIn) GetLabels() []SetLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NodeGroupIn) GetLabelsOk() (*[]SetLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NodeGroupIn) SetLabels(v []SetLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NodeGroupIn) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetIsAutoscaling

`func (o *NodeGroupIn) GetIsAutoscaling() bool`

GetIsAutoscaling returns the IsAutoscaling field if non-nil, zero value otherwise.

### GetIsAutoscalingOk

`func (o *NodeGroupIn) GetIsAutoscalingOk() (*bool, bool)`

GetIsAutoscalingOk returns a tuple with the IsAutoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscaling

`func (o *NodeGroupIn) SetIsAutoscaling(v bool)`

SetIsAutoscaling sets IsAutoscaling field to given value.

### HasIsAutoscaling

`func (o *NodeGroupIn) HasIsAutoscaling() bool`

HasIsAutoscaling returns a boolean if a field has been set.

### GetMinSize

`func (o *NodeGroupIn) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *NodeGroupIn) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *NodeGroupIn) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *NodeGroupIn) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### GetMaxSize

`func (o *NodeGroupIn) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *NodeGroupIn) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *NodeGroupIn) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *NodeGroupIn) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


