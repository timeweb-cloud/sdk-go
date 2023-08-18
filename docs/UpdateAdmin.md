# UpdateAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Пароль пользователя базы данных | [optional] 
**Privileges** | Pointer to **[]string** | Список привилегий пользователя базы данных | [optional] 
**Description** | Pointer to **string** | Описание пользователя базы данных | [optional] 

## Methods

### NewUpdateAdmin

`func NewUpdateAdmin() *UpdateAdmin`

NewUpdateAdmin instantiates a new UpdateAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdminWithDefaults

`func NewUpdateAdminWithDefaults() *UpdateAdmin`

NewUpdateAdminWithDefaults instantiates a new UpdateAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UpdateAdmin) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateAdmin) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateAdmin) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateAdmin) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPrivileges

`func (o *UpdateAdmin) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *UpdateAdmin) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *UpdateAdmin) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *UpdateAdmin) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAdmin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAdmin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAdmin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAdmin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


