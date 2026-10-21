# CreateMaxBot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **int32** | ID агента, который будет отвечать от имени бота | 
**Token** | **string** | Токен бота, выданный @MasterBot в MAX | 
**Description** | Pointer to **string** | Комментарий к боту | [optional] 

## Methods

### NewCreateMaxBot

`func NewCreateMaxBot(agentId int32, token string, ) *CreateMaxBot`

NewCreateMaxBot instantiates a new CreateMaxBot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMaxBotWithDefaults

`func NewCreateMaxBotWithDefaults() *CreateMaxBot`

NewCreateMaxBotWithDefaults instantiates a new CreateMaxBot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *CreateMaxBot) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CreateMaxBot) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CreateMaxBot) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.


### GetToken

`func (o *CreateMaxBot) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateMaxBot) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateMaxBot) SetToken(v string)`

SetToken sets Token field to given value.


### GetDescription

`func (o *CreateMaxBot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMaxBot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMaxBot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMaxBot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


