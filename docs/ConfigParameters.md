# ConfigParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mysql** | Pointer to [**ConfigParametersMysql**](ConfigParametersMysql.md) |  | [optional] 
**Postgres** | Pointer to [**ConfigParametersPostgres**](ConfigParametersPostgres.md) |  | [optional] 
**Valkey** | Pointer to [**ConfigParametersValkey**](ConfigParametersValkey.md) |  | [optional] 

## Methods

### NewConfigParameters

`func NewConfigParameters() *ConfigParameters`

NewConfigParameters instantiates a new ConfigParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigParametersWithDefaults

`func NewConfigParametersWithDefaults() *ConfigParameters`

NewConfigParametersWithDefaults instantiates a new ConfigParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMysql

`func (o *ConfigParameters) GetMysql() ConfigParametersMysql`

GetMysql returns the Mysql field if non-nil, zero value otherwise.

### GetMysqlOk

`func (o *ConfigParameters) GetMysqlOk() (*ConfigParametersMysql, bool)`

GetMysqlOk returns a tuple with the Mysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql

`func (o *ConfigParameters) SetMysql(v ConfigParametersMysql)`

SetMysql sets Mysql field to given value.

### HasMysql

`func (o *ConfigParameters) HasMysql() bool`

HasMysql returns a boolean if a field has been set.

### GetPostgres

`func (o *ConfigParameters) GetPostgres() ConfigParametersPostgres`

GetPostgres returns the Postgres field if non-nil, zero value otherwise.

### GetPostgresOk

`func (o *ConfigParameters) GetPostgresOk() (*ConfigParametersPostgres, bool)`

GetPostgresOk returns a tuple with the Postgres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgres

`func (o *ConfigParameters) SetPostgres(v ConfigParametersPostgres)`

SetPostgres sets Postgres field to given value.

### HasPostgres

`func (o *ConfigParameters) HasPostgres() bool`

HasPostgres returns a boolean if a field has been set.

### GetValkey

`func (o *ConfigParameters) GetValkey() ConfigParametersValkey`

GetValkey returns the Valkey field if non-nil, zero value otherwise.

### GetValkeyOk

`func (o *ConfigParameters) GetValkeyOk() (*ConfigParametersValkey, bool)`

GetValkeyOk returns a tuple with the Valkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValkey

`func (o *ConfigParameters) SetValkey(v ConfigParametersValkey)`

SetValkey sets Valkey field to given value.

### HasValkey

`func (o *ConfigParameters) HasValkey() bool`

HasValkey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


