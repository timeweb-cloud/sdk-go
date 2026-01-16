# DnsRecordV2Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **float32** | Приоритет DNS-записи (для MX и SRV записей). | [optional] 
**Subdomain** | Pointer to **NullableString** | Имя поддомена (только поддомен без основного домена, например &#x60;sub&#x60; для &#x60;sub.example.com&#x60;). Для записей на основном домене это поле отсутствует в ответе. | [optional] 
**Value** | Pointer to **string** | Значение DNS-записи (для A, AAAA, TXT, CNAME, MX записей). | [optional] 
**Host** | Pointer to **string** | Каноническое имя хоста, предоставляющего сервис (для SRV записей). | [optional] 
**Port** | Pointer to **float32** | TCP или UDP порт, на котором работает сервис (для SRV записей). | [optional] 
**Service** | Pointer to **string** | Имя сервиса (для SRV записей). | [optional] 
**Protocol** | Pointer to **string** | Протокол (для SRV записей). | [optional] 

## Methods

### NewDnsRecordV2Data

`func NewDnsRecordV2Data() *DnsRecordV2Data`

NewDnsRecordV2Data instantiates a new DnsRecordV2Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordV2DataWithDefaults

`func NewDnsRecordV2DataWithDefaults() *DnsRecordV2Data`

NewDnsRecordV2DataWithDefaults instantiates a new DnsRecordV2Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *DnsRecordV2Data) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DnsRecordV2Data) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DnsRecordV2Data) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DnsRecordV2Data) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSubdomain

`func (o *DnsRecordV2Data) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *DnsRecordV2Data) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *DnsRecordV2Data) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *DnsRecordV2Data) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### SetSubdomainNil

`func (o *DnsRecordV2Data) SetSubdomainNil(b bool)`

 SetSubdomainNil sets the value for Subdomain to be an explicit nil

### UnsetSubdomain
`func (o *DnsRecordV2Data) UnsetSubdomain()`

UnsetSubdomain ensures that no value is present for Subdomain, not even an explicit nil
### GetValue

`func (o *DnsRecordV2Data) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DnsRecordV2Data) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DnsRecordV2Data) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DnsRecordV2Data) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetHost

`func (o *DnsRecordV2Data) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DnsRecordV2Data) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DnsRecordV2Data) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DnsRecordV2Data) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *DnsRecordV2Data) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DnsRecordV2Data) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DnsRecordV2Data) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DnsRecordV2Data) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetService

`func (o *DnsRecordV2Data) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DnsRecordV2Data) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DnsRecordV2Data) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DnsRecordV2Data) HasService() bool`

HasService returns a boolean if a field has been set.

### GetProtocol

`func (o *DnsRecordV2Data) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DnsRecordV2Data) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DnsRecordV2Data) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *DnsRecordV2Data) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


