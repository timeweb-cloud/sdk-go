# ClusterIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название кластера | 
**Description** | Pointer to **string** | Описание кластера | [optional] 
**Ha** | **bool** | Описание появится позднее | 
**K8sVersion** | **string** | Версия Kubernetes | 
**NetworkDriver** | **string** | Тип используемого сетевого драйвера в кластере | 
**Ingress** | **bool** | Логическое значение, которое показывает, использовать ли Ingress в кластере | 
**PresetId** | **int32** | Идентификатор тарифа мастер-ноды | 
**WorkerGroups** | Pointer to [**[]NodeGroupIn**](NodeGroupIn.md) | Группы воркеров в кластере | [optional] 

## Methods

### NewClusterIn

`func NewClusterIn(name string, ha bool, k8sVersion string, networkDriver string, ingress bool, presetId int32, ) *ClusterIn`

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

### GetHa

`func (o *ClusterIn) GetHa() bool`

GetHa returns the Ha field if non-nil, zero value otherwise.

### GetHaOk

`func (o *ClusterIn) GetHaOk() (*bool, bool)`

GetHaOk returns a tuple with the Ha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHa

`func (o *ClusterIn) SetHa(v bool)`

SetHa sets Ha field to given value.


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


### GetIngress

`func (o *ClusterIn) GetIngress() bool`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *ClusterIn) GetIngressOk() (*bool, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *ClusterIn) SetIngress(v bool)`

SetIngress sets Ingress field to given value.


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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


