# UpdateClusterV2DiskAutoscaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включить автоматическое расширение диска. | 
**StepSize** | Pointer to **int32** | Шаг расширения диска (в Мб). Значение должно быть кратно 5120, минимум — 5120, максимум — 102400. | [optional] 
**ThresholdPercent** | Pointer to **int32** | Порог заполнения диска (в процентах), при достижении которого диск расширяется. | [optional] 

## Methods

### NewUpdateClusterV2DiskAutoscaling

`func NewUpdateClusterV2DiskAutoscaling(isEnabled bool, ) *UpdateClusterV2DiskAutoscaling`

NewUpdateClusterV2DiskAutoscaling instantiates a new UpdateClusterV2DiskAutoscaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterV2DiskAutoscalingWithDefaults

`func NewUpdateClusterV2DiskAutoscalingWithDefaults() *UpdateClusterV2DiskAutoscaling`

NewUpdateClusterV2DiskAutoscalingWithDefaults instantiates a new UpdateClusterV2DiskAutoscaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *UpdateClusterV2DiskAutoscaling) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateClusterV2DiskAutoscaling) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateClusterV2DiskAutoscaling) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetStepSize

`func (o *UpdateClusterV2DiskAutoscaling) GetStepSize() int32`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *UpdateClusterV2DiskAutoscaling) GetStepSizeOk() (*int32, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *UpdateClusterV2DiskAutoscaling) SetStepSize(v int32)`

SetStepSize sets StepSize field to given value.

### HasStepSize

`func (o *UpdateClusterV2DiskAutoscaling) HasStepSize() bool`

HasStepSize returns a boolean if a field has been set.

### GetThresholdPercent

`func (o *UpdateClusterV2DiskAutoscaling) GetThresholdPercent() int32`

GetThresholdPercent returns the ThresholdPercent field if non-nil, zero value otherwise.

### GetThresholdPercentOk

`func (o *UpdateClusterV2DiskAutoscaling) GetThresholdPercentOk() (*int32, bool)`

GetThresholdPercentOk returns a tuple with the ThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPercent

`func (o *UpdateClusterV2DiskAutoscaling) SetThresholdPercent(v int32)`

SetThresholdPercent sets ThresholdPercent field to given value.

### HasThresholdPercent

`func (o *UpdateClusterV2DiskAutoscaling) HasThresholdPercent() bool`

HasThresholdPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


