# CreateClusterAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** | Имя пользователя базы данных | [optional] 
**Password** | Pointer to **string** | Пароль пользователя базы данных | [optional] 
**Host** | Pointer to **string** | Хост пользователя | [optional] 
**Privileges** | Pointer to **[]string** | Список привилегий пользователя базы данных | [optional] 
**Description** | Pointer to **string** | Описание пользователя базы данных | [optional] 

## Methods

### NewCreateClusterAdmin

`func NewCreateClusterAdmin() *CreateClusterAdmin`

NewCreateClusterAdmin instantiates a new CreateClusterAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterAdminWithDefaults

`func NewCreateClusterAdminWithDefaults() *CreateClusterAdmin`

NewCreateClusterAdminWithDefaults instantiates a new CreateClusterAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *CreateClusterAdmin) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CreateClusterAdmin) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *CreateClusterAdmin) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *CreateClusterAdmin) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetPassword

`func (o *CreateClusterAdmin) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateClusterAdmin) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateClusterAdmin) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateClusterAdmin) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHost

`func (o *CreateClusterAdmin) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateClusterAdmin) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateClusterAdmin) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CreateClusterAdmin) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPrivileges

`func (o *CreateClusterAdmin) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *CreateClusterAdmin) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *CreateClusterAdmin) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *CreateClusterAdmin) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetDescription

`func (o *CreateClusterAdmin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateClusterAdmin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateClusterAdmin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateClusterAdmin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


