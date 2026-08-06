# Replica

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID реплики. Автоматически генерируется при создании. | 
**DbId** | **int32** | ID кластера базы данных, которому принадлежит реплика. | 
**Status** | **string** | Текущий статус реплики. | 
**LocalIp** | **string** | IP-адрес реплики в локальной сети. Возвращается пустая строка, если адрес ещё не назначен. | 
**Disk** | Pointer to [**NullableReplicaDisk**](ReplicaDisk.md) |  | [optional] 

## Methods

### NewReplica

`func NewReplica(id int32, dbId int32, status string, localIp string, ) *Replica`

NewReplica instantiates a new Replica object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaWithDefaults

`func NewReplicaWithDefaults() *Replica`

NewReplicaWithDefaults instantiates a new Replica object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Replica) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Replica) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Replica) SetId(v int32)`

SetId sets Id field to given value.


### GetDbId

`func (o *Replica) GetDbId() int32`

GetDbId returns the DbId field if non-nil, zero value otherwise.

### GetDbIdOk

`func (o *Replica) GetDbIdOk() (*int32, bool)`

GetDbIdOk returns a tuple with the DbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbId

`func (o *Replica) SetDbId(v int32)`

SetDbId sets DbId field to given value.


### GetStatus

`func (o *Replica) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Replica) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Replica) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLocalIp

`func (o *Replica) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *Replica) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *Replica) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.


### GetDisk

`func (o *Replica) GetDisk() ReplicaDisk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *Replica) GetDiskOk() (*ReplicaDisk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *Replica) SetDisk(v ReplicaDisk)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *Replica) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *Replica) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *Replica) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


