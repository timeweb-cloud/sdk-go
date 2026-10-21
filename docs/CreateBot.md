# CreateBot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **int32** | ID агента, который будет отвечать от имени бота | 
**Token** | **string** | Токен бота, выданный @BotFather | 
**Description** | Pointer to **string** | Комментарий к боту | [optional] 
**WhitelistEnabled** | Pointer to **bool** | Ограничить доступ к боту списком &#x60;whitelist_usernames&#x60; | [optional] 
**WhitelistUsernames** | Pointer to **[]string** | Username пользователей Telegram, которым разрешено писать боту. Можно с &#x60;@&#x60; и в любом регистре — они будут приведены к нижнему регистру без &#x60;@&#x60;, дубли удаляются. Действует, только если &#x60;whitelist_enabled&#x60; равен &#x60;true&#x60; и список не пуст | [optional] 

## Methods

### NewCreateBot

`func NewCreateBot(agentId int32, token string, ) *CreateBot`

NewCreateBot instantiates a new CreateBot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBotWithDefaults

`func NewCreateBotWithDefaults() *CreateBot`

NewCreateBotWithDefaults instantiates a new CreateBot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *CreateBot) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CreateBot) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CreateBot) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.


### GetToken

`func (o *CreateBot) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateBot) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateBot) SetToken(v string)`

SetToken sets Token field to given value.


### GetDescription

`func (o *CreateBot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateBot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWhitelistEnabled

`func (o *CreateBot) GetWhitelistEnabled() bool`

GetWhitelistEnabled returns the WhitelistEnabled field if non-nil, zero value otherwise.

### GetWhitelistEnabledOk

`func (o *CreateBot) GetWhitelistEnabledOk() (*bool, bool)`

GetWhitelistEnabledOk returns a tuple with the WhitelistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistEnabled

`func (o *CreateBot) SetWhitelistEnabled(v bool)`

SetWhitelistEnabled sets WhitelistEnabled field to given value.

### HasWhitelistEnabled

`func (o *CreateBot) HasWhitelistEnabled() bool`

HasWhitelistEnabled returns a boolean if a field has been set.

### GetWhitelistUsernames

`func (o *CreateBot) GetWhitelistUsernames() []string`

GetWhitelistUsernames returns the WhitelistUsernames field if non-nil, zero value otherwise.

### GetWhitelistUsernamesOk

`func (o *CreateBot) GetWhitelistUsernamesOk() (*[]string, bool)`

GetWhitelistUsernamesOk returns a tuple with the WhitelistUsernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistUsernames

`func (o *CreateBot) SetWhitelistUsernames(v []string)`

SetWhitelistUsernames sets WhitelistUsernames field to given value.

### HasWhitelistUsernames

`func (o *CreateBot) HasWhitelistUsernames() bool`

HasWhitelistUsernames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


