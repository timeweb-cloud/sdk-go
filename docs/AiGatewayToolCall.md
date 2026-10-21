# AiGatewayToolCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID вызова. Передайте его в &#x60;tool_call_id&#x60; сообщения с результатом | [optional] 
**Type** | Pointer to **string** | Тип вызова | [optional] 
**Function** | Pointer to [**AiGatewayToolCallFunction**](AiGatewayToolCallFunction.md) |  | [optional] 

## Methods

### NewAiGatewayToolCall

`func NewAiGatewayToolCall() *AiGatewayToolCall`

NewAiGatewayToolCall instantiates a new AiGatewayToolCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayToolCallWithDefaults

`func NewAiGatewayToolCallWithDefaults() *AiGatewayToolCall`

NewAiGatewayToolCallWithDefaults instantiates a new AiGatewayToolCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiGatewayToolCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiGatewayToolCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiGatewayToolCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiGatewayToolCall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AiGatewayToolCall) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiGatewayToolCall) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiGatewayToolCall) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AiGatewayToolCall) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFunction

`func (o *AiGatewayToolCall) GetFunction() AiGatewayToolCallFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *AiGatewayToolCall) GetFunctionOk() (*AiGatewayToolCallFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *AiGatewayToolCall) SetFunction(v AiGatewayToolCallFunction)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *AiGatewayToolCall) HasFunction() bool`

HasFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


