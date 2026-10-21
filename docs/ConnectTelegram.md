# ConnectTelegram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BotToken** | **string** | Токен бота, выданный @BotFather, в формате &#x60;&lt;id бота&gt;:&lt;секрет&gt;&#x60;, например &#x60;123456789:AAH...&#x60; | 

## Methods

### NewConnectTelegram

`func NewConnectTelegram(botToken string, ) *ConnectTelegram`

NewConnectTelegram instantiates a new ConnectTelegram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectTelegramWithDefaults

`func NewConnectTelegramWithDefaults() *ConnectTelegram`

NewConnectTelegramWithDefaults instantiates a new ConnectTelegram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBotToken

`func (o *ConnectTelegram) GetBotToken() string`

GetBotToken returns the BotToken field if non-nil, zero value otherwise.

### GetBotTokenOk

`func (o *ConnectTelegram) GetBotTokenOk() (*string, bool)`

GetBotTokenOk returns a tuple with the BotToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotToken

`func (o *ConnectTelegram) SetBotToken(v string)`

SetBotToken sets BotToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


