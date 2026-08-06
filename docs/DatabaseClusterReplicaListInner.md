# DatabaseClusterReplicaListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID реплики. | 
**DbId** | **int32** | ID кластера базы данных, которому принадлежит реплика. | 
**Status** | **string** | Текущий статус реплики. | 
**LocalIp** | **string** | Адрес реплики в локальной сети. | 
**Disk** | Pointer to [**NullableDatabaseClusterReplicaListInnerDisk**](DatabaseClusterReplicaListInnerDisk.md) |  | [optional] 

## Methods

### NewDatabaseClusterReplicaListInner

`func NewDatabaseClusterReplicaListInner(id int32, dbId int32, status string, localIp string, ) *DatabaseClusterReplicaListInner`

NewDatabaseClusterReplicaListInner instantiates a new DatabaseClusterReplicaListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterReplicaListInnerWithDefaults

`func NewDatabaseClusterReplicaListInnerWithDefaults() *DatabaseClusterReplicaListInner`

NewDatabaseClusterReplicaListInnerWithDefaults instantiates a new DatabaseClusterReplicaListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseClusterReplicaListInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseClusterReplicaListInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseClusterReplicaListInner) SetId(v int32)`

SetId sets Id field to given value.


### GetDbId

`func (o *DatabaseClusterReplicaListInner) GetDbId() int32`

GetDbId returns the DbId field if non-nil, zero value otherwise.

### GetDbIdOk

`func (o *DatabaseClusterReplicaListInner) GetDbIdOk() (*int32, bool)`

GetDbIdOk returns a tuple with the DbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbId

`func (o *DatabaseClusterReplicaListInner) SetDbId(v int32)`

SetDbId sets DbId field to given value.


### GetStatus

`func (o *DatabaseClusterReplicaListInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseClusterReplicaListInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseClusterReplicaListInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLocalIp

`func (o *DatabaseClusterReplicaListInner) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *DatabaseClusterReplicaListInner) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *DatabaseClusterReplicaListInner) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.


### GetDisk

`func (o *DatabaseClusterReplicaListInner) GetDisk() DatabaseClusterReplicaListInnerDisk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *DatabaseClusterReplicaListInner) GetDiskOk() (*DatabaseClusterReplicaListInnerDisk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *DatabaseClusterReplicaListInner) SetDisk(v DatabaseClusterReplicaListInnerDisk)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *DatabaseClusterReplicaListInner) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *DatabaseClusterReplicaListInner) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *DatabaseClusterReplicaListInner) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


