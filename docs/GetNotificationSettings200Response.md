# GetNotificationSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationSettings** | [**[]NotificationSetting**](NotificationSetting.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetNotificationSettings200Response

`func NewGetNotificationSettings200Response(notificationSettings []NotificationSetting, meta Meta, ) *GetNotificationSettings200Response`

NewGetNotificationSettings200Response instantiates a new GetNotificationSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationSettings200ResponseWithDefaults

`func NewGetNotificationSettings200ResponseWithDefaults() *GetNotificationSettings200Response`

NewGetNotificationSettings200ResponseWithDefaults instantiates a new GetNotificationSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationSettings

`func (o *GetNotificationSettings200Response) GetNotificationSettings() []NotificationSetting`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *GetNotificationSettings200Response) GetNotificationSettingsOk() (*[]NotificationSetting, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *GetNotificationSettings200Response) SetNotificationSettings(v []NotificationSetting)`

SetNotificationSettings sets NotificationSettings field to given value.


### GetMeta

`func (o *GetNotificationSettings200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetNotificationSettings200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetNotificationSettings200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


