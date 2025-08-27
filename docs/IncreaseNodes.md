# IncreaseNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Количество нод | 
**Labels** | Pointer to [**[]SetLabels**](SetLabels.md) | Лейблы добавляемых нод | [optional] 

## Methods

### NewIncreaseNodes

`func NewIncreaseNodes(count int32, ) *IncreaseNodes`

NewIncreaseNodes instantiates a new IncreaseNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncreaseNodesWithDefaults

`func NewIncreaseNodesWithDefaults() *IncreaseNodes`

NewIncreaseNodesWithDefaults instantiates a new IncreaseNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IncreaseNodes) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IncreaseNodes) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IncreaseNodes) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLabels

`func (o *IncreaseNodes) GetLabels() []SetLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IncreaseNodes) GetLabelsOk() (*[]SetLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IncreaseNodes) SetLabels(v []SetLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IncreaseNodes) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


