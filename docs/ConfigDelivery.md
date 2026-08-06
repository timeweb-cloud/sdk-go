# ConfigDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Http3** | Pointer to **bool** | Раздавать контент по HTTP/3 (QUIC). | [optional] 
**Gzip** | Pointer to **bool** | Сжимать ответы алгоритмом gzip. | [optional] 
**LargeFiles** | Pointer to **bool** | Режим раздачи больших файлов. | [optional] 
**SliceSize** | Pointer to **NullableInt32** | Размер слайса в мегабайтах — файл забирается с источника и кэшируется частями такого размера. &#x60;null&#x60; — не использовать слайсинг. | [optional] 
**ImageOptimization** | Pointer to **bool** | Оптимизировать изображения на лету. | [optional] 
**Packaging** | Pointer to [**NullableConfigDeliveryPackaging**](ConfigDeliveryPackaging.md) |  | [optional] 

## Methods

### NewConfigDelivery

`func NewConfigDelivery() *ConfigDelivery`

NewConfigDelivery instantiates a new ConfigDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigDeliveryWithDefaults

`func NewConfigDeliveryWithDefaults() *ConfigDelivery`

NewConfigDeliveryWithDefaults instantiates a new ConfigDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttp3

`func (o *ConfigDelivery) GetHttp3() bool`

GetHttp3 returns the Http3 field if non-nil, zero value otherwise.

### GetHttp3Ok

`func (o *ConfigDelivery) GetHttp3Ok() (*bool, bool)`

GetHttp3Ok returns a tuple with the Http3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp3

`func (o *ConfigDelivery) SetHttp3(v bool)`

SetHttp3 sets Http3 field to given value.

### HasHttp3

`func (o *ConfigDelivery) HasHttp3() bool`

HasHttp3 returns a boolean if a field has been set.

### GetGzip

`func (o *ConfigDelivery) GetGzip() bool`

GetGzip returns the Gzip field if non-nil, zero value otherwise.

### GetGzipOk

`func (o *ConfigDelivery) GetGzipOk() (*bool, bool)`

GetGzipOk returns a tuple with the Gzip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzip

`func (o *ConfigDelivery) SetGzip(v bool)`

SetGzip sets Gzip field to given value.

### HasGzip

`func (o *ConfigDelivery) HasGzip() bool`

HasGzip returns a boolean if a field has been set.

### GetLargeFiles

`func (o *ConfigDelivery) GetLargeFiles() bool`

GetLargeFiles returns the LargeFiles field if non-nil, zero value otherwise.

### GetLargeFilesOk

`func (o *ConfigDelivery) GetLargeFilesOk() (*bool, bool)`

GetLargeFilesOk returns a tuple with the LargeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeFiles

`func (o *ConfigDelivery) SetLargeFiles(v bool)`

SetLargeFiles sets LargeFiles field to given value.

### HasLargeFiles

`func (o *ConfigDelivery) HasLargeFiles() bool`

HasLargeFiles returns a boolean if a field has been set.

### GetSliceSize

`func (o *ConfigDelivery) GetSliceSize() int32`

GetSliceSize returns the SliceSize field if non-nil, zero value otherwise.

### GetSliceSizeOk

`func (o *ConfigDelivery) GetSliceSizeOk() (*int32, bool)`

GetSliceSizeOk returns a tuple with the SliceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceSize

`func (o *ConfigDelivery) SetSliceSize(v int32)`

SetSliceSize sets SliceSize field to given value.

### HasSliceSize

`func (o *ConfigDelivery) HasSliceSize() bool`

HasSliceSize returns a boolean if a field has been set.

### SetSliceSizeNil

`func (o *ConfigDelivery) SetSliceSizeNil(b bool)`

 SetSliceSizeNil sets the value for SliceSize to be an explicit nil

### UnsetSliceSize
`func (o *ConfigDelivery) UnsetSliceSize()`

UnsetSliceSize ensures that no value is present for SliceSize, not even an explicit nil
### GetImageOptimization

`func (o *ConfigDelivery) GetImageOptimization() bool`

GetImageOptimization returns the ImageOptimization field if non-nil, zero value otherwise.

### GetImageOptimizationOk

`func (o *ConfigDelivery) GetImageOptimizationOk() (*bool, bool)`

GetImageOptimizationOk returns a tuple with the ImageOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOptimization

`func (o *ConfigDelivery) SetImageOptimization(v bool)`

SetImageOptimization sets ImageOptimization field to given value.

### HasImageOptimization

`func (o *ConfigDelivery) HasImageOptimization() bool`

HasImageOptimization returns a boolean if a field has been set.

### GetPackaging

`func (o *ConfigDelivery) GetPackaging() ConfigDeliveryPackaging`

GetPackaging returns the Packaging field if non-nil, zero value otherwise.

### GetPackagingOk

`func (o *ConfigDelivery) GetPackagingOk() (*ConfigDeliveryPackaging, bool)`

GetPackagingOk returns a tuple with the Packaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackaging

`func (o *ConfigDelivery) SetPackaging(v ConfigDeliveryPackaging)`

SetPackaging sets Packaging field to given value.

### HasPackaging

`func (o *ConfigDelivery) HasPackaging() bool`

HasPackaging returns a boolean if a field has been set.

### SetPackagingNil

`func (o *ConfigDelivery) SetPackagingNil(b bool)`

 SetPackagingNil sets the value for Packaging to be an explicit nil

### UnsetPackaging
`func (o *ConfigDelivery) UnsetPackaging()`

UnsetPackaging ensures that no value is present for Packaging, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


