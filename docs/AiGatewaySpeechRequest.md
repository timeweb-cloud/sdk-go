# AiGatewaySpeechRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID модели синтеза речи (&#x60;mode: audio_speech&#x60; в &#x60;GET /v1/models&#x60;) | 
**Input** | **string** | Текст для озвучивания | 
**Voice** | **string** | Голос | 
**Instructions** | Pointer to **string** | Инструкции к манере речи — для моделей, которые их поддерживают | [optional] 
**ResponseFormat** | Pointer to **string** | Формат аудио, например &#x60;mp3&#x60;, &#x60;wav&#x60;, &#x60;opus&#x60;. По умолчанию — &#x60;mp3&#x60; | [optional] 
**Speed** | Pointer to **float32** | Скорость речи | [optional] 

## Methods

### NewAiGatewaySpeechRequest

`func NewAiGatewaySpeechRequest(model string, input string, voice string, ) *AiGatewaySpeechRequest`

NewAiGatewaySpeechRequest instantiates a new AiGatewaySpeechRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewaySpeechRequestWithDefaults

`func NewAiGatewaySpeechRequestWithDefaults() *AiGatewaySpeechRequest`

NewAiGatewaySpeechRequestWithDefaults instantiates a new AiGatewaySpeechRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiGatewaySpeechRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiGatewaySpeechRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiGatewaySpeechRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *AiGatewaySpeechRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AiGatewaySpeechRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AiGatewaySpeechRequest) SetInput(v string)`

SetInput sets Input field to given value.


### GetVoice

`func (o *AiGatewaySpeechRequest) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *AiGatewaySpeechRequest) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *AiGatewaySpeechRequest) SetVoice(v string)`

SetVoice sets Voice field to given value.


### GetInstructions

`func (o *AiGatewaySpeechRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AiGatewaySpeechRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AiGatewaySpeechRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AiGatewaySpeechRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetResponseFormat

`func (o *AiGatewaySpeechRequest) GetResponseFormat() string`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *AiGatewaySpeechRequest) GetResponseFormatOk() (*string, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *AiGatewaySpeechRequest) SetResponseFormat(v string)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *AiGatewaySpeechRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetSpeed

`func (o *AiGatewaySpeechRequest) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *AiGatewaySpeechRequest) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *AiGatewaySpeechRequest) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *AiGatewaySpeechRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


