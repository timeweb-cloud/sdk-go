# CreateAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** | Имя пользователя базы данных | 
**Password** | **string** | Пароль пользователя базы данных | 
**Host** | Pointer to **string** | Хост пользователя | [optional] 
**InstanceId** | Pointer to **float32** | ID инстанса базы данных для приминения привилегий. В данных момент поле доступно только для кластеров MySQL. Если поле не передано, то привилегии будут применены ко всем инстансам | [optional] 
**Privileges** | **[]string** | Список привилегий пользователя базы данных | 
**Description** | Pointer to **string** | Описание пользователя базы данных | [optional] 

## Methods

### NewCreateAdmin

`func NewCreateAdmin(login string, password string, privileges []string, ) *CreateAdmin`

NewCreateAdmin instantiates a new CreateAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdminWithDefaults

`func NewCreateAdminWithDefaults() *CreateAdmin`

NewCreateAdminWithDefaults instantiates a new CreateAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *CreateAdmin) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CreateAdmin) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *CreateAdmin) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *CreateAdmin) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateAdmin) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateAdmin) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetHost

`func (o *CreateAdmin) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateAdmin) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateAdmin) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CreateAdmin) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetInstanceId

`func (o *CreateAdmin) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CreateAdmin) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CreateAdmin) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *CreateAdmin) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetPrivileges

`func (o *CreateAdmin) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *CreateAdmin) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *CreateAdmin) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.


### GetDescription

`func (o *CreateAdmin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAdmin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAdmin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAdmin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


