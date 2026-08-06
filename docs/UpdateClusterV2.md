# UpdateClusterV2

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
**FloatingIp** | Pointer to **string** | Плавающий IP-адрес, который нужно привязать к кластеру базы данных. Передается сам адрес, а не его ID; адрес должен быть свободен (не привязан к другому сервису). | [optional] 
**IsSecureConnectionEnable** | Pointer to **bool** | Включить защищенное подключение к кластеру базы данных. Обратите внимание: в ответе это же значение возвращается под ключом &#x60;is_secure_connection_enabled&#x60;. | [optional] 
**MaintenanceSlot** | Pointer to [**UpdateClusterV2MaintenanceSlot**](UpdateClusterV2MaintenanceSlot.md) |  | [optional] 
**DiskAutoscaling** | Pointer to [**UpdateClusterV2DiskAutoscaling**](UpdateClusterV2DiskAutoscaling.md) |  | [optional] 

## Methods

### NewUpdateClusterV2

`func NewUpdateClusterV2() *UpdateClusterV2`

NewUpdateClusterV2 instantiates a new UpdateClusterV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterV2WithDefaults

`func NewUpdateClusterV2WithDefaults() *UpdateClusterV2`

NewUpdateClusterV2WithDefaults instantiates a new UpdateClusterV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateClusterV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClusterV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClusterV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateClusterV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPresetId

`func (o *UpdateClusterV2) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *UpdateClusterV2) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *UpdateClusterV2) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *UpdateClusterV2) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateClusterV2) GetConfiguration() UpdateClusterConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateClusterV2) GetConfigurationOk() (*UpdateClusterConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateClusterV2) SetConfiguration(v UpdateClusterConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UpdateClusterV2) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetConfigParameters

`func (o *UpdateClusterV2) GetConfigParameters() Mysql`

GetConfigParameters returns the ConfigParameters field if non-nil, zero value otherwise.

### GetConfigParametersOk

`func (o *UpdateClusterV2) GetConfigParametersOk() (*Mysql, bool)`

GetConfigParametersOk returns a tuple with the ConfigParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParameters

`func (o *UpdateClusterV2) SetConfigParameters(v Mysql)`

SetConfigParameters sets ConfigParameters field to given value.

### HasConfigParameters

`func (o *UpdateClusterV2) HasConfigParameters() bool`

HasConfigParameters returns a boolean if a field has been set.

### GetHashType

`func (o *UpdateClusterV2) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *UpdateClusterV2) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *UpdateClusterV2) SetHashType(v string)`

SetHashType sets HashType field to given value.

### HasHashType

`func (o *UpdateClusterV2) HasHashType() bool`

HasHashType returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateClusterV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateClusterV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateClusterV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateClusterV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEnabledPublicNetwork

`func (o *UpdateClusterV2) GetIsEnabledPublicNetwork() bool`

GetIsEnabledPublicNetwork returns the IsEnabledPublicNetwork field if non-nil, zero value otherwise.

### GetIsEnabledPublicNetworkOk

`func (o *UpdateClusterV2) GetIsEnabledPublicNetworkOk() (*bool, bool)`

GetIsEnabledPublicNetworkOk returns a tuple with the IsEnabledPublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledPublicNetwork

`func (o *UpdateClusterV2) SetIsEnabledPublicNetwork(v bool)`

SetIsEnabledPublicNetwork sets IsEnabledPublicNetwork field to given value.

### HasIsEnabledPublicNetwork

`func (o *UpdateClusterV2) HasIsEnabledPublicNetwork() bool`

HasIsEnabledPublicNetwork returns a boolean if a field has been set.

### GetIsEnabledPublicIpv6

`func (o *UpdateClusterV2) GetIsEnabledPublicIpv6() bool`

GetIsEnabledPublicIpv6 returns the IsEnabledPublicIpv6 field if non-nil, zero value otherwise.

### GetIsEnabledPublicIpv6Ok

`func (o *UpdateClusterV2) GetIsEnabledPublicIpv6Ok() (*bool, bool)`

GetIsEnabledPublicIpv6Ok returns a tuple with the IsEnabledPublicIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledPublicIpv6

`func (o *UpdateClusterV2) SetIsEnabledPublicIpv6(v bool)`

SetIsEnabledPublicIpv6 sets IsEnabledPublicIpv6 field to given value.

### HasIsEnabledPublicIpv6

`func (o *UpdateClusterV2) HasIsEnabledPublicIpv6() bool`

HasIsEnabledPublicIpv6 returns a boolean if a field has been set.

### GetFloatingIp

`func (o *UpdateClusterV2) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *UpdateClusterV2) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *UpdateClusterV2) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *UpdateClusterV2) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetIsSecureConnectionEnable

`func (o *UpdateClusterV2) GetIsSecureConnectionEnable() bool`

GetIsSecureConnectionEnable returns the IsSecureConnectionEnable field if non-nil, zero value otherwise.

### GetIsSecureConnectionEnableOk

`func (o *UpdateClusterV2) GetIsSecureConnectionEnableOk() (*bool, bool)`

GetIsSecureConnectionEnableOk returns a tuple with the IsSecureConnectionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecureConnectionEnable

`func (o *UpdateClusterV2) SetIsSecureConnectionEnable(v bool)`

SetIsSecureConnectionEnable sets IsSecureConnectionEnable field to given value.

### HasIsSecureConnectionEnable

`func (o *UpdateClusterV2) HasIsSecureConnectionEnable() bool`

HasIsSecureConnectionEnable returns a boolean if a field has been set.

### GetMaintenanceSlot

`func (o *UpdateClusterV2) GetMaintenanceSlot() UpdateClusterV2MaintenanceSlot`

GetMaintenanceSlot returns the MaintenanceSlot field if non-nil, zero value otherwise.

### GetMaintenanceSlotOk

`func (o *UpdateClusterV2) GetMaintenanceSlotOk() (*UpdateClusterV2MaintenanceSlot, bool)`

GetMaintenanceSlotOk returns a tuple with the MaintenanceSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceSlot

`func (o *UpdateClusterV2) SetMaintenanceSlot(v UpdateClusterV2MaintenanceSlot)`

SetMaintenanceSlot sets MaintenanceSlot field to given value.

### HasMaintenanceSlot

`func (o *UpdateClusterV2) HasMaintenanceSlot() bool`

HasMaintenanceSlot returns a boolean if a field has been set.

### GetDiskAutoscaling

`func (o *UpdateClusterV2) GetDiskAutoscaling() UpdateClusterV2DiskAutoscaling`

GetDiskAutoscaling returns the DiskAutoscaling field if non-nil, zero value otherwise.

### GetDiskAutoscalingOk

`func (o *UpdateClusterV2) GetDiskAutoscalingOk() (*UpdateClusterV2DiskAutoscaling, bool)`

GetDiskAutoscalingOk returns a tuple with the DiskAutoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskAutoscaling

`func (o *UpdateClusterV2) SetDiskAutoscaling(v UpdateClusterV2DiskAutoscaling)`

SetDiskAutoscaling sets DiskAutoscaling field to given value.

### HasDiskAutoscaling

`func (o *UpdateClusterV2) HasDiskAutoscaling() bool`

HasDiskAutoscaling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


