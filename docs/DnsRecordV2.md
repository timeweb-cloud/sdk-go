# DnsRecordV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип DNS-записи. | 
**Id** | Pointer to **NullableFloat32** | ID DNS-записи. | [optional] 
**Fqdn** | Pointer to **string** | Полное имя основного домена. | [optional] 
**Data** | [**DnsRecordV2Data**](DnsRecordV2Data.md) |  | 
**Ttl** | Pointer to **float32** | Время жизни DNS-записи. | [optional] 

## Methods

### NewDnsRecordV2

`func NewDnsRecordV2(type_ string, data DnsRecordV2Data, ) *DnsRecordV2`

NewDnsRecordV2 instantiates a new DnsRecordV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordV2WithDefaults

`func NewDnsRecordV2WithDefaults() *DnsRecordV2`

NewDnsRecordV2WithDefaults instantiates a new DnsRecordV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DnsRecordV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsRecordV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsRecordV2) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *DnsRecordV2) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsRecordV2) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsRecordV2) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *DnsRecordV2) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DnsRecordV2) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DnsRecordV2) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetFqdn

`func (o *DnsRecordV2) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DnsRecordV2) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DnsRecordV2) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *DnsRecordV2) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetData

`func (o *DnsRecordV2) GetData() DnsRecordV2Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DnsRecordV2) GetDataOk() (*DnsRecordV2Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DnsRecordV2) SetData(v DnsRecordV2Data)`

SetData sets Data field to given value.


### GetTtl

`func (o *DnsRecordV2) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsRecordV2) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsRecordV2) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsRecordV2) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


