# DatabaseAdminInstancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **float32** | ID базы данных | 
**Privileges** | **[]string** | Список привилегий пользователя базы данных | 

## Methods

### NewDatabaseAdminInstancesInner

`func NewDatabaseAdminInstancesInner(instanceId float32, privileges []string, ) *DatabaseAdminInstancesInner`

NewDatabaseAdminInstancesInner instantiates a new DatabaseAdminInstancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseAdminInstancesInnerWithDefaults

`func NewDatabaseAdminInstancesInnerWithDefaults() *DatabaseAdminInstancesInner`

NewDatabaseAdminInstancesInnerWithDefaults instantiates a new DatabaseAdminInstancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *DatabaseAdminInstancesInner) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *DatabaseAdminInstancesInner) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *DatabaseAdminInstancesInner) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.


### GetPrivileges

`func (o *DatabaseAdminInstancesInner) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *DatabaseAdminInstancesInner) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *DatabaseAdminInstancesInner) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


