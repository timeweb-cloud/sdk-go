# AiGatewayImagesResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**B64Json** | Pointer to **string** | Изображение в Base64 | [optional] 
**Url** | Pointer to **string** | Ссылка на изображение — для моделей, которые отдают ссылку | [optional] 
**RevisedPrompt** | Pointer to **string** | Промпт, уточнённый моделью | [optional] 

## Methods

### NewAiGatewayImagesResponseDataInner

`func NewAiGatewayImagesResponseDataInner() *AiGatewayImagesResponseDataInner`

NewAiGatewayImagesResponseDataInner instantiates a new AiGatewayImagesResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayImagesResponseDataInnerWithDefaults

`func NewAiGatewayImagesResponseDataInnerWithDefaults() *AiGatewayImagesResponseDataInner`

NewAiGatewayImagesResponseDataInnerWithDefaults instantiates a new AiGatewayImagesResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetB64Json

`func (o *AiGatewayImagesResponseDataInner) GetB64Json() string`

GetB64Json returns the B64Json field if non-nil, zero value otherwise.

### GetB64JsonOk

`func (o *AiGatewayImagesResponseDataInner) GetB64JsonOk() (*string, bool)`

GetB64JsonOk returns a tuple with the B64Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB64Json

`func (o *AiGatewayImagesResponseDataInner) SetB64Json(v string)`

SetB64Json sets B64Json field to given value.

### HasB64Json

`func (o *AiGatewayImagesResponseDataInner) HasB64Json() bool`

HasB64Json returns a boolean if a field has been set.

### GetUrl

`func (o *AiGatewayImagesResponseDataInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AiGatewayImagesResponseDataInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AiGatewayImagesResponseDataInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AiGatewayImagesResponseDataInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRevisedPrompt

`func (o *AiGatewayImagesResponseDataInner) GetRevisedPrompt() string`

GetRevisedPrompt returns the RevisedPrompt field if non-nil, zero value otherwise.

### GetRevisedPromptOk

`func (o *AiGatewayImagesResponseDataInner) GetRevisedPromptOk() (*string, bool)`

GetRevisedPromptOk returns a tuple with the RevisedPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisedPrompt

`func (o *AiGatewayImagesResponseDataInner) SetRevisedPrompt(v string)`

SetRevisedPrompt sets RevisedPrompt field to given value.

### HasRevisedPrompt

`func (o *AiGatewayImagesResponseDataInner) HasRevisedPrompt() bool`

HasRevisedPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


