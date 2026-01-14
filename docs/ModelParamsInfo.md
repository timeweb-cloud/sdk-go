# ModelParamsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Temperature** | Pointer to [**ModelParamsInfoTemperature**](ModelParamsInfoTemperature.md) |  | [optional] 
**MaxTokens** | Pointer to [**ModelParamsInfoMaxTokens**](ModelParamsInfoMaxTokens.md) |  | [optional] 
**ReasoningEffort** | Pointer to [**ModelParamsInfoReasoningEffort**](ModelParamsInfoReasoningEffort.md) |  | [optional] 

## Methods

### NewModelParamsInfo

`func NewModelParamsInfo() *ModelParamsInfo`

NewModelParamsInfo instantiates a new ModelParamsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelParamsInfoWithDefaults

`func NewModelParamsInfoWithDefaults() *ModelParamsInfo`

NewModelParamsInfoWithDefaults instantiates a new ModelParamsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemperature

`func (o *ModelParamsInfo) GetTemperature() ModelParamsInfoTemperature`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ModelParamsInfo) GetTemperatureOk() (*ModelParamsInfoTemperature, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ModelParamsInfo) SetTemperature(v ModelParamsInfoTemperature)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ModelParamsInfo) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *ModelParamsInfo) GetMaxTokens() ModelParamsInfoMaxTokens`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *ModelParamsInfo) GetMaxTokensOk() (*ModelParamsInfoMaxTokens, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *ModelParamsInfo) SetMaxTokens(v ModelParamsInfoMaxTokens)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *ModelParamsInfo) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetReasoningEffort

`func (o *ModelParamsInfo) GetReasoningEffort() ModelParamsInfoReasoningEffort`

GetReasoningEffort returns the ReasoningEffort field if non-nil, zero value otherwise.

### GetReasoningEffortOk

`func (o *ModelParamsInfo) GetReasoningEffortOk() (*ModelParamsInfoReasoningEffort, bool)`

GetReasoningEffortOk returns a tuple with the ReasoningEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningEffort

`func (o *ModelParamsInfo) SetReasoningEffort(v ModelParamsInfoReasoningEffort)`

SetReasoningEffort sets ReasoningEffort field to given value.

### HasReasoningEffort

`func (o *ModelParamsInfo) HasReasoningEffort() bool`

HasReasoningEffort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


