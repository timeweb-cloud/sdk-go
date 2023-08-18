# UpdateNotificationSettingsRequestSettingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | [**UpdateNotificationSettingsRequestSettingsInnerChannels**](UpdateNotificationSettingsRequestSettingsInnerChannels.md) |  | 
**Type** | [**NotificationSettingType**](NotificationSettingType.md) |  | 

## Methods

### NewUpdateNotificationSettingsRequestSettingsInner

`func NewUpdateNotificationSettingsRequestSettingsInner(channels UpdateNotificationSettingsRequestSettingsInnerChannels, type_ NotificationSettingType, ) *UpdateNotificationSettingsRequestSettingsInner`

NewUpdateNotificationSettingsRequestSettingsInner instantiates a new UpdateNotificationSettingsRequestSettingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationSettingsRequestSettingsInnerWithDefaults

`func NewUpdateNotificationSettingsRequestSettingsInnerWithDefaults() *UpdateNotificationSettingsRequestSettingsInner`

NewUpdateNotificationSettingsRequestSettingsInnerWithDefaults instantiates a new UpdateNotificationSettingsRequestSettingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *UpdateNotificationSettingsRequestSettingsInner) GetChannels() UpdateNotificationSettingsRequestSettingsInnerChannels`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *UpdateNotificationSettingsRequestSettingsInner) GetChannelsOk() (*UpdateNotificationSettingsRequestSettingsInnerChannels, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *UpdateNotificationSettingsRequestSettingsInner) SetChannels(v UpdateNotificationSettingsRequestSettingsInnerChannels)`

SetChannels sets Channels field to given value.


### GetType

`func (o *UpdateNotificationSettingsRequestSettingsInner) GetType() NotificationSettingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNotificationSettingsRequestSettingsInner) GetTypeOk() (*NotificationSettingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNotificationSettingsRequestSettingsInner) SetType(v NotificationSettingType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


