# GetDomainDNSRecordsV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**DnsRecords** | [**[]DnsRecordV2**](DnsRecordV2.md) |  | 

## Methods

### NewGetDomainDNSRecordsV2200Response

`func NewGetDomainDNSRecordsV2200Response(meta Meta, dnsRecords []DnsRecordV2, ) *GetDomainDNSRecordsV2200Response`

NewGetDomainDNSRecordsV2200Response instantiates a new GetDomainDNSRecordsV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainDNSRecordsV2200ResponseWithDefaults

`func NewGetDomainDNSRecordsV2200ResponseWithDefaults() *GetDomainDNSRecordsV2200Response`

NewGetDomainDNSRecordsV2200ResponseWithDefaults instantiates a new GetDomainDNSRecordsV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDomainDNSRecordsV2200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDomainDNSRecordsV2200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDomainDNSRecordsV2200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDnsRecords

`func (o *GetDomainDNSRecordsV2200Response) GetDnsRecords() []DnsRecordV2`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *GetDomainDNSRecordsV2200Response) GetDnsRecordsOk() (*[]DnsRecordV2, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *GetDomainDNSRecordsV2200Response) SetDnsRecords(v []DnsRecordV2)`

SetDnsRecords sets DnsRecords field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


