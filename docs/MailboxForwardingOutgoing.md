# MailboxForwardingOutgoing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включена ли пересылка исходящих писем | 
**OutgoingTo** | **string** | Адрес для пересылки исходящих писем | 

## Methods

### NewMailboxForwardingOutgoing

`func NewMailboxForwardingOutgoing(isEnabled bool, outgoingTo string, ) *MailboxForwardingOutgoing`

NewMailboxForwardingOutgoing instantiates a new MailboxForwardingOutgoing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxForwardingOutgoingWithDefaults

`func NewMailboxForwardingOutgoingWithDefaults() *MailboxForwardingOutgoing`

NewMailboxForwardingOutgoingWithDefaults instantiates a new MailboxForwardingOutgoing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *MailboxForwardingOutgoing) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MailboxForwardingOutgoing) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MailboxForwardingOutgoing) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetOutgoingTo

`func (o *MailboxForwardingOutgoing) GetOutgoingTo() string`

GetOutgoingTo returns the OutgoingTo field if non-nil, zero value otherwise.

### GetOutgoingToOk

`func (o *MailboxForwardingOutgoing) GetOutgoingToOk() (*string, bool)`

GetOutgoingToOk returns a tuple with the OutgoingTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTo

`func (o *MailboxForwardingOutgoing) SetOutgoingTo(v string)`

SetOutgoingTo sets OutgoingTo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


