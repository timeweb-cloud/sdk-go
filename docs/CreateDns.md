# CreateDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **float32** | Приоритет DNS-записи. | [optional] 
**Subdomain** | Pointer to **string** | Полное имя поддомена. | [optional] 
**Type** | **string** | Тип DNS-записи. | 
**Value** | **string** | Значение DNS-записи. | 
**Ttl** | Pointer to **float32** | Время жизни DNS-записи. | [optional] 

## Methods

### NewCreateDns

`func NewCreateDns(type_ string, value string, ) *CreateDns`

NewCreateDns instantiates a new CreateDns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDnsWithDefaults

`func NewCreateDnsWithDefaults() *CreateDns`

NewCreateDnsWithDefaults instantiates a new CreateDns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *CreateDns) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateDns) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateDns) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateDns) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSubdomain

`func (o *CreateDns) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *CreateDns) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *CreateDns) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *CreateDns) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetType

`func (o *CreateDns) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDns) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDns) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *CreateDns) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateDns) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateDns) SetValue(v string)`

SetValue sets Value field to given value.


### GetTtl

`func (o *CreateDns) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateDns) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateDns) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CreateDns) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


