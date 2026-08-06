# DatabaseClusterDiskAutoscaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включено ли автоматическое расширение диска. | 
**StepSize** | **int32** | Шаг расширения диска (в Мб). | 
**ThresholdPercent** | **int32** | Порог заполнения диска (в процентах), при достижении которого диск расширяется. | 

## Methods

### NewDatabaseClusterDiskAutoscaling

`func NewDatabaseClusterDiskAutoscaling(isEnabled bool, stepSize int32, thresholdPercent int32, ) *DatabaseClusterDiskAutoscaling`

NewDatabaseClusterDiskAutoscaling instantiates a new DatabaseClusterDiskAutoscaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterDiskAutoscalingWithDefaults

`func NewDatabaseClusterDiskAutoscalingWithDefaults() *DatabaseClusterDiskAutoscaling`

NewDatabaseClusterDiskAutoscalingWithDefaults instantiates a new DatabaseClusterDiskAutoscaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *DatabaseClusterDiskAutoscaling) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DatabaseClusterDiskAutoscaling) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DatabaseClusterDiskAutoscaling) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetStepSize

`func (o *DatabaseClusterDiskAutoscaling) GetStepSize() int32`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *DatabaseClusterDiskAutoscaling) GetStepSizeOk() (*int32, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *DatabaseClusterDiskAutoscaling) SetStepSize(v int32)`

SetStepSize sets StepSize field to given value.


### GetThresholdPercent

`func (o *DatabaseClusterDiskAutoscaling) GetThresholdPercent() int32`

GetThresholdPercent returns the ThresholdPercent field if non-nil, zero value otherwise.

### GetThresholdPercentOk

`func (o *DatabaseClusterDiskAutoscaling) GetThresholdPercentOk() (*int32, bool)`

GetThresholdPercentOk returns a tuple with the ThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPercent

`func (o *DatabaseClusterDiskAutoscaling) SetThresholdPercent(v int32)`

SetThresholdPercent sets ThresholdPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


