# AddonOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID дополнения | 
**Type** | **string** | Тип дополнения | 
**Status** | **string** | Статус дополнения | 
**CreatedAt** | **time.Time** | Дата и время создания дополнения в формате ISO8601 | 
**Version** | **string** | Версия дополнения | 
**Config** | Pointer to **map[string]interface{}** | Дополнительная конфигурация дополнения | [optional] 
**YamlConfig** | **string** | Yaml конфигурация дополнения | 
**ConfigType** | **string** | Тип конфигурации дополнения | 

## Methods

### NewAddonOut

`func NewAddonOut(id int32, type_ string, status string, createdAt time.Time, version string, yamlConfig string, configType string, ) *AddonOut`

NewAddonOut instantiates a new AddonOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonOutWithDefaults

`func NewAddonOutWithDefaults() *AddonOut`

NewAddonOutWithDefaults instantiates a new AddonOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddonOut) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonOut) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonOut) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *AddonOut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddonOut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddonOut) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *AddonOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddonOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddonOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *AddonOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AddonOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AddonOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetVersion

`func (o *AddonOut) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddonOut) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddonOut) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetConfig

`func (o *AddonOut) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddonOut) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddonOut) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddonOut) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetYamlConfig

`func (o *AddonOut) GetYamlConfig() string`

GetYamlConfig returns the YamlConfig field if non-nil, zero value otherwise.

### GetYamlConfigOk

`func (o *AddonOut) GetYamlConfigOk() (*string, bool)`

GetYamlConfigOk returns a tuple with the YamlConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYamlConfig

`func (o *AddonOut) SetYamlConfig(v string)`

SetYamlConfig sets YamlConfig field to given value.


### GetConfigType

`func (o *AddonOut) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *AddonOut) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *AddonOut) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


