# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID сертификата. Указывается в поле &#x60;config.security.certificate_id&#x60; при привязке сертификата к ресурсу. | 
**Type** | **string** | Тип сертификата. - &#x60;lets_encrypt&#x60; — сертификат, выпущенный через Let&#39;s Encrypt; - &#x60;uploaded&#x60; — сертификат, загруженный вами. | 
**Cn** | **string** | Основное доменное имя сертификата. | 
**Domains** | **[]string** | Все доменные имена сертификата, включая указанные в SAN. | 
**IssuedAt** | **time.Time** | Дата выпуска сертификата. | 
**ExpiresAt** | **time.Time** | Дата окончания действия сертификата. | 

## Methods

### NewCertificate

`func NewCertificate(id int32, type_ string, cn string, domains []string, issuedAt time.Time, expiresAt time.Time, ) *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Certificate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Certificate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Certificate) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *Certificate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Certificate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Certificate) SetType(v string)`

SetType sets Type field to given value.


### GetCn

`func (o *Certificate) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *Certificate) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *Certificate) SetCn(v string)`

SetCn sets Cn field to given value.


### GetDomains

`func (o *Certificate) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Certificate) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *Certificate) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetIssuedAt

`func (o *Certificate) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *Certificate) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *Certificate) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.


### GetExpiresAt

`func (o *Certificate) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Certificate) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Certificate) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


