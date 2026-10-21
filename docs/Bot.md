# Bot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID бота в Timeweb Cloud | 
**AccountId** | **string** | Логин аккаунта, которому принадлежит бот | 
**AgentId** | **int32** | ID агента, который отвечает от имени бота | 
**Username** | **string** | Username бота в мессенджере (без &#x60;@&#x60;) | 
**FirstName** | **NullableString** | Имя бота в мессенджере | 
**Description** | **string** | Комментарий к боту. Пустая строка, если не задан | 
**Status** | **string** | Статус подключения: &#x60;pending&#x60; — ожидает подключения, &#x60;active&#x60; — работает, &#x60;error&#x60; — ошибка подключения | 
**WhitelistEnabled** | **bool** | Включено ли ограничение доступа по списку username. У MAX-ботов всегда &#x60;false&#x60; | 
**WhitelistUsernames** | **[]string** | Username пользователей, которым разрешено писать боту: в нижнем регистре, без &#x60;@&#x60;. У MAX-ботов всегда пустой | 
**CreatedAt** | **time.Time** | Дата и время создания | 
**UpdatedAt** | **time.Time** | Дата и время последнего обновления | 

## Methods

### NewBot

`func NewBot(id string, accountId string, agentId int32, username string, firstName NullableString, description string, status string, whitelistEnabled bool, whitelistUsernames []string, createdAt time.Time, updatedAt time.Time, ) *Bot`

NewBot instantiates a new Bot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotWithDefaults

`func NewBotWithDefaults() *Bot`

NewBotWithDefaults instantiates a new Bot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bot) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *Bot) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Bot) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Bot) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAgentId

`func (o *Bot) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *Bot) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *Bot) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.


### GetUsername

`func (o *Bot) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Bot) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Bot) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetFirstName

`func (o *Bot) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Bot) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Bot) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *Bot) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *Bot) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetDescription

`func (o *Bot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Bot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Bot) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *Bot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Bot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Bot) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWhitelistEnabled

`func (o *Bot) GetWhitelistEnabled() bool`

GetWhitelistEnabled returns the WhitelistEnabled field if non-nil, zero value otherwise.

### GetWhitelistEnabledOk

`func (o *Bot) GetWhitelistEnabledOk() (*bool, bool)`

GetWhitelistEnabledOk returns a tuple with the WhitelistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistEnabled

`func (o *Bot) SetWhitelistEnabled(v bool)`

SetWhitelistEnabled sets WhitelistEnabled field to given value.


### GetWhitelistUsernames

`func (o *Bot) GetWhitelistUsernames() []string`

GetWhitelistUsernames returns the WhitelistUsernames field if non-nil, zero value otherwise.

### GetWhitelistUsernamesOk

`func (o *Bot) GetWhitelistUsernamesOk() (*[]string, bool)`

GetWhitelistUsernamesOk returns a tuple with the WhitelistUsernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistUsernames

`func (o *Bot) SetWhitelistUsernames(v []string)`

SetWhitelistUsernames sets WhitelistUsernames field to given value.


### GetCreatedAt

`func (o *Bot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Bot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Bot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Bot) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Bot) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Bot) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


