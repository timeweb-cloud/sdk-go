# AppV2DomainsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | **string** | Домен | 
**IsTechnical** | **bool** | Является ли домен техническим. | 
**IsExternal** | **bool** | Является ли домен внешним. | 
**DnsRecordId** | **NullableInt32** | ID DNS-записи; null, если запись отсутствует. | 

## Methods

### NewAppV2DomainsInner

`func NewAppV2DomainsInner(fqdn string, isTechnical bool, isExternal bool, dnsRecordId NullableInt32, ) *AppV2DomainsInner`

NewAppV2DomainsInner instantiates a new AppV2DomainsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppV2DomainsInnerWithDefaults

`func NewAppV2DomainsInnerWithDefaults() *AppV2DomainsInner`

NewAppV2DomainsInnerWithDefaults instantiates a new AppV2DomainsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *AppV2DomainsInner) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *AppV2DomainsInner) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *AppV2DomainsInner) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetIsTechnical

`func (o *AppV2DomainsInner) GetIsTechnical() bool`

GetIsTechnical returns the IsTechnical field if non-nil, zero value otherwise.

### GetIsTechnicalOk

`func (o *AppV2DomainsInner) GetIsTechnicalOk() (*bool, bool)`

GetIsTechnicalOk returns a tuple with the IsTechnical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTechnical

`func (o *AppV2DomainsInner) SetIsTechnical(v bool)`

SetIsTechnical sets IsTechnical field to given value.


### GetIsExternal

`func (o *AppV2DomainsInner) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *AppV2DomainsInner) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *AppV2DomainsInner) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.


### GetDnsRecordId

`func (o *AppV2DomainsInner) GetDnsRecordId() int32`

GetDnsRecordId returns the DnsRecordId field if non-nil, zero value otherwise.

### GetDnsRecordIdOk

`func (o *AppV2DomainsInner) GetDnsRecordIdOk() (*int32, bool)`

GetDnsRecordIdOk returns a tuple with the DnsRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecordId

`func (o *AppV2DomainsInner) SetDnsRecordId(v int32)`

SetDnsRecordId sets DnsRecordId field to given value.


### SetDnsRecordIdNil

`func (o *AppV2DomainsInner) SetDnsRecordIdNil(b bool)`

 SetDnsRecordIdNil sets the value for DnsRecordId to be an explicit nil

### UnsetDnsRecordId
`func (o *AppV2DomainsInner) UnsetDnsRecordId()`

UnsetDnsRecordId ensures that no value is present for DnsRecordId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


