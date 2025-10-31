# OutgoingIsEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включена ли пересылка исходящих писем | 
**OutgoingEmail** | **string** | Адрес для пересылки исходящих писем. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | 

## Methods

### NewOutgoingIsEnabled

`func NewOutgoingIsEnabled(isEnabled bool, outgoingEmail string, ) *OutgoingIsEnabled`

NewOutgoingIsEnabled instantiates a new OutgoingIsEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingIsEnabledWithDefaults

`func NewOutgoingIsEnabledWithDefaults() *OutgoingIsEnabled`

NewOutgoingIsEnabledWithDefaults instantiates a new OutgoingIsEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *OutgoingIsEnabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *OutgoingIsEnabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *OutgoingIsEnabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetOutgoingEmail

`func (o *OutgoingIsEnabled) GetOutgoingEmail() string`

GetOutgoingEmail returns the OutgoingEmail field if non-nil, zero value otherwise.

### GetOutgoingEmailOk

`func (o *OutgoingIsEnabled) GetOutgoingEmailOk() (*string, bool)`

GetOutgoingEmailOk returns a tuple with the OutgoingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingEmail

`func (o *OutgoingIsEnabled) SetOutgoingEmail(v string)`

SetOutgoingEmail sets OutgoingEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


