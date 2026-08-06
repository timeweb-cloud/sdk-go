# PreloadCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | **[]string** | Список путей к файлам, которые нужно заранее загрузить в кэш узлов CDN. | 

## Methods

### NewPreloadCache

`func NewPreloadCache(paths []string, ) *PreloadCache`

NewPreloadCache instantiates a new PreloadCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreloadCacheWithDefaults

`func NewPreloadCacheWithDefaults() *PreloadCache`

NewPreloadCacheWithDefaults instantiates a new PreloadCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *PreloadCache) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *PreloadCache) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *PreloadCache) SetPaths(v []string)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


