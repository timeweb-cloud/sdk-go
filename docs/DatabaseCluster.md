# DatabaseCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра базы данных. Автоматически генерируется при создании. | 
**CreatedAt** | **string** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда была создана база данных. | 
**Location** | **NullableString** | Локация сервера. | 
**Name** | **string** | Название кластера базы данных. | 
**Description** | **string** | Описание кластера базы данных. | 
**Networks** | [**[]DatabaseClusterNetworksInner**](DatabaseClusterNetworksInner.md) | Список сетей кластера базы данных. | 
**IsEnabledPublicIpv6** | **bool** | Использование публичного IPv6-адреса. | 
**Type** | **NullableString** | Тип базы данных. Список возможных значений шире, чем список типов, доступных при создании нового кластера. | 
**HashType** | **NullableString** | Тип хеширования кластера базы данных (mysql5 | mysql | postgres). | 
**AvatarLink** | **NullableString** | Ссылка на аватар для базы данных. | 
**Port** | **NullableInt32** | Порт | 
**Status** | **string** | Текущий статус кластера базы данных. Значение &#x60;read_only&#x60; означает, что запись в кластер заблокирована из-за переполнения диска — чтобы снять блокировку, освободите место или увеличьте размер диска. | 
**PresetId** | **NullableInt32** | ID тарифа. Равен &#x60;null&#x60; у кластеров, созданных через конфигуратор — в этом случае заполнен &#x60;configurator_id&#x60;. | 
**ConfiguratorId** | **NullableInt32** | ID конфигуратора. Равен &#x60;null&#x60; у кластеров, созданных по тарифу. | 
**Cpu** | **NullableInt32** | Количество ядер процессора. | 
**CpuFrequency** | **NullableString** | Частота процессора. | 
**IsDedicatedCpu** | **bool** | Используются ли выделенные ядра процессора. | 
**Ram** | **NullableInt32** | Объем оперативной памяти (в Мб). | 
**Disk** | [**NullableDatabaseClusterDisk**](DatabaseClusterDisk.md) |  | 
**HasAdditionalDisk** | **bool** | Подключен ли к кластеру дополнительный диск. | 
**DiskAutoscaling** | [**NullableDatabaseClusterDiskAutoscaling**](DatabaseClusterDiskAutoscaling.md) |  | 
**ConfigParameters** | [**Mysql**](Mysql.md) |  | 
**IsEnabledPublicNetwork** | **bool** | Доступность публичного IP-адреса | 
**IsSecureConnectionEnabled** | **bool** | Включено ли защищенное подключение к кластеру базы данных. | 
**IsAutobackupsEnabled** | **bool** | Включены ли автоматические резервные копии кластера базы данных. | 
**IsBackupScheduleEnabled** | **bool** | Включено ли расписание резервного копирования кластера базы данных. | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | 
**ProjectId** | Pointer to **int32** | ID проекта, в котором находится кластер базы данных. | [optional] 
**ReplicaList** | [**[]DatabaseClusterReplicaListInner**](DatabaseClusterReplicaListInner.md) | Список реплик кластера базы данных. | 
**Domains** | [**[]DatabaseClusterDomainsInner**](DatabaseClusterDomainsInner.md) | Список доменов кластера базы данных. Если публичная сеть отключена (&#x60;is_enabled_public_network: false&#x60;), список всегда пустой. | 
**ChildServices** | [**[]DatabaseClusterChildServicesInner**](DatabaseClusterChildServicesInner.md) | Список дочерних сервисов кластера базы данных. | 
**ParentServices** | [**[]DatabaseClusterParentServicesInner**](DatabaseClusterParentServicesInner.md) | Список родительских сервисов кластера базы данных. | 
**MaintenanceSlot** | [**DatabaseClusterMaintenanceSlot**](DatabaseClusterMaintenanceSlot.md) |  | 

## Methods

### NewDatabaseCluster

`func NewDatabaseCluster(id float32, createdAt string, location NullableString, name string, description string, networks []DatabaseClusterNetworksInner, isEnabledPublicIpv6 bool, type_ NullableString, hashType NullableString, avatarLink NullableString, port NullableInt32, status string, presetId NullableInt32, configuratorId NullableInt32, cpu NullableInt32, cpuFrequency NullableString, isDedicatedCpu bool, ram NullableInt32, disk NullableDatabaseClusterDisk, hasAdditionalDisk bool, diskAutoscaling NullableDatabaseClusterDiskAutoscaling, configParameters Mysql, isEnabledPublicNetwork bool, isSecureConnectionEnabled bool, isAutobackupsEnabled bool, isBackupScheduleEnabled bool, availabilityZone AvailabilityZone, replicaList []DatabaseClusterReplicaListInner, domains []DatabaseClusterDomainsInner, childServices []DatabaseClusterChildServicesInner, parentServices []DatabaseClusterParentServicesInner, maintenanceSlot DatabaseClusterMaintenanceSlot, ) *DatabaseCluster`

NewDatabaseCluster instantiates a new DatabaseCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterWithDefaults

`func NewDatabaseClusterWithDefaults() *DatabaseCluster`

NewDatabaseClusterWithDefaults instantiates a new DatabaseCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseCluster) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseCluster) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseCluster) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DatabaseCluster) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatabaseCluster) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatabaseCluster) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetLocation

`func (o *DatabaseCluster) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DatabaseCluster) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DatabaseCluster) SetLocation(v string)`

SetLocation sets Location field to given value.


### SetLocationNil

`func (o *DatabaseCluster) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DatabaseCluster) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetName

`func (o *DatabaseCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseCluster) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DatabaseCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseCluster) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNetworks

`func (o *DatabaseCluster) GetNetworks() []DatabaseClusterNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *DatabaseCluster) GetNetworksOk() (*[]DatabaseClusterNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *DatabaseCluster) SetNetworks(v []DatabaseClusterNetworksInner)`

SetNetworks sets Networks field to given value.


### GetIsEnabledPublicIpv6

`func (o *DatabaseCluster) GetIsEnabledPublicIpv6() bool`

GetIsEnabledPublicIpv6 returns the IsEnabledPublicIpv6 field if non-nil, zero value otherwise.

### GetIsEnabledPublicIpv6Ok

`func (o *DatabaseCluster) GetIsEnabledPublicIpv6Ok() (*bool, bool)`

GetIsEnabledPublicIpv6Ok returns a tuple with the IsEnabledPublicIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledPublicIpv6

`func (o *DatabaseCluster) SetIsEnabledPublicIpv6(v bool)`

SetIsEnabledPublicIpv6 sets IsEnabledPublicIpv6 field to given value.


### GetType

`func (o *DatabaseCluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseCluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseCluster) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *DatabaseCluster) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DatabaseCluster) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetHashType

`func (o *DatabaseCluster) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *DatabaseCluster) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *DatabaseCluster) SetHashType(v string)`

SetHashType sets HashType field to given value.


### SetHashTypeNil

`func (o *DatabaseCluster) SetHashTypeNil(b bool)`

 SetHashTypeNil sets the value for HashType to be an explicit nil

### UnsetHashType
`func (o *DatabaseCluster) UnsetHashType()`

UnsetHashType ensures that no value is present for HashType, not even an explicit nil
### GetAvatarLink

`func (o *DatabaseCluster) GetAvatarLink() string`

GetAvatarLink returns the AvatarLink field if non-nil, zero value otherwise.

### GetAvatarLinkOk

`func (o *DatabaseCluster) GetAvatarLinkOk() (*string, bool)`

GetAvatarLinkOk returns a tuple with the AvatarLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarLink

`func (o *DatabaseCluster) SetAvatarLink(v string)`

SetAvatarLink sets AvatarLink field to given value.


### SetAvatarLinkNil

`func (o *DatabaseCluster) SetAvatarLinkNil(b bool)`

 SetAvatarLinkNil sets the value for AvatarLink to be an explicit nil

### UnsetAvatarLink
`func (o *DatabaseCluster) UnsetAvatarLink()`

UnsetAvatarLink ensures that no value is present for AvatarLink, not even an explicit nil
### GetPort

`func (o *DatabaseCluster) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DatabaseCluster) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DatabaseCluster) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *DatabaseCluster) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *DatabaseCluster) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetStatus

`func (o *DatabaseCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseCluster) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPresetId

`func (o *DatabaseCluster) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *DatabaseCluster) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *DatabaseCluster) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### SetPresetIdNil

`func (o *DatabaseCluster) SetPresetIdNil(b bool)`

 SetPresetIdNil sets the value for PresetId to be an explicit nil

### UnsetPresetId
`func (o *DatabaseCluster) UnsetPresetId()`

UnsetPresetId ensures that no value is present for PresetId, not even an explicit nil
### GetConfiguratorId

`func (o *DatabaseCluster) GetConfiguratorId() int32`

GetConfiguratorId returns the ConfiguratorId field if non-nil, zero value otherwise.

### GetConfiguratorIdOk

`func (o *DatabaseCluster) GetConfiguratorIdOk() (*int32, bool)`

GetConfiguratorIdOk returns a tuple with the ConfiguratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguratorId

`func (o *DatabaseCluster) SetConfiguratorId(v int32)`

SetConfiguratorId sets ConfiguratorId field to given value.


### SetConfiguratorIdNil

`func (o *DatabaseCluster) SetConfiguratorIdNil(b bool)`

 SetConfiguratorIdNil sets the value for ConfiguratorId to be an explicit nil

### UnsetConfiguratorId
`func (o *DatabaseCluster) UnsetConfiguratorId()`

UnsetConfiguratorId ensures that no value is present for ConfiguratorId, not even an explicit nil
### GetCpu

`func (o *DatabaseCluster) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DatabaseCluster) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DatabaseCluster) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### SetCpuNil

`func (o *DatabaseCluster) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *DatabaseCluster) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetCpuFrequency

`func (o *DatabaseCluster) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *DatabaseCluster) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *DatabaseCluster) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.


### SetCpuFrequencyNil

`func (o *DatabaseCluster) SetCpuFrequencyNil(b bool)`

 SetCpuFrequencyNil sets the value for CpuFrequency to be an explicit nil

### UnsetCpuFrequency
`func (o *DatabaseCluster) UnsetCpuFrequency()`

UnsetCpuFrequency ensures that no value is present for CpuFrequency, not even an explicit nil
### GetIsDedicatedCpu

`func (o *DatabaseCluster) GetIsDedicatedCpu() bool`

GetIsDedicatedCpu returns the IsDedicatedCpu field if non-nil, zero value otherwise.

### GetIsDedicatedCpuOk

`func (o *DatabaseCluster) GetIsDedicatedCpuOk() (*bool, bool)`

GetIsDedicatedCpuOk returns a tuple with the IsDedicatedCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDedicatedCpu

`func (o *DatabaseCluster) SetIsDedicatedCpu(v bool)`

SetIsDedicatedCpu sets IsDedicatedCpu field to given value.


### GetRam

`func (o *DatabaseCluster) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *DatabaseCluster) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *DatabaseCluster) SetRam(v int32)`

SetRam sets Ram field to given value.


### SetRamNil

`func (o *DatabaseCluster) SetRamNil(b bool)`

 SetRamNil sets the value for Ram to be an explicit nil

### UnsetRam
`func (o *DatabaseCluster) UnsetRam()`

UnsetRam ensures that no value is present for Ram, not even an explicit nil
### GetDisk

`func (o *DatabaseCluster) GetDisk() DatabaseClusterDisk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *DatabaseCluster) GetDiskOk() (*DatabaseClusterDisk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *DatabaseCluster) SetDisk(v DatabaseClusterDisk)`

SetDisk sets Disk field to given value.


### SetDiskNil

`func (o *DatabaseCluster) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *DatabaseCluster) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetHasAdditionalDisk

`func (o *DatabaseCluster) GetHasAdditionalDisk() bool`

GetHasAdditionalDisk returns the HasAdditionalDisk field if non-nil, zero value otherwise.

### GetHasAdditionalDiskOk

`func (o *DatabaseCluster) GetHasAdditionalDiskOk() (*bool, bool)`

GetHasAdditionalDiskOk returns a tuple with the HasAdditionalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAdditionalDisk

`func (o *DatabaseCluster) SetHasAdditionalDisk(v bool)`

SetHasAdditionalDisk sets HasAdditionalDisk field to given value.


### GetDiskAutoscaling

`func (o *DatabaseCluster) GetDiskAutoscaling() DatabaseClusterDiskAutoscaling`

GetDiskAutoscaling returns the DiskAutoscaling field if non-nil, zero value otherwise.

### GetDiskAutoscalingOk

`func (o *DatabaseCluster) GetDiskAutoscalingOk() (*DatabaseClusterDiskAutoscaling, bool)`

GetDiskAutoscalingOk returns a tuple with the DiskAutoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskAutoscaling

`func (o *DatabaseCluster) SetDiskAutoscaling(v DatabaseClusterDiskAutoscaling)`

SetDiskAutoscaling sets DiskAutoscaling field to given value.


### SetDiskAutoscalingNil

`func (o *DatabaseCluster) SetDiskAutoscalingNil(b bool)`

 SetDiskAutoscalingNil sets the value for DiskAutoscaling to be an explicit nil

### UnsetDiskAutoscaling
`func (o *DatabaseCluster) UnsetDiskAutoscaling()`

UnsetDiskAutoscaling ensures that no value is present for DiskAutoscaling, not even an explicit nil
### GetConfigParameters

`func (o *DatabaseCluster) GetConfigParameters() Mysql`

GetConfigParameters returns the ConfigParameters field if non-nil, zero value otherwise.

### GetConfigParametersOk

`func (o *DatabaseCluster) GetConfigParametersOk() (*Mysql, bool)`

GetConfigParametersOk returns a tuple with the ConfigParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParameters

`func (o *DatabaseCluster) SetConfigParameters(v Mysql)`

SetConfigParameters sets ConfigParameters field to given value.


### GetIsEnabledPublicNetwork

`func (o *DatabaseCluster) GetIsEnabledPublicNetwork() bool`

GetIsEnabledPublicNetwork returns the IsEnabledPublicNetwork field if non-nil, zero value otherwise.

### GetIsEnabledPublicNetworkOk

`func (o *DatabaseCluster) GetIsEnabledPublicNetworkOk() (*bool, bool)`

GetIsEnabledPublicNetworkOk returns a tuple with the IsEnabledPublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledPublicNetwork

`func (o *DatabaseCluster) SetIsEnabledPublicNetwork(v bool)`

SetIsEnabledPublicNetwork sets IsEnabledPublicNetwork field to given value.


### GetIsSecureConnectionEnabled

`func (o *DatabaseCluster) GetIsSecureConnectionEnabled() bool`

GetIsSecureConnectionEnabled returns the IsSecureConnectionEnabled field if non-nil, zero value otherwise.

### GetIsSecureConnectionEnabledOk

`func (o *DatabaseCluster) GetIsSecureConnectionEnabledOk() (*bool, bool)`

GetIsSecureConnectionEnabledOk returns a tuple with the IsSecureConnectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecureConnectionEnabled

`func (o *DatabaseCluster) SetIsSecureConnectionEnabled(v bool)`

SetIsSecureConnectionEnabled sets IsSecureConnectionEnabled field to given value.


### GetIsAutobackupsEnabled

`func (o *DatabaseCluster) GetIsAutobackupsEnabled() bool`

GetIsAutobackupsEnabled returns the IsAutobackupsEnabled field if non-nil, zero value otherwise.

### GetIsAutobackupsEnabledOk

`func (o *DatabaseCluster) GetIsAutobackupsEnabledOk() (*bool, bool)`

GetIsAutobackupsEnabledOk returns a tuple with the IsAutobackupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutobackupsEnabled

`func (o *DatabaseCluster) SetIsAutobackupsEnabled(v bool)`

SetIsAutobackupsEnabled sets IsAutobackupsEnabled field to given value.


### GetIsBackupScheduleEnabled

`func (o *DatabaseCluster) GetIsBackupScheduleEnabled() bool`

GetIsBackupScheduleEnabled returns the IsBackupScheduleEnabled field if non-nil, zero value otherwise.

### GetIsBackupScheduleEnabledOk

`func (o *DatabaseCluster) GetIsBackupScheduleEnabledOk() (*bool, bool)`

GetIsBackupScheduleEnabledOk returns a tuple with the IsBackupScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupScheduleEnabled

`func (o *DatabaseCluster) SetIsBackupScheduleEnabled(v bool)`

SetIsBackupScheduleEnabled sets IsBackupScheduleEnabled field to given value.


### GetAvailabilityZone

`func (o *DatabaseCluster) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *DatabaseCluster) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *DatabaseCluster) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetProjectId

`func (o *DatabaseCluster) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DatabaseCluster) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DatabaseCluster) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DatabaseCluster) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReplicaList

`func (o *DatabaseCluster) GetReplicaList() []DatabaseClusterReplicaListInner`

GetReplicaList returns the ReplicaList field if non-nil, zero value otherwise.

### GetReplicaListOk

`func (o *DatabaseCluster) GetReplicaListOk() (*[]DatabaseClusterReplicaListInner, bool)`

GetReplicaListOk returns a tuple with the ReplicaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaList

`func (o *DatabaseCluster) SetReplicaList(v []DatabaseClusterReplicaListInner)`

SetReplicaList sets ReplicaList field to given value.


### GetDomains

`func (o *DatabaseCluster) GetDomains() []DatabaseClusterDomainsInner`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *DatabaseCluster) GetDomainsOk() (*[]DatabaseClusterDomainsInner, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *DatabaseCluster) SetDomains(v []DatabaseClusterDomainsInner)`

SetDomains sets Domains field to given value.


### GetChildServices

`func (o *DatabaseCluster) GetChildServices() []DatabaseClusterChildServicesInner`

GetChildServices returns the ChildServices field if non-nil, zero value otherwise.

### GetChildServicesOk

`func (o *DatabaseCluster) GetChildServicesOk() (*[]DatabaseClusterChildServicesInner, bool)`

GetChildServicesOk returns a tuple with the ChildServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildServices

`func (o *DatabaseCluster) SetChildServices(v []DatabaseClusterChildServicesInner)`

SetChildServices sets ChildServices field to given value.


### GetParentServices

`func (o *DatabaseCluster) GetParentServices() []DatabaseClusterParentServicesInner`

GetParentServices returns the ParentServices field if non-nil, zero value otherwise.

### GetParentServicesOk

`func (o *DatabaseCluster) GetParentServicesOk() (*[]DatabaseClusterParentServicesInner, bool)`

GetParentServicesOk returns a tuple with the ParentServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServices

`func (o *DatabaseCluster) SetParentServices(v []DatabaseClusterParentServicesInner)`

SetParentServices sets ParentServices field to given value.


### GetMaintenanceSlot

`func (o *DatabaseCluster) GetMaintenanceSlot() DatabaseClusterMaintenanceSlot`

GetMaintenanceSlot returns the MaintenanceSlot field if non-nil, zero value otherwise.

### GetMaintenanceSlotOk

`func (o *DatabaseCluster) GetMaintenanceSlotOk() (*DatabaseClusterMaintenanceSlot, bool)`

GetMaintenanceSlotOk returns a tuple with the MaintenanceSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceSlot

`func (o *DatabaseCluster) SetMaintenanceSlot(v DatabaseClusterMaintenanceSlot)`

SetMaintenanceSlot sets MaintenanceSlot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


