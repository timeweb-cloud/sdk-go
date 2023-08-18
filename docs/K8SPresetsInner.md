# K8SPresetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Уникальный идентификатор тарифа | 
**Description** | **string** | Описание тарифа | 
**DescriptionShort** | **string** | Краткое описание тарифа | 
**Price** | **float32** | Стоимость | 
**Cpu** | **int32** | Количество ядер ноды | 
**Ram** | **int32** | Количество памяти ноды | 
**Disk** | **int32** | Количество пространства на ноде | 
**Network** | **int32** | Пропускная способность ноды | 
**Type** | Pointer to **string** | Тип тарифа | [optional] [default to "worker"]
**Limit** | Pointer to **int32** | Ограничение по количеству воркер-нод в кластере | [optional] 

## Methods

### NewK8SPresetsInner

`func NewK8SPresetsInner(id int32, description string, descriptionShort string, price float32, cpu int32, ram int32, disk int32, network int32, ) *K8SPresetsInner`

NewK8SPresetsInner instantiates a new K8SPresetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8SPresetsInnerWithDefaults

`func NewK8SPresetsInnerWithDefaults() *K8SPresetsInner`

NewK8SPresetsInnerWithDefaults instantiates a new K8SPresetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *K8SPresetsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *K8SPresetsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *K8SPresetsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetDescription

`func (o *K8SPresetsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *K8SPresetsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *K8SPresetsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionShort

`func (o *K8SPresetsInner) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *K8SPresetsInner) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *K8SPresetsInner) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.


### GetPrice

`func (o *K8SPresetsInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *K8SPresetsInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *K8SPresetsInner) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetCpu

`func (o *K8SPresetsInner) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *K8SPresetsInner) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *K8SPresetsInner) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *K8SPresetsInner) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *K8SPresetsInner) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *K8SPresetsInner) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetDisk

`func (o *K8SPresetsInner) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *K8SPresetsInner) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *K8SPresetsInner) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetNetwork

`func (o *K8SPresetsInner) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *K8SPresetsInner) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *K8SPresetsInner) SetNetwork(v int32)`

SetNetwork sets Network field to given value.


### GetType

`func (o *K8SPresetsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *K8SPresetsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *K8SPresetsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *K8SPresetsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimit

`func (o *K8SPresetsInner) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *K8SPresetsInner) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *K8SPresetsInner) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *K8SPresetsInner) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


