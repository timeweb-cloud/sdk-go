# WorkerPresetOutApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID тарифа | 
**Description** | **string** | Описание тарифа | 
**DescriptionShort** | **string** | Краткое описание тарифа | 
**Price** | **float32** | Стоимость | 
**Cpu** | **int32** | Количество ядер ноды | 
**Ram** | **int32** | Количество памяти ноды | 
**Disk** | **int32** | Количество пространства на ноде | 
**Network** | **int32** | Пропускная способность ноды | 
**Type** | Pointer to **string** | Тип тарифа | [optional] [default to "worker"]

## Methods

### NewWorkerPresetOutApi

`func NewWorkerPresetOutApi(id int32, description string, descriptionShort string, price float32, cpu int32, ram int32, disk int32, network int32, ) *WorkerPresetOutApi`

NewWorkerPresetOutApi instantiates a new WorkerPresetOutApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerPresetOutApiWithDefaults

`func NewWorkerPresetOutApiWithDefaults() *WorkerPresetOutApi`

NewWorkerPresetOutApiWithDefaults instantiates a new WorkerPresetOutApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkerPresetOutApi) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkerPresetOutApi) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkerPresetOutApi) SetId(v int32)`

SetId sets Id field to given value.


### GetDescription

`func (o *WorkerPresetOutApi) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkerPresetOutApi) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkerPresetOutApi) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionShort

`func (o *WorkerPresetOutApi) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *WorkerPresetOutApi) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *WorkerPresetOutApi) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.


### GetPrice

`func (o *WorkerPresetOutApi) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *WorkerPresetOutApi) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *WorkerPresetOutApi) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetCpu

`func (o *WorkerPresetOutApi) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *WorkerPresetOutApi) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *WorkerPresetOutApi) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *WorkerPresetOutApi) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *WorkerPresetOutApi) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *WorkerPresetOutApi) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetDisk

`func (o *WorkerPresetOutApi) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *WorkerPresetOutApi) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *WorkerPresetOutApi) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetNetwork

`func (o *WorkerPresetOutApi) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WorkerPresetOutApi) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WorkerPresetOutApi) SetNetwork(v int32)`

SetNetwork sets Network field to given value.


### GetType

`func (o *WorkerPresetOutApi) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkerPresetOutApi) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkerPresetOutApi) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkerPresetOutApi) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


