# ClearCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurgeType** | **string** | Тип очистки. - &#x60;full&#x60; — очистить весь кэш ресурса; - &#x60;partial&#x60; — очистить кэш только по путям из &#x60;paths&#x60;. | 
**Paths** | Pointer to **[]string** | Список путей к файлам, кэш которых нужно очистить. Обязателен при &#x60;purge_type&#x60; &#x3D; &#x60;partial&#x60;. | [optional] 

## Methods

### NewClearCache

`func NewClearCache(purgeType string, ) *ClearCache`

NewClearCache instantiates a new ClearCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearCacheWithDefaults

`func NewClearCacheWithDefaults() *ClearCache`

NewClearCacheWithDefaults instantiates a new ClearCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurgeType

`func (o *ClearCache) GetPurgeType() string`

GetPurgeType returns the PurgeType field if non-nil, zero value otherwise.

### GetPurgeTypeOk

`func (o *ClearCache) GetPurgeTypeOk() (*string, bool)`

GetPurgeTypeOk returns a tuple with the PurgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeType

`func (o *ClearCache) SetPurgeType(v string)`

SetPurgeType sets PurgeType field to given value.


### GetPaths

`func (o *ClearCache) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *ClearCache) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *ClearCache) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *ClearCache) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


