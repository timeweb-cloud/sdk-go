# UpdateNotificationSettingsRequestSettingsInnerChannels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Telegram** | Pointer to [**SettingCondition**](SettingCondition.md) |  | [optional] 
**Sms** | Pointer to [**SettingCondition**](SettingCondition.md) |  | [optional] 
**Email** | Pointer to [**SettingCondition**](SettingCondition.md) |  | [optional] 

## Methods

### NewUpdateNotificationSettingsRequestSettingsInnerChannels

`func NewUpdateNotificationSettingsRequestSettingsInnerChannels() *UpdateNotificationSettingsRequestSettingsInnerChannels`

NewUpdateNotificationSettingsRequestSettingsInnerChannels instantiates a new UpdateNotificationSettingsRequestSettingsInnerChannels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationSettingsRequestSettingsInnerChannelsWithDefaults

`func NewUpdateNotificationSettingsRequestSettingsInnerChannelsWithDefaults() *UpdateNotificationSettingsRequestSettingsInnerChannels`

NewUpdateNotificationSettingsRequestSettingsInnerChannelsWithDefaults instantiates a new UpdateNotificationSettingsRequestSettingsInnerChannels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTelegram

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) GetTelegram() SettingCondition`

GetTelegram returns the Telegram field if non-nil, zero value otherwise.

### GetTelegramOk

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) GetTelegramOk() (*SettingCondition, bool)`

GetTelegramOk returns a tuple with the Telegram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegram

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) SetTelegram(v SettingCondition)`

SetTelegram sets Telegram field to given value.

### HasTelegram

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) HasTelegram() bool`

HasTelegram returns a boolean if a field has been set.

### GetSms

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) GetSms() SettingCondition`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) GetSmsOk() (*SettingCondition, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) SetSms(v SettingCondition)`

SetSms sets Sms field to given value.

### HasSms

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) GetEmail() SettingCondition`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) GetEmailOk() (*SettingCondition, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) SetEmail(v SettingCondition)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateNotificationSettingsRequestSettingsInnerChannels) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


