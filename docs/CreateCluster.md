# CreateCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название кластера базы данных. | 
**Type** | [**DbType**](DbType.md) |  | 
**Admin** | Pointer to [**CreateClusterAdmin**](CreateClusterAdmin.md) |  | [optional] 
**Instance** | Pointer to [**CreateClusterInstance**](CreateClusterInstance.md) |  | [optional] 
**HashType** | Pointer to **string** | Тип хеширования базы данных (mysql5 | mysql | postgres). | [optional] 
**PresetId** | **int32** | ID тарифа. | 
**ConfigParameters** | Pointer to [**ConfigParameters**](ConfigParameters.md) |  | [optional] 
**Network** | Pointer to [**Network**](Network.md) |  | [optional] 
**Description** | Pointer to **string** | Описание кластера базы данных | [optional] 
**AvailabilityZone** | Pointer to [**AvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**AutoBackups** | Pointer to [**CreateDbAutoBackups**](CreateDbAutoBackups.md) |  | [optional] 

## Methods

### NewCreateCluster

`func NewCreateCluster(name string, type_ DbType, presetId int32, ) *CreateCluster`

NewCreateCluster instantiates a new CreateCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterWithDefaults

`func NewCreateClusterWithDefaults() *CreateCluster`

NewCreateClusterWithDefaults instantiates a new CreateCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCluster) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateCluster) GetType() DbType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCluster) GetTypeOk() (*DbType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCluster) SetType(v DbType)`

SetType sets Type field to given value.


### GetAdmin

`func (o *CreateCluster) GetAdmin() CreateClusterAdmin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *CreateCluster) GetAdminOk() (*CreateClusterAdmin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *CreateCluster) SetAdmin(v CreateClusterAdmin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *CreateCluster) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetInstance

`func (o *CreateCluster) GetInstance() CreateClusterInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *CreateCluster) GetInstanceOk() (*CreateClusterInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *CreateCluster) SetInstance(v CreateClusterInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *CreateCluster) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetHashType

`func (o *CreateCluster) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *CreateCluster) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *CreateCluster) SetHashType(v string)`

SetHashType sets HashType field to given value.

### HasHashType

`func (o *CreateCluster) HasHashType() bool`

HasHashType returns a boolean if a field has been set.

### GetPresetId

`func (o *CreateCluster) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreateCluster) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreateCluster) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetConfigParameters

`func (o *CreateCluster) GetConfigParameters() ConfigParameters`

GetConfigParameters returns the ConfigParameters field if non-nil, zero value otherwise.

### GetConfigParametersOk

`func (o *CreateCluster) GetConfigParametersOk() (*ConfigParameters, bool)`

GetConfigParametersOk returns a tuple with the ConfigParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParameters

`func (o *CreateCluster) SetConfigParameters(v ConfigParameters)`

SetConfigParameters sets ConfigParameters field to given value.

### HasConfigParameters

`func (o *CreateCluster) HasConfigParameters() bool`

HasConfigParameters returns a boolean if a field has been set.

### GetNetwork

`func (o *CreateCluster) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateCluster) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateCluster) SetNetwork(v Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateCluster) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetDescription

`func (o *CreateCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *CreateCluster) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateCluster) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateCluster) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateCluster) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetAutoBackups

`func (o *CreateCluster) GetAutoBackups() CreateDbAutoBackups`

GetAutoBackups returns the AutoBackups field if non-nil, zero value otherwise.

### GetAutoBackupsOk

`func (o *CreateCluster) GetAutoBackupsOk() (*CreateDbAutoBackups, bool)`

GetAutoBackupsOk returns a tuple with the AutoBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBackups

`func (o *CreateCluster) SetAutoBackups(v CreateDbAutoBackups)`

SetAutoBackups sets AutoBackups field to given value.

### HasAutoBackups

`func (o *CreateCluster) HasAutoBackups() bool`

HasAutoBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


