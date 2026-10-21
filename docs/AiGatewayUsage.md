# AiGatewayUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromptTokens** | Pointer to **int32** | Входные токены | [optional] 
**CompletionTokens** | Pointer to **int32** | Выходные токены | [optional] 
**TotalTokens** | Pointer to **int32** | Всего токенов | [optional] 
**PromptTokensDetails** | Pointer to [**NullableAiGatewayUsagePromptTokensDetails**](AiGatewayUsagePromptTokensDetails.md) |  | [optional] 
**CompletionTokensDetails** | Pointer to [**NullableAiGatewayUsageCompletionTokensDetails**](AiGatewayUsageCompletionTokensDetails.md) |  | [optional] 

## Methods

### NewAiGatewayUsage

`func NewAiGatewayUsage() *AiGatewayUsage`

NewAiGatewayUsage instantiates a new AiGatewayUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayUsageWithDefaults

`func NewAiGatewayUsageWithDefaults() *AiGatewayUsage`

NewAiGatewayUsageWithDefaults instantiates a new AiGatewayUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromptTokens

`func (o *AiGatewayUsage) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *AiGatewayUsage) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *AiGatewayUsage) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *AiGatewayUsage) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *AiGatewayUsage) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *AiGatewayUsage) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *AiGatewayUsage) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *AiGatewayUsage) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetTotalTokens

`func (o *AiGatewayUsage) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *AiGatewayUsage) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *AiGatewayUsage) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *AiGatewayUsage) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetPromptTokensDetails

`func (o *AiGatewayUsage) GetPromptTokensDetails() AiGatewayUsagePromptTokensDetails`

GetPromptTokensDetails returns the PromptTokensDetails field if non-nil, zero value otherwise.

### GetPromptTokensDetailsOk

`func (o *AiGatewayUsage) GetPromptTokensDetailsOk() (*AiGatewayUsagePromptTokensDetails, bool)`

GetPromptTokensDetailsOk returns a tuple with the PromptTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokensDetails

`func (o *AiGatewayUsage) SetPromptTokensDetails(v AiGatewayUsagePromptTokensDetails)`

SetPromptTokensDetails sets PromptTokensDetails field to given value.

### HasPromptTokensDetails

`func (o *AiGatewayUsage) HasPromptTokensDetails() bool`

HasPromptTokensDetails returns a boolean if a field has been set.

### SetPromptTokensDetailsNil

`func (o *AiGatewayUsage) SetPromptTokensDetailsNil(b bool)`

 SetPromptTokensDetailsNil sets the value for PromptTokensDetails to be an explicit nil

### UnsetPromptTokensDetails
`func (o *AiGatewayUsage) UnsetPromptTokensDetails()`

UnsetPromptTokensDetails ensures that no value is present for PromptTokensDetails, not even an explicit nil
### GetCompletionTokensDetails

`func (o *AiGatewayUsage) GetCompletionTokensDetails() AiGatewayUsageCompletionTokensDetails`

GetCompletionTokensDetails returns the CompletionTokensDetails field if non-nil, zero value otherwise.

### GetCompletionTokensDetailsOk

`func (o *AiGatewayUsage) GetCompletionTokensDetailsOk() (*AiGatewayUsageCompletionTokensDetails, bool)`

GetCompletionTokensDetailsOk returns a tuple with the CompletionTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokensDetails

`func (o *AiGatewayUsage) SetCompletionTokensDetails(v AiGatewayUsageCompletionTokensDetails)`

SetCompletionTokensDetails sets CompletionTokensDetails field to given value.

### HasCompletionTokensDetails

`func (o *AiGatewayUsage) HasCompletionTokensDetails() bool`

HasCompletionTokensDetails returns a boolean if a field has been set.

### SetCompletionTokensDetailsNil

`func (o *AiGatewayUsage) SetCompletionTokensDetailsNil(b bool)`

 SetCompletionTokensDetailsNil sets the value for CompletionTokensDetails to be an explicit nil

### UnsetCompletionTokensDetails
`func (o *AiGatewayUsage) UnsetCompletionTokensDetails()`

UnsetCompletionTokensDetails ensures that no value is present for CompletionTokensDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


