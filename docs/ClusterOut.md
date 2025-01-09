# ClusterOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID кластера | 
**Name** | **string** | Название | 
**CreatedAt** | **time.Time** | Дата и время создания кластера в формате ISO8601 | 
**Status** | **string** | Статус | 
**Description** | **string** | Описание | 
**K8sVersion** | **string** | Версия Kubernetes | 
**NetworkDriver** | **string** | Используемый сетевой драйвер | 
**Ingress** | **bool** | Логическое значение, показывающее, включен ли Ingress | 
**PresetId** | **int32** | ID тарифа мастер-ноды | 
**Cpu** | Pointer to **int32** | Общее количество ядер | [optional] [default to 0]
**Ram** | Pointer to **int32** | Общее количество памяти | [optional] [default to 0]
**Disk** | Pointer to **int32** | Общее дисковое пространство | [optional] [default to 0]
**AvailabilityZone** | Pointer to **string** | Зона доступности | [optional] 
**ProjectId** | Pointer to **int32** | ID проекта | [optional] 

## Methods

### NewClusterOut

`func NewClusterOut(id int32, name string, createdAt time.Time, status string, description string, k8sVersion string, networkDriver string, ingress bool, presetId int32, ) *ClusterOut`

NewClusterOut instantiates a new ClusterOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOutWithDefaults

`func NewClusterOutWithDefaults() *ClusterOut`

NewClusterOutWithDefaults instantiates a new ClusterOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterOut) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterOut) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterOut) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ClusterOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterOut) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *ClusterOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStatus

`func (o *ClusterOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *ClusterOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetK8sVersion

`func (o *ClusterOut) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *ClusterOut) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *ClusterOut) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.


### GetNetworkDriver

`func (o *ClusterOut) GetNetworkDriver() string`

GetNetworkDriver returns the NetworkDriver field if non-nil, zero value otherwise.

### GetNetworkDriverOk

`func (o *ClusterOut) GetNetworkDriverOk() (*string, bool)`

GetNetworkDriverOk returns a tuple with the NetworkDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriver

`func (o *ClusterOut) SetNetworkDriver(v string)`

SetNetworkDriver sets NetworkDriver field to given value.


### GetIngress

`func (o *ClusterOut) GetIngress() bool`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *ClusterOut) GetIngressOk() (*bool, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *ClusterOut) SetIngress(v bool)`

SetIngress sets Ingress field to given value.


### GetPresetId

`func (o *ClusterOut) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *ClusterOut) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *ClusterOut) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetCpu

`func (o *ClusterOut) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ClusterOut) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ClusterOut) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ClusterOut) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *ClusterOut) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ClusterOut) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ClusterOut) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *ClusterOut) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetDisk

`func (o *ClusterOut) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ClusterOut) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ClusterOut) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *ClusterOut) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *ClusterOut) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ClusterOut) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ClusterOut) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *ClusterOut) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetProjectId

`func (o *ClusterOut) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ClusterOut) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ClusterOut) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ClusterOut) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


