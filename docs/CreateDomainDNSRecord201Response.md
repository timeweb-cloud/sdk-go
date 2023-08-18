# CreateDomainDNSRecord201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsRecord** | [**DnsRecord**](DnsRecord.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCreateDomainDNSRecord201Response

`func NewCreateDomainDNSRecord201Response(dnsRecord DnsRecord, responseId string, ) *CreateDomainDNSRecord201Response`

NewCreateDomainDNSRecord201Response instantiates a new CreateDomainDNSRecord201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainDNSRecord201ResponseWithDefaults

`func NewCreateDomainDNSRecord201ResponseWithDefaults() *CreateDomainDNSRecord201Response`

NewCreateDomainDNSRecord201ResponseWithDefaults instantiates a new CreateDomainDNSRecord201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsRecord

`func (o *CreateDomainDNSRecord201Response) GetDnsRecord() DnsRecord`

GetDnsRecord returns the DnsRecord field if non-nil, zero value otherwise.

### GetDnsRecordOk

`func (o *CreateDomainDNSRecord201Response) GetDnsRecordOk() (*DnsRecord, bool)`

GetDnsRecordOk returns a tuple with the DnsRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecord

`func (o *CreateDomainDNSRecord201Response) SetDnsRecord(v DnsRecord)`

SetDnsRecord sets DnsRecord field to given value.


### GetResponseId

`func (o *CreateDomainDNSRecord201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateDomainDNSRecord201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateDomainDNSRecord201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


