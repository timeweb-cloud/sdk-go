# ModelBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MultiplierIn** | **float32** | Множитель цены входных токенов | 
**MultiplierOut** | **float32** | Множитель цены выходных токенов | 
**MultiplierCacheRead** | **float32** | Множитель цены токенов, прочитанных из кэша | 
**MultiplierCacheWrite** | **float32** | Множитель цены токенов, записанных в кэш | 
**Endpoints** | **[]string** | Эндпоинты, допустимые в строках batch-файла | 
**SupportsImages** | **bool** | Принимает ли модель изображения в batch-запросах | 

## Methods

### NewModelBatch

`func NewModelBatch(multiplierIn float32, multiplierOut float32, multiplierCacheRead float32, multiplierCacheWrite float32, endpoints []string, supportsImages bool, ) *ModelBatch`

NewModelBatch instantiates a new ModelBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelBatchWithDefaults

`func NewModelBatchWithDefaults() *ModelBatch`

NewModelBatchWithDefaults instantiates a new ModelBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMultiplierIn

`func (o *ModelBatch) GetMultiplierIn() float32`

GetMultiplierIn returns the MultiplierIn field if non-nil, zero value otherwise.

### GetMultiplierInOk

`func (o *ModelBatch) GetMultiplierInOk() (*float32, bool)`

GetMultiplierInOk returns a tuple with the MultiplierIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierIn

`func (o *ModelBatch) SetMultiplierIn(v float32)`

SetMultiplierIn sets MultiplierIn field to given value.


### GetMultiplierOut

`func (o *ModelBatch) GetMultiplierOut() float32`

GetMultiplierOut returns the MultiplierOut field if non-nil, zero value otherwise.

### GetMultiplierOutOk

`func (o *ModelBatch) GetMultiplierOutOk() (*float32, bool)`

GetMultiplierOutOk returns a tuple with the MultiplierOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierOut

`func (o *ModelBatch) SetMultiplierOut(v float32)`

SetMultiplierOut sets MultiplierOut field to given value.


### GetMultiplierCacheRead

`func (o *ModelBatch) GetMultiplierCacheRead() float32`

GetMultiplierCacheRead returns the MultiplierCacheRead field if non-nil, zero value otherwise.

### GetMultiplierCacheReadOk

`func (o *ModelBatch) GetMultiplierCacheReadOk() (*float32, bool)`

GetMultiplierCacheReadOk returns a tuple with the MultiplierCacheRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierCacheRead

`func (o *ModelBatch) SetMultiplierCacheRead(v float32)`

SetMultiplierCacheRead sets MultiplierCacheRead field to given value.


### GetMultiplierCacheWrite

`func (o *ModelBatch) GetMultiplierCacheWrite() float32`

GetMultiplierCacheWrite returns the MultiplierCacheWrite field if non-nil, zero value otherwise.

### GetMultiplierCacheWriteOk

`func (o *ModelBatch) GetMultiplierCacheWriteOk() (*float32, bool)`

GetMultiplierCacheWriteOk returns a tuple with the MultiplierCacheWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierCacheWrite

`func (o *ModelBatch) SetMultiplierCacheWrite(v float32)`

SetMultiplierCacheWrite sets MultiplierCacheWrite field to given value.


### GetEndpoints

`func (o *ModelBatch) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ModelBatch) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ModelBatch) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.


### GetSupportsImages

`func (o *ModelBatch) GetSupportsImages() bool`

GetSupportsImages returns the SupportsImages field if non-nil, zero value otherwise.

### GetSupportsImagesOk

`func (o *ModelBatch) GetSupportsImagesOk() (*bool, bool)`

GetSupportsImagesOk returns a tuple with the SupportsImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsImages

`func (o *ModelBatch) SetSupportsImages(v bool)`

SetSupportsImages sets SupportsImages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


