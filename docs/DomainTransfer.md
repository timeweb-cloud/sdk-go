# DomainTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Тип создаваемой заявки. | 
**AuthCode** | **string** | Код авторизации для переноса домена. | 
**Fqdn** | **string** | Полное имя домена. | 

## Methods

### NewDomainTransfer

`func NewDomainTransfer(action string, authCode string, fqdn string, ) *DomainTransfer`

NewDomainTransfer instantiates a new DomainTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainTransferWithDefaults

`func NewDomainTransferWithDefaults() *DomainTransfer`

NewDomainTransferWithDefaults instantiates a new DomainTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DomainTransfer) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DomainTransfer) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DomainTransfer) SetAction(v string)`

SetAction sets Action field to given value.


### GetAuthCode

`func (o *DomainTransfer) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *DomainTransfer) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *DomainTransfer) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetFqdn

`func (o *DomainTransfer) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DomainTransfer) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DomainTransfer) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


