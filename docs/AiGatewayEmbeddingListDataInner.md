# AiGatewayEmbeddingListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** | Тип объекта | [optional] 
**Index** | Pointer to **int32** | Номер текста во входном списке | [optional] 
**Embedding** | Pointer to **[]float32** | Вектор | [optional] 

## Methods

### NewAiGatewayEmbeddingListDataInner

`func NewAiGatewayEmbeddingListDataInner() *AiGatewayEmbeddingListDataInner`

NewAiGatewayEmbeddingListDataInner instantiates a new AiGatewayEmbeddingListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayEmbeddingListDataInnerWithDefaults

`func NewAiGatewayEmbeddingListDataInnerWithDefaults() *AiGatewayEmbeddingListDataInner`

NewAiGatewayEmbeddingListDataInnerWithDefaults instantiates a new AiGatewayEmbeddingListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *AiGatewayEmbeddingListDataInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiGatewayEmbeddingListDataInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiGatewayEmbeddingListDataInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiGatewayEmbeddingListDataInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetIndex

`func (o *AiGatewayEmbeddingListDataInner) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AiGatewayEmbeddingListDataInner) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AiGatewayEmbeddingListDataInner) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *AiGatewayEmbeddingListDataInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetEmbedding

`func (o *AiGatewayEmbeddingListDataInner) GetEmbedding() []float32`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *AiGatewayEmbeddingListDataInner) GetEmbeddingOk() (*[]float32, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *AiGatewayEmbeddingListDataInner) SetEmbedding(v []float32)`

SetEmbedding sets Embedding field to given value.

### HasEmbedding

`func (o *AiGatewayEmbeddingListDataInner) HasEmbedding() bool`

HasEmbedding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


