# Subdomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | **string** | Полное имя поддомена. | 
**Id** | **float32** | ID поддомена. | 
**LinkedIp** | **NullableString** | Привязанный к поддомену IP-адрес. | 

## Methods

### NewSubdomain

`func NewSubdomain(fqdn string, id float32, linkedIp NullableString, ) *Subdomain`

NewSubdomain instantiates a new Subdomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubdomainWithDefaults

`func NewSubdomainWithDefaults() *Subdomain`

NewSubdomainWithDefaults instantiates a new Subdomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *Subdomain) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Subdomain) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Subdomain) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetId

`func (o *Subdomain) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subdomain) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subdomain) SetId(v float32)`

SetId sets Id field to given value.


### GetLinkedIp

`func (o *Subdomain) GetLinkedIp() string`

GetLinkedIp returns the LinkedIp field if non-nil, zero value otherwise.

### GetLinkedIpOk

`func (o *Subdomain) GetLinkedIpOk() (*string, bool)`

GetLinkedIpOk returns a tuple with the LinkedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIp

`func (o *Subdomain) SetLinkedIp(v string)`

SetLinkedIp sets LinkedIp field to given value.


### SetLinkedIpNil

`func (o *Subdomain) SetLinkedIpNil(b bool)`

 SetLinkedIpNil sets the value for LinkedIp to be an explicit nil

### UnsetLinkedIp
`func (o *Subdomain) UnsetLinkedIp()`

UnsetLinkedIp ensures that no value is present for LinkedIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


