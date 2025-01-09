# CreateDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** | Логин для подключения к базе данных. | [optional] 
**Password** | **string** | Пароль для подключения к базе данных. | 
**Name** | **string** | Название базы данных. | 
**Type** | [**DbType**](DbType.md) |  | 
**HashType** | Pointer to **string** | Тип хеширования базы данных (mysql5 | mysql | postgres). | [optional] 
**PresetId** | **int32** | ID тарифа. | 
**ConfigParameters** | Pointer to [**ConfigParameters**](ConfigParameters.md) |  | [optional] 
**Network** | Pointer to [**Network**](Network.md) |  | [optional] 

## Methods

### NewCreateDb

`func NewCreateDb(password string, name string, type_ DbType, presetId int32, ) *CreateDb`

NewCreateDb instantiates a new CreateDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDbWithDefaults

`func NewCreateDbWithDefaults() *CreateDb`

NewCreateDbWithDefaults instantiates a new CreateDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *CreateDb) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CreateDb) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *CreateDb) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *CreateDb) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetPassword

`func (o *CreateDb) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDb) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDb) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *CreateDb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDb) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateDb) GetType() DbType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDb) GetTypeOk() (*DbType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDb) SetType(v DbType)`

SetType sets Type field to given value.


### GetHashType

`func (o *CreateDb) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *CreateDb) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *CreateDb) SetHashType(v string)`

SetHashType sets HashType field to given value.

### HasHashType

`func (o *CreateDb) HasHashType() bool`

HasHashType returns a boolean if a field has been set.

### GetPresetId

`func (o *CreateDb) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreateDb) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreateDb) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetConfigParameters

`func (o *CreateDb) GetConfigParameters() ConfigParameters`

GetConfigParameters returns the ConfigParameters field if non-nil, zero value otherwise.

### GetConfigParametersOk

`func (o *CreateDb) GetConfigParametersOk() (*ConfigParameters, bool)`

GetConfigParametersOk returns a tuple with the ConfigParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParameters

`func (o *CreateDb) SetConfigParameters(v ConfigParameters)`

SetConfigParameters sets ConfigParameters field to given value.

### HasConfigParameters

`func (o *CreateDb) HasConfigParameters() bool`

HasConfigParameters returns a boolean if a field has been set.

### GetNetwork

`func (o *CreateDb) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateDb) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateDb) SetNetwork(v Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateDb) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


