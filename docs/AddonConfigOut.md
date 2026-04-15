# AddonConfigOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID конфигурации дополнения | 
**Type** | **string** | Тип дополнения | 
**Version** | **string** | Версия дополнения | 
**Dependencies** | **[]string** | Зависимости дополнения | 
**YamlConfig** | **string** | YAML конфигурации дополнения | 

## Methods

### NewAddonConfigOut

`func NewAddonConfigOut(id int32, type_ string, version string, dependencies []string, yamlConfig string, ) *AddonConfigOut`

NewAddonConfigOut instantiates a new AddonConfigOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonConfigOutWithDefaults

`func NewAddonConfigOutWithDefaults() *AddonConfigOut`

NewAddonConfigOutWithDefaults instantiates a new AddonConfigOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddonConfigOut) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonConfigOut) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonConfigOut) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *AddonConfigOut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddonConfigOut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddonConfigOut) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *AddonConfigOut) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddonConfigOut) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddonConfigOut) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDependencies

`func (o *AddonConfigOut) GetDependencies() []string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *AddonConfigOut) GetDependenciesOk() (*[]string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *AddonConfigOut) SetDependencies(v []string)`

SetDependencies sets Dependencies field to given value.


### GetYamlConfig

`func (o *AddonConfigOut) GetYamlConfig() string`

GetYamlConfig returns the YamlConfig field if non-nil, zero value otherwise.

### GetYamlConfigOk

`func (o *AddonConfigOut) GetYamlConfigOk() (*string, bool)`

GetYamlConfigOk returns a tuple with the YamlConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYamlConfig

`func (o *AddonConfigOut) SetYamlConfig(v string)`

SetYamlConfig sets YamlConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


