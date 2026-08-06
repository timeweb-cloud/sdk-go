# ConfigCacheQueryArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Режим учета query-параметров. - &#x60;all&#x60; — учитывать все параметры; - &#x60;whitelist&#x60; — учитывать только параметры из &#x60;list&#x60;; - &#x60;blacklist&#x60; — учитывать все параметры, кроме перечисленных в &#x60;list&#x60;. | 
**List** | Pointer to **[]string** | Список query-параметров для режимов &#x60;whitelist&#x60; и &#x60;blacklist&#x60;. | [optional] 

## Methods

### NewConfigCacheQueryArgs

`func NewConfigCacheQueryArgs(mode string, ) *ConfigCacheQueryArgs`

NewConfigCacheQueryArgs instantiates a new ConfigCacheQueryArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigCacheQueryArgsWithDefaults

`func NewConfigCacheQueryArgsWithDefaults() *ConfigCacheQueryArgs`

NewConfigCacheQueryArgsWithDefaults instantiates a new ConfigCacheQueryArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ConfigCacheQueryArgs) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ConfigCacheQueryArgs) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ConfigCacheQueryArgs) SetMode(v string)`

SetMode sets Mode field to given value.


### GetList

`func (o *ConfigCacheQueryArgs) GetList() []string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ConfigCacheQueryArgs) GetListOk() (*[]string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ConfigCacheQueryArgs) SetList(v []string)`

SetList sets List field to given value.

### HasList

`func (o *ConfigCacheQueryArgs) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


