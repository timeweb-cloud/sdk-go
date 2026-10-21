# ToolsCallResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**[]TextContent**](TextContent.md) | Блоки результата | 
**IsError** | Pointer to **bool** | &#x60;true&#x60;, если инструмент завершился ошибкой; в тексте — JSON с полем &#x60;error&#x60; | [optional] 

## Methods

### NewToolsCallResult

`func NewToolsCallResult(content []TextContent, ) *ToolsCallResult`

NewToolsCallResult instantiates a new ToolsCallResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolsCallResultWithDefaults

`func NewToolsCallResultWithDefaults() *ToolsCallResult`

NewToolsCallResultWithDefaults instantiates a new ToolsCallResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ToolsCallResult) GetContent() []TextContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ToolsCallResult) GetContentOk() (*[]TextContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ToolsCallResult) SetContent(v []TextContent)`

SetContent sets Content field to given value.


### GetIsError

`func (o *ToolsCallResult) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *ToolsCallResult) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *ToolsCallResult) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *ToolsCallResult) HasIsError() bool`

HasIsError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


