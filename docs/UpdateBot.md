# UpdateBot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | Новый токен того же бота (например, после перевыпуска в @BotFather). Токен другого бота не принимается | [optional] 
**Description** | Pointer to **string** | Комментарий к боту | [optional] 
**WhitelistEnabled** | Pointer to **bool** | Ограничить доступ к боту списком &#x60;whitelist_usernames&#x60; | [optional] 
**WhitelistUsernames** | Pointer to **[]string** | Username пользователей Telegram, которым разрешено писать боту. Можно с &#x60;@&#x60; и в любом регистре — они будут приведены к нижнему регистру без &#x60;@&#x60;, дубли удаляются. Действует, только если &#x60;whitelist_enabled&#x60; равен &#x60;true&#x60; и список не пуст. Переданный список полностью заменяет текущий | [optional] 

## Methods

### NewUpdateBot

`func NewUpdateBot() *UpdateBot`

NewUpdateBot instantiates a new UpdateBot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBotWithDefaults

`func NewUpdateBotWithDefaults() *UpdateBot`

NewUpdateBotWithDefaults instantiates a new UpdateBot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *UpdateBot) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateBot) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateBot) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateBot) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateBot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateBot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateBot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateBot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWhitelistEnabled

`func (o *UpdateBot) GetWhitelistEnabled() bool`

GetWhitelistEnabled returns the WhitelistEnabled field if non-nil, zero value otherwise.

### GetWhitelistEnabledOk

`func (o *UpdateBot) GetWhitelistEnabledOk() (*bool, bool)`

GetWhitelistEnabledOk returns a tuple with the WhitelistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistEnabled

`func (o *UpdateBot) SetWhitelistEnabled(v bool)`

SetWhitelistEnabled sets WhitelistEnabled field to given value.

### HasWhitelistEnabled

`func (o *UpdateBot) HasWhitelistEnabled() bool`

HasWhitelistEnabled returns a boolean if a field has been set.

### GetWhitelistUsernames

`func (o *UpdateBot) GetWhitelistUsernames() []string`

GetWhitelistUsernames returns the WhitelistUsernames field if non-nil, zero value otherwise.

### GetWhitelistUsernamesOk

`func (o *UpdateBot) GetWhitelistUsernamesOk() (*[]string, bool)`

GetWhitelistUsernamesOk returns a tuple with the WhitelistUsernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistUsernames

`func (o *UpdateBot) SetWhitelistUsernames(v []string)`

SetWhitelistUsernames sets WhitelistUsernames field to given value.

### HasWhitelistUsernames

`func (o *UpdateBot) HasWhitelistUsernames() bool`

HasWhitelistUsernames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


