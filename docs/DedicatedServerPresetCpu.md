# DedicatedServerPresetCpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Описание характеристик процессора выделенного сервера. | 
**DescriptionShort** | **string** | Краткое описание характеристик процессора выделенного сервера. | 
**Count** | **float32** | Количество ядер процессора выделенного сервера. | 

## Methods

### NewDedicatedServerPresetCpu

`func NewDedicatedServerPresetCpu(description string, descriptionShort string, count float32, ) *DedicatedServerPresetCpu`

NewDedicatedServerPresetCpu instantiates a new DedicatedServerPresetCpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerPresetCpuWithDefaults

`func NewDedicatedServerPresetCpuWithDefaults() *DedicatedServerPresetCpu`

NewDedicatedServerPresetCpuWithDefaults instantiates a new DedicatedServerPresetCpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DedicatedServerPresetCpu) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DedicatedServerPresetCpu) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DedicatedServerPresetCpu) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionShort

`func (o *DedicatedServerPresetCpu) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *DedicatedServerPresetCpu) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *DedicatedServerPresetCpu) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.


### GetCount

`func (o *DedicatedServerPresetCpu) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DedicatedServerPresetCpu) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DedicatedServerPresetCpu) SetCount(v float32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


