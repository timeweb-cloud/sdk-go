# AiGatewayChatCompletionRequestStreamOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeUsage** | Pointer to **bool** | Отправить перед &#x60;data: [DONE]&#x60; последний чанк с расходом токенов в &#x60;usage&#x60; | [optional] 

## Methods

### NewAiGatewayChatCompletionRequestStreamOptions

`func NewAiGatewayChatCompletionRequestStreamOptions() *AiGatewayChatCompletionRequestStreamOptions`

NewAiGatewayChatCompletionRequestStreamOptions instantiates a new AiGatewayChatCompletionRequestStreamOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayChatCompletionRequestStreamOptionsWithDefaults

`func NewAiGatewayChatCompletionRequestStreamOptionsWithDefaults() *AiGatewayChatCompletionRequestStreamOptions`

NewAiGatewayChatCompletionRequestStreamOptionsWithDefaults instantiates a new AiGatewayChatCompletionRequestStreamOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeUsage

`func (o *AiGatewayChatCompletionRequestStreamOptions) GetIncludeUsage() bool`

GetIncludeUsage returns the IncludeUsage field if non-nil, zero value otherwise.

### GetIncludeUsageOk

`func (o *AiGatewayChatCompletionRequestStreamOptions) GetIncludeUsageOk() (*bool, bool)`

GetIncludeUsageOk returns a tuple with the IncludeUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUsage

`func (o *AiGatewayChatCompletionRequestStreamOptions) SetIncludeUsage(v bool)`

SetIncludeUsage sets IncludeUsage field to given value.

### HasIncludeUsage

`func (o *AiGatewayChatCompletionRequestStreamOptions) HasIncludeUsage() bool`

HasIncludeUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


