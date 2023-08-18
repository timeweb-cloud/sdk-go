# MailboxSpamFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включен ли спам-фильтр | 
**Action** | **string** | Что делать с письмами, которые попадают в спам. \\  Параметры: \\  &#x60;move_to_directory&#x60; - переместить в паку спам; \\  &#x60;forward&#x60; - переслать письмо на другой адрес; \\  &#x60;delete&#x60; - удалить письмо; \\  &#x60;tag&#x60; - пометить письмо; | 
**ForwardTo** | **string** | Адрес для пересылки при выбранном действии &#x60;forward&#x60; из параметра &#x60;action&#x60; | 
**WhiteList** | **[]string** | Белый список адресов от которых письма не будут попадать в спам | 

## Methods

### NewMailboxSpamFilter

`func NewMailboxSpamFilter(isEnabled bool, action string, forwardTo string, whiteList []string, ) *MailboxSpamFilter`

NewMailboxSpamFilter instantiates a new MailboxSpamFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxSpamFilterWithDefaults

`func NewMailboxSpamFilterWithDefaults() *MailboxSpamFilter`

NewMailboxSpamFilterWithDefaults instantiates a new MailboxSpamFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *MailboxSpamFilter) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MailboxSpamFilter) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MailboxSpamFilter) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetAction

`func (o *MailboxSpamFilter) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MailboxSpamFilter) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MailboxSpamFilter) SetAction(v string)`

SetAction sets Action field to given value.


### GetForwardTo

`func (o *MailboxSpamFilter) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *MailboxSpamFilter) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *MailboxSpamFilter) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.


### GetWhiteList

`func (o *MailboxSpamFilter) GetWhiteList() []string`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *MailboxSpamFilter) GetWhiteListOk() (*[]string, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *MailboxSpamFilter) SetWhiteList(v []string)`

SetWhiteList sets WhiteList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


