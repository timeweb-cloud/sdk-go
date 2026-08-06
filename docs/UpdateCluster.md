# UpdateCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название кластера базы данных. | [optional] 
**PresetId** | Pointer to **int32** | ID тарифа. Нельзя передавать вместе с &#x60;configuration&#x60; | [optional] 
**Configuration** | Pointer to [**UpdateClusterConfiguration**](UpdateClusterConfiguration.md) |  | [optional] 
**ConfigParameters** | Pointer to [**Mysql**](Mysql.md) |  | [optional] 
**HashType** | Pointer to **string** | Тип хеширования базы данных (mysql | postgres). | [optional] 
**Description** | Pointer to **string** | Описание кластера базы данных | [optional] 
**IsEnabledPublicNetwork** | Pointer to **bool** | Доступность публичного IP-адреса | [optional] 
**IsEnabledPublicIpv6** | Pointer to **bool** | Использование публичного IPv6-адреса. | [optional] 
**IsSecureConnectionEnable** | Pointer to **bool** | Включить защищенное подключение к кластеру базы данных | [optional] 
**MaintenanceSlot** | Pointer to [**CreateClusterMaintenanceSlot**](CreateClusterMaintenanceSlot.md) |  | [optional] 
**DiskAutoscaling** | Pointer to [**CreateClusterDiskAutoscaling**](CreateClusterDiskAutoscaling.md) |  | [optional] 

## Methods

### NewUpdateCluster

`func NewUpdateCluster() *UpdateCluster`

NewUpdateCluster instantiates a new UpdateCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterWithDefaults

`func NewUpdateClusterWithDefaults() *UpdateCluster`

NewUpdateClusterWithDefaults instantiates a new UpdateCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPresetId

`func (o *UpdateCluster) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *UpdateCluster) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *UpdateCluster) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *UpdateCluster) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateCluster) GetConfiguration() UpdateClusterConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateCluster) GetConfigurationOk() (*UpdateClusterConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateCluster) SetConfiguration(v UpdateClusterConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UpdateCluster) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetConfigParameters

`func (o *UpdateCluster) GetConfigParameters() Mysql`

GetConfigParameters returns the ConfigParameters field if non-nil, zero value otherwise.

### GetConfigParametersOk

`func (o *UpdateCluster) GetConfigParametersOk() (*Mysql, bool)`

GetConfigParametersOk returns a tuple with the ConfigParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParameters

`func (o *UpdateCluster) SetConfigParameters(v Mysql)`

SetConfigParameters sets ConfigParameters field to given value.

### HasConfigParameters

`func (o *UpdateCluster) HasConfigParameters() bool`

HasConfigParameters returns a boolean if a field has been set.

### GetHashType

`func (o *UpdateCluster) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *UpdateCluster) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *UpdateCluster) SetHashType(v string)`

SetHashType sets HashType field to given value.

### HasHashType

`func (o *UpdateCluster) HasHashType() bool`

HasHashType returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEnabledPublicNetwork

`func (o *UpdateCluster) GetIsEnabledPublicNetwork() bool`

GetIsEnabledPublicNetwork returns the IsEnabledPublicNetwork field if non-nil, zero value otherwise.

### GetIsEnabledPublicNetworkOk

`func (o *UpdateCluster) GetIsEnabledPublicNetworkOk() (*bool, bool)`

GetIsEnabledPublicNetworkOk returns a tuple with the IsEnabledPublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledPublicNetwork

`func (o *UpdateCluster) SetIsEnabledPublicNetwork(v bool)`

SetIsEnabledPublicNetwork sets IsEnabledPublicNetwork field to given value.

### HasIsEnabledPublicNetwork

`func (o *UpdateCluster) HasIsEnabledPublicNetwork() bool`

HasIsEnabledPublicNetwork returns a boolean if a field has been set.

### GetIsEnabledPublicIpv6

`func (o *UpdateCluster) GetIsEnabledPublicIpv6() bool`

GetIsEnabledPublicIpv6 returns the IsEnabledPublicIpv6 field if non-nil, zero value otherwise.

### GetIsEnabledPublicIpv6Ok

`func (o *UpdateCluster) GetIsEnabledPublicIpv6Ok() (*bool, bool)`

GetIsEnabledPublicIpv6Ok returns a tuple with the IsEnabledPublicIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledPublicIpv6

`func (o *UpdateCluster) SetIsEnabledPublicIpv6(v bool)`

SetIsEnabledPublicIpv6 sets IsEnabledPublicIpv6 field to given value.

### HasIsEnabledPublicIpv6

`func (o *UpdateCluster) HasIsEnabledPublicIpv6() bool`

HasIsEnabledPublicIpv6 returns a boolean if a field has been set.

### GetIsSecureConnectionEnable

`func (o *UpdateCluster) GetIsSecureConnectionEnable() bool`

GetIsSecureConnectionEnable returns the IsSecureConnectionEnable field if non-nil, zero value otherwise.

### GetIsSecureConnectionEnableOk

`func (o *UpdateCluster) GetIsSecureConnectionEnableOk() (*bool, bool)`

GetIsSecureConnectionEnableOk returns a tuple with the IsSecureConnectionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecureConnectionEnable

`func (o *UpdateCluster) SetIsSecureConnectionEnable(v bool)`

SetIsSecureConnectionEnable sets IsSecureConnectionEnable field to given value.

### HasIsSecureConnectionEnable

`func (o *UpdateCluster) HasIsSecureConnectionEnable() bool`

HasIsSecureConnectionEnable returns a boolean if a field has been set.

### GetMaintenanceSlot

`func (o *UpdateCluster) GetMaintenanceSlot() CreateClusterMaintenanceSlot`

GetMaintenanceSlot returns the MaintenanceSlot field if non-nil, zero value otherwise.

### GetMaintenanceSlotOk

`func (o *UpdateCluster) GetMaintenanceSlotOk() (*CreateClusterMaintenanceSlot, bool)`

GetMaintenanceSlotOk returns a tuple with the MaintenanceSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceSlot

`func (o *UpdateCluster) SetMaintenanceSlot(v CreateClusterMaintenanceSlot)`

SetMaintenanceSlot sets MaintenanceSlot field to given value.

### HasMaintenanceSlot

`func (o *UpdateCluster) HasMaintenanceSlot() bool`

HasMaintenanceSlot returns a boolean if a field has been set.

### GetDiskAutoscaling

`func (o *UpdateCluster) GetDiskAutoscaling() CreateClusterDiskAutoscaling`

GetDiskAutoscaling returns the DiskAutoscaling field if non-nil, zero value otherwise.

### GetDiskAutoscalingOk

`func (o *UpdateCluster) GetDiskAutoscalingOk() (*CreateClusterDiskAutoscaling, bool)`

GetDiskAutoscalingOk returns a tuple with the DiskAutoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskAutoscaling

`func (o *UpdateCluster) SetDiskAutoscaling(v CreateClusterDiskAutoscaling)`

SetDiskAutoscaling sets DiskAutoscaling field to given value.

### HasDiskAutoscaling

`func (o *UpdateCluster) HasDiskAutoscaling() bool`

HasDiskAutoscaling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


