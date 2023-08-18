# NotificationSettingChannels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | [**NotificationSettingChannel**](NotificationSettingChannel.md) |  | 
**Sms** | [**NotificationSettingChannel**](NotificationSettingChannel.md) |  | 
**Telegram** | [**NotificationSettingChannel**](NotificationSettingChannel.md) |  | 

## Methods

### NewNotificationSettingChannels

`func NewNotificationSettingChannels(email NotificationSettingChannel, sms NotificationSettingChannel, telegram NotificationSettingChannel, ) *NotificationSettingChannels`

NewNotificationSettingChannels instantiates a new NotificationSettingChannels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingChannelsWithDefaults

`func NewNotificationSettingChannelsWithDefaults() *NotificationSettingChannels`

NewNotificationSettingChannelsWithDefaults instantiates a new NotificationSettingChannels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *NotificationSettingChannels) GetEmail() NotificationSettingChannel`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NotificationSettingChannels) GetEmailOk() (*NotificationSettingChannel, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NotificationSettingChannels) SetEmail(v NotificationSettingChannel)`

SetEmail sets Email field to given value.


### GetSms

`func (o *NotificationSettingChannels) GetSms() NotificationSettingChannel`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *NotificationSettingChannels) GetSmsOk() (*NotificationSettingChannel, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *NotificationSettingChannels) SetSms(v NotificationSettingChannel)`

SetSms sets Sms field to given value.


### GetTelegram

`func (o *NotificationSettingChannels) GetTelegram() NotificationSettingChannel`

GetTelegram returns the Telegram field if non-nil, zero value otherwise.

### GetTelegramOk

`func (o *NotificationSettingChannels) GetTelegramOk() (*NotificationSettingChannel, bool)`

GetTelegramOk returns a tuple with the Telegram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegram

`func (o *NotificationSettingChannels) SetTelegram(v NotificationSettingChannel)`

SetTelegram sets Telegram field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


