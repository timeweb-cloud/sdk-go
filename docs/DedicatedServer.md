# DedicatedServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра выделенного сервера. Автоматически генерируется при создании. | 
**CpuDescription** | **string** | Описание параметров процессора выделенного сервера. | 
**HddDescription** | **string** | Описание параметров жёсткого диска выделенного сервера. | 
**RamDescription** | **string** | Описание ОЗУ выделенного сервера. | 
**CreatedAt** | **time.Time** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был создан выделенный сервер. | 
**Ip** | **NullableString** | IP-адрес сетевого интерфейса IPv4. | 
**IpmiIp** | **NullableString** | IP-адрес сетевого интерфейса IPMI. | 
**IpmiLogin** | **NullableString** | Логин, используемый для входа в IPMI-консоль. | 
**IpmiPassword** | **NullableString** | Пароль, используемый для входа в IPMI-консоль. | 
**Ipv6** | **NullableString** | IP-адрес сетевого интерфейса IPv6. | 
**NodeId** | **NullableFloat32** | Внутренний дополнительный ID сервера. | 
**Name** | **string** | Удобочитаемое имя, установленное для выделенного сервера. | 
**Comment** | **string** | Комментарий к выделенному серверу. | 
**VncPass** | **NullableString** | Пароль для подключения к VNC-консоли выделенного сервера. | 
**Status** | **string** | Строка состояния, указывающая состояние выделенного сервера. Может быть «installing», «installed», «on» или «off». | 
**OsId** | **NullableFloat32** | ID операционной системы, установленной на выделенный сервер. | 
**CpId** | **NullableFloat32** | ID панели управления, установленной на выделенный сервер. | 
**BandwidthId** | **NullableFloat32** | ID интернет-канала, установленного на выделенный сервер. | 
**NetworkDriveId** | **[]float32** | Массив уникальных ID сетевых дисков, подключенных к выделенному серверу. | 
**AdditionalIpAddrId** | **[]float32** | Массив уникальных ID дополнительных IP-адресов, подключенных к выделенному серверу. | 
**PlanId** | **NullableFloat32** | ID списка дополнительных услуг выделенного сервера. | 
**Price** | **float32** | Стоимость выделенного сервера. | 
**Location** | **string** | Локация сервера. | 
**AutoinstallReady** | **float32** | Количество готовых к автоматической выдаче серверов. Если значение равно 0, сервер будет установлен через инженеров. | 
**Password** | **NullableString** | Пароль root сервера или пароль Администратора для серверов Windows. | 

## Methods

### NewDedicatedServer

`func NewDedicatedServer(id float32, cpuDescription string, hddDescription string, ramDescription string, createdAt time.Time, ip NullableString, ipmiIp NullableString, ipmiLogin NullableString, ipmiPassword NullableString, ipv6 NullableString, nodeId NullableFloat32, name string, comment string, vncPass NullableString, status string, osId NullableFloat32, cpId NullableFloat32, bandwidthId NullableFloat32, networkDriveId []float32, additionalIpAddrId []float32, planId NullableFloat32, price float32, location string, autoinstallReady float32, password NullableString, ) *DedicatedServer`

NewDedicatedServer instantiates a new DedicatedServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerWithDefaults

`func NewDedicatedServerWithDefaults() *DedicatedServer`

NewDedicatedServerWithDefaults instantiates a new DedicatedServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DedicatedServer) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DedicatedServer) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DedicatedServer) SetId(v float32)`

SetId sets Id field to given value.


### GetCpuDescription

`func (o *DedicatedServer) GetCpuDescription() string`

GetCpuDescription returns the CpuDescription field if non-nil, zero value otherwise.

### GetCpuDescriptionOk

`func (o *DedicatedServer) GetCpuDescriptionOk() (*string, bool)`

GetCpuDescriptionOk returns a tuple with the CpuDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuDescription

`func (o *DedicatedServer) SetCpuDescription(v string)`

SetCpuDescription sets CpuDescription field to given value.


### GetHddDescription

`func (o *DedicatedServer) GetHddDescription() string`

GetHddDescription returns the HddDescription field if non-nil, zero value otherwise.

### GetHddDescriptionOk

`func (o *DedicatedServer) GetHddDescriptionOk() (*string, bool)`

GetHddDescriptionOk returns a tuple with the HddDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddDescription

`func (o *DedicatedServer) SetHddDescription(v string)`

SetHddDescription sets HddDescription field to given value.


### GetRamDescription

`func (o *DedicatedServer) GetRamDescription() string`

GetRamDescription returns the RamDescription field if non-nil, zero value otherwise.

### GetRamDescriptionOk

`func (o *DedicatedServer) GetRamDescriptionOk() (*string, bool)`

GetRamDescriptionOk returns a tuple with the RamDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamDescription

`func (o *DedicatedServer) SetRamDescription(v string)`

SetRamDescription sets RamDescription field to given value.


### GetCreatedAt

`func (o *DedicatedServer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DedicatedServer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DedicatedServer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIp

`func (o *DedicatedServer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DedicatedServer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DedicatedServer) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *DedicatedServer) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *DedicatedServer) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetIpmiIp

`func (o *DedicatedServer) GetIpmiIp() string`

GetIpmiIp returns the IpmiIp field if non-nil, zero value otherwise.

### GetIpmiIpOk

`func (o *DedicatedServer) GetIpmiIpOk() (*string, bool)`

GetIpmiIpOk returns a tuple with the IpmiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiIp

`func (o *DedicatedServer) SetIpmiIp(v string)`

SetIpmiIp sets IpmiIp field to given value.


### SetIpmiIpNil

`func (o *DedicatedServer) SetIpmiIpNil(b bool)`

 SetIpmiIpNil sets the value for IpmiIp to be an explicit nil

### UnsetIpmiIp
`func (o *DedicatedServer) UnsetIpmiIp()`

UnsetIpmiIp ensures that no value is present for IpmiIp, not even an explicit nil
### GetIpmiLogin

`func (o *DedicatedServer) GetIpmiLogin() string`

GetIpmiLogin returns the IpmiLogin field if non-nil, zero value otherwise.

### GetIpmiLoginOk

`func (o *DedicatedServer) GetIpmiLoginOk() (*string, bool)`

GetIpmiLoginOk returns a tuple with the IpmiLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiLogin

`func (o *DedicatedServer) SetIpmiLogin(v string)`

SetIpmiLogin sets IpmiLogin field to given value.


### SetIpmiLoginNil

`func (o *DedicatedServer) SetIpmiLoginNil(b bool)`

 SetIpmiLoginNil sets the value for IpmiLogin to be an explicit nil

### UnsetIpmiLogin
`func (o *DedicatedServer) UnsetIpmiLogin()`

UnsetIpmiLogin ensures that no value is present for IpmiLogin, not even an explicit nil
### GetIpmiPassword

`func (o *DedicatedServer) GetIpmiPassword() string`

GetIpmiPassword returns the IpmiPassword field if non-nil, zero value otherwise.

### GetIpmiPasswordOk

`func (o *DedicatedServer) GetIpmiPasswordOk() (*string, bool)`

GetIpmiPasswordOk returns a tuple with the IpmiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiPassword

`func (o *DedicatedServer) SetIpmiPassword(v string)`

SetIpmiPassword sets IpmiPassword field to given value.


### SetIpmiPasswordNil

`func (o *DedicatedServer) SetIpmiPasswordNil(b bool)`

 SetIpmiPasswordNil sets the value for IpmiPassword to be an explicit nil

### UnsetIpmiPassword
`func (o *DedicatedServer) UnsetIpmiPassword()`

UnsetIpmiPassword ensures that no value is present for IpmiPassword, not even an explicit nil
### GetIpv6

`func (o *DedicatedServer) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *DedicatedServer) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *DedicatedServer) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.


### SetIpv6Nil

`func (o *DedicatedServer) SetIpv6Nil(b bool)`

 SetIpv6Nil sets the value for Ipv6 to be an explicit nil

### UnsetIpv6
`func (o *DedicatedServer) UnsetIpv6()`

UnsetIpv6 ensures that no value is present for Ipv6, not even an explicit nil
### GetNodeId

`func (o *DedicatedServer) GetNodeId() float32`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *DedicatedServer) GetNodeIdOk() (*float32, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *DedicatedServer) SetNodeId(v float32)`

SetNodeId sets NodeId field to given value.


### SetNodeIdNil

`func (o *DedicatedServer) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *DedicatedServer) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetName

`func (o *DedicatedServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DedicatedServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DedicatedServer) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *DedicatedServer) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DedicatedServer) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DedicatedServer) SetComment(v string)`

SetComment sets Comment field to given value.


### GetVncPass

`func (o *DedicatedServer) GetVncPass() string`

GetVncPass returns the VncPass field if non-nil, zero value otherwise.

### GetVncPassOk

`func (o *DedicatedServer) GetVncPassOk() (*string, bool)`

GetVncPassOk returns a tuple with the VncPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncPass

`func (o *DedicatedServer) SetVncPass(v string)`

SetVncPass sets VncPass field to given value.


### SetVncPassNil

`func (o *DedicatedServer) SetVncPassNil(b bool)`

 SetVncPassNil sets the value for VncPass to be an explicit nil

### UnsetVncPass
`func (o *DedicatedServer) UnsetVncPass()`

UnsetVncPass ensures that no value is present for VncPass, not even an explicit nil
### GetStatus

`func (o *DedicatedServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DedicatedServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DedicatedServer) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOsId

`func (o *DedicatedServer) GetOsId() float32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *DedicatedServer) GetOsIdOk() (*float32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *DedicatedServer) SetOsId(v float32)`

SetOsId sets OsId field to given value.


### SetOsIdNil

`func (o *DedicatedServer) SetOsIdNil(b bool)`

 SetOsIdNil sets the value for OsId to be an explicit nil

### UnsetOsId
`func (o *DedicatedServer) UnsetOsId()`

UnsetOsId ensures that no value is present for OsId, not even an explicit nil
### GetCpId

`func (o *DedicatedServer) GetCpId() float32`

GetCpId returns the CpId field if non-nil, zero value otherwise.

### GetCpIdOk

`func (o *DedicatedServer) GetCpIdOk() (*float32, bool)`

GetCpIdOk returns a tuple with the CpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpId

`func (o *DedicatedServer) SetCpId(v float32)`

SetCpId sets CpId field to given value.


### SetCpIdNil

`func (o *DedicatedServer) SetCpIdNil(b bool)`

 SetCpIdNil sets the value for CpId to be an explicit nil

### UnsetCpId
`func (o *DedicatedServer) UnsetCpId()`

UnsetCpId ensures that no value is present for CpId, not even an explicit nil
### GetBandwidthId

`func (o *DedicatedServer) GetBandwidthId() float32`

GetBandwidthId returns the BandwidthId field if non-nil, zero value otherwise.

### GetBandwidthIdOk

`func (o *DedicatedServer) GetBandwidthIdOk() (*float32, bool)`

GetBandwidthIdOk returns a tuple with the BandwidthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthId

`func (o *DedicatedServer) SetBandwidthId(v float32)`

SetBandwidthId sets BandwidthId field to given value.


### SetBandwidthIdNil

`func (o *DedicatedServer) SetBandwidthIdNil(b bool)`

 SetBandwidthIdNil sets the value for BandwidthId to be an explicit nil

### UnsetBandwidthId
`func (o *DedicatedServer) UnsetBandwidthId()`

UnsetBandwidthId ensures that no value is present for BandwidthId, not even an explicit nil
### GetNetworkDriveId

`func (o *DedicatedServer) GetNetworkDriveId() []float32`

GetNetworkDriveId returns the NetworkDriveId field if non-nil, zero value otherwise.

### GetNetworkDriveIdOk

`func (o *DedicatedServer) GetNetworkDriveIdOk() (*[]float32, bool)`

GetNetworkDriveIdOk returns a tuple with the NetworkDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriveId

`func (o *DedicatedServer) SetNetworkDriveId(v []float32)`

SetNetworkDriveId sets NetworkDriveId field to given value.


### SetNetworkDriveIdNil

`func (o *DedicatedServer) SetNetworkDriveIdNil(b bool)`

 SetNetworkDriveIdNil sets the value for NetworkDriveId to be an explicit nil

### UnsetNetworkDriveId
`func (o *DedicatedServer) UnsetNetworkDriveId()`

UnsetNetworkDriveId ensures that no value is present for NetworkDriveId, not even an explicit nil
### GetAdditionalIpAddrId

`func (o *DedicatedServer) GetAdditionalIpAddrId() []float32`

GetAdditionalIpAddrId returns the AdditionalIpAddrId field if non-nil, zero value otherwise.

### GetAdditionalIpAddrIdOk

`func (o *DedicatedServer) GetAdditionalIpAddrIdOk() (*[]float32, bool)`

GetAdditionalIpAddrIdOk returns a tuple with the AdditionalIpAddrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIpAddrId

`func (o *DedicatedServer) SetAdditionalIpAddrId(v []float32)`

SetAdditionalIpAddrId sets AdditionalIpAddrId field to given value.


### SetAdditionalIpAddrIdNil

`func (o *DedicatedServer) SetAdditionalIpAddrIdNil(b bool)`

 SetAdditionalIpAddrIdNil sets the value for AdditionalIpAddrId to be an explicit nil

### UnsetAdditionalIpAddrId
`func (o *DedicatedServer) UnsetAdditionalIpAddrId()`

UnsetAdditionalIpAddrId ensures that no value is present for AdditionalIpAddrId, not even an explicit nil
### GetPlanId

`func (o *DedicatedServer) GetPlanId() float32`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DedicatedServer) GetPlanIdOk() (*float32, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DedicatedServer) SetPlanId(v float32)`

SetPlanId sets PlanId field to given value.


### SetPlanIdNil

`func (o *DedicatedServer) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *DedicatedServer) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetPrice

`func (o *DedicatedServer) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DedicatedServer) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DedicatedServer) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetLocation

`func (o *DedicatedServer) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DedicatedServer) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DedicatedServer) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetAutoinstallReady

`func (o *DedicatedServer) GetAutoinstallReady() float32`

GetAutoinstallReady returns the AutoinstallReady field if non-nil, zero value otherwise.

### GetAutoinstallReadyOk

`func (o *DedicatedServer) GetAutoinstallReadyOk() (*float32, bool)`

GetAutoinstallReadyOk returns a tuple with the AutoinstallReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoinstallReady

`func (o *DedicatedServer) SetAutoinstallReady(v float32)`

SetAutoinstallReady sets AutoinstallReady field to given value.


### GetPassword

`func (o *DedicatedServer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DedicatedServer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DedicatedServer) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *DedicatedServer) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *DedicatedServer) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


