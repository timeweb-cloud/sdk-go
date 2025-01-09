# ClusterIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название кластера | 
**Description** | Pointer to **string** | Описание кластера | [optional] 
**K8sVersion** | **string** | Версия Kubernetes | 
**AvailabilityZone** | Pointer to **string** | Зона доступности | [optional] 
**NetworkDriver** | **string** | Тип используемого сетевого драйвера в кластере | 
**IsIngress** | Pointer to **bool** | Логическое значение, которое показывает, использовать ли Ingress в кластере | [optional] 
**IsK8sDashboard** | Pointer to **bool** | Логическое значение, которое показывает, использовать ли Kubernetes Dashboard в кластере | [optional] 
**PresetId** | **int32** | ID тарифа мастер-ноды | 
**WorkerGroups** | Pointer to [**[]NodeGroupIn**](NodeGroupIn.md) | Группы воркеров в кластере | [optional] 
**NetworkId** | Pointer to **string** | ID приватной сети | [optional] 
**ProjectId** | Pointer to **int32** | ID проекта | [optional] 

## Methods

### NewClusterIn

`func NewClusterIn(name string, k8sVersion string, networkDriver string, presetId int32, ) *ClusterIn`

NewClusterIn instantiates a new ClusterIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInWithDefaults

`func NewClusterInWithDefaults() *ClusterIn`

NewClusterInWithDefaults instantiates a new ClusterIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterIn) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ClusterIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetK8sVersion

`func (o *ClusterIn) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *ClusterIn) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *ClusterIn) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.


### GetAvailabilityZone

`func (o *ClusterIn) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ClusterIn) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ClusterIn) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *ClusterIn) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetNetworkDriver

`func (o *ClusterIn) GetNetworkDriver() string`

GetNetworkDriver returns the NetworkDriver field if non-nil, zero value otherwise.

### GetNetworkDriverOk

`func (o *ClusterIn) GetNetworkDriverOk() (*string, bool)`

GetNetworkDriverOk returns a tuple with the NetworkDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriver

`func (o *ClusterIn) SetNetworkDriver(v string)`

SetNetworkDriver sets NetworkDriver field to given value.


### GetIsIngress

`func (o *ClusterIn) GetIsIngress() bool`

GetIsIngress returns the IsIngress field if non-nil, zero value otherwise.

### GetIsIngressOk

`func (o *ClusterIn) GetIsIngressOk() (*bool, bool)`

GetIsIngressOk returns a tuple with the IsIngress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIngress

`func (o *ClusterIn) SetIsIngress(v bool)`

SetIsIngress sets IsIngress field to given value.

### HasIsIngress

`func (o *ClusterIn) HasIsIngress() bool`

HasIsIngress returns a boolean if a field has been set.

### GetIsK8sDashboard

`func (o *ClusterIn) GetIsK8sDashboard() bool`

GetIsK8sDashboard returns the IsK8sDashboard field if non-nil, zero value otherwise.

### GetIsK8sDashboardOk

`func (o *ClusterIn) GetIsK8sDashboardOk() (*bool, bool)`

GetIsK8sDashboardOk returns a tuple with the IsK8sDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsK8sDashboard

`func (o *ClusterIn) SetIsK8sDashboard(v bool)`

SetIsK8sDashboard sets IsK8sDashboard field to given value.

### HasIsK8sDashboard

`func (o *ClusterIn) HasIsK8sDashboard() bool`

HasIsK8sDashboard returns a boolean if a field has been set.

### GetPresetId

`func (o *ClusterIn) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *ClusterIn) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *ClusterIn) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetWorkerGroups

`func (o *ClusterIn) GetWorkerGroups() []NodeGroupIn`

GetWorkerGroups returns the WorkerGroups field if non-nil, zero value otherwise.

### GetWorkerGroupsOk

`func (o *ClusterIn) GetWorkerGroupsOk() (*[]NodeGroupIn, bool)`

GetWorkerGroupsOk returns a tuple with the WorkerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerGroups

`func (o *ClusterIn) SetWorkerGroups(v []NodeGroupIn)`

SetWorkerGroups sets WorkerGroups field to given value.

### HasWorkerGroups

`func (o *ClusterIn) HasWorkerGroups() bool`

HasWorkerGroups returns a boolean if a field has been set.

### GetNetworkId

`func (o *ClusterIn) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ClusterIn) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ClusterIn) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *ClusterIn) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetProjectId

`func (o *ClusterIn) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ClusterIn) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ClusterIn) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ClusterIn) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


