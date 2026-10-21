# AiGatewayEmbeddingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID модели эмбеддингов (&#x60;mode: embedding&#x60; в &#x60;GET /v1/models&#x60;) | 
**Input** | **string** | Текст или список текстов | 
**Dimensions** | Pointer to **int32** | Размерность вектора — для моделей, которые её поддерживают | [optional] 
**EncodingFormat** | Pointer to **string** | Формат векторов | [optional] [default to "float"]

## Methods

### NewAiGatewayEmbeddingRequest

`func NewAiGatewayEmbeddingRequest(model string, input string, ) *AiGatewayEmbeddingRequest`

NewAiGatewayEmbeddingRequest instantiates a new AiGatewayEmbeddingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayEmbeddingRequestWithDefaults

`func NewAiGatewayEmbeddingRequestWithDefaults() *AiGatewayEmbeddingRequest`

NewAiGatewayEmbeddingRequestWithDefaults instantiates a new AiGatewayEmbeddingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiGatewayEmbeddingRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiGatewayEmbeddingRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiGatewayEmbeddingRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *AiGatewayEmbeddingRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AiGatewayEmbeddingRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AiGatewayEmbeddingRequest) SetInput(v string)`

SetInput sets Input field to given value.


### GetDimensions

`func (o *AiGatewayEmbeddingRequest) GetDimensions() int32`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *AiGatewayEmbeddingRequest) GetDimensionsOk() (*int32, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *AiGatewayEmbeddingRequest) SetDimensions(v int32)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *AiGatewayEmbeddingRequest) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetEncodingFormat

`func (o *AiGatewayEmbeddingRequest) GetEncodingFormat() string`

GetEncodingFormat returns the EncodingFormat field if non-nil, zero value otherwise.

### GetEncodingFormatOk

`func (o *AiGatewayEmbeddingRequest) GetEncodingFormatOk() (*string, bool)`

GetEncodingFormatOk returns a tuple with the EncodingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingFormat

`func (o *AiGatewayEmbeddingRequest) SetEncodingFormat(v string)`

SetEncodingFormat sets EncodingFormat field to given value.

### HasEncodingFormat

`func (o *AiGatewayEmbeddingRequest) HasEncodingFormat() bool`

HasEncodingFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


