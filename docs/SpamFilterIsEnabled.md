# SpamFilterIsEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | Включен ли спам-фильтр | [optional] 
**Action** | Pointer to **string** | Что делать с письмами, которые попадают в спам. \\  Параметры: \\  &#x60;move_to_directory&#x60; - переместить в паку спам; \\  &#x60;forward&#x60; - переслать письмо на другой адрес; \\  &#x60;delete&#x60; - удалить письмо; \\  &#x60;tag&#x60; - пометить письмо; \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 
**ForwardTo** | Pointer to **string** | Адрес для пересылки при выбранном действии &#x60;forward&#x60; из параметра &#x60;action&#x60;. Не может быть пустым, если &#x60;action&#x60; выбран &#x60;forward&#x60;. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 
**WhiteList** | Pointer to **[]string** | Белый список адресов от которых письма не будут попадать в спам. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewSpamFilterIsEnabled

`func NewSpamFilterIsEnabled() *SpamFilterIsEnabled`

NewSpamFilterIsEnabled instantiates a new SpamFilterIsEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpamFilterIsEnabledWithDefaults

`func NewSpamFilterIsEnabledWithDefaults() *SpamFilterIsEnabled`

NewSpamFilterIsEnabledWithDefaults instantiates a new SpamFilterIsEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *SpamFilterIsEnabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SpamFilterIsEnabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SpamFilterIsEnabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *SpamFilterIsEnabled) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetAction

`func (o *SpamFilterIsEnabled) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SpamFilterIsEnabled) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SpamFilterIsEnabled) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SpamFilterIsEnabled) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetForwardTo

`func (o *SpamFilterIsEnabled) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *SpamFilterIsEnabled) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *SpamFilterIsEnabled) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *SpamFilterIsEnabled) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### GetWhiteList

`func (o *SpamFilterIsEnabled) GetWhiteList() []string`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *SpamFilterIsEnabled) GetWhiteListOk() (*[]string, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *SpamFilterIsEnabled) SetWhiteList(v []string)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *SpamFilterIsEnabled) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


