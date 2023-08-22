# GetDomainDNSRecords200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**DnsRecords** | [**[]DnsRecord**](DnsRecord.md) |  | 

## Methods

### NewGetDomainDNSRecords200Response

`func NewGetDomainDNSRecords200Response(meta Meta, dnsRecords []DnsRecord, ) *GetDomainDNSRecords200Response`

NewGetDomainDNSRecords200Response instantiates a new GetDomainDNSRecords200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainDNSRecords200ResponseWithDefaults

`func NewGetDomainDNSRecords200ResponseWithDefaults() *GetDomainDNSRecords200Response`

NewGetDomainDNSRecords200ResponseWithDefaults instantiates a new GetDomainDNSRecords200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDomainDNSRecords200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDomainDNSRecords200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDomainDNSRecords200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDnsRecords

`func (o *GetDomainDNSRecords200Response) GetDnsRecords() []DnsRecord`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *GetDomainDNSRecords200Response) GetDnsRecordsOk() (*[]DnsRecord, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *GetDomainDNSRecords200Response) SetDnsRecords(v []DnsRecord)`

SetDnsRecords sets DnsRecords field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


