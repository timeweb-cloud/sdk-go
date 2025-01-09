# Db

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра базы данных. Автоматически генерируется при создании. | 
**CreatedAt** | **string** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда была создана база данных. | 
**AccountId** | **string** | ID пользователя. | 
**Login** | **string** | Логин для подключения к базе данных. | 
**Location** | Pointer to **string** | Локация сервера. | [optional] 
**Password** | **string** | Пароль для подключения к базе данных. | 
**Name** | **string** | Название базы данных. | 
**Host** | **NullableString** | Хост. | 
**Type** | [**DbType**](DbType.md) |  | 
**HashType** | **NullableString** | Тип хеширования базы данных (mysql5 | mysql | postgres). | 
**Port** | **int32** | Порт | 
**Ip** | **NullableString** | IP-адрес сетевого интерфейса IPv4. | 
**LocalIp** | **NullableString** | IP-адрес локального сетевого интерфейса IPv4. | 
**Status** | **string** | Текущий статус базы данных. | 
**PresetId** | **int32** | ID тарифа. | 
**DiskStats** | [**NullableDbDiskStats**](DbDiskStats.md) |  | 
**ConfigParameters** | [**ConfigParameters**](ConfigParameters.md) |  | 
**IsOnlyLocalIpAccess** | **bool** | Это логическое значение, которое показывает, доступна ли база данных только по локальному IP адресу. | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | 

## Methods

### NewDb

`func NewDb(id float32, createdAt string, accountId string, login string, password string, name string, host NullableString, type_ DbType, hashType NullableString, port int32, ip NullableString, localIp NullableString, status string, presetId int32, diskStats NullableDbDiskStats, configParameters ConfigParameters, isOnlyLocalIpAccess bool, availabilityZone AvailabilityZone, ) *Db`

NewDb instantiates a new Db object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbWithDefaults

`func NewDbWithDefaults() *Db`

NewDbWithDefaults instantiates a new Db object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Db) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Db) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Db) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Db) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Db) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Db) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetAccountId

`func (o *Db) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Db) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Db) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetLogin

`func (o *Db) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *Db) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *Db) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetLocation

`func (o *Db) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Db) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Db) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Db) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPassword

`func (o *Db) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Db) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Db) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *Db) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Db) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Db) SetName(v string)`

SetName sets Name field to given value.


### GetHost

`func (o *Db) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Db) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Db) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *Db) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *Db) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetType

`func (o *Db) GetType() DbType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Db) GetTypeOk() (*DbType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Db) SetType(v DbType)`

SetType sets Type field to given value.


### GetHashType

`func (o *Db) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *Db) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *Db) SetHashType(v string)`

SetHashType sets HashType field to given value.


### SetHashTypeNil

`func (o *Db) SetHashTypeNil(b bool)`

 SetHashTypeNil sets the value for HashType to be an explicit nil

### UnsetHashType
`func (o *Db) UnsetHashType()`

UnsetHashType ensures that no value is present for HashType, not even an explicit nil
### GetPort

`func (o *Db) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Db) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Db) SetPort(v int32)`

SetPort sets Port field to given value.


### GetIp

`func (o *Db) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Db) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Db) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *Db) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *Db) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetLocalIp

`func (o *Db) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *Db) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *Db) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.


### SetLocalIpNil

`func (o *Db) SetLocalIpNil(b bool)`

 SetLocalIpNil sets the value for LocalIp to be an explicit nil

### UnsetLocalIp
`func (o *Db) UnsetLocalIp()`

UnsetLocalIp ensures that no value is present for LocalIp, not even an explicit nil
### GetStatus

`func (o *Db) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Db) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Db) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPresetId

`func (o *Db) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *Db) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *Db) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetDiskStats

`func (o *Db) GetDiskStats() DbDiskStats`

GetDiskStats returns the DiskStats field if non-nil, zero value otherwise.

### GetDiskStatsOk

`func (o *Db) GetDiskStatsOk() (*DbDiskStats, bool)`

GetDiskStatsOk returns a tuple with the DiskStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStats

`func (o *Db) SetDiskStats(v DbDiskStats)`

SetDiskStats sets DiskStats field to given value.


### SetDiskStatsNil

`func (o *Db) SetDiskStatsNil(b bool)`

 SetDiskStatsNil sets the value for DiskStats to be an explicit nil

### UnsetDiskStats
`func (o *Db) UnsetDiskStats()`

UnsetDiskStats ensures that no value is present for DiskStats, not even an explicit nil
### GetConfigParameters

`func (o *Db) GetConfigParameters() ConfigParameters`

GetConfigParameters returns the ConfigParameters field if non-nil, zero value otherwise.

### GetConfigParametersOk

`func (o *Db) GetConfigParametersOk() (*ConfigParameters, bool)`

GetConfigParametersOk returns a tuple with the ConfigParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParameters

`func (o *Db) SetConfigParameters(v ConfigParameters)`

SetConfigParameters sets ConfigParameters field to given value.


### GetIsOnlyLocalIpAccess

`func (o *Db) GetIsOnlyLocalIpAccess() bool`

GetIsOnlyLocalIpAccess returns the IsOnlyLocalIpAccess field if non-nil, zero value otherwise.

### GetIsOnlyLocalIpAccessOk

`func (o *Db) GetIsOnlyLocalIpAccessOk() (*bool, bool)`

GetIsOnlyLocalIpAccessOk returns a tuple with the IsOnlyLocalIpAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnlyLocalIpAccess

`func (o *Db) SetIsOnlyLocalIpAccess(v bool)`

SetIsOnlyLocalIpAccess sets IsOnlyLocalIpAccess field to given value.


### GetAvailabilityZone

`func (o *Db) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *Db) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *Db) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


