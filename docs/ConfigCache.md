# ConfigCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cdn** | Pointer to [**NullableConfigCacheCdn**](ConfigCacheCdn.md) |  | [optional] 
**Browser** | Pointer to [**NullableConfigCacheBrowser**](ConfigCacheBrowser.md) |  | [optional] 
**AlwaysOnline** | Pointer to [**NullableConfigCacheAlwaysOnline**](ConfigCacheAlwaysOnline.md) |  | [optional] 
**QueryArgs** | Pointer to [**NullableConfigCacheQueryArgs**](ConfigCacheQueryArgs.md) |  | [optional] 

## Methods

### NewConfigCache

`func NewConfigCache() *ConfigCache`

NewConfigCache instantiates a new ConfigCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigCacheWithDefaults

`func NewConfigCacheWithDefaults() *ConfigCache`

NewConfigCacheWithDefaults instantiates a new ConfigCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdn

`func (o *ConfigCache) GetCdn() ConfigCacheCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *ConfigCache) GetCdnOk() (*ConfigCacheCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *ConfigCache) SetCdn(v ConfigCacheCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *ConfigCache) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### SetCdnNil

`func (o *ConfigCache) SetCdnNil(b bool)`

 SetCdnNil sets the value for Cdn to be an explicit nil

### UnsetCdn
`func (o *ConfigCache) UnsetCdn()`

UnsetCdn ensures that no value is present for Cdn, not even an explicit nil
### GetBrowser

`func (o *ConfigCache) GetBrowser() ConfigCacheBrowser`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *ConfigCache) GetBrowserOk() (*ConfigCacheBrowser, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *ConfigCache) SetBrowser(v ConfigCacheBrowser)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *ConfigCache) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### SetBrowserNil

`func (o *ConfigCache) SetBrowserNil(b bool)`

 SetBrowserNil sets the value for Browser to be an explicit nil

### UnsetBrowser
`func (o *ConfigCache) UnsetBrowser()`

UnsetBrowser ensures that no value is present for Browser, not even an explicit nil
### GetAlwaysOnline

`func (o *ConfigCache) GetAlwaysOnline() ConfigCacheAlwaysOnline`

GetAlwaysOnline returns the AlwaysOnline field if non-nil, zero value otherwise.

### GetAlwaysOnlineOk

`func (o *ConfigCache) GetAlwaysOnlineOk() (*ConfigCacheAlwaysOnline, bool)`

GetAlwaysOnlineOk returns a tuple with the AlwaysOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysOnline

`func (o *ConfigCache) SetAlwaysOnline(v ConfigCacheAlwaysOnline)`

SetAlwaysOnline sets AlwaysOnline field to given value.

### HasAlwaysOnline

`func (o *ConfigCache) HasAlwaysOnline() bool`

HasAlwaysOnline returns a boolean if a field has been set.

### SetAlwaysOnlineNil

`func (o *ConfigCache) SetAlwaysOnlineNil(b bool)`

 SetAlwaysOnlineNil sets the value for AlwaysOnline to be an explicit nil

### UnsetAlwaysOnline
`func (o *ConfigCache) UnsetAlwaysOnline()`

UnsetAlwaysOnline ensures that no value is present for AlwaysOnline, not even an explicit nil
### GetQueryArgs

`func (o *ConfigCache) GetQueryArgs() ConfigCacheQueryArgs`

GetQueryArgs returns the QueryArgs field if non-nil, zero value otherwise.

### GetQueryArgsOk

`func (o *ConfigCache) GetQueryArgsOk() (*ConfigCacheQueryArgs, bool)`

GetQueryArgsOk returns a tuple with the QueryArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryArgs

`func (o *ConfigCache) SetQueryArgs(v ConfigCacheQueryArgs)`

SetQueryArgs sets QueryArgs field to given value.

### HasQueryArgs

`func (o *ConfigCache) HasQueryArgs() bool`

HasQueryArgs returns a boolean if a field has been set.

### SetQueryArgsNil

`func (o *ConfigCache) SetQueryArgsNil(b bool)`

 SetQueryArgsNil sets the value for QueryArgs to be an explicit nil

### UnsetQueryArgs
`func (o *ConfigCache) UnsetQueryArgs()`

UnsetQueryArgs ensures that no value is present for QueryArgs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


