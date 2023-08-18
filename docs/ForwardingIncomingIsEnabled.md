# ForwardingIncomingIsEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | Включена ли пересылка входящик писем | [optional] 
**IsDeleteMessages** | Pointer to **bool** | Удалять ли входящие письма. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 
**IncomingList** | Pointer to **[]string** | Список адресов для пересылки. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewForwardingIncomingIsEnabled

`func NewForwardingIncomingIsEnabled() *ForwardingIncomingIsEnabled`

NewForwardingIncomingIsEnabled instantiates a new ForwardingIncomingIsEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardingIncomingIsEnabledWithDefaults

`func NewForwardingIncomingIsEnabledWithDefaults() *ForwardingIncomingIsEnabled`

NewForwardingIncomingIsEnabledWithDefaults instantiates a new ForwardingIncomingIsEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ForwardingIncomingIsEnabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ForwardingIncomingIsEnabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ForwardingIncomingIsEnabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ForwardingIncomingIsEnabled) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDeleteMessages

`func (o *ForwardingIncomingIsEnabled) GetIsDeleteMessages() bool`

GetIsDeleteMessages returns the IsDeleteMessages field if non-nil, zero value otherwise.

### GetIsDeleteMessagesOk

`func (o *ForwardingIncomingIsEnabled) GetIsDeleteMessagesOk() (*bool, bool)`

GetIsDeleteMessagesOk returns a tuple with the IsDeleteMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteMessages

`func (o *ForwardingIncomingIsEnabled) SetIsDeleteMessages(v bool)`

SetIsDeleteMessages sets IsDeleteMessages field to given value.

### HasIsDeleteMessages

`func (o *ForwardingIncomingIsEnabled) HasIsDeleteMessages() bool`

HasIsDeleteMessages returns a boolean if a field has been set.

### GetIncomingList

`func (o *ForwardingIncomingIsEnabled) GetIncomingList() []string`

GetIncomingList returns the IncomingList field if non-nil, zero value otherwise.

### GetIncomingListOk

`func (o *ForwardingIncomingIsEnabled) GetIncomingListOk() (*[]string, bool)`

GetIncomingListOk returns a tuple with the IncomingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingList

`func (o *ForwardingIncomingIsEnabled) SetIncomingList(v []string)`

SetIncomingList sets IncomingList field to given value.

### HasIncomingList

`func (o *ForwardingIncomingIsEnabled) HasIncomingList() bool`

HasIncomingList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


