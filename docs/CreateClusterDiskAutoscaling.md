# CreateClusterDiskAutoscaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включить автоматическое расширение диска. | 
**StepSize** | Pointer to **int32** | Шаг расширения диска (в Мб). | [optional] 
**ThresholdPercent** | Pointer to **int32** | Порог заполнения диска (в процентах), при достижении которого диск расширяется. | [optional] 

## Methods

### NewCreateClusterDiskAutoscaling

`func NewCreateClusterDiskAutoscaling(isEnabled bool, ) *CreateClusterDiskAutoscaling`

NewCreateClusterDiskAutoscaling instantiates a new CreateClusterDiskAutoscaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterDiskAutoscalingWithDefaults

`func NewCreateClusterDiskAutoscalingWithDefaults() *CreateClusterDiskAutoscaling`

NewCreateClusterDiskAutoscalingWithDefaults instantiates a new CreateClusterDiskAutoscaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *CreateClusterDiskAutoscaling) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateClusterDiskAutoscaling) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateClusterDiskAutoscaling) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetStepSize

`func (o *CreateClusterDiskAutoscaling) GetStepSize() int32`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *CreateClusterDiskAutoscaling) GetStepSizeOk() (*int32, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *CreateClusterDiskAutoscaling) SetStepSize(v int32)`

SetStepSize sets StepSize field to given value.

### HasStepSize

`func (o *CreateClusterDiskAutoscaling) HasStepSize() bool`

HasStepSize returns a boolean if a field has been set.

### GetThresholdPercent

`func (o *CreateClusterDiskAutoscaling) GetThresholdPercent() int32`

GetThresholdPercent returns the ThresholdPercent field if non-nil, zero value otherwise.

### GetThresholdPercentOk

`func (o *CreateClusterDiskAutoscaling) GetThresholdPercentOk() (*int32, bool)`

GetThresholdPercentOk returns a tuple with the ThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPercent

`func (o *CreateClusterDiskAutoscaling) SetThresholdPercent(v int32)`

SetThresholdPercent sets ThresholdPercent field to given value.

### HasThresholdPercent

`func (o *CreateClusterDiskAutoscaling) HasThresholdPercent() bool`

HasThresholdPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


