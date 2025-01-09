# Clusterk8s

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра кластера. Автоматически генерируется при создании. | 
**Name** | **string** | Удобочитаемое имя, установленное для кластера. | 
**CreatedAt** | **time.Time** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был создан кластер. | 
**Status** | **string** | Статус кластера. | 
**Description** | **string** | Описание кластера. | 
**Ha** | **bool** | Описание появится позднее. | 
**K8sVersion** | **string** | Версия k8s. | 
**NetworkDriver** | **string** | Описание появится позднее. | 
**Ingress** | **bool** | Описание появится позднее. | 
**Cpu** | **float32** | Количество ядер процессора кластера. | 
**Ram** | **float32** | Количество (в Мб) оперативной памяти кластера. | 
**Disk** | **float32** | Размер (в Гб) диска кластера. | 
**PresetId** | **float32** | Тип сервиса кластера. | 

## Methods

### NewClusterk8s

`func NewClusterk8s(id float32, name string, createdAt time.Time, status string, description string, ha bool, k8sVersion string, networkDriver string, ingress bool, cpu float32, ram float32, disk float32, presetId float32, ) *Clusterk8s`

NewClusterk8s instantiates a new Clusterk8s object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterk8sWithDefaults

`func NewClusterk8sWithDefaults() *Clusterk8s`

NewClusterk8sWithDefaults instantiates a new Clusterk8s object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Clusterk8s) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Clusterk8s) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Clusterk8s) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *Clusterk8s) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Clusterk8s) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Clusterk8s) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *Clusterk8s) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Clusterk8s) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Clusterk8s) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStatus

`func (o *Clusterk8s) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Clusterk8s) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Clusterk8s) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *Clusterk8s) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Clusterk8s) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Clusterk8s) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHa

`func (o *Clusterk8s) GetHa() bool`

GetHa returns the Ha field if non-nil, zero value otherwise.

### GetHaOk

`func (o *Clusterk8s) GetHaOk() (*bool, bool)`

GetHaOk returns a tuple with the Ha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHa

`func (o *Clusterk8s) SetHa(v bool)`

SetHa sets Ha field to given value.


### GetK8sVersion

`func (o *Clusterk8s) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *Clusterk8s) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *Clusterk8s) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.


### GetNetworkDriver

`func (o *Clusterk8s) GetNetworkDriver() string`

GetNetworkDriver returns the NetworkDriver field if non-nil, zero value otherwise.

### GetNetworkDriverOk

`func (o *Clusterk8s) GetNetworkDriverOk() (*string, bool)`

GetNetworkDriverOk returns a tuple with the NetworkDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriver

`func (o *Clusterk8s) SetNetworkDriver(v string)`

SetNetworkDriver sets NetworkDriver field to given value.


### GetIngress

`func (o *Clusterk8s) GetIngress() bool`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *Clusterk8s) GetIngressOk() (*bool, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *Clusterk8s) SetIngress(v bool)`

SetIngress sets Ingress field to given value.


### GetCpu

`func (o *Clusterk8s) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Clusterk8s) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Clusterk8s) SetCpu(v float32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *Clusterk8s) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *Clusterk8s) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *Clusterk8s) SetRam(v float32)`

SetRam sets Ram field to given value.


### GetDisk

`func (o *Clusterk8s) GetDisk() float32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *Clusterk8s) GetDiskOk() (*float32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *Clusterk8s) SetDisk(v float32)`

SetDisk sets Disk field to given value.


### GetPresetId

`func (o *Clusterk8s) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *Clusterk8s) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *Clusterk8s) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


