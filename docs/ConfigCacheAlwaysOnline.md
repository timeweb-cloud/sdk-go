# ConfigCacheAlwaysOnline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaleConditions** | **[]string** | Условия, при которых из кэша отдается устаревший контент. | 

## Methods

### NewConfigCacheAlwaysOnline

`func NewConfigCacheAlwaysOnline(staleConditions []string, ) *ConfigCacheAlwaysOnline`

NewConfigCacheAlwaysOnline instantiates a new ConfigCacheAlwaysOnline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigCacheAlwaysOnlineWithDefaults

`func NewConfigCacheAlwaysOnlineWithDefaults() *ConfigCacheAlwaysOnline`

NewConfigCacheAlwaysOnlineWithDefaults instantiates a new ConfigCacheAlwaysOnline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaleConditions

`func (o *ConfigCacheAlwaysOnline) GetStaleConditions() []string`

GetStaleConditions returns the StaleConditions field if non-nil, zero value otherwise.

### GetStaleConditionsOk

`func (o *ConfigCacheAlwaysOnline) GetStaleConditionsOk() (*[]string, bool)`

GetStaleConditionsOk returns a tuple with the StaleConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleConditions

`func (o *ConfigCacheAlwaysOnline) SetStaleConditions(v []string)`

SetStaleConditions sets StaleConditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


