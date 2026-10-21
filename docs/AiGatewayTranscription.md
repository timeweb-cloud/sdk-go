# AiGatewayTranscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Текст | [optional] 
**Usage** | Pointer to **map[string]interface{}** | Расход токенов, если модель его возвращает | [optional] 

## Methods

### NewAiGatewayTranscription

`func NewAiGatewayTranscription() *AiGatewayTranscription`

NewAiGatewayTranscription instantiates a new AiGatewayTranscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayTranscriptionWithDefaults

`func NewAiGatewayTranscriptionWithDefaults() *AiGatewayTranscription`

NewAiGatewayTranscriptionWithDefaults instantiates a new AiGatewayTranscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *AiGatewayTranscription) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AiGatewayTranscription) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AiGatewayTranscription) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *AiGatewayTranscription) HasText() bool`

HasText returns a boolean if a field has been set.

### GetUsage

`func (o *AiGatewayTranscription) GetUsage() map[string]interface{}`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiGatewayTranscription) GetUsageOk() (*map[string]interface{}, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiGatewayTranscription) SetUsage(v map[string]interface{})`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiGatewayTranscription) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


