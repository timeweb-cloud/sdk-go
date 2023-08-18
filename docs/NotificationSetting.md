# NotificationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | [**NotificationSettingChannels**](NotificationSettingChannels.md) |  | 
**Group** | **string** | Строка, указывающая название группы уведомления. Может быть «security», «monitoring» или «finances». | 
**Type** | [**NotificationSettingType**](NotificationSettingType.md) |  | 

## Methods

### NewNotificationSetting

`func NewNotificationSetting(channels NotificationSettingChannels, group string, type_ NotificationSettingType, ) *NotificationSetting`

NewNotificationSetting instantiates a new NotificationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingWithDefaults

`func NewNotificationSettingWithDefaults() *NotificationSetting`

NewNotificationSettingWithDefaults instantiates a new NotificationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *NotificationSetting) GetChannels() NotificationSettingChannels`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationSetting) GetChannelsOk() (*NotificationSettingChannels, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationSetting) SetChannels(v NotificationSettingChannels)`

SetChannels sets Channels field to given value.


### GetGroup

`func (o *NotificationSetting) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *NotificationSetting) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *NotificationSetting) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetType

`func (o *NotificationSetting) GetType() NotificationSettingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationSetting) GetTypeOk() (*NotificationSettingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationSetting) SetType(v NotificationSettingType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


