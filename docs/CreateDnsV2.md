# CreateDnsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип DNS-записи. | 
**Value** | **string** | IPv4 адрес. | 
**Ttl** | Pointer to **float32** | Время жизни DNS-записи в секундах. | [optional] 

## Methods

### NewCreateDnsV2

`func NewCreateDnsV2(type_ string, value string, ) *CreateDnsV2`

NewCreateDnsV2 instantiates a new CreateDnsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDnsV2WithDefaults

`func NewCreateDnsV2WithDefaults() *CreateDnsV2`

NewCreateDnsV2WithDefaults instantiates a new CreateDnsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateDnsV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDnsV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDnsV2) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *CreateDnsV2) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateDnsV2) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateDnsV2) SetValue(v string)`

SetValue sets Value field to given value.


### GetTtl

`func (o *CreateDnsV2) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateDnsV2) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateDnsV2) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CreateDnsV2) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


