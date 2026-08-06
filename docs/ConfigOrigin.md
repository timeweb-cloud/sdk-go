# ConfigOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]OriginServer**](OriginServer.md) | Origin-серверы, с которых CDN забирает контент. Передача этого поля переключает ресурс с S3-хранилища на origin-сервер. | [optional] 
**UseHttps** | Pointer to **bool** | Обращаться к источнику контента по HTTPS. | [optional] 
**Aws** | Pointer to [**NullableConfigOriginAws**](ConfigOriginAws.md) |  | [optional] 

## Methods

### NewConfigOrigin

`func NewConfigOrigin() *ConfigOrigin`

NewConfigOrigin instantiates a new ConfigOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigOriginWithDefaults

`func NewConfigOriginWithDefaults() *ConfigOrigin`

NewConfigOriginWithDefaults instantiates a new ConfigOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *ConfigOrigin) GetServers() []OriginServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ConfigOrigin) GetServersOk() (*[]OriginServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ConfigOrigin) SetServers(v []OriginServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ConfigOrigin) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *ConfigOrigin) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *ConfigOrigin) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetUseHttps

`func (o *ConfigOrigin) GetUseHttps() bool`

GetUseHttps returns the UseHttps field if non-nil, zero value otherwise.

### GetUseHttpsOk

`func (o *ConfigOrigin) GetUseHttpsOk() (*bool, bool)`

GetUseHttpsOk returns a tuple with the UseHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHttps

`func (o *ConfigOrigin) SetUseHttps(v bool)`

SetUseHttps sets UseHttps field to given value.

### HasUseHttps

`func (o *ConfigOrigin) HasUseHttps() bool`

HasUseHttps returns a boolean if a field has been set.

### GetAws

`func (o *ConfigOrigin) GetAws() ConfigOriginAws`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *ConfigOrigin) GetAwsOk() (*ConfigOriginAws, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *ConfigOrigin) SetAws(v ConfigOriginAws)`

SetAws sets Aws field to given value.

### HasAws

`func (o *ConfigOrigin) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *ConfigOrigin) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *ConfigOrigin) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


