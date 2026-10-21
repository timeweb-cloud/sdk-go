# AiSolutionTelegramBot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKey** | **string** | ID бота в решении: &#x60;b&lt;ID бота в Telegram&gt;&#x60; | 
**BotUsername** | **NullableString** | Имя пользователя бота без &#x60;@&#x60;. &#x60;null&#x60;, если Telegram его не вернул | 
**State** | **string** | Состояние бота: &#x60;pending_pairing&#x60; — бот подключён, но владелец ещё не подтверждён (ожидается код подтверждения); &#x60;connected&#x60; — владелец подтверждён, агент отвечает в этом боте | 

## Methods

### NewAiSolutionTelegramBot

`func NewAiSolutionTelegramBot(accountKey string, botUsername NullableString, state string, ) *AiSolutionTelegramBot`

NewAiSolutionTelegramBot instantiates a new AiSolutionTelegramBot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSolutionTelegramBotWithDefaults

`func NewAiSolutionTelegramBotWithDefaults() *AiSolutionTelegramBot`

NewAiSolutionTelegramBotWithDefaults instantiates a new AiSolutionTelegramBot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountKey

`func (o *AiSolutionTelegramBot) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *AiSolutionTelegramBot) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *AiSolutionTelegramBot) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.


### GetBotUsername

`func (o *AiSolutionTelegramBot) GetBotUsername() string`

GetBotUsername returns the BotUsername field if non-nil, zero value otherwise.

### GetBotUsernameOk

`func (o *AiSolutionTelegramBot) GetBotUsernameOk() (*string, bool)`

GetBotUsernameOk returns a tuple with the BotUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUsername

`func (o *AiSolutionTelegramBot) SetBotUsername(v string)`

SetBotUsername sets BotUsername field to given value.


### SetBotUsernameNil

`func (o *AiSolutionTelegramBot) SetBotUsernameNil(b bool)`

 SetBotUsernameNil sets the value for BotUsername to be an explicit nil

### UnsetBotUsername
`func (o *AiSolutionTelegramBot) UnsetBotUsername()`

UnsetBotUsername ensures that no value is present for BotUsername, not even an explicit nil
### GetState

`func (o *AiSolutionTelegramBot) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AiSolutionTelegramBot) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AiSolutionTelegramBot) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


