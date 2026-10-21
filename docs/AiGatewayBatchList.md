# AiGatewayBatchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** | Тип объекта | [optional] 
**Data** | Pointer to [**[]AiGatewayBatch**](AiGatewayBatch.md) | Батчи | [optional] 
**FirstId** | Pointer to **NullableString** | ID первого батча на странице | [optional] 
**LastId** | Pointer to **NullableString** | ID последнего батча на странице — передайте его в &#x60;after&#x60; для следующей страницы | [optional] 
**HasMore** | Pointer to **bool** | Есть ли следующая страница | [optional] 

## Methods

### NewAiGatewayBatchList

`func NewAiGatewayBatchList() *AiGatewayBatchList`

NewAiGatewayBatchList instantiates a new AiGatewayBatchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayBatchListWithDefaults

`func NewAiGatewayBatchListWithDefaults() *AiGatewayBatchList`

NewAiGatewayBatchListWithDefaults instantiates a new AiGatewayBatchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *AiGatewayBatchList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiGatewayBatchList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiGatewayBatchList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiGatewayBatchList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *AiGatewayBatchList) GetData() []AiGatewayBatch`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AiGatewayBatchList) GetDataOk() (*[]AiGatewayBatch, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AiGatewayBatchList) SetData(v []AiGatewayBatch)`

SetData sets Data field to given value.

### HasData

`func (o *AiGatewayBatchList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFirstId

`func (o *AiGatewayBatchList) GetFirstId() string`

GetFirstId returns the FirstId field if non-nil, zero value otherwise.

### GetFirstIdOk

`func (o *AiGatewayBatchList) GetFirstIdOk() (*string, bool)`

GetFirstIdOk returns a tuple with the FirstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstId

`func (o *AiGatewayBatchList) SetFirstId(v string)`

SetFirstId sets FirstId field to given value.

### HasFirstId

`func (o *AiGatewayBatchList) HasFirstId() bool`

HasFirstId returns a boolean if a field has been set.

### SetFirstIdNil

`func (o *AiGatewayBatchList) SetFirstIdNil(b bool)`

 SetFirstIdNil sets the value for FirstId to be an explicit nil

### UnsetFirstId
`func (o *AiGatewayBatchList) UnsetFirstId()`

UnsetFirstId ensures that no value is present for FirstId, not even an explicit nil
### GetLastId

`func (o *AiGatewayBatchList) GetLastId() string`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *AiGatewayBatchList) GetLastIdOk() (*string, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *AiGatewayBatchList) SetLastId(v string)`

SetLastId sets LastId field to given value.

### HasLastId

`func (o *AiGatewayBatchList) HasLastId() bool`

HasLastId returns a boolean if a field has been set.

### SetLastIdNil

`func (o *AiGatewayBatchList) SetLastIdNil(b bool)`

 SetLastIdNil sets the value for LastId to be an explicit nil

### UnsetLastId
`func (o *AiGatewayBatchList) UnsetLastId()`

UnsetLastId ensures that no value is present for LastId, not even an explicit nil
### GetHasMore

`func (o *AiGatewayBatchList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *AiGatewayBatchList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *AiGatewayBatchList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *AiGatewayBatchList) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


