# AiGatewayEmbeddingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** | Тип объекта | [optional] 
**Model** | Pointer to **string** | Модель, которая ответила | [optional] 
**Data** | Pointer to [**[]AiGatewayEmbeddingListDataInner**](AiGatewayEmbeddingListDataInner.md) | Эмбеддинги | [optional] 
**Usage** | Pointer to [**AiGatewayEmbeddingListUsage**](AiGatewayEmbeddingListUsage.md) |  | [optional] 

## Methods

### NewAiGatewayEmbeddingList

`func NewAiGatewayEmbeddingList() *AiGatewayEmbeddingList`

NewAiGatewayEmbeddingList instantiates a new AiGatewayEmbeddingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayEmbeddingListWithDefaults

`func NewAiGatewayEmbeddingListWithDefaults() *AiGatewayEmbeddingList`

NewAiGatewayEmbeddingListWithDefaults instantiates a new AiGatewayEmbeddingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *AiGatewayEmbeddingList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiGatewayEmbeddingList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiGatewayEmbeddingList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiGatewayEmbeddingList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetModel

`func (o *AiGatewayEmbeddingList) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiGatewayEmbeddingList) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiGatewayEmbeddingList) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiGatewayEmbeddingList) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetData

`func (o *AiGatewayEmbeddingList) GetData() []AiGatewayEmbeddingListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AiGatewayEmbeddingList) GetDataOk() (*[]AiGatewayEmbeddingListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AiGatewayEmbeddingList) SetData(v []AiGatewayEmbeddingListDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AiGatewayEmbeddingList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUsage

`func (o *AiGatewayEmbeddingList) GetUsage() AiGatewayEmbeddingListUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiGatewayEmbeddingList) GetUsageOk() (*AiGatewayEmbeddingListUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiGatewayEmbeddingList) SetUsage(v AiGatewayEmbeddingListUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiGatewayEmbeddingList) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


