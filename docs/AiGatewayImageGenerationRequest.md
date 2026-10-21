# AiGatewayImageGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID модели генерации изображений (&#x60;mode: image_generation&#x60; в &#x60;GET /v1/models&#x60;) | 
**Prompt** | **string** | Описание изображения | 
**N** | Pointer to **int32** | Количество изображений | [optional] [default to 1]
**Size** | Pointer to **string** | Размер изображения | [optional] 
**Quality** | Pointer to **string** | Качество изображения | [optional] 
**ResponseFormat** | Pointer to **string** | Формат изображения в ответе. Поддерживается не всеми моделями | [optional] 
**OutputFormat** | Pointer to **string** | Формат файла изображения | [optional] 

## Methods

### NewAiGatewayImageGenerationRequest

`func NewAiGatewayImageGenerationRequest(model string, prompt string, ) *AiGatewayImageGenerationRequest`

NewAiGatewayImageGenerationRequest instantiates a new AiGatewayImageGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayImageGenerationRequestWithDefaults

`func NewAiGatewayImageGenerationRequestWithDefaults() *AiGatewayImageGenerationRequest`

NewAiGatewayImageGenerationRequestWithDefaults instantiates a new AiGatewayImageGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiGatewayImageGenerationRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiGatewayImageGenerationRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiGatewayImageGenerationRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetPrompt

`func (o *AiGatewayImageGenerationRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiGatewayImageGenerationRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiGatewayImageGenerationRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetN

`func (o *AiGatewayImageGenerationRequest) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AiGatewayImageGenerationRequest) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AiGatewayImageGenerationRequest) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *AiGatewayImageGenerationRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### GetSize

`func (o *AiGatewayImageGenerationRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AiGatewayImageGenerationRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AiGatewayImageGenerationRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *AiGatewayImageGenerationRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetQuality

`func (o *AiGatewayImageGenerationRequest) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *AiGatewayImageGenerationRequest) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *AiGatewayImageGenerationRequest) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *AiGatewayImageGenerationRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetResponseFormat

`func (o *AiGatewayImageGenerationRequest) GetResponseFormat() string`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *AiGatewayImageGenerationRequest) GetResponseFormatOk() (*string, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *AiGatewayImageGenerationRequest) SetResponseFormat(v string)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *AiGatewayImageGenerationRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetOutputFormat

`func (o *AiGatewayImageGenerationRequest) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *AiGatewayImageGenerationRequest) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *AiGatewayImageGenerationRequest) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *AiGatewayImageGenerationRequest) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


