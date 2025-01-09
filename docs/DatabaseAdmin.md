# DatabaseAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра пользователя базы данных. Автоматически генерируется при создании. | 
**CreatedAt** | **string** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда была создана база данных. | 
**Login** | **string** | Имя пользователя базы данных | 
**Password** | **string** | Пароль пользователя базы данных | 
**Description** | **string** | Описанеие пользователя базы данных | 
**Host** | **NullableString** | Хост пользователя | 
**Instances** | [**[]DatabaseAdminInstancesInner**](DatabaseAdminInstancesInner.md) |  | 

## Methods

### NewDatabaseAdmin

`func NewDatabaseAdmin(id float32, createdAt string, login string, password string, description string, host NullableString, instances []DatabaseAdminInstancesInner, ) *DatabaseAdmin`

NewDatabaseAdmin instantiates a new DatabaseAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseAdminWithDefaults

`func NewDatabaseAdminWithDefaults() *DatabaseAdmin`

NewDatabaseAdminWithDefaults instantiates a new DatabaseAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseAdmin) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseAdmin) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseAdmin) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DatabaseAdmin) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatabaseAdmin) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatabaseAdmin) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetLogin

`func (o *DatabaseAdmin) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *DatabaseAdmin) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *DatabaseAdmin) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *DatabaseAdmin) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DatabaseAdmin) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DatabaseAdmin) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDescription

`func (o *DatabaseAdmin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseAdmin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseAdmin) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHost

`func (o *DatabaseAdmin) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DatabaseAdmin) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DatabaseAdmin) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *DatabaseAdmin) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *DatabaseAdmin) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetInstances

`func (o *DatabaseAdmin) GetInstances() []DatabaseAdminInstancesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *DatabaseAdmin) GetInstancesOk() (*[]DatabaseAdminInstancesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *DatabaseAdmin) SetInstances(v []DatabaseAdminInstancesInner)`

SetInstances sets Instances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


