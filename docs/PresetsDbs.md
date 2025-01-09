# PresetsDbs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | ID для каждого экземпляра тарифа базы данных. | [optional] 
**Description** | Pointer to **string** | Описание тарифа. | [optional] 
**DescriptionShort** | Pointer to **string** | Краткое описание тарифа. | [optional] 
**Cpu** | Pointer to **float32** | Описание процессора тарифа. | [optional] 
**Ram** | Pointer to **float32** | Описание ОЗУ тарифа. | [optional] 
**Disk** | Pointer to **float32** | Описание диска тарифа. | [optional] 
**Type** | Pointer to [**DbType**](DbType.md) |  | [optional] 
**Price** | Pointer to **float32** | Стоимость тарифа базы данных | [optional] 
**Location** | Pointer to **string** | Географическое расположение тарифа. | [optional] 

## Methods

### NewPresetsDbs

`func NewPresetsDbs() *PresetsDbs`

NewPresetsDbs instantiates a new PresetsDbs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresetsDbsWithDefaults

`func NewPresetsDbsWithDefaults() *PresetsDbs`

NewPresetsDbsWithDefaults instantiates a new PresetsDbs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PresetsDbs) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PresetsDbs) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PresetsDbs) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *PresetsDbs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *PresetsDbs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PresetsDbs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PresetsDbs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PresetsDbs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionShort

`func (o *PresetsDbs) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *PresetsDbs) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *PresetsDbs) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.

### HasDescriptionShort

`func (o *PresetsDbs) HasDescriptionShort() bool`

HasDescriptionShort returns a boolean if a field has been set.

### GetCpu

`func (o *PresetsDbs) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *PresetsDbs) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *PresetsDbs) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *PresetsDbs) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *PresetsDbs) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *PresetsDbs) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *PresetsDbs) SetRam(v float32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *PresetsDbs) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetDisk

`func (o *PresetsDbs) GetDisk() float32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *PresetsDbs) GetDiskOk() (*float32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *PresetsDbs) SetDisk(v float32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *PresetsDbs) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetType

`func (o *PresetsDbs) GetType() DbType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PresetsDbs) GetTypeOk() (*DbType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PresetsDbs) SetType(v DbType)`

SetType sets Type field to given value.

### HasType

`func (o *PresetsDbs) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrice

`func (o *PresetsDbs) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PresetsDbs) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PresetsDbs) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PresetsDbs) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetLocation

`func (o *PresetsDbs) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PresetsDbs) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PresetsDbs) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PresetsDbs) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


