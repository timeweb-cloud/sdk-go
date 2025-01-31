# DnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип DNS-записи. | 
**Id** | Pointer to **NullableFloat32** | ID DNS-записи. | [optional] 
**Data** | [**DnsRecordData**](DnsRecordData.md) |  | 
**Ttl** | Pointer to **float32** | Время жизни DNS-записи. | [optional] 

## Methods

### NewDnsRecord

`func NewDnsRecord(type_ string, data DnsRecordData, ) *DnsRecord`

NewDnsRecord instantiates a new DnsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordWithDefaults

`func NewDnsRecordWithDefaults() *DnsRecord`

NewDnsRecordWithDefaults instantiates a new DnsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DnsRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsRecord) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *DnsRecord) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsRecord) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsRecord) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *DnsRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DnsRecord) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DnsRecord) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetData

`func (o *DnsRecord) GetData() DnsRecordData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DnsRecord) GetDataOk() (*DnsRecordData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DnsRecord) SetData(v DnsRecordData)`

SetData sets Data field to given value.


### GetTtl

`func (o *DnsRecord) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsRecord) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsRecord) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


