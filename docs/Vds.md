# Vds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра сервера. Автоматически генерируется при создании. | 
**Name** | **string** | Удобочитаемое имя, установленное для сервера. | 
**Comment** | **string** | Комментарий к серверу. | 
**CreatedAt** | **string** | Дата создания сервера в формате ISO8061. | 
**Os** | [**VdsOs**](VdsOs.md) |  | 
**Software** | [**NullableVdsSoftware**](VdsSoftware.md) |  | 
**PresetId** | **NullableFloat32** | ID тарифа сервера. | 
**Location** | **string** | Локация сервера. | 
**ConfiguratorId** | **NullableFloat32** | ID конфигуратора сервера. | 
**BootMode** | **string** | Режим загрузки ОС сервера. | 
**Status** | **string** | Статус сервера. | 
**StartAt** | **NullableTime** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был запущен сервер. | 
**IsDdosGuard** | **bool** | Это логическое значение, которое показывает, включена ли защита от DDoS у данного сервера. | 
**IsMasterSsh** | **bool** | Это логическое значение, которое показывает, доступно ли подключение по SSH для поддержки. | 
**IsDedicatedCpu** | **bool** | Это логическое значение, которое показывает, является ли CPU выделенным. | 
**Gpu** | **float32** | Количество видеокарт сервера. | 
**Cpu** | **float32** | Количество ядер процессора сервера. | 
**CpuFrequency** | **string** | Частота ядер процессора сервера. | 
**Ram** | **float32** | Размер (в Мб) ОЗУ сервера. | 
**Disks** | [**[]VdsDisksInner**](VdsDisksInner.md) | Список дисков сервера. | 
**AvatarId** | **NullableString** | ID аватара сервера. Описание методов работы с аватарами появится позднее. | 
**VncPass** | **string** | Пароль от VNC. | 
**RootPass** | **NullableString** | Пароль root сервера или пароль Администратора для серверов Windows. | 
**Image** | [**NullableVdsImage**](VdsImage.md) |  | 
**Networks** | [**[]VdsNetworksInner**](VdsNetworksInner.md) | Список сетей сервера. | 
**CloudInit** | **NullableString** | Cloud-init скрипт. | 
**IsQemuAgent** | **bool** | Это логическое значение, которое показывает, включен ли QEMU-agent на сервере. | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | 

## Methods

### NewVds

`func NewVds(id float32, name string, comment string, createdAt string, os VdsOs, software NullableVdsSoftware, presetId NullableFloat32, location string, configuratorId NullableFloat32, bootMode string, status string, startAt NullableTime, isDdosGuard bool, isMasterSsh bool, isDedicatedCpu bool, gpu float32, cpu float32, cpuFrequency string, ram float32, disks []VdsDisksInner, avatarId NullableString, vncPass string, rootPass NullableString, image NullableVdsImage, networks []VdsNetworksInner, cloudInit NullableString, isQemuAgent bool, availabilityZone AvailabilityZone, ) *Vds`

NewVds instantiates a new Vds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdsWithDefaults

`func NewVdsWithDefaults() *Vds`

NewVdsWithDefaults instantiates a new Vds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vds) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vds) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vds) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *Vds) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vds) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vds) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *Vds) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Vds) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Vds) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCreatedAt

`func (o *Vds) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Vds) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Vds) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetOs

`func (o *Vds) GetOs() VdsOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Vds) GetOsOk() (*VdsOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Vds) SetOs(v VdsOs)`

SetOs sets Os field to given value.


### GetSoftware

`func (o *Vds) GetSoftware() VdsSoftware`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *Vds) GetSoftwareOk() (*VdsSoftware, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *Vds) SetSoftware(v VdsSoftware)`

SetSoftware sets Software field to given value.


### SetSoftwareNil

`func (o *Vds) SetSoftwareNil(b bool)`

 SetSoftwareNil sets the value for Software to be an explicit nil

### UnsetSoftware
`func (o *Vds) UnsetSoftware()`

UnsetSoftware ensures that no value is present for Software, not even an explicit nil
### GetPresetId

`func (o *Vds) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *Vds) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *Vds) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.


### SetPresetIdNil

`func (o *Vds) SetPresetIdNil(b bool)`

 SetPresetIdNil sets the value for PresetId to be an explicit nil

### UnsetPresetId
`func (o *Vds) UnsetPresetId()`

UnsetPresetId ensures that no value is present for PresetId, not even an explicit nil
### GetLocation

`func (o *Vds) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Vds) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Vds) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetConfiguratorId

`func (o *Vds) GetConfiguratorId() float32`

GetConfiguratorId returns the ConfiguratorId field if non-nil, zero value otherwise.

### GetConfiguratorIdOk

`func (o *Vds) GetConfiguratorIdOk() (*float32, bool)`

GetConfiguratorIdOk returns a tuple with the ConfiguratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguratorId

`func (o *Vds) SetConfiguratorId(v float32)`

SetConfiguratorId sets ConfiguratorId field to given value.


### SetConfiguratorIdNil

`func (o *Vds) SetConfiguratorIdNil(b bool)`

 SetConfiguratorIdNil sets the value for ConfiguratorId to be an explicit nil

### UnsetConfiguratorId
`func (o *Vds) UnsetConfiguratorId()`

UnsetConfiguratorId ensures that no value is present for ConfiguratorId, not even an explicit nil
### GetBootMode

`func (o *Vds) GetBootMode() string`

GetBootMode returns the BootMode field if non-nil, zero value otherwise.

### GetBootModeOk

`func (o *Vds) GetBootModeOk() (*string, bool)`

GetBootModeOk returns a tuple with the BootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootMode

`func (o *Vds) SetBootMode(v string)`

SetBootMode sets BootMode field to given value.


### GetStatus

`func (o *Vds) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Vds) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Vds) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartAt

`func (o *Vds) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *Vds) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *Vds) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.


### SetStartAtNil

`func (o *Vds) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *Vds) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetIsDdosGuard

`func (o *Vds) GetIsDdosGuard() bool`

GetIsDdosGuard returns the IsDdosGuard field if non-nil, zero value otherwise.

### GetIsDdosGuardOk

`func (o *Vds) GetIsDdosGuardOk() (*bool, bool)`

GetIsDdosGuardOk returns a tuple with the IsDdosGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDdosGuard

`func (o *Vds) SetIsDdosGuard(v bool)`

SetIsDdosGuard sets IsDdosGuard field to given value.


### GetIsMasterSsh

`func (o *Vds) GetIsMasterSsh() bool`

GetIsMasterSsh returns the IsMasterSsh field if non-nil, zero value otherwise.

### GetIsMasterSshOk

`func (o *Vds) GetIsMasterSshOk() (*bool, bool)`

GetIsMasterSshOk returns a tuple with the IsMasterSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterSsh

`func (o *Vds) SetIsMasterSsh(v bool)`

SetIsMasterSsh sets IsMasterSsh field to given value.


### GetIsDedicatedCpu

`func (o *Vds) GetIsDedicatedCpu() bool`

GetIsDedicatedCpu returns the IsDedicatedCpu field if non-nil, zero value otherwise.

### GetIsDedicatedCpuOk

`func (o *Vds) GetIsDedicatedCpuOk() (*bool, bool)`

GetIsDedicatedCpuOk returns a tuple with the IsDedicatedCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDedicatedCpu

`func (o *Vds) SetIsDedicatedCpu(v bool)`

SetIsDedicatedCpu sets IsDedicatedCpu field to given value.


### GetGpu

`func (o *Vds) GetGpu() float32`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *Vds) GetGpuOk() (*float32, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *Vds) SetGpu(v float32)`

SetGpu sets Gpu field to given value.


### GetCpu

`func (o *Vds) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Vds) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Vds) SetCpu(v float32)`

SetCpu sets Cpu field to given value.


### GetCpuFrequency

`func (o *Vds) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *Vds) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *Vds) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.


### GetRam

`func (o *Vds) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *Vds) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *Vds) SetRam(v float32)`

SetRam sets Ram field to given value.


### GetDisks

`func (o *Vds) GetDisks() []VdsDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *Vds) GetDisksOk() (*[]VdsDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *Vds) SetDisks(v []VdsDisksInner)`

SetDisks sets Disks field to given value.


### GetAvatarId

`func (o *Vds) GetAvatarId() string`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *Vds) GetAvatarIdOk() (*string, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *Vds) SetAvatarId(v string)`

SetAvatarId sets AvatarId field to given value.


### SetAvatarIdNil

`func (o *Vds) SetAvatarIdNil(b bool)`

 SetAvatarIdNil sets the value for AvatarId to be an explicit nil

### UnsetAvatarId
`func (o *Vds) UnsetAvatarId()`

UnsetAvatarId ensures that no value is present for AvatarId, not even an explicit nil
### GetVncPass

`func (o *Vds) GetVncPass() string`

GetVncPass returns the VncPass field if non-nil, zero value otherwise.

### GetVncPassOk

`func (o *Vds) GetVncPassOk() (*string, bool)`

GetVncPassOk returns a tuple with the VncPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncPass

`func (o *Vds) SetVncPass(v string)`

SetVncPass sets VncPass field to given value.


### GetRootPass

`func (o *Vds) GetRootPass() string`

GetRootPass returns the RootPass field if non-nil, zero value otherwise.

### GetRootPassOk

`func (o *Vds) GetRootPassOk() (*string, bool)`

GetRootPassOk returns a tuple with the RootPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPass

`func (o *Vds) SetRootPass(v string)`

SetRootPass sets RootPass field to given value.


### SetRootPassNil

`func (o *Vds) SetRootPassNil(b bool)`

 SetRootPassNil sets the value for RootPass to be an explicit nil

### UnsetRootPass
`func (o *Vds) UnsetRootPass()`

UnsetRootPass ensures that no value is present for RootPass, not even an explicit nil
### GetImage

`func (o *Vds) GetImage() VdsImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Vds) GetImageOk() (*VdsImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Vds) SetImage(v VdsImage)`

SetImage sets Image field to given value.


### SetImageNil

`func (o *Vds) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *Vds) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetNetworks

`func (o *Vds) GetNetworks() []VdsNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *Vds) GetNetworksOk() (*[]VdsNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *Vds) SetNetworks(v []VdsNetworksInner)`

SetNetworks sets Networks field to given value.


### GetCloudInit

`func (o *Vds) GetCloudInit() string`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *Vds) GetCloudInitOk() (*string, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *Vds) SetCloudInit(v string)`

SetCloudInit sets CloudInit field to given value.


### SetCloudInitNil

`func (o *Vds) SetCloudInitNil(b bool)`

 SetCloudInitNil sets the value for CloudInit to be an explicit nil

### UnsetCloudInit
`func (o *Vds) UnsetCloudInit()`

UnsetCloudInit ensures that no value is present for CloudInit, not even an explicit nil
### GetIsQemuAgent

`func (o *Vds) GetIsQemuAgent() bool`

GetIsQemuAgent returns the IsQemuAgent field if non-nil, zero value otherwise.

### GetIsQemuAgentOk

`func (o *Vds) GetIsQemuAgentOk() (*bool, bool)`

GetIsQemuAgentOk returns a tuple with the IsQemuAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQemuAgent

`func (o *Vds) SetIsQemuAgent(v bool)`

SetIsQemuAgent sets IsQemuAgent field to given value.


### GetAvailabilityZone

`func (o *Vds) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *Vds) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *Vds) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


