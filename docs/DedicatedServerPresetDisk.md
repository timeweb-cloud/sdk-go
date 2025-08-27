# DedicatedServerPresetDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Описание характеристик диска выделенного сервера. | 
**Count** | **float32** | Количество дисков выделенного сервера. | 
**TotalSize** | **float32** | Общий размер дисков выделенного сервера. | 
**Type** | **string** | Тип дисков выделенного сервера. | 

## Methods

### NewDedicatedServerPresetDisk

`func NewDedicatedServerPresetDisk(description string, count float32, totalSize float32, type_ string, ) *DedicatedServerPresetDisk`

NewDedicatedServerPresetDisk instantiates a new DedicatedServerPresetDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerPresetDiskWithDefaults

`func NewDedicatedServerPresetDiskWithDefaults() *DedicatedServerPresetDisk`

NewDedicatedServerPresetDiskWithDefaults instantiates a new DedicatedServerPresetDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DedicatedServerPresetDisk) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DedicatedServerPresetDisk) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DedicatedServerPresetDisk) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCount

`func (o *DedicatedServerPresetDisk) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DedicatedServerPresetDisk) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DedicatedServerPresetDisk) SetCount(v float32)`

SetCount sets Count field to given value.


### GetTotalSize

`func (o *DedicatedServerPresetDisk) GetTotalSize() float32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *DedicatedServerPresetDisk) GetTotalSizeOk() (*float32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *DedicatedServerPresetDisk) SetTotalSize(v float32)`

SetTotalSize sets TotalSize field to given value.


### GetType

`func (o *DedicatedServerPresetDisk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DedicatedServerPresetDisk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DedicatedServerPresetDisk) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


