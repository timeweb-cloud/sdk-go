# CreateBalancerCertificates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Тип сертификата. | [optional] 
**Fqdn** | Pointer to **string** | Полное имя домена. | [optional] 
**CertData** | Pointer to **string** | Данные сертификата. Нужны только для типа custom. | [optional] 
**KeyData** | Pointer to **string** | Данные ключа. Нужны только для типа custom. | [optional] 

## Methods

### NewCreateBalancerCertificates

`func NewCreateBalancerCertificates() *CreateBalancerCertificates`

NewCreateBalancerCertificates instantiates a new CreateBalancerCertificates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBalancerCertificatesWithDefaults

`func NewCreateBalancerCertificatesWithDefaults() *CreateBalancerCertificates`

NewCreateBalancerCertificatesWithDefaults instantiates a new CreateBalancerCertificates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateBalancerCertificates) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateBalancerCertificates) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateBalancerCertificates) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateBalancerCertificates) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFqdn

`func (o *CreateBalancerCertificates) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *CreateBalancerCertificates) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *CreateBalancerCertificates) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *CreateBalancerCertificates) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetCertData

`func (o *CreateBalancerCertificates) GetCertData() string`

GetCertData returns the CertData field if non-nil, zero value otherwise.

### GetCertDataOk

`func (o *CreateBalancerCertificates) GetCertDataOk() (*string, bool)`

GetCertDataOk returns a tuple with the CertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertData

`func (o *CreateBalancerCertificates) SetCertData(v string)`

SetCertData sets CertData field to given value.

### HasCertData

`func (o *CreateBalancerCertificates) HasCertData() bool`

HasCertData returns a boolean if a field has been set.

### GetKeyData

`func (o *CreateBalancerCertificates) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *CreateBalancerCertificates) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *CreateBalancerCertificates) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *CreateBalancerCertificates) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


