# DnsRecordData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **float32** | Приоритет DNS-записи. | [optional] 
**Subdomain** | Pointer to **NullableString** | Поддомен. | [optional] 
**Value** | **string** | Значение DNS-записи. | 

## Methods

### NewDnsRecordData

`func NewDnsRecordData(value string, ) *DnsRecordData`

NewDnsRecordData instantiates a new DnsRecordData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordDataWithDefaults

`func NewDnsRecordDataWithDefaults() *DnsRecordData`

NewDnsRecordDataWithDefaults instantiates a new DnsRecordData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *DnsRecordData) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DnsRecordData) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DnsRecordData) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DnsRecordData) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSubdomain

`func (o *DnsRecordData) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *DnsRecordData) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *DnsRecordData) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *DnsRecordData) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### SetSubdomainNil

`func (o *DnsRecordData) SetSubdomainNil(b bool)`

 SetSubdomainNil sets the value for Subdomain to be an explicit nil

### UnsetSubdomain
`func (o *DnsRecordData) UnsetSubdomain()`

UnsetSubdomain ensures that no value is present for Subdomain, not even an explicit nil
### GetValue

`func (o *DnsRecordData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DnsRecordData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DnsRecordData) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


