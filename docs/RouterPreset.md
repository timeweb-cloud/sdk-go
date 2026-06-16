# RouterPreset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID тарифа | 
**NodeCount** | **int32** | Количество нод | 
**Cpu** | **int32** | Количество ядер процессора | 
**CpuFrequency** | **string** | Частота процессора | 
**Ram** | **int32** | Размер ОЗУ в ГБ | 
**Bandwidth** | **int32** | Пропускная способность | 
**Cost** | **int32** | Цена в рублях | 
**Location** | **string** | Локация | 

## Methods

### NewRouterPreset

`func NewRouterPreset(id int32, nodeCount int32, cpu int32, cpuFrequency string, ram int32, bandwidth int32, cost int32, location string, ) *RouterPreset`

NewRouterPreset instantiates a new RouterPreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterPresetWithDefaults

`func NewRouterPresetWithDefaults() *RouterPreset`

NewRouterPresetWithDefaults instantiates a new RouterPreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouterPreset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouterPreset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouterPreset) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeCount

`func (o *RouterPreset) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *RouterPreset) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *RouterPreset) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetCpu

`func (o *RouterPreset) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *RouterPreset) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *RouterPreset) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetCpuFrequency

`func (o *RouterPreset) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *RouterPreset) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *RouterPreset) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.


### GetRam

`func (o *RouterPreset) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *RouterPreset) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *RouterPreset) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetBandwidth

`func (o *RouterPreset) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *RouterPreset) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *RouterPreset) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetCost

`func (o *RouterPreset) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *RouterPreset) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *RouterPreset) SetCost(v int32)`

SetCost sets Cost field to given value.


### GetLocation

`func (o *RouterPreset) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RouterPreset) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RouterPreset) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


