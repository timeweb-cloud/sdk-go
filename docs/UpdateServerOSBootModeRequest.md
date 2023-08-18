# UpdateServerOSBootModeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootMode** | **string** | Тип загрузки операционной системы. \\  Параметры: &#x60;default&#x60; – стандартный режим, &#x60;single&#x60; – однопользовательский режим, &#x60;recovery_disk&#x60; – загрузка с диска восстановления. | 

## Methods

### NewUpdateServerOSBootModeRequest

`func NewUpdateServerOSBootModeRequest(bootMode string, ) *UpdateServerOSBootModeRequest`

NewUpdateServerOSBootModeRequest instantiates a new UpdateServerOSBootModeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerOSBootModeRequestWithDefaults

`func NewUpdateServerOSBootModeRequestWithDefaults() *UpdateServerOSBootModeRequest`

NewUpdateServerOSBootModeRequestWithDefaults instantiates a new UpdateServerOSBootModeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootMode

`func (o *UpdateServerOSBootModeRequest) GetBootMode() string`

GetBootMode returns the BootMode field if non-nil, zero value otherwise.

### GetBootModeOk

`func (o *UpdateServerOSBootModeRequest) GetBootModeOk() (*string, bool)`

GetBootModeOk returns a tuple with the BootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootMode

`func (o *UpdateServerOSBootModeRequest) SetBootMode(v string)`

SetBootMode sets BootMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


