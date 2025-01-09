# DatabaseCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра базы данных. Автоматически генерируется при создании. | 
**CreatedAt** | **string** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда была создана база данных. | 
**Location** | **NullableString** | Локация сервера. | 
**Name** | **string** | Название кластера базы данных. | 
**Networks** | [**[]DatabaseClusterNetworksInner**](DatabaseClusterNetworksInner.md) | Список сетей кластера базы данных. | 
**Type** | [**DbType**](DbType.md) |  | 
**HashType** | **NullableString** | Тип хеширования кластера базы данных (mysql5 | mysql | postgres). | 
**Port** | **NullableInt32** | Порт | 
**Status** | **string** | Текущий статус кластера базы данных. | 
**PresetId** | **int32** | ID тарифа. | 
**DiskStats** | [**NullableDatabaseClusterDiskStats**](DatabaseClusterDiskStats.md) |  | 
**ConfigParameters** | [**ConfigParameters**](ConfigParameters.md) |  | 
**IsEnabledPublicNetwork** | **bool** | Доступность публичного IP-адреса | 

## Methods

### NewDatabaseCluster

`func NewDatabaseCluster(id float32, createdAt string, location NullableString, name string, networks []DatabaseClusterNetworksInner, type_ DbType, hashType NullableString, port NullableInt32, status string, presetId int32, diskStats NullableDatabaseClusterDiskStats, configParameters ConfigParameters, isEnabledPublicNetwork bool, ) *DatabaseCluster`

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


### GetType

`func (o *DatabaseCluster) GetType() DbType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseCluster) GetTypeOk() (*DbType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseCluster) SetType(v DbType)`

SetType sets Type field to given value.


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


### GetDiskStats

`func (o *DatabaseCluster) GetDiskStats() DatabaseClusterDiskStats`

GetDiskStats returns the DiskStats field if non-nil, zero value otherwise.

### GetDiskStatsOk

`func (o *DatabaseCluster) GetDiskStatsOk() (*DatabaseClusterDiskStats, bool)`

GetDiskStatsOk returns a tuple with the DiskStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStats

`func (o *DatabaseCluster) SetDiskStats(v DatabaseClusterDiskStats)`

SetDiskStats sets DiskStats field to given value.


### SetDiskStatsNil

`func (o *DatabaseCluster) SetDiskStatsNil(b bool)`

 SetDiskStatsNil sets the value for DiskStats to be an explicit nil

### UnsetDiskStats
`func (o *DatabaseCluster) UnsetDiskStats()`

UnsetDiskStats ensures that no value is present for DiskStats, not even an explicit nil
### GetConfigParameters

`func (o *DatabaseCluster) GetConfigParameters() ConfigParameters`

GetConfigParameters returns the ConfigParameters field if non-nil, zero value otherwise.

### GetConfigParametersOk

`func (o *DatabaseCluster) GetConfigParametersOk() (*ConfigParameters, bool)`

GetConfigParametersOk returns a tuple with the ConfigParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParameters

`func (o *DatabaseCluster) SetConfigParameters(v ConfigParameters)`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


