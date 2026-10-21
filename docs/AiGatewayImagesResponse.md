# AiGatewayImagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int32** | Время создания, Unix-время в секундах | [optional] 
**Data** | Pointer to [**[]AiGatewayImagesResponseDataInner**](AiGatewayImagesResponseDataInner.md) | Изображения | [optional] 
**Usage** | Pointer to **map[string]interface{}** | Расход токенов, если модель его возвращает | [optional] 

## Methods

### NewAiGatewayImagesResponse

`func NewAiGatewayImagesResponse() *AiGatewayImagesResponse`

NewAiGatewayImagesResponse instantiates a new AiGatewayImagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayImagesResponseWithDefaults

`func NewAiGatewayImagesResponseWithDefaults() *AiGatewayImagesResponse`

NewAiGatewayImagesResponseWithDefaults instantiates a new AiGatewayImagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AiGatewayImagesResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AiGatewayImagesResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AiGatewayImagesResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AiGatewayImagesResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetData

`func (o *AiGatewayImagesResponse) GetData() []AiGatewayImagesResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AiGatewayImagesResponse) GetDataOk() (*[]AiGatewayImagesResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AiGatewayImagesResponse) SetData(v []AiGatewayImagesResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AiGatewayImagesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUsage

`func (o *AiGatewayImagesResponse) GetUsage() map[string]interface{}`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiGatewayImagesResponse) GetUsageOk() (*map[string]interface{}, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiGatewayImagesResponse) SetUsage(v map[string]interface{})`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiGatewayImagesResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


